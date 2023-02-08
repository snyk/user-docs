# Snyk Broker

Snyk Broker is an open source tool that can act as a proxy between Snyk and special integrations including:

* Your Source Code Management (SCM) system on-premise platforms
* Your publicly-accessible Git-based repositories, allowing you to view and control Snyk activity in those repositories for increased data security
* Your on-premise Jira installation or JFrog Artifactory installation
* Network-restricted [container registries](snyk-broker-container-registry-agent/)
* [Infrastructure as code (IaC) configuration](snyk-broker-infrastructure-as-code-detection/) files using Snyk IaC located on private Git-based repositories

Snyk Broker is an open-source project, hosted at [GitHub](https://github.com/snyk/broker) and published as a set of Docker images for specific integrations. For additional information see the [Broker documentation on GitHub](https://github.com/snyk/broker/blob/master/README.md).

{% hint style="info" %}
**Feature availability**\
Snyk Broker is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="warning" %}
**Multi-tenant settings for EU and AU**\
When you have up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example:  `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).
{% endhint %}
