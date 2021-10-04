# Detecting application vulnerabilities in container images

{% hint style="info" %}
The feature is currently supported for Node, Ruby, PHP, Python, Go binaries and Java.
{% endhint %}

Snyk allows detection of vulnerabilities in your application dependencies from container images, as well as from the operating system, all in one single scan.

After you integrate with a container registry and import your projects, we scan your image and test for vulnerabilities.

## Enable application vulnerabilities scan from container images

1. Navigate to your container registry integration settings
2. Enable the _**Detect application vulnerabilities**_ capability and save the changes:

![](../../../.gitbook/assets/mceclip1%20%281%29.png)

## View vulnerabilities and licensing issues

After the feature is enabled, you can see:

* Dependency vulnerabilities and licensing issues of manifest files detected in your container image.
* Vulnerabilities detected in operating system packages.

When an image is imported to Snyk, it appears under its registry record in the **Projects** view, showing the operating system vulnerabilities found in your image.

With this feature enabled, you can also see nested manifest files detected in the image and their vulnerabilities and licensing issues.

![](../../../.gitbook/assets/mceclip2%20%281%29%20%281%29%20%281%29%20%283%29%20%283%29%20%284%29%20%286%29%20%281%29%20%2818%29.png)

## Automated scanning

Snyk scans the image regularly based on your project’s settings, and updates you via email or Slack - based on your configuration - when any new vulnerabilities are identified in both the operating system and application dependencies.

For each project, you can choose the test frequency under its settings \(the default is daily testing\).

![](../../../.gitbook/assets/mceclip3%20%281%29.png)

**Supported registries**

This is supported across the following container registries:

* [ACR](https://docs.snyk.io/snyk-container/image-scanning-library/acr-image-scanning)
* [Amazon ECR](https://docs.snyk.io/snyk-container/image-scanning-library/ecr-image-scanning) 
* [JFrog Artifactory](https://docs.snyk.io/snyk-container/image-scanning-library/jfrog-artifactory-image-scanning)
* [Docker Hub](https://docs.snyk.io/snyk-container/image-scanning-library/docker-hub-image-scanning) 
* [GCR](https://docs.snyk.io/snyk-container/image-scanning-library/gcr-image-scanning)

**Supported Integrations**

The supported languages work on the following integrations:

| **Language** | **Container Registry** | **CLI** | **Kubernetes** |
| :--- | :--- | :--- | :--- |
| Node | Yes | Yes | Yes |
| Ruby | Yes |  |  |
| PHP | Yes |  |  |
| Python | Yes |  |  |
| Go Binaries | Yes | Yes | Yes |
| Java | Yes | Yes | Yes |

**Language Limitations**

These are the current limitations of the language support:

**Java**

Scan only the top level JAR - we do not currently unpack the JAR

