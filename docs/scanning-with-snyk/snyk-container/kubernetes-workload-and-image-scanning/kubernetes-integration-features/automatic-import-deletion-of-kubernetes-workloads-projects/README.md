# Automatic import/deletion of Kubernetes workloads projects

{% hint style="info" %}
With the same integration id, you can import **many** clusters to **one** Snyk organization by giving clusters a unique cluster name during installation.

The auto-import/delete also allows you to import from **one** cluster to multiple Snyk organizations, read [Advance use of automatic import/deletion](https://docs.snyk.io/products/snyk-container/kubernetes-workload-and-image-scanning/kubernetes-integration-features/automatic-import-deletion-of-kubernetes-workloads-projects/advanced-use-of-automatic-import-deletion#using-more-than-one-org).

It is important to note that 1 Snyk organization has a limit of 25,000 projects, please consult your SCM for more information.
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
Only Snyk organizations that share the same Kubernetes integration ID can provision the Snyk controller to do so. More information on how to share the same integration ID can be found in [Clone an integration across your Snyk orgs](../../../../../integrations/git-repository-and-ci-cd-integrations-comparisons/managing-integrations/clone-an-integration-across-your-snyk-orgs.md).
{% endhint %}

## Annotated Import

Once an administrator for your Snyk account has installed the Snyk controller on your Kubernetes cluster, add workloads for testing. Kubernetes collaborators can mark workloads from the cluster to be automatically added to Snyk

## Automatically add, update and remove workloads

Once you’ve configured the integration between Snyk and your cluster, you can annotate your workloads in order to have them automatically added as projects for testing in Snyk.

{% hint style="info" %}
The annotated import happens when the **image** itself changes (rescans the workload due to image change) or when the **workload details** change (which creates a new revision of the workload). Changing the annotation for the workload will not cause a workload change.

If the workload is only annotated after it has been scanned by `snyk monitor the` annotation will not be recognized until a significant change takes place that causes a full rescan. Terminating the `snyk monitor` pod is one way to force a rescan.
{% endhint %}

Annotate any of the following workload types:

* Deployments
* ReplicaSets
* DaemonSets
* StatefulSets
* Jobs
* CronJobs
* ReplicationControllers
* Pods

**The steps follow:**

1. Log in to your account and navigate to the relevant group and organization that you want to manage.
2. Click on settings ![](../../../../../.gitbook/assets/cog\_icon.png) > **General**.
3. Copy the **Organization ID** value.
4. Add an annotation to the workload with the key `orgs.k8s.snyk.io/v1`, entering the Organization ID as the value in a comma-separated list.

You can also annotate a single workload to be added to multiple organizations.

1.  The Snyk controller automatically picks up on the changes to your workload and ensures that the workload is automatically imported to Snyk as a Snyk project.

    Example: Deployment YAML file annotated to be automatically imported into an organization

    ```
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: my-app-deployment
      annotations:
        orgs.k8s.snyk.io/v1: cacb791e-07cc-4b10-b4be-64de19f532f1
    spec:
      template:
        spec:
          containers:
          …
    ```

    To annotate for multiple organizations use a comma-separated list.
2. Once imported, the project remains in your Snyk organization even if you remove the annotation. To remove the project from Snyk, you should delete the annotation and delete it from the Snyk UI or [with the API](https://snyk.docs.apiary.io/#reference/projects/individual-project/delete-a-project).
