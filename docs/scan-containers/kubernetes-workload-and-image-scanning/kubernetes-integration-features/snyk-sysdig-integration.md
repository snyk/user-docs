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

## Enriching Snyk vulnerability data and priority score

Snyk uses packages executed at runtime to enrich the priority score of vulnerabilities detected by Snyk. This allows to better prioritize which vulnerabilities to fix first. The priority score will then be available both on the Project page and in the [Snyk public API](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues).

![The "Executed" badge appearing on packages executed at runtime.](<../../../.gitbook/assets/image (113) (1) (2) (1) (1) (2) (1) (1) (1).png>)

In order to see which packages have been executed at runtime, you would need to wait for the next daily test, or import the workload manually into Snyk. After enabling the Sysdig integration, allow **4 hours** before manually importing the workload due to the following timing considerations:

{% hint style="info" %}
There are a couple timing considerations related to the collection of executed packages:

* Sysdig reports executed packages approximately 1 hour after they have been detected.
* The Snyk controller collects data about executed packages once every 4 hours.
* Snyk re-tests imported Kubernetes Projects for new vulnerabilities daily.
{% endhint %}

## Application support

For application vulnerabilities, we currently provide support for the following languages:

* Java
* JavaScript
* Go

You can find an up-to-date list of supported languages in [Snyk Container: Detecting Application Vulnerabilities in Container Images](https://docs.snyk.io/products/snyk-container/getting-around-the-snyk-container-ui/detecting-application-vulnerabilities-in-container-images).
