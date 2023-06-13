# Insights setup

### Prerequisites

To get value from Insights, you **must choose** an application where you are using [Snyk Container](../../../scan-containers/) to scan the images. You can get additional value by also scanning your open source dependencies with [Snyk Open Source](../../../scan-application-code/snyk-open-source/) and your source code with [Snyk Code](../../../scan-application-code/snyk-code/).

Snyk Insights operates by providing you with 3 risk factors for your vulnerabilities:&#x20;

* **OS Condition**: Does this vulnerability apply to my operating system?
* **Deployed**: Is my code and container image deployed anywhere?
* **Public Facing**: Does my container have any internet exposure?

In order to get data about these 3 risk factors, you need to meet the following criteria:

* **\[OS Condition]** Source code & dependencies are being built into a container image and scanned with Snyk Container - this is the minimum requirement to get value from Insights

<figure><img src="../../../.gitbook/assets/Example OS condition.png" alt=""><figcaption></figcaption></figure>

* **\[Deployed & Public Facing]** This container image is deployed to a kubernetes cluster, where you’re able to deploy the [Kubernetes Connector](insights-setup-kubernetes-connector.md) - by making sure these 2 requirements are satisfied, you get data for all 3 risk factors about the code in your scanned image\


Snyk’s **recommendation** is that you also perform the following steps in order to get the maximum value out of Insights:

* Scan the third party dependencies using Snyk Open Source
* Scan the source code using Snyk Code&#x20;

By doing these two actions, both the source code and the third party dependencies will get risk factors data, meaning you get the application context to better prioritize your open issues.

#### Good practice tip

We recommend picking one application to start with, and then expanding from there.&#x20;

## Process overview

1. [User Permissions](insights-setup-user-permissions.md): Granting users the Group Viewer role
2. [Kubernetes Connector](insights-setup-kubernetes-connector.md): Create the required Organization, roles, permissions, and deploy the agent. _Note: this Kubernetes Connector is separate from the Kubernetes Integration (snyk-monitor)_
3. [Image scanning](broken-reference): Scan your images properly so Insights has access to the right data
4. [Associating Snyk Open Source, Code & Container Projects](insights-setup-associating-snyk-open-source-code-and-container-projects.md): set up the required linking for the application you want to use Insights on.



