# Prioritization for Snyk AppRisk

{% hint style="warning" %}
**Release status** \
The Snyk Runtime Sensor and the thrid-party integrations providing the Deployed and Loaded Package risk factors are in Closed Beta and available only for Snyk AppRisk Pro.&#x20;
{% endhint %}

Snyk AppRisk uses holistic application intelligence to help you better identify and prioritize your Container, Code, and Open Source issues based on the risk they pose to your application. Users can also prioritize their issues based on asset classification as defined in Snyk AppRisk policies.&#x20;

If you use Snyk AppRisk, you can access the Issues page from the [Snyk Web UI](../../getting-started/snyk-web-ui.md#view-and-prioritize-issues).

## Insights for Snyk AppRisk

Snyk is introducing a new Issues page, providing a centralized view of all the issues identified by Snyk with additional asset context. This will help empower AppSec teams to better triage and remediate issues in Snyk.

**Issues** is available at the Group level or at the Organization level. The [evidence graph](using-the-issues-ui-with-snyk-apprisk/evidence-graph.md) is available only for Snyk AppRisk Pro users and only at the Group level.&#x20;

The Snyk approach looks holistically at your application to understand the following:

* What source code and dependencies were built into a container image
* The operating system it is running on
* Where the image was deployed
* How the supporting Kubernetes and cloud infrastructure is configured

The following video explains the initial steps of setting up your prioritization with Issues for Snyk AppRisk:

{% embed url="https://youtu.be/xnSUFsDzlMg" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)!
{% endembed %}

## Issues page: Snyk AppRisk Essentials vs Snyk AppRisk Pro

Issues is nested in the main left menu and provides information about the identified issues. Use the available filters to customize and prioritize the issues list.

{% hint style="info" %}
The Risk Factor column and filter, and with it, the evidence graph information, are available only for Snyk AppRisk Pro users.&#x20;
{% endhint %}

**Issues - Snyk AppRisk Essentials**

<figure><img src="../../.gitbook/assets/issues-apprisk-essentials.png" alt="Issues page - Snyk AppRisk Essentials"><figcaption><p>Issues page - Snyk AppRisk Essentials</p></figcaption></figure>

**Issues - Snyk AppRisk Pro**

<figure><img src="../../.gitbook/assets/image (451).png" alt="Issues page - Snyk AppRisk Pro"><figcaption><p>Issues page - Snyk AppRisk Pro</p></figcaption></figure>

The insights presented under the Issues menu for Snyk AppRisk Pro work as illustrated in the following example. The same principles apply to Snyk AppRisk Essentials without the availability of Risk Factors or evidence graph information.

* Snyk Open Source has identified a high-severity Remote Code Exploit (RCE).&#x20;
* That RCE is built into a container image, which is deployed onto a production Kubernetes cluster, and the running container is configured to have access to the internet.&#x20;
* The combination of an RCE in a running image with internet access enables Snyk to determine that this particular vulnerability poses more risk to your application than an RCE that is not deployed.

The following video explains the initial steps of setting up your prioritization with Issues insights for Snyk AppRisk Pro:

{% embed url="https://youtu.be/wGbcsSDDZE8" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)!
{% endembed %}
