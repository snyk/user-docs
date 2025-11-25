# Use policies in the SDLC

You can apply policies across all stages of the SDLC, from the developer’s local development environment, in the IDE or CLI, through to Git-based workflows and CI/CD, and into production.

These multiple security and compliance controls ensure issues are flagged as early as possible in the development process when it is least costly and time-consuming to fix them.

{% hint style="info" %}
In addition, the `.snyk` file is a policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins. See [The .snyk file](the-.snyk-file.md) for details
{% endhint %}

## Assign policies to Projects or Organizations

For both [security policies](security-policies/) and [license policies](license-policies/), you can apply a policy to Project attributes and to an Organization. This enables you to assign policies to Projects and to Organizations. For details, see [Assign policies to Projects](assign-policies-to-projects.md) and [Assign a policy to an Organization](assign-a-policy-to-an-organization.md).

### Example: assign a license policy to Projects

Scenario:\
The legal team in your company requires strict license compliance controls for business-critical front-end applications but is less concerned about internal development Projects.

To meet this requirement, first add the `Critical`, `Production`, and `Frontend` attributes to the Snyk Projects you want this policy to apply to:

<figure><img src="../../.gitbook/assets/image (38).png" alt="Add relevant attributes to a Project from the Issues tab"><figcaption><p>Add relevant attributes to a Project from the Issues tab</p></figcaption></figure>

Next, create a new license policy and apply the policy to those attributes:

<figure><img src="../../.gitbook/assets/image (115).png" alt="Apply license policy to selected attributes"><figcaption><p>Apply license policy to selected attributes</p></figcaption></figure>

{% hint style="info" %}
In the policy itself, a high severity can be applied to any copyleft license identified in Projects, such as the [GPL-3.0](https://snyk.io/learn/what-is-gpl-license-gplv3-explained/) and [AGPL-3.0 licenses](https://snyk.io/learn/agpl-license/).\
When you create license policies, Snyk recommends that you describe why Snyk will fail the test. Thus, for example, if a build fails due to the GPL license, developers can see the explanation, and they will know what action to take. See [Create a license policy and rules](license-policies/create-a-license-policy-and-rules.md) for details.
{% endhint %}

This policy is now assigned to all Projects with the selected attributes applied and takes effect the next time Snyk scans those Projects.

See [License policies](license-policies/) for more details.

### Example: assign a **security policy to Projects**

Using a process similar to the one in the previous example, you can define a security policy to automatically ignore all `Medium` severity vulnerabilities in the `FrontEnd` environment without a known exploit:

<div align="left"><figure><img src="../../.gitbook/assets/image (81).png" alt="Snyk security policy - specify vulnerabilities to ignore"><figcaption><p>Snyk security policy - specify vulnerabilities to ignore</p></figcaption></figure></div>

This policy is now assigned to all Projects with the selected attributes applied and takes effect the next time Snyk scans those Projects.

See [Security policies](security-policies/) for more details.

## Apply policies in GitHub repos

For GitHub Projects monitored by Snyk, any new pull request from a contributing developer can be checked against policies assigned to that Project. This ensures that policy-breaking code cannot be committed to the repository.

{% hint style="info" %}
See [PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/) for details of Snyk’s PR Checks feature.
{% endhint %}

An example follows of a PR check on a JavaScript package license.

This example shows a pull request to add the `fullpage.js` package to a JavaScript application. Although this change passes the security policy check, because the latest version of the package has no known vulnerability, it fails the license policy check because the GPLv3 license is included in violation of the license policy of the company.

<figure><img src="../../.gitbook/assets/image (76) (1).png" alt="PR Check fails on license compliance"><figcaption><p>PR Check fails on license compliance</p></figcaption></figure>

## Apply policies in CI/CD

Assigned policies take effect in CI/CD, ensuring builds comply with security and compliance boundaries.
