# Overview

{% hint style="info" %}
The majority of Snyk APIs are restricted to use by Enterprise plan customers only.

For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

The Snyk API allows customers to integrate programmatically with Snyk.

The [Snyk REST API ](rest-api/about-the-rest-api.md) and the [V1 API](v1-api.md) are available in the Snyk API [Reference](reference/). Additional endpoints are available in the [OAuth2 API reference](oauth2-api.md).

## Automating Snyk processes

The Snyk API enables developers to automate Snyk processes to accomplish their specific workflows, ensuring consistency in both developer experience and platform governance.

Use the API when you want to customize, integrate, and automate Snyk processes as part of your specific workflows.

## Deciding to use the API, the CLI, or an SCM integration

You may decide to use an API rather than the CLI or an integration. When you decide, consider possible differences in the output for the API, the CLI, and integrations.

For example, for many package managers, using the API will be less accurate than running the Snyk CLI as part of your build pipe or locally on your package. More than one version of a package may fit the requirements in manifest files.

Running the CLI locally tests the actual deployed code and creates an accurate snapshot of the dependency versions in use. The API infers a snapshot, with inferior accuracy. Note that the Snyk CLI can output machine-readable JSON (`snyk test --json`).

You can allow Snyk access to your development flow by using Snyk integrations. The advantage is having Snyk monitor every new pull request and suggest fixes by opening new pull requests. You can integrate Snyk directly with your source code management (SCM) tool, or by using a Broker to allow greater security and auditability.
