# Snyk Open Source Scans (SCA) of large manifest files, Docker setup

When large manifest files are detected by Snyk, it is sometimes necessary to use a different method (different SCM endpoint) to retrieve the file. Since the alternative method requires a more permissive rule, it is disabled by default.&#x20;

{% hint style="info" %}
This applies only to Github and Github Enterprise Broker integrations.&#x20;
{% endhint %}

To add this rule easily, allowing the Broker client to retrieve the larger manifests file using a different endpoint, add the following environment variable:

```console
-e ACCEPT_LARGE_MANIFESTS=true
```

If you are using a [custom accept.json](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/custom-approved-listing-filter-with-docker) instead of the ACCEPT environment variables, please add this to your accept.json in the private section

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
