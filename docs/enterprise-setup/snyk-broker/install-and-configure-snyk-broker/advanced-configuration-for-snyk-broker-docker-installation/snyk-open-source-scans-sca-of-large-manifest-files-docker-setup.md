# Snyk Open Source Scans (SCA) of large manifest files, Docker setup

When large manifest files are detected by Snyk, it is sometimes necessary to use a different method (different SCM endpoint) to retrieve the file. Since the alternative method requires a more permissive rule, it is disabled by default.

To add this rule easily, allowing the Broker client to retrieve the larger manifests file using a different endpoint, add the following environment variable:

```console
-e ACCEPT_LARGE_MANIFESTS=true
```
