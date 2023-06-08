# Test your Kubernetes files with Snyk CLI

With Snyk Infrastructure as Code, you can test your configuration files with the CLI. Snyk Infrastructure as Code for Kubernetes supports the following:

* Deployments, Pods, and Services
* CronJobs, Jobs, StatefulSet, ReplicaSet, DaemonSet, and ReplicationController

Use the following Snyk CLI command to test for an issue on specified files:

```
snyk iac test
```

Example:

```
snyk iac test deploy.yaml
```

You can also specify multiple files by appending the file names after each other, for example:

```
snyk iac test file-1.yaml file-2.yaml
```

For the steps to scan a Helm chart using the Snyk CLI, see [Test your Helm charts with Snyk CLI](test-your-helm-charts-with-snyk-cli.md).

For the steps to scan a Kustomize template using the Snyk CLI, see [Test your Kustomize files with Snyk CLI](test-your-kustomize-files-with-snyk-cli.md).
