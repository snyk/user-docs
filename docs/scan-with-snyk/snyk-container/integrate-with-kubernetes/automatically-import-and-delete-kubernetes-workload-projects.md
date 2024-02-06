# Automatically import and delete Kubernetes workload Projects

{% hint style="info" %}
**Feature availability**

This feature is still in Beta. Auto import and deletion are not guaranteed and can depend on the network.

If you don't find your workload in Snyk, please import the workload manually.

If your workload persists in Snyk after it was deleted from your cluster, please delete the workload manually.
{% endhint %}

Using the same integration ID, you can import multiple clusters to one Snyk Organization by giving clusters a unique cluster name during installation.

The automatic import and delete feature also allows you to import from one cluster to multiple Snyk organizations. For more information, see [Advanced use of the automatic import and delete feature](automatically-import-and-delete-kubernetes-workload-projects.md#advanced-use-of-the-automatic-import-and-delete-feature).

One Snyk Organization has a limit of 25,000 Projects. For more information, consult Snyk admin.

## **Prerequisites for automatic import and delete**

You must first have an account with Snyk and be onboarded to your Organization by an administrator.

In addition, you must configure the integration between Snyk and your Kubernetes environment per Organization. To ensure the integration is configured, ensure you have an Integration ID. For more information, see [Enable the Kubernetes integration](overview-of-the-kubernetes-integration/enable-the-kubernetes-integration.md).

You can configure the Snyk Controller to automatically import and update scanned workloads directly in Snyk to scan and monitor for vulnerabilities. You can also automatically delete imported Projects once workloads are deleted from the cluster.

## Enable workload auto-import and auto-delete

The Helm chart of the Snyk Controller is already provisioned with a default policy to process events for any workload except Jobs and Pods. To enable this feature, provide your Snyk Organization public ID in the Helm chart installation.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
    --namespace snyk-monitor \
    --set clusterName="Production cluster" \
    --set policyOrgs={19982df2-0ed5-4a16-b355-e6535cfc41ef}
```

`policyOrgs` is a list of Organization public IDs. You can add more than one Organization to use the auto-import and auto-delete capabilities. Your public ID is available on your Organization's settings page.

{% hint style="info" %}
Only Snyk Organizations that share the same Kubernetes integration ID can provision the Snyk Controller to do so. For more information on how to share the same integration ID, see [Clone an integration across your Snyk Organizations](../../../enterprise-setup/snyk-broker/clone-an-integration-across-your-snyk-organizations.md).
{% endhint %}

## Advanced use of the automatic import and delete feature

If you have an advanced use case for automatically importing and deleting Kubernetes workload Projects, you can write your own rules. The Snyk Controller evaluates policy decisions using a policy file written in [Rego policy language](https://www.openpolicyagent.org/docs/latest/policy-language/). Ensure the file is named "workload-events.rego".

### Policy syntax

Provide the policy file to the Snyk Controller in a ConfigMap:

```
package snyk
orgs := ["<org-id>"]
default workload_events = false
```

{% hint style="warning" %}
It is important that the `default workload_events` is `false`. If the value is `true`, it will automatically import or delete everything in the cluster, including any `workload_events` defined in your policy.
{% endhint %}

Snyk does not recommend setting the `workload_events` key to `true`, as some workload types like Jobs and Pods can be noisy and generate lots of workload imports in your Snyk Organization.

{% hint style="info" %}
Both `package snyk` and the `workload_events` key are required by the Snyk Controller.
{% endhint %}

### Use more than one Organization

{% hint style="info" %}
This feature is only available with an Enterprise Plan. For details, see [pricing plans](https://snyk.io/plans/?\_gl=1\*myjr92\*\_ga\*MTYzMjUyMDYwNC4xNjg4OTkzNjQ2\*\_ga\_X9SH3KP7B4\*MTY5NTYzNDE0NC4xNDAuMS4xNjk1NjQxNTIyLjAuMC4w).
{% endhint %}

`Orgs` is a list of Organization public IDs. You can add more than one Organization to use the auto-import and auto-delete capabilities. Your public ID is available on your Organization's settings page.

```
package snyk
orgs := ["<org-id-1>","<org-id-2>"]
default workload_events = false
```

### Define rules

To define your own rules, set a condition on the `workload_events` key and provide your Organization public ID.

For example, to import workloads from the `default` namespace and to automatically delete them from Snyk after they are deleted from the cluster, the policy looks as follows:

```
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.metadata.namespace == "default"
}
```

In this example, `input` refers to the Kubernetes metadata of the workload scanned by the Snyk Controller.

You can also create a policy for workload events (creation/deletion) with a specific annotation:

```
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.metadata.annotations.team == "apollo"
}
```

### Exclude workload types

As best practice, Snyk recommends excluding specific workload types such as Pods and Jobs from workload events (creation/deletion), as they can be noisy and can generate lots of workload imports in your Snyk Organization.

You can exclude workload types with the following example policy:

```
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.kind != "Job"
    input.kind != "Pod"
}
```

### Configure the Snyk Controller to use the policy

```
kubectl create configmap snyk-monitor-custom-policies \
    -n snyk-monitor \
    --from-file=workload-events.rego # This name is hardcoded
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
    --namespace snyk-monitor \
    --set clusterName="Production cluster" \
    --set workloadPoliciesMap=snyk-monitor-custom-policies
```

Now, you can deploy the Snyk Controller, or restart it if it is already running, in order for it to pick up the policy. New workloads are now visible in Snyk.

