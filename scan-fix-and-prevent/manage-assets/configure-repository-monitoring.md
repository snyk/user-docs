---
description: How to configure repository monitoring in Snyk
---

# Configure repository monitoring

{% hint style="info" %}
**Release status**

Repository monitoring configuration is in Early Access and available only with Enterprise plans. To enable the feature, visit [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

## Overview

{% hint style="info" %}
To perform bulk monitoring configuration and maintain visibility into the security health of all repositories, you must be a Group Admin.
{% endhint %}

You can manage repository coverage and monitoring configurations centrally across your entire Snyk Group from the **Inventory** page, at Group level. This means you can monitor and manage repositories without navigating between individual Snyk Organizations.

Repository monitoring configuration provides the following capabilities:

* Centralized asset monitoring: view monitoring status for all products, identify health status, and see required actions (such as enabling Snyk Code or resolving SCM integration issues) in one view.
* Bulk import: import repositories directly from the Group **Inventory** page into specific Snyk Organizations.
* On-demand retesting: trigger a retest for specific repositories directly from **Inventory**.
* Actionable error resolution: clear guidance ia available when testing fails due to integration issues or entitlements. After the underlying issue is resolved, testing resumes automatically.

## Configure settings for repository monitoring

You can configure repository monitoring settings for your Projects across multiple repositories simultaneously. These configurations are asset-based. To do this, navigate to **Inventory** > **All Assets**, at the Group level.

{% stepper %}
{% step %}
**Select the repositories to update**

Select the repositories you want to update, then click **Configure monitoring.**
{% endstep %}

{% step %}
**Select the configuration type**

Select one of the following:

* **Create new configuration (override)**: create a new configuration for the repositories you selected and override previous configurations.
* **Update specific settings (partial)**: update the settings for an existing configuration.

Click **Continue**.
{% endstep %}

{% step %}
**Configure settings**

On the configuration page, the following configurations are available:

* **Test frequency**: update the test frequency for Open Source and Code Projects to **Daily**, **Weekly**, or **Never**.

{% hint style="info" %}
It is not possible to manage these settings at Organization level for assets tracked on the **Inventory** page.
{% endhint %}

* **Report to Organization**: map repositories to Snyk Organizations. To avoid duplicate issues, Snyk recommends one Snyk Organization per repository.
* **Testing exclusions (optional)**: manage file exclusions for Snyk Open Source and Snyk Container Projects. Exclusions apply at the asset level. You cannot exclude specific files for the same repository in different Snyk Organizations.

{% hint style="info" %}
For Snyk Code, you can manage exclusions using the `.snyk` files, in order to maintain developer-first workflows.
{% endhint %}

Click **Continue**.
{% endstep %}

{% step %}
**Enable monitoring**

Review your configuration settings and click **Enable monitoring**.
{% endstep %}
{% endstepper %}

### Disable recurring testing

For on-demand testing or to disable recurring testing for specific repositories:

1. Click the asset you want to disable recurring testing for.
2. In the asset card's **Settings** tab, click **Disable recurrent tests** to disable recurrent testing.

To trigger a retest, click **Retest.**
