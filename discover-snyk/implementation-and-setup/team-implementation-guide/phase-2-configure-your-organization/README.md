---
description: 'Phase 2 of the Snyk team implementation: name and configure your Organization and its integrations'
nav_context: new
---

# Phase 2: Configure your Organization

## Name your Organization

This step is easiest performed prior to purchasing the license. If you did not, Snyk support can assist you.

### Changing the name prior to purchase

Open the **Organization** dropdown in the scope selector and click **+ Create new Organization**. Then purchase a Team plan for the new Organization. [ACTION REQUIRED: confirm the new-navigation location for purchasing a Team plan. The classic path was Settings-Plans and Billing.]

### Changing the name after purchase

After you determine the name of your Organization in phase 1, navigate to **Settings** > **Organization settings** > **General**, update the **Organization name** box, and click **Update**.

{% hint style="info" %}
You must contact Snyk support and submit a ticket to update the url slug. This can impact any existing CI/CD scripts that your team has previously created, so this step should be performed as early as possible and be communicated to your team if CLI has been in use for a while.
{% endhint %}

## Additional settings to consider

* Enable Snyk Code, if purchased, in **Settings** > **Products and features** > **Snyk Code**.
* Require a reason to ignore issues in the **Ignores** section of **Settings** > **Organization settings** > **General**.
* Disable notifications at the Organization level in **Settings** > **Organization settings** > **Notifications**. You can re-enable notifications after importing in the later step. This reduces noise during early testing/importing and is more appropriate for a steady state of the implementation.

### License policy (optional)

Snyk comes with a default policy. You can find this under **Settings-Licenses** [ACTION REQUIRED: confirm the new-navigation location for license policies. No Licenses entry appears in the Organization Settings menu.] if your team wants to add custom text when an issue is found or change severities for specific licenses.

Administrators can set license policies to define Snyk behavior for treating license issues. For example, you can allow or disallow packages with certain license types to avoid using packages containing incompatible licenses.

By default, Snyk determines the severity of licenses as follows:

* High severity - licenses that present issues for commercial software.
* Medium severity - licenses that have clauses that may be of concern and should be reviewed.

Configure policies to match your requirements.

For details, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management).
