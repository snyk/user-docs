# How Snyk Container works

## What are container images?

Container images comprise a layered file system and associated metadata, as defined by the [Open Container Initiative](https://opencontainers.org) (OCI) specifications.

Container images often include several layers containing third-party software from:

* Operating system distributions, such as Debian, Ubuntu or CentOS.
* Application package managers, such as npm, pip and RubyGems.

## What Snyk Container detects

When Snyk Container scans an image, using any of the available integrations, we first find the software installed in the image, including:

* dpkg, rpm and apk operating systems packages.
* Popular unmanaged software, ie. installed outside a package manager.
* Application packages based on the presence of a manifest file.

**Note:** the container does not need to be run as Snyk reads the info from the file system; therefore, no container or foreign code needs to be run in order to successfully scan.

After we have the list of installed software, we look that up against our vulnerability database, which combines public sources with proprietary research.

## Supported operating systems

We detect vulnerabilities in images based on:

* Alpine
* Debian
* Ubuntu
* CentOS
* Red Hat Enterprise Linux (RHEL), including Universal Base Image (UBI)
* Amazon Linux
* Oracle Linux
* SUSE Linux Enterprise Server

Check out the [Operating Systems Support](https://docs.snyk.io/snyk-container/snyk-container-security-basics/supported-operating-system-distributions) page for specific version support and our [updates](https://updates.snyk.io) page for all the latest updates.

{% hint style="info" %}
**Note:** Snyk also supports images using packages from those distributions but without the associated package manager, such as Distroless images.
{% endhint %}

## Unmanaged software

Some software components from upstream providers are not installed using a package manager, but are downloaded as executables from third-parties. Snyk uses file fingerprinting to detect versions of the following components:

* Node.js
* OpenJDK

## Recurring scans

New vulnerabilities are disclosed continuously. Snyk can alert you to new vulnerabilities in your image as they are announced, even when your image software installed has not changed.

If you use an integration which saves a snapshot of the installed software on Snyk’s service, Snyk Container automatically rescans without accessing the image, alerting you to new vulnerabilities quicker.

{% hint style="info" %}
Note that recurring scans do not detect updates to the dependencies of your applications. The recurring scans simply test for new vulnerabilities using a snapshot of your application dependencies at the time the application was imported.\
\
To detect changes in your application, such as updated dependencies, please re-import your container image in Snyk. See [getting-started-snyk-container.md](../getting-started-snyk-container.md "mention") for an example on how to import your image.
{% endhint %}

Learn more about [container security](https://snyk.io/learn/container-security/).
