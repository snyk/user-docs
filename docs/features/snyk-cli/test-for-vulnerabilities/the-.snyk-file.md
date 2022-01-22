# The .snyk file

The `.snyk` file is a policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins.

The file can be generated in a number of ways and can be used in a number of different scenarios. The `.snyk` file is generally located at the root of your project.

This page provides detailed information about the contents and use of the `.snyk` file as well as about creating the file.

{% hint style="info" %}
You can create and save a policy where no attributes are selected, for example, if you have not yet decided which attributes should be associated with that policy. This policy does not apply to projects where all attributes are left blank.
{% endhint %}

The following video provides an introduction to the `.snyk` file.

{% embed url="https://youtu.be/QSIBt-hQ0Xo" %}

### Capabilities and behaviors

The `.snyk` policy file in a project is used to apply ignores and other settings for the `snyk test`and `snyk monitor` commands, as well as any tests done through the API or website. For IaC ignore rules see [IaC ignores using the .snyk policy file](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file).

* The `.snyk` file defines **Snyk patches** to be applied at build time, to resolve vulnerabilities that cannot be fixed with upgrades or do not have a version to get to (npm and Yarn only).
* The `.snyk` file defines **Ignores**.
  * Snyk checks the Snyk database and the `.snyk` policy file for ignore rules when performing CLI and CI/CD scanning.
  * The `.snyk` policy file is used to apply ignores and other settings for the `snyk test` and `snyk monitor` commands, as well as any tests through the API or website.
  * **Note**: if **admin users only** is enabled (**Settings > General > Ignores**), the ignore rules in the database are used, unless there is a `.snyk` file in the project. If there a `.snyk` file in the project, **** `snyk test` uses that file as the ignore mechanism, instead of the ignores set from the web UI.
  * Developers can ignore issues by using the `.snyk` policy file when running`snyk monitor`.
  * When the `.snyk` file is included in a Source Control Management (SCM) project, Snyk considers both the database ignores and the `.snyk` ignores.
  * Specify the project-level Python version in SCM or CLI scans.
* The `.snyk` file defines certain **analysis configuration items** such as language settings and Python version.
  * CLI and CI/CD setting: the language setting for the current project
  * SCM scans (for example, GitHub): the Snyk web UI currently limits users to setting Python versions at the organization level. When you include the `.snyk` file in your code repository and the language setting value is set, then when you run code repository scans you gain the advantage of creating project-level Python settings. You may need to re-import the project if the `.snyk` file was not present at the initial import of the project into Snyk.

### `.snyk` file creation

The `.snyk` file can be created in a number of ways:

* **Snyk vulnerability fix pull request (PR)** - When you select the **fix a vulnerability** button on a git code repository scan, and a Snyk patch is available and an upgrade is not possible, a `.snyk` file is added to the pull request. (Currently Snyk patches are for npm and Yarn only.)
* **Snyk CLI** - Using the `snyk ignore` command creates a `.snyk` file..
* **Manual creation** - You can create a new `.snyk` file and populate it with the code that follows. The version is the current version of the snyk-policy package; you can find this at [https://www.npmjs.com/package/snyk-policy](https://www.npmjs.com/package/snyk-policy). Note that in order to ignore by path you must edit the `.snyk` file manually.

```
 # Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
 version: v1.22.1
```

### Syntax

The `.snyk` file has the following content:

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
language-settings:
ignore: 
patch:
```

The `language-settings:` value is the Python version you are currently using. See the example  [Setting the language version for Python](the-.snyk-file.md#setting-the-language-version-for-python).

The `patch`: is in the form of

```
'npm:library:yyyymmdd’ :
  - path to library using > seperator:
    patched: 'datetime string'
  - path to library using > seperator > to > another > path:
    patched: 'datetime string'
```

The `ignore:` is an ignore rule in the form of

```
Ignore:
  snyk-vulnid:
    - path to library using > seperator :
      reason: 'text string'
      expires: 'datetime string'
```

### Monorepos and complex project considerations

Snyk CLI expects the `.snyk` file to be relative to the manifest being analyzed. In the case of a complex project or monorepo, there may be many manifests in subfolders, and you may wish to use a centralized ignore policy.

If you create a `.snyk` ignore policy with the CLI and Snyk does not successfully ignore the vulnerability, use the option `--policy-path=/path/path/file.`

Your complete statement should be `snyk ignore --id=IssueID [--expiry=expiry] [--reason='reason for ignoring'] [--policy-path=/path/path/file].`

Ignores can be specified in the web interface only after an issue is detected and monitored; thus the advantage of using the `.snyk` policy file is to reduce friction and be proactive.

The ignore rules can be overridden if admin users only is enabled for the relevant organization (click on **Settings** > **General** > **Ignores**).

### Examples

#### Creating a .snyk  file

Generate a patch rule using a vulnerability fix PR:

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
ignore: {}
# patches apply the minimum changes required to fix a vulnerability
patch:
  'npm:hawk:20160119':
    - tap > codecov.io > request > hawk:
        patched: '2020-01-20T14:26:34.404Z'
```

#### Setting the language version for Python

Manually modify the `.snyk` file to set `language-settings:` for the project to Python 2.7:

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
language-settings: 
python: "2.7"
```

Manually modify the `.snyk` file to set `language-settings:` for the project to Python 3.6.2:

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
language-settings: 
python: "3.6.2"
```

**Note:** When you include the `.snyk` file in your code repository and the `language-settings:` value is set, then when you run code repository scans you gain the advantage of creating project-level Python settings.&#x20;

#### Setting vulnerability ignore rules

Ignore a specific vulnerability for a given path:

```
ignore:
 SNYK-JS-BSON-561052:
    - mongodb > mongodb-core > bson:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
```

Ignore a vulnerability for all paths:

```
ignore:
  'npm:mem:20180117':
    - '*':
        reason: None Given
        expires: 2020-04-04T17:33:45.004Z
```

Ignore a specific vulnerability on multiple paths:

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

To ignore the license issue for package, find the  ID for the license in the output of the `snyk test` command.&#x20;

The license ID is part of the license issue URL, for example, in this URL: [https://snyk.io/vuln/snyk:lic:npm:symbol:MPL-2.0](https://snyk.io/vuln/snyk:lic:npm:symbol:MPL-2.0), the license ID is `snyk:lic:npm:symbol:MPL-2.0`.

#### **Ignoring the license with the CLI**

Enter the license ID in lowercase to avoid causing an error. Only the proper name of the license can be in uppercase. In the example that follows, everything is in lowercase except the proper name of the license, GPL-2.0.

`snyk ignore --id=snyk:lic:npm:goof:GPL-2.0`

This command results in the following `.snyk` file:

```
ignore:
  'snyk:lic:npm:goof:GPL-2.0':
    - '*':
        reason: None Given
        expires: 2020-11-07T11:38:28.614Z
```

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

The `.snyk` file should be versioned in the code repository, the same as other applications and build resources.

### Additional information

[Ignore vulnerabilities using Snyk CLI](https://docs.snyk.io/snyk-cli/fix-vulnerabilities-from-the-cli/ignore-vulnerabilities-using-snyk-cli)\
Support KB - [Ignoring via the CLI is not enabled for this organization](https://support.snyk.io/hc/en-us/articles/360001558598)\
[CLI reference](../cli-reference/)
