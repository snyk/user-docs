# Insights

{% hint style="info" %}
Insights is currently in Open Beta for Enterprise customers only. If you would like to participate, contact Snyk to add Insights to your plan. See [Feature release process](../../more-info/snyk-feature-release-process.md) for more details.
{% endhint %}

Insights is a platform capability providing advanced application intelligence to help you better identify and prioritize your Container, Code, and Open Source issues based on the risk they pose to your application.

The Snyk approach looks holistically at your application to understand the following:

* What source code and dependencies were built into a container image
* The operating system it is running on
* Where the image was deployed
* How the supporting Kubernetes and cloud infrastructure is configured

Insights works as illustrated in this example:

* Snyk Open Source has identified a high-severity Remote Code Exploit (RCE).&#x20;
* That RCE is built into a container image, which is deployed onto a production Kubernetes cluster, and the running container is configured to have access to the internet.&#x20;
* The combination of an RCE in a running image with internet access enables Snyk to determine that this particular vulnerability poses more risk to your application.

The documentation on these pages explains more about how Insights works, how to set Insights up, and how to use it. For more information about how to use Insights, see the video [Using Insights (Snyk)](https://www.youtube.com/watch?v=dVlamm9gIGI\&t=8s).&#x20;

