# Preflight checks for Snyk Broker

The main objective of preflight checks is to catch errors and misconfigurations early, on Broker Client startup, rather than later during use. Whether or not the checks are successful, the Broker Client starts. The following checks are available.

## `broker-server-status`

The Broker Server Healthcheck validates the connectivity to the Broker Server. It performs a GET request to `{BROKER_SERVER_URL}/healthcheck.`

If it is not specified, BROKER\_SERVER\_URL is [https://broker.snyk.io](https://broker.snyk.io/)

## `rest-api-status`

The REST API Healthcheck validates the connectivity to the Snyk REST API by performing a GET request to `{API_BASE_URL}/rest/openapi`. This check is conditional and will be executed only if high availability mode is enabled.

If it is not specified,  the `API_BASE_URL` is [https://api.snyk.io](https://api.snyk.io/). For additional URLs, see [Regional hosting and data residency](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md).

## `broker-client-url-validation`

The Broker Client Configuration Check validates the `BROKER_CLIENT_URL` value to the extent possible. The check verifies that the configuration contains a scheme (`http or https`) and if the scheme is `https`, that an SSL certificate+key is loaded or TLS-termination upstream is used.

If you are using TLS termination and you do not require a certificate+key in the Broker Client, add the environment variable `BROKER_CLIENT_URL_TLS_TERMINATED` to signal TLS termination in the preflight check.

There is no default.

{% hint style="info" %}
You can use the environment variable PREFLIGHT\_CHECKS\_ENABLED=false to disable the Preflight Checks feature, so no checks will be executed when the Broker Client starts.
{% endhint %}

