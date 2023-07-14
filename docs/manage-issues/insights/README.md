# Insights

{% hint style="info" %}
Insights is currently in Open Beta for Enterprise customers only. If you would like to participate, contact Snyk to add Insights to your plan. See [Feature release process](../../more-info/snyk-feature-release-process.md) for more details.
{% endhint %}

Insights is a platform capability providing advanced application intelligence to help you better identify and prioritize your Code, Open Source, and Container issues based on the risk they pose to your application.

The Snyk approach looks holistically at your application to understand the following:

* What source code and dependencies were built into the container image
* The operating system it is running on
* Where the image was deployed
* How the supporting Kubernetes and cloud infrastructure is configured

An example of how Insights works follows:

* Snyk Open Source has identified a high-severity Remote Code Exploit (RCE).&#x20;
* That RCE is built into a container image, which is deployed onto a production Kubernetes cluster, and the running container is configured to have access to the internet.&#x20;
* The combination of an RCE in a running image with internet access enables us to determine that this particular vulnerability poses more risk to your application.

