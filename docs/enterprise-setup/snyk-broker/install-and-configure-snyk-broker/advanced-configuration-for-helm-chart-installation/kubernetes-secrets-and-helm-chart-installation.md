# Kubernetes secrets and Helm Chart installation

Beginning with version `2.8.0` of the Snyk Broker Helm Chart, external secrets are supported.&#x20;

To enable this functionality, set `useExternalSecrets` to `true` in `values.yaml` or  `--set externalSecrets=true`.&#x20;

To obtain a list of required secrets, perform a dry run of a Helm installation. This will not make any changes to your Kubernetes environment, but does require the following:

```
helm install snyk-broker-chart \
  snyk-broker/snyk-broker \
  --set externalSecrets=true \
  --set scmType=<your-scm-type> \
  --dry-run=client
```

A list of secrets with their expected names and values will be generated. The following example uses `scmType=nexus` :

```
### Secret Creation Disabled ###

Ensure secrets are present on your cluster in the default namespace:

-> NAME:KEY <VALUE>
-> nexus-broker-token-snyk-broker-chart:nexus-broker-token-key <your-broker-token>
-> nexus-base-nexus-url-snyk-broker-chart:nexus-base-nexus-url <BASE_NEXUS_URL>
-> nexus-nexus-url-snyk-broker-chart:nexus-nexus-url <NEXUS_URL>
-> nexus-broker-client-validation-url-snyk-broker-chart:nexus-broker-client-validation-url <BROKER_CLIENT_VALIDATION_URL>
```

In this example, four secrets must exist within the same namespace to which the Broker will be installed, each containing one key-value pair. Any values that are shown in `<>` characters are indicators to add your own secret data.

## Renaming secrets and keys

Each of the following Helm values supports `name` and `key`, to allow the Snyk Broker Helm Chart to reference a specific secret name and key within that secret:

* `externalCredentialSecret` (used for any Broker type that is not `artifactory`, `nexus` or `nexus2` for the required password or PAT associated with the Broker type)
* `brokerTokenSecret` (used for your Broker token)
* `scmTokenPoolSecret` (used if [Credential Pooling](../advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker-and-helm.md) is enabled)
* `artifactoryUrlSecret` (required for `artifactory` only)
* `baseNexusUrlSecret` (required for `nexus` and `nexus2` only)
* `nexusUrlSecret` (required for `nexus` and `nexus2` only)
* `brokerClientValidationUrlSecret` (required for `nexus` and `nexus2` only, optionally set for `artifactory`)

For example, if your Kubernetes cluster has a secret with a Broker token in the following form:

```
apiVersion: v1
kind: Secret
metadata:
  name: snyk-broker-secrets
type: Opaque
data:
  org-x-broker-token: <broker-token-here>
```

Set the following:

```
useExternalSecrets: true

brokerTokenSecret:
  name: snyk-broker-secrets
  key: org-x-broker-token
```

The Helm Chart will reference the contents under the `org-x-broker-token` key in Secret `snyk-broker-secrets` for the Broker token.

## Partial external secrets

When `useExternalSecrets` is true, the Broker Helm Chart will check whether a value is provided for a secret (for example, `brokerToken=<your-broker-token>`)

* If a value exists, create a secret as usual.
* If no value exists, look for an external secret.

By this means, some secrets may be controlled by the Broker Helm chart, and others controlled externally:

```
scmType: github-com
brokerToken: <my-broker-token>
useExternalSecrets: true
githubToken: ""
```

This set of values will:

* Create a secret for the provided Broker token
* Reference an external secret for the required GitHub token

Performing a dry run of a Helm installation will provide the required secret names and keys:

```
### Secret Creation Disabled ###

Ensure secrets are present on your cluster in the default namespace:

-> NAME:KEY <VALUE>
-> github-com-token-snyk-broker-chart:github-com-token-key <GITHUB_TOKEN>
```

Note the Broker token secret is excluded from this list as a value is directly provided to the Broker Helm Chart.

## Using a single external secret with multiple keys

A single Kubernetes secret may contain all required credentials for the Snyk Broker to operate. Using a Broker of type `nexus` as an example, assume this secret is present in Kubernetes:

```
apiVersion: v1
kind: Secret
metadata:
  name: snyk-broker-secrets
type: Opaque
data:
  nexus-broker-token-key: <broker-token-here>
  nexus-nexus-url: https://user:pass@nexus.tld/myrepository
  nexus-base-nexus-url: https://user:pass@nexus.tld
  nexus-broker-client-validation-url: https://user:pass@nexus.tld/service/rest/v1/status/check
```

To specify this secret for all required values for `scmType=nexus`, set:

```
brokerTokenSecret:
  name: snyk-broker-secrets
  
nexusUrlSecret:
  name: snyk-broker-secrets
  
baseNexusUrlSecret:
  name: snyk-broker-secrets
    
brokerClientValidationUrlSecret:
  name: snyk-broker-secrets
```
