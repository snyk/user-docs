# Build your own custom rules

{% hint style="info" %}
This feature is currently in beta. We would appreciate any feedback you might have.
{% endhint %}

Snyk IaC includes a comprehensive list of security rules, covering AWS, Azure, GCP & Kubernetes. These rules are based on security research, best practices and recognised standards and benchmarks. They are actively maintained by Snyk’s security engineering team and new rules are released on a regular basis.&#x20;

These rules aim to meet the majority of your needs on your first scan, but you may need to enforce additional security rules for your system, such as tagging standards.&#x20;

#### Creating additional Snyk IaC Custom Rules

The IaC SDK helps security teams define their own rules, to be run by the [Snyk CLI](../snyk-cli-for-infrastructure-as-code/) giving feedback to developers.

Using this SDK, you can add your own custom rules to Snyk IaC, to run alongside the standard provided rules, giving comprehensive security feedback to your development teams in one place.

Initial instructions to get you started with the Snyk Infrastructure as Code (IaC) SDK:

* [Install the SDK](install-the-sdk.md)
* [Getting started with the SDK](getting-started-with-the-sdk/)
* [How to run custom rules with the Snyk CLI](use-IaC-custom-rules-with-CLI/)
* [How to integrate custom rules within a pipeline](integrating-iac-custom-rules-within-a-pipeline.md)
* [SDK reference](sdk-reference.md)

![End to end flow of writing your own custom rules to distributing and using them to scan files with the Snyk CLI](<../../../.gitbook/assets/image (77) (1) (1).png>)

#### Snyk platform policies and Snyk IaC custom rules

{% hint style="info" %}
Summary:&#x20;

* Snyk platform policies: manage issues
* Snyk IaC custom rules: generate issues&#x20;
{% endhint %}

The Snyk platform allows you to create your own [policies](../../../features/fixing-and-prioritizing-issues/policies/), to manage how you prioritize and triage the issues Snyk identifies during scanning. For example, you can define policies to change the priority of an issue from medium to high if it has specific attributes, or to bulk ignore issues if they meet certain criteria.

The Snyk IaC custom rules functionality enables you to define your own rules for misconfiguration checks that you would like to enforce. The result of a custom rule failing on a configuration file will generate an issue.



\
