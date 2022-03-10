# CLI reference

## Usage

`snyk [COMMAND] [SUBCOMMAND] [OPTIONS] [PACKAGE] [CONTEXT`-SPECIFIC\_OPTIONS`]`

## Description

The Snyk CLI is a build-time tool to find and fix known vulnerabilities in your projects. For a more detailed description of the Snyk CLI and Snyk, see [Snyk CLI](../). For an introduction on how to use the Snyk CLI, see [Getting started with the CLI](../getting-started-with-the-cli/).

## Available CLI commands

To learn more about each Snyk CLI command, use the `--help` option, for example, `snyk auth --help` or `snyk container --help`. You can also use the help command: `snyk help [<COMMAND>]`.

Each command in this list is linked to the corresponding help page in these docs. A list of all the [options](./#options) for Snyk CLI commands is at the end of this page. The options are explained in detail in the help for each command.

### [`snyk auth`](../../features/snyk-cli/commands/auth.md)

Authenticate Snyk CLI with a Snyk account.

### [`snyk test`](../../features/snyk-cli/commands/test.md)

Test a project for open source vulnerabilities and license issues.

### [`snyk monitor`](../../features/snyk-cli/commands/monitor.md)

Snapshot and continuously monitor a project for open source vulnerabilities and license issues.

### [`snyk container`](../../features/snyk-cli/commands/container.md)

Test container images for vulnerabilities.

### ``[`snyk iac`](https://docs.snyk.io/snyk-cli/commands/iac)``

Find security issues in Infrastructure as Code files.

### [`snyk code`](../../features/snyk-cli/commands/code.md)

Find security issues using static code analysis.

### [`snyk log4shell`](../../features/snyk-cli/commands/log4shell.md)

Find Log4Shell vulnerability.

### [`snyk config`](../../features/snyk-cli/commands/config.md)

Manage Snyk CLI configuration.

### [`snyk policy`](../../features/snyk-cli/commands/policy.md)

Display the `.snyk` policy for a package.

### [`snyk ignore`](../../features/snyk-cli/commands/ignore.md)

Modify the `.snyk` policy to ignore stated issues.

## New CLI commands

### [snyk fix](../../features/snyk-cli/fix-vulnerabilities-from-the-cli/automatic-remediation-with-snyk-fix.md)

Apply the recommended updates for supported ecosystems automatically.

### [snyk apps](../create-a-snyk-app-using-the-snyk-cli.md)

Create a Snyk App using the Snyk CLI.

### `snyk unmanaged test` and `snyk unmanaged monitor`

For information about these commands see [Snyk CLI for C / C++ projects](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-c-c++#snyk-cli-for-c-c++-projects) and the sections that follow on the page.

## Sub-commands of CLI commands

The following is a list of the sub-commands for Snyk CLI commands. Each sub-command is followed by the command(s) to which the sub-command applies. The commands are linked to their help docs. For details concerning each sub-command, see the help docs.

`test`: sub-command of [`code`](https://docs.snyk.io/features/snyk-cli/commands/code), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container), and [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)

`monitor`: sub-command of [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)

`get <KEY>`: sub-command of [`config`](https://docs.snyk.io/features/snyk-cli/commands/config)\`\`

`set <KEY>=<VALUE>`: sub-command of [`config`](https://docs.snyk.io/features/snyk-cli/commands/config)\`\`

`unset <KEY>`: sub-command of [`config`](https://docs.snyk.io/features/snyk-cli/commands/config)\`\`

`clear`: sub-command of [`config`](https://docs.snyk.io/features/snyk-cli/commands/config)\`\`

## Exit codes

Possible exit codes and their meaning:

**0**: success, no vulnerabilities found\
**1**: action\_needed, vulnerabilities found\
**2**: failure, try to re-run command\
**3**: failure, no supported projects

## Configure the Snyk CLI

You can use environment variables to configure the Snyk CLI and also set variables to configure the Snyk CLI to connect with the Snyk API. See [Configure the Snyk CLI](https://docs.snyk.io/features/snyk-cli/configure-the-snyk-cli).

## Debug

Use `-d` option to output the debug logs.

## Options

The following is a list of the options for Snyk CLI commands. Each option is followed by the command(s) to which the option applies. The commands are linked to their help docs. For details concerning each option, see the help docs.

`--all-projects`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--detection-depth=<DEPTH>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--exclude=<DIRECTORY>[,<DIRECTORY>]...>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--prune-repeated-subdependencies, -p`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--print-deps`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [container](https://docs.snyk.io/features/snyk-cli/commands/container)

`--remote-repo-url=<URL>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--dev`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--org=<ORG_ID>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`code`](https://docs.snyk.io/features/snyk-cli/commands/code), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--file=<FILE>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [container](https://docs.snyk.io/features/snyk-cli/commands/container)

`--package-manager=<PACKAGE_MANAGER_NAME>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--ignore-policy`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--trust-policies` [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--show-vulnerable-paths=<none|some|all>` [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--project-name=<PROJECT_NAME>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)\`\`

`--target-reference=<TARGET_REFERENCE>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--policy-path=<PATH_TO_POLICY_FILE>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac), [`ignore`](https://docs.snyk.io/features/snyk-cli/commands/ignore)

`--json`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`code`](https://docs.snyk.io/features/snyk-cli/commands/code), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)

`--sarif`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`code`](https://docs.snyk.io/features/snyk-cli/commands/code), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--sarif-file-output=<OUTPUT_FILE_PATH>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--severity-threshold=<low|medium|high|critical>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`code`](https://docs.snyk.io/features/snyk-cli/commands/code), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container), [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--fail-on=<all|upgradable|patchable>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test)

`--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...>`: [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)\`\`

`--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE>]...>`: [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)\`\`

`--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...>`: [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)\`\`

`--project-tags=<TAG>[,<TAG>]...>`: [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)\`\`

`--tags=<TAG>[,<TAG>]...>`: [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor), [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)\`\`

`--exclude-base-image-vulns`: [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)

`--platform=<PLATFORM>`: [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)

`--username=<CONTAINER_REGISTRY_USERNAME>`: [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)

`--password=<CONTAINER_REGISTRY_PASSWORD>`: [`container`](https://docs.snyk.io/features/snyk-cli/commands/container)

`--scan=<TERRAFORM_PLAN_SCAN_MODE>`: [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--rules=<PATH_TO_CUSTOM_RULES_BUNDLE>`: [`iac`](https://docs.snyk.io/features/snyk-cli/commands/iac)\`\`

`--id=<ISSUE_ID>`: [`ignore`](https://docs.snyk.io/features/snyk-cli/commands/ignore)

`--expiry=<EXPIRY>`: [`ignore`](https://docs.snyk.io/features/snyk-cli/commands/ignore)

`--reason=<REASON>`: [`ignore`](https://docs.snyk.io/features/snyk-cli/commands/ignore)

`--path=<PATH_TO_RESOURCE>`: [`ignore`](https://docs.snyk.io/features/snyk-cli/commands/ignore)

## Debug

`-d`: [all](https://docs.snyk.io/features/snyk-cli/commands)

## Options for Maven projects

`--scan-all-unmanaged`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--reachable`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--reachable-timeout=<TIMEOUT>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

## Options for Gradle projects

`--sub-project=<NAME>`, `--gradle-sub-project=<NAME>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--all-sub-projects`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--configuration-matching=<CONFIGURATION_REGEX>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--configuration-attributes=<ATTRIBUTE>[,<ATTRIBUTE>]...`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--reachable`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--reachable-timeout=<TIMEOUT>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--init-script=<FILE`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

## Options for .Net and NuGet projects

`--assets-project-name`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--packages-folder`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--project-name-prefix=<PREFIX_STRING>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--project-name-prefix=my-group/`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

## Options for npm projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

## Options for Yarn projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--yarn-workspaces`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

## Options for CocoaPods projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

## Options for Python projects

`--command=<COMMAND>`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

`--skip-unresolved=true|false`: [`test`](https://docs.snyk.io/features/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor)

## Options for Go projects

Currently the following options are not supported:

`--fail-on=<all|upgradable|patchable>`: `test`

## `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

These options are used with the [`test`](https://docs.snyk.io/features/snyk-cli/commands/test) and [`monitor`](https://docs.snyk.io/features/snyk-cli/commands/monitor) commands. See the help docs for details.
