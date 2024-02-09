# Current IaC custom rules

{% hint style="warning" %}
**Release status**

IaC custom rules are available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

Snyk IaC includes a comprehensive list of security rules, covering AWS, Azure, GCP, and Kubernetes. These rules are based on security research, best practices, recognized standards. and benchmarks. They are actively maintained by Snyk’s security engineering team, and new rules are released on a regular basis.

These rules are intended to meet most of your needs on your first scan, but you may need to enforce additional security rules for your system, such as tagging standards.

#### Training - custom rules

[See Snyk Learn](https://learn.snyk.io/lesson/iac-custom-rules/) for interactive training on IaC custom rules.

## Creating additional Snyk IaC Custom Rules

The IaC SDK helps security teams define their own rules, to be run by the [Snyk CLI](../../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/), providing feedback to developers.

Using this SDK, you can add your own custom rules to Snyk IaC to run alongside the standard provided rules, giving comprehensive security feedback to your development teams in one place.

This section provides initial instructions to help you use the Snyk Infrastructure as Code (IaC) SDK:

* [Install the SDK](install-the-sdk.md)
* [Getting started with the SDK](writing-rules-using-the-sdk/)
* [Use IaC custom rules with the Snyk CLI](use-iac-custom-rules-with-cli/)
* [Integrating custom rules within a pipeline](integrated-iac-custom-rules-within-a-pipeline.md)
* [SDK reference](sdk-reference.md)

<figure><img src="../../../../.gitbook/assets/image (159) (1) (1) (2).png" alt="End to end flow of writing your own custom rules, distributing them, and using them to scan files with the Snyk CLI"><figcaption><p>End to end flow of writing your own custom rules, distributing them, and using them to scan files with the Snyk CLI</p></figcaption></figure>

## Snyk platform policies and Snyk IaC custom rules

{% hint style="info" %}
Summary:

* Snyk platform policies: manage issues
* Snyk IaC custom rules: generate issues
{% endhint %}

The Snyk platform allows you to create your own [policies](../../../policies/) to manage how you prioritize and triage the issues Snyk identifies during scanning. For example, you can define policies to change the priority of an issue from medium to high if it has specific attributes, or to bulk ignore issues if they meet certain criteria.

The Snyk IaC custom rules functionality enables you to define your own rules for misconfiguration checks that you would like to enforce. The result of a custom rule failing on a configuration file is generating an issue.
