# Install the Snyk controller with Helm

To get vulnerability details about your Kubernetes workloads, a Snyk admin must first install the Snyk controller onto your cluster. The Snyk controller is published in [Helm Hub](https://hub.helm.sh/charts/snyk/snyk-monitor).

This section covers:

* Snyk integration for most Kubernetes platforms
* Additional configuration steps for integration when using Amazon Elastic Container Registry (ECR) with your Amazon Elastic Kubernetes Service (EKS) clusters

**Prerequisites**

{% hint style="info" %}
**Feature availability**\
This feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

* An administrator account for your Snyk organization.
* A minimum of 50 GB of storage must be available in the form of an [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) on the cluster.
* Your Kubernetes cluster needs to be able to communicate with Snyk outbound over HTTPS.
* When configuring Snyk to integrate with an Amazon Elastic Kubernetes Services (EKS) cluster, if you wish to scan images hosted on your Amazon Elastic Container Registry (ECR), you need to first follow the prerequisites outlined in the [AWS documentation](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR\_on\_EKS.html).

**Steps**

1.  Access your Kubernetes environment and run the following command in order to add the Snyk Charts repository to Helm:

    ```
    helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor
    ```
2.  Once the repository is added, create a unique namespace for the Snyk controller:

    ```
    kubectl create namespace snyk-monitor
    ```

    **Tip:** Use a unique namespace to isolate the controller resources more easily. This is generally good practice for Kubernetes applications. Notice our namespace is called snyk-monitor, you’ll need this later when configuring other resources.
3. Now, log in to your Snyk account and navigate to **Integrations**.
4. Search for and click **Kubernetes**.
5. Click **Connect** from the page that loads, copy the **Integration ID.** The Snyk **Integration ID** is a UUID, similar to this format: `abcd1234-abcd-1234-abcd-1234abcd1234`. Save it for use from your Kubernetes environment in the next step.
6.  Snyk monitor runs by using your Snyk **Integration ID**, and using a `dockercfg` file. If you are not using any private registries, create a Kubernetes secret called `snyk-monitor` containing the Snyk **Integration ID** from the previous step by running the following command:

    ```
    kubectl create secret generic snyk-monitor -n snyk-monitor \
            --from-literal=dockercfg.json={} \
            --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234
    ```

    **Note:** The secret must be called `snyk-monitor` in order for the integration to work.
7. If any of the images you need to scan are located in private registries, you need to provide credentials to access those registries by creating a secret (which must be called snyk-monitor) using both the Snyk Integration ID as well as a `dockercfg.json` file. The `dockercfg.json` file is necessary to allow the monitor to look up images in private registries. Usually, your credentials reside in `$HOME/.docker/config.json`. These credential must also be added to the `dockerconfig.json` file.
   1.  &#x20;Create a file named `dockercfg.json`. Store your credentials in there; it should look like this:

       ```
       {
         // If your cluster does not run on GKE or it runs on GKE and pulls images from other private registries, add the following:
         "auths": {
           "gcr.io": {
             "auth": "BASE64-ENCODED-AUTH-DETAILS"
           }
           // Add other registries as necessary
         },
         
         // If your cluster runs on GKE and you are using GCR, add the following:
         "credHelpers": {
           "us.gcr.io": "gcloud",
           "asia.gcr.io": "gcloud",
           "marketplace.gcr.io": "gcloud",
           "gcr.io": "gcloud",
           "eu.gcr.io": "gcloud",
           "staging-k8s.gcr.io": "gcloud"
         }
         
         // If your cluster runs on EKS and you are using ECR, add the following:
         {
           "credsStore": "ecr-login"
         }
         
         With Docker 1.13.0 or greater, you can configure Docker to use different credential helpers for different registries.
         To use this credential helper for a specific ECR registry, create a credHelpers section with the URI of your ECR registry:
         {
           "credHelpers": {
             "public.ecr.aws": "ecr-login",
       	"<aws_account_id>.dkr.ecr.<region>.amazonaws.com": "ecr-login"
           }
         }
       }
       ```

       2\. Create a secret with the file added:

       ```
       kubectl create secret generic snyk-monitor \
               -n snyk-monitor --from-file=dockercfg.json \
               --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234
       ```
8.  If your registry is using self-signed or other additional certificates you must make those available to Snyk monitor. First place the `.crt`, `.cert`, and/or `.key` files in a directory and create a ConfigMap:

    ```
    kubectl create configmap snyk-monitor-certs \
            -n snyk-monitor --from-file=<path_to_certs_folder>
    ```
9.  If you are using an insecure registry or your registry is using unqualified images, you can provide a `registries.conf` file.

    ```
    [[registry]]
    location = "internal-registry-for-example.net/bar"
    insecure = true
    ```

    See the [documentation](https://github.com/containers/image/blob/master/docs/containers-registries.conf.5.md) for information on the format and further examples. Once you've created the file, you can use it to create the following ConfigMap:

    ```
    kubectl create configmap snyk-monitor-registries-conf \
            -n snyk-monitor \
            --from-file=<path_to_registries_conf_file>
    ```
10. Install the Snyk Helm chart:

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set clusterName="Production cluster"
    ```

    If you are running your own instance of Snyk you need to specify the API endpoint when installing the controller. Replace below with the full hostname of your Snyk instance.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set clusterName="Production cluster" \
                 --set integrationApi=https://<server>/kubernetes-upstream
    ```

    **Tip**: Replace the name **Production cluster** with a name based on the cluster you are monitoring. You’ll use this label to find workloads in Snyk later. Please note that `/` in cluster name is disallowed. Any `/` in cluster names will be removed.\
    Also, to avoid naming the cluster on every update, you can use Helm's existing option for **--reuse-values.** This means that when upgrading, it'll reuse the last release's values and merge in any overrides from the command line via **--set** and -f. If '**--reset-values**' is specified, this is ignored.
11. If you are using a proxy for the outbound connection to Snyk then you need to configure the integration to use that proxy. To configure the proxy set the following values provided in the Helm chart:

    * `http_proxy`
    * `https_proxy`
    * `no_proxy`

    For instance:

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set clusterName="Production cluster" \
                 --set https_proxy=http://192.168.99.100:8080
    ```

    Note that the integration does not support CIDR address ranges or wildcards in the `no_proxy` value. Only exact matches are supported.
12. If you would like to alter the logging verbosity you can do so as follows. Valid levels are `INFO`, `WARN` and `ERROR`. We default to `INFO`.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set clusterName="Production cluster" \
                 --set log_level="WARN"
    ```
13. By default the controller will run without a [Pod Security Policy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/). However, this can be enabled by passing a setting.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set clusterName="Production cluster" \
                 --set psp.enabled=true
    ```

    You can reuse an existing Pod Security Policy by specifying the name. If you don't specify a name then a new policy will be automatically created.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set clusterName="Production cluster" \
                 --set psp.enabled=true \
                 --set psp.name=something
    ```
14. You can configure the Snyk controller to use a **PersistentVolumeClaim** (PVC) instead of the default emptyDir storage medium for temporarily pulling images. The PVC can either be created by the Helm template provided by the Snyk chart or you can use an already provisioned PVC.

    Use the following flags to control the PVC:

    * pvc.enabled - instructs the Helm chart to use a PVC instead of an emptyDir
    * pvc.create - whether to create the PVC - this is useful when provisioning for the first time
    * pvc.storageClassName - controls the StorageClass of the PVC
    * pvc.name - the name of the PVC to use in Kubernetes

    For example, you can run the following command on installation to provision/create the PVC:

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set pvc.enabled=true \
                 --set pvc.create=true \
                 --set pvc.name="snyk-monitor-pvc"
    ```

    On subsequent upgrades you can drop the "pvc.create" flag because the PVC already exists:

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set pvc.enabled=true \
                 --set pvc.name="snyk-monitor-pvc"
    ```
15. By default, we purposely ignore scanning certain namespaces which we believe are internal to Kubernetes (any namespace starting with _**kube-\*,**_ full list can be found [here](https://github.com/snyk/kubernetes-monitor/blob/master/src/supervisor/watchers/internal-namespaces.ts)). If you wish to change that, we allow configuring the excluded namespaces.\
    By adding your own list of namespaces to exclude using _excludedNamespaces_ setting, we will override our default settings and use the list of namespaces you provide.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set excludedNamespaces="{kube-node-lease,local-path-storage,some_namespace}"
    ```
16. If more resources are required in order to deploy the controller, configure the helm charts default value for requests and limits with the `--set` flag.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set requests."ephemeral-storage"="50Gi"
                 --set limits."ephemeral-storage"="50Gi"
    ```







![](../../../../.gitbook/assets/uuid-26f9c2cd-2755-07d5-61a0-bdb0261d87ab-en.gif)
