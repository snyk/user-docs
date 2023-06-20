# Snyk Open Source Scans (SCA) of large manifest files, Helm setup

When large manifest files are detected by Snyk, it is sometimes necessary to use a different method (different SCM endpoint) to retrieve the file. Since the alternative method requires a more permissive rule, it is disabled by default.

To add this rule easily, allowing the Broker client to retrieve the larger manifests file using a different endpoint, add the following environment variable:

```console
--set env[0].name=ACCEPT_LARGE_MANIFESTS \
--set env[0].value=true
```

{% hint style="info" %}
If you are using other custom environment variables, you may need to adjust the 0 index in this code to avoid having any of the environment variables overwrite each other.
{% endhint %}
