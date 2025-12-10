# Set up Insights

## Prerequisites

Customize prioritization using the Set up Insights option and an application that scans images using [Snyk Container](../../../scan-with-snyk/snyk-container/). You can get additional value by also scanning your open-source dependencies with [Snyk Open Source](../../../scan-with-snyk/snyk-open-source/) and your source code with [Snyk Code](../../../scan-with-snyk/snyk-code/).

### What risk factors do I need?

The Set up insights product operates by providing you with the following risk factors for your vulnerabilities:

* [Deployed](../assets-and-risk-factors/risk-factor-deployed.md): Is my code and container image deployed anywhere?
* [Loaded package](/broken/pages/x5b9XFg7zGtut8eTWBQp): Has a third-party package that is the dependency of an image been loaded?
* [OS condition](../assets-and-risk-factors/risk-factor-os-condition.md): Does this vulnerability apply to my operating system?
* [Public facing](../assets-and-risk-factors/risk-factor-public-facing.md): Does my container have any internet exposure?

### Risk factors usage criteria

To get data about these four risk factors, you must meet the following criteria:

#### **Loaded package risk factor**

You can use the Loaded package risk factor if there is a package that is loaded more often than others poses a higher risk to your application compared to one that is rarely loaded.

#### **OS condition risk factor**

You need to meet the following conditions to use the OS condition risk factor:

* Source code and dependencies are being built into a container image and scanned with Snyk Container. This is the minimum requirement to get value from the prioritization.

<figure><img src="../../../.gitbook/assets/Example OS condition.png" alt="Source code and dependencies built into a container image"><figcaption><p>Source code and dependencies built into a container image</p></figcaption></figure>

#### **Deployed and public facing risk factors**

You need to meet the following conditions to use the Deployed and Public facing risk factors:

* For both risk factors, you have a container image that is deployed onto a Kubernetes cluster, where you can deploy the [Kubernetes Connector](broken-reference/).

Ensure you meet these requirements to gather data for all four risk factors for the code in your scanned image.

### Maximize your insights

Snyk recommends that you also perform the following steps to get the maximum value out of insights:

* Scan the third-party dependencies using Snyk Open Source,
* Scan the source code using Snyk Code
* Start with one application and expand from there

By scanning both the source code and the third-party dependencies, you will get risk factors data, which provides the application context to better prioritize your open issues.

## Prioritization process overview

The major steps in setting up issues are as follows:

1. Grant users the Group Viewer role or the Organization Collaborator role. See [Prioritization setup: User permissions](broken-reference/).
2. Create the required Organization, roles, and permissions, and deploy the agent. See [Prioritization setup: Kubernetes Connector](broken-reference/).

{% hint style="info" %}
The Kubernetes Connector is different from the Kubernetes Controller, Snyk-Monitor.
{% endhint %}

3. Scan your images properly so Snyk has access to the right data. See [Prioritization setup: Image scanning](broken-reference/).
4. Set up the required link for the application on which you want to use prioritization. See [Prioritization setup: Associating Snyk Open Source, Code, and Container Projects](broken-reference/).

## Set up Insights UI settings

The Insights settings are organized into three main categories:

* [Risk factors](./#risk-factors) - allows you to enable or disable the risk factors involved in calculating the risk of your identified issues.
* [Provider selection](./#provider-selection) - allows you to configure the Kubernetes providers.
* [Kubernetes cluster mapping](./#kubernetes-cluster-mapping) - allows you to match the clusters identified by providers with the preferred names.

All these settings can be found in the Snyk Web UI, under Group Settings, the Settings option.

### Risk factors

You can enable or disable the available risk factors: [Deployed](../assets-and-risk-factors/risk-factor-deployed.md), [Loaded package](/broken/pages/x5b9XFg7zGtut8eTWBQp), [OS condition](../assets-and-risk-factors/risk-factor-os-condition.md), [Public facing](../assets-and-risk-factors/risk-factor-public-facing.md). When a risk factor is disabled, it will not be used to calculate issues.

You can enable or disable the risk factors from Snyk web UI under Group **Settings** > **Settings** option > **Risk factors**.

### Provider selection

You can set up multiple Kubernetes providers to gather relevant risk factors from your current integrations. When multiple Kubernetes providers report the same resources, some details, such as the loaded package, may only be available from one of them.

The default provider setting is used when two or more providers report data for the same Kubernetes cluster. When the same deployment is identified, select which one should take priority. When no default provider is selected, the earliest data available is used.

Individual providers can also be disabled here if they should not be used when calculating Insights.

{% hint style="info" %}
Consider setting up a default provider. When a default provider is not configured, Snyk prioritizes data from the first provider added.
{% endhint %}

### Kubernetes cluster mapping

Runtime providers may report different names that refer to the same Kubernetes cluster. You can add a cluster name mapping to let Insights correlate resources from the two data sets.

For example, if Snyk reports a cluster name as `dev`, and an integration reports the same cluster name as `dev-foo`, you can add a mapping for the integration with a source name of `dev-foo` and a target name of `dev`.

{% hint style="info" %}
Ensure each source name is assigned to only one cluster mapping.
{% endhint %}

## Prioritize your integrations

When prioritizing issues, it is important to understand the available integration options and associated risk factors.

Here are the integration options that you can choose from when setting up issues prioritization. You can customize the settings by navigating to the Group level Snyk web UI, the Setting menu, and then the Insights option.

* [Kubernetes Connector](broken-reference/): Offers comprehensive monitoring for your Kubernetes deployments. This integration helps identify vulnerabilities within Kubernetes clusters and provides data on workload vulnerabilities, infrastructure misconfigurations, and potential malicious activity.

By leveraging these integration options, you can ensure comprehensive coverage and accurate prioritization of security risks.

The **Deployed** risk factor offers more specific and actionable insights into the actual state and behavior of your applications, enabling better prioritization and remediation efforts.

The **Public facing** risk factor is less commonly used in all available integration options because it covers broader, more external threats that are harder to quantify and track within an internal monitoring system.
