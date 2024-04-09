# Troubleshooting .NET

## Snyk Open Source limitations for .NET

* [`Directory.Build.props`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) and [`Directory.Build.targets`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) are not  supported for SCM integration. You can scan private dependencies using central package management using the Snyk CLI. You must run `dotnet restore` and then run `snyk test` with `-all-projects`, as each sub-folder will contain its own `project.assets.json` file.
* `<ProjectReference>`elements are not supported.
* Private dependency scanning is not supported for SCM integration. You can scan private dependencies using the Snyk CLI.

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;

## Support articles for .NET

* [Project import errors](https://support.snyk.io/hc/en-us/articles/360001373118-Project-import-errors)
* [Changing .NET Target Framework not reflected in Snyk Portal](https://support.snyk.io/hc/en-us/articles/360001421457-Changing-NET-Target-Framework-not-reflected-in-Snyk-Portal)
* [How does Snyk aggregate .NET Projects?](https://support.snyk.io/hc/en-us/articles/360002941078-How-does-Snyk-aggregate-NET-Projects-)
