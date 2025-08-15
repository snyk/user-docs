# The .snyk file

The `.snyk` file is a capability of Snyk that all users can employ locally or as part of their workflow to control Snyk ignores of issues, to exclude files from scanning, to set the Python version at the Project level, and to specify patches for the CLI and CI/CD plugins.

How the `.snyk` file works varies among Snyk products. When you deploy the `.snyk` file, start by reviewing how the file is created, where it can be used, and what it is used for. For details, see [Use the `.snyk` file with Snyk Open Source](the-.snyk-file.md#use-the-.snyk-file-with-snyk-open-source), [Use the `.snyk` file with Snyk Code](the-.snyk-file.md#use-the-.snyk-file-with-snyk-code) and [Use the `.snyk` file with Snyk IaC](the-.snyk-file.md#use-the-.snyk-file-with-snyk-iac).

You can create the `.snyk` file by using the `snyk ignore` CLI command. This generates the file and an ignore rule. You can also create the file using a text or code editor. The format is YAML. For details, see [How to create the `.snyk` file](the-.snyk-file.md#how-to-create-the-.snyk-file).

Key considerations regarding how the `.snyk` file is used are as follows:

* The CLI automatically uses the `.snyk` file if it is present, along with the ignore rules in the database created in the Snyk Web UI, if the product supports using ignores in the database or Snyk rules.
* The CLI used as part of a build system and the CI/CD plugins use the `.snyk` file during scanning if the file is present.
* If you merge the `.snyk` file with the rest of your code, when you import an SCM to Snyk, the rules in the `.snyk` file are applied on top of the database rules created in the Snyk Web UI.
* If you use the `.snyk` file to specify ignores, you avoid having to specify them in the Snyk Web UI, which you can do only after an issue is detected and monitored. You can use the `.snyk` file to override the ignore rules in the Snyk database. For details, see [How to override the ignore rules in the database](the-.snyk-file.md#how-to-override-the-ignore-rules-in-the-database).
* For more information, see [Exclude files and ignore issues FAQs](../prioritize-issues-for-fixing/ignore-issues/exclude-files-and-ignore-issues-faqs.md).

## How to create the `.snyk` file

You can create the `.snyk` file by using the `snyk ignore` command. For details, see the [Ignore](../../developer-tools/snyk-cli/commands/ignore.md) command CLI help.

If you do not have an existing `.snyk` file, you can create one and populate it with the following code:\
`# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities`\
`version: v1.25.0`

* You must set the `version` to `v1.25.0` as this is the current policy schema version.
* The ignore block or blocks must follow the relevant syntax as shown in the [description of the `ignore` command](../../developer-tools/snyk-cli/commands/ignore.md#usage-and-description) or the examples on this page.

For more information, see [Syntax of the `.snyk` file](the-.snyk-file.md#syntax-of-the-.snyk-file).

{% hint style="info" %}
Snyk also provides the [snyk-policy package](https://www.npmjs.com/package/snyk-policy) to create a policy file, typically named `.snyk`. The version of the package is not the same as the policy schema version to be entered in the `.snyk` file.
{% endhint %}

### Where to create the `.snyk` file

Generally, you must ensure the `.snyk` file is created in the code repository, the same as other applications and build resources.

The `.snyk` file is generally located at the root of your Project. However, for SCM imports, the `.snyk` file must be in the same directory as any file needed for scanning to which it relates, for example, a manifest file. See [Use the `.snyk` file with monorepos and complex Projects](the-.snyk-file.md#use-the-.snyk-file-with-monorepos-and-complex-projects).

### Create a `.snyk` file for a patch

When you select the **Fix a vulnerability** button on a Git repository Open Source scan, and a Snyk patch is available, and an upgrade is not possible, a `.snyk` file is added to the pull request to specify a patch. Creating Snyk patches is supported for npm and Yarn only.

The following example shows how to create a `.snyk` file to generate a patch rule using a vulnerability fix PR:

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.25.0
ignore: {}
# patches apply the minimum changes required to fix a vulnerability
patch:
  'npm:hawk:20160119':
    - tap > codecov.io > request > hawk:
        patched: '2020-01-20T14:26:34.404Z'
```

## Use the `.snyk` file with Snyk Code

You can use the `.snyk` file to specify files or directories in a repository that are to be excluded from the Snyk Code scan that will import files for Snyk Code testing and generate the Code Analysis Project. The `exclude from import` option is supported only in Snyk Code, and only for imports that are performed using the Snyk Web UI and CLI.

For Projects imported using a code repository integration as opposed to using the `snyk monitor` command, the `--policy-path` option is not available. The `.snyk` file applies only to Projects found on the same path as the `.snyk` file.

For details, see [Excluding directories and files from the import process](../../scan-with-snyk/import-project-repository/exclude-directories-and-files-from-project-import.md).

## Use the `.snyk` file with Snyk IaC

For IaC ignore rules, see [IaC ignores using the `.snyk` policy file](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/iac-ignores-using-the-.snyk-policy-file.md). For more information, see [Ignore unmanaged resources](../../scan-with-snyk/snyk-iac/detect-manually-created-resources/ignore-unmanaged-resources.md).

## Use the `.snyk` file with Snyk Open Source

The `.snyk` file in a Project is used to apply ignores and other settings for the `snyk test` and `snyk monitor` commands and for any tests done through the API or Snyk Web UI. The `.snyk` file defines Snyk patches to be applied at build time, to resolve vulnerabilities that cannot be fixed with upgrades, and to apply the `@snyk/protect` [package](https://github.com/snyk/snyk/tree/master/packages/snyk-protect) that replaced the `snyk protect` command. The `.snyk` file defines analysis configuration items such as `language settings:` for the Python (Pip) version.

### How the `.snyk` file works with Open Source Projects

Snyk checks the Snyk database and the `.snyk` file for ignore rules when scanning by means of an SCM integration, the Snyk CLI, and a CI/CD integration.

If there is a `.snyk` file in the Project, the`snyk test` command uses that file as the ignore mechanism, instead of the ignores set from the Web UI.

When the `.snyk` file is included in an SCM Project, Snyk considers both the database ignores and the `.snyk` ignores.

When you include the `.snyk` file in your code repository and the `language-settings:` value is set, you gain the advantage of creating Project-level Python settings when you run code repository scans.

* For SCM scans, for example, GitHub scans, the Snyk Web UI controls the Python version at the Organization level, from the **Organization** > **Settings** > **Snyk Open Source** > **Python** > **Pip Python version** option.
* By including a `.snyk` file in your code repository with the `language settings:` value set to one of the available UI language settings options, you can override the Organization level settings for SCM scans of that repository to use any Python version that is available in the UI options.

{% hint style="info" %}
if the `.snyk` file was not present at the initial import of the Project into Snyk., you must re-import the Project.
{% endhint %}

For more information about Python version support, see [Python version support](../../supported-languages/supported-languages-list/python/#python-version-support).

For more information about using the `.snyk` file with Open Source Projects, see the following:

[Ignore vulnerabilities using the Snyk CLI](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/ignore-vulnerabilities-using-the-snyk-cli.md)

[Error message: Ignoring via the CLI is not enabled for this organization. Please ignore issues via our website](https://support.snyk.io/s/article/Error-message-Ignoring-via-the-CLI-is-not-enabled-for-this-organization-Please-ignore-issues-via-our-website)

### Examples of the .snyk file for Open Source

#### Set the language version for Python

Manually modify the `.snyk` file to set `language-settings:` for the Project to Python 3.7:

```
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.25.0
language-settings: 
  python: "3.7"
```

For more information, see [Setting Python version in Git Projects](../../supported-languages/supported-languages-list/python/scm-integrations-and-python.md#set-the-python-version-in-git-projects).

#### Set vulnerability ignore rules

{% hint style="warning" %}
The `expires` field is optional. If you need a permanent ignore, omit the field.

To ensure that expiration dates are enforced for ignores that are created through the `.snyk` file, you must specify a valid expiration date. The date must be in the Date Time String Javascript format like YYYY-MM-DDThh:mm:ss.fffZ. If the specified expiration date does not adhere to this format, the ignore will be respected and persist indefinitely. For details, see the [--expiry option in the snyk ignore command help](../../developer-tools/snyk-cli/commands/ignore.md#expiry-less-than-expiry-greater-than).
{% endhint %}

Ignore a specific vulnerability for a given path:

<pre><code><strong>ignore:
</strong>  SNYK-JS-BSON-561052:
    - mongodb > mongodb-core > bson:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
</code></pre>

Ignore a vulnerability for all paths:

```
ignore:
  SNYK-JS-BSON-561052:
    - '*':
        reason: None Given
        expires: 2020-04-04T17:33:45.004Z
```

Ignore a specific vulnerability on multiple paths:

<pre><code><strong>ignore:
</strong><strong>  SNYK-JS-DOTPROP-543489:
</strong>    - configstore > dot-prop:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
    - snyk > configstore > dot-prop:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
</code></pre>

#### Set license ignore rules

To ignore the license issue for a package, find the ID for the license in the output of the `snyk test` command.

The license ID is part of the license issue URL, for example, in this URL: [https://snyk.io/vuln/snyk:lic:npm:symbol:MPL-2.0](https://snyk.io/vuln/snyk:lic:npm:symbol:MPL-2.0), the license ID is `snyk:lic:npm:symbol:MPL-2.0`.

## Use the Snyk CLI and the `.snyk` file for Snyk Open Source

### CLI commands to use

The Snyk CLI has commands to create and view a `.snyk` file.

The `snyk policy` command displays the `.snyk` policy for a package.

The `snyk ignore` command modifies the `.snyk` file to ignore a stated issue.

```
snyk ignore --id='vulnerabilityID' --expiry='date-string' --reason='text string'
```

The following example shows using the `snyk ignore` command to generate a rule for ignoring the `SNYK-JS-BSON-561052` vulnerability for all paths that lead to that library on disk.

```
snyk ignore --id='SNYK-JS-BSON-561052' --expiry='2018-04-01' --reason='testing'
```

### Use the `.snyk` file with monorepos and complex Projects

The Snyk CLI expects the `.snyk` file to apply to the manifest being analyzed. In the case of a complex Project or monorepo, there may be many manifests in subfolders, and you may wish to use a centralized ignore policy. The `.snyk` file is expected to be the root of your Project, with your manifest file. If the `.snyk` file is not in the root of your Project, for example, in the case of a centralized policy, you must specify the path explicitly using the `--policy-path` option.

If you create a `.snyk` ignore policy using the CLI and Snyk does not successfully ignore the vulnerability, use the option `--policy-path=/path/path/file.`

Your complete statement should be `snyk ignore --id=IssueID [--expiry=expiry] [--reason='reason for ignoring'] [--policy-path=/path/path/file].`

## How to override the ignore rules in the database

If there is a `.snyk` file in the Project, the`snyk test` CLI command uses that file as the ignore mechanism instead of the ignores set in the Web UI. This means that if you have a `.snyk` file in the Project and you are using the `snyk test` command through the CLI, Snyk overrides all settings made in the Snyk Web UI.

However, when the `.snyk` file is included in an SCM Project, Snyk considers both the database ignores and the `.snyk` ignores.

If **Admin users only** is enabled by using **Settings > General > Ignores,** you can use a `.snyk` file to override the ignore rules in the database. To override the ignore rules set in the Web UI, you must specify that Admin users only can ignore an issue or edit the ignore settings for an issue.

To set these ignore preferences for use by the Snyk Web UI and API:

1. Log in to your Snyk account.
2. Select **Settings** > **General**.
3. Select an option as follows:
   * **Admin users only** - only admins can customize the ignore settings.
   * **All users in any environment** - all users can customize the ignore settings.

## Syntax of the `.snyk` file

The `.snyk` file has the following top-level keys:

* `language-settings:`
* `ignore:`
* `patch:`

The `language-settings:` value is the Python version you are using. See the examples in the section [Set the language version for Python](the-.snyk-file.md#set-the-language-version-for-python) on this page.

The `ignore:` is an ignore rule in the form of:

```
ignore:
  snyk-vulnid:
    - path to library using > separator :
        reason: 'text string'
        expires: 'YYYY-MM-DDThh:mm:ss.fffZ'
```

{% hint style="info" %}
Note that a double indent is required for the `reason` and `expires` fields.
{% endhint %}

The`expires` field is optional. If you need a permanent ignore, omit this field as shown in the following example:

```
ignore:
  snyk-vulnid:
    - path to library using > separator :
        reason: 'text string'
```

The `patch`: is in the form of:

```
'npm:library:yyyymmdd’ :
  - path to library using > separator:
    patched: 'datetime string'
  - path to library using > separator > to > another > path:
    patched: 'YYYY-MM-DDThh:mm:ss.fffZ'
```
