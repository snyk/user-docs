# Detecting application vulnerabilities in container images

{% hint style="info" %}
For the Container Registry integration, the feature is supported for Node, Ruby, PHP, Python, Go binaries, and Java. For the CLI and Kubernetes, the feature is supported for Node, Go binaries, PHP, Python, and Java.
{% endhint %}

{% hint style="warning" %}
For applications projects created from images imported from **Container Registry integrations**, the applications will not be re-imported during recurring tests or manual re-test.

Instead, the applications dependencies that were found during the initial image import will be tested for new vulnerabilities.

This means that if new dependencies were introduced in an application within an image, they will not be detected by the recurring tests or manual re-test.

In order to detect **new or updated** applications within images from container registries, the image will need to be re-imported to snyk.

For applications found in images imported from **the Kubernetes integration**, existing applications will be re-imported, but new apps added to the image will not be imported during recurring tests.

In order to detect **new** applications within images from Kubernetes, the image will need to be re-imported to snyk.
{% endhint %}

Snyk allows detection of vulnerabilities in your application dependencies from container images, as well as from the operating system, all in one single scan.

After you integrate with a container registry and import your projects, Snyk scans your image and test for vulnerabilities.

### Enable application vulnerabilities scan from container registries

1. Navigate to your container registry integration settings
2. Enable the _**Detect application vulnerabilities**_ capability and save the changes:

![](../../../.gitbook/assets/detect-app-vulns.png)

When you are scanning an image using a container registry or Kubernetes integration, the scan also uses the `--app-vulns` flag by default. You can opt out of the flag in the container registry only. Do so by disabling the ‘_detect application vulnerabilities_’ toggle in the integration settings.

### Notes

1. For Java, when you use the flag, Snyk scans one level of nested jars by default.
2. For Python, Snyk supports Poetry and Pip (in all integration points).
3. For Go binaries, Snyk supports any kind of a Go binary built with Go module support.

### Using CLI to detect vulnerabilities

#### App Vulns Flag

In CLI versions 1.1090.0 (2023-01-24) and higher, Snyk scans for application dependencies in your image by default; you do not need to specify the `--app-vulns` flag.

If you wish to opt out of application vulnerability scanning, you can do so by specifying the `--exclude-app-vulns` flag — which will omit the application vulnerabilities section from the results, mimicking the previous behavior. The `--exclude-app-vulns` flag is available in CLI version 1.1021.0 and above.

#### Nested Jars Depth Flag

For Java applications, when `--app-vulns` is enabled, you can also use the `--nested-jars-depth=n` flag to set how many levels of nested jars Snyk will unpack. The implicit default is 1. When you specify 2, it means that Snyk unzips jars in jars; 3 means Snyk unzips jars in jars in jars, and so on.

Users can use `--nested-jar-depth=0` to opt out of any scans they feel are unnecessary.

### View vulnerabilities and licensing issues

After the feature is enabled, you can see:

* Dependency vulnerabilities and licensing issues of manifest files detected in your container image.
* Vulnerabilities detected in operating system packages.

When an image is imported to Snyk, it appears under its registry record in the **Projects** view, showing the operating system vulnerabilities found in your image.

With this feature enabled, you can also see nested manifest files detected in the image and their vulnerabilities and licensing issues.

![](<../../../.gitbook/assets/mceclip2 (1) (1) (1) (3) (3) (4) (6) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (  (9).png>)

### Automated scanning

Snyk scans the image regularly based on your project’s settings, and updates you via email or Slack - based on your configuration - when any new vulnerabilities are identified in both the operating system and application dependencies.

For each project, you can choose the test frequency under its settings (the default is daily testing).

![](<../../../.gitbook/assets/mceclip3 (1).png>)

**Supported container registries**

This is supported across the following container registries:

* [ACR](https://docs.snyk.io/snyk-container/image-scanning-library/acr-image-scanning)
* [Amazon ECR](https://docs.snyk.io/snyk-container/image-scanning-library/ecr-image-scanning)
* [JFrog Artifactory](https://docs.snyk.io/snyk-container/image-scanning-library/jfrog-artifactory-image-scanning)
* [Docker Hub](https://docs.snyk.io/snyk-container/image-scanning-library/docker-hub-image-scanning)
* [GCR](https://docs.snyk.io/snyk-container/image-scanning-library/gcr-image-scanning)
* [DigitalOcean](https://docs.snyk.io/products/snyk-container/image-scanning-library/digitalocean-image-scanning)
* [GitHub](https://docs.snyk.io/products/snyk-container/image-scanning-library/github-container-registry-image-scanning)
* [GitLab](https://docs.snyk.io/products/snyk-container/image-scanning-library/gitlab-container-registry-image-scanning)
* [Harbor](https://docs.snyk.io/products/snyk-container/image-scanning-library/harbor-image-scanning)
* [Nexus](https://docs.snyk.io/products/snyk-container/image-scanning-library/nexus-image-scanningexsd)
* [Quay](https://docs.snyk.io/products/snyk-container/image-scanning-library/quay-image-scanning)

**Supported Integrations**

The supported languages work on the following integrations:

| **Language** | **Container Registry** | **CLI** | **Kubernetes** |
| ------------ | ---------------------- | ------- | -------------- |
| Node         | Yes                    | Yes     | Yes            |
| Ruby         | Yes                    |         |                |
| PHP          | Yes                    | Yes     | Yes            |
| Python       | Yes                    | Yes     | Yes            |
| Go Binaries  | Yes                    | Yes     | Yes            |
| Java         | Yes                    | Yes     | Yes            |

***
