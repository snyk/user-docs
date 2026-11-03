---
description: 'Phase 2 of the Snyk team implementation: name and configure your Organization and its integrations'
nav_context: new
---

# Phase 2: Configure your Organization

## Name your Organization

This step is easiest to perform before you purchase the license. If you did not do so, Snyk Support can help you.

### Change the name before purchase

Open the **Organization** dropdown in the scope selector and click **+ Create new Organization**. Then purchase a Team plan for the new Organization.

### Change the name after purchase

After you determine the name of your Organization in Phase 1, navigate to **Settings** > **Organization settings** > **General**, update the **Organization name** box, and click **Update**.

{% hint style="info" %}
To update the URL slug, you must contact Snyk Support and submit a ticket. This change can affect any existing CI/CD scripts that your team created, so Snyk recommends performing this step as early as possible. If your team has used the CLI for a while, communicate the change to them.
{% endhint %}

## Additional settings to consider

* Enable Snyk Code, if purchased, in **Settings** > **Products and features** > **Snyk Code**.
* Require a reason to ignore issues in the **Ignores** section of **Settings** > **Organization settings** > **General**.
* Disable notifications at the Organization level in **Settings** > **Organization settings** > **Notifications**. This reduces noise during early testing and importing. You can re-enable notifications after importing, in a later step, when the implementation reaches a steady state.

## License policy (optional)

{% hint style="info" %}
**Feature availability**

License policies are available only with Enterprise plans and apply only to Snyk Open Source scans. For more information, visit [License policies](https://docs.snyk.io/scan-fix-and-prevent/prevent/policies/license-policies).
{% endhint %}

Snyk provides a default license policy. You can find it under **Group** > **Policies** > **License policies** > **Snyk Default License Policy**.

Administrators can set license policies to define Snyk behavior for treating license issues. For example, you can allow or disallow packages with certain license types to avoid using packages containing incompatible licenses.

By default, Snyk determines the severity of licenses as follows:

* High severity: licenses that present issues for commercial software
* Medium severity: licenses with clauses of potential concern that require review

Configure policies to match your requirements.

For details, visit [Getting started with Snyk License Compliance Management](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management).
