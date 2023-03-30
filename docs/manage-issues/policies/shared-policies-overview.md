# Introduction to policies

{% hint style="warning" %}
Snyk policies currently only apply to Snyk Open Source scans.
{% endhint %}

Snyk Policies contain rules defining how Snyk behaves when encountering certain types of issues. With policies, you can identify types of issues based on conditions (such as “no exploit available”), then apply actions to these issues (such as changing the severity).

### Benefits of Snyk policies

Policies give you a quick and automated way to identify and triage issues that are not important or relevant to your application development. This saves valuable development time, and allows developers to take more responsibility and ownership for security, reducing the “noise” level.

Policies help prioritize which issues to address, and can ensure vulnerable or non-compliant components do not slip through the cracks. Policies are part of your company’s governance framework.

### Categories of policy

Snyk policies include:

* [Security policies](security-policies/): defines Snyk behavior for treating vulnerabilities. For example, to change severity levels or ignore issues.
* [License policies](license-policies/)**:** defines Snyk behavior for treating license issues. For example, to allow or disallow packages with certain license types, to avoid using packages containing incompatible licenses.

### Apply **policies to  Projects or Organizations**

Different applications may need different policies applied; mission-critical applications will likely need more control than internal applications in a sandbox environment.

You can enable this control, by applying policies to:

* [Projects](https://docs.snyk.io/manage-issues/policies/assign-a-policy-to-project-attributes), using Project tags and attributes.
* [Organizations](https://docs.snyk.io/manage-issues/policies/assign-a-policy-to-organizations) in a Snyk Group.

