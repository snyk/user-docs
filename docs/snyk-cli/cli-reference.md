# CLI commands and options summary

**Note:** This page **only summarizes** the CLI commands and the options for each command. Be sure to **use the links in this summary to look at the help** for the command you are using **for details**. (The help in the docs is the same as the help in the CLI.)

## Usage

`snyk [COMMAND] [SUBCOMMAND] [OPTIONS] [PACKAGE] [CONTEXT-SPECIFIC-OPTIONS]`

## Description

The Snyk CLI is a build-time tool to find and fix known vulnerabilities in your projects. For a more detailed description of Snyk CLI and Snyk, see [Snyk CLI](https://docs.snyk.io/snyk-cli). For an introduction on how to use the Snyk CLI, see [Getting started with the CLI](https://docs.snyk.io/snyk-cli/getting-started-with-the-cli).

## Available CLI commands

To learn more about each Snyk CLI command, use the `--help` option, for example, `snyk auth --help` or `snyk container --help`. Each command in this list is linked to the corresponding help page in these docs.

**Note:** Lists of all the options for Snyk CLI commands are on this page. The options are explained in detail in the help for each command.

### [`snyk auth`](https://docs.snyk.io/snyk-cli/commands/auth)

Authenticate Snyk CLI with a Snyk account.

### [`snyk test`](https://docs.snyk.io/snyk-cli/commands/test)

Test a project for open source vulnerabilities and license issues.

### [`snyk monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

Snapshot and continuously monitor a project for open source vulnerabilities and license issues.

### [`snyk container`](https://docs.snyk.io/snyk-cli/commands/container)

Print a list of the `snyk container` commands, `snyk container monitor` and `snyk container test`.

### ``[`snyk container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

Capture the container image layers and dependencies and monitor for vulnerabilities on [snyk.io](https://snyk.io)

### ``[`snyk container test`](https://docs.snyk.io/snyk-cli/commands/container-test)``

Test container images for any known vulnerabilities.

### ``[`snyk iac`](https://docs.snyk.io/snyk-cli/commands/iac)``

Print a list of the `snyk iac` commands: `snyk iac describe`, `snyk iac update-exclude-policy`, and `snyk iac test`.

### [`snyk iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

Detect, track, and alert on infrastructure drift and unmanaged resources.

### [`snyk iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

Generate exclude policy rules to be used by `snyk iac describe`.

### [`snyk iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

Test for any known security issue.

### [`snyk code`](https://docs.snyk.io/snyk-cli/commands/code)

`Print the name of the snyk code` command `with its help option: snyk code test`

### ``[`snyk code test`](https://docs.snyk.io/snyk-cli/commands/code-test)``

Test for any known security issues using Static Code Analysis.

### [`snyk log4shell`](https://docs.snyk.io/snyk-cli/commands/log4shell)

Find Log4Shell vulnerability.

### [`snyk config`](https://docs.snyk.io/snyk-cli/commands/config)

Manage Snyk CLI configuration.

### [`snyk policy`](https://docs.snyk.io/snyk-cli/commands/policy)

Display the `.snyk` policy for a package.

### [`snyk ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

Modify the `.snyk` policy to ignore stated issues.

## New CLI commands

### [`snyk fix`](https://docs.snyk.io/snyk-cli/fix-vulnerabilities-from-the-cli/automatic-remediation-with-snyk-fix)

Apply the recommended updates for supported ecosystems automatically.

### [`snyk apps`](https://docs.snyk.io/snyk-cli/create-a-snyk-app-using-the-snyk-cli)

Create a Snyk App using the Snyk CLI.

## Subcommands of CLI commands

The following is a list of the sub-commands for Snyk CLI commands. Each sub-command is followed by the command(s) to which the sub-command applies. The commands are linked to their help docs. For details concerning each sub-command, see the help docs.

`get <KEY>`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

`set <KEY>=<VALUE>`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

`unset <KEY>`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

`clear`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

## Configure the Snyk CLI

You can use environment variables to configure the Snyk CLI and also set variables to configure the Snyk CLI to connect with the Snyk API. See [Configure the Snyk CLI](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli).

## Debug

Use `-d` option to output the debug logs for any command.

## Options for multiple commands

Lists of the options for Snyk CLI commands follow. Each option is followed by the command(s) to which the option applies. The commands are linked to their help docs. For details concerning each option, see the [help docs](https://docs.snyk.io/snyk-cli/commands).

`--all-projects`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--fail-fast`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--detection-depth=<DEPTH>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac/iac-test)

`--exclude=<NAME>[,<NAME>]...>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--prune-repeated-subdependencies, -p`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--print-deps`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test)``

`--remote-repo-url=<URL>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [iac test](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--dev`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--org=<ORG_ID>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [container monitor](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac/iac-describe)

`--file=<FILE>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--package-manager=<PACKAGE_MANAGER_NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--unmanaged:` [test](commands/test.md)`,` [monitor](commands/monitor.md). See also [Options for scanning using `--unmanaged`](cli-reference.md#option-for-c-c++-projects).

`--ignore-policy`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--trust-policies` [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--show-vulnerable-paths=<none|some|all>` [`test`](https://docs.snyk.io/snyk-cli/commands/test)

`--project-name=<PROJECT_NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

`--target-reference=<TARGET_REFERENCE>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)``

`--policy-path=<PATH_TO_POLICY_FILE>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe), [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--json`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--sarif`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--sarif-file-output=<OUTPUT_FILE_PATH>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--severity-threshold=<low|medium|high|critical>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--fail-on=<all|upgradable|patchable>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [test](https://docs.snyk.io/snyk-cli/commands/test)

`--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)``

`--project-tags=<TAG>[,<TAG>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)``

`--tags=<TAG>[,<TAG>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

## `snyk container` command options

`--file=<FILE_PATH>`: [container test](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

`--app-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container`monitor](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--exclude-app-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

`--nested-jars-depth`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

`--exclude-base-image-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

`--platform=<PLATFORM>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

`--username=<CONTAINER_REGISTRY_USERNAME>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

`--password=<CONTAINER_REGISTRY_PASSWORD>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)``

## `snyk iac test` command options

`--report`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)``

`--scan=<TERRAFORM_PLAN_SCAN_MODE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--target-name=<TARGET_NAME>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)``

`--rules=<PATH_TO_CUSTOM_RULES_BUNDLE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--var-file=<PATH_TO_VARIABLE_FILE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)``

## `snyk iac describe` command options

`--from=<STATE>[,<STATE>...]`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--to=<PROVIDER+TYPE>`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac/iac-describe)

`--service=<SERVICE>[,<SERVICE]...>`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--all`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--only-managed` or `--drift`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--only-unmanaged`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--quiet`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--filter`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--html`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac/iac-describe)

`--html-file-output=<OUTPUT`_`FILE`_`PATH>`: [`iac-describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--fetch-tfstate-headers`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--tfc-token`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--tfc-endpoint`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--tf-provider-version`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--strict`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--deep`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--tf-lockfile`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

\-`-config-dir`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

## `snyk iac update-exclude-policy command options`

`--exclude-changed`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)``

`--exclude-missing`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)``

`--exclude-unmanaged`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)``

## `snyk ignore` command options

`--id=<ISSUE_ID>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--expiry=<EXPIRY>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--reason=<REASON>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--path=<PATH_TO_RESOURCE>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

## Option for Maven projects

`--scan-all-unmanaged`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Gradle projects

`--sub-project=<NAME>`, `--gradle-sub-project=<NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--all-sub-projects`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--configuration-matching=<CONFIGURATION_REGEX>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--configuration-attributes=<ATTRIBUTE>[,<ATTRIBUTE>]...`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--init-script=<FILE`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for .Net and NuGet projects

`--assets-project-name`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--packages-folder`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--project-name-prefix=<PREFIX_STRING>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--project-name-prefix=my-group/`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for npm projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Yarn projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--yarn-workspaces`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for CocoaPods projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Python projects

`--command=<COMMAND>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--skip-unresolved=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Go projects

Currently the following options are not supported:

`--fail-on=<all|upgradable|patchable>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test)

## Options for scanning using `--unmanaged`

`--org=<ORG_ID>`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--json`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md)

`--remote-repo-url=<URL>`: [`test`](commands/test.md)

`--severity-threshold=<low|medium|high|critical>:` [`test`](commands/test.md)

`--target-reference=<TARGET_REFERENCE>`: [`test`](commands/test.md), [monitor](commands/monitor.md)

`--max-depth`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)``

`--print-dep-paths:` [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)``

`--project-name=c-project`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)``

## `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

These options are used with the `snyk test` and `snyk monitor` commands. See the help docs for [`snyk test`](https://docs.snyk.io/snyk-cli/commands/test) and [`snyk monitor`](https://docs.snyk.io/snyk-cli/commands/monitor) for details.
