---
description: >-
  Where to find features in the new Snyk interface, with a mapping from the
  classic navigation.
---

# Navigate the Snyk Web UI

Snyk introduces a new navigation with a unified side menu, a top scope selector for **Tenant**, **Group**, and **Organization**, **Navigation search**, **Dark Mode**. Use this page to find where features live in the new interface and to map paths from the classic navigation.

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

* **Security and access:** Members. View and manage all users in the Tenant. Assign Tenant-level roles: **Tenant Admin**, **Tenant Viewer**, or **Tenant Member**.
* **Plan and billing:** Your plan and billing, including contract details and licensed capabilities (Enterprise plans only).

### Group scope

* **Group settings:** General, Notifications.
* **Security and access:** SSO, Member roles, Service accounts, Members.
* **Products and features:** Snyk Agent Fix, Snyk Open Source, Snyk Code, and other licensed products.
* **Plan and billing:** Your plan and billing, Available plans.
* **Integrations:** General, Snyk Broker, All integrations.

### Organization scope

* **Organization settings:** General, Service accounts, Notifications, Automated collections.
* **Security and access:** Members.
* **Products and features:** Snyk Open Source, Snyk Code, Snyk Container, Snyk IaC, and other licensed products.
* **Integrations:** General, Snyk Broker, Authorized Snyk Apps, All integrations, and individual integrations such as ECR and GitHub.
* **Snyk Preview:** enable controls for preview features.

## Where things moved

The following classic items now live elsewhere in the new interface.

| Classic navigation                    | New location                                                | How to get there                                                                                 |
| ------------------------------------- | ----------------------------------------------------------- | ------------------------------------------------------------------------------------------------ |
| Organizations (left sidebar)          | Top breadcrumb bar > **Organization** dropdown              | Open the **Organization** dropdown, then select an Organization or **+ Create new Organization** |
| Tenant / Group / Organization headers | Scope selector                                              | Use the three dropdowns from left to right                                                       |
| Dependencies                          | **Analytics** > **Reports** > **Dependencies and licenses** | Navigate to **Analytics** > **Reports**, or press **⌘K** and enter `dependencies`                |
| Integrations                          | **Settings** > **Integrations**                             | Navigate to **Settings** > **Integrations**                                                      |
| Members                               | **Settings** > **Security and access** > **Members**        | Navigate to **Settings** > **Security and access** > **Members**                                 |
| Product updates                       | Notifications bell in the side menu                         | Click the bell at the bottom of the side menu                                                    |
| Help                                  | **Help** at the bottom of the side menu                     | Click **Help** in the side menu                                                                  |

New areas without a direct classic equivalent:

| New area              | Where it is         | What it does                                                         |
| --------------------- | ------------------- | -------------------------------------------------------------------- |
| **Projects**          | Side menu           | Top-level entry to the Snyk Projects list                            |
| **Navigation search** | Opens over any page | Jump to any page by name using **⌘K** or **Ctrl+K**                  |
| **Inventory**         | Side menu           | Unified view of your assets, starting with Container Images and SBOM |

## Switch between new and classic navigation

You can return to the classic navigation at any time.

1. Open your account menu.
2. Select **Switch new navigation off**.

To return to the new navigation, open the same menu and select **Switch new navigation on**.

## What stays the same

* Your **Projects**, **Issues**, **Integrations**, and **Settings** values do not change under the new navigation.
* API tokens, service accounts, and Snyk CLI behavior are unaffected.
* Snyk continues to release features to both interfaces during the transition period.

## Dark mode

The Snyk Web UI supports light and dark themes.

1. Open your account menu.
2. Under **Theme**, select the light icon, the dark icon, or the system icon to follow your operating system settings.

Snyk is completing dark mode coverage across the platform. Some areas do not yet render in dark mode.

## Navigation search

Press **⌘K** (macOS) or **Ctrl+K** (Windows and Linux) from anywhere in Snyk to open the **Navigation search**. Enter a page name or keyword to jump directly to that page.

Snyk scopes each result to a level. Each result shows an **Organization** or **Group** tag for the level it applies to. You can move between **Organization** and **Group** settings without leaving the **Navigation search**.
