# Releases and channels for the Snyk CLI

This page describes Snyk CLI releases and support policy, and also explains how to opt-in to different channels and the purpose of each channel.

## Releases

Beginning with v.1.1291.0, the Snyk CLI follows the industry standard [Semantic Versioning](https://semver.org/) three-part notation as follows.

Given a version number MAJOR.MINOR.PATCH, increment the:

1. MAJOR version when you make \[breaking] changes
2. MINOR version when you add functionality in a backward-compatible manner
3. PATCH version when you make backward-compatible bug fixes

Additional labels are added for CLI releases as needed based on the standard.

In the Snyk CLI context, Snyk defines a breaking change as a change that could break automated workflows and cause failures in your existing working setup, such as CI/CD integrations. Breaking changes will be indicated by incrementing MAJOR and mentioned in the release notes too.

Some examples of breaking changes are the following:

* Deprecating or changing output fields, field names, or environment variables
* Introducing mandatory configuration changes
* Changes in error or exit codes

## Channels

Beginning with v.1.1291.0, Snyk is providing different channels to enable customers to opt-in to a channel based on their needs and preferences.

When you select a channel, you are selecting the stability level you want to use: **preview**, **rc,** or **stable**.

### preview&#x20;

{% hint style="info" %}
Snyk offers a preview channel for those who want to test _in-progress_ features. However, keep in mind that this channel may contain bugs and is not officially supported.
{% endhint %}

A preview version is not recommended for production environments. It may contain bugs and is best tested in a local environment. For instructions on installing a preview version, see [Install standalone executables from a channel.](releases-and-channels-for-the-snyk-cli.md#install-standalone-executables-from-a-channel)

* Preview: pre-release builds are deployed regularly up to multiple times a day and contain the latest changes.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}-preview
* Cadence: Varying
* Availability:
  * [https://downloads.snyk.io/cli/preview/](https://downloads.snyk.io/cli/preview/)
  * [https://downloads.snyk.io/fips/cli/preview/](https://static.snyk.io/fips/cli/preview/)

### rc&#x20;

* Release candidate: pre-releases are deployed at distinct points in time and contain a version of the CLI that is expected to be promoted to stable after additional testing.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}-rc
* Cadence: every eight weeks, two weeks before a stable release (hotfix releases possible)
* Availability:
  * [https://downloads.snyk.io/cli/rc/](https://downloads.snyk.io/cli/rc/)
  * [https://downloads.snyk.io/fips/cli/rc/](https://downloads.snyk.io/fips/cli/rc/)

### **stable**&#x20;

* Stable: builds are deployed at distinct points in time after additional testing and are considered stable.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}
* Cadence: every eight weeks (hotfix releases possible)
* Availability:
  * [https://github.com/snyk/cli/releases/](https://github.com/snyk/cli/releases/)
  * [https://downloads.snyk.io/cli/stable/](https://downloads.snyk.io/cli/stable/)
  * [https://downloads.snyk.io/fips/cli/stable/](https://downloads.snyk.io/fips/cli/stable/)
  * npm
  * brew
  * scoop
  * Snyk-images

Snyk recommends opting in to a stable channel for the following reasons:

* A stable build is tested extensively over the course of eight weeks during which Snyk development teams use the CLI in the SDLC process
* Accompanying release notes help you decide which version best suits your needs

However, customers who would like to receive code changes as soon they are merged can opt in to the preview channel. Note that Snyk does not offer support for the preview channel and expects known issues to be present in this channel.

{% hint style="info" %}
Existing Snyk customers who are opted in to the previously known latest channel will be automatically opted in to the stable channel. Snyk is mirroring the latest channel and the stable channel to avoid disruption to existing customers. However, Snyk encourages you to switch to the new channels as shown above.
{% endhint %}

## Support policy

{% hint style="info" %}
This policy will be effective beginning on June 24, 2025.
{% endhint %}

Snyk supports the latest 12 months of CLI versions in the stable channel, ensuring functionality and performance. Older versions are considered End-of-Support (EOS) and will not receive bug fixes or support for troubleshooting.&#x20;

Snyk provides fixes only in new versions and cannot fix older versions. Customers must upgrade to benefit from improvements.&#x20;

This policy fosters innovation while optimizing resources.

If you need help, submit a [request](https://support.snyk.io) to Snyk Support.

## Install standalone executables from a channel

Use the`release.json` in each channel. The download links are provided here, followed by an example for the preview version on the MacOS platform:

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

To select a channel from the IDE, choose a CLI release channel using the dropdown, as shown in the screenshot that follows. Users can switch between channels, for example, switch to release-candidate (**rc**) to receive a hotfix.&#x20;

However, Snyk also recommends the **stable** channel as the default for IDE users.

<figure><img src="../.gitbook/assets/Screenshot 2024-09-02 at 10.32.41.png" alt="Choose a CLI release channel"><figcaption><p>Choose a CLI release channel</p></figcaption></figure>
