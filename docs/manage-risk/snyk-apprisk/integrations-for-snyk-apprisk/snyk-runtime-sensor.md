# Snyk runtime sensor

The Runtime Sensor watches your deployments on a Kubernetes cluster and sends the collected data to Snyk.

## Installation

There is a [Helm chart](https://helm.sh) within this repo in [helm/runtime-sensor](https://github.com/snyk/runtime-sensor), that is hosted through GitHub pages in `https://snyk.github.io/runtime-sensor`.

Create a Kubernetes secret that contains the API token for the [service account](https://docs.snyk.io/snyk-admin/service-accounts). The service account must have one of the following roles:

* Group Admin
* Custom Group Level Role with `AppRisk edit` permission enabled.

To install the Snyk runtime sensor using Helm Charts, you can follow these steps:

1. Ensure Helm is installed
2.  Create the `snyk-runtime-sensor` namespace:

    ```
    kubectl create namespace snyk-runtime-sensor
    ```
3.  Create a secret with your service account token, which has the appropriate permissions under the created namespace:

    {% code overflow="wrap" %}
    ```
    kubectl create secret generic <<YOUR_SECRET_NAME>> --from-literal=snykToken=<<YOUR_TOKEN>> -n snyk-runtime-sensor
    ```
    {% endcode %}
4.  Add the Helm repo:

    ```
    helm repo add runtime-sensor https://snyk.github.io/runtime-sensor
    ```
5.  Install the Helm chart:

    ```
    helm install my-runtime-sensor \
    --set secretName=<<YOUR_SECRET_NAME>> \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```

Release versions can be found on[ GitHub](https://github.com/snyk/runtime-sensor/releases).



