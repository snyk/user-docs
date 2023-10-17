# Integrate with Sysdig

To enhance its capabilities when detecting workload information, Snyk has partnered with Sysdig. The integration enriches the workload issues that Snyk detects with the runtime data provided by Sysdig.

## Enable the Sysdig integration

For a successful integration with Sysdig, the Snyk Controller requires an extra Sysdig Secret in the `snyk-monitor` namespace. The Sysdig Secret name is `sysdig-eve-secret`.

{% hint style="info" %}
You must execute the commands below immediately after installing Sysdig, in order to allow the Snyk Controller to detect Sysdig in the cluster.
{% endhint %}

To install the Sysdig Secret, see [Sysdig Secret installation guide](https://docs.sysdig.com/en/docs/sysdig-secure/integrate-effective-vulnerability-exposure-with-snyk/#copy-the-sysdig-secret). After you have installed the Sysdig Secret, you must copy it to the `snyk-monitor` namespace:

```
kubectl get secret sysdig-eve-secret -n sysdig-agent -o yaml \
  | grep -v '^\s*namespace:\s' \
  | kubectl apply -n snyk-monitor -f -
```

To enable Snyk to integrate with Sysdig and collect information about packages executed at runtime, use `--set sysdig.enabled=true` when installing the Snyk Controller:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
  --namespace snyk-monitor \
  --set clusterName="Production cluster" \
  --set sysdig.enabled=true
```

Your Snyk Controller now collects data from Sysdig every four hours.&#x20;

## Enrich Snyk vulnerability data and priority score

To enrich the priority score of vulnerabilities it detects, Snyk uses packages executed at runtime.  This allows Snyk to better prioritize which vulnerabilities to fix first. The priority score is available on both the **Project** page and in the [Snyk public API](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues).

![Packages executed at runtime](<../../../.gitbook/assets/image (113) (1) (2) (1) (1) (2) (1) (1) (1).png>)

To see which packages have been executed at runtime, you must wait for the next daily scan or manually import the workload into Snyk.

After enabling the Sysdig integration, allow four hours before manually importing the workload. This is because of the following timing considerations related to the collection of executed packages:

* Sysdig reports executed packages approximately one hour after they have been detected.
* The Snyk Controller collects data about executed packages once every four hours.
* Snyk re-scans imported Kubernetes Projects for new vulnerabilities daily.

## Application support

For application vulnerabilities, Snyk currently provides support for the following languages:

* Java
* JavaScript
* Go

To see the updated list of supported languages, see [Detect vulnerabilities in container images](../use-snyk-container/detect-application-vulnerabilities-in-container-images.md).
