# Snyk Container for self-hosted container registries (with broker)

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can integrate to private container registries you host that do not have internet access , and help you to better secure container images in those registries.

To enable and configure your hosted container registry, contact our support team at [support@snyk.io](mailto:support@snyk.io)

## **Self-hosted container registries solution components**

* **Broker server:** running on Snyk SaaS backend
* **Broker client & Container Registry Agent:** Two Docker images deployed in your infrastructure, creating two separate services, responsible for sampling your container registries in a secured manner and sending the allowed information to Snyk.

The Broker client provides the Agent with the connection details. The Agent uses these details to connect to the container registry, scan the images, and send the scan results through the brokered communication using callbacks. The brokered communication happens when a Broker Client connects (using your Broker ID) to a Broker server which runs within the Snyk environment.

See [Snyk Broker Container Registry Agent](../integrations/snyk-broker/snyk-broker-container-registry-agent/) documentation for more details.

![](../.gitbook/assets/mceclip0-8-.png)

#### **Supported Container registries**

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
