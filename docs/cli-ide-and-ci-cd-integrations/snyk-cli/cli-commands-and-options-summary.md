# CLI commands and options summary



{% hint style="info" %}
This page **only summarizes** the CLI commands and the options for each command. **For details, use the links in this summary to open the help** docs page for the command you are using. The help docs pages are the same as the help in the CLI.
{% endhint %}

## Usage

`snyk [COMMAND] [SUBCOMMAND] [OPTIONS] [PACKAGE] [CONTEXT-SPECIFIC-OPTIONS]`

## Description

The Snyk CLI is a build-time tool to find and fix known vulnerabilities in your projects. For a more detailed description of Snyk CLI and Snyk, see [Snyk CLI](./). For an introduction on how to use the Snyk CLI, see [Getting started with the CLI](getting-started-with-the-snyk-cli.md).

## Available CLI commands

To learn more about each Snyk CLI command, use the `--help` option, for example, `snyk auth --help` or `snyk container --help`. Each command in this list is linked to the corresponding help page in these docs.

**Note:** Lists of all the options for Snyk CLI commands are on this page. The options are explained in detail in the help for each command.

### [`snyk auth`](commands/auth.md)

Authenticate Snyk CLI with a Snyk account.

### [`snyk config`](commands/config.md)

Manage Snyk CLI configuration.

### [`snyk config environment`](commands/config-environment.md)

Use to set your environment for the region before you run the `snyk auth` command.

### [`snyk test`](commands/test.md)

Test a Project for open-source vulnerabilities and license issues.

### [`snyk monitor`](commands/monitor.md)

Snapshot and continuously monitor a project for open-source vulnerabilities and license issues.

### [`snyk code`](commands/code.md)

Print the name of the `snyk code` command with its help option: `snyk code test`

### [`snyk code test`](commands/code-test.md)

Test source code for any known security issues (Static Application Security Testing).

### [`snyk container`](commands/container.md)

Print a list of the `snyk container` commands, `snyk container monitor` and `snyk container test`.

### [`snyk container monitor`](commands/container-monitor.md)

Capture the container image layers and dependencies and monitor for vulnerabilities on [snyk.io](https://snyk.io).

### [`snyk container SBOM`](commands/container-sbom.md)

Generate an SBOM for a container image

### [`snyk container test`](commands/container-test.md)

Test container images for any known vulnerabilities.

### [`snyk iac`](commands/iac.md)

Print a list of the `snyk iac` commands: `snyk iac describe`, `snyk iac update-exclude-policy`, and `snyk iac test`.

### [`snyk iac test`](commands/iac-test.md)

Test for any known security issue.

### [`snyk iac describe`](commands/iac-describe.md)

Detect, track, and alert on unmanaged resources.

### [`snyk iac update-exclude-policy`](commands/iac-update-exclude-policy.md)

Generate exclude policy rules to be used by `snyk iac describe`.

### [`snyk ignore`](commands/ignore.md)

Modify the `.snyk` policy to ignore stated issues.

### [`snyk log4shell`](commands/log4shell.md)

Find Log4Shell vulnerability.

### [`snyk policy`](commands/policy.md)

Display the `.snyk` policy for a package.

### [`snyk sbom`](commands/sbom.md)

Generate an SBOM for a local software project in an ecosystem supported by Snyk.

### &#x20;[`snyk sbom monitor`](broken-reference)&#x20;

Create a target and projects in your Snyk account to be continuously monitored for open-source vulnerabilities and license issues, sending the results to [snyk.io](https://snyk.io).

### [`snyk sbom test`](commands/sbom-test.md)

Check an SBOM for vulnerabilities in open-source packages.

### [`snyk aibom`](commands/aibom.md)

Generates an AIBOM for a local software project that is written in Python, to help you understand what AI models, datasets, tools, and so on are used in that project.

### [`snyk apps`](../../snyk-api/using-specific-snyk-apis/snyk-apps-apis/create-a-snyk-app-using-the-snyk-cli.md)

Create a Snyk App using the Snyk CLI. For more information, see [Snyk Apps](../../snyk-api/using-specific-snyk-apis/snyk-apps-apis/).

## Subcommands of CLI commands

The following is a list of the sub-commands for Snyk CLI commands. Each sub-command is followed by the command(s) to which the sub-command applies. The commands are linked to their help docs. For details concerning each sub-command, see the help docs.

`get <KEY>`: subcommand of [`config`](commands/config.md)

`set <KEY>=<VALUE>`: subcommand of [`config`](commands/config.md)

`unset <KEY>`: subcommand of [`config`](commands/config.md)

`clear`: subcommand of [`config`](commands/config.md)

`environment`: subcommand of [`config`](commands/config.md)

## Configure the Snyk CLI

You can use environment variables to configure the Snyk CLI and also set variables to configure the Snyk CLI to connect with the Snyk API. See [Configure the Snyk CLI](configure-the-snyk-cli/).

## Debug

See [Debugging the Snyk CLI](debugging-the-snyk-cli.md) for detailed information about the `--d` option.

## Exit codes for CLI commands

Exit codes for the `test` commands are all the same. See the exit codes in the following help docs:

* [`snyk test` exit codes](commands/test.md#exit-codes)
* [`snyk container test` exit codes](commands/container-test.md#exit-codes)
* [`snyk iac test` exit codes](commands/iac-test.md#exit-codes)
* [`snyk code test` exit codes](commands/code-test.md#exit-codes)

Additional CLI commands have exit codes as listed in the following help docs:

* [`snyk monitor` exit codes](commands/monitor.md#exit-codes)
* [`snyk container monitor` exit codes](commands/container-monitor.md#exit-codes)
* [`snyk iac describe` exit codes](commands/iac-describe.md#exit-codes)
* [`snyk iac update-exclude-policy` exit codes](commands/iac-update-exclude-policy.md#exit-codes)
* [`snyk log4shell` exit codes](commands/log4shell.md#exit-codes)
* [`snyk sbom` exit codes](commands/sbom.md#exit-codes)
* [`snyk sbom monitor` exit codes](broken-reference)
* &#x20;[`snyk sbom test` exit codes](commands/sbom-test.md#exit-codes)
* [`snyk container sbom` exit codes](commands/container-sbom.md#exit-codes)

## Options for multiple commands

Lists of the options for Snyk CLI commands follow. Each option is followed by the command(s) to which the option applies. The commands are linked to their help docs. For details concerning each option, see the [help docs](commands/).

`--all-projects`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--fail-fast`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--detection-depth=<DEPTH>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`iac test`](commands/iac-test.md), [`sbom`](commands/sbom.md)

`--exclude=<NAME>[,<NAME>]...>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--prune-repeated-subdependencies`, `-p`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--print-deps`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`container test`](commands/container-test.md)

`--remote-repo-url=<URL>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`iac test`](commands/iac-test.md)

`--dev`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--org=<ORG_ID>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`code test`](commands/code-test.md), [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md), [`iac test`](commands/iac-test.md), [`iac describe`](commands/iac-describe.md), [`sbom`](commands/sbom.md), [`container sbom`](commands/container-sbom.md), [`aibom`](commands/aibom.md)

`--file=<FILE>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--file=<FILE_PATH>`: [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md), [`sbom test`](commands/sbom-test.md)

`--package-manager=<PACKAGE_MANAGER_NAME>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--unmanaged`: [`test`](commands/test.md), [`monitor`](commands/monitor.md). See also [Options for scanning using `--unmanaged`](cli-commands-and-options-summary.md#options-for-scanning-using-unmanaged) and the [`sbom`](commands/sbom.md) command help for another use of this option.

`--ignore-policy`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`iac test`](commands/iac-test.md), [`iac describe`](commands/iac-describe.md)

`--trust-policies`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--show-vulnerable-paths=<none|some|all>`: [`test`](commands/test.md)

`--project-name=<PROJECT_NAME>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md)

`--target-reference=<TARGET_REFERENCE>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`iac test`](commands/iac-test.md), [`container monitor`](commands/container-monitor.md)

`--policy-path=<PATH_TO_POLICY_FILE>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md), [`iac test`](commands/iac-test.md), [`iac describe`](commands/iac-describe.md), [`ignore`](commands/ignore.md)

`--json`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`code test`](commands/code-test.md), [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md), [`iac test`](commands/iac-test.md), [`iac describe`](commands/iac-describe.md), [`sbom test`](commands/sbom-test.md)

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md), [`code test`](commands/code-test.md), [`container test`](commands/container-test.md), [`iac test`](commands/iac-test.md), [`sbom`](commands/sbom.md)

`--sarif`: [`test`](commands/test.md), [`code test`](commands/code-test.md), [`container test`](commands/container-test.md), [`iac test`](commands/iac-test.md)

`--sarif-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md), [`code test`](commands/code-test.md), [`container test`](commands/container-test.md), [`iac test`](commands/iac-test.md)

`--severity-threshold=<low|medium|high|critical>`: [`test`](commands/test.md), [`code test`](commands/code-test.md), [`container test`](commands/container-test.md), [`iac test`](commands/iac-test.md)

`--fail-on=<all|upgradable|patchable>`: [`container test`](commands/container-test.md), [`test`](commands/test.md)

`--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...>`: [`monitor`](commands/monitor.md), [`container monitor`](commands/container-monitor.md), [`iac test`](commands/iac-test.md)

`--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE>]...>`: [`monitor`](commands/monitor.md), [`container monitor`](commands/container-monitor.md), [`iac test`](commands/iac-test.md)

`--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...>`: [`monitor`](commands/monitor.md), [`container monitor`](commands/container-monitor.md), [`iac test`](commands/iac-test.md)

`--project-tags=<TAG>[,<TAG>]...>`: [`monitor`](commands/monitor.md), [`container monitor`](commands/container-monitor.md), [`iac test`](commands/iac-test.md)

`--tags=<TAG>[,<TAG>]...>`: [`monitor`](commands/monitor.md), [`container monitor`](commands/container-monitor.md)&#x20;

## `snyk aibom` command options

`--html`\
`--json-file-output`: [snyk aibom](commands/aibom.md)

## `snyk auth` command options

`--auth-type=<TYPE>`\
`--client-secret=<SECRET>`\
`--client-id=<ID>`: [`snyk auth`](commands/auth.md)

## `snyk code test` command option&#x20;

`--include-ignores`: [`code test`](commands/code-test.md)

## `snyk config environment` command option

`--no-check` [snyk config environment](commands/config-environment.md)

## `snyk container` command options

`--app-vulns`: [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md)

`--exclude-app-vulns`: [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md), [`container sbom`](commands/container-sbom.md)

`--nested-jars-depth`: [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md)

`--exclude-base-image-vulns`: [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md)

`--platform=<PLATFORM>`: [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md) ; [`container sbom`](commands/container-sbom.md)&#x20;

`--username=<CONTAINER_REGISTRY_USERNAME>`: [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md)

`--password=<CONTAINER_REGISTRY_PASSWORD>`: [`container test`](commands/container-test.md), [`container monitor`](commands/container-monitor.md)

## `snyk iac test` command options

`--scan=<TERRAFORM_PLAN_SCAN_MODE>`: [`iac test`](commands/iac-test.md)

`--target-name=<TARGET_NAME>`: [`iac test`](commands/iac-test.md)

`--rules=<PATH_TO_CUSTOM_RULES_BUNDLE>`: [`iac test`](commands/iac-test.md)

`--var-file=<PATH_TO_VARIABLE_FILE>`: [`iac test`](commands/iac-test.md)

`--report`: [`iac test`](commands/iac-test.md)

## `snyk iac describe` command options

`--from=<STATE>[,<STATE>...]`: [`iac describe`](commands/iac-describe.md)

`--to=<PROVIDER+TYPE>`: [`iac describe`](commands/iac-describe.md)

`--service=<SERVICE>[,<SERVICE]...>`: [`iac describe`](commands/iac-describe.md)

`--quiet`: [`iac describe`](commands/iac-describe.md)

`--filter`: [`iac describe`](commands/iac-describe.md)

`--html`: [`iac describe`](commands/iac-describe.md)&#x20;

`--html-file-output=<OUTPUT_FILE_PATH>`: [`iac-describe`](commands/iac-describe.md)

`--fetch-tfstate-headers`: [`iac describe`](commands/iac-describe.md)

`--tfc-token`: [`iac describe`](commands/iac-describe.md)

`--tfc-endpoint`: [`iac describe`](commands/iac-describe.md)

`--tf-provider-version`: [`iac describe`](commands/iac-describe.md)

`--strict`: [`iac describe`](commands/iac-describe.md)

`--deep`: [`iac describe`](commands/iac-describe.md)

`--tf-lockfile`: [`iac describe`](commands/iac-describe.md)

-`-config-dir`: [`iac describe`](commands/iac-describe.md)

## `snyk iac update-exclude-policy` command options

`--exclude-changed`: [`iac update-exclude-policy`](commands/iac-update-exclude-policy.md)

`--exclude-missing`: [`iac update-exclude-policy`](commands/iac-update-exclude-policy.md)

`--exclude-unmanaged`: [`iac update-exclude-policy`](commands/iac-update-exclude-policy.md)

## `snyk ignore` command options

`--id=<ISSUE_ID>`: [`ignore`](commands/ignore.md)

`--expiry=<EXPIRY>`: [`ignore`](commands/ignore.md)

`--reason=<REASON>`: [`ignore`](commands/ignore.md)

`--path=<PATH_TO_RESOURCE>`: [`ignore`](commands/ignore.md)

## `snyk sbom` and `snyk container sbom` command options

`--format=<cyclonedx1.4+json|cyclonedx1.4+xml|cyclonedx1.5+json|cyclonedx1.5+xml|cyclonedx1.6+json|cyclonedx1.6+xml|spdx2.3+json>`: [`snyk sbom`](commands/sbom.md), [`snyk container sbom`](commands/container-sbom.md)

`[--file=] or [--f=]`: [`snyk sbom`](commands/sbom.md)

`[--name=<NAME>]`: [`snyk sbom`](commands/sbom.md)

`[--version=<VERSION>]`: [`snyk sbom`](commands/sbom.md)

`[<TARGET_DIRECTORY>]`: [`snyk sbom`](commands/sbom.md)

`<IMAGE>`: [`snyk container sbom`](commands/container-sbom.md)

## Option for Maven projects

`--maven-aggregate-project`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--scan-unmanaged`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--scan-all-unmanaged`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

## Options for Gradle projects

`--sub-project=<NAME>`, `--gradle-sub-project=<NAME>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--all-sub-projects`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--configuration-matching=<CONFIGURATION_REGEX>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--configuration-attributes=<ATTRIBUTE>[,<ATTRIBUTE>]...`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--init-script=<FILE`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

## Options for .Net and NuGet projects

`--file=.sln`: [`test`](commands/test.md)

`--file=<filename>.sln`: [`sbom`](commands/sbom.md)

-`-file=packages.config`: [`test`](commands/test.md), [`sbom`](commands/sbom.md)

`--assets-project-name`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--packages-folder`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--project-name-prefix=<PREFIX_STRING>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--project-name-prefix=my-group/`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--dotnet-runtime-resolution`:  [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--dotnet-target-framework`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

## Options for npm projects

`--strict-out-of-sync=true|false`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

## Options for pnpm projects

`--dev:` [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--all-projects`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--fail-on`: [`test`](commands/test.md)

`--prune-repeated-subdependencies`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

## Options for Yarn projects

`--strict-out-of-sync=true|false`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--yarn-workspaces`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

## Options for CocoaPods projects

`--strict-out-of-sync=true|false`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

## Options for Python projects

`--command=<COMMAND>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--skip-unresolved=true|false`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--File=<filename>`: [`sbom`](commands/sbom.md)

`--pakage-manager=<package manager>`: [`sbom`](commands/sbom.md)

## Options for Go projects

The following options are not supported:

`--fail-on=<all|upgradable|patchable>`: [`test`](commands/test.md)

## Options for scanning using `--unmanaged`

`--org=<ORG_ID>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--json`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md)

`--remote-repo-url=<URL>`: [`test`](commands/test.md)

`--severity-threshold=<low|medium|high|critical>`: [`test`](commands/test.md)

`--target-reference=<TARGET_REFERENCE>`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--max-depth`: [`test`](commands/test.md), [`monitor`](commands/monitor.md), [`sbom`](commands/sbom.md)

`--print-dep-paths`: [`test`](commands/test.md), [`monitor`](commands/monitor.md)

`--project-name=c-project`: [`monitor`](commands/monitor.md)

## `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

These options are used with the `snyk test` and `snyk monitor` commands. See the help docs for [`snyk test`](commands/test.md) and [`snyk monitor`](commands/monitor.md) for details.
