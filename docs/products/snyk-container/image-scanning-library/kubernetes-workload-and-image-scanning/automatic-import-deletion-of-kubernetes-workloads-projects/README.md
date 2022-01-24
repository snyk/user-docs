# Automatic import/deletion of Kubernetes workloads projects

{% hint style="info" %}
This feature is currently in beta. We would appreciate any feedback you might have.
{% endhint %}

You can configure the Snyk controller to automatically import and update scanned workloads directly in Snyk to test and monitor for vulnerabilities. You can also automatically delete imported projects once workloads are deleted from the cluster.&#x20;

## Enabling workload auto-import and auto-delete

The Helm chart of the Snyk controller is already provisioned with a default policy to process events for any workload except Jobs and Pods. To enable this feature, provide your Snyk Organization public ID in the Helm chart installation.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
    --namespace snyk-monitor \
    --set clusterName="Production cluster" \
    --set policyOrgs={19982df2-0ed5-4a16-b355-e6535cfc41ef}
```

Note that _**policyOrgs**_ is a list of organization public IDs. You can add more than one Organization to use the auto-import and auto-delete capabilities. You can locate this public ID under your organization's settings page.

{% hint style="info" %}
You can only use organizations that share the same Kubernetes integration ID used to provision the Snyk controller.
{% endhint %}
