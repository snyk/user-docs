# IaC+ and cloud custom rules

{% hint style="info" %}
**Feature availability**\
Custom rules are available to organizations with Enterprise plans enabled with Integrated IaC. [Contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new) for more information.
{% endhint %}

## Prerequisites for IaC to cloud custom roles

Install the following in your PATH: Snyk CLI >= 1.1168.0\
For details, see [Install or update the Snyk CLI](../../../../snyk-cli/install-or-update-the-snyk-cli/).

Configure the CLI appropriately and set your default Organization:\
`snyk config set org=<org id>`

If you work with multiple Snyk Organizations, you can add `--org=<your org id>` to your commands to specify your desired Organization.

## Predefined security rules in Snyk IaC

Snyk IaC includes a set of security rules that work out of the box, covering AWS, Azure, GCP, and Kubernetes. These rules are based on security research, best practices, recognized standards, and benchmarks. New rules are released regularly. Snyk’s security engineering team actively maintains them.

These rules are intended to meet most of your needs on your first scan, but you may need to enforce additional security rules for your system, such as tagging standards.

## Purpose of Cloud Custom Rules

Complementing Snyk’s predefined rules, IaC to Cloud Custom Rules enables you to enforce your internal security controls across your SDLC (software development lifecycle). Using Cloud Custom Rules, you can identify and highlight the following:

* Issues on cloud configurations across the SDLC, including SCM, CLI, Terraform Cloud, and deployed cloud environments
* Issues on any Terraform IaC configurations using Terraform providers beyond cloud (AWS, Azure, Google Cloud) configurations, such as GitHub or Snowflake configurations.

The following are the steps in using Cloud Custom Rules:

1. [Create a Custom Rules Project](create-a-custom-rules-project.md)
2. [Create a Custom Rule](create-a-custom-rule.md)
3. [(Recommended) Write a Custom Rule Spec (Unit Test)](recommended-write-a-custom-rule-spec-unit-test.md)
4. [(Recommended) Test your Custom Rule](recommended-test-your-custom-rule.md)
5. [Build, bundle, and upload your Custom Rules](build-bundle-and-upload-your-custom-rules.md)
6. [Enforce your Custom Rules](enforce-your-custom-rules.md)
7. [Explore example Custom Rules](explore-example-custom-rules.md)
