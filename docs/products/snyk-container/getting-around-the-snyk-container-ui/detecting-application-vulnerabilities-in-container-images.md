# Detecting application vulnerabilities in container images

{% hint style="info" %}
For the Container Registry integration, the feature is supported for Node, Ruby, PHP, Python, Go binaries and Java. For the CLI and Kubernetes, the feature is supported for Node, Golang, and Java.
{% endhint %}

Snyk allows detection of vulnerabilities in your application dependencies from container images, as well as from the operating system, all in one single scan.

After you integrate with a container registry and import your projects, we scan your image and test for vulnerabilities.

### Enable application vulnerabilities scan from container images

1. Navigate to your container registry integration settings
2. Enable the _**Detect application vulnerabilities**_ capability and save the changes:

![](../../../.gitbook/assets/detect-app-vulns.png)

When scanning an image using a container registry, Kubernetes integration, or through the Docker scan command, the scan also uses the `--app-vulns` flag by default. You are able to opt out of the flag in the container registry only. Do so by disabling the ‘_detect application vulnerabilities_’ toggle in the integration settings.

### Using CLI to detect vulnerabilities

#### App Vulns Flag

Use the `--app-vulns` flag to detect vulnerabilities in your application dependencies from container images.&#x20;

For Java, when specifying the flag, we scan one level of nested jars by default.&#x20;

#### Nested Jars Depth Flag

For Java applications, when using `--app-vulns`, you can also use the `--nested-jars-depth=n` flag to set how many levels of nested jars Snyk will unpack. The implicit default is 1. I.e when specifying 2, it means that Snyk unzips jars in jars, 3 means Snyk unzips jars in jars in jars, etc.

Users can use --nested-jar-depth=0 to opt out of any scans they feel are unnecessary.&#x20;

### View vulnerabilities and licensing issues

After the feature is enabled, you can see:

* Dependency vulnerabilities and licensing issues of manifest files detected in your container image.
* Vulnerabilities detected in operating system packages.

When an image is imported to Snyk, it appears under its registry record in the **Projects** view, showing the operating system vulnerabilities found in your image.

With this feature enabled, you can also see nested manifest files detected in the image and their vulnerabilities and licensing issues.

![](<../../../.gitbook/assets/mceclip2 (1) (1) (1) (3) (3) (4) (6) (1) (23).png>)

### Automated scanning

Snyk scans the image regularly based on your project’s settings, and updates you via email or Slack - based on your configuration - when any new vulnerabilities are identified in both the operating system and application dependencies.

For each project, you can choose the test frequency under its settings (the default is daily testing).

![](<../../../.gitbook/assets/mceclip3 (1).png>)

**Supported registries**

This is supported across the following container registries:

* [ACR](https://docs.snyk.io/snyk-container/image-scanning-library/acr-image-scanning)
* [Amazon ECR](https://docs.snyk.io/snyk-container/image-scanning-library/ecr-image-scanning)&#x20;
* [JFrog Artifactory](https://docs.snyk.io/snyk-container/image-scanning-library/jfrog-artifactory-image-scanning)
* [Docker Hub](https://docs.snyk.io/snyk-container/image-scanning-library/docker-hub-image-scanning)&#x20;
* [GCR](https://docs.snyk.io/snyk-container/image-scanning-library/gcr-image-scanning)

**Supported Integrations**

The supported languages work on the following integrations:

| **Language** | **Container Registry** | **CLI** | **Kubernetes** |
| ------------ | ---------------------- | ------- | -------------- |
| Node         | Yes                    | Yes     | Yes            |
| Ruby         | Yes                    |         |                |
| PHP          | Yes                    |         |                |
| Python       | Yes                    |         |                |
| Go Binaries  | Yes                    | Yes     | Yes            |
| Java         | Yes                    | Yes     | Yes            |

****
