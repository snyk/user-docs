# Releases and channels for the Snyk CLI

This page describes Snyk CLI releases and channels, and also explains how to opt-in to a channel from the IDE.

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

Note that the preview channel contains features that are being developed. This channel may have known as well as unknown issues that are not yet officially released and supported.

* Preview: pre-release builds are deployed regularly up to multiple times a day and contain the latest changes.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}-preview
* Cadence: Varying
* Availability:
  * [https://static.snyk.io/cli/preview/](https://static.snyk.io/cli/preview/)
  * [https://static.snyk.io/fips/cli/preview/](https://static.snyk.io/cli/preview/)

{% hint style="info" %}
For instructions on installing a preview version, see [Install standalone executables from a channel.](releases-and-channels-for-the-snyk-cli.md#install-standalone-executables-from-a-channel)
{% endhint %}

### rc&#x20;

* Release candidate: pre-releases are deployed at distinct points in time and contain a version of the CLI that is expected to be promoted to stable after additional testing.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}-rc
* Cadence: every eight weeks, two weeks before a stable release (hotfix releases possible)
* Availability:
  * [https://static.snyk.io/cli/rc/](https://static.snyk.io/cli/rc/)
  * [https://static.snyk.io/fips/cli/rc/](https://static.snyk.io/cli/rc/)

### **stable**&#x20;

* Stable: builds are deployed at distinct points in time after additional testing and are considered stable.
* Version Pattern: v{MAJOR}.{MINOR}.{PATCH}
* Cadence: every eight weeks (hotfix releases possible)
* Availability:
  * [https://github.com/snyk/cli/releases/](https://github.com/snyk/cli/releases/)
  * [https://static.snyk.io/cli/stable/](https://static.snyk.io/cli/stable/)
  * [https://static.snyk.io/fips/cli/stable/](https://static.snyk.io/cli/stable/)
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

## Install standalone executables from a channel

See the `release.json` in each channel for the download links. Examples for the preview version on the MacOS platform follow:

* [https://static.snyk.io/cli/preview/release.json](https://static.snyk.io/cli/preview/release.json)
* [https://static.snyk.io/cli/preview/snyk-macos](https://static.snyk.io/cli/preview/snyk-macos)

For example, to download and run the preview release of the Snyk CLI for macOS, you could run:

```
curl --compressed https://static.snyk.io/cli/preview/snyk-macos -o snyk
chmod +x ./snyk
mv ./snyk /usr/local/bin/
```

## Selecting a channel from the IDE

{% hint style="info" %}
This functionality is available in the IntelliJ IDE. Snyk is extending this capability to other supported IDEs.
{% endhint %}

The default channel for all IDEs is the stable channel.

To select a channel from the IDE, choose a CLI release channel using the dropdown, as shown in the screenshot that follows. Users can switch between channels, for example, switch to release-candidate (**rc**) to receive a hotfix.&#x20;

However, Snyk also recommends the **stable** channel as the default for IDE users.

<figure><img src="https://lh7-us.googleusercontent.com/RPWcJ4UOtDfSqXNn0V11tljHUDmz3blBCUlK8U3uAtUaodbKi4KYLeHFrqKSdQxGjlLy8VS07DsGAgsEixHnofj6igcaSduNKgdi2GTCi9FcpfuuGha5hnyVm6xypSuh2bgmpKXBQTMyUtUNL_i3RdQ" alt="Choose a CLI release channel"><figcaption><p>Choose a CLI release channel</p></figcaption></figure>
