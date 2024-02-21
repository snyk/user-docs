# Prioritization setup

## Prerequisites for Risk-based Prioritization

To get value from Snyk AppRisk Essentials, you should use an application where you are using [Snyk Container](../../../../scan-using-snyk/snyk-container/) to scan the images. You can get additional value by also scanning your open-source dependencies with [Snyk Open Source](../../../../scan-using-snyk/snyk-open-source/) and your source code with [Snyk Code](../../../../scan-using-snyk/snyk-code/).

Snyk risk-based prioritization product operates by providing you with three risk factors for your vulnerabilities:&#x20;

* **OS condition**: Does this vulnerability apply to my operating system?
* **Deployed**: Is my code and container image deployed anywhere?
* **Public facing**: Does my container have any internet exposure?

To get data about these three risk factors, you must meet the following criteria:

* **OS condition:** Source code and dependencies are being built into a container image and scanned with Snyk Container. This is the minimum requirement to get value from the prioritization.

<figure><img src="../../../../.gitbook/assets/Example OS condition.png" alt="Source code and dependencies built into a container image"><figcaption><p>Source code and dependencies built into a container image</p></figcaption></figure>

* **Deployed and Public facing:** This container image is deployed onto a Kubernetes cluster, where you can deploy the [Kubernetes Connector](prioritization-setup-kubernetes-connector.md).

By ensuring these two requirements are satisfied, you get data for all three risk factors for the code in your scanned image

Snyk **recommends** that you also perform the following steps to get the maximum value out of Insights:

* Scan the third-party dependencies using Snyk Open Source,
* Scan the source code using Snyk Code,

By scanning both the source code and the third-party dependencies, you will get risk factors data, which provides the application context to better prioritize your open issues.

{% hint style="info" %}
Snyk recommends starting with one application and expanding from there.
{% endhint %}

## Prioritization process overview

The major steps in setting up risk-based prioritization are as follows:

1. Grant users the Group Viewer role or the Organization Collaborator role. See [Prioritization setup: User permissions](prioritization-setup-user-permissions.md).
2. Create the required Organization, roles, and permissions, and deploy the agent. See [Prioritization setup: Kubernetes Connector](prioritization-setup-kubernetes-connector.md).\
   &#x20;Note: The Kubernetes Connector is different from the Kubernetes Controller, Snyk-Monitor.
3. Scan your images properly so Snyk has access to the right data. See [Prioritization setup: Image scanning](prioritization-setup-image-scanning.md).
4. Set up the required linking for the application on which you want to use risk-based prioritization. See [Prioritization setup: Associating Snyk Open Source, Code, and Container Projects](prioritization-setup-associating-snyk-open-source-code-and-container-projects.md).
5. To ensure you have properly set up the prioritization capability, navigate to the **Set up Snyk AppRisk** tab on the **Insights** page and view the data Snyk has access to.\
   You can also filter relevant sections by Organization for a granular view of your progress.



