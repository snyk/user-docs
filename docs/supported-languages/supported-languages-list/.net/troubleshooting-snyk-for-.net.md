# Troubleshooting Snyk for .NET

## Snyk Open Source limitations for .NET

* [`Directory.Build.props`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) and [`Directory.Build.targets`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) are not supported for SCM integration. You can scan private dependencies using central package management with the Snyk CLI. You must run `dotnet restore` and then run `snyk test` with `--all-projects`, as each sub-folder will contain its own `project.assets.json` file.
* `<ProjectReference>`elements are not supported.
* Private dependency scanning is supported for SCM integrations by enabling improved .NET scanning for your Organization or Group through the [Snyk Preview](../../../snyk-platform-administration/snyk-preview.md) menu. Navigate to the [Improved .NET scanning](improved-.net-scanning.md) page for more details. You can also scan private dependencies using the Snyk CLI

If you need help, [contact Snyk Support](https://support.snyk.io).

## Snyk technical support for .NET

The following support articles are available:

* [Project import errors](https://support.snyk.io/s/article/Project-import-errors)
* [Changing .NET Target Framework not reflected in Snyk Portal](https://support.snyk.io/s/article/Changing-NET-Target-Framework-not-reflected-in-Snyk-Portal)
* [How does Snyk aggregate .NET Projects?](https://support.snyk.io/s/article/How-does-Snyk-aggregate-NET-Projects)
