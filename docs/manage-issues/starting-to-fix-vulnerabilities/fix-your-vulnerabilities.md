# Fix your vulnerabilities

Snyk helps you to fix vulnerabilities, by upgrading the direct dependencies to a vulnerability-free version, or by patching the vulnerability. After Snyk scans your Projects, scan results allow you to resolve those issues in your code, with the help of clear suggestions and explanations.

You can:

* [View scan results on the Snyk Web UI](fix-your-vulnerabilities.md#view-scan-results-on-the-snyk-web-ui)
* [View scan results using Snyk CLI](fix-your-vulnerabilities.md#view-scan-results-using-snyk-cli)
* [Apply fixes](fix-your-vulnerabilities.md#apply-fixes)

## View scan results on the Snyk Web UI

From our app, for each tab (upgrade and patch) in the fix advice area of your project details, results are displayed as follows:

* the total number of packages that can be fixed is displayed on the tab title
* in groups of vulnerabilities by package, entitled by the upgrade or fix that’s recommended
* packages can be expanded in order to view the full list of vulnerabilities affecting the package
* All the vulnerabilities found in your dependencies are displayed further below, together with contextual information that can help you prioritize the issues and start fixing them if required.

<figure><img src="../../.gitbook/assets/Screenshot 2023-03-15 at 12.14.06.png" alt="Scan results on Web UI"><figcaption><p>Scan results on Web UI</p></figcaption></figure>



### View Fix Advice

The Fix Advice area appears in the project details page:

Snyk offers you one of these solutions:

* an **upgrade** - an upgrade to the original package
* **Pinning** a package - installing a package as a top-level dependency; that is, a specific version of an indirect dependency. This avoids a direct dependency pulling in a vulnerable version
* a Snyk precision **patch** - if an upgrade to fix any of the vulnerabilities in the package is not currently available, Snyk offers patches to fix the issues

The summary area groups advice per package, and is displayed based on the best available fix. Advice in these summary lists includes these details per package:

* All vulnerability names and severity details affecting that package
* The recommended fix - a link to the recommended fix for this package and its listed vulnerabilities: either the specific version to which to upgrade or the name of the patch

![Upgrade issues tabs](<../../.gitbook/assets/Screenshot 2021-10-12 at 14.08.13.png>)

![Patchable issues tabs](<../../.gitbook/assets/Screenshot 2021-10-12 at 14.10.00 (1).png>)

You can also find additional advice and details further down on the Project details page:

* from the **Issues**, tab, a full description per vulnerability
* from the **Dependencies** tab, the entire tree of your project dependencies, enabling you to clearly visualize affected paths

## View scan results using Snyk CLI

From the CLI, for each list (upgrade and patch), results are displayed in groups based on the packages we recommend that you fix, and including:

* details for all vulnerabilities introduced per package; to view all dependency paths affected, use `--show-vulnerable-paths=all` when running `snyk test` or `snyk monitor`
* links to full descriptions of each vulnerability

Upgrade and patch results appear similar to the following:

<figure><img src="../../.gitbook/assets/image (17) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Upgrade results in the CLI"><figcaption><p>Upgrade results in the CLI</p></figcaption></figure>

Patch recommendations:

<figure><img src="../../.gitbook/assets/uuid-1afca091-a9a5-d42c-40b6-f48aa0e72584-en.png" alt="Patch results in the CLI"><figcaption><p>Patch results in the CLI</p></figcaption></figure>

## Apply fixes

To apply fixes, you can:

* Click **Fix this vulnerability** on a specific [issue card](../introduction-to-snyk-projects/issue-card-information.md) on the relevant project page.
* If you are using a [Source code integration](../../integrations/git-repository-scm-integrations/):
  * Click **Open a fix PR** on the project page.
  * Use [automated pull requests](../../products/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities.md) when new fixes become available that help you to fix a vulnerability.

{% hint style="info" %}
**Automatic Fix PRs**\
When a new fixable vulnerability is found, Snyk can attempt to open a new pull request automatically. See [Automated pull request creation for new fixes](../../products/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities.md) for details.


{% endhint %}

