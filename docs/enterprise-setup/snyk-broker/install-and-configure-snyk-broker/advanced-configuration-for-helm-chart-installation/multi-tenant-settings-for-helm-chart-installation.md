# Multi-tenant settings for Helm chart installation

## **Broker multi-tenant settings**

To use the Helm chart with different multi-tenant environments, set the `brokerServerUrl` to be one of the following URLs, depending on which environment you are using:

Europe: `https://broker.eu.snyk.io`\
Australia: `https://broker.au.snyk.io`

```
--set brokerServerUrl=<BROKER_SERVER_URL>
```

## **Code Agent multi-tenant settings**

{% hint style="warning" %}
**Deprecated**

The Code Agent is deprecated and is no longer maintained.

The preferred method of running Snyk Code analysis using Snyk Broker is through [Brokered Code](../advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md). The Code Agent is an alternative method without advantages. For details, contact your Snyk Integration Consultant or Technical Success Manager or contact [Snyk Support](https://support.snyk.io).

The automatic [PR Checks](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks) feature is not supported for Snyk Broker - Code Agent.
{% endhint %}

If you are using Code Agent, the `upstreamUrlCodeAgent` value must be one of the following URLs, depending on which environment you are using:

Europe: `https://deeproxy.eu.snyk.io`\
Australia: `https://deeproxy.au.snyk.io`

```
--set upstreamUrlCodeAgent=<UPSTREAM_URL>
```
