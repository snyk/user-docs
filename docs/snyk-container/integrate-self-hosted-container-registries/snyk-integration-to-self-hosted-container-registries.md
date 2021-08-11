# Snyk integration to self-hosted container registries

**Feature availability**  
This feature is currently in beta and is available with Enterprise plans. See [Pricing plans](https://snyk.io/plans/) for more details.

Snyk can integrate to private container registries you host \(currently in open beta\), and help you to better secure container images in those registries.   


To enable and configure your hosted container registry, contact our support team at [support@snyk.io](mailto:support@snyk.io)

**Introduction**

Integration with private container registries allows you to:

* Keep sensitive data such as your access tokens inside your private network, never sharing that information with Snyk.
* Provide controlled access to the network by Snyk, limiting Snyk access and the actions that Snyk can perform.

**Self-hosted container registries solution components**

* Broker server: running on Snyk SaaS backend
* Broker client & Container Registry Agent: Two Docker images deployed in your infrastructure, creating two separate services, responsible for sampling your container registries in a secured manner and sending the allowed information to Snyk.

The Broker client provides the Agent with the connection details. The Agent uses these details to connect to the container registry, scan the images, and send the scan results through the brokered communication using callbacks. The brokered communication happens when a Broker Client connects \(using your Broker ID\) to a Broker server which runs in Snyk environment. See [Snyk Broker](https://support.snyk.io/hc/en-us/sections/360001138138-Snyk-Broker) documentation for more details.  

![mceclip0.png](https://support.snyk.io/hc/article_attachments/360016081258/mceclip0.png)

**Supported Container registries**

* Artifactory \(type: ArtifactoryCR\)
* Harbor \(type: HarborCR\)
* Azure \(type: AzureCR\)
* GCR \(type: GoogleCR\)
* Docker Hub \(type: DockerHub\)
* Quay \(type: QuayCR\)
* Nexus \(type: nexus-cr\)
* GitHub \(type: github-cr\)
* DigitalOcean \(type: digitalocean-cr\)
* GitLab \(type: gitlab-cr\)

### **Note**

The integration pattern using broker with open source container registries from the above list is designed for users who require images to be scanned in their own environment, instead of inside the Snyk service.

If such a requirement is not relevant for you, you do not need the architecture described in this article, and can integrate to it in the standard way from the integrations page.

**Settings prerequisites**

* * Broker Client machine system requirements: 1 CPU, 256MB of Ram.
  * Container registry agent machine system requirements should be \(given MAX\_ACTIVE\_OPERATIONS=1\): 
    * CPU: 1 vcpu
    * Memory: 2Gb \(should be reflected in node memory setting\)
    * Storage: 5Gb
  * Docker configured to pull components images from Docker Hub
  * Connection between broker and agent
  * Agent image can be found for download [here](https://hub.docker.com/r/snyk/container-registry-agent/tags?page=1&ordering=last_updated)
  * Broker Client image can be found for download [here](https://hub.docker.com/r/snyk/broker/tags?page=1&ordering=last_updated&name=container-registry-agent)

**Set up the remote connection**

To use the Broker client with a container registry agent deployment, run  
\`docker pull snyk/broker:container-registry-agent\`. 

The following environment variables are mandatory to configure the Broker client:

* \`BROKER\_TOKEN\` - The Snyk Broker token, obtained from your Container registry integration \(provided by Snyk support\)
* \`BROKER\_CLIENT\_URL\` - The URL of your broker client \(including scheme and - port\) used by the container registry agent to call back to Snyk, for example: "http://my.broker.client:8000".
* \`CR\_AGENT\_URL\` - The URL of your container registry agent, for example "http://my.container-registry-agent".
* \`CR\_TYPE\` - The container registry type as listed in supporter registries, for example "DockerHub", "GoogleCR", "ArtifactoryCR".
* \`CR\_BASE\` - The hostname of the container registry api to connect to, for example: "cr.host.com".
* \`CR\_USERNAME\` - The username for authenticating to container registry api. Not used for DigitalOcean container registry.
* \`CR\_PASSWORD\` - The password for authenticating to container registry api. Not used for DigitalOcean container registry.
* \`CR\_TOKEN\` - Authentication token for DigitalOcean container registry.
* \`PORT\` - The local port at which the Broker client accepts connections. Default is 7341.

**Note for Artifactory users**

In case you are using **Repository path** as your Docker access method, the container registry hostname in CR\_BASE variable should be set in this structure:_/artifactory/api/docker/_

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

**Configuring and using system check:**

You can use the \`/systemcheck\` endpoint to verify connectivity between the Broker Client and the Container Registry Agent.

In order to use it, specify the following environment variable when running the broker client:  
BROKER\_CLIENT\_VALIDATION\_URL = "&lt;agent-url&gt;/healthcheck"

**Secure your images:**

You can now start scanning your container images directly from your private registry. See [scanning images from container registry](https://support.snyk.io/hc/en-us/articles/360003915998-Configuring-your-JFrog-Artifactory-container-registry-integration) \(Artifactory example\) for more details.

