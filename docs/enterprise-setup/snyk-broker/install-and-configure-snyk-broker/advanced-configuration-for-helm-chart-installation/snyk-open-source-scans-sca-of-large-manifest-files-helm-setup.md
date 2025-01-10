# Snyk Open Source scans (SCA) of large manifest files, Helm setup

When large manifest files are detected by Snyk, it is sometimes necessary to use a different method (different SCM endpoint) to retrieve the file. Since the alternative method requires a more permissive rule, it is disabled by default.&#x20;

{% hint style="info" %}
This only applies to Github and Github Enterprise Broker integrations.&#x20;
{% endhint %}

To add this rule easily, allowing the Broker client to retrieve the larger manifests file using a different endpoint, add the following environment variable:

```console
--set 'env[0].name=ACCEPT_LARGE_MANIFESTS' \
--set 'env[0].value=true'
```

{% hint style="info" %}
If you are using other custom environment variables, you may need to adjust the 0 index in this code to avoid having any of the environment variables overwrite each other.
{% endhint %}

If you are using a custom `accept.json` instead of the ACCEPT environment variables, add this to your accept.json in the private section

```
{
    "//": "used to get given manifest file",
    "method": "GET",
    "path": "/repos/:owner/:repo/git/blobs/:sha",
    "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

{% hint style="info" %}
To ensure the maximum possible security, Snyk does not enable this rule by default, as use of this endpoint means that the Snyk platform can theoretically access all files in this repository because the path does not include specific allowed file names.
{% endhint %}
