# The .snyk file

The `.snyk` file is a policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI & CI/CD plugins.

The file can be generated in a number of ways and can be used in a number of different scenarios. The `.snyk` file is generally located at the root of your project.

This article will help you understand the `.snyk` file.

{% hint style="info" %}
You can create and save a policy where no attributes are selected, for example, if you have not yet decided which attributes should be associated with that policy. This policy does not apply to projects where all attributes are left blank.
{% endhint %}

## Capabilities and behaviors

When present, the `.snyk` policy is used to apply ignores and other settings for `snyk test`, `snyk protect` and `snyk monitor` commands as well as any tests via the API or website.

* Defines **Snyk Patches** to be applied at build time, to resolve vulnerabilities that can't be fixed via upgrades or do not have a version to get to \(NPM/Yarn only\).
* Defines **Ignores** 
  * Snyk will check the Snyk database and `.snyk` file for ignore rules when performing CLI & CI/CD scanning. 
    * When present Snyk policy is used to apply ignores and other settings for `snyk test`, `snyk protect` and `snyk monitor` commands as well as any tests via the API or website.
  * **NOTE**: only the database ignore rules are applied if **admin users only is** enabled \(click on  **Settings &gt; General &gt; Ignores**\). The ignore rules already present in the `.snyk` file are applied regardless of the admin setting for the organization. Developers are able to ignore issues via `.snyk` when using `snyk monitor` .
  * When the `.snyk` file is included in an SCM project, Snyk will consider both the database ignores and the `.snyk` ignores.
  * Specify project level Python version in SCM or CLI scans
* Defines certain **analysis configuration items**, such as language settings/python version. 
  * CLI & CI/CD setting the language setting for the current project
  * SCM Scans \(i.e. Github\): The Snyk web UI currently limits users to setting python versions at the organization level. By including the `.snyk` file in your code repository, when running code repository scans, has the added advantage of creating project-level python settings when the language-setting value is set, resulting in project-level settings! You may need to re-import the project if the `.snyk` file if it was not present on the initial import of the project into Snyk.

## `.snyk` file creation

The `.snyk` file can be created in a number of ways

* **Snyk vulnerability Fix Pull Request \(PR\)** - When a user selects to _fix a vulnerability_ button on a git code repository scan, and a Snyk Patch is available when an upgrade is not possible, a `.snyk` file is added to the pull request \(c_urrently Snyk Patches are for NPM and Yarn only\)_
* **Snyk CLI Wizard** - the CLI allows users to generate a `.snyk` policy \(_currently available for NPM/Yarn only_\)
* **Snyk CLI** - utilizing the `snyk ignore` command \(.snyk file must exist\)
* **Manual creation** - a user may decide to create the file manually. Simply create a new

  `.snyk` file and populate it with the following. Note that in order to ignore by path--_manual editing is required except for where the `snyk wizard` is used_

```text
  # Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.version: v1.14.0
```

## Syntax

The `.snyk` file

```text
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.14.0
language-settings:
ignore: 
patch:
```

Where patch is in the form of

```text
'npm:library:yyyymmdd’ :
  - path to library using > seperator:
    patched: 'datetime string'
  - path to library using > seperator > to > another > path:
    patched: 'datetime string'
```

And where an ignore rule is in the form of

```text
Ignore:
  snyk-vulnid:
    - path to library using > seperator :
      reason: 'text string'
      expires: 'datetime string'
```

## Mono-repos and complex project considerations

Snyk CLI expects the `.snyk` file to be relative to the manifest being analyzed, in the case of a complex project or mono-repo, there may be many manifests in subfolders, and you may wish to use a centralized ignore policy.

If you create a `.snyk` ignore policy in CLI and Snyk does not successfully ignore the vulnerability, add the param `--policy-path=/path/path/file`

Your complete statement should be `snyk ignore --id=IssueID [--expiry=expiry] [--reason='reason for ignoring'] [--policy-path=/path/path/file]`

Ignores can be specified in the web interface, only after an issue is detected and monitored, so the advantage of using the policy file is to reduce friction and be proactive.

The ignore rules can be overridden if admin users only is enabled \(click on **Settings** &gt; **General** &gt; Ignores\) for the relevant organization.

## Examples

## .snyk file creation

**Patch rule generated by vulnerability fix PR**

```text
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.14.0
ignore: {}
# patches apply the minimum changes required to fix a vulnerability
patch:
  'npm:hawk:20160119':
    - tap > codecov.io > request > hawk:
        patched: '2020-01-20T14:26:34.404Z'
```

## Setting language version for Python

**Manually modifying the `.snyk` file for setting analysis for the project at Python 2.7 \***

```text
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.14.0
language-settings: python: "2.7"
```

**Manually modifying the `.snyk` file for setting analysis for the project at Python 3.6.2 \***

```text
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.14.0
language-settings: python: "3.6.2"
```

Note by including the `.snyk` file in your code repository, when running code repository scans, this has an added advantage of creating project level python languages settings when the language-setting value is set

## Setting vulnerability ignore rules

**Ignore Example - Specific vulnerability for a given path**

```text
ignore:
 SNYK-JS-BSON-561052:
    - mongodb > mongodb-core > bson:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
```

**Ignore Example - Ignore a vulnerability for all paths**

```text
ignore:
  'npm:mem:20180117':
    - '*':
        reason: None Given
        expires: 2020-04-04T17:33:45.004Z
```

**Ignore specific vulnerability on multiple paths**

```text
ignore:
 SNYK-JS-DOTPROP-543489:
   - configstore > dot-prop:
       reason: None given
       expires: '2020-06-19T20:36:54.553Z'
   - snyk > configstore > dot-prop:
       reason: None given
       expires: '2020-06-19T20:36:54.553Z'
```

## Setting license ignore rules

**Ignore example - license issue for package**

You can find the ID for a License as a substring of the License issue URL, e.g.

Ignoring the license via the CLI:

```text
snyk ignore --id=snyk:lic:npm:goof:GPL-2.0
```

Result in `.snyk` file:

```text
ignore:
  'snyk:lic:npm:goof:GPL-2.0':
    - '*':
        reason: None Given
        expires: 2020-11-07T11:38:28.614Z
```

{% hint style="info" %}
**Note**  
For IaC ignore rules see [IaC ignores using the .snyk policy file](https://snyk.gitbook.io/user-docs/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file).
{% endhint %}

## .snyk related CLI commands

`policy` ............. Display the `.snyk` policy for a package.

```text
snyk policy
```

`ignore` ............. Modifies the `.snyk` policy to ignore stated issues.

```text
snyk ignore --id=’vulnerabilityID’ --expiry=’date-string’ --reason=’text string’
```

**Example - Generating an ignore rule using CLI. Ignore the SNYK-JS-BSON-561052 vulnerability for all paths that lead to that library on disk.**

```text
snyk ignore --id=’SNYK-JS-BSON-561052’ --expiry=’2018-04-01’ –reason=’testing’
```

## Best practices

The `.snyk` file should be versioned in the code repository, like other applications and build resources.

## Additional articles

Support KB - [Ignore Vulnerabilities](https://snyk.gitbook.io/user-docs/snyk-cli/fix-vulnerabilities-from-the-cli/ignore-vulnerabilities-using-snyk-cli)  
Support KB - [Ignoring via the CLI is not enabled for this organization](https://support.snyk.io/hc/en-us/articles/360001558598)  
Support KB - [CLI reference](https://support.snyk.io/hc/en-us/articles/360003812578)  
Support KB - [Manage Vulnerabilities with the Snyk Wizard in the CLI](https://snyk.gitbook.io/user-docs/snyk-cli/fix-vulnerabilities-from-the-cli/manage-vulnerability-results-with-the-snyk-cli-wizard) \(NPM/Yarn Only\)

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

