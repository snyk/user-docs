# Secure Kubernetes

![](../../../.gitbook/assets/snykk8s.png)

In this module, we will provide some hands-on examples on how to deploy [Dynatrace ActiveGate](https://www.dynatrace.com/support/help/setup-and-configuration/dynatrace-activegate/) as well as [Snyk Monitor for Kubernetes Helm Chart](https://artifacthub.io/packages/helm/snyk/snyk-monitor) into an existing Kubernetes cluster. You will then be guided on how to import your Kubernetes project into Snyk and correlate findings in Dynatrace with remediation advice in Snyk.

{% hint style="danger" %}
This feature is available with all paid plans. See [Pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

### Prerequisites:

1. An existing Dynatrace environment or a [free trial](https://www.dynatrace.com/trial/).
2. A Snyk **Business or Enterprise** Plan. See our [pricing page](https://snyk.io/plans/) for details.
3. A running Kubernetes cluster.
4. Snyk's Goof [sample application](https://github.com/snyk-partners/k8s-goof).

### Objectives:

* Configure Snyk's Kubernetes integration.
* Deploy Goof sample app.
* Activate Dynatrace Application Security module to detect vulnerabilities in your environment at runtime.
* Remediate issues with Snyk Container.

