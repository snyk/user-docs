# Policies

{% hint style="warning" %}
**Release status**&#x20;

Policies are available for Enterprise plans and apply only to Snyk Open Source scans.&#x20;

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

Snyk policies contain rules to define how Snyk behaves when encountering specific types of issues. With policies, you can identify types of issues based on conditions, such as `no exploit available`, and then apply actions to these issues, such as changing the severity. Thus by using customizable Snyk policies, you can define actions for specific types of issues encountered in scanning.

<div align="left">

<figure><img src="../../.gitbook/assets/image (112) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1) (2) (3).png" alt="Snyk Policy manager"><figcaption><p>Snyk Policy manager</p></figcaption></figure>

</div>

Using the Snyk Policy Manager, you can [view](view-policies.md), [create, and edit](create-and-edit-policies.md) policies.

#### The .snyk file

The `.snyk` file is a policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins. See [The .snyk file](the-.snyk-file.md) for details

## Benefits of Snyk policies

Policies give you a quick and automated way to identify and triage issues that are irrelevant to or unimportant in your application development. This reduces the noise level, saving valuable development time and allowing developers to take more responsibility for and ownership of security.

Policies help prioritize issues to address and can ensure vulnerable or non-compliant components do not escape notice. Policies are part of the governance framework of your company.

For more information, see [Use policies in the SDLC](use-policies-in-the-sdlc.md).

## Categories of policies

Snyk has security and license policies.

* [Security policies](security-policies/) define Snyk behavior in treating vulnerabilities, for example, according to severity levels or ignored issues.
* [License policies](license-policies/) define Snyk behavior in treating license issues, such as allowing or disallowing packages with certain license types and avoiding the use of packages containing incompatible licenses.

## Assign **policies to  Projects or Organizations**

Different applications may need to be scanned according to different policies. Mission-critical applications are likely to need more control than internal applications in a sandbox environment. You can establish the needed control by assigning policies to:

* [Projects](assign-policies-to-projects.md), after applying attributes to Projects and policies to attributes
* [Organizations](assign-a-policy-to-an-organization.md) in a Snyk Group
