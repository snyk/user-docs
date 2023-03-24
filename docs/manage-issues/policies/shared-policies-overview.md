# Introduction to policies

{% hint style="info" %}
**Feature availability**\
This feature is available to Enterprise customers. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="info" %}
You must be a Group administrator to access policies for that Group.
{% endhint %}

**Snyk** **Policies** are rules defining how Snyk behaves when encountering certain types of issues. With policies, you can identify types of issues based on conditions (such as “no exploit available”), then apply actions to these issues (such as changing the severity).

Snyk policies include:

* [Security policies](../security-policies/): defines Snyk behavior for treating vulnerabilities. For example, to change severity levels or ignore issues.
* [License policies](../../scan-application-code/snyk-open-source/license-policies/)**:** defines Snyk behavior for treating license issues. For example, to allow or disallow packages with certain license types, to avoid using packages containing incompatible licenses.

{% hint style="info" %}
Snyk policies currently only apply to Snyk Open Source scans.
{% endhint %}

### Benefits of Snyk policies

Policies give you a quick and automated way to identify and triage issues that are not important or relevant to your application development. This saves valuable development time, and allows developers to take more responsibility and ownership for security, reducing the “noise” level.

Policies help prioritize which issues to address, and can ensure vulnerable or non-compliant components do not slip through the cracks. Policies are part of your company’s governance framework.

### Using policies

* [View policies](view-policies.md) (in the policy manager)
* [Create and edit policies](create-and-edit-policies.md)
