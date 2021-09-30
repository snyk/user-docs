# Getting Started with the SDK

To get you started with the SDK, you will learn how to:

1. [Parse a fixture file](parsing-an-input-file.md) in order to help you understand how to write a rule
2. [​Write a rule with the SDK](writing-a-rule.md) using Rego
3. [Add unit tests for the rules you’ve written](testing-a-rule.md) to verify your rules
4. [Build a bundle containing your custom rules](bundling-rules.md) so that you can [use it with the Snyk CLI](../how-to-run-custom-rules-with-the-snyk-cli.md)

### Rules in Rego

Rules are written in Rego. When you are writing Rego, you do two things:

1. Write **rules** that make policy decisions. A rule is a conditional assignment.
2. Organise rules into **policies**. A policy is a set of rules with a hierarchical name

A Rego decision policy is a JSON document. Rego values are JSON values plus Sets.

{% hint style="info" %}
To learn more about the Policy Language, please visit the official [OPA Policy Language Documentation Page](https://www.openpolicyagent.org/docs/latest/policy-language/).
{% endhint %}

{% hint style="info" %}
You can also use the [OPA Playground](https://play.openpolicyagent.org/) to try out Rego, or run examples of this guide.
{% endhint %}

