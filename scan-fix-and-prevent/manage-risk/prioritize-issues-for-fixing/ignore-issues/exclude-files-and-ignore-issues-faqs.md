---
description: Frequently asked questions about excluding files and ignoring issues in Snyk
nav_context: agnostic
---

# Exclude files and ignore issues FAQs

There are many considerations in determining how excluding files and ignoring issues will work, depending on several factors:

* How the Project was imported: through an SCM integration, or through the CLI or an IDE
* The scanning method being used: Open Source, Code, Container, or IaC
* How the test is being done, in the UI, or through the CLI or an IDE
* How the exclude or ignore was set: in a policy, through the UI or the API, or in the `.snyk` file

This document collects questions the support team receives frequenty and provides the answers.

## Questions related to scanning methods

### How do I ignore issues and vulnerabilities in Code (SAST) scans?

You cannot use the `.snyk` file to ignore a specific Snyk Code finding. For Snyk Code, the `.snyk` file excludes files and directories from the scan; it does not suppress individual findings.

To ignore a Snyk Code finding, use one of the following supported methods:

* Ignore the finding on its issue card in the Snyk Web UI. This is the method to use for a single finding. For details, see [Ignore issues in the Snyk Web UI](./#ignore-issues-in-the-snyk-web-ui).
* Create the ignore with the `snyk ignore create` command. This command is an Early Access feature of the Ignore Approval Workflow. For details, visit [Ignore create](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/commands/ignore-create).
* Ignore findings in bulk with a Snyk Code Security policy at the Group level, matching on CWE, Snyk Code rule ID, or severity. This method is available to Enterprise customers. For details, see [Manage ignores at the Group level through Snyk Code Security policies](consistent-ignores-for-snyk-code/#manage-ignores-at-the-group-level-through-snyk-code-security-policies).

With [Consistent Ignores for Snyk Code](consistent-ignores-for-snyk-code/), these ignores also apply when you run `snyk code test` in the CLI, in your IDE, and in pull request checks. To display the ignored findings, run `snyk code test --include-ignores`.

The `snyk-to-html` tool displays all issues for Code scans, whether the issues are ignored or not.

### How do I avoid scanning certain files for Open Source scans?

* Use the --`exclude` option when scanning with the CLI to omit scanning directories or files but not paths. This option excludes all directories or all files with specified names. For details, see [the --exclude option](https://docs.snyk.io/developer-tools/snyk-cli/commands/test#exclude-less-than-name-greater-than-less-than-name-greater-than-...greater-than) in the CLI `test` command help.
*   If you import a Project through an SCM integration, add the exclusions, folders only, to the bottom of the import window; see [Stage 2: Import Project](https://docs.snyk.io/developer-tools/integrations/scm-integrations/deployment-recommendations#stage-2-import-projects) in Git repositories deployment recommendations.\\

    <figure><img src="https://lh7-us.googleusercontent.com/stHVnzk1ZuP6oUm0zAImt0zROcajuZMm5iB4qX7vTbHkjPWklSgD9NxUdZ6UGgT1kV-dBjrcLyOp0SP1CqFzbNuq9S7qgl4cOD6T9UwuWlEk5SWVHUiHRlO-KfAyq_UppnGNvE67p7ZsSwuWok0_2RM" alt="Exclude folders"><figcaption><p>Exclude folders</p></figcaption></figure>
* You cannot use an `exclude` block in a `.snyk` file for Open Source scans except for unmanaged scans. For details, see [Ignore files or folders using glob expression - Snyk Code and `unmanaged`only](https://docs.snyk.io/developer-tools/snyk-cli/commands/ignore#ignore-files-or-folders-using-glob-expression-snyk-code-and-unmanaged-only).

### How do I avoid scanning certain files for Code scans?

* Use an exclude in a `.snyk` file to omit all scanning of certain files or folders from a Snyk Code scan. For details, see the [`--file-path` option for the `snyk ignore` command](https://docs.snyk.io/developer-tools/snyk-cli/commands/ignore#file-path-less-than-path_to_resource-greater-than), [Ignore files or folders using glob expression - Snyk Code and `unmanaged`only](https://docs.snyk.io/developer-tools/snyk-cli/commands/ignore#ignore-files-or-folders-using-glob-expression-snyk-code-and-unmanaged-only), and [Exclude directories and files from Snyk Code CLI tests](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests).
* When you import a repository to test using Snyk Code, use an `exclude:` statement in the `.snyk` file to omit certain directories and files from the import. For details see For details see [Exclude directories and files from Snyk Code CLI tests](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests).
* A `.snyk` file with file or folder exclusions and contained in the root directory of your repository or SCM will exclude those files and folders from being scanned when you import using an SCM.
* The CLI `--exclude` option used with `snyk test` and `snyk monitor` does not apply for Code scans.
* The Exclude Folders option in the import windows in the Web UI does not apply for Code scans.
* The `.snyk` file does not apply for excluding files and directories from IDE scanning of Code.
* For Code and Container scans only., you can use exclusion globs in API import, including an import with the `snyk-api-import` tool. This exclusion works the same way as an SCM integration exclusion.

### Can I ignore a Snyk Code finding by rule ID in the `.snyk` file?

No. The `ignore` block in the `.snyk` file does not accept a Snyk Code rule ID, and there is no `.snyk` syntax for ignoring a Code finding by rule ID and file path. The `ignore` block applies to Snyk Open Source, Snyk Container, and Snyk IaC issues only.

To ignore Snyk Code findings by rule ID, use a Snyk Code Security policy at the Group level. You can find the rule ID in the SARIF output of `snyk code test --sarif`. For details, see [Manage ignores at the Group level through Snyk Code Security policies](consistent-ignores-for-snyk-code/#manage-ignores-at-the-group-level-through-snyk-code-security-policies).

### How do I avoid scanning certain files for Container scans?

See the last bullet in the previous section and the [`--exclude-app-vulns`](https://docs.snyk.io/developer-tools/snyk-cli/commands/container-test#exclude-app-vulns), [`--exclude-base-image-vulns`](https://docs.snyk.io/developer-tools/snyk-cli/commands/container-test#exclude-base-image-vulns), and [`--exclude-node-modules`](https://docs.snyk.io/developer-tools/snyk-cli/commands/container-test#exclude-node-modules)options in the `snyk container test` help.

### How do I avoid scanning certain files for IaC scans?

See [IaC exclusions using the command line](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/iac-exclusions-using-the-command-line).

## Question related to ways excludes and ignores are set

Why do ignores in the `.snyk` file in the root directory of my monorepo apply to all Projects when I scan with the CLI but not when I import using my SCM?

For SCM scanning, the `.snyk` file must be present in each relevant subdirectory. See [Monorepos and complex Project considerations](../../policies/the-.snyk-file.md#use-the-.snyk-file-with-monorepos-and-complex-projects) with the `.snyk` file.
