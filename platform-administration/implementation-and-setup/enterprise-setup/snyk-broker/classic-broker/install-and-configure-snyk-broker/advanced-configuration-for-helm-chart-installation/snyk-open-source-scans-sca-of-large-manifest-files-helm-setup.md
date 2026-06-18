# Snyk Open Source scans (SCA) of large manifest files, Helm setup

{% hint style="info" %}
This applies only to Github and Github Enterprise Broker integrations.
{% endhint %}

When Snyk detects large manifest files, you sometimes must use a different method (a different SCM endpoint) to retrieve the file. Because the alternative method requires a more permissive rule, it is disabled by default. This ensures the maximum possible security. Use of this endpoint means that the Snyk platform can theoretically access all files in this repository because the path does not include specific allowed file names.

To add this rule, allowing the Broker client to retrieve the larger manifests file using a different endpoint, add the following environment variable:

```console
--set 'env[0].name=ACCEPT_LARGE_MANIFESTS' \
--set 'env[0].value=true'
```

{% hint style="info" %}
If you use other custom environment variables, you might need to adjust the 0 index in this code to avoid having any of the environment variables overwrite each other.
{% endhint %}
