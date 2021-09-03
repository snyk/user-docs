# Licenses

## Licenses overview

{% hint style="info" %}
**Feature availability**  
Licenses are available to all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Every time you test your projects either with the CLI or from our app, your projects are scanned for not only vulnerabilities but also for license compliance. This includes all of your direct and indirect dependencies. Snyk scans your manifest files, and then checks for license issues against Snyk’s known licenses.

In order to enable customers to get started with license compliance faster, we’ve created an out of the box default policy. The default policy is a baseline which tries to answer the requirements of multiple types of applications \(SaaS, distributed, etc.\), and may be used as a starting point to calibrate additional license policies. The default policy doesn’t endorse or criticize any license.  
Different customers may have different needs and tolerance for different license types. We encourage you to make sure you made the needed changes or created new policies that fit your company’s specific requirements.

By [default](https://snyk.gitbook.io/user-docs/fixing-and-prioritizing-issues/policies/shared-policies-overview) we determine the severity of licenses in the following way:

* High severity - licenses that definitely present issues for commercial software.
* Medium severity - licenses that have clauses that may be of concern and should be reviewed.

New licenses added by Snyk will inherit the Unknown license type severity. In cases where this severity was not set to None, newly added licenses will appear in the licenses compliance results.

In case you notice a license with the wrong license assigned to it, you can reach out to our support team. We will investigate the request and update the license if needed.

To facilitate onboarding of your developers, we recommend that your teams check these defaults, update severities and add instructions per license type based on the policies outlined specifically by your Legal teams. Once updated, when Snyk detects a license violation it displays the violation for all users in the organization from our UI project area, or from the CLI Snyk test results, in the same way as a security vulnerability, and including the severity and instructions you configured.

![](../../.gitbook/assets/image%20%282%29.png)

### **An inventory of your licenses**

Within the **Reports** area you can view an inventory of all of your licenses across all your projects. Snyk also lists packages that have dual licenses and multiple licenses. See [here](https://support.snyk.io/hc/articles/360003557857#UUID-627d8387-74c4-9228-477b-263417eb2a61) for more information.

### **Supported package managers**

* npm
* RubyGems
* Maven
* Pip
* Nuget
* Go
* Composer
* Cocoapods

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}