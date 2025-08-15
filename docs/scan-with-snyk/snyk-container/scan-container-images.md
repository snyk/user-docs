# Scan container images

Snyk Container helps you find and fix vulnerabilities in container images, based on container registry scans.

You can scan your container images using Snyk Container:&#x20;

* In the [Snyk Web UI](use-snyk-container/)
* Through the [Snyk CLI](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)
* With [Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/integrate-with-self-hosted-container-registries-broker.md) (for self-hosted container registries)

## **Prerequisites for using Snyk Container in the Web UI**

Before scanning your container images with Snyk Container, ensure you:

* Create or log in to a Snyk account.
* Set up an integration with a supported container registry, such as Docker Hub. See [Container security integrations](container-registry-integrations/).

For more information, see [Getting started](../../discover-snyk/getting-started/).

## View vulnerabilities in your container images

In the **Projects** tab, you can see vulnerability results for Snyk Projects that you have imported. The imported Projects are grouped into **Targets**.

{% hint style="info" %}
You can see the history of all the repositories and container registry images imported into an Organization. For details, see [Import Log](../../snyk-platform-administration/snyk-projects/import-log.md).
{% endhint %}

To see vulnerability information for that Project, select an imported Project from the target list.

Click on a Project entry to see details of the vulnerabilities found, including where it was introduced, how to fix it, and other details about the vulnerability.

## Fix vulnerabilities in your container images

To fix vulnerabilities in your container images:

1. Open a PR based on Snyk recommendation by clicking **Open a fix PR**.
2. Upgrade or rebuild your image.

After the updated image is pushed, Snyk automatically rescans your new image.&#x20;

For more details on fixing vulnerabilities, see [Analysis and fixes for your images from the Snyk Web UI](use-snyk-container/analyze-and-fix-container-images.md).
