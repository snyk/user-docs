# Snyk Container - Integrations

Snyk Container provides a range of integrations to import Projects into Snyk. These integrations support different workflows for Snyk users and customers.

<figure><img src="../../.gitbook/assets/Snyk container integrations.png" alt="Container Project on Projects listing page"><figcaption><p>Container Project on Projects listing page</p></figcaption></figure>

The integrations you use depend on your requirements and workflow. You can start with one integration and move to others later, or use a combination of integrations as your team grows.

For example, it’s common to use Snyk CI integrations to provide fast feedback to development teams when you build an image, and then use the Kubernetes integration to provide assurance for running images in production.

The main Container integrations are:

* CLI: Use for local investigation or testing an image you have built. For details, see [Snyk CLI for container security](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/).
* SCM: Snyk can detect Dockerfiles directly from Git repositories and provide recommendations for updating the base image to a less vulnerable one. For details, see [Scan your Dockerfile](../../scan-using-snyk/snyk-container/scan-your-dockerfile/).
* CI: can act as a gate, for example, by breaking the build on new high-severity vulnerabilities.
* Container registries: Use to test a large number of images, or if you cannot modify lots of CI pipelines.
* Kubernetes: Use the same way as container registries, but Kubernetes provides more context about the running workload Snyk can use to prioritize vulnerabilities or group Projects. For details, see [Kubernetes integration overview.](../../scan-using-snyk/snyk-container/integrate-with-kubernetes/overview-of-the-kubernetes-integration/)

{% hint style="info" %}
For cloud-hosted container registries, Snyk does not import and scan images that are larger than 2GB in size. To scan images that exceed this size, use the [Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/).
{% endhint %}
