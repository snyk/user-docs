# Exclude files and ignore issues FAQs

There are many considerations in determining how excluding files and ignoring issues will work, depending on several factors:

* How the Project was imported: through an SCM integration, or through the CLI or an IDE
* The scanning method being used: Open Source, Code, Container, or IaC
* How the test is being done, in the UI, or through the CLI or an IDE
* How the exclude or ignore was set: in a policy, through the UI or the API, or in the `.snyk` file

This document collects questions the support team receives frequenty and provides the answers.

## Questions related to scanning methods

### How do I ignore issues and vulnerabilities in Code (SAST) scans?

* To ignore a code vulnerability, import the Project into the Snyk UI, and use the ignore button.&#x20;
* You cannot use the `.snyk` file to ignore issues in Code scans.
* The `snyk-to-html` tool will display all issues for Code scans whether the issues are ignored or not.

### How do I avoid scanning certain files for Open Source scans?

* Use the --`exclude` option when scanning with the CLI to omit scanning directories or files but not paths. This option excludes all directories or all files with specified names. For details, see [the --exclude option](../../../cli-ide-and-ci-cd-integrations/snyk-cli/commands/test.md#exclude-less-than-name-greater-than-less-than-name-greater-than-...greater-than) in the CLI `test` command help.
*   If you import a Project through an SCM integration, add the exclusions, folders only, to the bottom of the import window; see [Stage 2: Import Project](../../../developer-tools/scms/deployment-recommendations.md#stage-2-import-projects) in Git repositories deployment recommendations.\


    <figure><img src="https://lh7-us.googleusercontent.com/stHVnzk1ZuP6oUm0zAImt0zROcajuZMm5iB4qX7vTbHkjPWklSgD9NxUdZ6UGgT1kV-dBjrcLyOp0SP1CqFzbNuq9S7qgl4cOD6T9UwuWlEk5SWVHUiHRlO-KfAyq_UppnGNvE67p7ZsSwuWok0_2RM" alt="Exclude folders"><figcaption><p>Exclude folders</p></figcaption></figure>
* You cannot use an `exclude` block in a `.snyk` file for Open Source scans except for unmanaged scans. For details, see [Ignore files or folders using glob expression - Snyk Code and `unmanaged`only](../../../cli-ide-and-ci-cd-integrations/snyk-cli/commands/ignore.md#ignore-files-or-folders-using-glob-expression-snyk-code-and-unmanaged-only).

### How do I avoid scanning certain files for Code scans?

* Use an exclude in a `.snyk` file to omit all scanning of certain files or folders from a Snyk Code scan. For details, see the [`--file-path` option for the `snyk ignore` command](../../../cli-ide-and-ci-cd-integrations/snyk-cli/commands/ignore.md#file-path-less-than-path_to_resource-greater-than), [Ignore files or folders using glob expression - Snyk Code and `unmanaged`only](../../../cli-ide-and-ci-cd-integrations/snyk-cli/commands/ignore.md#ignore-files-or-folders-using-glob-expression-snyk-code-and-unmanaged-only), and [Exclude directories and files from Snyk Code CLI tests](../../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md).
* When you import a repository to test using Snyk Code, use an `exclude:` statement in the `.snyk` file to omit certain directories and files from the import. For details see For details see [Exclude directories and files from Snyk Code CLI tests](../../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md).
* A `.snyk` file with file or folder exclusions and contained in the root directory of your repository or SCM will exclude those files and folders from being scanned when you import using an SCM.
* The CLI `--exclude` option used with `snyk test` and `snyk monitor` does not apply for Code scans.
* The Exclude Folders option in the import windows in the Web UI does not apply for Code scans.
* The `.snyk` file does not apply for excluding files and directories from IDE scanning of Code.
* For Code and Container scans only., you can use exclusion globs in API import, including an import with the `snyk-api-import` tool. This exclusion works the same way as an SCM integration exclusion.

### How do I avoid scanning certain files for Container scans?

See the last bullet in the previous section and the [`--exclude-app-vulns`](../../../cli-ide-and-ci-cd-integrations/snyk-cli/commands/container-test.md#exclude-app-vulns), [`--exclude-base-image-vulns`](../../../cli-ide-and-ci-cd-integrations/snyk-cli/commands/container-test.md#exclude-base-image-vulns), and [`--exclude-node-modules`](../../../cli-ide-and-ci-cd-integrations/snyk-cli/commands/container-test.md#exclude-node-modules)options in the `snyk container test` help.

### How do I avoid scanning certain files for IaC scans?

See [IaC exclusions using the command line](../../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/iac-exclusions-using-the-command-line.md).

## Question related to ways excludes and ignores are set

Why do ignores in the `.snyk` file in the root directory of my monorepo apply to all Projects when I scan with the CLI but not when I import using my SCM?

For SCM scanning, the `.snyk` file must be present in each relevant subdirectory. See [Monorepos and complex Project considerations](../../policies/the-.snyk-file.md#use-the-.snyk-file-with-monorepos-and-complex-projects) with the `.snyk` file.
