# Snyk runtime sensor

{% hint style="warning" %}
**Release status**

The Snyk Runtime Sensor is available in a Closed Beta state and applies only to Snyk AppRisk Pro.
{% endhint %}

The Snyk Runtime Sensor is a [Kubernetes DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) that watches your deployments on a Kubernetes cluster and sends the collected data to Snyk.

The following [risk factors](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk#risk-factors) are reported from the Snyk Runtime Sensor: [Deployed](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed), and [Loaded package](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package) (only for application packages).

On this page, you can find the following information:

* [Prerequisites for Snyk Runtime Sensor](snyk-runtime-sensor.md#prerequisites-for-snyk-runtime-sensor)
* [Install Snyk Runtime Sensor](snyk-runtime-sensor.md#install-snyk-runtime-sensor)
  * [Using a Helm chart](snyk-runtime-sensor.md#using-a-helm-chart)
  * [Using a Helm chart and the AWS Secrets Manager](snyk-runtime-sensor.md#using-a-helm-chart-and-the-aws-secrets-manager)
  * [On OpenShift](snyk-runtime-sensor.md#on-openshift)
  * [Through the AWS Marketplace as an EKS add-on](snyk-runtime-sensor.md#aws-eks-deployment)
* [Troubleshooting for Snyk Runtime Sensor](snyk-runtime-sensor.md#troubleshooting-for-snyk-runtime-sensor)

## Prerequisites for Snyk Runtime Sensor

Ensure that your environment meets the following technical prerequisites to properly use the Snyk Runtime Sensor:

* Kubernetes supported version - Use Kubernetes v.1.19 or higher.

{% hint style="info" %}
Managed Kubernetes services such as EKS Fargate or GKE Autopilot, are not supported, as the cluster nodes are managed by the cloud provider.
{% endhint %}

* Linux kernel - version 5.8 or higher.
* Privileged access - you need either root or the following Linux capabilities: `BPF`, `PERFMON`, `SYS_RESOURCES`, `DAC_READ_SEARCH`, `SYS_PTRACE`, `NET_ADMIN`
* Cluster nodes must support BTF.
* Language support - Go, Java v8 or higher, .NET v2.0.9 or higher, Node.js v10 or higher, Python 3.6 or higher.
* Network policy - if your cluster does not allow all outgoing traffic, set up the policy to enable outgoing traffic on port 443 for the following hosts:
  * `api.snyk.io` or `api.<<REGION>>.snyk.io` if hosted in a different region.
  * `api.sentry.io`
  * `kubernetes.default.svc.cluster.local`

{% hint style="warning" %}
If you encounter network restrictions, ensure that port 443 is enabled and that the policy is stateful.
{% endhint %}

You also need a token for a [service account](https://docs.snyk.io/snyk-admin/service-accounts). The service account must have one of the following roles:

* Group Admin
* Custom Group Level Role with `AppRisk edit` permission enabled.

## Install Snyk Runtime Sensor

* The Snyk Runtime Sensor DaemonSet must meet the following minimum requirements:
  * `CPU: 100m`
  * `Memory: 512Mi`
* Choose one of the following methods to deploy the Snyk Runtime Sensor:&#x20;
  * [Install the Snyk Runtime Sensor using a Helm chart ](snyk-runtime-sensor.md#using-a-helm-chart)
  * [Install the Snyk Runtime Sensor using a Helm chart and the AWS Secrets Manager ](snyk-runtime-sensor.md#using-a-helm-chart-and-the-aws-secrets-manager)
  * [Install the Snyk Runtime Sensor on OpenShift ](snyk-runtime-sensor.md#on-openshift)
  * [Install the Snyk Runtime Sensor through the AWS Marketplace as an EKS add-on](snyk-runtime-sensor.md#through-the-aws-marketplace-as-an-eks-add-on)

### Using a Helm chart

There is a [Helm chart](https://helm.sh) within this repo in [helm/runtime-sensor](https://github.com/snyk/runtime-sensor), that is hosted through GitHub pages in `https://snyk.github.io/runtime-sensor`.

To install the Snyk runtime sensor using Helm Charts, you can follow these steps:

1. Ensure Helm is installed.
2.  Create the `snyk-runtime-sensor` namespace:

    <pre><code><strong>kubectl create namespace snyk-runtime-sensor
    </strong></code></pre>
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
6.  (Optional) If you want to configure custom resources (CPU/memory) for the runtime sensor image, set the following values as well when running the next step (default values are used here):

    ```
    ...
    --set sensor.resources.requests.memory=512Mi
    --set sensor.resources.requests.cpu=100m
    --set sensor.resources.limits.memory=1024Mi
    --set sensor.resources.limits.cpu=500m
    ...
    ```
7.  Install the Helm chart:

    ```
    helm install my-runtime-sensor \
    --set secretName=<<YOUR_SECRET_NAME>> \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    --set snykAPIBaseURL=<<YOUR_REGIONS_API_URL>> \ # Optional
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```

#### Upgrading to the latest version

1.  Check the name that was given to the sensor

    {% code overflow="wrap" %}
    ```
    helm repo list
    ```
    {% endcode %}
2.  Update the repo (with the name from (1)):

    {% code overflow="wrap" %}
    ```
    helm repo update <<SENSOR_REPO_NAME>>
    ```
    {% endcode %}
3.  Upgrade installation:

    ```
    helm upgrade --install <<SENSOR_REPO_NAME>> \
    --set secretName=<<YOUR_SECRET_NAME>> \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```

### Using a Helm Chart and the AWS Secrets Manager

There is a [Helm chart](https://helm.sh) within this repo in [helm/runtime-sensor](https://github.com/snyk/runtime-sensor), that is hosted through GitHub pages in `https://snyk.github.io/runtime-sensor`.

To install the Snyk runtime sensor using Helm Charts and the AWS Secrets Manager, you can follow these steps:

Prerequisite: Install AWS Provider and CSI Secrets Store in your cluster, as instructed [here](https://github.com/aws/secrets-store-csi-driver-provider-aws).

1. Ensure Helm is installed.
2.  Create the `snyk-runtime-sensor` namespace:

    <pre><code><strong>kubectl create namespace snyk-runtime-sensor
    </strong></code></pre>
3.  Create the Snyk Runtime Sensor Secret containing your service account token under the `snykToken` key in your AWS account, and obtain the resulted ARN:

    ```
    aws secretsmanager create-secret \
    --name snyk-runtime-sensor-secret \
    --secret-string '{"snykToken":"<<YOUR_SERVICE_ACCOUNT_TOKEN>>"}'
    ```
4.  Create an access policy for the newly created secret:

    ```
    POLICY_ARN=$(aws --query Policy.Arn --output text iam create-policy --policy-name snyk-runtime-sensor-secret-policy --policy-document '{
        "Version": "2012-10-17",
        "Statement": [ {
            "Effect": "Allow",
            "Action": ["secretsmanager:GetSecretValue", "secretsmanager:DescribeSecret"],
            "Resource": ["<<YOUR_SECRET'S_ARN>>"]
        } ]
    }')
    ```
5.  Create an IAM OIDC provider for your cluster if you haven't done so already (only run this once):

    ```
    eksctl utils associate-iam-oidc-provider \
    --cluster="<<CLUSTER_NAME>>" \
    --approve
    ```
6.  Add the Helm repo:

    ```
    helm repo add runtime-sensor https://snyk.github.io/runtime-sensor
    ```
7. If your data is hosted in a [different region](../../../working-with-snyk/regional-hosting-and-data-residency.md) than the default region (USA), you need to set the `snykAPIBaseURL` while installing the Helm chart in the following format: `api.<<REGION>>.snyk.io:443`, for example `api.eu.snyk.io:443`
8.  (Optional) If you want to configure custom resources (CPU/memory) for the runtime sensor image, set the following values as well when running the next step (default values are used here):

    ```
    ...
    --set sensor.resources.requests.memory=512Mi
    --set sensor.resources.requests.cpu=100m
    --set sensor.resources.limits.memory=1024Mi
    --set sensor.resources.limits.cpu=500m
    ...
    ```
9.  Install the Helm chart:

    ```
    helm install my-runtime-sensor \
    --set secretProvider=aws \
    --set secretName=snyk-runtime-sensor-secret \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    --set snykAPIBaseURL=<<YOUR_REGIONS_API_URL>> \ # Optional
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```
10. Attach the ARN of the policy created in step 4 to the newly created service account, by creating a new role:

    {% code overflow="wrap" %}
    ```
    eksctl create iamserviceaccount \
    --name runtime-sensor \
    --region=<<REGION>> \
    --cluster "<<CLUSTER_NAME>>" \
    --attach-policy-arn "$POLICY_ARN" \
    --approve \
    --override-existing-serviceaccounts \
    --namespace=snyk-runtime-sensor
    ```
    {% endcode %}
11. Verify that the secret was mounted successfully into the `snyk-runtime-sensor` namespace (`kubectl get secrets -n snyk-runtime-sensor`), and that the sensor pods are running successfully (`kubectl get pods -n snyk-runtime-sensor`).

#### Upgrading to the latest version

1.  Check the name that was given to the sensor

    {% code overflow="wrap" %}
    ```
    helm repo list
    ```
    {% endcode %}
2.  Update the repo (with the name from (1)):

    {% code overflow="wrap" %}
    ```
    helm repo update <<SENSOR_REPO_NAME>>
    ```
    {% endcode %}
3.  Upgrade installation:

    ```
    helm upgrade --install <<SENSOR_REPO_NAME>> \
    --set secretProvider=aws \
    --set secretName=snyk-runtime-sensor-secret \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```

### On OpenShift

When running your Kubernetes cluster in OpenShift, you will have to apply the `privileged` Security Context Constraint to the service account of the Snyk Runtime Sensor by running the following command:

```
oc adm policy add-scc-to-user privileged \
system:serviceaccount:<<YOUR_NAMESPACE>>:runtime-sensor
```

Run this command after the sensor is installed as the service account will not be available until the installation is complete.

### Through the AWS Marketplace as an EKS add-on  <a href="#aws-eks-deployment" id="aws-eks-deployment"></a>

Snyk provides a straightforward process for installing the Snyk Runtime Sensor on your AWS EKS cluster. The following steps explain how to integrate this security feature into your environment, enhancing the security of your cluster.

To deploy the Snyk Runtime Sensor on Amazon EKS with EKS Add-on, you need to meet the following prerequisites:

1. Subscribe to Snyk Runtime Sensor on AWS Marketplace [here](https://aws.amazon.com/marketplace/pp/prodview-i23vvrxuamcya).
2. Install the following tools: `kubectl`, `AWS CLI`, and optionally `eksctl`.&#x20;
3. Ensure you have access to the Amazon EKS cluster where you want to install the sensor.&#x20;
4. Ensure you have a Snyk service account token ready with the right permissions, as described in the [prerequisites](snyk-runtime-sensor.md#prerequisites).

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
* `snykAPIBaseURL` - should be configured to be `api.snyk.io:443` , unless your data is hosted in a [different region](../../../working-with-snyk/regional-hosting-and-data-residency.md#what-regions-are-available) than the default (US).

Here is a base configuration to copy:

```
secretName: snyk-secret
clusterName: <<MY_CLUSTER>>
snykGroupId: <<MY_SNYK_GROUP_ID>>
snykAPIBaseURL: api.snyk.io:443
```

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.58.12.png" alt="Set the appropriate configuration values under &#x22;Optional configuraiton settings&#x22;"><figcaption><p>Set the appropriate configuration values under "Optional configuraiton settings"</p></figcaption></figure>

After you select the **Next** and **Create** options you will see the `Add-on snyk-runtimesensor successfully added to cluster <<YOUR_CLUSTER>>` notification on top of the page.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 16.26.13.png" alt="The success message."><figcaption><p>The success message.</p></figcaption></figure>

#### **Enable Snyk Runtime Sensor add-on using AWS CLI**

Run the following command on your workspace to enable the Snyk Runtime Sensor add-on for your Amazon EKS cluster. You have to set the following environment variables to match your Snyk account and your targeted EKS cluster:

* $CLUSTER\_NAME
* $AWS\_REGION
* $SNYK\_GROUP\_ID&#x20;
* $SNYK\_API\_BASE\_URL (should be set to `api.snyk.io:443` unless your account is hosted on a different region than US).

```
aws eks create-addon \
--cluster-name $CLUSTER_NAME \
--region $AWS_REGION \
--addon-name snyk_runtime-sensor \
--configuration-values '{"secretName":"snyk-secret","clusterName":"$CLUSTER_NAME","snykGroupId":"$SNYK_GROUP_ID","snykAPIBaseURL": "$SNYK_API_BASE_URL"}' \
--resolve-conflicts OVERWRITE
```

After you have added the Snyk service account token as described [below](snyk-runtime-sensor.md#add-your-snyk-service-account-token-to-the-eks-cluster), ensure installation has been completed successfully by running the following command:

```
aws eks describe-addon --addon-name snyk_runtime-sensor --cluster-name $CLUSTER_NAME --region $AWS_REGION
```

Ensure the response you get is similar to this one and that the status is ACTIVE - it could take a few minutes until it reaches this status.

```
{
    "addon": {
        "addonName": "snyk_runtime-sensor",
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

#### **Add your Snyk Service Account Token to the EKS cluster**

* Set your `kubectl` context to control your cluster using `aws eks`:&#x20;

```
aws eks update-kubeconfig --name $CLUSTER_NAME --region $AWS_REGION
```

* Create a secret name `snyk-secret` under the `snyk-runtime-sensor` namespace that contains the `snykToken`. The `snykToken` will be your service account token:&#x20;

```
kubectl create secret generic snyk-secret \
--from-literal=snykToken=$SNYK_TOKEN \
-n snyk-runtime-sensor
```

* Data from your AWS EKS Cluster will be reported to Snyk using the Snyk Runtime Sensor.

#### **Disable the Snyk Runtime Sensor add-on**

You can disable the Snyk Runtime Sensor add-on by running the following command:

```
aws eks delete-addon --addon-name snyk_runtime-sensor --cluster-name $CLUSTER_NAME --region $AWS_REGION
```

## Troubleshooting for Snyk Runtime Sensor

If the Snyk Runtime Sensor is not properly reporting the `is_loaded` risk factor, it may be caused by a non-default value of the Linux kernel `perf_event_paranoid` configuration.\
In such cases, install the helm chart with either `--set securityContext.privileged=true` or add `SYS_ADMIN` as a required Linux capability `--set "securityContext.capabilities={SYS_ADMIN}"`.

{% hint style="info" %}
The Loaded package risk factor is not supported by Snyk for operating system packages (such as Debian packages), only for packages which are hosted under package managers such as npm, Maven, or PyPI.
{% endhint %}

Release versions can be found on[ GitHub](https://github.com/snyk/runtime-sensor/releases).
