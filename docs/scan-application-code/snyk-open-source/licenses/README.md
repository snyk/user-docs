# Licenses

## Licenses overview

{% hint style="info" %}
**Feature availability**\
Licenses are available to all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Every time you test your code in the [Snyk Web UI](../../../getting-started/exploring-the-snyk-web-ui.md), the [Snyk CLI](../../../snyk-cli/), or using [PR Checks](../../run-pr-checks/), your repositories are scanned not only for vulnerabilities but also for license compliance. This includes all of your direct and indirect dependencies. Snyk scans your manifest files, and then checks for license issues against Snyk’s known licenses.

### Default license policy

To enable customers to get started with license compliance faster, we created an out-of-the-box default policy. The default policy is a baseline that tries to answer the requirements of multiple types of applications (SaaS, distributed, etc.), and may be used as a starting point to calibrate additional license policies. The default policy does not endorse or criticize any license.

By [default](../../../manage-issues/policies/introduction-to-policies.md) we determine the severity of licenses in the following way:

* High severity - licenses that definitely present issues for commercial software.
* Medium severity - licenses that have clauses that may be of concern and should be reviewed.

### Manage license policies

Different customers may have different needs and tolerance for different license types. We encourage you to make sure you made the needed changes or created new policies that fit your company’s specific requirements.

See [License policies](../../../manage-issues/policies/license-policies/) for more details.

New licenses added by Snyk will inherit the **Unknown** license type severity. In cases where this severity was not set to **None**, newly added licenses will appear in the licenses compliance results.

If you notice a license with the wrong license assigned to it, you can reach out to our support team. We will investigate the request and update the license if needed.

### Using licenses

To facilitate the onboarding of your developers, we recommend that your teams check these defaults, update severities, and add instructions per license type based on the policies outlined specifically by your Legal teams.&#x20;

After being updated, when Snyk detects a license violation, it displays the violation for all users in the Organization on the test results, on the [Snyk Web UI](../../../getting-started/exploring-the-snyk-web-ui.md), the [Snyk CLI](../../../snyk-cli/), or using [PR Checks](../../run-pr-checks/). in the same way as a security vulnerability, and including the severity and instructions you configured.

For example:

<div align="left">

<figure><img src="../../../.gitbook/assets/image5 (2).png" alt="License card overview."><figcaption><p>License card overview</p></figcaption></figure>

</div>

## **An inventory of your licenses**

You can view an inventory of all of your licenses across all your Projects. Snyk also lists packages that have dual licenses and multiple licenses. See [View licenses](../../../manage-issues/dependencies-and-licenses/view-licenses.md) for more information.

## **Supported package managers**

{% hint style="warning" %}
Snyk does not support package versions which include a git commit hash.
{% endhint %}

* npm
* RubyGems
* Maven
* Pip
* Nuget
* Go
* Composer
* Cocoapods

## **Supported unmanaged Ecosystems**

{% hint style="warning" %}
Snyk does not support package versions which include a git commit hash.
{% endhint %}

* C/C++
