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

The server doesn’t recognize the request method, or it cannot fulfill it. Review the request and try again.

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
# OpenAPI
### [SNYK-OPENAPI-0001](#snyk-openapi-0001)

#### Bad request

The server cannot process the request due to invalid or corrupt data. Review the request, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api ](https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api )

### [SNYK-OPENAPI-0002](#snyk-openapi-0002)

#### Forbidden

Access to the requested resource is forbidden.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)


### [SNYK-OPENAPI-0003](#snyk-openapi-0003)

#### Not acceptable

The server cannot provide a response that matches the provided accept headers. Review the request, then try again.

**HTTP Status:** [406](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406)


### [SNYK-OPENAPI-0004](#snyk-openapi-0004)

#### Not found

The server cannot find the requested resource.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)


### [SNYK-OPENAPI-0005](#snyk-openapi-0005)

#### Method not allowed

The target endpoint does not support your request method. Review the request, then try again.

**HTTP Status:** [405](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405)


### [SNYK-OPENAPI-0006](#snyk-openapi-0006)

#### Request entity too large

The request entity exceeds server limitations.

**HTTP Status:** [413](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413)


### [SNYK-OPENAPI-0007](#snyk-openapi-0007)

#### Unauthorized

The request lacks authentication credentials for the requested resource.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**Help Links:**
- [https://docs.snyk.io/snyk-api-info/authentication-for-api ](https://docs.snyk.io/snyk-api-info/authentication-for-api )

### [SNYK-OPENAPI-0008](#snyk-openapi-0008)

#### Unsupported media type

The media format of the request is not supported.

**HTTP Status:** [415](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415)


---
# [Open Source Languages & Package Managers](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support)
### [SNYK-OS-0001](#snyk-os-0001)

#### Unsupported Ecosystem

The language or package manager is not supported.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support)

### [SNYK-OS-0002](#snyk-os-0002)

#### Unable to parse manifest file

The provided manifest file could not be parsed as it has invalid syntax or does not match the expected schema. Review the manifest file, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-0003](#snyk-os-0003)

#### Lock file out of sync with manifest file

Some of the dependencies that are expected to be in the lock file are missing, usually indicating that the lock file is out of sync with the provided manifest file. Re-sync the lock file, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-0004](#snyk-os-0004)

#### Unable to parse lock file

The provided lock file could not be parsed as it has invalid syntax or does not match the expected schema. Review the lock file, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-0005](#snyk-os-0005)

#### Unknown dependency version

Dependency version could not be resolved.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**
- [https://support.snyk.io/hc/en-us/articles/360001373178-Could-not-determine-version-for-dependencies](https://support.snyk.io/hc/en-us/articles/360001373178-Could-not-determine-version-for-dependencies)

### [SNYK-OS-0006](#snyk-os-0006)

#### Payload missing required elements

The server could not process the request.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0007](#snyk-os-0007)

#### Files cannot be processed

The dependency service could not process the files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0008](#snyk-os-0008)

#### Source is not supported

The source used is not supported by fetcher. The supported sources are: github, bitbucket, gitlab.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0009](#snyk-os-0009)

#### Cannot get file from source

Could not get the file from the source URL.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-0010](#snyk-os-0010)

#### No released version for versions range

There was no version released for the specified versions range.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0011](#snyk-os-0011)

#### Timeout when processing the dependency tree

There was an timeout when processing the dependecy tree.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0012](#snyk-os-0012)

#### Received more manifests than expected

Too many manifest files were provided in the request body.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0013](#snyk-os-0013)

#### Failed to apply dependency updates

An error occured while updating dependencies.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0014](#snyk-os-0014)

#### Unknown blob encoding on Github

Unknown blob encoding on Github.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0015](#snyk-os-0015)

#### No result from forked process

No result from forked process.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-0016](#snyk-os-0016)

#### Child Process Execution Error

The child process encountered an error during execution.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-0017](#snyk-os-0017)

#### No valid package upgrades

The system attempted to find valid upgrades for the packages specified in the lock file, but none were available.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0018](#snyk-os-0018)

#### No dependency updates

There are no available updates for the dependencies.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0019](#snyk-os-0019)

#### Missing required request header

The server encountered a request that is missing a mandatory request header.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0020](#snyk-os-0020)

#### Could not parse JSON file

An error occurred while attempting to parse a JSON file.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0021](#snyk-os-0021)

#### Could not Base64 encode

An error occurred while attempting to perform Base64 encoding.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0022](#snyk-os-0022)

#### Could not Base64 decode

An error occurred while attempting to perform Base64 decoding.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-0023](#snyk-os-0023)

#### Missing supported file

Could not find supported file.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-0024](#snyk-os-0024)

#### Missing environment variable

The server encountered a critical operation that requires a specific environment variable, but the variable is not set or is not accessible within the current environment.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-0025](#snyk-os-0025)

#### No output from isolated builds

The response from isolated builds had no output.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-0026](#snyk-os-0026)

#### Failed to relock

An error occurred while attempting to relock.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


### [SNYK-OS-0027](#snyk-os-0027)

#### Invalid configuration

The configuration parameter does not meet the expected data type. Please ensure the provided value is of the correct data type.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


### [SNYK-OS-DOTNET-0001](#snyk-os-dotnet-0001)

#### Unsupported manifest file type for remediation

The provided manifest file is not supported by Snyk for .NET.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net#git-services-for-.net-projects](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net#git-services-for-.net-projects)
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net#dependencies-managed-by-packagereference](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net#dependencies-managed-by-packagereference)
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net#dependencies-managed-by-packages.config](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net#dependencies-managed-by-packages.config)
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net#paket-dependencies-managed-by-paket](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-.net#paket-dependencies-managed-by-paket)

### [SNYK-OS-GO-0001](#snyk-os-go-0001)

#### Failed to access private module

Snyk could not access the private modules within your go.mod files.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-golang#go-modules-and-snyk-cli](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-golang#go-modules-and-snyk-cli)
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-golang#go-modules-and-git](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-golang#go-modules-and-git)

### [SNYK-OS-GO-0002](#snyk-os-go-0002)

#### Go mod file not found

A go.mod file was not found in the current directory or any parent directory.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-golang#go-modules-and-snyk-cli](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-golang#go-modules-and-snyk-cli)
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-golang#go-modules-and-git](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-golang#go-modules-and-git)

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

Generating the dependency graph requires Snyk to run go list `go list -deps -json` inside the Project. If the operation fails, creating a full dependency graph cannot continue.  

This error usually means that you need some cleanup, such as `go mod tidy`) or your Project deployment process contains a code generation step such as `protobuf` or something similar that is not currently supported by Snyk. 

To verify if this is the case, clone your Project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. 

If Snyk cannot process your code successfully, insert the Snyk CLI as part of your deployment pipeline.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://docs.snyk.io/snyk-cli](https://docs.snyk.io/snyk-cli)
- [https://github.com/snyk/snyk-go-plugin](https://github.com/snyk/snyk-go-plugin)
- [https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go](https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go)

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

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0004](#snyk-os-maven-0004)

#### Cyclic property detected in POM file

There is circular dependency among properties in the Maven project's configuration file (POM), preventing proper resolution and causing an error.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)


### [SNYK-OS-MAVEN-0005](#snyk-os-maven-0005)

#### Error parsing the XML file

There is an error parsing the XML file. This could be referring to either pom.xml or maven-metadata.xml.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)


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


### [SNYK-OS-PIP-0001](#snyk-os-pip-0001)

#### Unsupported manifest file type for remediation

The provided requirements file is not supported by Snyk for Python.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**
- [https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-python#git-services-for-pip-projects](https://docs.snyk.io/scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-python#git-services-for-pip-projects)

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

The request cannot be processed due to an internal error. Review the request, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)


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

Unable to find the coordinates for the provided SHA1.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**
- [https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files)

---
# [PURL Vulnerabilities](https://docs.snyk.io/introducing-snyk/getting-started-snyk-intel-vuln-db-access#about-the-snyk-vulnerability-database)
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

A list of components of the purl spec is required. The purl did not specify all the required components.

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

The given Package URL does not contain a supported vendor.

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


--- Generated at 2023-07-24T12:18:17.726Z
