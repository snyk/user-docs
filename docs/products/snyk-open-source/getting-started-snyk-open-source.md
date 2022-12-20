# Getting started with Snyk Open Source

Use Snyk Open Source to scan and fix vulnerabilities in your application's Open Source libraries, for a [supported language and package manager](language-and-package-manager-support/), such as Java.

{% hint style="info" %}
This process describes getting started using the [Snyk Web UI](../../snyk-web-ui/) with a Git-based source repository.

* You can also use an [IDE tool](https://docs.snyk.io/integrations/ide-tools) or a [CI/CD integration](https://docs.snyk.io/integrations/ci-cd-integrations). See [Integrations](https://docs.snyk.io/integrations).
* You can also [use Snyk Open Source from the CLI](use-snyk-open-source-from-the-cli/).
{% endhint %}

### **Prerequisites**

Ensure you have:

* [Created a Snyk account](../../getting-started/quickstart/create-a-snyk-account.md).
* [Set up integration](../../getting-started/quickstart/set-up-an-integration.md) with your code repository on a supported system such as GitHub.
* [Imported a Snyk Project for scanning](../../getting-started/quickstart/import-a-project.md)

See the [Getting started](../../getting-started/) section for more details.

### View vulnerabilities

You can view vulnerability results for imported Projects. The **Projects** tab appears by default after import, showing vulnerability information for Snyk Projects you've imported, grouped into Targets.

You can expand a Target to see vulnerability information for Projects, including the number of issues found, grouped by severity level:

<figure><img src="../../.gitbook/assets/Getting started with open source.png" alt="Screenshot of a Snyk Project listing"><figcaption><p>.</p></figcaption></figure>

Click an entry to open the issues view for that entry, including the module where it was introduced, how to fix it, plus more details about the vulnerability itself.

![](../../.gitbook/assets/project-details.png)

See [view-project-information.md](../../snyk-web-ui/introduction-to-snyk-projects/view-project-information.md "mention") for more details.

### Fix vulnerabilities

For some languages, Snyk can fix vulnerabilities using fix pull/merge requests.

{% hint style="info" %}
See [what-languages-do-we-support-fix-pull-requests-or-merge-requests.md](../../features/fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/what-languages-do-we-support-fix-pull-requests-or-merge-requests.md "mention")
{% endhint %}

Navigate to the **Issues** view for a project:

![](../../.gitbook/assets/Issues-view.png)

To fix vulnerabilities:

1. Click **Fix this vulnerability** to raise a fix PR for that issue (or click **Fix these vulnerabilities** to fix multiple issues).
2. The **Open a Fix PR** screen opens and indicates the selected vulnerabilities.
3. Check or uncheck the issues you want to fix or remove from this fix.
4. Scroll to the bottom of the screen and click **Open a Fix PR**.
5. Snyk acts on the PR and displays a results screen.
6. Optionally, select the **Files changed** tab to see details of the changes made.

![](../../.gitbook/assets/screenshot\_2021-04-09\_at\_17.46.22.png)

See [fix-your-vulnerabilities.md](../../features/fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/fix-your-vulnerabilities.md "mention") for more details.
