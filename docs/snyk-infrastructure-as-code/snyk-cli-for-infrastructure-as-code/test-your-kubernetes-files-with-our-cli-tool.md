# Test your Kubernetes files with our CLI tool

With Snyk Infrastructure as Code, you can test your configuration files directly from the CLI.

Snyk Infrastructure as Code for Kubernetes supports:

* Deployments, Pods and Services.
* CronJobs, Jobs, StatefulSet, ReplicaSet, DaemonSet, and ReplicationController.

You can use the CLI as follows:

## To test for an issue on specified files:

```text
snyk iac test
```

For example, from the CLI enter the following:

```text
snyk iac test deploy.yaml
```

You can also specify multiple files by appending the file names after each other, such as:

```text
snyk iac test file-1.yaml file-2.yaml
```

## To scan a Helm chart using the CLI

You scan a Helm chart by converting the template to a rendered Kuberenetes manifest file and then scanning this using the Snyk IaC CLI.

```text
helm template ./iac-helm > helm.yaml
snyk iac test helm.yaml
```

change \`iac-helm\` for your Helm chart name.

## To scan a Kustomize template using the CLI

You scan a Kustomize template by building the Kuberenetes manifest file and then scanning this using the Snyk IaC CLI.

```text
kustomize build > kubernetes.yaml
snyk iac test kubernetes.yaml
```

Depending on your kustomize templates, you may need to provide a name after the build argument.



