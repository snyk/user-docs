# Detect vulnerabilities in container images

{% hint style="info" %}
**Feature availability**

For Container Registry integration, the feature is supported for Node, Ruby, PHP, Python, Go binaries, and Java.

For the CLI and Kubernetes, the feature is supported for Node, PHP, Python, Go binaries, and Java.
{% endhint %}

Snyk is able to detect in one scan the vulnerabilities in your application dependencies from container images, as well as from the operating system.

After you integrate with a container registry and import your Projects, Snyk scans your image for vulnerabilities.

{% hint style="warning" %}
For application Projects created from images that are imported from container registry integrations, the applications are not re-imported during recurring scans or manual rescans.

Instead, the application dependencies that are found during the initial image import are scanned for new vulnerabilities.
{% endhint %}

If new dependencies are introduced in an application within an image, they will not be detected by recurring scans or manual rescans.

To detect new or updated applications within images from container registries, you must re-import the image to Snyk.

For applications found in images imported from the Kubernetes integration, existing applications will be re-imported, but new apps added to the image will not be imported during recurring scans.

To detect new applications within images from Kubernetes, you must re-import the image to Snyk.

## Enable container registries vulnerability scan

To enable the application vulnerability scan from container registries:

1. Navigate to your container registry integration settings.
2. Enable **Detect application vulnerabilities** and save the changes.

<figure><img src="../../../.gitbook/assets/enable_detect_app_vuln (1).png" alt=""><figcaption><p>Enable Detect application vulnerabilties</p></figcaption></figure>

When you are scanning an image using a container registry or Kubernetes integration, the scan also uses the `--app-vulns` flag by default. You can opt out of the flag in the container registry only. To do this, disable the **Detect application vulnerabilities** feature.

{% hint style="info" %}
* For Java, when you use the flag, Snyk scans one level of nested jars by default.
* For Python, Snyk supports Poetry and Pip (in all integration points).
* For Go binaries, Snyk supports any type of Go binary built with Go module support.
{% endhint %}

## Snyk Container CLI options to detect vulnerabilities

### App vulns option

In CLI versions 1.1090.0 (2023-01-24) and higher, Snyk scans for application dependencies in your image by default; you do not need to specify the `--app-vulns` option.

If you want to opt out of application vulnerability scanning, you can specify the `--exclude-app-vulns` flag. This omits the application vulnerabilities section from the results, mimicking the previous behavior. The `--exclude-app-vulns` option is available in CLI version 1.1021.0 and above.

### Nested jars depth option

For Java applications, when `--app-vulns` is enabled, you can also use the `--nested-jars-depth=n` option to set how many levels of nested jars Snyk will unpack. The implicit default is 1. When you specify 2, it means that Snyk unzips jars in jars; 3 means Snyk unzips jars in jars in jars, and so on.

To opt out of any scans you feel are unnecessary, use `--nested-jar-depth=0` .

## View vulnerabilities and licensing issues

When the **Detect application vulnerabilities** feature is enabled, it allows you to see:

* Dependency vulnerabilities and licensing issues of manifest files detected in your container image
* Vulnerabilities detected in operating system packages.

When an image is imported to Snyk, it appears under its registry record in the **Projects** view, showing the operating system vulnerabilities found in your image.

With this feature enabled, you can also see nested manifest files detected in the image and their vulnerabilities and licensing issues.

<figure><img src="../../../.gitbook/assets/mceclip2 (1) (1) (1) (3) (3) (4) (6) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) ( (31).png" alt="Images listed in Projects view"><figcaption><p>Images listed in Projects view</p></figcaption></figure>

## Set up automated scanning

Depending on your Project settings, Snyk scans the image regularly. Based on your configuration, Snyk updates you through email or Slack whenever new vulnerabilities are identified in both the operating system and application dependencies.

You can set the scan frequency for each Project. The default is daily testing. To update the scan frequency, navigate to the **Settings** tab on the Project page and select the frequency from the dropdown. The options are never, **daily**, or **weekly**.

<figure><img src="../../../.gitbook/assets/scan_frequency.png" alt=""><figcaption><p>Update scan frequency</p></figcaption></figure>

## **Supported container registries**

The following container registries are supported:&#x20;

* Docker Hub
* GCR
* ACR
* Amazon ECR
* JFrog Artifactory
* Harbor
* Quay
* GitHub
* Nexus
* DigitalOcean
* GitLab

## **Supported integrations**

The supported languages work with the following integrations:

| **Language** | **Container Registry** | **CLI** | **Kubernetes** |
| ------------ | ---------------------- | ------- | -------------- |
| Node         | Yes                    | Yes     | Yes            |
| Ruby         | Yes                    |         |                |
| PHP          | Yes                    | Yes     | Yes            |
| Python       | Yes                    | Yes     | Yes            |
| Go Binaries  | Yes                    | Yes     | Yes            |
| Java         | Yes                    | Yes     | Yes            |
