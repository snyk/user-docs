# Optional installation steps for the Snyk Controller with Helm

The installation steps depend on how you want to configure the Snyk Controller to fit your environment. Follow the applicable steps that match your situation.

## **Use a registry with a self-signed or additional certificate**

If your registry is using self-signed or other additional certificates, you must make them available to Snyk monitor. To do this, first place the `.crt`, `.cert`, and `.key` files in a directory and create a ConfigMap:

```
kubectl create configmap snyk-monitor-certs \
        -n snyk-monitor --from-file=<path_to_certs_folder>
```

## **Use the insecure container registry or a registry that uses unqualified images**

If you are using an insecure registry or your registry is using unqualified images, you can provide a `registries.conf` file.

```
[[registry]]
location = "internal-registry-for-example.net/bar"
insecure = true
```

For information on the format and more examples, see the [GitHub container image documentation](https://github.com/containers/image/blob/master/docs/containers-registries.conf.5.md). After you create the file, you can use it to create the following ConfigMap:

```
kubectl create configmap snyk-monitor-registries-conf \
        -n snyk-monitor \
        --from-file=<path_to_registries_conf_file>
```

## **Use a proxy for outbound connection to Snyk**

If you are using a proxy for the outbound connection to Snyk, you must configure the integration to use that proxy. To configure the proxy, set the following values provided in the Helm Chart:

* `http_proxy`
* `https_proxy`
* `no_proxy`

For example:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster" \
             --set https_proxy=http://192.168.99.100:8080
```

{% hint style="info" %}
The integration does not support CIDR address ranges or wildcards in the `no_proxy` value. Only exact matches are supported.
{% endhint %}

## **Change the logging verbosity**

You can change the logging verbosity. Valid levels are `INFO`, `WARN,` `ERROR` and `DEBUG`. Snyk defaults to `INFO`. To change the logging verbosity, use the following command:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster" \
             --set log_level="WARN"
```

## **Enable a Pod Security Policy**

By default, the controller runs without a [Pod Security Policy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/). You can enable it by passing a setting:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster" \
             --set psp.enabled=true
```

You can reuse an existing Pod Security Policy by specifying the name. If you do not specify a name, a new policy is created automatically.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster" \
             --set psp.enabled=true \
             --set psp.name=something
```

## **Configure the Snyk Controller with PersistentVolumeClaim (PVC)**

You can configure the Snyk Controller to use a PersistentVolumeClaim (PVC) instead of the default emptyDir storage medium to temporarily pull images.

To create the PVC, you can use the Helm template provided by the Snyk chart or an already provisioned PVC.

To control the PVC, use the following flags:

* `pvc.enabled` - instructs the Helm Chart to use a PVC instead of an emptyDir.
* `pvc.create` - creates the PVC. This is useful when provisioning for the first time.
* `pvc.storageClassName` - controls the StorageClass of the PVC.
* `pvc.name` - the name of the PVC to use with Kubernetes.

For example, you can run the following command on installation to provision and create the PVC:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set pvc.enabled=true \
             --set pvc.create=true \
             --set pvc.name="snyk-monitor-pvc"
```

On subsequent upgrades, you can remove the `pvc.create` flag because the PVC already exists:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \             
             --namespace snyk-monitor \
             --set pvc.enabled=true \
             --set pvc.name="snyk-monitor-pvc"
```

## **Configure the Snyk Controller to exclude specific namespaces**

By default, Snyk ignores scanning certain namespaces believed to be internal to Kubernetes. For the full list, see [Configuring excluded namespaces](https://github.com/snyk/kubernetes-monitor/tree/master/snyk-monitor#configuring-excluded-namespaces).\
\
You can change the default, as Snyk allows you to configure the excluded namespaces.

When you add your own list of namespaces to exclude with the `excludedNamespaces` setting, Snyk overrides the default settings and uses the list of namespaces you provide.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \ 
             --namespace snyk-monitor \
             --set excludedNamespaces="{kube-node-lease,local-path-storage,some_namespace}"
```

## **Configure Snyk Controller resources**

If more resources are required to deploy the controller, configure the Helm Chart default value for requests and limits with the `--set` flag.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set requests."ephemeral-storage"="50Gi"
             --set limits."ephemeral-storage"="50Gi"
```

## Configure Snyk Controller workers count

You can adjust the workers number (XX) with the`--set` flag.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set workers.count="XX"
```

## Configure Snyk Controller CPU

If more or less CPU is required in order to deploy the controller, configure the HelmChart default value for requests (X) and limits (Y) with the `--set` flag.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set requests.cpu="X"
             --set limits.cpu="Y"
```

## Configure Snyk Controller initContainers setting

If you use a **PersistentVolumeClaim** (PVC), instead of the default **emptyDir** storage, then you must provision volume permissions. To do this, disable the [initContainer](https://github.com/snyk/kubernetes-monitor/blob/master/snyk-monitor/templates/deployment.yaml#L56)[s](https://github.com/snyk/kubernetes-monitor/blob/master/snyk-monitor/templates/deployment.yaml#L56) in the Helm Chart using the `--set` flag:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set initContainers.enabled = false
```

{% hint style="info" %}
For Openshift platform, the `initContainers` setting is mandatory.
{% endhint %}
