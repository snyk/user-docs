# Define policies

Policies define how Snyk behaves when identifying issues. Policies give you a quick and automated way to identify, prioritize, and triage issues. This saves valuable development time and allows developers to take more responsibility and ownership for security, reducing the “noise” level.

## Security policies

{% hint style="success" %}
**Key decision:** Decide which conditions automatically increase or decrease the priority or severity of an issue to match your risk appetite, and which specific issues or types of issues are automatically ignored to reduce "noise" and save development time.
{% endhint %}

Group administrators can define security policies, thus providing an automated way to identify certain issues or types of issues, and apply actions like changing the severity or ignoring the issue based on your conditions.

* Configure policies to increase priority or decrease it as needed.
* Create ignores where needed

## License policies <a href="#license-policies" id="license-policies"></a>

{% hint style="success" %}
**Key decision:** Decide which specific license types to explicitly allow or disallow to avoid using packages with incompatible or problematic licenses, and use severity types to decide how to configure these policies to match your specific legal and compliance requirements.
{% endhint %}

Group administrators can set license policies to define Snyk behavior for treating license issues. For example, you can allow or disallow packages with certain license types, to avoid using packages containing incompatible licenses.

By default, Snyk determines the severity of licenses as follows:

* **High severity**: licenses that definitely present issues for commercial software.
* **Medium severity**: licenses that have clauses that may be of concern and should be reviewed.

Configure policies to match your requirements.

## Asset policies <a href="#asset-policies" id="asset-policies"></a>

{% hint style="success" %}
**Key decision:** Decide how to automate the governance, tracking, and remediation workflows for your assets to ensure continuous security visibility and compliance
{% endhint %}

Asset policies in Snyk Essentials automate business context and notification workflows. Use policies to identify coverage gaps and manage assets at scale.&#x20;

## About Policies

### Policy components

A policy consists of the following elements:

* Filters: Define criteria (for example tags or asset names) to group specific assets.
* Actions: Define what happens to filtered assets, for example, assigning a classification or sending a Slack notification.&#x20;

### Key filter types

Use the following filters to refine your asset groups:

* **Name**: Match assets based on naming conventions.
* **Asset Type**: Target specific assets like repositories or packages.
* **Class**: Focus on business-critical assets (for example, Class A) to reduce noise.
* **Tags**: Use system or custom-defined tags to identify assets.

### **Policy execution**

* Automatic: Policies run every three hours.
* Manual: Click **Run** in the Policy view to apply changes immediately.

## **Creating a policy**

To create a policy:

1. In the Snyk web UI, navigate to **Policies** > **New policy**.
2. Enter a **Name** and a **Description**, then click **Next**.
3. In the policy builder, define your **Filters** and click **Apply**.
4. Click the **+** icon to **Set actions**.
   * Trigger an action: Apply a change or notification to filtered assets.
   * Logic node: Combine multiple filters before triggering an action.
5. Click **Save**.

## **Common use cases**

* **Coverage control**: Define where specific security controls must be active.
* **Classification**: Categorize assets by business importance.
* **Tagging**: Apply consistent metadata to matched assets.
* **Notification**: Alert teams through email or Slack when asset states change.
