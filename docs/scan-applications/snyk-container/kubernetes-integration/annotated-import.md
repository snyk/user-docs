# Annotated import



{% hint style="warning" %}
This feature is deprecated and Snyk no longer maintains it, for auto import/deletion of Kubernetes workload please follow instruction [here](automatically-import-and-delete-kubernetes-workload-projects.md)
{% endhint %}

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
2. Click on settings ![](../../../.gitbook/assets/cog\_icon.png) > **General**.
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
