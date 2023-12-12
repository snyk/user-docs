# Risk-Based Prioritization for Snyk AppRisk

Snyk AppRisk Essentials uses holistic application intelligence to help you better identify and prioritize your Container, Code, and Open Source issues based on the risk they pose to your application. Users can also prioritize their issues based on asset classification as defined in Snyk AppRisk policies.&#x20;

{% hint style="info" %}
Currently, asset class prioritization is available only for Snyk Code issues within the Insights UI.
{% endhint %}

Risk-based prioritization is available at the Group level or at the Organization level. The [evidence graph](../../../manage-issues/insights/using-insights/evidence-graph.md) is available only at the Group level.

The Snyk approach looks holistically at your application to understand the following:

* What source code and dependencies were built into a container image
* The operating system it is running on
* Where the image was deployed
* How the supporting Kubernetes and cloud infrastructure is configured

Risk-based prioritization works as illustrated in this example:

* Snyk Open Source has identified a high-severity Remote Code Exploit (RCE).&#x20;
* That RCE is built into a container image, which is deployed onto a production Kubernetes cluster, and the running container is configured to have access to the internet.&#x20;
* The combination of an RCE in a running image with internet access enables Snyk to determine that this particular vulnerability poses more risk to your application than an RCE that is not deployed.
