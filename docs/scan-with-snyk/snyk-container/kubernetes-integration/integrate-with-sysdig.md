# Integrate with Sysdig

To enhance its capabilities when detecting workload information, Snyk has partnered with Sysdig. The integration enriches the workload issues that Snyk detects with the runtime data provided by Sysdig.

## Enable the Sysdig integration

For a successful integration with Sysdig, the Snyk Controller requires an extra Sysdig Secret in the `snyk-monitor` namespace. The Sysdig Secret name is `snyk-sysdig-secret`.

{% hint style="info" %}
Execute the commands below after installing Sysdig, in order to allow the Snyk Controller to detect Sysdig in the cluster.
{% endhint %}

Create the `snyk-sysdig-secret` in the `snyk-monitor` namespace:

```
kubectl create secret generic snyk-sysdig-secret -n snyk-monitor \
  --from-literal=token=$SYSDIG_RISK_SPOTLIGHT_TOKEN \
  --from-literal=endpoint=$SYSDIG_ENDPOINT_URL \
  --from-literal=cluster=$SYSDIG_AGENT_CLUSTER
```

SYSDIG\_RISK\_SPOTLIGHT\_TOKEN is the "Risk Spotlight Integrations Token" and must be generated through the Sysdig UI. To create this API token, see the[ Sysdig Risk Spotlight guide](https://docs.sysdig.com/en/docs/sysdig-secure/integrations-for-sysdig-secure/risk-spotlight-integrations/#generate-a-token-for-the-integration).

SYSDIG\_ENDPOINT\_URL is associated with your Sysdig SaaS application and region. To identify it, see [SaaS Regions and IP Ranges](https://docs.sysdig.com/en/docs/administration/saas-regions-and-ip-ranges). For example, for US West (Oregon), the domain is [us2.app.sysdig.com](https://us2.app.sysdig.com/) (you must omit the prefix "https://").

SYSDIG\_AGENT\_CLUSTER are the ones that you configured when [installing the Sysdig Agent](https://docs.sysdig.com/en/docs/installation/sysdig-secure/install-agent-components/kubernetes/#parameter-definitions) - global.clusterConfig.name.

To enable Snyk to integrate with Sysdig and collect information about packages executed at runtime, use `--set sysdig.enabled=true` when installing the Snyk Controller:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
  --namespace snyk-monitor \
  --set clusterName="Production cluster" \
  --set sysdig.enabled=true
```

Your Snyk Controller now collects data from Sysdig every 30 minutes.&#x20;

## Enrich Snyk vulnerability data and priority score

To enrich the priority score of vulnerabilities it detects, Snyk uses packages executed at runtime.  This allows Snyk to better prioritize which vulnerabilities to fix first. The priority score is available on both the **Project** page and in the [Snyk public API](../../../snyk-api/reference/projects-v1.md#org-orgid-project-projectid-aggregated-issues).

![Packages executed at runtime](<../../../.gitbook/assets/image (238).png>)

To see which packages have been executed at runtime, you must wait for the next daily scan or manually import the workload into Snyk.

After enabling the Sysdig integration, allow 30 minutes before manually importing the workload. This is because of the following timing considerations related to the collection of executed packages:

* The Snyk Controller collects data about executed packages once every 30 minutes.
* Snyk re-scans imported Kubernetes Projects for new vulnerabilities daily.

## Application support

For application vulnerabilities, Snyk currently provides support for the following languages:

* Java
* JavaScript
* Go

To see the updated list of supported languages, see [Detect application vulnerabilities in container images](../use-snyk-container/detect-application-vulnerabilities-in-container-images.md).
