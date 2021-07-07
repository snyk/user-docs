# How do I scan .NET project in CLI?

##  How do I scan .NET project in CLI?

In order to scan a .NET project, ensure the project is built before the Snyk scan takes place. Additionally, consult run snyk help to view unique .NET parameters, which support complex projects as well.

For instance:

```text
dotnet build proj/myexample.csproj
snyk test --file=proj/myexample.sln
```

This is particularly important to do when using Snyk in a CI/CD environment because the scan completes locally but not in a pipeline test.

**Note:** Snyk does not support the ".net website" type of .net application. This is different from Snyk "webapplications" and is missing some key files we require.

See [.NET for Snyk](https://support.snyk.io/hc/en-us/articles/360004519138--NET-for-Snyk) for more information.

