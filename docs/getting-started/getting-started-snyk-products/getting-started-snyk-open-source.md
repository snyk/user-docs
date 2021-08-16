# Getting started with Snyk Open Source

Get started with Snyk Open Source to inspect, find and fix vulnerabilities in your code.

This process describes using the Snyk.io UI and a source code management system.  
You can also use an [IDE tool](https://support.snyk.io/hc/en-us/sections/360001138118-IDE-tools) or a [CI/CD integration](https://support.snyk.io/hc/en-us/sections/360001152577-CI-CD-integrations). See [Integrations](https://docs.snyk.io/integrations) for more details.

**Using the CLI tool**

The [Snyk CLI tool](https://docs.snyk.io/snyk-cli/guides-for-our-cli/getting-started-with-the-cli) allows you to get started using the command line - for example, to install on npm:

```text
npm install -g snyk
```

See [Getting started with the CLI](https://docs.snyk.io/snyk-cli/guides-for-our-cli/getting-started-with-the-cli) for details.

#### Stage 1: Add source control integration

if you already have an integration set up, you can go to Step 3.

Choose a source code integration, to allow Snyk to work on a project. 

1. Log in to Snyk.io.
2. Select **Integrations &gt; Source control**.
3. Click the source control system \(for example, GitHub\) to integrate with Snyk: ![Screenshot\_2021-06-22\_at\_10.58.18.png](https://support.snyk.io/hc/article_attachments/4402931344657/Screenshot_2021-06-22_at_10.58.18.png)
4. Fill in the account credentials as prompted \(or authenticate with your account in GitHub\), to grant Snyk access permissions for integration.

See [DevOps integrations & languages](https://support.snyk.io/hc/en-us/articles/360011733538-DevOps-integrations-languages) for more details

Add projects to test with Snyk, by choosing repositories for Snyk to test and monitor.

1. Select **Projects** from snyk.io.
2. Select the tool to add the project from \(for example GitHub\): ![Open-Source-Add-Projects.png](https://support.snyk.io/hc/article_attachments/360012555458/Open-Source-Add-Projects.png)
3. In **Personal and Organization repositories**, select the repositories to use: ![Screenshot\_2021-04-09\_at\_17.08.24.png](https://support.snyk.io/hc/article_attachments/360018814357/Screenshot_2021-04-09_at_17.08.24.png)
4. Click **Add selected repositories** to import the selected repositories into your projects. This also:
   * Sets Snyk to run a regular check \(daily by default\) for vulnerabilities.
   * Creates a Webhook, so when you change code, Snyk tests your pull / merge requests, to check that new dependencies do not introduce more vulnerabilities.
5. A progress bar appears: click **View log** to see log results. 
6. Project import completes.

If you encounter any errors during import, see the [Importing projects](https://support.snyk.io/hc/en-us/sections/360000923478-Importing-projects) information.

#### Stage 3: View vulnerabilities

You can now view vulnerability results for imported projects. The **Projects** tab appears by default after import, showing vulnerability information for project you've imported.

1. Click on an imported project to see vulnerability information for that project, including the number of issues found, grouped by severity level:

   ![View\_vulns\_\_overview.png](https://support.snyk.io/hc/article_attachments/360012480237/View_vulns__overview.png)

2. Click on an entry to open the issues view for that entry, including the module, where it was introduced, and the remediation to fix it, plus more details about the vulnerability itself: ![Detailed\_vuln\_information.png](https://support.snyk.io/hc/article_attachments/360012555638/Detailed_vuln_information.png)

See  [View project information](https://docs.snyk.io/getting-started/introduction-to-snyk-projects/view-project-information) for more details.

#### Stage 4: Fix vulnerabilities

For JavaScript, Ruby and Java projects, Snyk can remediate your vulnerabilities via fix pull/merge requests:

1. Open the issues view for a project

   ![Screenshot\_2021-04-09\_at\_17.35.25.png](https://support.snyk.io/hc/article_attachments/360018815217/Screenshot_2021-04-09_at_17.35.25.png)

   Click **Fix this vulnerability** to upgrade \(or patch\) to fix an individual issue, or click **Fix these vulnerabilities** to to fix multiple issues at once.

2. The **Open a Fix PR** screen opens and indicates the vulnerabilities you selected: ![Screenshot\_2021-04-09\_at\_17.40.00.png](https://support.snyk.io/hc/article_attachments/360018815237/Screenshot_2021-04-09_at_17.40.00.png)
3. Check any additional issues you want to fix, or uncheck items to remove them from the fix.
4. Scroll down to the bottom of the screen and click **Open a Fix PR**.
5. Snyk now actions this PR, then a results screen appears: ![Screenshot\_2021-04-09\_at\_17.44.26.png](https://support.snyk.io/hc/article_attachments/360018815277/Screenshot_2021-04-09_at_17.44.26.png)
6. Optionally, select the **Files changed** tab to see details of the changes made. ![Screenshot\_2021-04-09\_at\_17.46.22.png](https://support.snyk.io/hc/article_attachments/360018815297/Screenshot_2021-04-09_at_17.46.22.png)

See [Fixing vulnerabilities](https://docs.snyk.io/snyk-open-source/open-source-basics/fixing-vulnerabilities) for more details.

### For more information

See [Snyk Open Source](https://docs.snyk.io/snyk-open-source).

