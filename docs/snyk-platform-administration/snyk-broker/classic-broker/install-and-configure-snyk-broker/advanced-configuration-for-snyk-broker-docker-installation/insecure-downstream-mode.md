---
description: Usage of this mode is discouraged.
---

# Insecure downstream mode

In some situations, you may need to use only `http` for your downstream connection. These cases are infrequent and usually occur because of historical `http-`only setups. Snyk recommends upgrading these setups to use `https` instead.

If upgrading is not possible in the near term, the insecure downstream mode introduces a way to force downstream requests to your SCM, JIRA, or other service to take place over `http` instead of `https`.

In most cases, you should avoid insecure downstream mode, and to use it, you must opt-in.\
Insecure downstream mode makes all requests go over `http`, thus no longer benefitting from the safety of TLS encryption. Insecure downstream mode means all your credentials and data will appear unencrypted, which is only tolerable in tightly secured networks.

Export the environment variable `INSECURE_DOWNSTREAM="true"` to use this mode, passing it as an environment value using the `-e INSECURE_DOWNSTREAM="true"` option.

{% hint style="warning" %}
Using HTTP is highly insecure! Your data and credentials will be transmitted in clear form over the network exchanges.

Snyk is not responsible for any credential leaks that may occur as a result of the use of an insecure downstream mode.
{% endhint %}
