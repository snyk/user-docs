# Snyk Sysdig integration

To enhance the Snyk workload information detection, we have partnered with Sysdig to enrich the issues detected by Snyk for workloads, with runtime data provided by Sysdig.

## Enabling the Sysdig integration

For the integration with Sysdig to work, the Snyk controller requires an extra Secret in the **`snyk-monitor`** namespace. The Secret name is **`sysdig-eve-secret`**.

{% hint style="info" %}
The commands below should be executed right after installing Sysdig, to allow the Snyk controller to detect Sysdig in the cluster.
{% endhint %}

Please refer to the [Sysdig Secret installation guide](https://docs.sysdig.com/en/docs/sysdig-secure/integrate-effective-vulnerability-exposure-with-snyk/#copy-the-sysdig-secret) to install this Secret. Once the Sysdig Secret is installed, you need to copy it over to the **`snyk-monitor`** namespace:

```
kubectl get secret sysdig-eve-secret -n sysdig-agent -o yaml \
  | grep -v '^\s*namespace:\s' \
  | kubectl apply -n snyk-monitor -f -
```

To enable Snyk to integrate with Sysdig and collect information about packages executed at runtime, use `--set sysdig.enabled=true` when installing the Snyk controller:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
  --namespace snyk-monitor \
  --set clusterName="Production cluster" \
  --set sysdig.enabled=true
```

Your Snyk controller will now collect data from Sysdig every 4 hours. 🎊

## Application support

For application vulnerabilities, we currently provide support for the following languages:

* Java
* JavaScript
* Go

You can find an up-to-date list of supported languages in [Snyk Container: Detecting Application Vulnerabilities in Container Images](https://docs.snyk.io/products/snyk-container/getting-around-the-snyk-container-ui/detecting-application-vulnerabilities-in-container-images).
