# Use policies in the SDLC

You can apply policies across all stages of the SDLC, from the developer’s local development environment, in the IDE or CLI, through to Git-based workflows and CI/CD, and into production.

These multiple security and compliance controls ensure issues are flagged as early as possible in the development process when it is less costly and time-consuming to fix.

{% hint style="info" %}
Additionally, the **.snyk** file is a policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins. See [The .snyk file](the-.snyk-file.md) for details
{% endhint %}

## Apply policies to Projects or Organizations

For both [security policies](security-policies/) and [license policies](license-policies/), you can&#x20;

* [Apply a policy to Projects](apply-a-policy-to-projects.md)
* [Apply a policy to Organizations](apply-a-policy-to-organizations.md)

### Example: apply a license policy to Projects

Your company’s legal team requires strict license compliance controls for business-critical frontend applications, but is less concerned about internal development projects.

First, add the `Critical`, `Production`, and `Frontend` attributes to the Snyk Projects you want this policy to apply to:

<figure><img src="../../.gitbook/assets/image (1) (3).png" alt="Add relevant attributes to a Project"><figcaption><p>Add relevant attributes to a Project</p></figcaption></figure>

Next, create a new license policy with those attributes:

<figure><img src="../../.gitbook/assets/image (7) (1).png" alt="Create license policy matching attributes"><figcaption><p>Create license policy matching attributes</p></figcaption></figure>

{% hint style="info" %}
In the policy itself, a high severity can be applied to any copyleft license identified in projects, such as the [**GPL-3.0**](https://snyk.io/learn/what-is-gpl-license-gplv3-explained/) and [**AGPL-3.0 licenses**](https://snyk.io/learn/agpl-license/). \
When creating license policies, we recommend you describe why Snyk fails the test. For example, if their build failed due to the GPL license, developers can see the explanation, so they know what action to take. See [Create a license policy and rules](license-policies/create-a-license-policy-and-rules.md) for details.
{% endhint %}

This policy is now applied to all Projects containing those attributes, and takes effect the next time Snyk scans those Projects.

See [License policies](license-policies/) for more details.

### Example: apply a **security policy to Projects**

Using a similar process to the previous example, you can define a security policy to automatically ignore all **Medium** severity vulnerabilities in the **FrontEnd** environment without a known exploit:

<div align="left">

<figure><img src="../../.gitbook/assets/image (14) (3).png" alt="Snyk security policy - ignore Medium vulns"><figcaption><p>Snyk security policy - ignore Medium vulns</p></figcaption></figure>

</div>

This policy is now applied to all Projects containing those attributes., and takes effect the next time Snyk scans those Projects.

See [Security policies](security-policies/) for more details.

## Apply policies in GitHub repos

For GitHub Projects monitored by Snyk, any new pull request from a contributing developer can be checked against policies assigned to that Project. This ensures that policy-breaking code cannot be committed to the repository.

{% hint style="info" %}
See [PR Checks](../../scan-application-code/run-pr-checks/) for details of Snyk’s PR Checks feature.
{% endhint %}

### Example: PR check on JavaScript package license

This example shows a pull request to add the `fullpage.js` package to a JavaScript application. Although this change passes the security policy check (the latest version of the package has no known vulnerability), it fails the license policy check (because of the GPLv3 license included which violates the company’s license policy).

<figure><img src="../../.gitbook/assets/image (5) (1) (1).png" alt="PR Check fail on license compliance"><figcaption><p>PR Check fail on license compliance</p></figcaption></figure>

## Apply policies in CI/CD

Policies take effect in CI/CD, ensuring builds comply with security and compliance boundaries.

### Example: Workflow High-severity vulnerability

This example shows a GitHub Actions build workflow failing because of a high-severity vulnerability identified in Snyk’s testing:

<figure><img src="../../.gitbook/assets/image (6) (4).png" alt="CI/CD check fail on security policy breach"><figcaption><p>CI/CD check fail on security policy breach</p></figcaption></figure>
