# Getting started with Snyk Open Source

* [ Create a Snyk account](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360017098237-Create-a-Snyk-account/README.md)
* [ Select a Snyk product / tool](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360014959818-Select-a-Snyk-product-tool/README.md)
* [ Getting started with Snyk Open Source](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360014875297-Getting-started-with-Snyk-Open-Source/README.md)
* [ Getting started with Snyk Code](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360016765157-Getting-started-with-Snyk-Code/README.md)
* [ Getting started with Snyk Container](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360014877957-Getting-started-with-Snyk-Container/README.md)
* [ Getting started with Snyk Infrastructure as Code \(IaC\)](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360014938398-Getting-started-with-Snyk-Infrastructure-as-Code-IaC-/README.md)
* [ Getting Started with Snyk License Compliance Management](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015235618-Getting-Started-with-Snyk-License-Compliance-Management/README.md)
* [ Getting started with Snyk intel vulnerability DB access](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015452178-Getting-started-with-Snyk-intel-vulnerability-DB-access/README.md)

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
4. Fill in the account credentials as prompted \(or authenticate with your account in GitHub\), to grant Snyk access permissions for integration.

See [DevOps integrations & languages](https://support.snyk.io/hc/en-us/articles/360011733538-DevOps-integrations-languages) for more details

Add projects to test with Snyk, by choosing repositories for Snyk to test and monitor.

1. Select **Projects** from snyk.io.
2. Select the tool to add the project from \(for example GitHub\): 
3. In **Personal and Organization repositories**, select the repositories to use: 
4. Click **Add selected repositories** to import the selected repositories into your projects. This also:
   * Sets Snyk to run a regular check \(daily by default\) for vulnerabilities.
   * Creates a Webhook, so when you change code, Snyk tests your pull / merge requests, to check that new dependencies do not introduce more vulnerabilities.
5. A progress bar appears: click **View log** to see log results. 
6. Project import completes.

If you encounter any errors during import, see the [Importing projects](https://support.snyk.io/hc/en-us/sections/360000923478-Importing-projects) information.

#### Stage 3: View vulnerabilities

You can now view vulnerability results for imported projects. The **Projects** tab appears by default after import, showing vulnerability information for project you've imported.

1. Click on an imported project to see vulnerability information for that project, including the number of issues found, grouped by severity level:
2. Click on an entry to open the issues view for that entry, including the module, where it was introduced, and the remediation to fix it, plus more details about the vulnerability itself: 

See [View project information](https://support.snyk.io/hc/en-us/articles/360011450838-View-project-information) for more details.

#### Stage 4: Fix vulnerabilities

For JavaScript, Ruby and Java projects, Snyk can remediate your vulnerabilities via fix pull/merge requests:

1. Open the issues view for a project

   Click **Fix this vulnerability** to upgrade \(or patch\) to fix an individual issue, or click **Fix these vulnerabilities** to to fix multiple issues at once.

2. The **Open a Fix PR** screen opens and indicates the vulnerabilities you selected:
3. Check any additional issues you want to fix, or uncheck items to remove them from the fix.
4. Scroll down to the bottom of the screen and click **Open a Fix PR**.
5. Snyk now actions this PR, then a results screen appears: 
6. Optionally, select the **Files changed** tab to see details of the changes made. 

See [Fixing vulnerabilities](https://support.snyk.io/hc/en-us/articles/360011484018-Fixing-vulnerabilities) for more details.

### For more information

See [Snyk Open Source](https://support.snyk.io/hc/en-us/categories/360003049458-Snyk-Open-Source).

