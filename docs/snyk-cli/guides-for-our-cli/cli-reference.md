# CLI reference

**NAME**

`snyk` - CLI and build-time tool to find & fix known vulnerabilities in open-source dependencies

## SYNOPSIS

`snyk` \[COMMAND\] \[SUBCOMMAND\] \[OPTIONS\] \[PACKAGE\] \[-- COMPILER\_OPTIONS\]

## DESCRIPTION

Snyk helps you find, fix and monitor known vulnerabilities in open-source dependencies.  
For more information see [https://snyk.io](https://snyk.io/)

## Not sure where to start? <a id="Not-sure-where-to-start-"></a>

1. authenticate with `$ snyk auth`
2. test your local project with `$ snyk test`
3. get alerted for new vulnerabilities with `$ snyk monitor`

## COMMANDS

To see command-specific flags and usage, see `help` command, e.g. `snyk container --help` \(available in [Advanced Snyk Container CLI usage](snyk-container/snyk-cli-for-container-security/advanced-snyk-container-cli-usage)\).

Available top-level CLI commands:

`auth` \[API\_TOKEN\]Authenticate Snyk CLI with a Snyk account

`test`Test local project for vulnerabilities

`monitor`Snapshot and continuously monitor your project

`container`Test container images for vulnerabilities. See `snyk container --help` for full instructions.

`iac`Find security issues in your Infrastructure as Code files. See `snyk iac --help` for full instructions.

`code` Find security issues using static code analysis. See `snyk code --help` for full instructions.

`config`Manage Snyk CLI configuration

`protect`Applies the patches specified in your .snyk file to the local file system

`policy`Display the .snyk policy for a package.

`ignore`Modifies the .snyk policy to ignore stated issues

`wizard`Configure your policy file to update, auto patch, and ignore vulnerabilities. Snyk wizard updates your .snyk file.

## OPTIONS

To see command-specific flags and usage, see `help` command, e.g. `snyk container --help`. For advanced usage, we offer language and context-specific flags, listed further down this document.

`--all-projects`\(only in `test` and `monitor` commands\) Auto-detect all projects in the working directory

`--detection-depth`=DEPTH\(only in `test` and `monitor` commands\) Use with --all-projects or --yarn-workspaces to indicate how many subdirectories to search. `DEPTH` must be a number.  
Default: 4 \(the current working directory and 3 sub-directories\/)

`--exclude`=DIRECTORY\[,DIRECTORY\]...&gt;\(only in `test` and `monitor` commands\) Can be used with --all-projects and --yarn-workspaces to indicate sub-directories to exclude. Directories must be comma-separated. If using with `--detection-depth` exclude ignores directories at any level deep.

`--prune-repeated-subdependencies`, `-p`\(only in `test` and `monitor` commands\) Prune dependency trees, removing duplicate sub-dependencies. Will still find all vulnerabilities, but potentially not all of the vulnerable paths.

`--print-deps`\(only in `test` and `monitor` commands\) Print the dependency tree before sending it for analysis.

`--remote-repo-url`=URLSet or override the remote URL for the repository that you would like to monitor.

`--dev`Include devDependencies.Default: scan only production dependencies

`--org`=ORG\_NAMESpecify the ORG\_NAME to run Snyk commands tied to a specific organization. This will influence where will new projects be created after running `monitor` command, some features availability, and private test limits. If you have multiple organizations, you can set a default from the CLI using:

`$ snyk config set org`=ORG\_NAME

Setting a default will ensure all newly monitored projects will be created under your default organization. If you need to override the default, you can use the `--org`=ORG\_NAME argument. Default: uses ORG\_NAME that sets as default in your [Account settings](https://app.snyk.io/account)\`\`

`--file`=FILESets a package file. When testing locally or monitoring a project, you can specify the file that Snyk should inspect for package information. When omitted Snyk will try to detect the appropriate file for your project.

`--ignore-policy`Ignores all set policies. The current policy in the `.snyk` file, Org level ignores and the project policy on snyk.io.

`--trust-policies`Applies and uses ignore rules from your dependencies' Snyk policies, otherwise ignore policies are only shown as a suggestion.

`--show-vulnerable-paths`=none\|some\|allDisplay the dependency paths from the top-level dependencies, down to the vulnerable packages. Doesn't affect output when using JSON `--json` output. Default: some \(a few example paths shown\) false is an alias for none.

`--project-name`=PROJECT\_NAMESpecify a custom Snyk project name.

`--policy-path`=PATH\_TO\_POLICY\_FILE\`Manually pass a path to a Snyk policy file.

`--json`Prints result in JSON format.

`--json-file-output`=OUTPUT\_FILE\_PATH\(only in `test` command\) Save test output in JSON format directly to the specified file, regardless of whether or not you use the `--json` option. This is especially useful if you want to display the human-readable test output via stdout and at the same time save the JSON format output to a file.

`--severity-threshold`=low\|medium\|high\|criticalOnly report vulnerabilities of provided level or higher.

`--fail-on`=all\|upgradable\|patchableOnly fail when there are vulnerabilities that can be fixed.

**all** - fails when there is at least one vulnerability that can be either upgraded or patched.  
**upgradable** - fails when there is at least one vulnerability that can be upgraded.  
**patchable** - fails when there is at least one vulnerability that can be patched.

If vulnerabilities do not have a fix and this option is being used, tests will pass.

`--dry-run`\(only in `protect` command\) Don't apply updates or patches during `protect` command run.

`--` \[COMPILER\_OPTIONS\]Pass extra arguments directly to Gradle or Maven. E.g. `snyk test -- --build-cache`

Below are flags that are influencing CLI behavior for specific projects, languages, and contexts:

### Maven options <a id="Maven-options"></a>

`--scan-all-unmanaged`Auto-detects maven jar, war, and aar files in a given directory.

**NOTE**

custom-built jar files, even with open source dependencies, are out of scope.

`--scan-unmanaged`Auto-detects maven jar, war, and aar files in a given directory, and individual testing can be done with `--file`=FILE\_NAME

`--reachable`\(only in `test` and `monitor` commands\) Analyze your source code to find which vulnerable functions and packages are called.

`--reachable-timeout`=TIMEOUTThe amount of time \(in seconds\) to wait for Snyk to gather reachability data. If it takes longer than TIMEOUT, Reachable Vulnerabilities are not reported. This does not affect regular test or monitor output. Default: 300 \(5 minutes\).

### Gradle options <a id="Gradle-options"></a>

[More information about Gradle CLI options](https://snyk.co/ucT6P)\`\`

`--sub-project`=NAME, `--gradle-sub-project`=NAMEFor Gradle "multi-project" configurations, test a specific sub-project.

`--all-sub-projects`For "multi-project" configurations, test all sub-projects.

`--configuration-matching`=CONFIGURATION\_REGEXResolve dependencies using only configuration\(s\) that match the provided Java regular expression, e.g. `^releaseRuntimeClasspath$`.

`--configuration-attributes`=ATTRIBUTE\[,ATTRIBUTE\]...Select certain values of configuration attributes to resolve the dependencies. E.g. `buildtype:release,usage:java-runtime`

### .Net & NuGet options <a id="-Net-NuGet-options"></a>

`--assets-project-name`When monitoring a .NET project using NuGet `PackageReference` use the project name in project.assets.json, if found.

`--packages-folder`Custom path to packages folder

### npm options <a id="npm-options"></a>

`--strict-out-of-sync`=true\|falseControl testing out of sync lockfiles.

Default: true

### Yarn options <a id="Yarn-options"></a>

`--strict-out-of-sync`=true\|falseControl testing out of sync lockfiles. Default: true

`--yarn-workspaces`\(only in `test` and `monitor` commands\) Detect and scan yarn workspaces. You can specify how many sub-directories to search using `--detection-depth` and exclude directories using `--exclude`.

### CocoaPods options <a id="CocoaPods-options"></a>

`--strict-out-of-sync`=true\|falseControl testing out of sync lockfiles.

Default: false

### Python options <a id="Python-options"></a>

`--command`=COMMANDIndicate which specific Python commands to use based on the Python version. The default is `python` which executes your system's default python version. Run 'python -V' to find out what version is it. If you are using multiple Python versions, use this parameter to specify the correct Python command for execution. Default: `python` Example: `--command=python3`

`--skip-unresolved`=true\|falseAllow skipping packages that are not found in the environment.

### Go options <a id="Python-options"></a>

Please note that currently the following flags are not supported:

`--fail-on`=all\|upgradable\|patchable

### Flags available across all commands <a id="Flags-available-accross-all-commands"></a>

`--insecure`Ignore unknown certificate authorities.

`-d`Output debug logs.

`--quiet`, `-q`Silence all output.

`--version`, `-v`Prints versions.\[

COMMAND\] `--help`, `--help` \[COMMAND\], `-h`Prints a help text. You may specify a COMMAND to get more details.

## EXAMPLES <a id="EXAMPLES"></a>

`Authenticate in your CI without user interaction`$ snyk auth MY\_API\_TOKEN`Test a project in current folder for known vulnerabilities`$ snyk test`Test a specific dependency for vulnerabilities`$ snyk test ionic@1.6.5

More examples:

```text
$ snyk test --show-vulnerable-paths=false
$ snyk monitor --org=my-team
$ snyk monitor --project-name=my-project
```

### Container scanning <a id="Container-scanning"></a>

See `snyk container --help` for more details and examples:

```text
$ snyk container test ubuntu:18.04 --org=my-team
$ snyk container test app:latest --file=Dockerfile --policy-path=path/to/.snyk
```

### Infrastructure as Code \(IAC\) scanning <a id="Infrastructure-as-Code-IAC-scanning"></a>

See `snyk iac --help` for more details and examples:

```text
$ snyk iac test path/to/Kubernetes.yaml
$ snyk iac test path/to/terraform_file.tf
```

## EXIT CODES <a id="EXIT-CODES"></a>

Possible exit codes and their meaning:

**0**: success, no vulns found  
**1**: action\_needed, vulns found  
**2**: failure, try to re-run the command  
**3**: failure, no supported projects detected

## ENVIRONMENT <a id="ENVIRONMENT"></a>

You can set these environment variables to change CLI run settings. These keys will override settings in your `snyk config`:

`SNYK_TOKEN`Snyk authorization token. Setting this envvar will override the token that may be available in your `snyk config` settings.  
[How to get your account token](https://snyk.co/ucT6J/)  
[How to use Service Accounts](https://snyk.co/ucT6L)\`\`

`SNYK_API`Sets API host to use for Snyk requests. Useful for on-premise instances and configuring proxies.

`SNYK_CFG_<KEY>`Allows you to override any key that's also available as `snyk config` option.

## NOTICES <a id="NOTICES"></a>

### Snyk API usage policy <a id="Snyk-API-usage-policy"></a>

The use of Snyk's API, whether through the use of the 'snyk' npm package or otherwise, is subject to the [terms & conditions](https://snyk.co/ucT6N/)

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page/)
{% endhint %}

