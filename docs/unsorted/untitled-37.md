# Are Snyk patches available indefinitely to apply/view through Snyk's website ?

* [ Kubernetes integration overview](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360003916138-Kubernetes-integration-overview/README.md)
* [ Install the Snyk controller with Helm](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360003916158-Install-the-Snyk-controller-with-Helm/README.md)
* [ Automatic import/deletion of Kubernetes workloads projects](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360020835037-Automatic-import-deletion-of-Kubernetes-workloads-projects/README.md)
* [ Install the Snyk controller with OpenShift 4 and OperatorHub](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360006548317-Install-the-Snyk-controller-with-OpenShift-4-and-OperatorHub/README.md)
* [ Install the Snyk controller on Amazon Elastic Kubernetes Service \(Amazon EKS\)](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360011128137-Install-the-Snyk-controller-on-Amazon-Elastic-Kubernetes-Service-Amazon-EKS-/README.md)
* [ Adding Kubernetes workloads for security scanning](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360003947117-Adding-Kubernetes-workloads-for-security-scanning/README.md)
* [ Viewing project details and test results](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360003916178-Viewing-project-details-and-test-results/README.md)
* [ Snyk Priority Score and Kubernetes](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360010906897-Snyk-Priority-Score-and-Kubernetes/README.md)
* [ Viewing your Kubernetes integration settings](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360006368657-Viewing-your-Kubernetes-integration-settings/README.md)
* [ Disable the Kubernetes integration](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360003947137-Disable-the-Kubernetes-integration/README.md)

## Automatic import/deletion of Kubernetes workloads projects

This feature is currently in beta. We would appreciate any feedback you might have.

You can configure the Snyk controller to automatically import and update scanned workloads directly in Snyk to test and monitor for vulnerabilities. You can also automatically delete imported projects once workloads are deleted from the cluster. The controller evaluates policy decisions using a policy file written in [Rego policy language](https://www.openpolicyagent.org/docs/latest/policy-language/).

### Enabling workload auto-import and auto-delete

The Helm chart of the Snyk controller is already provisioned with a default Rego policy to process events for any workload except Jobs. To enable this feature, provide your Snyk Organization public ID in the Helm chart installation.

```text
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \\
    --namespace snyk-monitor \\
    --set clusterName="Production cluster" \\
    --set policyOrgs={19982df2-0ed5-4a16-b355-e6535cfc41ef}
```

Note that _**policyOrgs**_ is a list of organization public IDs. You can add more than one Organization to use the auto-import and auto-delete capabilities. You can locate this public ID under your organization's settings page.

You can only use organizations that share the same Kubernetes integration ID used to provision the Snyk controller.

### Policy syntax

Provide the policy file to the Snyk controller in a ConfigMap. The policy syntax looks like this:

```text
package snyk
orgs := [ ]
default workload_events = false
```

You can flip the value to **true** to automatically import or delete everything in the cluster. **Tip**: exclude Jobs from auto-import as they can be noisy and generate lots of workload imports in your Snyk organization.

Both **package snyk** and the key **workload\_events** are mandatory by Snyk Controller.

#### Defining rules

To define your own rules, set a condition on the **workload\_events** key and by providing your organization public ID. For example, to import workloads from the **default** namespace and automatically delete them on Snyk side once they are deleted from the cluster, the policy would look like this:

```text
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.metadata.namespace == "default"
}
```

Here **input** refers to the Kubernetes metadata of the workload scanned by the Snyk controller.

You can also create a policy for workload events \(creation/deletion\) with a specific annotation:

```text
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.metadata.annotations.team == "apollo"
}
```

#### Excluding Jobs

You should exclude Jobs from workload events \(creation/deletion\) as they can be really noisy and can generate lots of workload imports in your Snyk organization. You can do this with the following example policy:

```text
package snyk
orgs := ["19982df2-0ed5-4a16-b355-e6535cfc41ef"]
default workload_events = false
workload_events {
    input.kind != "Job"
}
```

### Configure Snyk controller to use the policy

```text
kubectl create configmap snyk-monitor-custom-policies \\
    -n snyk-monitor \\
    --from-file=workload-events.rego # This name is hardcoded
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \\
    --namespace snyk-monitor \\
    --set clusterName="Production cluster" \\
    --set workloadPoliciesMap=snyk-monitor-custom-policies
```

**Note**: Ensure the file is named **workload-events.rego**.

Now you can deploy the Snyk controller, or restart it if it is already running in order to pick up the policy.

