# Snyk API

Snyk’s [extensibility and API](https://snyk.io/blog/extensibility-and-the-snyk-api/) enable developers to tune Snyk’s security automation to their specific workflows, ensuring consistency in both developer experience and platform governance.

## Snyk API v1

The [**Snyk API v1**](https://snyk.docs.apiary.io/https://snyk.docs.apiary.io/) has the ability to test a package for issues as they are defined by Snyk, and to provide Snyk security automation according to your own workflows, unconstrained by security processes in Snyk products. Customers and partners can perform functions including:

* Accessing vulnerability data
* Scanning projects and applications
* Receiving remediation advice
* Viewing user data to build custom security solutions

## Snyk REST API

{% hint style="info" %}
The Snyk REST API was formerly known as Snyk API v3.
{% endhint %}

The [**Snyk REST API**](https://apidocs.snyk.io/) is available for you to try out as endpoints are released. It is based on the OpenAPI and JSON:API standards and represents an evolutionary approach to API development, with each endpoint versioned (see [Versioning](https://apidocs.snyk.io/#overview) in the reference docs). The Snyk REST API ultimately will replace the API v1.

## When to use the API versus the CLI or an integration

Note that there may be differences in output of the API, the CLI, and integrations.

For example, for many package managers using the API will be less accurate than running the Snyk CLI as part of your build pipe or locally on your package. More than one version of a package may fit the requirements in manifest files. Running the CLI locally tests the actual deployed code, and has an accurate snapshot of the dependency versions in use. The API infers a snapshot, with inferior accuracy. Note that the Snyk CLI can output machine-readable JSON (`snyk test --json`).

You can allow Snyk access to your development flow by using Snyk integrations. The advantage is having Snyk monitor every new pull request and suggest fixes by opening new pull requests. You can integrate Snyk directly to your source code management (SCM) tool, or via a broker to allow greater security and auditability.

Use the API when you want to customize, integrate, and automate Snyk security as part of your specific workflows.

{% hint style="info" %}
**Feature availability**\
The Snyk API is available with Business and Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}
