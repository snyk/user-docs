# Install and configure Broker using Helm

Installing the Snyk Broker with the [Broker Helm Chart](https://github.com/snyk/snyk-broker-helm) is the easiest way to deploy Snyk Broker if you are using Kubernetes.

{% hint style="info" %}
The Helm chart does not manage connectivity and thus you will be responsible for managing [ingress](ingress-options-with-snyk-broker-helm-installation.md) in the Kubernetes cluster.
{% endhint %}

## Kubernetes secrets and Helm Chart installation

API tokens and passwords use Kubernetes Secrets. You can use existing secrets that are created in the following formats.

### Broker tokens for Helm Chart installation

```
apiVersion: v1
kind: Secret
metadata:
  name: <ENTER_SCM_TYPE>-broker-token
  labels:
    app: <ENTER_SCM_TYPE>-broker-token
type: Opaque
data:
  <ENTER_SCM_TYPE>-broker-token-key: <BASE64_ENCODED_SECRET>
```

### SCM token for Helm Chart installation

```
apiVersion: v1
kind: Secret
metadata:
  name: <ENTER_SCM_TYPE>-token
type: Opaque
data:
  <ENTER_SCM_TYPE>-token-key: <BASE64_ENCODED_SECRET>
```

## Service accounts for Helm Chart installation

To use an existing service account, add the following parameters to the install command:

```
--set serviceAccount.create=false \
--set serviceAccount.name=<ENTER_EXISTING_SERVICE_ACCOUNT> \
```

## Installation using the Snyk Broker Helm Chart

To use this chart you must first add the chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

Then run the [commands for each SCM ](../install-broker-for-scms-using-helm.md)or the Snyk Broker - Code Agent or Snyk Broker - Container Registry Agent.

* [GitHub](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#github.com-helm-install)
* [GitHub Enterprise](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#github-enterprise-helm-install)
* [Bitbucket Server/Data Centre](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#bitbucket-helm-install)
* [Gitlab](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#gitlab-helm-install)
* [Azure Repos](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#azure-repos-helm-install)
* [JFrog Artifactory](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#artifactory-respository-helm-install)
* [Jira](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#jira-notifications-helm-install)
* [Snyk Broker - Container Registry Agent (needed to connect to Container Registries)](../install-broker-for-container-registry-agent-using-helm.md)
* [Snyk Broker Code Agent (needed to enable SAST analysis)](../install-broker-for-code-agent-using-helm.md)
