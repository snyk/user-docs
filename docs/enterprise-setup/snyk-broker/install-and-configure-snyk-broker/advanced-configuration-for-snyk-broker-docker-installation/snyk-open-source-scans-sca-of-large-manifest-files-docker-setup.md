# Snyk Open Source Scans (SCA) of large manifest files, Docker setup

{% hint style="info" %}
This applies only to Github and Github Enterprise Broker integrations.&#x20;
{% endhint %}

When large manifest files are detected by Snyk, it is sometimes necessary to use a different method (different SCM endpoint) to retrieve the file. Since the alternative method requires a more permissive rule, it is disabled by default. This ensures the maximum possible security. Use of this endpoint means that the Snyk platform can theoretically access all files in this repository because the path does not include specific allowed file names.

To add this rule easily, allowing the Broker client to retrieve the larger manifests file using a different endpoint, add the following environment variable:

```console
-e ACCEPT_LARGE_MANIFESTS=true
```
