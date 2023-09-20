# Automatic import/deletion of Kubernetes workloads projects

{% hint style="warning" %}
This feature is still in beta. Auto import and deletion are not guaranteed and can depend on the network.

If you don't find your workload in Snyk, please import the workload through manual import.

If your workload persists in Snyk after deletion from your cluster, please delete the workload through manual deletion.
{% endhint %}

{% hint style="info" %}
With the same integration id, you can import **many** clusters to **one** Snyk organization by giving clusters a unique cluster name during installation.

The auto-import/delete also allows you to import from **one** cluster to multiple Snyk organizations, read [Advance use of automatic import/deletion](https://docs.snyk.io/products/snyk-container/kubernetes-workload-and-image-scanning/kubernetes-integration-features/automatic-import-deletion-of-kubernetes-workloads-projects/advanced-use-of-automatic-import-deletion#using-more-than-one-org).

It is important to note that 1 Snyk organization has a limit of 25,000 projects. For more information, consult your Snyk admin.
{% endhint %}

## **Prerequisites**

You must have an account with Snyk and be onboarded to your organization by an administrator.

In addition the integration must be configured between Snyk and your Kubernetes environment per organization. To verify the integration is configured, ensure you have an [**Integration ID**](../../kubernetes-integration-overview/viewing-your-kubernetes-integration-settings.md)\*\*\*\*

You can configure the Snyk controller to automatically import and update scanned workloads directly in Snyk to test and monitor for vulnerabilities. You can also automatically delete imported projects once workloads are deleted from the cluster.

## Enabling workload auto-import and auto-delete

The Helm chart of the Snyk controller is already provisioned with a default policy to process events for any workload except Jobs and Pods. To enable this feature, provide your Snyk Organization public ID in the Helm chart installation.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
    --namespace snyk-monitor \
    --set clusterName="Production cluster" \
    --set policyOrgs={19982df2-0ed5-4a16-b355-e6535cfc41ef}
```

Note that _**policyOrgs**_ is a list of organization public IDs. You can add more than one organization to use the auto-import and auto-delete capabilities. You can locate this public ID under your organization's settings page.

{% hint style="info" %}
Only Snyk organizations that share the same Kubernetes integration ID can provision the Snyk controller to do so. More information on how to share the same integration ID can be found in [Clone an integration across your Snyk orgs](../../../../enterprise-setup/snyk-broker/clone-an-integration-across-your-snyk-organizations.md).
{% endhint %}

##
