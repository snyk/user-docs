# Prioritization setup

## Prerequisites for Prioritization with Insights

To get value from Snyk Insights, you should use an application that scans images using [Snyk Container](../../../../scan-with-snyk/snyk-container/). You can get additional value by also scanning your open-source dependencies with [Snyk Open Source](../../../../scan-with-snyk/snyk-open-source/) and your source code with [Snyk Code](../../../../scan-with-snyk/snyk-code/).

Snyk prioritization with Insights product operates by providing you with four risk factors for your vulnerabilities:&#x20;

* [**Deployed**](../how-prioritization-with-insights-works/risk-factors/deployed.md): Is my code and container image deployed anywhere?
* [**Loaded package**](../how-prioritization-with-insights-works/risk-factors/loaded-package.md): Has a third-party package that is the dependency of an image been loaded?
* [**OS condition**](../how-prioritization-with-insights-works/risk-factors/os-condition.md): Does this vulnerability apply to my operating system?
* [**Public facing**](../how-prioritization-with-insights-works/risk-factors/public-facing.md): Does my container have any internet exposure?

To get data about these four risk factors, you must meet the following criteria:

* **Loaded package**: a package that is loaded more often than others poses a higher risk to your application compared to one that is rarely loaded. These are the minimum requirements you need to set up for the loaded package risk factor to be applied for prioritization with Insights:

You have to set up the Dynatrace or Sysdig integrations with Snyk AppRisk, or the Snyk Runtime Sensor. You can find more details on the [runtime third-party integrations](../../integrations-for-snyk-apprisk/connect-a-third-party-integration.md) page.&#x20;

* **OS condition:** Source code and dependencies are being built into a container image and scanned with Snyk Container. This is the minimum requirement to get value from the prioritization.

<figure><img src="../../../../.gitbook/assets/Example OS condition.png" alt="Source code and dependencies built into a container image"><figcaption><p>Source code and dependencies built into a container image</p></figcaption></figure>

* **Deployed and Public facing:** This container image is deployed onto a Kubernetes cluster, where you can deploy the [Kubernetes Connector](prioritization-setup-kubernetes-connector.md).

By ensuring these two requirements are satisfied, you get data for all four risk factors for the code in your scanned image.

Snyk **recommends** that you also perform the following steps to get the maximum value out of Insights:

* Scan the third-party dependencies using Snyk Open Source,
* Scan the source code using Snyk Code,

By scanning both the source code and the third-party dependencies, you will get risk factors data, which provides the application context to better prioritize your open issues.

{% hint style="info" %}
Snyk recommends starting with one application and expanding from there.
{% endhint %}

## Prioritization process overview

The major steps in setting up prioritization with Insights are as follows:

1. Grant users the Group Viewer role or the Organization Collaborator role. See [Prioritization setup: User permissions](prioritization-setup-user-permissions.md).
2. Create the required Organization, roles, and permissions, and deploy the agent. See [Prioritization setup: Kubernetes Connector](prioritization-setup-kubernetes-connector.md).

{% hint style="info" %}
The Kubernetes Connector is different from the Kubernetes Controller, Snyk-Monitor.
{% endhint %}

3. Scan your images properly so Snyk has access to the right data. See [Prioritization setup: Image scanning](prioritization-setup-image-scanning.md).
4. Set up the required link for the application on which you want to use prioritization. See [Prioritization setup: Associating Snyk Open Source, Code, and Container Projects](prioritization-setup-associating-snyk-open-source-code-and-container-projects.md).
5. Set up third-party runtime integrations or the Snyk Runtime Sensor to get even more runtime data.
6. To ensure you have properly set up the prioritization capability, navigate to the **Set up Snyk AppRisk** tab on the **Insights** page and view the data Snyk has access to.\
   You can also filter relevant sections by Organization for a granular view of your progress.

## Prioritize your integrations

When prioritizing with Insights, it is important to understand the available integration options and associated risk factors.&#x20;

Here are the integration options that you can choose from when setting up prioritization with Insights. You can customize the settings by navigating to the Group level [Snyk Web UI](../../../../getting-started/snyk-web-ui.md), the Setting menu, and then the Insights option.&#x20;

* **Kubernetes Connector**: Offers comprehensive monitoring for your Kubernetes deployments. This integration helps identify vulnerabilities within Kubernetes clusters and provides data on workload vulnerabilities, infrastructure misconfigurations, and potential malicious activity.
* **Snyk Runtime Sensor**: Used for deeper runtime analysis, providing detailed insight into the actual usage and potential vulnerabilities of your applications during execution. This sensor helps in gathering live traffic data, identifying runtime vulnerabilities, and assessing potential risks based on real-time application behavior.
* **Third-party Integrations**: These integrations, such as cloud providers or CI/CD tools, provide additional context and data sources for better vulnerability assessment. They help identify misconfigurations, exposure points, and integration-specific vulnerabilities.

Each of these integration options uses different risk factors:

* **Kubernetes Connector**: Deployed, Public facing
* **Snyk Runtime Sensor**: Deployed and Loaded Package
* **Third-party Integrations**: Deployed and Loaded Package

By leveraging these integration options, you can ensure comprehensive coverage and accurate prioritization of security risks.

The Public facing risk factor is less commonly used in all available integration options because it covers broader, more external threats that are harder to quantify and track within an internal monitoring system. In contrast, other risk factors such as the Deployed and Loaded Package offer more specific and actionable insights into the actual state and behavior of your applications, enabling better prioritization and remediation efforts.

