# Snyk integration to self-hosted container registries

{% hint style="info" %}
**Feature availability**  
This feature is currently in beta and is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can integrate to private container registries you host \(currently in open beta\), and help you to better secure container images in those registries.

To enable and configure your hosted container registry, contact our support team at [support@snyk.io](mailto:support@snyk.io/)

## **Introduction**

Integration with private container registries allows you to:

* Keep sensitive data such as your access tokens inside your private network, never sharing that information with Snyk.
* Provide controlled access to the network by Snyk, limiting Snyk access and the actions that Snyk can perform.

## **Self-hosted container registries solution components**

* Broker server: running on Snyk SaaS backend
* Broker client & Container Registry Agent: Two Docker images deployed in your infrastructure, creating two separate services, responsible for sampling your container registries in a secured manner and sending the allowed information to Snyk.

The Broker client provides the Agent with the connection details. The Agent uses these details to connect to the container registry, scan the images, and send the scan results through the brokered communication using callbacks. The brokered communication happens when a Broker Client connects \(using your Broker ID\) to a Broker server which runs in Snyk environment. See [Snyk Broker](integrations/snyk-broker/) documentation for more details.

![](../../.gitbook/assets/mceclip0-8-.png)

**Supported Container registries**

* Artifactory \(type: ArtifactoryCR\/)
* Harbor \(type: HarborCR\/)
* Azure \(type: AzureCR\/)
* GCR \(type: GoogleCR\/)
* Docker Hub \(type: DockerHub\/)
* Quay \(type: QuayCR\/)
* Nexus \(type: nexus-cr\/)
* GitHub \(type: github-cr\/)
* DigitalOcean \(type: digitalocean-cr\/)
* GitLab \(type: gitlab-cr\/)

{% hint style="info" %}
**Note:**  
The integration pattern using broker with open source container registries from the above list is designed for users who require images to be scanned in their own environment, instead of inside the Snyk service**.**
{% endhint %}

If such a requirement is not relevant for you, you do not need the architecture described in this article, and can integrate to it in the standard way from the integrations page.

**Settings prerequisites**

* Broker Client machine system requirements: 1 CPU, 256MB of Ram.
* Container registry agent machine system requirements should be \(given MAX\_ACTIVE\_OPERATIONS=1\): 
  * CPU: 1 vcpu
  * Memory: 2Gb \(should be reflected in node memory setting\/)
  * Storage: 5Gb
  * 
* Docker configured to pull components images from Docker Hub
* Connection between broker and agent
* Agent image can be found for download [here](https://hub.docker.com/r/snyk/container-registry-agent/tags?page=1&ordering=last_updated/)
* Broker Client image can be found for download [here](https://hub.docker.com/r/snyk/broker/tags?page=1&ordering=last_updated&name=container-registry-agent/)

{% hint style="info" %}
### Scaling to adjust scan capacity

With the above configuration of 1 vCPU and 2GB RAM, scanning capacity would be approximately 160 images of ~350MB in one go. You can scale this up based on your image sizes, and in case you have a specific use case that doesn't allow scaling and doesn't fit the limitations, please contact our support team.
{% endhint %}

## **Set up the remote connection**

To use the Broker client with a container registry agent deployment, run  
\`docker pull snyk/broker:container-registry-agent\`.

There are different environment variable required to configure the Broker client, which differ between some of the container registries.

The following environment variables are mandatory to configure the Broker client:

**For DigitalOcean, GCR and Google Artifact Registry, there are a few values to notice. Specifications to follow.**

* \`BROKER\_TOKEN\` - The Snyk Broker token, obtained from your Container registry integration \(provided by Snyk support\/)
* \`BROKER\_CLIENT\_URL\` - The URL of your broker client \(including scheme and - port\) used by the container registry agent to call back to Snyk, for example: "[http://my.broker.client:8000](http://my.broker.client:8000)".
* \`CR\_AGENT\_URL\` - The URL of your container registry agent, for example "[http://my.container-registry-agent](http://my.container-registry-agent)".
* \`CR\_TYPE\` - The container registry type as listed in supporter registries, for example "DockerHub", "GoogleCR", "ArtifactoryCR".
* \`CR\_BASE\` - The hostname of the container registry api to connect to, for example: "cr.host.com".
* \`CR\_USERNAME\` - The username for authenticating to container registry api. Not used for DigitalOcean container registry.
* \`CR\_PASSWORD\` - The password for authenticating to container registry api. Not used for DigitalOcean container registry.
* \`CR\_TOKEN\` - Authentication token for DigitalOcean container registry.
* \`PORT\` - The local port at which the Broker client accepts connections. Default is 7341.

### **DigitalOcean**

To set up Broker Client for **DigitalOcean**,`CR_USERNAME` and `CR_PASSWORD` are not required. Instead, you need to specify `CR_TOKEN` - authentication token for DigitalOcean container registry.

### **GCR and Google Artifact Registry**

To set up the Broker Client for those container registries, all the above applies. The only thing to note is that `CR_USERNAME` value is permanent and should be `_json_key`, and the `CR_PASSWORD` value should be the JSON key used to authenticate to google.

{% hint style="info" %}
**NOTE for Artifactory users:**  
In case you are using **Repository path** as your Docker access method, the container registry hostname in CR\_BASE variable should be set in this structure:_/artifactory/api/docker/_
{% endhint %}

To run the docker container, provide the relevant configuration:

```text
docker run --restart=always \
       -p 8000:8000 \
       -e BROKER_TOKEN="" \
       -e BROKER_CLIENT_URL="" \
       -e CR_AGENT_URL="" \
       -e CR_TYPE="" \
       -e CR_BASE="" \
       -e CR_USERNAME="" \
       -e CR_PASSWORD="" \
       -e PORT=8000 \
       snyk/broker:container-registry-agent
```

The container registry agent can be pulled from Docker Hub using the link provided above in the settings prerequisites. To run the image you can use a single env variable for specifying the port:

```text
docker run --restart=always \
       -p 8081:8081 \
       -e SNYK_PORT=8081 \
       snyk/container-registry-agent:latest
```

## **Configuring and using system check:**

You can use the \`/systemcheck\` endpoint to verify connectivity between the Broker Client and the Container Registry Agent.

In order to use it, specify the following environment variable when running the broker client:  
BROKER\_CLIENT\_VALIDATION\_URL = "&lt;agent-url&gt;/healthcheck"

{% hint style="info" %}
**Note:**  
The /systemcheck endpoint is not mandatory for the brokered integration to function. More information can be found here: [https://github.com/snyk/broker\#systemcheck](https://github.com/snyk/broker#systemcheck/)
{% endhint %}

**Secure your images:**

You can now start scanning your container images directly from your private registry. See [scanning images from container registry](snyk-container/jfrog-artifactory-image-scanning/configuring-your-jfrog-artifactory-container-registry-integration/) \(Artifactory example\) for more details.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

