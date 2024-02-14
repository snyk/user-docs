---
description: Usage of this mode is discouraged.
---

# Insecure Downstream Mode

In some situations, you may need to use only http for your downstream connection. Usually because of historical http only setups, these cases are relatively infrequent, and upgrading them to use https instead is our recommendations.

In the cases where this is not possible in the near term, the insecure downstream mode introduces a way to force downstream requests to your SCM/JIRA/others to take place over http instead of https.

Using this mode should be avoided in most cases and remains opt-in.\
It makes all requests go over http, therefore not benefiting from the safety of TLS encryption. It means all your credentials and data will appear unencrypted, which is only tolerable in tightly secure networks.

Use the [Custom additional options for Broker Helm Chart installation](custom-additional-options-for-broker-helm-chart-installation.md) to inject this environment variable:

`--set env[0].name=INSECURE_DOWNSTREAM --set env[0].value="true"`

{% hint style="danger" %}
Using HTTP is highly insecure ! Your data and credentials will transit in clear over the network exchanges.

Snyk **will not be held responsible** for any credential leaks that may occur as a result of the usage this mode.
{% endhint %}
