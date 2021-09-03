# Automatic import/deletion of Kubernetes workloads projects

{% hint style="info" %}
This feature is currently in beta. We would appreciate any feedback you might have.
{% endhint %}

You can configure the Snyk controller to automatically import and update scanned workloads directly in Snyk to test and monitor for vulnerabilities. You can also automatically delete imported projects once workloads are deleted from the cluster. The controller evaluates policy decisions using a policy file written in [Rego policy language](https://www.openpolicyagent.org/docs/latest/policy-language/).

### Enabling workload auto-import and auto-delete

The Helm chart of the Snyk controller is already provisioned with a default Rego policy to process events for any workload except Jobs. To enable this feature, provide your Snyk Organization public ID in the Helm chart installation.

```text
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
    --namespace snyk-monitor \
    --set clusterName="Production cluster" \
    --set policyOrgs={19982df2-0ed5-4a16-b355-e6535cfc41ef}
```

Note that _**policyOrgs**_ is a list of organization public IDs. You can add more than one Organization to use the auto-import and auto-delete capabilities. You can locate this public ID under your organization's settings page.

{% hint style="info" %}
You can only use organizations that share the same Kubernetes integration ID used to provision the Snyk controller.
{% endhint %}

### Policy syntax

Provide the policy file to the Snyk controller in a ConfigMap. The policy syntax looks like this:

```text
package snyk
orgs := [ ]
default workload_events = false
```

You can flip the value to **true** to automatically import or delete everything in the cluster. **Tip**: exclude Jobs from auto-import as they can be noisy and generate lots of workload imports in your Snyk organization.

Both **package snyk** and the key **workload\_events** are mandatory by Snyk Controller.

### Defining rules

To define your own rules, set a condition on the **workload\_events** key and by providing your organization public ID. For example, to import workloads from the **default** namespace and automatically delete them on Snyk side once they are deleted from the cluster, the policy would look like this:

```text
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.metadata.namespace == "default"
}
```

Here, **input** refers to the Kubernetes metadata of the workload scanned by the Snyk controller.

You can also create a policy for workload events \(creation/deletion\) with a specific annotation:

```text
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.metadata.annotations.team == "apollo"
}
```

### Excluding workload types

As best practice, we recommend excluding specific workload types such as Pods and Jobs from workload events \(creation/deletion\), as they can be really noisy and can generate lots of workload imports in your Snyk organization. You can do this with the following example policy:

```text
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.kind != "Job"
    input.kind != "Pod"
}
```

### Configure Snyk controller to use the policy

```text
kubectl create configmap snyk-monitor-custom-policies \
    -n snyk-monitor \
    --from-file=workload-events.rego # This name is hardcoded
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
    --namespace snyk-monitor \
    --set clusterName="Production cluster" \
    --set workloadPoliciesMap=snyk-monitor-custom-policies
```

{% hint style="info" %}
**NOTE**

Ensure the file is named **workload-events.rego**.
{% endhint %}

Now you can deploy the Snyk controller, or restart it if it is already running in order to pick up the policy.

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}