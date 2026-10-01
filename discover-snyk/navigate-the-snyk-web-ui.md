---
description: >-
  Where to find features in the new Snyk interface, with a mapping from the
  classic navigation.
---

# Navigate the Snyk Web UI

Snyk introduces a new navigation with a unified side menu, a top scope selector for **Tenant**, **Group**, and **Organization**, **Navigation search**, and **Dark Mode**. Use this page to find where features live in the new interface and to map paths from the classic navigation.

## Side menu

The side menu groups the main areas of Snyk:

* **Analytics:** dashboards and key performance indicators.
* **Inventory:** a unified view of your assets.
* **Projects:** the Snyk Projects list, scoped to the current Organization.
* **Issues:** vulnerabilities and license issues across your Snyk Projects.
* **Policies:** Snyk policies and rules.
* **Settings:** a unified hub for Organization and Group settings, security, integrations, and plan management.

**Notifications**, **Help**, and **More from Snyk** details are all available from the main navigation across all menus in the Snyk Web UI.

## Scope selector

The top scope selector replaces the classic sidebar headers for **Tenant**, **Group**, and **Organization**. It contains three dropdowns:

* **Tenant**: switch between the Tenants you can access.
* **Group**: switch between the Groups in the selected Tenant.
* **Organization**: switch between the Organizations in the selected Group. Select All Organizations to view the Group-level context.\
  \
  Each dropdown is searchable. To create an **Organization**, open the **Organization** dropdown and select **+ Create new Organization**.

### Tenant scope

* **Security and access:** Members, Capability Access. View and manage all users in the Tenant, and control which capabilities are available across the Tenant. Assign Tenant-level roles: **Tenant Admin**, **Tenant Viewer**, or **Tenant Member**.
* **Plan and billing:** Your Plan, Credit breakdown, including contract details, licensed capabilities, and credit usage (Enterprise plans only).

### Group scope

* **Group settings:** General, Notifications.
* **Security and access:** SSO, Member roles, Service accounts, Members.
* **Products and features:** Snyk Agent Fix, Snyk Open Source, Snyk Code, Snyk Assist, and other licensed products.
* **Plan and billing:** Your plan and billing, Available plans.
* **Integrations:** General, Snyk Broker, All integrations.

### Organization scope

* **Organization settings:** General, Service accounts, Notifications, Automated collections.
* **Security and access:** Members.
* **Products and features:** Snyk Open Source, Snyk Code, Snyk Container, Snyk IaC, Snyk Assist, and other licensed products.
* **Integrations:** General, Snyk Broker, Authorized Snyk Apps, All integrations, and individual integrations such as ECR and GitHub.
* **Snyk Preview:** enable controls for preview features.

## Where things moved

The following classic items now live elsewhere in the new interface.

| Classic navigation                    | New location                                                | How to get there                                                                                 |
| ------------------------------------- | ----------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| Organizations                         | Scope selector > **Organization** dropdown                  | Open the **Organization** dropdown, then select an Organization or **+ Create new Organization** |
| Tenant / Group / Organization headers | Scope selector                                              | Use the three dropdowns from left to right                                                       |
| Dependencies                          | **Analytics** > **Reports** > **Dependencies and licenses** | Navigate to **Analytics** > **Reports**, or press **⌘K** and enter `dependencies`                |
| Ignore requests                       | **Issues** > **Ignore requests**                            | Navigate to **Issues**, then select the **Ignore requests** tab                                  |
| Cloud                                 | **More from Snyk** > **Cloud**                              | Select **More from Snyk** at the bottom of the side menu, then select **Cloud**                  |
| Custom rules                          | **Settings** > **Snyk Code** > **Rule Extensions**          | Set the scope selector to a Group, then navigate to **Settings** > **Snyk Code**. Snyk manages Rule Extensions at the Group level |
| Integrations                          | **Settings** > **Integrations**                             | Navigate to **Settings** > **Integrations**                                                      |
| Members                               | **Settings** > **Security and access** > **Members**        | Navigate to **Settings** > **Security and access** > **Members**                                 |
| Product updates                       | Side menu                                                   | Click **Product updates** in the side menu                                                       |
| Help                                  | Side menu                                                   | Click **Help** in the side menu                                                                  |

{% hint style="info" %}
The Organization **Dashboard** from the classic navigation has no equivalent in the new navigation.
{% endhint %}

New areas without a direct classic equivalent:

| New area              | Where it is                                     | What it does                                                         |
| --------------------- | ----------------------------------------------- | -------------------------------------------------------------------- |
| **Projects**          | Side menu                                       | Top-level entry to the Snyk Projects list                            |
| **Navigation search** | Opens over any page                             | Jump to any page by name using **⌘K** or **Ctrl+K**                  |
| **Inventory**         | Side menu                                       | Unified view of your assets, starting with Container Images and SBOM |
| **Snyk Assist**       | Top-right corner, next to **Navigation search** | In-product AI assistant for product, security, and account questions |

## Switch between new and classic navigation

You can return to the classic navigation at any time.

1. Open your account menu.
2. Select **Switch new navigation off**.

To return to the new navigation, open the same menu and select **Switch new navigation on**.

## What stays the same

* Your **Projects**, **Issues**, **Integrations**, and **Settings** values do not change under the new navigation.
* API tokens, service accounts, and Snyk CLI behavior are unaffected.
* Snyk continues to release features to both interfaces during the transition period.

## Snyk Assist

{% hint style="info" %}
**Release** **status**\
\
Snyk Assist is in Early Access and available only on the Snyk Enterprise plan. Snyk Assist is not available under classic navigation. Group Admins manage Snyk Assist from **Settings** > **Products and features**.
{% endhint %}

Snyk Assist is an in-product AI assistant. To open Snyk Assist, click the **Snyk Assist** icon in the top-right corner, next to **Navigation search**. The Snyk Assist panel opens on the right.

Snyk Assist can help with:

* security questions and best practices
* Snyk product and feature guidance
* information about your Organization and Group, such as issues, Projects, and integrations
* known vulnerabilities for a package or package version
* support ticket creation
* feature requests for capabilities Snyk does not support

Snyk Assist does not see your screen. It returns only what you can already see in the Snyk Web UI, and cannot change settings, run scans, or modify Projects on your behalf.

Snyk Assist is scoped to your current Organization and Group. It cannot see across your Tenant or compare Organizations.

Snyk Assist runs on Google Vertex AI. Snyk does not use your prompts or responses to train foundation models.

{% hint style="info" %}
Do not include confidential or sensitive information in your questions. For details, see the [Snyk Assist disclaimer](https://snyk.io/policies/snyk-assist-disclaimer).
{% endhint %}

## Dark mode

The Snyk Web UI supports light and dark themes.

1. Open your account menu.
2. Under **Theme**, select the light icon, the dark icon, or the system icon to follow your operating system settings.

Snyk is completing dark mode coverage across the platform. Some areas do not yet render in dark mode.

## Navigation search

Press **⌘K** (macOS) or **Ctrl+K** (Windows and Linux) from anywhere in Snyk to open the **Navigation search**. Enter a page name or keyword to jump directly to that page.

Snyk scopes each result to a level. Each result shows an **Organization** or **Group** tag for the level it applies to. You can move between **Organization** and **Group** settings without leaving the **Navigation search**.
