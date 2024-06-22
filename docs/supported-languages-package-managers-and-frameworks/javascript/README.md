# JavaScript

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## Getting started with Snyk for Javascript across environments

### Snyk CLI&#x20;

#### Prerequisites

* [Create a Snyk account](../../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md)
* [Install Snyk CLI and authenticate your machine](../../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
* [Set the default Organization for all Snyk tests](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/set-the-snyk-organization-for-the-cli-tests.md) (code analysis)

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Explore test results in a JSON or SARIF format in the terminal ](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
* [Exporting the test results to a JSON or SARIF file](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)

#### Open source and licensing

The following table lists the steps to start scanning your dependencies. It covers basic commands, such as `snyk test` and `snyk monitor`. To check the full list, see [CLI commands and options summary](../../snyk-cli/cli-commands-and-options-summary.md).

{% hint style="info" %}
To scan your dependencies, ensure you install the relevant package manager and that your Project contains the supported manifest files.
{% endhint %}

| Package manager | Getting started                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                                                    |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| npm             | <ol><li>Install NPM.</li><li>Make sure you are in a directory with NPM Project files, that is, <code>package.json</code> and <code>package-lock.json</code>.</li><li>Run <code>npm install</code>.</li><li>Run <a href="../../snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../snyk-cli/commands/test.md#options-for-npm-projects">snyk test</a> and <a href="../../snyk-cli/commands/monitor.md#options-for-npm-projects">snyk monitor</a>.<br></li></ol> | <p>Snyk analyzes your <code>package.json</code> and <code>package-lock.json</code> files to build a fully structured dependency tree. </p><p></p><p>If the <code>package-lock.json</code> is missing, Snyk analyzes your <code>node_modules</code> folder.</p> |
| Yarn            | <ol><li>Install Yarn.</li><li>Make sure you are in a directory with Yarn Project files, that is, <code>package.json</code> and <code>yarn.lock</code>.</li><li>Run <code>yarn</code></li><li>Run <a href="../../snyk-cli/cli-commands-and-options-summary.md">Snyk commands</a>.</li><li>(Optional) Run command options for <a href="../../snyk-cli/commands/test.md#options-for-yarn-projects">snyk test</a> and <a href="../../snyk-cli/commands/monitor.md#options-for-yarn-projects">snyk monitor</a>.</li></ol>                 | <p>Snyk analyzes your <code>package.json</code> and <code>yarn.lock</code> files to build a fully structured dependency tree. </p><p></p><p>If the <code>yarn.lock</code> is missing, Snyk analyzes your <code>node_modules</code> folder.</p>                 |

## Monorepo Projects

Snyk only scans one manifest file at a time. To scan all manifest files, you can use an argument in the CLI (see [Does the Snyk CLI support monorepos or multiple manifest files?](https://support.snyk.io/hc/en-us/articles/360000910577-Does-the-Snyk-CLI-support-monorepos-or-multiple-manifest-files-)).

**npm workspaces**

NPM v7 introduced the concept of **workspaces**. See [lockfile version and Snyk feature availability matrix](./#npm).&#x20;

**Example**

To detect and scan all workspaces in your npm project, use the `--all-projects` Snyk CLI parameter, as follows:

```javascript
snyk test --all-projects
```

### **Yarn workspaces**

{% hint style="info" %}
`nohoist`is **not** supported by Yarn Workspaces.
{% endhint %}

For Yarn Workspaces, you can use the following flags:

* `--all-projects` : Test and monitor your packages with other Projects or `--yarn-workspaces` , scanning only Yarn Workspaces Projects. The root lock file is referenced when scanning all the packages.&#x20;
* `--detection-depth` : Find sub-folders that are not auto-discovered by default.
* `--strict-out-of-sync=false` : Relax strict synchronization requirements for packages in a Yarn workspace. When set to `false` , you can run Snyk tests with unsynchronized `package.json` and the `package-lock.json` files without throwing errors. Using different dependency versions can introduce potential risks, such as compatibility issues or security vulnerabilities.
* `--policy-path` : Specify the path to a policy used by Snyk during testing.

{% hint style="info" %}
For Yarn Workspaces, only the `package.json` file is updated for Snyk Fix PRs. The `yarn.lock` file is not updated.
{% endhint %}

### **Examples**&#x20;

Scan the packages that belong to any discovered workspaces in this directory, five deep sub-directories, and any other Projects detected.

```javascript
snyk test --all-projects --strict-out-of-sync=false --detection-depth=6 
```



Scan only the Yarn Workspace packages that belong to any discovered workspaces in this directory and five deep sub-directories.

```javascript
snyk test --yarn-workspaces --strict-out-of-sync=false --detection-depth=6
```



Use a common `.snyk` policy file, if you maintain ignores and patches in one place to be applied for all detected workspaces by using the policy path (see [The .snyk file](../../manage-risk/policies/the-.snyk-file.md)).

```javascript
snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk
```

### Snyk Web UI (Git repository integration)

You can import JavaScript repositories from any Git services (Source Control Managers) Snyk supports (see [Git repositories](../../scm-ide-and-ci-cd-workflow-and-integrations/git-repositories-scms-integrations-with-snyk/)). After the import, Snyk analyzes your Projects based on their supported manifest files.

Navigate to the [How Snyk works for open source and licensing](../technical-specifications.md#how-snyk-works-for-open-source-and-licensing) page for more details.

#### Import Project

To import Projects from a Git repository integration:

1. Open Snyk Web UI and go to your Group and Organization.
2. Go to **Projects**.&#x20;
3. Click **Add Projects**, select the import source, and choose the repository\
   \
   If you have an integrated Git repository (GitHub) it shows up as an option to choose from.

Navigate to the [Import a Project](../../getting-started/quickstart/import-a-project.md) page for more details.

#### Configure language settings for open source&#x20;

Configure language settings for your open source and licensing at the Organization level. The configuration settings apply to all Projects in that Organization.

1. Open Snyk Web UI and go to **Settings >** **Languages** section.
2. Under **Languages**, go to **JavaScript** and select **Edit settings.**
3. Configure the settings based on your package manager, **npm** or **Yarn**.

* [ ] **Scan and fix dev dependencies**: If this is selected, Snyk reads the `devDependencies` property on the `package.json` and reports and fixes any vulnerabilities accordingly.
* [ ] **Require package.json and package-lock.json/yarn.lock files to be in sync**: When this is selected, if the `package.json` and `package-lock.json`/`yarn.lock`files are out-of-sync, Snyk fails the import.
* [ ] **Exclude package-lock.json from being generated when fixing vulnerabilities**: If you are using private mirrors or registries, a Snyk-generated lockfile might not be appropriate for you because Snyk uses the npm registry to update the lockfile. This setting allows you to opt out of getting lockfiles generated for you in Snyk fix pull requests and merge requests.

4. **Update Settings** to save changes.

<figure><img src="../../.gitbook/assets/language_settings.png" alt="JavaScript language settings for open source and licensing."><figcaption><p>JavaScript language settings for open source and licensing</p></figcaption></figure>

#### Workspaces

{% hint style="info" %}
Yarn and npm workspaces are not explicitly supported in Snyk git repository integration scans.
{% endhint %}

Root-level `package.json`manifest files with adjacent lockfiles will be scanned as normal.&#x20;

For nested manifest files with no lockfiles Snyk will approximate what the tree will look like at build time without using the root lockfile.

#### Fix PRs and npm save-prefix

When creating a fix for vulnerabilities using npm v7+ Projects, Snyk will use the default npm `save-prefix` rather than inferring it from your Project.

This means if you have dependencies using a range format other than the caret range (`^`), you may see additional changes to the `version` fields in the `package-lock.json` file.&#x20;

These changes should not affect day-to-day functionality, as the ranges will be read from the `package.json`.

#### Fix PRs for Yarn zero-installs users

In Yarn v2, the [zero-installs](https://yarnpkg.com/features/zero-installs) feature was released, which allowed Yarn developers to work on a Project without having to run `yarn` to install dependencies on their machine.&#x20;

Zero-installs achieved this by installing all the dependencies of a Project inside of the `.yarn/cache` directory and asking users to commit this to their version control system , allowing the next developer to pull any new dependencies directly from the repo.

{% hint style="info" %}
If you use the **zero-installs** feature, Snyk Fix PRs do not update the **.yarn/cache** directory. You must run `yarn` to update this directory.
{% endhint %}

#### What's next?

* [Open a Fix PR](./#open-a-fix-pr)&#x20;
* [Configure PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

### Snyk integrations&#x20;

For integrated development environments, see [Use Snyk in your IDE](../../scm-ide-and-ci-cd-workflow-and-integrations/use-snyk-in-your-ide/).

If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../../scm-ide-and-ci-cd-workflow-and-integrations/snyk-ci-cd-integrations/) and [Snyk API](../../snyk-api/)).

## Best practices

To apply best practices for Javascript environments, see [Best practices for Javascript](best-practices-for-javascript-and-node.js.md).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
