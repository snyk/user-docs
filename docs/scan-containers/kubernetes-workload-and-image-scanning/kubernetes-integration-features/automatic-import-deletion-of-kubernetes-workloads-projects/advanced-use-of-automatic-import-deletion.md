# Advanced use of automatic import/deletion

If you have an advanced use case for automatic import/deletion of Kubernetes workload projects, you can write your own rules. The controller evaluates policy decisions using a policy file written in [Rego policy language](https://www.openpolicyagent.org/docs/latest/policy-language/). Ensure the file is named **workload-events.rego**.

## Policy syntax

Provide the policy file to the Snyk controller in a ConfigMap. The policy syntax looks like this:

```
package snyk
orgs := ["<org-id>"]
default workload_events = false
```

It is important that the default workload\_events is **false**. If you flip the value to true, it will automatically import or delete **everything** in the cluster including any workload\_events defined in your policy.

{% hint style="warning" %}
Setting the workload\_events key to true is **not recommended** as some workload types like Jobs and Pods can be noisy and generate lots of workload imports in your Snyk organization
{% endhint %}

Both **package snyk** and the key **workload\_events** are mandatory by Snyk Controller.

## Using more than one org

{% hint style="info" %}
This feature is only available if you are on an Enterprise Plan.
{% endhint %}

Orgs is a list of organization public IDs. You can add more than one organization to use the auto-import and auto-delete capabilities. You can locate this public ID in the settings page for the organization.

```
package snyk
orgs := ["<org-id-1>","<org-id-2>"]
default workload_events = false
```

## Defining rules

To define your own rules, set a condition on the **workload\_events** key and by providing your organization public ID. For example, to import workloads from the **default** namespace and automatically delete them on Snyk side once they are deleted from the cluster, the policy would look like this:

```
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.metadata.namespace == "default"
}
```

Here, **input** refers to the Kubernetes metadata of the workload scanned by the Snyk controller.

You can also create a policy for workload events (creation/deletion) with a specific annotation:

```
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.metadata.annotations.team == "apollo"
}
```

## Excluding workload types

As best practice, we recommend excluding specific workload types such as Pods and Jobs from workload events (creation/deletion), as they can be really noisy and can generate lots of workload imports in your Snyk organization. You can do this with the following example policy:

```
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.kind != "Job"
    input.kind != "Pod"
}
```

## Configure Snyk controller to use the policy

```
kubectl create configmap snyk-monitor-custom-policies \
    -n snyk-monitor \
    --from-file=workload-events.rego # This name is hardcoded
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
    --namespace snyk-monitor \
    --set clusterName="Production cluster" \
    --set workloadPoliciesMap=snyk-monitor-custom-policies
```

Now you can deploy the Snyk controller, or restart it if it is already running in order to pick up the policy. You will now see new workloads in Snyk.
