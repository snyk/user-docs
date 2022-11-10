# Snyk Container - Integrations

Snyk Container provides a range of integrations to import projects into Snyk. These integrations support different workflows for our users and customers.

<figure><img src="../../.gitbook/assets/Snyk container integrations.png" alt="Screenshot of a Snyk Container project in the Projects listing page"><figcaption></figcaption></figure>

The integrations you use depends on your requirements and workflow. You can start with one integration and move to others later, or use a combination of integrations as your team grows.

For example, it’s common to use our CI integrations to provide fast feedback to development teams when you build an image, then use the Kubernetes integration to provide assurance around running images in production.

### Integrations available

The main integrations are:

* **CLI**: useful for local investigation, or testing an image you have built. See [snyk-cli-for-container-security](snyk-cli-for-container-security/ "mention").
* SCM: Snyk can detect Dockerfiles directly from Git repositories, and provide recommendations for updating the base image to a less vulnerable one. See [scan-your-dockerfile](scan-your-dockerfile/ "mention").
* CI: can act as a gate, for example breaking the build on new high severity vulnerabilities.
* Container registries: useful to test a large number of images, or if you cannot modify lots of CI pipelines.
* [Kubernetes](https://support.snyk.io/hc/en-us/articles/360003916138-Kubernetes-integration-overview): similar to container registries, but with more context about the running workload Snyk can use to prioritize vulnerabilities or group projects.

See [image-scanning-library](image-scanning-library/ "mention") for more details.
