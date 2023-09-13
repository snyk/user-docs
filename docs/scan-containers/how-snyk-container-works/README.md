# How Snyk Container works

## Snyk tests container images

As defined by the [Open Container Initiative](https://opencontainers.org) (OCI) specifications, container images comprise a layered file system and associated metadata.&#x20;

Container images often include several layers containing third-party software from:

* Operating system distributions, such as Debian, Ubuntu, or CentOS
* Application package managers, such as npm, pip, and RubyGems.

## What Snyk Container detects

When Snyk Container tests an image using any of the available integrations, Snyk first finds the software installed in the image, including:

* dpkg, rpm, and apk operating systems packages
* Popular unmanaged software (installed outside a package manager)
* Application packages that are based on the presence of a manifest file.

Because Snyk reads the information from the file system, the container does not need to be run. This means that for a successful scan, no container or foreign code must be run.

After Snyk has the list of installed software, Snyk looks that up against the Snyk Vulnerability Database, which combines public sources with proprietary research.

{% hint style="warning" %}
Snyk supports testing OCI compliant and Docker v2 complaint images but does not support images that combine both OCI and Docker v2 standards into a single archive.
{% endhint %}

## Operating systems supported by Snyk Container

Snyk detects vulnerabilities in images based on the following operating systems:

* AlmaLinux
* Alpine Linux
* Amazon Linux
* CentOS Linux & CentOS Stream
* Chainguard
* Debian
* Oracle Linux
* Red Hat Enterprise Linux (RHEL), including Universal Base Image (UBI)
* Rocky Linux
* SUSE Linux Enterprise Server (SLES)
* Ubuntu
* Wolfi

{% hint style="info" %}
Snyk also supports images using packages from these distributions, but without the associated package manager, such as Distroless images.
{% endhint %}

For specific version support, see [supported operating system distributions](../../scan-applications/snyk-container/how-snyk-container-works/supported-operating-system-distributions.md). For the latest updates, see [Snyk updates](https://updates.snyk.io).

## Official container advisory sources

Snyk works directly with the security teams of the supported Linux distributions to provide the most accurate and reliable information on the affected packages (including fix availability). The specific package versions listed are those distributed by the official container source and may differ from the upstream package versions.

## Unmanaged container software

Some software components from upstream providers are not installed using a package manager but are downloaded as executables from third parties. Snyk uses file fingerprinting to detect versions of the following components:

* Node.js
* OpenJDK 8 binaries

## Recurring scans with Snyk Container

Snyk continuously discloses new vulnerabilities. Snyk can alert you to new vulnerabilities in your image as they are announced, even when your image software installed has not changed.

If you use an integration that saves a snapshot of the installed software on Snyk’s service and the image has not changed, Snyk Container automatically rescans without accessing the image, alerting you to new vulnerabilities.

If the image has changed, you must reimport your image, so that Snyk can access the updated image and its metadata before rescanning it.&#x20;

{% hint style="info" %}
Recurring scans do not detect updates to the dependencies of your applications. The recurring scans test for new vulnerabilities using a snapshot of your application dependencies at the time the application was imported.\
\
To detect changes in your application, such as updated dependencies, you must re-import your container image in Snyk. To see an example of how to import your image, see [Getting started with Snyk Container](../getting-started-with-snyk-container.md).
{% endhint %}

To learn more about container security, see [container security](https://snyk.io/learn/container-security/).
