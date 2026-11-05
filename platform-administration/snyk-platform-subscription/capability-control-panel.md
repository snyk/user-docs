---
description: >-
  How Tenant Admins enable and disable Snyk capabilities in the capability
  control panel, which controls credit consumption for the Tenant
last_reviewed: 2026-10-30
last_reviewed_source: focused-edit
---

# Enable and disable capabilities

The capability control panel is the menu for Tenant Admins to enable or disable Snyk capabilities.

To enable capabilities:

1. In the Snyk Web UI, open the Tenant settings.
2. Under **Security and access**, select **Capability Access** to open the capability control panel.
3. Click the **Enable** toggle for each capability you want to use.

You can toggle a capability off at any time. For capabilities entitled through credits, this stops credit consumption.

All capabilities default to off for a new customer, so you cannot access most of the Snyk Platform until a Tenant Admin chooses which capabilities to enable. Configure this access as your first step when you onboard as a new customer on the Snyk Platform Subscription.

Settings at the Tenant level determine capability access for all associated Groups and Organizations. This off-by-default policy ensures that Snyk consumes credits only with Tenant Admin permission. Only Tenant Admins can adjust the control panel.

{% hint style="info" %}
If your contract has hybrid licensing, with active subscriptions for both credits and SKU-based licensing, the control panel can toggle only the capabilities that are not subscribed to on SKUs.
{% endhint %}

The **Your Plan** page under **Plan and billing** shows a summary of the capabilities included in your subscriptions, by licensing type. For details, visit [View your plan](viewing-your-plan.md). Capabilities you can toggle in the control panel show **Credits** as their billing model. Any SKU-based subscriptions show their quantity and SKU-level licensing measurement instead.
