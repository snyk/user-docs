# The .snyk file

The `.snyk` file is a YAML [policy](./) file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins.

The file can be generated in a number of ways and can be used in a number of different scenarios. The `.snyk` file is generally located at the root of your Project.

This page provides detailed information about the contents and use of the `.snyk` file, as well as about creating the file.

The following video provides an introduction to the `.snyk` file.

{% embed url="https://youtu.be/QSIBt-hQ0Xo" %}

## Capabilities and behaviors of the `.snyk` file

The `.snyk` policy file in a Project is used to apply ignores and other settings for the `snyk test` and `snyk monitor` commands and for any tests done through the API or Snyk Web UI. and to apply the `@snyk/protect` [package](https://github.com/snyk/snyk/tree/master/packages/snyk-protect) that replaced the `snyk protect` command.

The `.snyk` file can also be used to exclude directories and files from repositories that are imported for Snyk Code testing. The `exclude from import` option is supported only in Snyk Code, and only for imports that are performed using the Snyk Web UI and CLI. For details, see [Excluding directories and files from the import process](../../scan-application-code/snyk-code/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process.md).

For IaC ignore rules, see [IaC ignores using the .snyk policy file](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/iac-ignores-using-the-.snyk-policy-file).

The following explains how the `.snyk` file works.

* The `.snyk` file defines **Snyk patches** to be applied at build time, to resolve vulnerabilities that cannot be fixed with upgrades.
* The `.snyk` file defines **Ignores**.
* Snyk checks the Snyk database and the `.snyk` policy file for ignore rules when scanning through an SCM integration, the Snyk CLI, and a CI/CD integration.
* If **Admin users only** is enabled (by using **Settings > General > Ignores**), the ignore rules in the database are used, unless there is a `.snyk` file in the Project. For details, see the section [How to override the ignore rules in the database](the-.snyk-file.md#how-to-override-the-ignore-rules-in-the-database) on this page.
* If there is a `.snyk` file in the Project, `snyk test` uses that file as the ignore mechanism, instead of the ignores set from the Web UI.
* When the `.snyk` file is included in an SCM Project, Snyk considers both the database ignores and the `.snyk` ignores.
* Developers can ignore issues by using the .snyk policy file when running `snyk monitor`.
* The `.snyk` file defines certain analysis configuration items such as `language settings:` for the Python version.
  * For SCM scans, for example, GitHub, the Snyk Web UI limits users to setting Python versions at the Organization level.
  * When you include the `.snyk` file in your code repository and the `language settings:` value is set, you gain the advantage of creating Project-level Python settings when you run code repository scans.
  * You may need to re-import the Project if the `.snyk` file was not present at the initial import of the Project into Snyk.

## How to create the `.snyk` file

{% hint style="info" %}
The `.snyk` file should be versioned in the code repository, the same as other applications and build resources.
{% endhint %}

You can create the `.snyk` file in a number of ways:

* **Snyk vulnerability fix pull request (PR)** - When you select the **fix a vulnerability** button on a Git code repository scan, and a Snyk patch is available and an upgrade is not possible, a `.snyk` file is added to the pull request. Creating Snyk patches is supported for npm and Yarn only.
* **Snyk CLI** - When you use the `snyk ignore` command, Snyk creates a `.snyk` file.
* **Manual creation** - If you want to ignore by path, you must edit the `.snyk` file manually. You can create a new `.snyk` file and populate it with the code that follows:\
  &#x20;`# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities`\
  `version: v1.25.0`\
  The `version` is the current version of the snyk-policy package found at [https://www.npmjs.com/package/snyk-policy](https://www.npmjs.com/package/snyk-policy).

## Syntax of the `.snyk` file

The `.snyk` file may have the following top-level keys:

* `language-settings:`
* `ignore:`
* `patch:`

The `language-settings:` value is the Python version you are currently using. See the examples in the section [Setting the language version for Python](the-.snyk-file.md#setting-the-language-version-for-python) on this page.

The `ignore:` is an ignore rule in the form of:

```
ignore:
  snyk-vulnid:
    - path to library using > seperator :
      reason: 'text string'
      expires: 'datetime string'
```

The `patch`: is in the form of:

```
'npm:library:yyyymmdd’ :
  - path to library using > seperator:
    patched: 'datetime string'
  - path to library using > seperator > to > another > path:
    patched: 'datetime string'
```

## How to override the ignore rules in the database

You can use a `.snyk` file to override the ignore rules in the database if **Admin users only** is enabled for the Organization.

When Admin users only is enabled, if there is a `.snyk` file in the Project, the`snyk test` CLI command uses that file as the ignore mechanism, instead of the ignores set in the Web UI. This means that if you have a `.snyk` file in the Project and you are using the `snyk test` command through the CLI, Snyk overrides all settings made in the Snyk Web UI.

You can choose who can ignore an issue or edit the ignore settings for an issue by using the Snyk Web UI or the API.

Follow these steps to set the ignore preferences for use by the Snyk Web UI and API:

1. Log in to your [Snyk account](../../getting-started/quickstart/create-a-snyk-account/).
2. Select **Settings**, then **General**.
3. Select an option as follows:
   * **Admin users only** - only admins can customize the ignore settings.
   * **All users in any environment** - all users can customize the ignore settings.

{% hint style="info" %}
Snyk does not process the ignores from changes added in a newer commit or in the same PR. Only ignores from the same PR are applied.
{% endhint %}

## Monorepos and complex Project considerations

The Snyk CLI expects the `.snyk` file to apply to the manifest being analyzed. In the case of a complex Project or monorepo, there may be many manifests in subfolders, and you may wish to use a centralized ignore policy. The `.snyk` file is expected to be the root of your Project, with your manifest file. If the `.snyk` file is not in the root of your Project, you must specify the path explicitly using the `--policy-path`, for example, in the case of a centralized policy.

If you create a `.snyk` ignore policy using the CLI and Snyk does not successfully ignore the vulnerability, use the option `--policy-path=/path/path/file.`

Your complete statement should be `snyk ignore --id=IssueID [--expiry=expiry] [--reason='reason for ignoring'] [--policy-path=/path/path/file].`

{% hint style="info" %}
If you use the `.snyk` policy file, you avoid having to specify ignores in the web interface, which you can do only after an issue is detected and monitored.
{% endhint %}

{% hint style="warning" %}
For Projects imported using a code repository integration as opposed to using the `snyk monitor` command, the `--policy-path` option is not available. The `.snyk` file  applies only to Projects found on the same path as the `.snyk` file.
{% endhint %}

## Examples for the `.snyk` file

### Create a `.snyk` file

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

### Set the language version for Python

Manually modify the `.snyk` file to set `language-settings:` for the Project to Python 2.7:

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
language-settings: 
  python: "2.7"
```

Manually modify the `.snyk` file to set `language-settings:` for the Project to Python 3.6.2:

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
language-settings: 
  python: "3.6.2"
```

{% hint style="info" %}
When you include the `.snyk` file in your code repository and the `language-settings:` value is set, you gain the advantage of creating Project-level Python settings when you run code repository scans.
{% endhint %}

### Set vulnerability ignore rules

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
  SNYK-JS-BSON-561052:
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

### Set license ignore rules

To ignore the license issue for a package, find the ID for the license in the output of the `snyk test` command.

The license ID is part of the license issue URL, for example, in this URL: [https://snyk.io/vuln/snyk:lic:npm:symbol:MPL-2.0](https://snyk.io/vuln/snyk:lic:npm:symbol:MPL-2.0), the license ID is `snyk:lic:npm:symbol:MPL-2.0`.

### **Ignoring the license with the CLI**

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

## `.snyk`-related CLI commands

The `snyk policy` command displays the `.snyk` policy for a package.

The `snyk ignore` command modifies the `.snyk` policy to ignore a stated issue.

```
snyk ignore --id='vulnerabilityID' --expiry='date-string' --reason='text string'
```

The following example shows using the `snyk ignore` command to generate a rule for ignoring the `SNYK-JS-BSON-561052` vulnerability for all paths that lead to that library on disk.

```
snyk ignore --id='SNYK-JS-BSON-561052' --expiry='2018-04-01' --reason='testing'
```

## More information about the `.snyk` file

For more information, see the following:

[Ignore vulnerabilities using the Snyk CLI](https://docs.snyk.io/snyk-cli/fix-vulnerabilities-from-the-cli/ignore-vulnerabilities-using-snyk-cli)

[Error message: Ignoring via the CLI is not enabled for this organization. Please ignore issues via our website](https://support.snyk.io/hc/en-us/articles/360001569438-Error-message-Ignoring-via-the-CLI-is-not-enabled-for-this-organization-Please-ignore-issues-via-our-website)
