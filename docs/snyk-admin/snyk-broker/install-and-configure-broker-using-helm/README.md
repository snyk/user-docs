# Install and configure Broker using Helm

Installing the Snyk Broker with the [Broker Helm Chart](https://github.com/snyk/snyk-broker-helm) is the easiest way to deploy Snyk Broker if you are using Kubernetes.

{% hint style="info" %}
The Helm chart does not manage connectivity, and thus you will be responsible for managing [ingress](ingress-options-with-snyk-broker-helm-installation.md) in the Kubernetes cluster.
{% endhint %}

{% hint style="warning" %}
**Multi-tenant settings for EU and AU**\
When you set up Broker, Code Agent, or both for use in EU or AU Multi-tenant environments, additional environment variables with the specific URLs are required.\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
For the URLs, see [EU and AU account datacenter creation](https://docs.snyk.io/snyk-processes/data-residency-at-snyk#eu-and-au-datacenter-account-creation).
{% endhint %}

## Considerations in using the Helm Chart to install Broker

### Kubernetes secrets and Helm Chart installation

API tokens and passwords use Kubernetes Secrets. You can use existing secrets that are created in the following formats.

#### Broker tokens for Helm Chart installation

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

#### SCM token for Helm Chart installation

```
apiVersion: v1
kind: Secret
metadata:
  name: <ENTER_SCM_TYPE>-token
type: Opaque
data:
  <ENTER_SCM_TYPE>-token-key: <BASE64_ENCODED_SECRET>
```

### Image repository, tag, and Image Pull Secret

You can choose to use your own container registry and tag instead of the public images by customizing the `values.yaml` file to specify your container registry uri and tag.

If your container registry requires an image pull secret, you can specify an image secret.

Note that the Image Pull Secret is NOT created by the chart; rather, the image Pull Secret is expected to be present on your cluster.

### Service accounts for Helm Chart installation

To use an existing service account, add the following parameters to the install command:

```
--set serviceAccount.create=false \
--set serviceAccount.name=<ENTER_EXISTING_SERVICE_ACCOUNT> \
```

## Installation using the Snyk Broker Helm Chart

To use this chart, you must first add the chart by adding the repo:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

Then run the commands for[ each SCM, registry, or Jira](../snyk-broker-client-integration-setups-with-helm.md).

* [GitHub](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#github.com-helm-install)
* [GitHub Enterprise](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#github-enterprise-helm-install)
* [Bitbucket Server/Data Centre](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#bitbucket-helm-install)
* [Gitlab](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#gitlab-helm-install)
* [Azure Repos](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#azure-repos-helm-install)
* [JFrog Artifactory](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#artifactory-respository-helm-install)
* [Jira](https://docs.snyk.io/snyk-admin/snyk-broker/install-broker-for-scms-using-helm#jira-notifications-helm-install)

Additional commands are available to Install the Snyk Broker Code Agent and Container Registry Agent.

* [Snyk Broker Code Agent (needed to enable SAST analysis)](../snyk-broker-code-agent/install-broker-for-code-agent-using-helm.md)
* [Snyk Broker - Container Registry Agent (needed to connect to Container Registries)](../snyk-broker-container-registry-agent/install-broker-for-container-registry-agent-using-helm.md)
