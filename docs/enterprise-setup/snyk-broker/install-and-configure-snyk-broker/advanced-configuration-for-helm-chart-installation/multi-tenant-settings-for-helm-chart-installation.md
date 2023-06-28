# Multi-tenant settings for Helm chart installation

## **Broker multi-tenant settings**

To use the Helm chart with different multi-tenant environments, set the `brokerServerUrl` to be one of the following URLs, depending which environment you are using:

Europe: `https://broker.eu.snyk.io`\
Australia: `https://broker.au.snyk.io`

```
--set brokerServerUrl=<BROKER_SERVER_URL>
```

## **Code Agent multi-tenant settings**

If you are using Code Agent, the `upstreamUrlCodeAgent` value must be one of the following URLs, depending which environment you are using:

Europe: `https://deeproxy.eu.snyk.io`\
Australia: `https://deeproxy.au.snyk.io`

```
--set upstreamUrlCodeAgent=<UPSTREAM_URL>
```
