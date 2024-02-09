# Integrate with self-hosted container registries (Broker)

{% hint style="warning" %}
**Release status**&#x20;

Self-hosted container registries (Broker) are available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

Snyk can integrate with self-hosted private container registries that do not have internet access and can help you to better secure container images in those registries.

To enable and configure your hosted container registry, contact Snyk admin.

## **Solution components of self-hosted container registries**

Integration with self-hosted container registries contains the following components:

* **Broker Server** - running on Snyk SaaS backend
* **Broker Client and Container Registry Agent** - two Docker images deployed in your infrastructure, creating two separate services responsible for sampling your container registries in a secured manner and sending the allowed information to Snyk.

<figure><img src="../../.gitbook/assets/mceclip0-8-.png" alt="High-level architecture of the Snyk Broker Container Registry Agent"><figcaption><p>High-level architecture of the Snyk Broker Container Registry Agent</p></figcaption></figure>

The Broker Client provides the Agent with the connection details. The Agent uses these details to connect to the container registry, scan the images, and send the scan results through the brokered communication using callbacks.

The brokered communication happens when a Broker Client connects (using your Broker ID) to a Broker server that runs within the Snyk environment.

For more details, see [Snyk Broker - Container Registry Agent](../../enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/).

## **Supported container registries**

* Docker Hub (type: docker-hub)
* GCR (type: gcr)
* ECR (type: ecr)
* Azure (type: acr)
* Artifactory (type: artifactory-cr)
* Harbor (type: harbor-cr)
* Quay (type: quay-cr)
* GitHub (type: github-cr)
* Nexus (type: nexus-cr)
* DigitalOcean (type: digitalocean-cr)
* GitLab (type: gitlab-cr)
* Google Artifact Registry (type: google-artifact-cr)
