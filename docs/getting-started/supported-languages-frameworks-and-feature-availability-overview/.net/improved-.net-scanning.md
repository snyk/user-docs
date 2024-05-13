# Improved .NET scanning

{% hint style="warning" %}
**Release status**&#x20;

Improved Snyk Open Source scanning for NuGet .NET applications is now available in Early Access.&#x20;

Whilst in Early Access, this feature can potentially cause scans to fail or produce different results. Snyk recommends gradually enabling this feature, for example, starting with a subset of your Organizations.

You can report issues by submitting a [support ticket](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

## Limitations of existing solution

The existing .NET scanning solution for SCM integration in the Snyk Web UI and CLI uses two [different approaches](../../supported-languages-and-frameworks/.net/best-practices-for-.net.md#dependency-analysis). The Dependency Analysis for SCM integrations can produce [false positives](../../supported-languages-and-frameworks/.net/best-practices-for-.net.md#tackling-vulnerabilities-from-runtime-dependencies) that have no remediation available and must be manually ignored.&#x20;

If you use the CLI to scan Projects, you can expect more accurate results compared to importing the same Project using an SCM integration. You can [encounter errors](../../supported-languages-and-frameworks/.net/troubleshooting-.net.md#not-supported-in-snyk-open-source-for-.net) when scanning Projects that use specific .NET features.

## Scanning improvements

Snyk improved the .NET scanning process to ensure that dependency results are consistent across the CLI and SCM integrations. This update also eliminates false positives from runtime dependencies that were previously displayed in the UI. The approach involves using the internal workings of the .NET ecosystem.

The improved .NET scanning also provides the capability of scanning any Project that can be successfully built by the `dotnet` SDK itself, removing the previous limitations of not being able to scan Projects using certain .NET features, such as `.props` files, `global.json`, or Central Package Management.

### Git repository integration in the Snyk Web UI

The .NET scanning improvements are available when importing projects using Git repository integrations.

Follow these steps to enable the improvements:

1. Activate improved .NET scanning for your Organization or Group through the [Snyk Preview](../../../snyk-admin/snyk-preview.md) menu.
2. Re-import any repositories with .NET applications; **re-testing existing Projects is not sufficient.**

#### Private package repository integration

{% hint style="info" %}
This feature is currently available only to customers on the [enterprise plan](https://snyk.io/plans/). For more information, see [Enterprise setup](https://docs.snyk.io/enterprise-configuration).
{% endhint %}

Since the improved .NET solution will build your .NET Project, Snyk requires access to any private NuGet repositories.&#x20;

The recommended approach is to use [`nuget.config`](https://learn.microsoft.com/en-us/nuget/reference/nuget-config-file) files in concurrence with registering the credentials in Snyk's NuGet private package repository integration (**Settings > Integrations > NuGet**).

If you are not using `nuget.config`, but another way of informing the .NET ecosystem of where to look for private packages, Snyk will attempt to add all private NuGet repository credentials defined in the private package repository integration as a `dotnet nuget` source before restoring the Project.

#### Limitations

* Brokered connections to private package repositories are not currently supported&#x20;
* `Directory.Build.props` , `global.json` and other .NET-specific manifest files are supported, but the file names must use upper and lower-case as Microsoft [describes](https://learn.microsoft.com/en-us/visualstudio/msbuild/customize-by-directory?view=vs-2022#directorybuildprops-and-directorybuildtargets).&#x20;
* For `global.json`, Snyk does not support all `major.minor.patch` versions that are currently supported by Microsoft, only a subset thereof. For more information, see this [error code](https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0008).

### Snyk CLI

The .NET scanning improvements are also available in the Snyk CLI.

Follow these steps to enable the improvements:

1. [Install](../../../snyk-cli/install-or-update-the-snyk-cli/) the latest version of the CLI.
2. Scan your .NET Projects using `snyk test` or `snyk monitor` as usual, but add the  `--dotnet-runtime-resolution` option.

{% hint style="info" %}
Ensure that `dotnet restore` has been run on your Project before you run `snyk test`, so the Snyk CLI can detect .NET.
{% endhint %}

The `--dotnet-runtime-resolution` option works with`--all-projects`.

Example: `snyk test --dotnet-runtime-resolution`

#### Multiple target frameworks support

The Improved .NET scanning for the CLI supports multiple multiple target frameworks in your build manifests.&#x20;

By default, the solution scans all target frameworks. If you want to scan individual target frameworks, you must add the `--dotnet-target-framework=<targetFramework>` option to your CLI command.

Example: `snyk test --dotnet-runtime-resolution --dotnet-target-framework=net8.0`

## Supported .NET versions

The improved .NET scanning feature supports all versions of .NET core that are [currently supported](https://dotnet.microsoft.com/en-us/platform/support/policy/dotnet-core) by Microsoft.

If you import an unsupported .NET Project using an SCM integration, the improved .NET scanning feature will not be enabled and will fall back to the legacy scanning method.

