# Annotated import



{% hint style="warning" %}
**Deprecated**

Annotated import is deprecated and no longer maintained.&#x20;

See [Automatically import and delete Kubernetes workload projects](automatically-import-and-delete-kubernetes-workload-projects.md).
{% endhint %}

After a Snyk administrator has installed the Snyk Controller on your Kubernetes cluster, you can add workloads for scanning. Kubernetes collaborators can mark workloads from the cluster to be automatically added to Snyk.

## Automatically add, update, and remove workloads

After you have configured the integration between Snyk and your cluster, you can annotate your workloads in order to have them automatically added to Snyk as Projects to scan.

{% hint style="info" %}
The annotated import occurs when the image itself changes (the workload is rescanned due to image change) or when the workload details change (which creates a new revision of the workload). Changing the annotation for the workload does not cause a workload change.

If the workload is annotated only after it has been scanned with `snyk monitor`, the annotation is not recognized until a significant change takes place that causes a full rescan. One way to force a rescan is to terminate the `snyk monitor` pod.
{% endhint %}

You can annotate any of the following workload types:

* Deployments
* ReplicaSets
* DaemonSets
* StatefulSets
* Jobs
* CronJobs
* ReplicationControllers
* Pods

To do this:

1. In Snyk, navigate to the relevant Group and Organization that you want to manage.
2. Navigate to **Settings** > **General**.
3. Copy the **Organization ID**.
4. Add an annotation to the workload with the key `orgs.k8s.snyk.io/v1`, and add the Organization ID as the value in a comma-separated list.

You can also annotate a single workload to be added to multiple Organizations. To do this:

1. The Snyk Controller automatically picks up on the changes to your workload and ensures that the workload is automatically imported to Snyk as a Snyk Project.

For example:

Deployment YAML file annotated to be automatically imported into an Organization:

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

To annotate for multiple Organizations, use a comma-separated list.

2. After it is imported, the Project remains in your Snyk Organization even if you remove the annotation. To remove the Project from Snyk, you must delete the annotation and delete the Project from the Snyk UI or by using the API endpoint [Delete a project](../../../snyk-api/reference/projects-v1.md#org-orgid-project-projectid-2).
