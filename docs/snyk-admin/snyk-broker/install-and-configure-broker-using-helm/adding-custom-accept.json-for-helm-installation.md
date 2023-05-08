# Adding custom accept.json for Helm installation

To add a custom `accept.json` file, include it in `values.yaml`.

See the example `accept.json` files in the [client templates on the Snyk Broker Helm repository](https://github.com/snyk/broker/tree/master/client-templates).

Structure the `values.yaml` file like this:

```
scmType: github-com
brokerToken: <ENTER_BROKER_TOKEN>
scmToken: <ENTER_REPO_TOKEN>
acceptJson: |-
  {
    "public": [
      {
        "//": "used for pushing up webhooks from github",
        "method": "POST",
        "path": "/webhook/github",
        "valid": [
          {
            "//": "accept all pull request state changes (these don't have files in them)",
            "path": "pull_request.state",
            "value": "open"
          },
          {
            "path": "commits.*.added.*",
            "value": "package.json"
          },
  ...
    ]
  }
```

You can then install:

```
helm install snyk-broker-gitub-com snyk-broker/snyk-broker -f values.yaml -n snyk-broker --create-namespace
```
