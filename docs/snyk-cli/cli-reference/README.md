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

`test`: subcommand of [`code`](../commands/code.md), [`container`](../commands/container/), and [`iac`](../commands/iac-1.md)\`\`

`monitor`: subcommand of [`container`](../commands/container/)\`\`

`get <KEY>`: subcommand of [`config`](../commands/config.md)\`\`

`set <KEY>=<VALUE>`: subcommand of [`config`](../commands/config.md)\`\`

`unset <KEY>`: sub-command of [`config`](../commands/config.md)\`\`

`clear`: subcommand of [`config`](../commands/config.md)\`\`

## Exit codes

Possible exit codes and their meaning:

**0**: success, no vulnerabilities found\
**1**: action\_needed, vulnerabilities found\
**2**: failure, try to re-run command\
**3**: failure, no supported projects

## Configure the Snyk CLI

You can use environment variables to configure the Snyk CLI and also set variables to configure the Snyk CLI to connect with the Snyk API. See [Configure the Snyk CLI](../configure-the-snyk-cli.md).

## Debug

Use `-d` option to output the debug logs.

## Options

The following is a list of the options for Snyk CLI commands. Each option is followed by the command(s) to which the option applies. The commands are linked to their help docs. For details concerning each option, see the help docs.

`--all-projects`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--detection-depth=<DEPTH>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`iac`](../commands/iac-1.md)\`\`

`--exclude=<DIRECTORY>[,<DIRECTORY>]...>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--prune-repeated-subdependencies, -p`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--print-deps`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`container`](../commands/container/)\`\`

`--remote-repo-url=<URL>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--dev`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--org=<ORG_ID>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`code`](../commands/code.md), [`iac`](../commands/iac-1.md)\`\`

`--file=<FILE>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`container`](../commands/container/)\`\`

`--package-manager=<PACKAGE_MANAGER_NAME>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--ignore-policy`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`iac`](../commands/iac-1.md)\`\`

`--trust-policies` [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--show-vulnerable-paths=<none|some|all>` [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--project-name=<PROJECT_NAME>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`container`](../commands/container/)\`\`

`--target-reference=<TARGET_REFERENCE>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--policy-path=<PATH_TO_POLICY_FILE>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`container`](../commands/container/), [`iac`](../commands/iac-1.md), [`ignore`](../commands/ignore.md)\`\`

`--json`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`code`](../commands/code.md), [`container`](../commands/container/), [`iac`](../commands/iac-1.md)\`\`

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`container`](../commands/container/), [`iac`](../commands/iac-1.md)\`\`

`--sarif`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`code`](../commands/code.md), [`container`](../commands/container/), [`iac`](../commands/iac-1.md)\`\`

`--sarif-file-output=<OUTPUT_FILE_PATH>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md), [`container`](../commands/container/), [`iac`](../commands/iac-1.md)\`\`

`--severity-threshold=<low|medium|high|critical>`: [`test`](../commands/test.md), [`code`](../commands/code.md), [`container`](../commands/container/), [`iac`](../commands/iac-1.md)\`\`

`--fail-on=<all|upgradable|patchable>`: [`test`](../commands/test.md)

`--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...>`: [`monitor`](../commands/monitor.md), [`container`](../commands/container/)\`\`

`--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE>]...>`: [`monitor`](../commands/monitor.md), [`container`](../commands/container/)\`\`

`--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...>`: [`monitor`](../commands/monitor.md), [`container`](../commands/container/)\`\`

`--project-tags=<TAG>[,<TAG>]...>`: [`monitor`](../commands/monitor.md), [`container`](../commands/container/)\`\`

`--tags=<TAG>[,<TAG>]...>`: [`monitor`](../commands/monitor.md), [`container`](../commands/container/)\`\`

`--exclude-base-image-vulns`: [`container`](../commands/container/)\`\`

`--platform=<PLATFORM>`: [`container`](../commands/container/)\`\`

`--username=<CONTAINER_REGISTRY_USERNAME>`: [`container`](../commands/container/)\`\`

`--password=<CONTAINER_REGISTRY_PASSWORD>`: [`container`](../commands/container/)\`\`

`--scan=<TERRAFORM_PLAN_SCAN_MODE>`: [`iac`](../commands/iac-1.md)\`\`

`--rules=<PATH_TO_CUSTOM_RULES_BUNDLE>`: [`iac`](../commands/iac-1.md)\`\`

`--id=<ISSUE_ID>`: [`ignore`](../commands/ignore.md)\`\`

`--expiry=<EXPIRY>`: [`ignore`](../commands/ignore.md)\`\`

`--reason=<REASON>`: [`ignore`](../commands/ignore.md)\`\`

`--path=<PATH_TO_RESOURCE>`: [`ignore`](../commands/ignore.md)\`\`

## Debug

`-d`: [all](../commands/)

## Options for Maven projects

`--scan-all-unmanaged`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--reachable`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--reachable-timeout=<TIMEOUT>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

## Options for Gradle projects

`--sub-project=<NAME>`, `--gradle-sub-project=<NAME>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--all-sub-projects`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--configuration-matching=<CONFIGURATION_REGEX>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--configuration-attributes=<ATTRIBUTE>[,<ATTRIBUTE>]...`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--reachable`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--reachable-timeout=<TIMEOUT>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--init-script=<FILE`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

## Options for .Net and NuGet projects

`--assets-project-name`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--packages-folder`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--project-name-prefix=<PREFIX_STRING>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--project-name-prefix=my-group/`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

## Options for npm projects

`--strict-out-of-sync=true|false`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

## Options for Yarn projects

`--strict-out-of-sync=true|false`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--yarn-workspaces`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

## Options for CocoaPods projects

`--strict-out-of-sync=true|false`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

## Options for Python projects

`--command=<COMMAND>`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

`--skip-unresolved=true|false`: [`test`](../commands/test.md), [`monitor`](../commands/monitor.md)\`\`

## Options for Go projects

Currently the following options are not supported:

`--fail-on=<all|upgradable|patchable>`: [`test`](../commands/test.md)\`\`

## Options for unmanaged projects

For information about these options see [Snyk CLI for C / C++ projects](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-c-c++#snyk-cli-for-c-c++-projects).

## `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

These options are used with the `test` and `monitor` commands. See the [help docs](../commands/) for details.
