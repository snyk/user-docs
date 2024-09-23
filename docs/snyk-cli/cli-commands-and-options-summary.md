# CLI commands and options summary

{% hint style="info" %}
This page **only summarizes** the CLI commands and the options for each command. **For details, use the links in this summary to open the help** docs page for the command you are using. The help docs pages are the same as the help in the CLI.
{% endhint %}

## Usage

`snyk [COMMAND] [SUBCOMMAND] [OPTIONS] [PACKAGE] [CONTEXT-SPECIFIC-OPTIONS]`

## Description

The Snyk CLI is a build-time tool to find and fix known vulnerabilities in your projects. For a more detailed description of Snyk CLI and Snyk, see [Snyk CLI](https://docs.snyk.io/snyk-cli). For an introduction on how to use the Snyk CLI, see [Getting started with the CLI](https://docs.snyk.io/snyk-cli/getting-started-with-the-cli).

## Available CLI commands

To learn more about each Snyk CLI command, use the `--help` option, for example, `snyk auth --help` or `snyk container --help`. Each command in this list is linked to the corresponding help page in these docs.

**Note:** Lists of all the options for Snyk CLI commands are on this page. The options are explained in detail in the help for each command.

### [`snyk auth`](https://docs.snyk.io/snyk-cli/commands/auth)

Authenticate Snyk CLI with a Snyk account.

### [`snyk config`](https://docs.snyk.io/snyk-cli/commands/config)

Manage Snyk CLI configuration.

### [`snyk config environment`](https://docs.snyk.io/snyk-cli/commands/config-environment)

Use to set your environment for the region before you run the `snyk auth` command.

### [`snyk test`](https://docs.snyk.io/snyk-cli/commands/test)

Test a Project for open-source vulnerabilities and license issues.

### [`snyk monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

Snapshot and continuously monitor a project for open-source vulnerabilities and license issues.

### [`snyk code`](https://docs.snyk.io/snyk-cli/commands/code)

`Print the name of the snyk code` command `with its help option: snyk code test`

### [`snyk code test`](https://docs.snyk.io/snyk-cli/commands/code-test)

Test source code for any known security issues (Static Application Security Testing).

### [`snyk container`](https://docs.snyk.io/snyk-cli/commands/container)

Print a list of the `snyk container` commands, `snyk container monitor` and `snyk container test`.

### [`snyk container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

Capture the container image layers and dependencies and monitor for vulnerabilities on [snyk.io](https://snyk.io).

### [`snyk container SBOM`](https://docs.snyk.io/snyk-cli/commands/container-sbom)

Generate an SBOM for a container image

### [`snyk container test`](https://docs.snyk.io/snyk-cli/commands/container-test)

Test container images for any known vulnerabilities.

### [`snyk iac`](https://docs.snyk.io/snyk-cli/commands/iac)

Print a list of the `snyk iac` commands: `snyk iac describe`, `snyk iac update-exclude-policy`, and `snyk iac test`.

### [`snyk iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

Test for any known security issue.

### [`snyk iac capture`](https://docs.snyk.io/snyk-cli/commands/iac-capture)

Generate a mapping artifact that contains the minimum amount of information needed to generate resource mappings from code to Cloud from Terraform state files, such as resource IDs and names, and send the mapping artifact to Snyk.

### [`snyk iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

Detect, track, and alert on infrastructure drift and unmanaged resources.

### [`snyk iac rules init`](https://docs.snyk.io/snyk-cli/commands/iac-rules-init)

Initialize custom rules project structure, relation, rule, or spec

### [`snyk iac rules test`](https://docs.snyk.io/snyk-cli/commands/iac-rules-test)

Run tests for all custom rules.

### [`snyk iac rules push`](https://docs.snyk.io/snyk-cli/commands/iac-rules-push)

Bundle and upload custom rule bundles to Snyk Cloud API.

### [`snyk iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

Generate exclude policy rules to be used by `snyk iac describe`.

### [`snyk ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

Modify the `.snyk` policy to ignore stated issues.

### [`snyk log4shell`](https://docs.snyk.io/snyk-cli/commands/log4shell)

Find Log4Shell vulnerability.

### [`snyk policy`](https://docs.snyk.io/snyk-cli/commands/policy)

generates a mapping artifact that contains the minimum amount of information needed to generate, from Terraform state files, resource mappings from code to Cloud, such as resource IDs and names, and sends the mapping artifact to Snyk.Display the `.snyk` policy for a package.

### [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

Generate an SBOM for a local software project in an ecosystem supported by Snyk.

### [`snyk sbom test`](https://docs.snyk.io/snyk-cli/commands/sbom-test)

Checks an SBOM for vulnerabilities in open-source packages.

## New CLI commands

### [`snyk fix`](https://docs.snyk.io/snyk-cli/fix-vulnerabilities-from-the-cli/automatic-remediation-with-snyk-fix)

Apply the recommended updates for supported ecosystems automatically. For more information, see [Automatic fixing with `snyk fix`](scan-and-maintain-projects-using-the-cli/automatic-fixing-with-snyk-fix.md).

### [`snyk apps`](https://docs.snyk.io/snyk-cli/create-a-snyk-app-using-the-snyk-cli)

Create a Snyk App using the Snyk CLI. For more information, see [Snyk Apps](../snyk-api/how-to-use-snyk-apps-apis/).

## Subcommands of CLI commands

The following is a list of the sub-commands for Snyk CLI commands. Each sub-command is followed by the command(s) to which the sub-command applies. The commands are linked to their help docs. For details concerning each sub-command, see the help docs.

`get <KEY>`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

`set <KEY>=<VALUE>`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

`unset <KEY>`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

`clear`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

`environment`: subcommand of [`config`](https://docs.snyk.io/snyk-cli/commands/config)

## Configure the Snyk CLI

You can use environment variables to configure the Snyk CLI and also set variables to configure the Snyk CLI to connect with the Snyk API. See [Configure the Snyk CLI](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli).

## Debug

Use `-d` option to output the debug logs for any command.

## Exit codes for CLI commands

Exit codes for the `test` commands are all the same. See the exit codes in the following help docs:

* [`snyk test` exit codes](https://docs.snyk.io/snyk-cli/commands/test#exit-codes)
* [`snyk container test` exit codes](https://docs.snyk.io/snyk-cli/commands/container-test#exit-codes)
* [`snyk iac test` exit codes](https://docs.snyk.io/snyk-cli/commands/iac-test#exit-codes)
* [`snyk code test` exit codes](https://docs.snyk.io/snyk-cli/commands/code-test#exit-codes)

Additional CLI commands have exit codes as listed in the following help docs:

* [`snyk monitor` exit codes](https://docs.snyk.io/snyk-cli/commands/container-monitor#exit-codes)
* [`snyk container monitor` exit codes](https://docs.snyk.io/snyk-cli/commands/container-monitor#exit-codes)
* [`snyk iac describe` exit codes](https://docs.snyk.io/snyk-cli/commands/iac-describe#exit-codes)
* [`snyk iac update-exclude-policy` exit codes](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy#exit-codes)
* [`snyk log4shell` exit codes](https://docs.snyk.io/snyk-cli/commands/log4shell#exit-codes)
* [`snyk sbom` exit codes](https://docs.snyk.io/snyk-cli/commands/sbom#exit-codes)
* [`snyk container sbom` exit codes](https://docs.snyk.io/snyk-cli/commands/container-sbom#exit-codes)

## Options for multiple commands

Lists of the options for Snyk CLI commands follow. Each option is followed by the command(s) to which the option applies. The commands are linked to their help docs. For details concerning each option, see the [help docs](https://docs.snyk.io/snyk-cli/commands).

`--all-projects`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor),[`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--fail-fast`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--detection-depth=<DEPTH>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac/iac-test), [`sbom`](commands/sbom.md)

`--exclude=<NAME>[,<NAME>]...>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](commands/sbom.md)

`--prune-repeated-subdependencies, -p`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](commands/sbom.md)

`--print-deps`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test)

`--remote-repo-url=<URL>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [iac test](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--dev`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](commands/sbom.md)

`--org=<ORG_ID>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [container monitor](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe), [`iac capture`](https://docs.snyk.io/snyk-cli/commands/iac-capture), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom), [`container sbom`](https://docs.snyk.io/snyk-cli/commands/container-sbom)

`--file=<FILE>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--file=<FILE_PATH>`: [container test](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`sbom test`](https://docs.snyk.io/snyk-cli/commands/sbom-test)

`--package-manager=<PACKAGE_MANAGER_NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--unmanaged:` [test](commands/test.md)`,` [monitor](commands/monitor.md). See also [Options for scanning using `--unmanaged`](https://docs.snyk.io/snyk-cli/cli-reference#options-for-scanning-using-unmanaged) and the [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom) command help for another use of this option.

`--ignore-policy`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--trust-policies` [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--show-vulnerable-paths=<none|some|all>` [`test`](https://docs.snyk.io/snyk-cli/commands/test)

`--project-name=<PROJECT_NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--target-reference=<TARGET_REFERENCE>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)`,`[`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--policy-path=<PATH_TO_POLICY_FILE>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe), [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--json`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe), [`sbom test`](https://docs.snyk.io/snyk-cli/commands/sbom-test)

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test),[`sbom`](https://docs.snyk.io/snyk-cli/commands/test)

`--sarif`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--sarif-file-output=<OUTPUT_FILE_PATH>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--severity-threshold=<low|medium|high|critical>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--fail-on=<all|upgradable|patchable>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [test](https://docs.snyk.io/snyk-cli/commands/test)

`--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--project-tags=<TAG>[,<TAG>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--tags=<TAG>[,<TAG>]...>`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

## `snyk auth` command options

\--`auth-type=<TYPE>`\
`--client-secret=<SECRET>`\
`--client-id=<ID>` [`snyk auth`](https://docs.snyk.io/snyk-cli/commands/auth)

## `snyk code test` command option&#x20;

`--include-ignores`: [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test)

## `snyk code test` and `snyk iac test` command option

`--report`: [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test),  [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

## `snyk config environment command option`

`--no-check` [snyk config environment](https://docs.snyk.io/snyk-cli/commands/config-environment)

## `snyk container` command options

`--app-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container`monitor](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--exclude-app-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`container sbom`](https://docs.snyk.io/snyk-cli/commands/container-sbom)

`--nested-jars-depth`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--exclude-base-image-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--platform=<PLATFORM>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--username=<CONTAINER_REGISTRY_USERNAME>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--password=<CONTAINER_REGISTRY_PASSWORD>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

## `snyk iac test` command options

`--scan=<TERRAFORM_PLAN_SCAN_MODE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--target-name=<TARGET_NAME>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--rules=<PATH_TO_CUSTOM_RULES_BUNDLE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--var-file=<PATH_TO_VARIABLE_FILE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

## `snyk iac capture` command options

`--stdin`: [`iac capture`](https://docs.snyk.io/snyk-cli/commands/iac-capture)

`PATH`: [`iac capture`](https://docs.snyk.io/snyk-cli/commands/iac-capture)

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

## `snyk iac update-exclude-policy` command options

`--exclude-changed`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

`--exclude-missing`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

`--exclude-unmanaged`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

## `snyk iac rules push` command option

`--delete`: [`iac rules push`](commands/iac-rules-push.md)

## `snyk iac rules test` command option

`--update-expected`: [`iac rules test`](commands/iac-rules-test.md)

## `snyk ignore` command options

`--id=<ISSUE_ID>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--expiry=<EXPIRY>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--reason=<REASON>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--path=<PATH_TO_RESOURCE>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

## `snyk sbom` and `snyk container sbom` command options

`--format=<cyclonedx1.4+json|cyclonedx1.4+xml|spdx2.3+json>`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom), [snyk container sbom](https://docs.snyk.io/snyk-cli/commands/container-sbom)

`[--file=] or [--f=]`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`[--name=<NAME>]`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`[--version=<VERSION>]`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`[<TARGET_DIRECTORY>]`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`<IMAGE>`: [`snyk container sbom`](https://docs.snyk.io/snyk-cli/commands/container-sbom)

## Option for Maven projects

`--maven-aggregate-project`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--scan-unmanaged`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--scan-all-unmanaged`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Gradle projects

`--sub-project=<NAME>`, `--gradle-sub-project=<NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--all-sub-projects`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--configuration-matching=<CONFIGURATION_REGEX>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--configuration-attributes=<ATTRIBUTE>[,<ATTRIBUTE>]...`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--init-script=<FILE`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for .Net and NuGet projects

`--file=.sln`: [test](https://docs.snyk.io/snyk-cli/commands/test)

`--file=<filename>.sln`: [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

\-`-file=packages.config`: [test](https://docs.snyk.io/snyk-cli/commands/test), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--assets-project-name`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--packages-folder`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--project-name-prefix=<PREFIX_STRING>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--project-name-prefix=my-group/`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--dotnet-runtime-resolution`:  `test,` [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--dotnet-target-framework`: `test,` [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for npm projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [sbom](https://docs.snyk.io/snyk-cli/commands/sbom)

## Options for pnpm projects

`--dev:` [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

\--all-projects: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

\--fail-on: [`test`](https://docs.snyk.io/snyk-cli/commands/test)

\--prune-repeated-subdependencies: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Yarn projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--yarn-workspaces`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

## Options for CocoaPods projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Python projects

`--command=<COMMAND>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--skip-unresolved=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--File=<filename>:` [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--pakage-manager=<package manager>`: [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

## Options for Go projects

The following options are not supported:

`--fail-on=<all|upgradable|patchable>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test)

## Options for scanning using `--unmanaged`

`--org=<ORG_ID>`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--json`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md)

`--remote-repo-url=<URL>`: [`test`](commands/test.md)

`--severity-threshold=<low|medium|high|critical>:` [`test`](commands/test.md)

`--target-reference=<TARGET_REFERENCE>`: [`test`](commands/test.md), [monitor](commands/monitor.md)

`--max-depth`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--print-dep-paths:` [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--project-name=c-project`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

These options are used with the `snyk test` and `snyk monitor` commands. See the help docs for [`snyk test`](https://docs.snyk.io/snyk-cli/commands/test) and [`snyk monitor`](https://docs.snyk.io/snyk-cli/commands/monitor) for details.
