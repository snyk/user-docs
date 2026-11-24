---
description: Choose where you work in a Snyk-managed account
nav_context: agnostic
last_reviewed: 2026-11-19
last_reviewed_source: reviewed
---

# Select a Group and an Organization

{% hint style="info" %}
This page applies to Snyk-managed accounts. If the bar at the top of your pages does not show a Tenant, Group, and Organization, your account is not managed in Snyk. For details, visit [Understand Snyk-managed accounts](snyk-managed-accounts.md).
{% endhint %}

In a Snyk-managed account, you work in one Organization at a time. You choose a Group first, then an Organization in that Group. Pages such as **Dashboard**, **Targets**, **Findings**, **Scans**, and **Users** show only the Organization you chose.

## The bar

A bar at the top of every page shows where you are, as Tenant / Group / Organization:

* Tenant: your account. It is always read-only.
* Group: a dropdown of your Groups. If you belong to one Group, it shows as text.
* Organization: a dropdown of the Organizations in the chosen Group. If you belong to one Organization, it shows as text.

Pages that depend only on a Group, such as **Settings** > **Organizations** and **API Keys**, show the Tenant and Group. Account-level pages, such as **Discovery** and **Account**, show only the Tenant.

## Choose an Organization

1. Log in. The last Organization you visited in this browser is preselected.
2. If no Organization is remembered, the **Choose a Group and an Organization** screen appears. In **Group**, select a Group.
3. In **Organization**, select an Organization. The list shows only the Organizations of the chosen Group, and it is disabled until you choose a Group.

The page you were trying to open appears, or **Targets** if you came from the login.

In either dropdown, type to search by name, and scroll to load more.

## Switch Organization or Group

Use the dropdowns in the bar. When you switch:

* Lists and overviews stay on the same page for the new Organization.
* Anything inside an Organization, such as a target, finding, or scan, returns to the list for the new Organization.
* Filters, search text, page number, and row selection reset. Page size is kept.
* Scans and reports already running continue.
* If you have unsaved changes, Snyk API & Web asks you to choose **Discard and switch** or **Stay on this page**.

When you choose another Group, Snyk API & Web asks you to choose an Organization in that Group.

## View all Organizations in a Group

If you have a role on a whole Group in Snyk, **All organizations** appears at the top of the Organization list. Select it to view targets, findings, scans, and more across every Organization in that Group.

In this view:

* Lists show an **Organization** column and an **Organization** filter, where you can select one or more Organizations.
* Actions look available, because only your Group permissions are loaded. If you do not have permission in the Organization of an item, Snyk API & Web refuses the action, and the message names the Organization.
* **Save report** saves the report for the Group.

No view across the whole account exists.

## If you cannot see an Organization

* **You're not a member of any organization**: you have no Organization in Snyk. Organization pages stay empty, and account-level pages still work. Ask your Snyk admin for access.
* **You don't have access to this group/organization.**: your access was removed in Snyk, or a link points to an Organization you cannot open. Select another one.
* **Discovery** requires a Tenant-level role, even if you have a Group role.
* **A new Organization is not listed**: Organizations created in Snyk can take up to 5 minutes to appear. Open the Group again after that.

## Related information

* [Understand Snyk-managed accounts](snyk-managed-accounts.md)
* [Understanding permissions](understanding-permissions.md)
