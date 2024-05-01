# Prioritization with Insights for Snyk AppRisk

{% hint style="warning" %}
**Release status** \
The Loaded Package, Deployed from third-party integrations risk factors, and the Snyk Runtime Sensor are currently in Closed Beta and available only for Snyk AppRisk Pro plans.&#x20;

Contact your account manager if you are interested in Snyk AppRisk Pro.
{% endhint %}

Snyk AppRisk uses holistic application intelligence to help you better identify and prioritize your Container, Code, and Open Source issues based on the risk they pose to your application. Users can also prioritize their issues based on asset classification as defined in Snyk AppRisk policies.&#x20;

{% hint style="info" %}
Currently, asset class prioritization is available only for Snyk Code issues within the Insights UI.
{% endhint %}

Prioritization with Insights is available at the Group level or at the Organization level. The [evidence graph](using-the-insights-ui/evidence-graph.md) is available only at the Group level.

The Snyk approach looks holistically at your application to understand the following:

* What source code and dependencies were built into a container image
* The operating system it is running on
* Where the image was deployed
* How the supporting Kubernetes and cloud infrastructure is configured

Prioritization with Insights works as illustrated in this example:

* Snyk Open Source has identified a high-severity Remote Code Exploit (RCE).&#x20;
* That RCE is built into a container image, which is deployed onto a production Kubernetes cluster, and the running container is configured to have access to the internet.&#x20;
* The combination of an RCE in a running image with internet access enables Snyk to determine that this particular vulnerability poses more risk to your application than an RCE that is not deployed.

The following video explains the initial steps of setting up your prioritization with Insights for Snyk AppRisk:

{% embed url="https://www.youtube.com/watch?v=TqcJ2UIAOvU" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)!
{% endembed %}
