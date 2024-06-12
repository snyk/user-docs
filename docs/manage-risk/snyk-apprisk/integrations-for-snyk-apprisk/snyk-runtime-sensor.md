# Snyk runtime sensor

{% hint style="warning" %}
**Release status**

The Snyk runtime sensor is available in a Closed Beta state and is applicable only to the Snyk AppRisk Pro version. &#x20;

Contact your salesperson if you are interested in Snyk AppRisk Pro.
{% endhint %}

The Runtime Sensor watches your deployments on a Kubernetes cluster and sends the collected data to Snyk.

## Prerequisites

Ensure that your environment meets the following technical prerequisites to properly use the Snyk Runtime Sensor:

* Kubernetes supported version - Use Kubernetes v.1.19 or higher.

{% hint style="info" %}
Managed Kubernetes services such as EKS Fargate or GKE Autopilot, are not supported, as the cluster nodes are managed by the cloud provider.
{% endhint %}

* Privileged access - you need either root or the following Linux capabilities: `BPF`, `PERFMON`, `SYS_RESOURCES`, `DAC_READ_SEARCH`, `SYS_PTRACE`, `NET_ADMIN`
* Cluster nodes must support BTF.
* Language support - Go, Java v8 or higher, .NET v2.0.9 or higher, Node.js v10 or higher, Python 3.6 or higher.

You also need a token for a [service account](https://docs.snyk.io/snyk-admin/service-accounts). The service account must have one of the following roles:

* Group Admin
* Custom Group Level Role with `AppRisk edit` permission enabled.

## Installation

There is a [Helm chart](https://helm.sh) within this repo in [helm/runtime-sensor](https://github.com/snyk/runtime-sensor), that is hosted through GitHub pages in `https://snyk.github.io/runtime-sensor`.

To install the Snyk runtime sensor using Helm Charts, you can follow these steps:

1. Ensure Helm is installed.
2.  Create the `snyk-runtime-sensor` namespace:

    ```
    kubectl create namespace snyk-runtime-sensor
    ```
3.  Create a secret with your service account token, which has the appropriate permissions (as instructed in the prerequisites section) under the created namespace:

    {% code overflow="wrap" %}
    ```
    kubectl create secret generic <<YOUR_SECRET_NAME>> --from-literal=snykToken=<<YOUR_TOKEN>> -n snyk-runtime-sensor
    ```
    {% endcode %}
4.  Add the Helm repo:

    ```
    helm repo add runtime-sensor https://snyk.github.io/runtime-sensor
    ```
5. If your data is hosted in a [different region](../../../working-with-snyk/regional-hosting-and-data-residency.md) than the default region (USA), you need to set the `snykAPIBaseURL` while installing the Helm chart in the following format: `api.<<REGION>>.snyk.io:443`, for example `api.eu.snyk.io:443`
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

## OpenShift

When running your Kubernetes cluster in OpenShift, you will have to apply the `privileged` Security Context Constraint to the Runtime Sensor's service account by running the following command:

```
oc adm policy add-scc-to-user privileged \
system:serviceaccount:<<YOUR_NAMESPACE>>:runtime-sensor
```

This command must be run after the sensor has been installed, as the service account will not be available prior to the installation.&#x20;

## AWS EKS Deployment

Snyk provides a straightforward process for installing the Snyk Runtime Sensor on your AWS EKS cluster. The following steps explain how to integrate this security feature into your environment, enhancing the security of your cluster.

### Deploy the Snyk Runtime Sensor on an Amazon EKS cluster using Amazon EKS add-on Prerequisites

1. Subscribe to Snyk Runtime Sensor on AWS Marketplace [here](https://aws.amazon.com/marketplace/pp/prodview-i23vvrxuamcya).
2. Install the following tools: `kubectl`, `AWS CLI`, and optionally `eksctl`.&#x20;
3. Ensure you have access to the Amazon EKS cluster where you want to install the sensor.&#x20;
4. Ensure you have a Snyk service account token ready with the right permissions.

### Install the add-on using AWS console or AWS CLI

#### **Enable the Snyk Runtime Sensor add-on from AWS console**

After you have successfully set up a subscription to Snyk Runtime Sensor on AWS Marketplace and followed the on-screen instructions, you will be redirected to Amazon EKS console.&#x20;

To enable the Snyk Runtime Sensor for your Amazon EKS cluster, select your cluster on the Amazon EKS console. Then, navigate to the Add-ons tab and choose "Get more add-ons". Use the search bar to find "runtime" and follow the on-screen instructions to enable the add-on for your cluster.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.53.23.png" alt="Select the Snyk Runtime Sensor add-on"><figcaption><p>Select the Snyk Runtime Sensor add-on</p></figcaption></figure>

On the next screen, select the latest version (even if already selected) and open the "Optional configuration settings".

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.55.19.png" alt=""><figcaption><p>Select the latest available version and open the 'Optional configuration settings'</p></figcaption></figure>

Under the "configuration values", set the following attributes in a YAML or JSON format:

* `secretName` - the secret name that will be created later in the process. The default value is  `snyk-secret` .
* `clusterName` - the name of the cluster where the add-on is installed.
* `snykGroupId` - the Group ID associated with the used service account.
* `snykAPIBaseURL` - should be configured to be `api.snyk.io:443` unless your data is hosted in a [different region](../../../working-with-snyk/regional-hosting-and-data-residency.md) than the default (USA).&#x20;

Here is a base configuration to copy:

```
secretName: snyk-secret
clusterName: <<MY_CLUSTER>>
snykGroupId: <<MY_SNYK_GROUP_ID>>
snykAPIBaseURL: api.snyk.io:443
```

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.58.12.png" alt="Set the appropriate configuration values under &#x22;Optional configuraiton settings&#x22;"><figcaption><p>Set the appropriate configuration values under "Optional configuraiton settings"</p></figcaption></figure>

After you select the **Next** and **Create** options you will see the "Add-on snyk-runtimesensor successfully added to cluster <\<YOUR\_CLUSTER>>" notification on top of the page.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 16.26.13.png" alt="The success message."><figcaption><p>The success message.</p></figcaption></figure>

#### **Enable Snyk Runtime Sensor add-on using AWS CLI**

Run the following command on your workspace to enable the Snyk Runtime Sensor add-on for your Amazon EKS cluster. You have to set the following parameters in your targeted EKS cluster:

* $CLUSTER\_NAME,&#x20;
* $AWS\_REGION,&#x20;
* $SNYK\_GROUP\_ID&#x20;
* Snyk Group ID

```
aws eks create-addon \
--cluster-name $CLUSTER_NAME \
--region $AWS_REGION \
--addon-name snyk_runtime-sensor \
--configuration-values '{"secretName":"snyk-secret","clusterName":"$CLUSTER_NAME","snykGroupId":"$SNYK_GROUP_ID","snykAPIBaseURL": "api.snyk.io:443"}' \
--resolve-conflicts OVERWRITE
```

After you have added the Snyk Service Account Token as described below, ensure installation has been completed successfully by running the following command:

```
aws eks describe-addon --addon-name snyk_runtime-sensor --cluster-name $CLUSTER_NAME --region $AWS_REGION
```

Ensure the response you get is similar to this one and that the status is ACTIVE.

```
{
    "addon": {
        "addonName": "snyk-runtimesensor",
        "clusterName": "<<YOUR_CLUSTER>>",
        "status": "ACTIVE",
        "addonVersion": "v1.17.2-eksbuild.1",
        "health": {
            "issues": []
        },
        "addonArn": "arn:aws:eks:us-east-1:XXXX:addon/<<YOUR_CLUSTER>>/snyk-runtimesensor/ffffffff-ffff-ffff-ffff-ffffffffffff",
        "createdAt": "2024-05-26T16:43:15.551000+03:00",
        "modifiedAt": "2024-05-26T16:43:28.714000+03:00",
        "tags": {},
        "configurationValues": "{\"secretName\":\"snyk-secret\",\"clusterName\":\"YOUR_CLUSTER\",\"snykGroupId\":\"YOUR_GROUP_ID\",\"snykAPIBaseURL\": \"api.snyk.io:443\"}"
    }
}
```

**Add your Snyk Service Account Token to the EKS cluster**

* Set your `kubectl` context to control your cluster using `aws eks`:&#x20;

```
aws eks update-kubeconfig --name $CLUSTER_NAME --region $AWS_REGION
```

* Create a secret name `snyk-secret` under the `snyk-runtime-sensor` namespace that contains the `snykToken` . The `snykToken` will be your service account token:&#x20;

```
kubectl create secret generic snyk-secret \
--from-literal=snykToken=$SNYK_TOKEN \
-n snyk-runtime-sensor
```

* Now, data from your AWS EKS Cluster will be reported to Snyk using the Snyk Runtime Sensor.

#### **Disable the Snyk Runtime Sensor add-on**

You can disable the Snyk Runtime Sensor add-on by running the following command:

```
aws eks delete-addon --addon-name snyk-runtimesensor --cluster-name $CLUSTER_NAME --region $AWS_REGION
```

## Troubleshooting

* If the Snyk Runtime Sensor is not properly reporting the `is_loaded` risk factor, it may be caused by a non-default value of the Linux kernel `perf_event_paranoid` configuration.\
  In such cases, install the helm chart with either `--set securityContext.privileged=true` or add `SYS_ADMIN` as a required Linux capability `--set "securityContext.capabilities={SYS_ADMIN}"`.



Release versions can be found on[ GitHub](https://github.com/snyk/runtime-sensor/releases).
