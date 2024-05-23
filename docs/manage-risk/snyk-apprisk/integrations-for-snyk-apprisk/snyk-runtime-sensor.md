# Snyk runtime sensor

{% hint style="warning" %}
**Release status**

The Snyk runtime sensor is available in a Closed Beta state and is applicable only to the Snyk AppRisk Pro version. &#x20;

Contact your salesperson if you are interested in Snyk AppRisk Pro.
{% endhint %}

The Runtime Sensor watches your deployments on a Kubernetes cluster and sends the collected data to Snyk.

## Prerequisites

To ensure proper use of the Snyk Runtime Sensor, ensure that your environment meets the following technical prerequisites:

* Kubernetes supported version - Use Kubernetes v.1.19 or higher.

{% hint style="info" %}
Managed Kubernetes are not supported, for example, EKS Fargate or GKE Autopilot.
{% endhint %}

* Privileged access - you need either root or the following Linux capabilities: `BPF`, `PERFMON`, `SYS_RESOURCES`, `DAC_READ_SEARCH`, `SYS_PTRACE`, `NET_ADMIN`
* Cluster nodes must support BTF
* Language support - Go, Java v8 or higher, .NET v2.0.9 of higher, Node.js v10 or higher, Python 3.6 or higher

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
5. If your data is hosted in a [different region](../../../working-with-snyk/regional-hosting-and-data-residency.md) than the default (USA), you also need to set the `snykAPIBaseURL` while installing the Helm chart in the following format: `api.<<REGION>>.snyk.io:443`, for example `api.eu.snyk.io:443`
6.  Install the Helm chart:

    ```
    helm install my-runtime-sensor \
    --set secretName=<<YOUR_SECRET_NAME>> \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    --set snykAPIBaseURL=<<YOUR_REGIONS_API_URL>> \ # Optional
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor

    ```

## Troubleshooting

* In case the `is_loaded` risk factor is not properly reported by the sensor, it may be caused by a non-default value of the Linux kernel `perf_event_paranoid` configuration.\
  In such cases, install the helm chart with either `--set securityContext.privileged=true` or add `SYS_ADMIN` as a required Linux capability `--set "securityContext.capabilities={SYS_ADMIN}"`.

Release versions can be found on[ GitHub](https://github.com/snyk/runtime-sensor/releases).
