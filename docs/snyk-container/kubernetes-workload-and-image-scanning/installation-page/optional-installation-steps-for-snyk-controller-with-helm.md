# Optional installation steps for Snyk Controller with Helm

The installation steps to use depend on how you want to configure Snyk Controller to fit your environment. Follow the relevant steps from the list that follows.

### **Registry with self-signed or additional certificate**

*   If your registry is using self-signed or other additional certificates you must make those available to Snyk monitor. First place the `.crt`, `.cert`, and/or `.key` files in a directory and create a ConfigMap:

    ```
    kubectl create configmap snyk-monitor-certs \
            -n snyk-monitor --from-file=<path_to_certs_folder>
    ```

### **Use the insecure container registry or your registry is using unqualified images**

*   If you are using an insecure registry or your registry is using unqualified images, you can provide a `registries.conf` file.

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

### **Using a proxy for outbound connection to Snyk**

*   If you are using a proxy for the outbound connection to Snyk then you need to configure the integration to use that proxy. To configure the proxy set the following values provided in the Helm chart:

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

### **Alter the logging verbosity**&#x20;

*   If you would like to alter the logging verbosity you can do so as follows. Valid levels are `INFO`, `WARN` and `ERROR`. We default to `INFO`.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set clusterName="Production cluster" \
                 --set log_level="WARN"
    ```

### **Enable Pod Security Policy**

*   By default the controller will run without a [Pod Security Policy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/). However, this can be enabled by passing a setting.

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

### **Configure Snyk Controller with PersistentVolumeClaim (PVC)**

*   You can configure the Snyk controller to use a **PersistentVolumeClaim** (PVC) instead of the default emptyDir storage medium for temporarily pulling images. The PVC can either be created by the Helm template provided by the Snyk chart or you can use an already provisioned PVC. Use the following flags to control the PVC:\
    `pvc.enabled` - instructs the Helm chart to use a PVC instead of an emptyDir\
    `pvc.create` - whether to create the PVC - this is useful when provisioning for the first time \
    `pvc.storageClassName` - controls the StorageClass of the PVC\
    `pvc.name` - the name of the PVC to use in Kubernetes\
    \
    For example, you can run the following command on installation to provision/create the PVC: &#x20;

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                  --namespace snyk-monitor \
                 --set pvc.enabled=true \
                 --set pvc.create=true \
                 --set pvc.name="snyk-monitor-pvc"
    ```

    On subsequent upgrades you can drop the `pvc.create` flag because the PVC already exists:

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \             
                 --namespace snyk-monitor \
                 --set pvc.enabled=true \
                 --set pvc.name="snyk-monitor-pvc"
    ```

### **Configure Snyk Controller to exclude specific namespaces**

*   By default, Snyk purposely ignores scanning certain namespaces which are believed to be internal to Kubernetes (any namespace starting with _**kube-\*,**_ see the full list on [GitHub](https://github.com/snyk/kubernetes-monitor/blob/master/src/supervisor/watchers/internal-namespaces.ts)). You can change the default; Snyk allows configuring the excluded namespaces.\


    When you add your own list of namespaces to exclude with the _excludedNamespaces_ setting, Snyk overrides the default settings and uses the list of namespaces you provide.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \ 
                 --namespace snyk-monitor \
                 --set excludedNamespaces="{kube-node-lease,local-path-storage,some_namespace}"
    ```

### **Configure Snyk Controller resources**

*   If more resources are required in order to deploy the controller, configure the Helm charts default value for requests and limits with the `--set` flag.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set requests."ephemeral-storage"="50Gi"
                 --set limits."ephemeral-storage"="50Gi"
    ```

### Configure Snyk Controller workers count

*   You can adjust the workers number (XX) with the`--set` flag.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set workers.count="XX"
    ```

### Configure Snyk Controller CPU

*   If more or less CPU is required in order to deploy the controller, configure the Helm charts default value for requests (X) and limits (Y) with the `--set` flag.

    ```
    helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
                 --namespace snyk-monitor \
                 --set requests.cpu="X"
                 --set limits.cpu="Y"
    ```
