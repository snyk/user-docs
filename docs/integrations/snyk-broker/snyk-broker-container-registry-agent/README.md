# Snyk Broker - Container Registry Agent

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can integrate with private container registries you host, and help you to better secure container images in those registries.

{% hint style="warning" %}
For this feature to work, you must have two separate containers deployed in your infrastructure, creating two separate services.
{% endhint %}

To enable and configure your hosted container registry, contact Snyk support at [support@snyk.io](mailto:support@snyk.io).

{% hint style="info" %}
The integration pattern using the Broker with open-source container registries from the list described on this page is designed for users who require images to be scanned in their own environment, instead of inside the Snyk service.\
If you do not have this requirement, you do not need the architecture described here, and you can integrate in the standard way from the Integrations page on the Web UI.
{% endhint %}

## **Introduction to Container Registry Agent**

Integration with private container registries allows you to:

* Keep sensitive data such as your access tokens inside your private network, never sharing that information with Snyk.
* Provide controlled access to the network for Snyk, limiting Snyk access and the actions that Snyk can perform.

### **Self-hosted container registries solution components**

The following components are needed with self-hosted container registries:

* Broker Server: running on the Snyk SaaS backend.
* Broker Client and Container Registry Agent: two Docker images deployed in your infrastructure, creating two separate services, responsible for sampling your container registries in a secured manner and sending the allowed information to Snyk

The Broker Client provides the Container Registry Agent with the connection details. The Agent uses these details to connect to the container registry, scan the images, and send the scan results through the brokered communication using callbacks. The brokered communication happens when a Broker Client connects, using your Broker ID, to a Broker Server which runs in the Snyk environment. See the [Snyk Broker](../) introductory information for more details.

<figure><img src="../../../.gitbook/assets/mceclip0-8-.png" alt="Flow for Snyk Broker Container Registry Agent"><figcaption><p>Flow for Snyk Broker Container Registry Agent</p></figcaption></figure>

### **Supported container registries**

* Artifactory (type: artifactory-cr)
* Harbor (type: harbor-cr)
* Azure (type: acr)
* GCR (type: gcr)
* ECR (type: ecr)
* Google Artifact Registry (type: google-artifact-cr)
* Docker Hub (type: docker-hub)
* Quay (type: quay-cr)
* Nexus (type: nexus-cr)
* GitHub (type: github-cr)
* DigitalOcean (type: digitalocean-cr)
* GitLab (type: gitlab-cr)

### **Settings prerequisites for Container Registry Agent**

* Broker Client machine system requirements: 1 CPU, 256MB of RAM
* Container registry agent machine system requirements should be (given MAX\_ACTIVE\_OPERATIONS=1):
  * CPU: 1 vcpu
  * Memory: 2Gb (should be reflected in node memory setting)
  * Storage: 5Gb
* Docker configured to pull components images from Docker Hub
* Connection between Broker and Agent
* HTTPS connection between the Agent and the registry. Support for HTTP-only registries can be resolved by deploying a reverse proxy between the Agent and the SCM.
* [Download for the Broker Client image on Docker](https://hub.docker.com/r/snyk/broker/tags?page=1\&ordering=last\_updated\&name=container-registry-agent)
* [Download for the Container Registry Agent image on Docker](https://hub.docker.com/r/snyk/container-registry-agent/tags?page=1\&ordering=last\_updated)

{% hint style="info" %}
**Scaling to adjust scan capacity**

With the listed configuration of 1 vCPU and 2GB RAM, scanning capacity would be approximately 160 images of \~350MB in one run. You can scale this up based on your image sizes, and if you have a specific use case that does not allow scaling and does not fit the limitations, contact [Snyk support](https://support.snyk.io/hc/en-us/).
{% endhint %}

## **Set up the remote connection**

### **Running the Broker Client**

You can pull the Broker Client image from Docker Hub using the link listed in the [Settings prerequisites](./#settings-prerequisites-for-container-registry-agent).

There are environment variables required to configure the Broker Client.

{% hint style="info" %}
For **DigitalOcean**, **GCR**, **Google Artifact Registry**, and **Artifactory**, there are a few values to note. For **ECR**, additional setup is required. [Specifications](./#container-registry-specific-configurations) are provided.
{% endhint %}

* `BROKER_TOKEN` - The Snyk Broker token, obtained from your Container registry integration provided by Snyk support.
* `BROKER_CLIENT_URL` - The URL of your Broker Client, including scheme and port, which is used by the container registry agent to call back to Snyk through the brokered connection, for example: "[http://my.broker.client:8000](http://my.broker.client:8000)".
* `CR_AGENT_URL` - The URL of your Container Registry Agent, to which the Broker Client will route the requests, for example: "[http://my.container-registry-agent](http://my.container-registry-agent)".
* `CR_TYPE` - The container registry type as listed in supported registries, for example, "docker-hub", "gcr", "artifactory-cr".
* `CR_BASE` - The hostname of the container registry api to connect to, for example: "cr.host.com".
* `CR_USERNAME` - The username for authenticating to the container registry api.
* `CR_PASSWORD` - The password for authenticating to the container registry api.
* `CR_TOKEN` - Authentication token for DigitalOcean container registry.
* `PORT` - The local port at which the Broker client accepts connections. Default is 7341.

Run the Broker Client container with relevant configuration:

```
docker run --restart=always \
       -p 8000:8000 \
       -e BROKER_TOKEN="<secret-broker-token>" \
       -e BROKER_CLIENT_URL="<broker-client-url>" \
       -e CR_AGENT_URL="<container-registry-agent-url>" \
       -e CR_TYPE="<container-registry-type>" \
       -e CR_BASE="<container-registry-hostname>" \
       -e CR_USERNAME="<username>" \
       -e CR_PASSWORD="<password>" \
       -e PORT=8000 \
       snyk/broker:container-registry-agent
```

### Run**ning the Container Registry Agent**

You can pull the Container Registry Agent image from Docker Hub using the link provided in the [Settings prerequisites](./#settings-prerequisites-for-container-registry-agent). To run the image you can use a single environment variable for specifying the port:

```
docker run --restart=always \
       -p 8081:8081 \
       -e SNYK_PORT=8081 \
       snyk/container-registry-agent:latest
```

### **Container registry-specific configurations**

The following container registries require specific environment variables, setup, or both.

#### **DigitalOcean**

To set up Broker Client for **DigitalOcean**,`CR_USERNAME` and `CR_PASSWORD` are not required. Instead, you need to specify `CR_TOKEN`, the authentication token for DigitalOcean container registry.

#### **GCR and Google Artifact Registry**

To set up the Broker Client for these container registries, all the preceding information applies. Note that the `CR_USERNAME` value is permanent and should be `_json_key`, and the `CR_PASSWORD` value should be the JSON key used to authenticate to Google.

#### **Artifactory**

If you are using **Repository path** as your Docker access method, the container registry hostname in `CR_BASE` variable should be set in this structure: `<your artifactory host>/artifactory/api/docker/<artifactory-repo-name>`

Catalog endpoint `https://<artifactory-host>/artifactory/api/docker/<artifactory-repository>/v2/_catalog` is not required for importing a project in Artifactory; this is used for listing the image repositories.

#### **ECR**

<figure><img src="../../../.gitbook/assets/untitled (1) (1) (1) (1).png" alt="High-level architecture of the brokered ECR integration"><figcaption><p>High-level architecture of the brokered ECR integration</p></figcaption></figure>

#### **Required AWS Resource with ECR**

ECR setup requires the following kinds of IAM resources to be created:

* Container Registry Agent IAM Role or IAM User: an IAM Role or IAM User the Container Registry Agent uses to assume a cross-account role with access to ECR. It should have the following permissions: `"sts:AssumeRole"`
*   Snyk ECR Service Role: an IAM Role with access to ECR which is assumed by the Container Registry Agent IAM Role or IAM User to gain read-only access to ECR.\
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

The listed resources can be used as follows, so that a single Container Registry Agent instance can access ECR repositories located in different accounts:

**Run this step once only.** Create the Container Registry Agent IAM Role or IAM User and use it to run the Container Registry Agent. The IAM Role or IAM User could be provided to the Container Registry Agent using one of the methods described in the  [AWS docs](https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/setting-credentials-node.html).

**Run the following steps for each of your ECR accounts, using a separate Broker instance for each ECR account:**

1. In the AWS account where your ECR resides, create the Snyk ECR Service Role with read access to your ECR, and edit the trust relationship to allow this role to be assumed only by the specific Container Registry Agent IAM Role or IAM User created in the one-time step.
2. Restrict the Container Registry Agent IAM Role or IAM User to be allowed only to assume the Snyk ECR Service Role(s).
3. Provide the Broker Client with the Role ARN of the Snyk ECR Service Role together with the ECR region. The Broker Client passes this Role ARN to the Container Registry Agent, and the Container Registry Agent assumes it to access your ECR. The following environment variables are required:
   * CR\_ROLE\_ARN=\<the role ARN of SnykEcrServiceRole>
   * CR\_REGION=\<AWS Region of ECR>
   * CR\_EXTERNAL\_ID=\<Optional. An external ID found in the trust relationship condition>

For more information about the brokered ECR setup see [Configure AWS ECR](../../../more-info/getting-started/atlassian-integrations/atlassian/devops-pipeline-with-bitbucket-cloud-and-kubernetes/module-2-bitbucket-pipelines/configure-aws-ecr.md).

## **Configuring and using system check**

You can use the `/systemcheck` endpoint to verify connectivity between the Broker Client, the Container Registry Agent, and your container registry.

In order to use the endpoint, provide the following environment variable to the Broker Client:\
`BROKER_CLIENT_VALIDATION_URL=<agent-url>/systemcheck`

When you call the `/systemcheck` endpoint of the Broker Client, it uses the `BROKER_CLIENT_VALIDATION_URL` to make a request to the `/systemcheck` endpoint Container Registry Agent's, with the credentials provided to the Broker Client. The Container Registry Agent then makes a request to the container registry to validate connectivity.

{% hint style="info" %}
The `/systemcheck` endpoint is **not mandatory** for the brokered integration to function. For  more information see [Systemcheck](https://github.com/snyk/broker#systemcheck) in the Snyk Broker docs on GitHub.
{% endhint %}

## **Debugging methods for Container Registry Agent**

The `LOG_LEVEL` environment variable can be set to the desired level (debug/info/warn/error) in order to change the level of the Container Registry Agent and Broker Client logs.

For more verbose debugging, the Container Registry Agent can be run with the `DEBUG=*` environment variable. This allows printing the logs of the Node [Debug](https://www.npmjs.com/package/debug) package. The Debug package is used by several packages in the Container Registry Agent, among them the [Needle](https://www.npmjs.com/package/needle) package, which is used for making HTTP requests. Printing debug logs specifically from Needle can be done by setting `DEBUG=needle`.

{% hint style="danger" %}
**Warning:**\
Using the debugging options of third-party tools is not recommended for production environments, as it may result in logging sensitive information in logs that are not maintained by Snyk, for example, header information of HTTP requests.
{% endhint %}

## **Secure your images**

You can now start scanning your container images directly from your private registry. See [Configuring your JFrog Artifactory container registry integration](../../../scan-containers/image-scanning-library/jfrog-artifactory-image-scanning/configuring-your-jfrog-artifactory-container-registry-integration.md) for more details.
