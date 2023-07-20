# How Snyk Container works

## Snyk tests container images

Container images comprise a layered file system and associated metadata, as defined by the [Open Container Initiative](https://opencontainers.org) (OCI) specifications.

Container images often include several layers containing third-party software from:

* Operating system distributions, such as Debian, Ubuntu, or CentOS
* Application package managers, such as npm, pip, and RubyGems

## What Snyk Container detects

When Snyk Container tests an image using any of the available integrations, Snyk first finds the software installed in the image, including:

* dpkg, rpm, and apk operating systems packages
* Popular unmanaged software, that is, installed outside a package manager
* Application packages based on the presence of a manifest file

The container does not need to be run as Snyk reads the information from the file system; therefore, no container or foreign code must be run to scan successfully.

After Snyk has the list of installed software, Snyk looks that up against the Snyk Vulnerability Database, which combines public sources with proprietary research.

{% hint style="warning" %}
Snyk supports testing OCI compliant and Docker v2 complaint images but does not support images which combine both OCI and Docker v2 standards into a single archive.
{% endhint %}

## Operating systems supported by Snyk Container

Snyk detects vulnerabilities in images based on the following:

* Alpine Linux
* Amazon Linux
* CentOS Linux & CentOS Stream
* Debian
* Oracle Linux
* Red Hat Enterprise Linux (RHEL), including Universal Base Image (UBI)
* Rocky Linux
* SUSE Linux Enterprise Server (SLES)
* Ubuntu

{% hint style="info" %}
Snyk also supports images using packages from those distributions, but without the associated package manager, such as Distroless images.
{% endhint %}

See the [supported operating system distributions ](supported-operating-system-distributions.md)page for specific version support and  [Snyk updates](https://updates.snyk.io) for the latest updates.

## Official container advisory sources

Snyk works directly with the security teams of the supported Linux distributions to provide the most accurate and reliable information on the affected packages (including fix availability). The specific package versions listed are those distributed by the official container source and may differ from the upstream package versions.

## Unmanaged container software

Some software components from upstream providers are not installed using a package manager but are downloaded as executables from third parties. Snyk uses file fingerprinting to detect versions of the following components:

* Node.js
* OpenJDK 8 binaries

## Recurring scans with Snyk Container

New vulnerabilities are disclosed continuously. Snyk can alert you to new vulnerabilities in your image as they are announced, even when your image software installed has not changed.

If you use an integration that saves a snapshot of the installed software on Snyk’s service and the image has not changed, Snyk Container automatically rescans without accessing the image, alerting you to new vulnerabilities. However, if the image has changed, Snyk will access and pull it before rescanning it.

{% hint style="info" %}
Note that recurring scans do not detect updates to the dependencies of your applications. The recurring scans test for new vulnerabilities using a snapshot of your application dependencies at the time the application was imported.\
\
To detect changes in your application, such as updated dependencies, re-import your container image in Snyk. See [Getting started with Snyk Container](../getting-started-with-snyk-container.md) for an example of how to import your image.
{% endhint %}

## More about container security

Learn more about [container security](https://snyk.io/learn/container-security/).
