# How Snyk Container works

## Container images

As defined by the [Open Container Initiative](https://opencontainers.org) (OCI) specifications, container images comprise a layered file system and associated metadata.

Container images often include several layers containing third-party software from:

* Operating system distributions, such as Debian, Ubuntu, or CentOS. For details see [Snyk Container - Supported operating system distributions.](operating-system-distributions-supported-by-snyk-container.md)
* Application package managers, such as npm, pip, and RubyGems

## What Snyk Container detects

When Snyk Container tests an image using any of the available integrations, Snyk first finds the software installed in the image, including:

* dpkg, rpm, and apk operating systems packages
* Popular unmanaged software (installed outside a package manager)
* Application packages, based on the presence of a manifest file

{% hint style="info" %}
Snyk scans for packages and vulnerabilities in the final built image, not in the intermediate layers.
{% endhint %}

Because Snyk reads the information from the file system, the container does not need to be run. This means that for a successful scan, no container or foreign code must be run.

After Snyk has the list of installed software, Snyk looks that up in the Snyk Vulnerability Database, which combines public sources with proprietary research.

{% hint style="warning" %}
Snyk supports testing OCI-compliant and Docker v2-compliant images but does not support images that combine both OCI and Docker v2 standards into a single archive.
{% endhint %}

When providing public base image recommendations, Snyk's logic considers the origin repository, flavor, and version of the base image it detects. This means that Snyk's upgrade recommendations aim to reduce the number or severity of found vulnerabilities, with minor version upgrades typically maintaining compatibility with the currently used base image.

## Official advisory resources used by Snyk Container

Snyk works directly with the security teams of the supported Linux distributions to provide the most accurate and reliable information on the affected packages (including fix availability). The specific package versions listed are those distributed by the official container source and may differ from the upstream package versions.

## Unmanaged container software

Some software components from upstream providers are not installed using a package manager but are downloaded as executables from third parties. Snyk uses file fingerprinting to detect versions of the following components:

* Node.js
* OpenJDK 8 binaries

## Recurring container image scans

Snyk continuously discloses new vulnerabilities. Snyk can alert you to new vulnerabilities in your image as they are announced, even when your installed image software has not changed.

If you use an integration that saves a snapshot of the installed software on Snyk’s service and the image has not changed, Snyk Container automatically rescans without accessing the image, alerting you to new vulnerabilities.

If the image has changed, you must reimport your image, so that Snyk can access the updated image and its metadata before rescanning it.

{% hint style="info" %}
Recurring scans do not detect updates to the dependencies of your applications. The recurring scans test for new vulnerabilities using a snapshot of your application dependencies at the time the application was imported.\
\
To detect changes in your application, such as updated dependencies, you must re-import your container image in Snyk.
{% endhint %}
