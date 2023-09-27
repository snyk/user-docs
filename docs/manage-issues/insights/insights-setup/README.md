# Insights setup

## Prerequisites for Insights

To get value from Insights, you **must choose** an application where you are using [Snyk Container](../../../scan-applications/snyk-container/) to scan the images. You can get additional value by also scanning your open-source dependencies with [Snyk Open Source](../../../scan-application-code/snyk-open-source/) and your source code with [Snyk Code](../../../scan-application-code/snyk-code/).

Snyk Insights operates by providing you with three risk factors for your vulnerabilities:&#x20;

* **OS condition**: Does this vulnerability apply to my operating system?
* **Deployed**: Is my code and container image deployed anywhere?
* **Public facing**: Does my container have any internet exposure?

To get data about these three risk factors, you must meet the following criteria:

* **OS condition:** Source code and dependencies are being built into a container image and scanned with Snyk Container. This is the minimum requirement to get value from Insights.

<figure><img src="../../../.gitbook/assets/Example OS condition.png" alt="Source code and dependencies built into a container image"><figcaption><p>Source code and dependencies built into a container image</p></figcaption></figure>

* **Deployed and Public facing:** This container image is deployed onto a Kubernetes cluster, where you can deploy the [Kubernetes Connector](insights-setup-kubernetes-connector.md).

By ensuring these two requirements are satisfied, you get data for all three risk factors for the code in your scanned image

Snyk **recommends** that you also perform the following steps to get the maximum value out of Insights:

* Scan the third-party dependencies using Snyk Open Source,
* Scan the source code using Snyk Code,

By scanning both the source code and the third-party dependencies, you will get risk factors data, which provides the application context to better prioritize your open issues.

{% hint style="info" %}
Snyk recommends starting with one application and expanding from there.
{% endhint %}

## Insights process overview

The major steps in setting up Insights are as follows:

1. Grant users the Group Viewer role. See [Insights setup: User permissions](insights-setup-user-permissions.md).
2. Create the required Organization, roles, and permissions, and deploy the agent. See [Insights setup: Kubernetes Connector](insights-setup-kubernetes-connector.md).\
   &#x20;Note: The Kubernetes Connector is different from the Kubernetes Controller, Snyk-Monitor.
3. Scan your images properly so Insights has access to the right data. See [Insights setup: Image scanning](insights-setup-image-scanning.md).
4. Set up the required linking for the application on which you want to use Insights. See [Insights setup: Associating Snyk Open Source, Code, and Container Projects](insights-setup-associating-snyk-open-source-code-and-container-projects.md).
5. To ensure you have set up Insights properly, navigate to the **Set up Insights** tab on the **Insights** page and view the data Insights has access to.\
   You can also filter relevant sections by Organization for a granular view of your progress.



