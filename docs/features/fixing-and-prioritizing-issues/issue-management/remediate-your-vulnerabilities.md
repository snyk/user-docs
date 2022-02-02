# Fix your vulnerabilities

Snyk helps you to fix vulnerabilities, by upgrading the direct dependencies to a vulnerability-free version, or by patching the vulnerability.

To fix a vulnerability with Snyk:

* Click **Fix this vulnerability** on a specific [issue card](../../../getting-started/introduction-to-snyk-projects/issue-card-information.md) on the relevant project page.
* If you are using a [Source code integration](../../integrations/git-repository-scm-integrations/):
  * Click **Open a fix PR** on the project page.
  * Use [automated pull requests](../../../products/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities.md) when new fixes become available that help you to fix a vulnerability.

### How it works

When a new fixable vulnerability is found, Snyk attempts to open a new pull request on your behalf (in a repository for which we support automatic fix pull requests), or suggests you open one manually based on your settings.

When Snyk automates the fix, we check if there are an existing branch and pull request for the exact fix; if there is, we reopen the existing, already closed pull request on that branch.

When there’s no existing branch and pull request for the issue a new branch and pull request are created.

### Fix advice

Once Snyk tests your manifest files, we then provide a summary and detailed advice for vulnerabilities that have fixes available, enabling you to resolve those vulnerabilities in your code with the help of clear suggestions and explanations.

Snyk offers you one of these solutions:

* an **upgrade** - an upgrade to the original package
* **Pinning** a package - installing a package as a top-level dependency; that is, a specific version of an indirect dependency. This avoids a direct dependency pulling in a vulnerable version
* a Snyk precision **patch** - if an upgrade to fix any of the vulnerabilities in the package is not currently available, Snyk offers patches to fix the issues

The summary area groups advice per package, and is displayed based on the best available fix. Advice in these summary lists includes these details per package:

* All vulnerability names and severity details affecting that package
* The recommended fix - a link to the recommended fix for this package and its listed vulnerabilities: either the specific version to which to upgrade or the name of the patch

### Actionable advice from our app

From our app, for each tab (upgrade and patch) in the fix advice area of your project details, results are displayed as follows:

* the total number of packages that can be fixed is displayed on the tab title
* in groups of vulnerabilities by package, entitled by the upgrade or fix that’s recommended
* packages can be expanded in order to view the full list of vulnerabilities affecting the package
* All the vulnerabilities found in your dependencies are displayed further below, together with contextual information that can help you prioritize the issues and start fixing them if required.

The Fix Advice area appears in the project details page near the top, similar to the following examples:

![Upgrade issues tabs](<../../../.gitbook/assets/Screenshot 2021-10-12 at 14.08.13.png>)

![Patchable issues tabs](<../../../.gitbook/assets/Screenshot 2021-10-12 at 14.10.00.png>)

You can also find additional advice and details further down on the Project details page:

* from the **Issues**, tab, a full description per vulnerability
* from the **Dependencies** tab, the entire tree of your project dependencies, enabling you to clearly visualize affected paths

### Actionable advice from our CLI tool

From the CLI, for each list (upgrade and patch), results are displayed in groups based on the packages we recommend that you fix, and including:

* details for all vulnerabilities introduced per package; to view all dependency paths affected, use `--show-vulnerable-paths=all` when running `snyk test` or `snyk monitor`
* links to full descriptions of each vulnerability

Upgrade and patch results appear similar to the following:

![](<../../../.gitbook/assets/image (17).png>)

![](<../../../.gitbook/assets/image (49).png>)

Patch recommendations with some and with all paths:

![](../../../.gitbook/assets/uuid-1afca091-a9a5-d42c-40b6-f48aa0e72584-en.png)

![](<../../../.gitbook/assets/image (3).png>)
