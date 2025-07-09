# Fix your vulnerabilities

Snyk helps you to fix vulnerabilities by upgrading the direct dependencies to a vulnerability-free version or by patching the vulnerability. After Snyk scans your Projects, the scan results allow you to resolve issues in your code with the help of clear suggestions and explanations.

Using Snyk Open Source, you can do the following:

* [View scan results on the Snyk Web UI](fix-your-vulnerabilities.md#view-scan-results-on-the-snyk-web-ui)
* [View scan results using Snyk CLI](fix-your-vulnerabilities.md#fixing-vulnerabilities-based-on-scan-results-using-snyk-cli)
* [Apply fixes](fix-your-vulnerabilities.md#apply-fixes)

## View scan results on the Snyk Web UI

From Snyk Open Source, for each tab, both upgrade and patch, in the fix advice area of your Project details, results are displayed as follows:

* The total number of packages that can be fixed is displayed in the tab title.
* Groups of vulnerabilities by package are shown, labeled by the upgrade or fix that is recommended.
* Packages are listed and can be expanded to show the full list of vulnerabilities affecting the package.
* All the vulnerabilities found in your dependencies are displayed further down the page, together with contextual information that can help you prioritize the issues and start fixing them if required.

<figure><img src="../../../.gitbook/assets/projects_issues_scan_results.png" alt=""><figcaption><p>Example of project issues</p></figcaption></figure>

## View Fix Advice

The Fix Advice area appears on the Project details page. Snyk offers you one of these solutions:

* An upgrade to the original package.
* Pinning a package, installing a package as a top-level dependency; that is, a specific version of an indirect dependency. This avoids having a direct dependency pull in a vulnerable version.
* A Snyk precision patch to fix the issue. If an upgrade to fix any of the vulnerabilities in the package is not currently available, Snyk offers patches to fix the issues

The summary area groups advice by package, and is displayed based on the best available fix. Advice in these summary lists includes these details for each package:

* All vulnerability names and severity details affecting that package
* The recommended fix, a link to the recommended fix for this package and its listed vulnerabilities: either the specific version to which to upgrade or the name of the patch.

<figure><img src="../../../.gitbook/assets/project_details_fixes.png" alt=""><figcaption><p>Upgradable issues tab</p></figcaption></figure>

You can also find additional advice and details further down on the Project details page:

* From the **Issues**, tab, a full description per vulnerability
* From the **Dependencies** tab, the entire tree of your Project dependencies, enabling you to clearly visualize affected paths

## Fixing vulnerabilities based on scan results using Snyk CLI

For information about fixing vulnerabilities, see [Fix vulnerabilities using the Snyk CLI](../../../cli-ide-and-ci-cd-integrations/snyk-cli/scan-and-maintain-projects-using-the-cli/fix-vulnerabilities-using-the-snyk-cli.md).

## Apply fixes

To apply fixes, you can:

* Click **Fix this vulnerability** on a specific [issue card](../../../snyk-admin/snyk-projects/issue-card-information.md) on the relevant Project page.
* If you are using a [Source code integration](../../../scm-integrations/snyk-scm-integrations/):
  * Click **Open a fix PR** on the Project page.
  * Use [automated pull requests](../../pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-new-fixes-fix-prs.md) when new fixes become available that help you to fix a vulnerability.

{% hint style="info" %}
**Automatic Fix PRs**\
When a new fixable vulnerability is found, Snyk can attempt to open a new pull request automatically. See [Automated pull request creation for new fixes](../../pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-new-fixes-fix-prs.md) for details.
{% endhint %}
