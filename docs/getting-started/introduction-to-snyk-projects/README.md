# Introduction to Snyk projects

Projects define the items Snyk scans.

A project includes:

* A scannable item external to Snyk.
* Configuration to define how to run that scan.

Projects appear on the **Projects** menu on the Snyk dashboard:

![](../../.gitbook/assets/code1.png/)

## Target

The address of the item to scan \(external to Snyk\), such as a Kubernetes cluster or a GitHub repo. One target may relate to many projects. The structure of the target depends on the origin.

## Origin

The project ecosystem, such as CLI, API, or Kubernetes.

## Targetfile

The specific item to scan in a target, such as a pom file in a GitHub repo.

[Snyk Code](snyk-code/) scans do not use targetfiles.

## Type

The scanning method to use for this project, such as static application security testing \(SAST, for scanning using Snyk Code\) or maven for a maven project using Snyk Open Source\). Part of the configuration for scanning.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

