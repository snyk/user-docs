# The .snyk file

The `.snyk` file is a policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins.

The file can be generated in a number of ways and can be used in a number of different scenarios. The `.snyk` file is generally located at the root of your project.

This article is provided to help you understand the `.snyk` file.

{% hint style="info" %}
You can create and save a policy where no attributes are selected, for example, if you have not yet decided which attributes should be associated with that policy. This policy does not apply to projects where all attributes are left blank.
{% endhint %}

{% embed url="https://youtu.be/QSIBt-hQ0Xo" %}

### Capabilities and behaviors

When present, the `.snyk` policy is used to apply ignores and other settings for the `snyk test`, `snyk protect` and `snyk monitor` commands, as well as any tests done through the API or website.

* The `.snyk` file defines **Snyk patches** to be applied at build time, to resolve vulnerabilities that can't be fixed with upgrades or do not have a version to get to (npm and Yarn only).
* The `.snyk` file defines **Ignores**.
  * Snyk checks the Snyk database and `.snyk` file for ignore rules when performing CLI and CI/CD scanning.
  * When present, the Snyk policy is used to apply ignores and other settings for the `snyk test`, `snyk protect` and `snyk monitor` commands as well as any tests  through the API or website.
  * **Note**: if **admin users only** is enabled (click on ![](<../../../.gitbook/assets/image (90).png>) settings **> General > Ignores**), the ignore rules present in the database are used, unless a `.snyk` file exists in the project. If a `.snyk` file exists, **** `snyk test` uses this file as the ignore mechanism, instead of the ignores set from the web UI.
  * Developers are able to ignore issues with `.snyk` when using `snyk monitor`.
  * When the `.snyk` file is included in an SCM project, Snyk considers both the database ignores and the `.snyk` ignores.
  * Specify the project-level Python version in SCM or CLI scans.
* The `.snyk` file defines certain **analysis configuration items**, such as language settings and Python version.
  * CLI and CI/CD setting: the language setting for the current project
  * SCM scans (that is, Github): the Snyk web UI currently limits users to setting Python versions at the organization level. Including the `.snyk` file in your code repository has the advantage when running code repository scans of creating project-level Python settings when the language-setting value is set, resulting in project-level settings. You may need to re-import the project if the `.snyk` file was not present on the initial import of the project into Snyk.

### `.snyk` file creation

The `.snyk` file can be created in a number of ways

* **Snyk vulnerability fix pull request (PR)** - When a user selects the **fix a vulnerability** button on a git code repository scan, and a Snyk patch is available when an upgrade is not possible, a `.snyk` file is added to the pull request (Currently Snyk patches are for npm and Yarn only.)
* **Snyk CLI Wizard** - The CLI Wizard allows users to generate a `.snyk` policy (_currently available for npm and Yarn only_)
* **Snyk CLI** - Using the `snyk ignore` command (.snyk file must exist)
*   **Manual creation** - A user may decide to create the file manually. Simply create a new

    `.snyk` file and populate it with the following. Note that in order to ignore by path, manual editing is required except when `snyk wizard` is used.

```
 # Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
 version: v1.14.0
```

### Syntax

The `.snyk` file

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.14.0
language-settings:
ignore: 
patch:
```

Where patch is in the form of

```
'npm:library:yyyymmdd’ :
  - path to library using > seperator:
    patched: 'datetime string'
  - path to library using > seperator > to > another > path:
    patched: 'datetime string'
```

And where an ignore rule is in the form of

```
Ignore:
  snyk-vulnid:
    - path to library using > seperator :
      reason: 'text string'
      expires: 'datetime string'
```

### Monorepos and complex project considerations

Snyk CLI expects the `.snyk` file to be relative to the manifest being analyzed. In the case of a complex project or monorepo, there may be many manifests in subfolders, and you may wish to use a centralized ignore policy.

If you create a `.snyk` ignore policy in CLI and Snyk does not successfully ignore the vulnerability, use the option `--policy-path=/path/path/file`

Your complete statement should be `snyk ignore --id=IssueID [--expiry=expiry] [--reason='reason for ignoring'] [--policy-path=/path/path/file]`

Ignores can be specified in the web interface only after an issue is detected and monitored; thus the advantage of using the policy file is to reduce friction and be proactive.

The ignore rules can be overridden if admin users only is enabled (click on **Settings** > **General** > Ignores) for the relevant organization.

### Examples

#### Creating a .snyk  file

Generating a patch rule using a vulnerability fix PR

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.14.0
ignore: {}
# patches apply the minimum changes required to fix a vulnerability
patch:
  'npm:hawk:20160119':
    - tap > codecov.io > request > hawk:
        patched: '2020-01-20T14:26:34.404Z'
```

#### Setting the language version for Python

Manually modifying the `.snyk` file for setting analysis for the project at Python 2.7 \*

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.14.0
language-settings: 
python: "2.7"
```

Manually modifying the `.snyk` file for setting analysis for the project at Python 3.6.2 \*

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.14.0
language-settings: 
python: "3.6.2"
```

**Note:** Including the `.snyk` file in your code repository when running code repository scans provides an added advantage of creating project-level Python languages settings when the language-setting value is set.

#### Setting vulnerability ignore rules

Ignore a specific vulnerability for a given path

```
ignore:
 SNYK-JS-BSON-561052:
    - mongodb > mongodb-core > bson:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
```

Ignore a vulnerability for all paths

```
ignore:
  'npm:mem:20180117':
    - '*':
        reason: None Given
        expires: 2020-04-04T17:33:45.004Z
```

Ignore a specific vulnerability on multiple paths

```
ignore:
 SNYK-JS-DOTPROP-543489:
   - configstore > dot-prop:
       reason: None given
       expires: '2020-06-19T20:36:54.553Z'
   - snyk > configstore > dot-prop:
       reason: None given
       expires: '2020-06-19T20:36:54.553Z'
```

#### Setting license ignore rules

Ignore license issue for package

You can find the ID for a License as a substring of the License issue URL, for example,

#### **Ignoring the license with the CLI**

**Note:** The license ID must be entered in lowercase or it causes an error. In the code that follows, everything is in lowercase except the proper name of the license. The proper name, `GPL-2.0` is the only piece of the command that can be uppercase without causing an error.

```
snyk ignore --id=snyk:lic:npm:goof:GPL-2.0
```

Results in `.snyk` file:

```
ignore:
  'snyk:lic:npm:goof:GPL-2.0':
    - '*':
        reason: None Given
        expires: 2020-11-07T11:38:28.614Z
```

{% hint style="info" %}
**Note**\
For IaC ignore rules see [IaC ignores using the .snyk policy file](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file).
{% endhint %}

### .snyk related CLI commands

The `snyk policy` command displays the `.snyk` policy for a package.

`snyk policy`

The `snyk ignore` command modifies the `.snyk` policy to ignore a stated issue.

```
snyk ignore --id=’vulnerabilityID’ --expiry=’date-string’ --reason=’text string’
```

The following example shows using the `snyk ignore` command to generate a rule to ignore the SNYK-JS-BSON-561052 vulnerability for all paths that lead to that library on disk.

```
snyk ignore --id=’SNYK-JS-BSON-561052’ --expiry=’2018-04-01’ --reason=’testing’
```

### Best practices

The `.snyk` file should be versioned in the code repository, like other applications and build resources.

### Additional information

[Ignore vulnerabilities using Snyk CLI](https://docs.snyk.io/snyk-cli/fix-vulnerabilities-from-the-cli/ignore-vulnerabilities-using-snyk-cli)\
Support KB - [Ignoring via the CLI is not enabled for this organization](https://support.snyk.io/hc/en-us/articles/360001558598)\
[CLI reference](../cli-reference/)\
[Manage vulnerability results with the Snyk CLI Wizard](https://docs.snyk.io/snyk-cli/fix-vulnerabilities-from-the-cli/manage-vulnerability-results-with-the-snyk-cli-wizard) (NPM/Yarn only)
