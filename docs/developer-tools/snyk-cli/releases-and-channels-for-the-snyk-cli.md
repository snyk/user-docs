# Releases and channels for the Snyk CLI

This page describes Snyk CLI releases and support policy, and also explains how to opt in to different channels and the purpose of each channel.

## Releases

Beginning with v.1.1291.0, the Snyk CLI follows the industry standard [Semantic Versioning](https://semver.org/) three-part notation as follows.

Given a version number MAJOR.MINOR.PATCH, increment the:

1. MAJOR version when you make \[breaking] changes
2. MINOR version when you add functionality in a backward-compatible manner
3. PATCH version when you make backward-compatible bug fixes

Additional labels are added for CLI releases as needed based on the standard.

In the Snyk CLI context, Snyk defines a breaking change as a change that could break automated workflows and cause failures in your existing working setup, such as CI/CD integrations. Breaking changes will be indicated by incrementing MAJOR and also mentioned in the release notes.

Some examples of breaking changes are the following:

* Deprecating or changing output fields, field names, or environment variables
* Introducing mandatory configuration changes
* Changes in error or exit codes

## Channels

Beginning with v.1.1291.0, Snyk is providing different channels to enable customers to opt in to a channel based on their needs and preferences.

When you select a channel, you are selecting the stability level you want to use: **preview**, **rc,** or **stable**.

### preview

{% hint style="info" %}
Snyk offers a preview channel for those who want to test _in-progress_ features. However, keep in mind that this channel may contain bugs and is not officially supported.
{% endhint %}

A preview version is not recommended for production environments. It may contain bugs and is best tested in a local environment. For instructions on installing a preview version, see [Install standalone executables from a channel.](releases-and-channels-for-the-snyk-cli.md#install-standalone-executables-from-a-channel)

* Preview: pre-release builds are deployed regularly up to multiple times a day and contain the latest changes.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}-preview
* Cadence: Varying
* Availability:
  * Linux: [https://downloads.snyk.io/cli/preview/snyk-linux](https://downloads.snyk.io/cli/preview/snyk-linux)
  * Windows: [https://downloads.snyk.io/cli/preview/snyk-win.exe](https://downloads.snyk.io/cli/preview/snyk-win.exe)
  * MacOS: [https://downloads.snyk.io/cli/preview/snyk-macos](https://downloads.snyk.io/cli/preview/snyk-macos)
  * Alpine: [https://downloads.snyk.io/cli/preview/snyk-alpine](https://downloads.snyk.io/cli/preview/snyk-alpine)
  * MacOS arm64: [https://downloads.snyk.io/cli/preview/snyk-macos-arm64](https://downloads.snyk.io/cli/preview/snyk-macos-arm64)
  * Linux arm64: [https://downloads.snyk.io/cli/preview/snyk-linux-arm64](https://downloads.snyk.io/cli/preview/snyk-linux-arm64)
  * Alpine arm64: [https://downloads.snyk.io/cli/preview/snyk-alpine-arm64](https://downloads.snyk.io/cli/preview/snyk-alpine-arm64)
  * For FIPS, add `fips` to the base URL, for example, `https://downloads.snyk.io/fips/cli/preview/snyk-linux`&#x20;
  * For experimental builds, add `experimental` to the base URL, for example, `https://downloads.snyk.io/experimental/cli/preview/snyk-linux`&#x20;

### rc

* Release candidate: pre-releases are deployed at distinct points in time and contain a version of the CLI that is expected to be promoted to stable after additional testing.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}-rc
* Cadence: every eight weeks, two weeks before a stable release (hotfix releases possible)
* Availability:
  * Linux: [https://downloads.snyk.io/cli/rc/snyk-linux](https://downloads.snyk.io/cli/rc/snyk-linux)
  * Windows: [https://downloads.snyk.io/cli/rc/snyk-win.exe](https://downloads.snyk.io/cli/rc/snyk-win.exe)
  * MacOS: [https://downloads.snyk.io/cli/rc/snyk-macos](https://downloads.snyk.io/cli/rc/snyk-macos)
  * Alpine: [https://downloads.snyk.io/cli/rc/snyk-alpine](https://downloads.snyk.io/cli/rc/snyk-alpine)
  * MacOS arm64: [https://downloads.snyk.io/cli/rc/snyk-macos-arm64](https://downloads.snyk.io/cli/rc/snyk-macos-arm64)
  * Linux arm64: [https://downloads.snyk.io/cli/rc/snyk-linux-arm64](https://downloads.snyk.io/cli/rc/snyk-linux-arm64)
  * Alpine arm64: [https://downloads.snyk.io/cli/rc/snyk-alpine-arm64](https://downloads.snyk.io/cli/rc/snyk-alpine-arm64)
  * For fips, add `fips` to the base URL, for example, `https://downloads.snyk.io/fips/cli/rc/snyk-linux`&#x20;
  * For experimental builds, add `experimental` to the base URL, for example, `https://downloads.snyk.io/experimental/cli/rc/snyk-linux`&#x20;

### stable

* Stable: builds are deployed at distinct points in time after additional testing and are considered stable.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}
* Cadence: every eight weeks (hotfix releases possible)
* Availability:
  * [https://github.com/snyk/cli/releases/](https://github.com/snyk/cli/releases/)
  * Linux: [https://downloads.snyk.io/cli/stable/snyk-linux](https://downloads.snyk.io/cli/stable/snyk-linux)
  * Windows: [https://downloads.snyk.io/cli/stable/snyk-win.exe](https://downloads.snyk.io/cli/stable/snyk-win.exe)
  * MacOS: [https://downloads.snyk.io/cli/stable/snyk-macos](https://downloads.snyk.io/cli/stable/snyk-macos)
  * Alpine: [https://downloads.snyk.io/cli/stable/snyk-alpine](https://downloads.snyk.io/cli/stable/snyk-alpine)
  * MacOS arm64: [https://downloads.snyk.io/cli/stable/snyk-macos-arm64](https://downloads.snyk.io/cli/stable/snyk-macos-arm64)
  * Linux arm64: [https://downloads.snyk.io/cli/stable/snyk-linux-arm64](https://downloads.snyk.io/cli/stable/snyk-linux-arm64)
  * Alpine arm64: [https://downloads.snyk.io/cli/stable/snyk-alpine-arm64](https://downloads.snyk.io/cli/stable/snyk-alpine-arm64)
  * For fips, add `fips` to the base URL, for example, `https://downloads.snyk.io/fips/cli/stable/snyk-linux`
  * For experimental builds, add `experimental` to the base URL, for example, `https://downloads.snyk.io/experimental/cli/stable/snyk-linux`&#x20;
* Installation methods:
  * [npm](install-or-update-the-snyk-cli/#install-the-snyk-cli-with-npm-or-yarn)
  * [Homebrew](install-or-update-the-snyk-cli/#install-with-homebrew-macos-linux)
  * [Scoop](install-or-update-the-snyk-cli/#install-with-scoop-windows)
  * [Snyk-images](install-or-update-the-snyk-cli/#snyk-cli-in-a-docker-image)

Snyk recommends opting in to a stable channel for the following reasons:

* A stable build is tested extensively over the course of eight weeks during which Snyk development teams use the CLI in the SDLC process
* Accompanying release notes help you decide which version best suits your needs

However, customers who would like to receive code changes as soon as they are merged can opt in to the preview channel. Note that Snyk does not offer support for the preview channel and expects known issues to be present in this channel.

{% hint style="info" %}
Existing Snyk customers who are opted in to the previously known latest channel will be automatically opted in to the stable channel. Snyk is mirroring the latest channel and the stable channel to avoid disruption to existing customers. However, Snyk encourages you to switch to the new channels as shown above.
{% endhint %}

## Experimental builds

{% hint style="info" %}
Experimental builds are exempt from the [12 month support policy](releases-and-channels-for-the-snyk-cli.md#support-policy), and are subject to change at any time.
{% endhint %}

Snyk deploys experimental builds across release channels for its binaries. Experimental builds are functionally the same as non-experimental builds, but can differ in their build process. This differentiation can impact runtime dependancies, for example.

Experimental builds are available by adding `experimental` to the base download URL as outlined in the [Channels](releases-and-channels-for-the-snyk-cli.md#channels) section.

### linuxstatic builds

{% hint style="info" %}
linuxstatic builds are experimental builds on the CLI side, which can be changed at any time without prior notice. We're exploring this type of build before confirming full support.
{% endhint %}

The Snyk CLI Linux binaries (`snyk-linux`, `snyk-linux-arm64`) come with a dynamically linked `nodejs` runtime. The main dependency for this runtime is the `glibc` library. This dependency is directly impacted by the customer environment and limits customers using the Snyk CLI in older Linux environments (as outlined in the [Runtime requirements](releases-and-channels-for-the-snyk-cli.md#runtime-requirements) section).

linuxstatic builds, as opposed to the regular Linux builds, are Linux binaries compiled to be self-contained, including all necessary dependencies, so they don't rely on a user's system having those libraries installed.

## Support policy

Snyk supports the latest 12 months of CLI versions in the stable channel, ensuring functionality and performance. Older versions are considered End-of-Support (EOS) and will not receive bug fixes or support for troubleshooting.

Snyk provides fixes only in new versions and cannot fix older versions. Customers must upgrade to benefit from improvements.

This policy fosters innovation while optimizing resources.

If you need help, submit a [request](https://support.snyk.io) to Snyk Support.

## Runtime requirements

This section specifies the minimum requirements for the latest CLI, covering aspects such as the operating system, dynamically linked libraries, and runtimes.

{% hint style="warning" %}
Snyk CLI is designed and expected to run on operating system environments that are actively supported by their vendors and have not reached their End of Support (EOS) or End of Life (EOL). Once an operating system environment reaches its end of support, Snyk will cease testing and providing support for Snyk CLI in that specific environment.
{% endhint %}

| OS\Architecture | amd64 (minimum requirements)                                                                  | arm64 (minimum requirements)                                                                                       |
| --------------- | --------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------ |
| macOS           | 10.13                                                                                         | 10.13                                                                                                              |
| Linux           | glibc ≥ 2.28 & OpenSSL ≥ 3.0 (Enterprise Linux 8, Ubuntu 20.04, Debian buster, other distros) | <p>glibc ≥ 2.31 &#x26; OpenSSL ≥ 3.0 </p><p>(Enterprise Linux 9, Ubuntu 20.04, Debian bullseye, other distros)</p> |
| Alpine          | musl libc >= 1.1.18 (3.7.3, other distros)                                                    | musl libc >= 1.1.18 (3.7.3, other distros)                                                                         |
| Windows         | 8.1                                                                                           | -                                                                                                                  |

## Install standalone executables from a channel

Use the `release.json` in each channel. The download links are provided here, followed by an example for the preview version on the MacOS platform:

* [https://downloads.snyk.io/cli/preview/release.json](https://downloads.snyk.io/cli/preview/release.json)
* [https://downloads.snyk.io/cli/preview/snyk-macos](https://downloads.snyk.io/cli/preview/snyk-macos)

For MacOS, download and run a preview version of the CLI in a temporary folder named `snyk-preview`. To do this, you can run the following set of commands.

<pre class="language-sh"><code class="lang-sh"><strong>mkdir snyk-preview
</strong>cd snyk-preview
curl --compressed https://downloads.snyk.io/cli/preview/snyk-macos -o snyk
chmod +x ./snyk
./snyk -version
</code></pre>

## Selecting a channel from the IDE

{% hint style="info" %}
This functionality is available in the IntelliJ IDE. Snyk is extending this capability to other supported IDEs.
{% endhint %}

The default channel for all IDEs is the stable channel.

To select a channel from the IDE, choose a CLI release channel using the dropdown, as shown in the screenshot that follows. Users can switch between channels, for example, switch to release-candidate (**rc**) to receive a hotfix.

However, Snyk also recommends the **stable** channel as the default for IDE users.

<figure><img src="../../.gitbook/assets/Choose a CLI release channel.png" alt=""><figcaption><p>Choose a CLI release channel</p></figcaption></figure>
