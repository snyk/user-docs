# CLI reference

## Usage

`snyk [COMMAND] [SUBCOMMAND] [OPTIONS] [PACKAGE] [CONTEXT-SPECIFIC-OPTIONS]`

## Description

The Snyk CLI is a build-time tool to find and fix known vulnerabilities in your projects. For a more detailed description of the Snyk CLI and Snyk, see [Snyk CLI](../). For an introduction on how to use the Snyk CLI, see [Getting started with the CLI](../getting-started-with-the-cli/).

## Available CLI commands

To learn more about each Snyk CLI command, use the `--help` option, for example, `snyk auth --help` or `snyk container --help`. You can also use the help command: `snyk help [<COMMAND>]`.

Each command in this list is linked to the corresponding help page in these docs. A list of all the [options](./#options) for Snyk CLI commands is at the end of this page. The options are explained in detail in the help for each command.

### ``[`snyk auth`](../commands/auth.md)``

Authenticate Snyk CLI with a Snyk account.

### ``[`snyk test`](../commands/test.md)``

Test a project for open source vulnerabilities and license issues.

### ``[`snyk monitor`](../commands/monitor.md)``

Snapshot and continuously monitor a project for open source vulnerabilities and license issues.

### ``[`snyk container`](../commands/container.md)``

Test container images for vulnerabilities.

### ``[`snyk iac`](../commands/iac.md)``

Find security issues in Infrastructure as Code files.

### ``[`snyk code`](../commands/code.md)``

Find security issues using static code analysis.

### ``[`snyk log4shell`](../commands/log4shell.md)``

Find Log4Shell vulnerability.

### ``[`snyk config`](../commands/config.md)``

Manage Snyk CLI configuration.

### ``[`snyk policy`](../commands/policy.md)``

Display the `.snyk` policy for a package.

### ``[`snyk ignore`](../commands/ignore.md)``

Modify the `.snyk` policy to ignore stated issues.

## New CLI commands

### ``[`snyk fix`](../fix-vulnerabilities-from-the-cli/automatic-remediation-with-snyk-fix.md)``

Apply the recommended updates for supported ecosystems automatically.

### [`snyk apps`](../create-a-snyk-app-using-the-snyk-cli.md)

Create a Snyk App using the Snyk CLI.

## Subcommands of CLI commands

The following is a list of the sub-commands for Snyk CLI commands. Each sub-command is followed by the command(s) to which the sub-command applies. The commands are linked to their help docs. For details concerning each sub-command, see the help docs.

`test`: subcommand of [`code`](../commands/code.md), [`container`](../commands/container.md), and [`iac`](../commands/iac.md)``

`monitor`: subcommand of [`container`](../commands/container.md)``

`get <KEY>`: subcommand of [`config`](../commands/config.md)``

`set <KEY>=<VALUE>`: subcommand of [`config`](../commands/config.md)``

`unset <KEY>`: sub-command of [`config`](../commands/config.md)``

`clear`: subcommand of [`config`](../commands/config.md)``

## Exit codes

Possible exit codes and their meaning:

**0**: success, no vulnerabilities found\
**1**: action\_needed, vulnerabilities found\
**2**: failure, try to re-run command\
**3**: failure, no supported projects

## Configure the Snyk CLI

You can use environment variables to configure the Snyk CLI and also set variables to configure the Snyk CLI to connect with the Snyk API. See [Configure the Snyk CLI](./#configure-the-snyk-cli).

## Debug

Use `-d` option to output the debug logs.

## Options

The following is a list of the options for Snyk CLI commands. Each option is followed by the command(s) to which the option applies. The commands are linked to their help docs. For details concerning each option, see the help docs.

`--all-projects`: `test`, `monitor`

`--detection-depth=<DEPTH>`: `test`, `monitor`, `iac`

`--exclude=<DIRECTORY>[,<DIRECTORY>]...>`: `test`, `monitor`

`--prune-repeated-subdependencies, -p`: `test`, `monitor`

`--print-deps`: `test`, `monitor`, container

`--remote-repo-url=<URL>`: `test`, `monitor`

`--dev`: `test`, `monitor`

`--org=<ORG_ID>`: `test`, `monitor`, `code`, `iac`

`--file=<FILE>`: `test`, `monitor`, container

`--package-manager=<PACKAGE_MANAGER_NAME>`: `test`, `monitor`

`--ignore-policy`: `test`, `monitor`, `iac`

`--trust-policies` `test`, `monitor`

`--show-vulnerable-paths=<none|some|all>` `test`, `monitor`

`--project-name=<PROJECT_NAME>`: `test`, `monitor`, `container`

`--target-reference=<TARGET_REFERENCE>`: `test`, `monitor`

`--policy-path=<PATH_TO_POLICY_FILE>`: `test`, `monitor`, `container`, `iac`, `ignore`

`--json`: `test`, `monitor`, `code`, `container`, `iac`

`--json-file-output=<OUTPUT_FILE_PATH>`: `test`, `monitor`, `container`, `iac`

`--sarif`: `test`, `monitor`, `code`, `container`, `iac`

`--sarif-file-output=<OUTPUT_FILE_PATH>`: `test`, `monitor`, `container`, `iac`

`--severity-threshold=<low|medium|high|critical>`: `test`, `code`, `container`, `iac`

`--fail-on=<all|upgradable|patchable>`: `test`

`--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...>`: `monitor`, `container`

`--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE>]...>`: `monitor`, `container`

`--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...>`: `monitor`, `container`

`--project-tags=<TAG>[,<TAG>]...>`: `monitor`, `container`

`--tags=<TAG>[,<TAG>]...>`: `monitor`, `container`

`--exclude-base-image-vulns`: `container`

`--platform=<PLATFORM>`: `container`

`--username=<CONTAINER_REGISTRY_USERNAME>`: `container`

`--password=<CONTAINER_REGISTRY_PASSWORD>`: `container`

`--scan=<TERRAFORM_PLAN_SCAN_MODE>`: `iac`

`--rules=<PATH_TO_CUSTOM_RULES_BUNDLE>`: `iac`

`--id=<ISSUE_ID>`: `ignore`

`--expiry=<EXPIRY>`: `ignore`

`--reason=<REASON>`: `ignore`

`--path=<PATH_TO_RESOURCE>`: `ignore`

## Debug

`-d`: [all](../commands/)

## Options for Maven projects

`--scan-all-unmanaged`: `test`, `monitor`

`--reachable`: `test`, `monitor`

`--reachable-timeout=<TIMEOUT>`: `test`, `monitor`

## Options for Gradle projects

`--sub-project=<NAME>`, `--gradle-sub-project=<NAME>`: `test`, `monitor`

`--all-sub-projects`: `test`, `monitor`

`--configuration-matching=<CONFIGURATION_REGEX>`: `test`, `monitor`

`--configuration-attributes=<ATTRIBUTE>[,<ATTRIBUTE>]...`: `test`, `monitor`

`--reachable`: `test`, `monitor`

`--reachable-timeout=<TIMEOUT>`: `test`, `monitor`

`--init-script=<FILE`: `test`, `monitor`

## Options for .Net and NuGet projects

`--assets-project-name`: `test`, `monitor`

`--packages-folder`: `test`, `monitor`

`--project-name-prefix=<PREFIX_STRING>`: `test`, `monitor`

`--project-name-prefix=my-group/`: `test`, `monitor`

## Options for npm projects

`--strict-out-of-sync=true|false`: `test`, `monitor`

## Options for Yarn projects

`--strict-out-of-sync=true|false`: `test`, `monitor`

`--yarn-workspaces`: `test`, `monitor`

## Options for CocoaPods projects

`--strict-out-of-sync=true|false`: `test`, `monitor`

## Options for Python projects

`--command=<COMMAND>`: `test`, `monitor`

`--skip-unresolved=true|false`: `test`, `monitor`

## Options for Go projects

Currently the following options are not supported:

`--fail-on=<all|upgradable|patchable>`: `test`

## Options for unmanaged projects

For information about these options see [Snyk CLI for C / C++ projects](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-c-c++#snyk-cli-for-c-c++-projects).

## `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

These options are used with the `test` and `monitor` commands. See the help docs for details.
