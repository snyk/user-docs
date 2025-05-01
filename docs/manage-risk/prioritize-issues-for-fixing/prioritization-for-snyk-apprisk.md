# Prioritization for Snyk AppRisk

Snyk uses holistic application intelligence to help you better identify and prioritize your Container, Code, and Open Source issues based on the risk they pose to your application. Users can also prioritize their issues based on asset classification as defined in asset-related policies.&#x20;

You can access the [Issues](../../getting-started/snyk-web-ui.md#view-and-prioritize-issues) page from the [Snyk Web UI](../../getting-started/snyk-web-ui.md#view-and-prioritize-issues).

The Issues page provides a centralized view of all the issues identified by Snyk with additional asset context. This will help empower AppSec teams to better triage and remediate issues in Snyk.

**Issues** is available at the Group level or at the Organization level.&#x20;

## Insights for Snyk AppRisk

The [evidence graph](using-the-issues-ui-with-snyk-apprisk/evidence-graph.md) is available only for Snyk AppRisk users and only at the Group level.&#x20;

The Snyk approach looks holistically at your application to understand the following:

* What source code and dependencies were built into a container image
* The operating system it is running on
* Where the image was deployed
* How the supporting Kubernetes and cloud infrastructure is configured

## Issues page for Snyk AppRisk&#x20;

**Issues** is nested in the main left menu and provides information about the identified issues. Use the available filters to customize and prioritize the issues list.

{% hint style="info" %}
The Risk Factor column and filter, and with it, the evidence graph information, are available only for Snyk AppRisk users.&#x20;
{% endhint %}

**Issues - Snyk AppRisk**&#x20;

<figure><img src="../../.gitbook/assets/image (451) (1).png" alt="Issues page - Snyk AppRisk"><figcaption><p>Issues menu - Snyk AppRisk </p></figcaption></figure>

The insights presented under the Issues menu for Snyk AppRisk work as illustrated in the following example.&#x20;

* Snyk Open Source has identified issues.&#x20;
* The vulnerabilities were built into a container image, which is deployed onto a production Kubernetes cluster, and the running container is configured to have access to the internet.&#x20;
* The combination of a critical vulnerability in a running image with internet access enables Snyk to determine that this particular vulnerability poses more risk to your application than one that is not deployed.

The following video demonstrates prioritizing issues with Snyk AppRisk using business, application, and runtime context:

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737745405/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-_7b_-_v1_-_Issue_Prioritization_with_Runtime_Insights.mp4" %}
Prioritizing issues with runtime insights
{% endembed %}

