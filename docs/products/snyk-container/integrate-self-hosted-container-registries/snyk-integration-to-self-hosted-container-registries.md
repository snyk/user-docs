# Snyk integration to self-hosted container registries

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can integrate to private container registries you host (currently in open beta), and help you to better secure container images in those registries.

{% hint style="warning" %}
For this feature to work, you need two separate containers deployed in your infrastructure, creating two separate services.
{% endhint %}

To enable and configure your hosted container registry, contact our support team at [support@snyk.io](mailto:support@snyk.io)

## **Introduction**

Integration with private container registries allows you to:

* Keep sensitive data such as your access tokens inside your private network, never sharing that information with Snyk.
* Provide controlled access to the network by Snyk, limiting Snyk access and the actions that Snyk can perform.

## **Self-hosted container registries solution components**

* Broker server: running on Snyk SaaS backend
* Broker client & Container Registry Agent: Two Docker images deployed in your infrastructure, creating two separate services, responsible for sampling your container registries in a secured manner and sending the allowed information to Snyk.

The Broker client provides the Agent with the connection details. The Agent uses these details to connect to the container registry, scan the images, and send the scan results through the brokered communication using callbacks. The brokered communication happens when a Broker Client connects (using your Broker ID) to a Broker server which runs in Snyk environment. See [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker) documentation for more details.

![](../../../.gitbook/assets/mceclip0-8-.png)

**Supported Container registries**

* Artifactory (type: artifactory-cr)
* Harbor (type: harbor-cr)
* Azure (type: acr)
* GCR (type: gcr)
* Docker Hub (type: docker-hub)
* Quay (type: quay-cr)
* Nexus (type: nexus-cr)
* GitHub (type: github-cr)
* DigitalOcean (type: digitalocean-cr)
* GitLab (type: gitlab-cr)

{% hint style="info" %}
**Note:**\
The integration pattern using broker with open source container registries from the above list is designed for users who require images to be scanned in their own environment, instead of inside the Snyk service**.**\
If such a requirement is not relevant for you, you do not need the architecture described in this article, and can integrate to it in the standard way from the integrations page.
{% endhint %}

**Settings prerequisites**

* Broker Client machine system requirements: 1 CPU, 256MB of RAM.
* Container registry agent machine system requirements should be (given MAX_ACTIVE_OPERATIONS=1): 
  * CPU: 1 vcpu
  * Memory: 2Gb (should be reflected in node memory setting)
  * Storage: 5Gb
* Docker configured to pull components images from Docker Hub
* Connection between broker and agent
* Agent image can be found for download [here](https://hub.docker.com/r/snyk/container-registry-agent/tags?page=1\&ordering=last_updated)
* Broker Client image can be found for download [here](https://hub.docker.com/r/snyk/broker/tags?page=1\&ordering=last_updated\&name=container-registry-agent)

{% hint style="info" %}
### Scaling to adjust scan capacity

With the above configuration of 1 vCPU and 2GB RAM, scanning capacity would be approximately 160 images of \~350MB in one go. You can scale this up based on your image sizes, and in case you have a specific use case that doesn't allow scaling and doesn't fit the limitations, please contact our support team.
{% endhint %}

## **Set up the remote connection**

To use the Broker client with a container registry agent deployment, run\
`docker pull snyk/broker:container-registry-agent`.

There are different environment variable required to configure the Broker client, which differ between some of the container registries.

The following environment variables are mandatory to configure the Broker client:

**For DigitalOcean, GCR, Google Artifact Registry and Artifactory, there are a few values to notice. For ECR additional setup is required. Specifications to follow.**

* `BROKER_TOKEN` - The Snyk Broker token, obtained from your Container registry integration (provided by Snyk support)
* `BROKER_CLIENT_URL` - The URL of your broker client (including scheme and - port) used by the container registry agent to call back to Snyk, for example: "[http://my.broker.client:8000](http://my.broker.client:8000)".
* `CR_AGENT_URL` - The URL of your container registry agent, for example "[http://my.container-registry-agent](http://my.container-registry-agent)".
* `CR_TYPE` - The container registry type as listed in supporter registries, for example "DockerHub", "GoogleCR", "ArtifactoryCR".
* `CR_BASE` - The hostname of the container registry api to connect to, for example: "cr.host.com".
* `CR_USERNAME` - The username for authenticating to container registry api. Not used for DigitalOcean container registry.
* `CR_PASSWORD` - The password for authenticating to container registry api. Not used for DigitalOcean container registry.
* `CR_TOKEN` - Authentication token for DigitalOcean container registry.
* `PORT` - The local port at which the Broker client accepts connections. Default is 7341.

### **DigitalOcean**

To set up Broker Client for **DigitalOcean**,`CR_USERNAME` and `CR_PASSWORD` are not required. Instead, you need to specify `CR_TOKEN` - authentication token for DigitalOcean container registry.

### **GCR and Google Artifact Registry**

To set up the Broker Client for those container registries, all the above applies. The only thing to note is that `CR_USERNAME` value is permanent and should be `_json_key`, and the `CR_PASSWORD` value should be the JSON key used to authenticate to google.

### **Artifactory**

{% hint style="info" %}
**NOTE for Artifactory users:**

In case you are using **Repository path** as your Docker access method, the container registry hostname in CR_BASE variable should be set in this structure:** **\<your artifactory host>_/artifactory/api/docker/\<artifactory-repo-name>_
{% endhint %}

### **ECR**

![A high-level architecture of the brokered ECR integration](<../../../.gitbook/assets/untitled (1).png>)

#### Required AWS Resource

ECR setup requires two kinds of IAM resources to be created:

* Container Registry Agent IAM Role / IAM User: an IAM Role / IAM User the Container Registry Agent uses to assume a cross-account role with access to ECR. It should have the following permissions:  `"sts:AssumeRole"`
*   Snyk ECR Service Role: an IAM Role with access to ECR which is assumed by the Container Registry Agent IAM Role / IAM User to gain read-only access to ECR.ddd\
    It should have the following permissions:

    ```
    [
      "ecr:GetLifecyclePolicyPreview",
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:DescribeImages",
      "ecr:GetAuthorizationToken",
      "ecr:DescribeRepositories",
      "ecr:ListTagsForResource",
      "ecr:ListImages",
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetRepositoryPolicy",
      "ecr:GetLifecyclePolicy"
    ]
    ```

#### **Setup steps for ECR**

The above resources can be used as follows, so that a single Container Registry Agent instance can access ECR repositories located in different accounts**:**

1. **(Run this step once only)** Create the Container Registry Agent IAM Role / IAM User and use it to run the Container Registry Agent. The IAM Role / IAM User could be provided to the Container Registry Agent using one of the methods described [here](https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/setting-credentials-node.html).  **Run the following steps for each of your ECR accounts, using a separate broker instance for each ECR account**
2. In the AWS account where your ECR reside, create the Snyk ECR Service Role with read access to your ECR, and edit the trust relationship to allow this role to be assumed only by the specific Container Registry Agent IAM Role / IAM User created in the previous step. 
3. Restrict the Container Registry Agent IAM Role / IAM User to only be allowed to assume the your Snyk ECR Service Role(s).
4. Provide the Broker Client with the Role ARN of the Snyk ECR Service Role together with the ECR region. The Broker Client passes this Role ARN to the Container Registry Agent, and the Container Registry Agent will assume it, to access your ECR. The following environment variables are needed:
   * CR_ROLE_ARN=\<the role ARN of SnykEcrServiceRole> 
   * CR_REGION=\<AWS Region of ECR>
   * CR_EXTERNAL_ID=\<Optional. An external ID found in the trust relationship condition>

For detailed information about the brokered ECR setup, click [here](setting-up-the-container-registry-agent-for-a-brokered-ecr-integration.md).

### Run**ning the images**

To run the docker container, provide the relevant configuration:

```
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

```
docker run --restart=always \
       -p 8081:8081 \
       -e SNYK_PORT=8081 \
       snyk/container-registry-agent:latest
```

## **Configuring and using system check:**

You can use the \`/systemcheck\` endpoint to verify connectivity between the Broker Client, the Container Registry Agent, and your container registry.

In order to use it, specify the following environment variable when running the broker client:\
BROKER_CLIENT_VALIDATION_URL = "\<agent-url>/systemcheck

{% hint style="info" %}
**Note:**\
The /systemcheck endpoint is not mandatory for the brokered integration to function. More information can be found here: [https://github.com/snyk/broker#systemcheck](https://github.com/snyk/broker#systemcheck)
{% endhint %}

## **Debugging Methods**

The `LOG_LEVEL` environment variable can be set to the desired level (debug/info/warn/error), in order to change the level of the Container Registry Agent and Broker Client logs.

For more verbose debugging, the Container Registry Agent can be run with the `DEBUG=*` environment variable. This allows printing the logs of the Node [Debug](https://www.npmjs.com/package/debug) package. The Debug package is used by several packages in the Container Registry Agent, among them is the [Needle](https://www.npmjs.com/package/needle) package, which is used for making HTTP requests. Printing debug logs specifically from Needle can be done by setting `DEBUG=needle`.

{% hint style="danger" %}
**Warning:**\
Using the debugging options of third-party tools is not recommended for production environments, as it may result in logging sensitive information in logs that are not maintained by Snyk. For example, header information of HTTP requests.
{% endhint %}

**Secure your images:**

You can now start scanning your container images directly from your private registry. See [scanning images from container registry](https://docs.snyk.io/snyk-container/jfrog-artifactory-image-scanning/configuring-your-jfrog-artifactory-container-registry-integration) (Artifactory example) for more details.
