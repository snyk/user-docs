# Remediate your vulnerabilities

Snyk helps you to fix vulnerabilities in two ways. Either by upgrading the direct dependencies to a vulnerability free version or by patching the vulnerability.

Fixing with Snyk can either be performed in 4 different ways

1. by using the Source code integrations and clicking the Open a fix PR button on the project page
2. by clicking the Fix this vulnerability link on a specific issue card on the project page.
3. automatic pull requests - When new remediation becomes available that helps you to fix a vulnerability Snyk can open an automated pull request.
4. by using the CLI and running the `snyk wizard` command to fix node.js projects.

### How it works

When a new fixable vulnerability is found, Snyk attempts to open a new pull request on your behalf \(in a repository for which we support automatic fix pull requests\), or suggests you open one manually based on your settings.

When Snyk automates the fix, we check if there are an existing branch and pull request for the exact fix; if there is, we reopen the existing, already closed pull request on that branch.

When there’s no existing branch and pull request for the issue a new branch and pull request are created.

### Actionable remediation advice

Once Snyk tests your manifest files, we then provide summary and detailed remediation advice for vulnerabilities that have fixes available, enabling you to resolve those vulnerabilities in your code with the help of clear suggestions and explanations.

Snyk offers you one of these solutions:

* an upgrade - an upgrade to the original package
* a Snyk precision patch - if an upgrade to fix any of the vulnerabilities in the package is not currently available, Snyk offers patches to remediate the issues

The summary area groups advice per package, and is displayed based on the best available fix. Advice in these summary lists includes these details per package:

* All vulnerability names and severity details affecting that package
* The recommended fix - a link to the recommended fix for this package and its listed vulnerabilities: either the specific version to which to upgrade, or the name of the patch

### Actionable advice from our app

From our app, for each tab \(upgrade and patch\) in the Remediation Advice area of your project details, results are displayed as follows:

* the total number of packages that can be fixed is displayed on the tab title
* in groups of vulnerabilities by package, entitled by the upgrade or fix that’s recommended
* packages can be expanded in order to view the full list of vulnerabilities affecting the package
* All the vulnerabilities found in your dependencies are displayed further below, together with contextual information that can help you prioritize the issue and commence remediation if required.

The Remediation Advice area appears in the project details page near the top, similar to the following examples:

![](../../.gitbook/assets/image%20%2813%29.png)

![](../../.gitbook/assets/image%20%2816%29.png)

Upgrade and Patchable issues tabs

You can also find additional advice and details farther down in the Project details page:

* from the **Issues**, tab, a full description per vulnerability
* from the **Dependencies** tab, the entire tree of your project dependencies, enabling you to clearly visualize affected paths

### Actionable advice from our CLI tool

From the CLI, for each list \(upgrade and patch\), results are displayed in groups based on the packages we recommend that you fix, and including:

* details for all vulnerabilities introduced per package; to view all dependency paths affected, use `--show-vulnerable-paths=all` when running `snyk test` or `snyk monitor`
* links to full descriptions of each vulnerability

Upgrade and patch results appear similar to the following:

![](../../.gitbook/assets/image%20%2817%29.png)

![](../../.gitbook/assets/image%20%2849%29.png)

Patch recommendations with some and with all paths:

![](../../.gitbook/assets/uuid-1afca091-a9a5-d42c-40b6-f48aa0e72584-en.png)

![](../../.gitbook/assets/image%20%283%29.png)

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}