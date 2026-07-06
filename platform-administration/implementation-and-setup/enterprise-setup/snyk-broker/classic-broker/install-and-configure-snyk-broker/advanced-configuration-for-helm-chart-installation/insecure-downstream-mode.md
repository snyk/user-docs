---
description: Usage of this mode is discouraged.
---

# Insecure downstream mode

In some situations, you might need to use only HTTP for your downstream connection. These cases are infrequent and usually occur because of historical HTTP-only setups. Snyk recommends upgrading to use HTTPS instead.

If upgrading is not possible in the near term, the insecure downstream mode introduces a way to force downstream requests to your SCM, JIRA, or other service to take place over HTTP instead of HTTPS.

In most cases, avoid this mode; it is opt-in. It makes all requests go over HTTP, without the benefit of TLS encryption. This mode means all your credentials and data appear unencrypted, which is tolerable only in tightly secured networks.

Use the [Custom additional options for Broker Helm Chart installation](custom-additional-options-for-broker-helm-chart-installation.md) to inject this environment variable:

`--set env[0].name=INSECURE_DOWNSTREAM --set env[0].value="true"`

{% hint style="warning" %}
Using HTTP is highly insecure. Your data and credentials transit in the clear over the network exchanges.

Snyk is not responsible for any credential leaks that occur as a result of the use of HTTP.
{% endhint %}
