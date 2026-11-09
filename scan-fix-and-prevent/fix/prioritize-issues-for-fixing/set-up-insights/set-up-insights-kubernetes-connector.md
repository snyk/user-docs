---
description: The Kubernetes connector for Snyk Insights reached end of life on November 2, 2026. Uninstall it from your Kubernetes clusters.
nav_context: classic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

{% include "../../../.gitbook/includes/new-navigation-banner.md" %}

# Insights Kubernetes connector (end of life)

{% hint style="warning" %}
The Kubernetes connector reached end of life on November 2, 2026. Snyk no longer accepts or uses the data that the connector collects, and does not release updates for the connector, including security fixes.

Do not install the Kubernetes connector in new clusters. If the connector runs in any of your clusters, [uninstall it](#uninstall-the-kubernetes-connector). For details, visit the [Kubernetes connector end-of-life announcement](https://updates.snyk.io/kubernetes-connector-reaches-end-of-life-02-nov-202/).
{% endhint %}

The Kubernetes connector was an agent that you deployed in your Kubernetes clusters with the `kubernetes-scanner` Helm chart. It collected the following information for Insights:

* The deployed images, with their IDs and SHAs
* The services and their configuration, for example, the networking services you use

Snyk used this information to compute the **Deployed** and **Public facing** risk factors.

## What end of life means for the Kubernetes connector

* Snyk stopped accepting data from the connector on November 2, 2026. A connector that is still installed receives an error each time it sends data, logs the error, and retries. The connector does not crash, but it continues to use cluster resources until you uninstall it.
* Snyk no longer computes the **Deployed** and **Public facing** risk factors.
* Snyk archived the [kubernetes-scanner GitHub repository](https://github.com/snyk/kubernetes-scanner) on October 26, 2026. A connector that stays installed keeps running an image that receives no security fixes.

## Kubernetes connector compared with the Kubernetes monitor

The Kubernetes connector and the [Kubernetes monitor](../../../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/) are separate integrations. The end of life of the Kubernetes connector does not affect the Kubernetes monitor.

| Integration | Helm chart | What it does |
| --- | --- | --- |
| Kubernetes connector | `kubernetes-scanner` | Extracted workload configurations from a Kubernetes cluster for Insights. End of life. |
| Kubernetes monitor, also called Snyk Controller or Snyk-Monitor | `snyk-monitor` | Extracts images from the workloads in a Kubernetes cluster and scans them for vulnerabilities. Not affected. |

## Uninstall the Kubernetes connector

Repeat these steps for each cluster where the connector runs. To run the Helm and `kubectl` commands, you must have permission to delete the resources that the Helm chart created in the cluster.

1. Find the release name and namespace of the connector. In the output, the **CHART** column shows `kubernetes-scanner` for the connector release.

    ```bash
    helm list --all-namespaces | grep kubernetes-scanner
    ```
2. Uninstall the release:

    ```bash
    helm uninstall <release-name> --namespace <namespace>
    ```
3. Delete the Kubernetes secret that stores the service account token for the connector. Helm does not remove this secret, because you created it outside the chart.

    ```bash
    kubectl delete secret <secret-name> --namespace <namespace>
    ```
4. In Snyk, delete the service account that you created for the connector. Snyk invalidates the service account token immediately. For instructions, visit [How to edit or delete a service account](https://docs.snyk.io/platform-administration/service-accounts/service-accounts#how-to-edit-or-delete-a-service-account).
5. If you created a custom role only for the connector, delete the role. For instructions, visit [Delete a custom role](https://docs.snyk.io/platform-administration/user-management/user-role-management#delete-a-custom-role).

{% hint style="warning" %}
If you created a separate Snyk Organization for the Kubernetes connector, confirm that it contains no Projects or integrations that you use before you delete it. If the Organization is in a Group, only a Group Admin can delete it. For instructions, visit [Delete an Organization](https://docs.snyk.io/platform-administration/snyk-hierarchy/organizations/create-and-delete-organizations#delete-an-organization).
{% endhint %}
