# Scan container images

Snyk Container helps you find and fix vulnerabilities in container images, based on container registry scans.

You can scan your container images using Snyk Container:&#x20;

* In the [Snyk Web UI](../snyk-container/use-snyk-container-from-the-web-ui/)
* Through the [Snyk CLI](../snyk-container/use-snyk-container-from-the-cli/)
* With [Broker](../snyk-container/integrate-with-self-hosted-container-registries-broker.md) (for self-hosted container registries)

## **Prerequisites for using Snyk Container in the Web UI**

Before scanning your container images with Snyk Container, ensure you:

* Have completed the [quickstart](../../getting-started/quickstart/) steps.
* Set up an integration with a supported container registry, like Docker Hub. See Snyk Container integrations.

## View vulnerabilities in your container images

In the **Projects** tab, you can see vulnerability results for Snyk Projects that you have imported. The imported Projects are grouped into **Targets**.

To see vulnerability information for that Project, select an imported Project from the target list.

<figure><img src="../../.gitbook/assets/vuln_info_project.png" alt=""><figcaption><p>Vulnerability information for an imported Snyk Project</p></figcaption></figure>

Click on a Project entry to see details of the vulnerabilities found, including where it was introduced, how to fix it, and other details about the vulnerability.

<figure><img src="../../.gitbook/assets/vuln_details.png" alt=""><figcaption><p>Vulnerability information</p></figcaption></figure>

## Fix vulnerabilities in your container images

To fix vulnerabilities in your container images:

1. Open a PR based on Snyk recommendation by clicking **Open a fix PR**.
2. Upgrade or rebuild your image.

After it is pushed, Snyk automatically rescans your new image.&#x20;

For more details on how to fix vulnerabilities, see [Analysis and fixes for your images from the Snyk Web UI](../snyk-container/use-snyk-container/analyze-and-fix-container-images.md).
