# Getting started with Snyk Open Source

## Getting started with Snyk Open Source

Get started with Snyk Open Source to inspect, find and fix vulnerabilities in your code.

This process describes using the Snyk.io UI and a source code management system.  
You can also use an [IDE tool](https://support.snyk.io/hc/en-us/sections/360001138118-IDE-tools) or a [CI/CD integration](https://support.snyk.io/hc/en-us/sections/360001152577-CI-CD-integrations). See [Integrations](https://support.snyk.io/hc/en-us/categories/360000598398-Integrations) for more details.

**Using the CLI tool**

The [Snyk CLI tool](https://support.snyk.io/hc/en-us/articles/360003812458-Getting-started-with-the-CLI) allows you to get started using the command line - for example, to install on npm:

```text
npm install -g snyk
```

See [Getting started with the CLI](https://support.snyk.io/hc/en-us/articles/360003812458-Getting-started-with-the-CLI) for details.

#### Stage 1: Add source control integration

if you already have an integration set up, you can go to Step 3.

Choose a source code integration, to allow Snyk to work on a project. 

1. Log in to Snyk.io.
2. Select **Integrations &gt; Source control**.
3. Click the source control system \(for example, GitHub\) to integrate with Snyk:

![](../.gitbook/assets/screenshot-2021-06-22-at-10.58.18.png)

1. Fill in the account credentials as prompted \(or authenticate with your account in GitHub\), to grant Snyk access permissions for integration.

See [DevOps integrations & languages](https://support.snyk.io/hc/en-us/articles/360011733538-DevOps-integrations-languages) for more details

Add projects to test with Snyk, by choosing repositories for Snyk to test and monitor.

1. Select **Projects** from snyk.io.
2. Select the tool to add the project from \(for example GitHub\):

![](../.gitbook/assets/open-source-add-projects.png)

1. In **Personal and Organization repositories**, select the repositories to use:

![](../.gitbook/assets/screenshot-2021-04-09-at-17.08.24.png)

1. Click **Add selected repositories** to import the selected repositories into your projects. This also:
   * Sets Snyk to run a regular check \(daily by default\) for vulnerabilities.
   * Creates a Webhook, so when you change code, Snyk tests your pull / merge requests, to check that new dependencies do not introduce more vulnerabilities.
2. A progress bar appears: click **View log** to see log results. 
3. Project import completes.

If you encounter any errors during import, see the [Importing projects](https://support.snyk.io/hc/en-us/sections/360000923478-Importing-projects) information.

#### Stage 3: View vulnerabilities

You can now view vulnerability results for imported projects. The **Projects** tab appears by default after import, showing vulnerability information for project you've imported.

* Click on an imported project to see vulnerability information for that project, including the number of issues found, grouped by severity level:

![](../.gitbook/assets/view-vulns-overview.png)

* Click on an entry to open the issues view for that entry, including the module, where it was introduced, and the remediation to fix it, plus more details about the vulnerability itself:

![](../.gitbook/assets/detailed-vuln-information.png)

See  [View project information](https://support.snyk.io/hc/en-us/articles/360011450838-View-project-information) for more details.

#### Stage 4: Fix vulnerabilities

For JavaScript, Ruby and Java projects, Snyk can remediate your vulnerabilities via fix pull/merge requests:

* Open the issues view for a project

![](../.gitbook/assets/screenshot-2021-04-09-at-17.35.25.png)

Click **Fix this vulnerability** to upgrade \(or patch\) to fix an individual issue, or click **Fix these vulnerabilities** to to fix multiple issues at once.

* The **Open a Fix PR** screen opens and indicates the vulnerabilities you selected:

![](../.gitbook/assets/screenshot-2021-04-09-at-17.40.00.png)

* Check any additional issues you want to fix, or uncheck items to remove them from the fix.
* Scroll down to the bottom of the screen and click **Open a Fix PR**.
* Snyk now actions this PR, then a results screen appears:

![](../.gitbook/assets/screenshot-2021-04-09-at-17.44.26.png)



* Optionally, select the **Files changed** tab to see details of the changes made.

![](../.gitbook/assets/screenshot-2021-04-09-at-17.46.22.png)

See [Fixing vulnerabilities](https://support.snyk.io/hc/en-us/articles/360011484018-Fixing-vulnerabilities) for more details.

### For more information

See [Snyk Open Source](https://support.snyk.io/hc/en-us/categories/360003049458-Snyk-Open-Source).



