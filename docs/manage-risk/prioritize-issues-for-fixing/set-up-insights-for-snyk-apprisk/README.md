# Set up Insights for Snyk AppRisk

{% hint style="info" %}
The Set up Insights feature is available only with Snyk AppRisk.
{% endhint %}

## Prerequisites for prioritization with Snyk AppRisk

Customize the Snyk AppRisk prioritization using the Set up Insights option and an application that scans images using [Snyk Container](../../../scan-with-snyk/snyk-container/). You can get additional value by also scanning your open-source dependencies with [Snyk Open Source](../../../scan-with-snyk/snyk-open-source/) and your source code with [Snyk Code](../../../scan-with-snyk/snyk-code/).&#x20;

### What risk factors do I need?

{% hint style="info" %}
The Risk Factors are available only with Snyk AppRisk.
{% endhint %}

Snyk AppRisk Insights product operates by providing you with the following risk factors for your vulnerabilities:&#x20;

* [Deployed](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md): Is my code and container image deployed anywhere?
* [Loaded package](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md): Has a third-party package that is the dependency of an image been loaded?
* [OS condition](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-os-condition.md): Does this vulnerability apply to my operating system?
* [Public facing](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-public-facing.md): Does my container have any internet exposure?

### Risk factors usage criteria

To get data about these four risk factors, you must meet the following criteria:

#### **Loaded package risk factor**

You need to meet the following conditions to use the Loaded package risk factor:&#x20;

* There is a package that is loaded more often than others poses a higher risk to your application compared to one that is rarely loaded.
* This is the minimum requirement you need to set up for the loaded package risk factor to be applied for runtime prioritization with Insights:
* You must set up the Dynatrace or Sysdig integrations with Snyk AppRisk, or the [Snyk Runtime Sensor](../../../integrations/snyk-runtime-sensor.md). You can find more details on the [runtime third-party integrations](../../../integrations/connect-a-third-party-integration.md) page.&#x20;

#### **OS condition risk factor**

You need to meet the following conditions to use the OS condition risk factor:&#x20;

* Source code and dependencies are being built into a container image and scanned with Snyk Container. This is the minimum requirement to get value from the prioritization.

<figure><img src="../../../.gitbook/assets/Example OS condition.png" alt="Source code and dependencies built into a container image"><figcaption><p>Source code and dependencies built into a container image</p></figcaption></figure>

#### **Deployed and public facing risk factors**

You need to meet the following conditions to use the Deployed and Public facing risk factors:&#x20;

* For both risk factors, you have a container image that is deployed onto a Kubernetes cluster, where you can deploy the [Kubernetes Connector](set-up-insights-kubernetes-connector.md).

Ensure you meet these requirements to gather data for all four risk factors for the code in your scanned image.

### Maximize your insights

Snyk recommends that you also perform the following steps to get the maximum value out of insights:

* Scan the third-party dependencies using Snyk Open Source,
* Scan the source code using Snyk Code
* Start with one application and expand from there

By scanning both the source code and the third-party dependencies, you will get risk factors data, which provides the application context to better prioritize your open issues.

## Prioritization process overview

{% hint style="info" %}
Snyk recommends installing the [Snyk Runtime Sensor](../../../integrations/snyk-runtime-sensor.md) to achieve the most effective integration and access its continuously expanded set of features.
{% endhint %}

The major steps in setting up issues are as follows:

1. Grant users the Group Viewer role or the Organization Collaborator role. See [Prioritization setup: User permissions](set-up-insights-user-permissions.md).
2. Create the required Organization, roles, and permissions, and deploy the agent. See [Prioritization setup: Kubernetes Connector](set-up-insights-kubernetes-connector.md).

{% hint style="info" %}
The Kubernetes Connector is different from the Kubernetes Controller, Snyk-Monitor.
{% endhint %}

3. Scan your images properly so Snyk has access to the right data. See [Prioritization setup: Image scanning](set-up-insights-image-scanning.md).
4. Set up the required link for the application on which you want to use prioritization. See [Prioritization setup: Associating Snyk Open Source, Code, and Container Projects](set-up-insights-associating-snyk-open-source-code-and-container-projects.md).
5. Set up third-party runtime integrations or the Snyk Runtime Sensor to get even more runtime data.
6. To ensure you have properly set up the prioritization capability, navigate to the **Set up Snyk AppRisk** tab on the **Issues** page and view the data Snyk has access to.\
   You can also filter relevant sections by Organization for a granular view of your progress.

{% hint style="info" %}
The Insights page, from the **Set up Insights** tab, is available only for Snyk AppRisk users.
{% endhint %}

## Set up Insights UI settings

The Insights settings are organized into three main categories:

* [Risk factors](./#risk-factors) - allows you to enable or disable the risk factors involved in calculating the risk of your identified issues.&#x20;
* [Provider selection](./#provider-selection) - allows you to configure the Kubernetes runtime providers.
* [Kubernetes cluster mapping](./#kubernetes-cluster-mapping) - allows you to match the clusters identified by providers with the preferred names.

All these settings can be found in the Snyk Web UI, under Group Settings, the Settings option.

### Risk factors

You can enable or disable the available risk factors: [Deployed](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md), [Loaded package](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md), [OS condition](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-os-condition.md), [Public facing](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-public-facing.md). When a risk factor is disabled, it will not be used to calculate issues.

You can enable or disable the risk factors from Snyk web UI under Group **Settings** > **Settings** option > **Risk factors**.

### Provider selection

You can set up multiple Kubernetes runtime providers to gather relevant runtime risk factors from your current integrations. When multiple Kubernetes runtime providers report the same resources, some details, such as the loaded package, may only be available from one of them.&#x20;

The default provider setting is used when two or more providers report data for the same Kubernetes cluster. When the same deployment is identified, select which one should take priority. When no default provider is selected, the earliest data available is used.

Individual providers can also be disabled here if they should not be used when calculating Insights.

{% hint style="info" %}
Consider setting up a default provider. When a default provider is not configured, Snyk prioritizes data from the first provider added.
{% endhint %}

### Kubernetes cluster mapping

Runtime providers may report different names that refer to the same Kubernetes cluster. You can add a cluster name mapping to let Insights correlate resources from the two data sets.&#x20;

For example, if Snyk reports a cluster name as `dev`, and an integration reports the same cluster name as `dev-foo`, you can add a mapping for the integration with a source name of `dev-foo` and a target name of `dev`.

{% hint style="info" %}
Ensure each source name is assigned to only one cluster mapping.
{% endhint %}

## Prioritize your integrations

When prioritizing issues, it is important to understand the available integration options and associated risk factors.&#x20;

Here are the integration options that you can choose from when setting up issues prioritization. You can customize the settings by navigating to the Group level Snyk web UI, the Setting menu, and then the Insights option.&#x20;

* [Snyk Runtime Sensor](../../../integrations/snyk-runtime-sensor.md): Used for deeper runtime analysis, providing detailed insight into the actual usage and potential vulnerabilities of your applications during execution. This sensor helps in gathering live traffic data, identifying runtime vulnerabilities, and assessing potential risks based on real-time application behavior.
* [Kubernetes Connector](set-up-insights-kubernetes-connector.md): Offers comprehensive monitoring for your Kubernetes deployments. This integration helps identify vulnerabilities within Kubernetes clusters and provides data on workload vulnerabilities, infrastructure misconfigurations, and potential malicious activity.
* [Third-party Integrations](../../../integrations/connect-a-third-party-integration.md): These integrations, such as cloud providers or CI/CD tools, provide additional context and data sources for better vulnerability assessment. They help identify misconfigurations, exposure points, and integration-specific vulnerabilities.

### Risk factors mapped to integration options

Each of these integration options uses different risk factors:

* Snyk Runtime Sensor: **Deployed** and **Loaded Package**
* Kubernetes Connector: **Deployed**, **Public facing**
* Third-party Integrations: **Deployed** and **Loaded Package**

By leveraging these integration options, you can ensure comprehensive coverage and accurate prioritization of security risks.

The **Deployed** and **Loaded Package** risk factors offer more specific and actionable insights into the actual state and behavior of your applications, enabling better prioritization and remediation efforts.

The **Public facing** risk factor is less commonly used in all available integration options because it covers broader, more external threats that are harder to quantify and track within an internal monitoring system.&#x20;
