# Snyk Error Codes

The error codes in the table below describe the codes that you may encounter while working with the [Snyk API](../snyk-api-info/README.md) or [CLI](../snyk-cli/README.md). When errors are encountered using the API, they will also have an appropriate [HTTP status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). If you encounter errors without an error code, use the HTTP status code to determine the appropriate action.

---
# [Snyk](https://docs.snyk.io/introducing-snyk)
### [SNYK-0001](#snyk-0001)

#### Service temporarily throttled

The request rate limit has been exceeded. Wait a few minutes, then try again.

**HTTP Status:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)


### [SNYK-0002](#snyk-0002)

#### Server error response

The server doesn't recognize the request method, or it cannot fulfill it. Review the request and try again.

**HTTP Status:** [501](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501)

**Help Links:**
- [https://docs.snyk.io/snyk-api-info](https://docs.snyk.io/snyk-api-info)

### [SNYK-0003](#snyk-0003)

#### Client request cannot be processed

The server cannot process the request due to a client error, such as malformed request syntax, size too large, invalid request message framing, or deceptive request routing. Review the request and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-0004](#snyk-0004)

#### Server communication error

The server timed out during the request. Check Snyk status, then try again.

**HTTP Status:** [504](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504)

**Help Links:**
- [https://status.snyk.io/](https://status.snyk.io/)

### [SNYK-0005](#snyk-0005)

#### Authentication error

Authentication credentials not recognized, or user access is not provisioned. Revise credentials and try again, or request access from your Snyk administrator.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)


### [SNYK-9999](#snyk-9999)

#### Request not fulfilled due to server error 

The server cannot process the request due to an unexpected error. Check Snyk status, then try again.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://status.snyk.io/](https://status.snyk.io/)

---
# Code
### [SNYK-CODE-0001](#snyk-code-0001)

#### Analysis file count limit exceeded

This error occurs when the analysis target has a supported file count which exceeds current system limits.

To reduce the file count, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code)
- [https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process](https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process)
- [https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli](https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli)

### [SNYK-CODE-0002](#snyk-code-0002)

#### Analysis result size limit exceeded

This error occurs when the analysis target generates a result with a byte size that exceeds current system limits.

To reduce the overall result size, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process](https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process)
- [https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli](https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli)

---
# OpenAPI
### [SNYK-OPENAPI-0001](#snyk-openapi-0001)

#### Bad request

The server cannot process the request due to invalid or corrupt data. Review the request, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api ](https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api )

### [SNYK-OPENAPI-0002](#snyk-openapi-0002)

#### Forbidden

Access to the requested resource is forbidden. Review the request, then try again.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)


### [SNYK-OPENAPI-0003](#snyk-openapi-0003)

#### Not acceptable

The server cannot provide a response that matches the provided accept headers. Review the request, then try again.

**HTTP Status:** [406](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406)


### [SNYK-OPENAPI-0004](#snyk-openapi-0004)

#### Not found

The server cannot find the requested resource. Review the request, then try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OPENAPI-0005](#snyk-openapi-0005)

#### Method not allowed

The target endpoint does not support your request method. Review the request, then try again.

**HTTP Status:** [405](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405)


### [SNYK-OPENAPI-0006](#snyk-openapi-0006)

#### Request entity too large

The request entity exceeds server limitations. Reduce the size of the request entity, then try again.

**HTTP Status:** [413](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413)


### [SNYK-OPENAPI-0007](#snyk-openapi-0007)

#### Unauthorized

The request lacks authentication credentials for the requested resource. Ensure you are sending valid credentials, then try again.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**Help Links:**
- [https://docs.snyk.io/snyk-api-info/authentication-for-api](https://docs.snyk.io/snyk-api-info/authentication-for-api)

### [SNYK-OPENAPI-0008](#snyk-openapi-0008)

#### Unsupported media type

The media format of the request is not supported. Change media format, then try again.

**HTTP Status:** [415](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415)


---
# [Open Source Languages & Package Managers](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source)
### [SNYK-OS-0001](#snyk-os-0001)

#### Unable to parse manifest file

The provided manifest file could not be parsed as it has invalid syntax or does not match the expected schema. Review the manifest file, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-0002](#snyk-os-0002)

#### Unable to parse lock file

The provided lock file could not be parsed as it has invalid syntax or does not match the expected schema. Review the lock file, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-0003](#snyk-os-0003)

#### Unknown dependency version

Dependency version could not be resolved.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360001373178-Could-not-determine-version-for-dependencies](https://support.snyk.io/hc/en-us/articles/360001373178-Could-not-determine-version-for-dependencies)

### [SNYK-OS-0004](#snyk-os-0004)

#### Missing required request header

The server encountered a request that is missing a mandatory request header.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0005](#snyk-os-0005)

#### Payload missing required elements

The server could not process the request.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0006](#snyk-os-0006)

#### Files cannot be processed

The dependency service could not process the files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0007](#snyk-os-0007)

#### Cannot get file from source

Could not get the file from the source URL.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-0008](#snyk-os-0008)

#### Missing environment variable

The server encountered a critical operation that requires a specific environment variable, but the variable is not set or is not accessible within the current environment.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-DOTNET-0001](#snyk-os-dotnet-0001)

#### Unsupported manifest file type for remediation

The provided manifest file is not supported by Snyk for .NET.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net)

### [SNYK-OS-DOTNET-0002](#snyk-os-dotnet-0002)

#### Target framework not supported

The provided manifest file defines a `<TargetFramework>` or `<TargetFrameworks>` that is not currently supported by Snyk's .NET scanning solution.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-DOTNET-0003](#snyk-os-dotnet-0003)

#### Your C# code is missing a static Main function

This error occurs when no static Main method with a correct signature is found in the code that produces an executable file. 
It also occurs if the entry point function, `Main`, is defined with the wrong case, such as lower-case main.

In order to fix this issue, ensure that your program has a .cs file that contains a main function, such as
```c#
namespace Example
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("hello world");
        }
    }
}
```

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://learn.microsoft.com/en-us/dotnet/csharp/misc/cs5001](https://learn.microsoft.com/en-us/dotnet/csharp/misc/cs5001)

### [SNYK-OS-DOTNET-0004](#snyk-os-dotnet-0004)

#### The dotnet CLI is unable to generate a self-contained binary

This error occurs when running `dotnet publish --sc --framework <your-target-framework>` fails to generate a 
self-contained binary. Snyk needs to run this command in order to adequately determine the dependency tree for your project. If this command fails, Snyk cannot continue.

Steps to determine why this happened:

* Checkout a clean version of your project in a temporary folder
* Run `dotnet publish --sc --framework <your-target-framework> ` on your project, and confirm this step fails.

If this step is successful locally, it is possible that Snyk is running another version of the .NET SDK. To tell Snyk which version of the .NET SDK to use, consider using the [global.json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json) solution provided by Microsoft.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/](https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/)
- [https://learn.microsoft.com/en-us/dotnet/core/tools/global-json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json)
- [https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L67](https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L67)

### [SNYK-OS-GO-0001](#snyk-os-go-0001)

#### Failed to access private module

Snyk could not access the private modules within your go.mod files.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go)

### [SNYK-OS-GO-0002](#snyk-os-go-0002)

#### Go mod file not found

A go.mod file was not found in the current directory or any parent directory.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go)

### [SNYK-OS-GO-0003](#snyk-os-go-0003)

#### OAuth re-authorization required

Your code is cloned on an isolated environment using Git as it is required by Snyk to analyze its dependencies.

Your Organization has enabled or enforced SAML SSO after you authorized Snyk to access your code, and a re-authentication is therefore required.

The error you're seeing is usually reproducible by attempting to do a `git clone` of your repository with incorrectly configured credentials.
Verify your authentication configuration with your Git cloud provider and try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso](https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso)

### [SNYK-OS-GO-0004](#snyk-os-go-0004)

#### Your project repository is missing required files

Generating the dependency graph requires Snyk to run go list `go list -deps -json` inside the project. If the operation fails, creating a full dependency graph cannot continue.  

This error means that you need some cleanup, such as `go mod tidy`) or your project deployment process contains a code generation step such as `protobuf` or similar that is not currently supported by Snyk. 

To verify if this is the case, clone your project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. 

If Snyk cannot process your code successfully, insert the Snyk CLI as part of your deployment pipeline.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://docs.snyk.io/snyk-cli](https://docs.snyk.io/snyk-cli)
- [https://github.com/snyk/snyk-go-plugin](https://github.com/snyk/snyk-go-plugin)
- [https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go](https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go)

### [SNYK-OS-GO-0005](#snyk-os-go-0005)

#### Your project repository has inconsistent vendoring information

Generating the dependency graph requires Snyk to run `go list -deps -json` inside the project. If the operation fails, creating a full dependency graph cannot continue.  

This error means that there is inconsistency between your `vendor/modules.txt` file and your `go.mod` file. To remediate, you need to:

* `go mod vendor`
* `go mod tidy`

Next, commit those changes to your repo. Snyk does not manipulate with your code on our end by design, which is why this is not done automatically.

To verify if this is the case, clone your project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. 
Then try and run the above mentioned commands and see if your SCM system reports changes in files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://go.dev/ref/mod#go-mod-vendor](https://go.dev/ref/mod#go-mod-vendor)

### [SNYK-OS-GO-0006](#snyk-os-go-0006)

#### Unsupported external file generation

Snyk currently does not support external file generation in your project. This limitation is due to Snyk's lack of visibility into the third-party generator tools you may be using and the specific commands required to generate these files.

Snyk can only work with the files available in your repository and does not have insight into the generation process for external files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-GO-0007](#snyk-os-go-0007)

#### Unable to access private dependencies

The Go tool encountered a `DepsError` while trying to download a private dependency. Private repositories that are not accessible to the public internet and are not available on the official Go proxy mirror are cloned with a version control system and built on demand. 
This requires the VCS to have the correct access rights to that repository.

Snyk supports private repositories that are hosted in the same Organization and on the same Project that is scanned for vulnerabilities. The authentication to the private repository is the same as the authentication used to integrate that repository with Snyk. 

This error appears when the authorization credentials do not allow access to the requested private dependency. 

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://go.dev/ref/mod#vcs](https://go.dev/ref/mod#vcs)

### [SNYK-OS-MAVEN-0001](#snyk-os-maven-0001)

#### Missing property

The required property is missing from the pom object.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0002](#snyk-os-maven-0002)

#### Unable to resolve value for property

The targeted property could not be resolved with a valid value.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0003](#snyk-os-maven-0003)

#### Unable to resolve version for property

The targeted property could not be resolved with a valid version.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-MAVEN-0004](#snyk-os-maven-0004)

#### Cyclic property detected in POM file

There is circular dependency among properties in the Maven project's configuration file (POM), preventing proper resolution and causing an error.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0005](#snyk-os-maven-0005)

#### Error parsing the XML file

There is an error parsing the XML file. This could be referring to either pom.xml or maven-metadata.xml.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0006](#snyk-os-maven-0006)

#### Invalid coordinates provided

The coordinates provided for a project were invalid.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0007](#snyk-os-maven-0007)

#### Skipping group

Skipping a specific groupId starting due to remapped coordinates.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0008](#snyk-os-maven-0008)

#### Pom file not found

The pom file was not found in Maven repository.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0009](#snyk-os-maven-0009)

#### Missing project from POM

A project element is missing from POM.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0010](#snyk-os-maven-0010)

#### Cannot resolve the target POM from the input XML

Cannot resolve the targeted POM from the input XML.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0011](#snyk-os-maven-0011)

#### Cannot resolve the target POM from the repository

Cannot resolve the targeted POM from the repository.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OS-MAVEN-0012](#snyk-os-maven-0012)

#### Cannot get the build file repository

Cannot get the build file repository.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OS-MAVEN-0013](#snyk-os-maven-0013)

#### Unable to create hosted git info

Cannot create source URL.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-MAVEN-0014](#snyk-os-maven-0014)

#### No released version for versions range

There was no version released for the specified versions range.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0015](#snyk-os-maven-0015)

#### Source is not supported

The source used is not supported by fetcher. The supported sources are: github, bitbucket, gitlab.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0016](#snyk-os-maven-0016)

#### Timeout when processing the dependency tree

There was an timeout when processing the dependency tree.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0001](#snyk-os-nodejs-0001)

#### No repository found for A NPM package

No repository found for the NPM package.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0002](#snyk-os-nodejs-0002)

#### Could not parse NPM registry URL

Could not parse NPM registry URL.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0003](#snyk-os-nodejs-0003)

#### Could not find a broker resolved URL

Could not find a broker resolved URL.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0004](#snyk-os-nodejs-0004)

#### Unable to replace broker URL

Unable to replace all broker urls in lock file.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0005](#snyk-os-nodejs-0005)

#### Bad NPM version

The NPM version is not supported.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0006](#snyk-os-nodejs-0006)

#### Unknown blob encoding on Github

Unknown blob encoding on Github.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0007](#snyk-os-nodejs-0007)

#### No result from forked process

No result from forked process.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-NODEJS-0008](#snyk-os-nodejs-0008)

#### Child Process Execution Error

The child process encountered an error during execution.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-NODEJS-0009](#snyk-os-nodejs-0009)

#### No valid package upgrades

The system attempted to find valid upgrades for the packages specified in the lock file, but none were available.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0010](#snyk-os-nodejs-0010)

#### No dependency updates

There are no available updates for the dependencies.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0011](#snyk-os-nodejs-0011)

#### Could not parse JSON file

An error occurred while attempting to parse a JSON file.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0012](#snyk-os-nodejs-0012)

#### Could not Base64 encode

An error occurred while attempting to perform Base64 encoding.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0013](#snyk-os-nodejs-0013)

#### Could not Base64 decode

An error occurred while attempting to perform Base64 decoding.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-NODEJS-0014](#snyk-os-nodejs-0014)

#### Missing supported file

Could not find supported file.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-NODEJS-0015](#snyk-os-nodejs-0015)

#### Invalid configuration

The configuration parameter does not meet the expected data type. Please ensure the provided value is of the correct data type.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-PIP-0001](#snyk-os-pip-0001)

#### Unsupported manifest file type for remediation

The provided requirements file is not supported by Snyk for Python.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/python](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/python)

### [SNYK-OS-PIP-0002](#snyk-os-pip-0002)

#### Received more manifests than expected

Too many manifest files were provided in the request body.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-PIP-0003](#snyk-os-pip-0003)

#### Failed to apply dependency updates

An error occurred while updating dependencies.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-PIP-0004](#snyk-os-pip-0004)

#### No matching distribution found for one or more of the packages

At least one of the packages requires a Python version that doesn't match the one used in the project scan.
Make sure to select a suitable Python version from the organization Python language settings.
Alternatively, add a `.snyk` file for Python version selection override.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


---
# Builds
### [SNYK-OS-8001](#snyk-os-8001)

#### Invalid request

The provided request payload is not valid for the selected ecosystem. Please review the API documentation.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://apidocs.snyk.io/](https://apidocs.snyk.io/)

### [SNYK-OS-8002](#snyk-os-8002)

#### Build environment not found

The build environment for the provided context could not be found. Please ensure you have created the build environment first.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OS-8003](#snyk-os-8003)

#### Unsupported Ecosystem

The language or package manager is not supported. Please refer to the supported package managers in the links.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source)

---
# SbomExport
### [SNYK-OS-9000](#snyk-os-9000)

#### SBOM generation export server error

An unexpected error occurred during the SBOM generation. Review the request, then try again. If the error persists, contact Snyk Support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-9001](#snyk-os-9001)

#### Dependency graph error

An unexpected dependency graph error occurred. Review the request, then try again. If the error persists, contact Snyk Support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-9002](#snyk-os-9002)

#### Error parsing dependency graph

The dependency graph cannot be parsed due to an unexpected error. Review the request, then try again. If the error persists, contact Snyk Support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-9003](#snyk-os-9003)

#### SBOM not supported due to project type

Only SBOMs for Snyk Open Source or Snyk Container projects are supported.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OS-9004](#snyk-os-9004)

#### SBOM not supported

Only SBOMs for open source projects are supported (Snyk Open Source).

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OS-9005](#snyk-os-9005)

#### Dependency graph request cannot be processed

The server cannot process the request due to incomplete data. Review the request, then try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OS-9006](#snyk-os-9006)

#### Authorization failed due to missing API token

The API token is misconfigured or expired. Configure or generate the API token, then try again.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**Help Links:**
- [https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens](https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens)

### [SNYK-OS-9007](#snyk-os-9007)

#### Client request cannot be processed

The body of the request is empty. Review the request, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-9008](#snyk-os-9008)

#### Invalid dependency graph

The supplied dependency graph was not valid. Review the request, then try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


---
# [Open Source Unmanaged](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files)
### [SNYK-OSJVM-001](#snyk-osjvm-001)

#### Maven search service unavailable

The upstream Maven search service is not available.

**HTTP Status:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

**Help Links:**
- [https://search.maven.org](https://search.maven.org)
- [https://status.maven.org](https://status.maven.org)

### [SNYK-OSJVM-002](#snyk-osjvm-002)

#### SHA1 not found

Unable to find the coordinates for the provided SHA1. Please verify the data you are sending and try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**
- [https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files)

---
# [PURL Vulnerabilities](https://docs.snyk.io/scan-applications/snyk-open-source/manage-vulnerabilities/snyk-vulnerability-database#about-the-vulnerability-database)
### [SNYK-OSSI-1040](#snyk-ossi-1040)

#### Your Organisation is not authorized to perform this action

You likely don’t have access to the features in Beta. To get access, you can request access to features in Beta through your account manager or team.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)


### [SNYK-OSSI-1050](#snyk-ossi-1050)

#### Authorization request failure

Unexpected error when authenticating. Try again, and if the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-2010](#snyk-ossi-2010)

#### Invalid purl

Make sure that the purl is valid. See the Package URL specification link for further information.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst](https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst)

### [SNYK-OSSI-2011](#snyk-ossi-2011)

#### Namespace not specified

You have requested a package type that requires a namespace (e.g. maven group id). Provide the namespace to retrieve the package.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst](https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst)

### [SNYK-OSSI-2020](#snyk-ossi-2020)

#### Unsupported ecosystem

The package type is not supported. Check the List issues for a package in Snyk API.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-2021](#snyk-ossi-2021)

#### Purl components required

A list of components of the purl spec is required. The purl did not specify all the required components. Please add the missing components to the purl and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-2022](#snyk-ossi-2022)

#### Unsupported purl components

Remove the unsupported component and retry the request.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-2030](#snyk-ossi-2030)

#### Requested package not found

The package you specified in the purl cannot be found in the vulnerability database. Check the package name, ecosystem, and version, then try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OSSI-2031](#snyk-ossi-2031)

#### Vulnerability service not available

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)


### [SNYK-OSSI-2032](#snyk-ossi-2032)

#### This issue is unexpected and the service should recover quickly if not please contact support

An unexpected error occurred. Please try again, and if you continue to experience issues please contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-2033](#snyk-ossi-2033)

#### This issue is unexpected and the service should recover quickly if not please contact support

An unexpected error occurred with the vulnerability service. Please try again, and if you continue to experience issues please contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-2040](#snyk-ossi-2040)

#### Request not processed due to unexpected error

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-2041](#snyk-ossi-2041)

#### Invalid pagination parameters

The pagination limit is > 1 and ≤ 1000, and the offset is ≥0.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-2042](#snyk-ossi-2042)

#### purls exceed limit

The number of purls sent in the request exceeds the limit of 1000 set by the service.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-2043](#snyk-ossi-2043)

#### Number of issues exceeds limit

The number of issues found for the provided purls exceeds the limit defined by the API. Reduce the number of purls sent in a single request.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-2044](#snyk-ossi-2044)

#### Expected distro to be present

The given Package URL does not have a required distro qualifier.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-containers/how-snyk-container-works/supported-operating-system-distributions#debian](https://docs.snyk.io/scan-containers/how-snyk-container-works/supported-operating-system-distributions#debian)

### [SNYK-OSSI-2045](#snyk-ossi-2045)

#### Unsupported Debian distro

This Debian distro is currently not supported.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-2046](#snyk-ossi-2046)

#### Expected namespace to be present

The given Package URL does not have a required namespace.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-2047](#snyk-ossi-2047)

#### Unsupported vendor

The given Package URL does not contain a supported vendor. Please use one of the listed vendors and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


---
# OpenSourceProjectIssues
### [SNYK-OSSI-OSPI-1001](#snyk-ossi-ospi-1001)

#### Invalid request

Check the body of your request and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-OSPI-1002](#snyk-ossi-ospi-1002)

#### Unable to return valid API response

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-OSPI-2001](#snyk-ossi-ospi-2001)

#### Failed to process data

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-OSPI-3001](#snyk-ossi-ospi-3001)

#### Failed to store issue data

Check inputs and then try again. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-OSPI-4001](#snyk-ossi-ospi-4001)

#### Internal server error

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


---
# Open Source Project Snapshots
### [SNYK-OSSI-OSPSS-1001](#snyk-ossi-ospss-1001)

#### Invalid request

Check the body of your request and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OSSI-OSPSS-1002](#snyk-ossi-ospss-1002)

#### Unable to return valid API response

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-OSPSS-2001](#snyk-ossi-ospss-2001)

#### Failed to process data

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-OSPSS-3001](#snyk-ossi-ospss-3001)

#### Failed to store snapshot data

Check inputs and then try again. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OSSI-OSPSS-4001](#snyk-ossi-ospss-4001)

#### Internal server error

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


---
# PRChecks
### [SNYK-PR-CHECK-0001](#snyk-pr-check-0001)

#### Error reading manifest

Snyk failed to read 1 or more manifest files.
Sometimes things go wrong: a flaky connection, 3rd party services go down and Snyk is unable to read the files needed in order to test your project. 

If this happens, you could try:

- Opening and re-opening your Pull Request / Merge Request, to kick off a new test
- Removing and re-adding the repo to Snyk

Ultimately, you should contact support@snyk.io if the issue persists

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360000910517-Failed-to-read-manifest-file](https://support.snyk.io/hc/en-us/articles/360000910517-Failed-to-read-manifest-file)

### [SNYK-PR-CHECK-0002](#snyk-pr-check-0002)

#### Manifest not found

Snyk uses your project manifest file to analyze your projects for vulnerabilities. When you import a project for monitoring, Snyk scans the project to locate the manifest file and then remembers where that file is. 
When a project manifest file is moved or deleted, we still try to look for in it in the last known location in order to run tests on commit statuses. If we can't find the file, this error can occur.

If this happens, you could try the following:
1. Delete the matching project from your account in the Snyk app (UI or CLI).
2. Now import the same project from scratch.

As during the original import, Snyk scans the project and locates the manifest file.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360000910537-Manifest-not-found](https://support.snyk.io/hc/en-us/articles/360000910537-Manifest-not-found)

### [SNYK-PR-CHECK-0003](#snyk-pr-check-0003)

#### Rate limit hit while testing project

Snyk makes requests to your SCM when testing a project, in order to analyze your projects for vulnerabilities. If we need to make a lot of requests in a short time period, we may encounter third party rate limits, and this error can occur.

If you receive any of these errors, try re-running the tests, by closing and reopening the pull request.

**HTTP Status:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)


### [SNYK-PR-CHECK-0004](#snyk-pr-check-0004)

#### Out of Sync Error

Sometimes a project may become out of sync between the lockfile and the manifest file. This might happen if the package.json is modified or updated but the lockfile is not. 

This can be resolved by ensuring the lockfile and manifest file are correctly synced, by executing npm install or yarn install.

In some cases, it may be necessary to delete the node_modules folder and the package-lock.json and run npm install again to force a full reinstall. 

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360000912457-Out-of-sync-manifest-lockfile-in-the-project](https://support.snyk.io/hc/en-us/articles/360000912457-Out-of-sync-manifest-lockfile-in-the-project)

### [SNYK-PR-CHECK-0005](#snyk-pr-check-0005)

#### Failed determining project target

An internal error occurred, whereby Snyk was unable to determine the correct target for a given project in your PR Check.

If you receive this error, try re-running the tests, by closing and reopening the pull request.

Ultimately, you should contact support@snyk.io if the issue persists.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-PR-CHECK-0006](#snyk-pr-check-0006)

#### Failed to complete the test

A "Failed to complete testing check status" appears in your commit checks when an unknown error occurs while Snyk was trying to test your projects for vulnerabilities or license issues.

If you receive this error, try re-running the tests, by closing and reopening the pull request.

Ultimately, you should contact support@snyk.io if the issue persists.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360004358517-Unknown-PR-test-error](https://support.snyk.io/hc/en-us/articles/360004358517-Unknown-PR-test-error)

### [SNYK-PR-CHECK-0007](#snyk-pr-check-0007)

#### Failed to fetch merge commit SHA

In order for snyk test to run, we need the merge commit SHA from the GitHub. For some reason, we couldn’t get it.

Try closing and then reopening the pull request, or you can Skip the Pull Request Check if it is consistent.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360005281837](https://support.snyk.io/hc/en-us/articles/360005281837)

### [SNYK-PR-CHECK-0008](#snyk-pr-check-0008)

#### Merge conflict error

Merge Conflict Error is not a Snyk specific issue but rather some issues on your SCM environment. As an example, merge conflicts could happen when people make different changes to the same line of the same file, or when one person edits a file and another person deletes the same file.

To resolve this, you might need to figure out all the merge conflicts on your SCM environment and resolve them to fully remediate these types of errors on Snyk. As a note, this cannot be modified/changed on Snyk's side.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360005281098](https://support.snyk.io/hc/en-us/articles/360005281098)

### [SNYK-PR-CHECK-0009](#snyk-pr-check-0009)

#### Failed to detect issues

Snyk is always trying to check for new issues and vulnerabilities to keep you safe. We do so by testing on your code on webhook Pull Request events and Push events.

Occasionally you might see a "Failed to detect issues" commit status which may block your PR. This means that we tried to run a test against your changes but unfortunately something went wrong / we encountered an internal problem. If this happens to you try recreating the pull request and if it still occurs reach out and let us know which user, organization and project and commit sha you experienced the issue with on support@snyk.io

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360000920678-Failed-to-detect-issues](https://support.snyk.io/hc/en-us/articles/360000920678-Failed-to-detect-issues)

### [SNYK-PR-CHECK-0010](#snyk-pr-check-0010)

#### No valid credentials to process PR check

Snyk uses credentials configured on your integration to test your code and to update your PR Check.

If this error occurs, please ensure your integration and credentials are correctly set up, by following the instructions for your SCM here: https://docs.snyk.io/integrate-with-snyk/git-repository-scm-integrations

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)


### [SNYK-PR-CHECK-0011](#snyk-pr-check-0011)

#### Failed to generate a commit status

Snyk is always trying to check for new issues and vulnerabilities to keep you safe. We do so by testing on your code on webhook Pull Request events and Push events.

Occasionally you might see a "Failed to generate a commit status" which may block your PR. This means that we tried to run a test against your changes but unfortunately something went wrong / we encountered an internal problem. If this happens to you try recreating the pull request and if it still occurs reach out and let us know which user, organization and project and commit sha you experienced the issue with on support@snyk.io

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


---
# Fix
### [SNYK-PR-TEMPLATE-0001](#snyk-pr-template-0001)

#### Failed to get pull request attributes

Snyk could not get the custom pull request template attributes, using the given variables and the fetched pr template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0002](#snyk-pr-template-0002)

#### Not found

We could not find your pull request template, have you created one yet? Please check the attached link for instructions on how to setup your pull request template.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0003](#snyk-pr-template-0003)

#### Failed to compile pull request template

Could not compile your customize pull request template. Please check for syntax errors using the Snyk variables inside the template.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0004](#snyk-pr-template-0004)

#### Failed to parse pull request attributes

Snyk could not parse the custom pull request template, using the given variables and assigning them to the fetched pr template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0005](#snyk-pr-template-0005)

#### Failed to load YAML file after substituting Snyk variables

Could not load YAML file after substituting Snyk variables into the custom PR template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0006](#snyk-pr-template-0006)

#### Failed to generate hash for custom PR template

Snyk could not generate hash using the customer PR files and projects vulnIds.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0007](#snyk-pr-template-0007)

#### Unable to create pull request template

Snyk could not create pull request template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0008](#snyk-pr-template-0008)

#### Unable to get pull request template

Snyk could not get pull request template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0009](#snyk-pr-template-0009)

#### Unable to delete pull request template

Snyk could not delete pull request template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0010](#snyk-pr-template-0010)

#### Invalid payload

The pull request template payload is invalid.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-PR-TEMPLATE-0011](#snyk-pr-template-0011)

#### Failed to load JSON file after substituting Snyk variables

Could not load JSON file after substituting Snyk variables into the custom PR template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

--- Generated at 2023-11-29T11:48:49.798Z
