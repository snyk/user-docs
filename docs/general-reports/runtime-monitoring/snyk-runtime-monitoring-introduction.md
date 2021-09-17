# Snyk runtime monitoring: introduction

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

Snyk Runtime Monitoring does a couple of important things for your open source components during runtime. It:

* Determines whether a vulnerable dependency is indeed being used at runtime in a way that can be exploited
* Flags vulnerable dependencies identified at runtime, and prioritizes them first by whether those dependencies were called at runtime and then by severity

With the data Snyk retrieves and highlights for you, you can focus your remediation efforts where they matter the most - fixing the vulnerabilities whose vulnerable functions are actually invoked at runtime.

## Snyk runtime monitoring: how it works

The Snyk runtime agent does the following:

1. The agent inspects every dependency of your application.
2. It then creates an execution hook on vulnerable functions in relevant dependencies.
3. Using these hooks, the agent detects the actual use of vulnerable functions.
4. The agent sends this data in beacons to Snyk, adding relevant data to the Snyk project.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page/)
{% endhint %}

