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

For Go with Snyk Code, Snyk support:

* Go Standard Library comprehensive as a library&#x20;
* .`go` as a file format

Available features:

* Reports
* Custom rules
* Interfile analysis

## Go for Snyk Open Source

### Available features

Available features for Go Projects with dependencies managed by Go Modules and dep:

* PR checks
* License scanning
* Reports
* Test your app's SBOM and packages using `pkg:golang` PURLs through the [SBOM test](../../developer-tools/snyk-cli/commands/sbom-test.md) command.

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

### Go Modules and dep support

{% hint style="info" %}
**Feature availability**\
Some features may not be available, depending on your plan. For more information, see [Plans and pricing.](https://snyk.io/plans/)
{% endhint %}

#### **Go Modules and the CLI**

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

#### **Dep and the CLI**

To build the dependency tree, Snyk analyzes your `Gopkg.lock` files.

When you test dep Projects using the CLI, Snyk requires installation of dependencies. Run `dep ensure` to achieve this.

#### **Dep and SCM integrations**

To build the dependency tree, Snyk analyzes the `Gopkg.lock` files in your SCM repository.

#### **Go Modules and SCM integrations**

Snyk resolves dependencies for Go Modules Projects imported using an SCM integration at the module level. In contrast, the CLI resolves dependencies at the package level.

Because of this difference, SCM integrations report more dependencies and issues than the CLI, including false positives.

To obtain the best possible resolution, enable [full source code analysis](go.md#enable-full-source-code-analysis).

When full source code analysis is enabled, Snyk uses the `go list -json -deps ./...` command to build the dependency tree the same way the CLI test does. Otherwise, it uses `go mod graph` .

#### Enable full source code analysis

To build the most accurate dependency tree for Go Modules Projects imported from SCM integrations, Snyk needs to access all the files in your repository.

This enables Snyk to see the `import` statements in your `.go` source files, and determine which specific packages are used in your application. Without this access, Snyk includes all packages from the modules listed in your `go.mod` file.

To enable full source code analysis, adjust your settings as follows:

1. Log in to your account and select your Organization.
2. Navigate to **Settings** > **Snyk Open Source**.
3. Select **Edit settings** for **Go**.
4. Toggle **Enable full source code analysis** on or off.

<figure><img src="../../.gitbook/assets/image (119).png" alt=""><figcaption><p>Enable full source code analysis</p></figcaption></figure>

For more details on levels of access to your repository required by different Snyk features, see [How Snyk handles your data](../../snyk-data-and-governance/how-snyk-handles-your-data.md).

#### **Private modules**

Go modules Projects that rely on modules from private SCM repositories are supported if those repositories are in the same SCM organization as the main project repository.

If you have private modules in repositories from other SCM organizations, it is possible that your Project imports do not work properly. The same is true if your code uses SCM submodules from another organization.

If your private modules have other private modules from another SCM organization, your Project imports do not work. All private modules, including the ones within other modules, need to be part of the same SCM organization as the main project repository.

Private module support in different SCMs varies depending on whether full source code analysis is enabled or disabled.

| Full source code analysis enabled                                                                                                      | Full source code analysis disabled                                         |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| <ul><li>Azure Repos</li><li>Bitbucket Cloud</li><li>Bitbucket Server</li><li>GitHub</li><li>GitLab</li><li>GitHub Enterprise</li></ul> | <ul><li>Bitbucket Cloud</li><li>GitHub</li><li>GitHub Enterprise</li></ul> |

#### **Snyk Broker support for Go**

{% hint style="warning" %}
To use Snyk Broker with Go, you must disable [full source code analysis](go.md#enable-full-source-code-analysis).
{% endhint %}

Go Modules Projects imported using new [Snyk Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/) clients should work as expected.

To add support to clients created before December 30, 2020, add `go.mod` and `go.sum` to your `accept.json` file, as per the changes in this [pull request](https://github.com/snyk/broker/pull/299/files).

If you are using private Go Modules integrated through the Broker, each private module must have a `go.mod` file defined.
