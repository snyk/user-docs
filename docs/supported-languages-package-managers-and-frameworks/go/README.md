# Go

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## Getting started with Snyk for Go across environments

### Snyk CLI&#x20;

#### Prerequisites

1. [Create a Snyk account](../../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md).
2. [Install Snyk CLI and authenticate your machine](../../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine).
3. [Set the Snyk Organization for CLI tests](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests.md).
4. Ensure you have installed the relevant package manager before you begin using the Snyk CLI (open source).
5. Ensure you have included the relevant manifest files supported by Snyk before testing.

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [`snyk-to-html`](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md)

#### Open source and licensing

{% hint style="info" %}
To scan your dependencies in the CLI, ensure you have installed the relevant package manager and that your Project contains the supported manifest files.
{% endhint %}

#### **Go Modules and Snyk CLI**

{% hint style="info" %}
Snyk scans Go Modules Projects in the CLI at the package level rather than the module level, as Snyk has full access to your local source code.
{% endhint %}

{% hint style="info" %}
Packages from the [Go standard library](https://pkg.go.dev/std) are not supported or included in the dependency tree.&#x20;

Packages under `golang.org/x/` which are [part of the Go Project](https://pkg.go.dev/golang.org/x) but outside the main Go tree are supported.
{% endhint %}

To build the dependency tree, Snyk uses the `go list -json -deps ./...` command, and the dependencies found in `Imports` .

`TestImports` and `XTestImports` are not supported.

When you test Go Modules Projects using the CLI, Snyk does not require their dependencies to be installed, but you must have a `go.mod` file at the root of your Project. `go list` uses this and your Project source code to build a complete dependency tree.

{% hint style="info" %}
Different versions of Go generate different results for the `go list -json -deps` command. This can affect the dependency tree and the vulnerabilities that the Snyk CLI finds.
{% endhint %}

#### **Dep and Snyk CLI**

To build the dependency tree, Snyk analyzes your `Gopkg.lock` files.

When you test dep Projects using the CLI, Snyk requires installation of dependencies. Run `dep ensure` to achieve this.

### Snyk Web UI (Git repository integration)

#### **Go Modules and Git**

{% hint style="warning" %}
By default, dependencies for Go Modules Projects imported using Git are resolved at the **module** level rather than the **package** level, as with Projects tested in the CLI. Thus, when importing using Git, you may see more dependencies and issues reported, including potential false positives, than with the CLI.

To obtain the best possible resolution, enable [full source code analysis](./#enable-full-source-code-analysis).
{% endhint %}

When full source code analysis is enabled, Snyk uses the `go list -json -deps ./...` command to build the dependency tree the same way the CLI test does. Otherwise, it uses `go mod graph` .

#### Enable full source code analysis

To build the most accurate dependency tree for Go Modules Projects imported from Git, Snyk needs to access all the files in your repository.

This enables Snyk to see the `import` statements in your `.go` source files, and determine which specific packages are used in your application. Without this access, Snyk will include all packages from the modules listed in your `go.mod` file.

To enable full source code analysis, adjust your settings as follows:

1. Log in to your account and select your Group and Organization.
2. Navigate to **Settings,** then **Languages**.
3. Select **Edit settings** for **Go**.
4. Toggle full source code analysis on or off.

<figure><img src="../../.gitbook/assets/image (149) (1).png" alt=""><figcaption><p>Enable full source code analysis</p></figcaption></figure>

For more details on levels of access to your repository required by different Snyk features, see [How Snyk handles your data](../../working-with-snyk/how-snyk-handles-your-data.md).

#### **Private modules**

Go modules projects that rely on modules from private Git repositories are supported if those repositories are in the same Git organization as the main project repository.

If you have private modules in repositories from other Git organizations, your Project imports may not work properly. The same is true if your code uses Git Submodules from another organization.

Private module support in different SCMs varies depending on whether [full source code analysis](./#enable-full-source-code-analysis) is enabled or disabled.

| Full source code analysis enabled                                                                                                      | Full source code analysis disabled                                                |
| -------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| <ul><li>Azure Repos</li><li>Bitbucket Cloud</li><li>Bitbucket Server</li><li>GitHub</li><li>GitLab</li><li>GitHub Enterprise</li></ul> | <p></p><ul><li>Bitbucket Cloud</li><li>GitHub</li><li>GitHub Enterprise</li></ul> |

#### **Snyk Broker**

{% hint style="warning" %}
Snyk Broker is currently supported only  when [full source code analysis](./#enable-full-source-code-analysis) is disabled
{% endhint %}

Go Modules Projects imported using new [Snyk Broker](../../enterprise-configuration/snyk-broker/) clients should work as expected.

To add support to clients created before December 30, 2020, add `go.mod` and `go.sum` to your `accept.json` file, as per the changes in this [pull request](https://github.com/snyk/broker/pull/299/files).

If you're using private Go Modules integrated through the Broker, each private module must have a `go.mod` file defined.

#### **Dep and Git**

To build the dependency tree, Snyk analyzes the `Gopkg.lock` files in your Git repository.

#### What's next?

* [Open a Fix PR](./#open-a-fix-pr)&#x20;
* [Configure PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

### Snyk integrations&#x20;

* For integrated development environments, see [Use Snyk in your IDE](../../scm-ide-and-ci-cd-workflow-and-integrations/use-snyk-in-your-ide/).
* If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../../scm-ide-and-ci-cd-workflow-and-integrations/snyk-ci-cd-integrations/) and [Snyk API](../../snyk-api/)).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;

