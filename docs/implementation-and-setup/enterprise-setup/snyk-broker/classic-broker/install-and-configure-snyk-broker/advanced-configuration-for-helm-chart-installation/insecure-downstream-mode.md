---
description: Usage of this mode is discouraged.
---

# Insecure Downstream Mode

In some situations, you may need to use only HTTP for your downstream connection. These cases are relatively infrequent and usually occur because of historical http-only setups. Snyk recommends upgrading to use HTTPS instead.

In the cases where this is not possible in the near term, the insecure downstream mode introduces a way to force downstream requests to your SCM/JIRA/others to take place over HTTP instead of HTTPS.

You should avoid using this mode in most cases; it is opt-in. It makes all requests go over HTTP, thus without the benefit of the safety of TLS encryption. This mode means all your credentials and data will appear unencrypted, which is tolerable only in tightly-secured networks.

Use the [Custom additional options for Broker Helm Chart installation](custom-additional-options-for-broker-helm-chart-installation.md) to inject this environment variable:

`--set env[0].name=INSECURE_DOWNSTREAM --set env[0].value="true"`

{% hint style="warning" %}
Using HTTP is highly insecure. Your data and credentials will transit in the clear over the network exchanges.

Snyk is not responsible for any credential leaks that may occur as a result of the use of HTTP.
{% endhint %}
