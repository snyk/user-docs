# Getting started with Snyk Open Source

Use Snyk Open Source to scan and fix vulnerabilities in your application's Open Source libraries, for a [supported language and package manager](../../scan-applications/supported-languages-and-frameworks/), such as Java.

{% hint style="info" %}
This process describes getting started using the [Snyk Web UI](../../getting-started/quickstart/create-a-snyk-account/logging-in-to-an-existing-account.md) with a Git-based source repository.

* You can also use an [IDE tool](../../integrations/ide-tools/) or a [CI/CD integration](../../integrations/snyk-ci-cd-integrations/). See [Integrations](../../integrations/).
* You can also [use Snyk Open Source from the CLI](use-snyk-open-source-from-the-cli/).
{% endhint %}

### **Prerequisites**

Ensure you have:

* [Created a Snyk account](../../getting-started/quickstart/create-a-snyk-account/).
* [Set up integration](../../getting-started/quickstart/set-up-an-integration.md) with your code repository on a supported system such as GitHub.
* [Imported a Snyk Project for scanning](../../getting-started/quickstart/import-a-project.md)

See the [Getting started](../../getting-started/) section for more details.

### View vulnerabilities

You can view vulnerability results for imported Projects. The **Projects** tab appears by default after import, showing vulnerability information for Snyk Projects you've imported, grouped into Targets.

You can expand a Target to see vulnerability information for Projects, including the number of issues found, grouped by severity level:

<figure><img src="../../.gitbook/assets/Getting started with open source.png" alt="List of projects in Snyk Web UI."><figcaption><p>Projects overview</p></figcaption></figure>

Click an entry to open the issues view for that entry, including the module where it was introduced, how to fix it, plus more details about the vulnerability itself.

<figure><img src="../../.gitbook/assets/project-details.png" alt="Open source project overview in Snyk Web UI."><figcaption><p>Open source project overview</p></figcaption></figure>

See [View Project information](../../manage-issues/snyk-projects/view-project-information.md) for more details.

### Fix vulnerabilities

For some languages, Snyk can fix vulnerabilities using fix pull/merge requests (see [What languages do we support](snyk-open-source-supported-languages-and-package-managers/), [Fix Pull Request or Merge Requests](open-source-basics/)).

Navigate to the **Issues** view for a project:

<figure><img src="../../.gitbook/assets/Issues-view.png" alt="The Issues tab in an open source project in Snyk Web UI."><figcaption><p>Issues tab in open source project</p></figcaption></figure>

To fix vulnerabilities:

1. Click **Fix this vulnerability** to raise a fix PR for that issue (or click **Fix these vulnerabilities** to fix multiple issues).
2. The **Open a Fix PR** screen opens and indicates the selected vulnerabilities.
3. Check or uncheck the issues you want to fix or remove from this fix.
4. Scroll to the bottom of the screen and click **Open a Fix PR**.
5. Snyk acts on the PR and displays a results screen.
6. Optionally, select the **Files changed** tab to see details of the changes made.

<figure><img src="../../.gitbook/assets/screenshot_2021-04-09_at_17.46.22.png" alt="The files changed tab in GitHub after triggering Fix PR for an open source project."><figcaption><p>Files changed tab in GitHub after triggering Fix PR for an open source project</p></figcaption></figure>

See [Fix your vulnerabilities](starting-to-fix-vulnerabilities/fix-your-vulnerabilities.md) for more details.
