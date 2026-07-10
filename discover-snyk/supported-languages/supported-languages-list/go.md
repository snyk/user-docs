---
description: Snyk support for Go with Snyk Open Source and Snyk Code, including SCM import, CLI and IDE testing, and supported frameworks
---

# Go

## Applicability and integration

{% hint style="info" %}
Snyk for Go is supported for Snyk Open Source and Snyk Code.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

## Technical specifications

### Supported frameworks and libraries

* Azure/azure-sdk-for-go/sdk/ai/azopenai
* gage-technologies/mistral-go
* Gin
* google/generative-ai-go/genai
* GORM library
* grpc-go
* labstack/echo
* lib/pq
* sashabaranov/go-openai
* spf13/pflag
* sqlx

### Supported package managers

For Go, Snyk supports [Go Modules](https://go.dev/ref/mod) and [dep](https://github.com/golang/dep) as package managers.

## Go for Snyk Code

For Go with Snyk Code, Snyk supports:

* Go Standard Library comprehensive as a library
* .`go` as a file format

Available features:

* Reports
* Interfile analysis

## Go for Snyk Open Source

### Available features

Available features for Go Projects with dependencies managed by Go Modules and dep:

* PR checks
* License scanning
* Reports
* Test your app's SBOM and packages using `pkg:golang` PURLs through the [SBOM test](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/commands/sbom-test) command.

{% hint style="info" %}
If the **Snyk Fix PR** feature is enabled, this means that you will be notified if the PR checks fail when the following conditions are met:

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}

Snyk supports all versions of Go, including the latest stable version listed on the Go [All releases](https://go.dev/dl/) page.

Snyk tracks only official releases. Snyk does not identify commits, including those in the default branch, unless they are included in an official release or tag. For Projects with a package manager, Snyk requires a release to the package manager. For Go and unamanaged scans (C/C++), Snyk requires an official release or tag in the GitHub repository.

{% hint style="warning" %}
Since January 1, 2023, Snyk has not supported govendor Projects. As a general security best practice, Snyk recommends using tools that are consistently maintained and up-to-date.

Since Snyk no longer supports scanning of govendor Projects, a warning is issued and no results are provided.
{% endhint %}

### Support for Go Modules

{% hint style="info" %}
**Feature availability**\
Some features may not be available, depending on your plan. For more information, visit [Plans and pricing.](https://snyk.io/plans/)
{% endhint %}

#### CLI support for Go Modules

Snyk scans Go Modules Projects in the CLI at the package level rather than the module level, as Snyk has full access to your local source code.

Packages from the [Go standard library](https://pkg.go.dev/std) are supported and included in the dependency tree.

Packages under `golang.org/x/` that are [part of the Go Project](https://pkg.go.dev/golang.org/x) but outside the main Go tree are also supported.

To build the dependency tree for all third party packages, Snyk uses

* The `go list -json -deps ./...` command and the dependencies found in `Imports` .
* The `toolchain` directive in the `go.mod` file
* The `go version` command to determine the Golang version to apply to standard libraries.

{% hint style="info" %}
`TestImports` and `XTestImports` are not supported.
{% endhint %}

When you test Go Modules Projects using the CLI, Snyk does not require that their dependencies are installed, but you must have a `go.mod` file at the root of your Project. `go list` uses this and your Project source code to build a complete dependency tree.

Different versions of Go generate different results for the `go list -json -deps` command. This can affect the dependency tree and the vulnerabilities that the Snyk CLI finds.

#### SCM integrations for Go Modules

The source code management (SCM) integration resolves dependencies using one of two scopes, depending on your configuration:

* Standard SCM scan (default): resolves dependencies at the module level. Snyk parses the `go.mod` file using the `go mod graph` command, which maps the entire dependency tree regardless of whether the application code imports specific packages.
* Full source code analysis scan: resolves dependencies at the package level. When you enable this feature, Snyk clones the repository and runs the `go list -json -deps ./...` command. This forces the SCM integration to analyze active package imports and generate a dependency tree that aligns with the package-level baseline.

Because the default SCM integration evaluates the entire module graph using the `go mod graph` command, it reports a higher number of dependencies and vulnerabilities than the package-level baseline. This introduces vulnerabilities from unused or unimported packages within a module, resulting in findings that do not affect the compiled binary.

Enabling full source code analysis aligns the SCM integration with CLI resolution and eliminates findings from unimported code.

#### Enable full source code analysis

To build the most accurate dependency tree for Go modules Projects imported from SCM integrations, Snyk must access all files in your repository.

This allows Snyk to see the import statements in your `.go` source files and determine which specific packages your application uses. Without this access, Snyk includes all packages from the modules listed in your `go.mod` file.

To enable full source code analysis, adjust your settings as follows:

1. Log in to your account and select your Organization.
2. Navigate to **Settings** > **Snyk Open Source**.
3. Select **Edit settings** for **Go**.
4. Toggle **Enable full source code analysis** on or off.

<figure><img src="../../.gitbook/assets/image (119).png" alt=""><figcaption><p>Enable full source code analysis</p></figcaption></figure>

For more details on levels of access to your repository required by different Snyk features, see [How Snyk handles your data](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/how-snyk-handles-your-data).

**Private modules**

Snyk supports Go modules Projects that rely on modules from private SCM repositories if those repositories are in the same SCM organization as the main Project repository.

Private module support for different SCMs varies based on whether you enable or disable full source code analysis.

| Full source code analysis enabled                                                                                                      | Full source code analysis disabled                                         |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| <ul><li>Azure Repos</li><li>Bitbucket Cloud</li><li>Bitbucket Server</li><li>GitHub</li><li>GitLab</li><li>GitHub Enterprise</li></ul> | <ul><li>Bitbucket Cloud</li><li>GitHub</li><li>GitHub Enterprise</li></ul> |

**Private and Brokered package sources**

To give Snyk access to privately hosted Go modules, configure your private module sources in **Settings > Open Source > Go**. Configure them under one of the following:

* Brokered package sources: for Universal Broker connections that have been enabled for Open Source.
* Private package registries: for direct access.

#### Source requirements and routing

Snyk handles configuration requirements differently based on the source type and how Snyk accesses it.

For package registries or proxies (Artifactory, Nexus), Snyk always requires an explicit registry URL for both brokered and non-brokered setups. Snyk appends the URL to the `GOPROXY` chain so Go can resolve and fetch packages from your private registry. For example: `https://artifactory.example.com/artifactory/api/go/team-go`

For source control private dependencies, (GitHub and GitHub Enterprise, Bitbucket, GitLab, Azure Repos):

* Snyk only requires an explicit URL if you use a brokered connection. For non-brokered setups, Snyk relies natively on your existing SCM Organization permissions, and you do not need to configure a URL here.
* When you provide a URL for a brokered SCM, Snyk adds the host to `GONOSUMDB` so Go correctly tunnels traffic and skips the public checksum database. For example: `https://github.snyk-customer.com/owner/internal-shared-lib`&#x20;

#### Configuration rules

The **Registry type** or **Source type** dropdowns only display options that have an active integration configured in your Organization and valid Universal Broker connections to SCM and package registries. The SCM integration must have permission to access the repositories containing the private Go modules.

Do not include credentials or authentication tokens within the URLs. You must configure credentials directly on the underlying SCM integration or the Universal Broker Client.

You can also view and configure these settings programmatically using the Snyk API, for both [brokered](https://apidocs.snyk.io/?version=2026-03-25#get-/orgs/-org_id-/settings/opensource/-ecosystem-/broker) and [direct](https://apidocs.snyk.io/?version=2026-03-25#get-/orgs/-org_id-/settings/opensource/-ecosystem-/private-registries) connections.

### Support for dep

#### CLI support for dep

To build the dependency tree, Snyk analyzes your `Gopkg.lock` files.

When you test dep Projects using the CLI, Snyk requires installation of dependencies. Run `dep ensure` to achieve this.

#### SCM integrations for dep&#x20;

To build the dependency tree, Snyk analyzes the `Gopkg.lock` files in your SCM repository.
