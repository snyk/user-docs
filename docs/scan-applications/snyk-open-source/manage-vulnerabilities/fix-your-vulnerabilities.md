# Fix your vulnerabilities

Snyk helps you to fix vulnerabilities, by upgrading the direct dependencies to a vulnerability-free version, or by patching the vulnerability. After Snyk scans your Projects, scan results allow you to resolve those issues in your code, with the help of clear suggestions and explanations.

You can:

* [View scan results on the Snyk Web UI](fix-your-vulnerabilities.md#view-scan-results-on-the-snyk-web-ui)
* [View scan results using Snyk CLI](fix-your-vulnerabilities.md#fixing-vulnerabilities-based-on-scan-results-using-snyk-cli)
* [Apply fixes](fix-your-vulnerabilities.md#apply-fixes)

## View scan results on the Snyk Web UI

From our app, for each tab (upgrade and patch) in the fix advice area of your Project details, results are displayed as follows:

* The total number of packages that can be fixed is displayed on the tab title.
* In groups of vulnerabilities by package, entitled by the upgrade or fix that’s recommended.
* Packages can be expanded in order to view the full list of vulnerabilities affecting the package.
* All the vulnerabilities found in your dependencies are displayed further below, together with contextual information that can help you prioritize the issues and start fixing them if required.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-15 at 12.14.06.png" alt="Scan results on Web UI"><figcaption><p>Scan results on Web UI</p></figcaption></figure>

### View Fix Advice

The Fix Advice area appears on the Project details page:

Snyk offers you one of these solutions:

* An **upgrade** - an upgrade to the original package
* **Pinning** a package - installing a package as a top-level dependency; that is, a specific version of an indirect dependency. This avoids a direct dependency pulling in a vulnerable version
* A Snyk precision **patch** - if an upgrade to fix any of the vulnerabilities in the package is not currently available, Snyk offers patches to fix the issues

The summary area groups advice per package, and is displayed based on the best available fix. Advice in these summary lists includes these details per package:

* All vulnerability names and severity details affecting that package
* The recommended fix - a link to the recommended fix for this package and its listed vulnerabilities: either the specific version to which to upgrade or the name of the patch

<figure><img src="../../../.gitbook/assets/Screenshot 2021-10-12 at 14.08.13.png" alt="Upgrade issues tabs"><figcaption><p>Upgrade issues tabs</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/Screenshot 2021-10-12 at 14.10.00.png" alt="Patchable issues tabs"><figcaption><p>Patchable issues tabs</p></figcaption></figure>

You can also find additional advice and details further down on the Project details page:

* From the **Issues**, tab, a full description per vulnerability
* From the **Dependencies** tab, the entire tree of your project dependencies, enabling you to clearly visualize affected paths

## Fixing vulnerabilities based on scan results using Snyk CLI

For information about using the CLI in fixing vulnerabilities, see [Fix vulnerabilities using the Snyk CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/fix-vulnerabilities-using-the-snyk-cli.md).

## Apply fixes

To apply fixes, you can:

* Click **Fix this vulnerability** on a specific [issue card](../../../manage-issues/snyk-projects/issue-card-information.md) on the relevant Project page.
* If you are using a [Source code integration](../../../integrations/git-repository-scm-integrations/):
  * Click **Open a fix PR** on the Project page.
  * Use [automated pull requests](../../../scan-application-code/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities.md) when new fixes become available that help you to fix a vulnerability.

{% hint style="info" %}
**Automatic Fix PRs**\
When a new fixable vulnerability is found, Snyk can attempt to open a new pull request automatically. See [Automated pull request creation for new fixes](../../../scan-application-code/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities.md) for details.
{% endhint %}
