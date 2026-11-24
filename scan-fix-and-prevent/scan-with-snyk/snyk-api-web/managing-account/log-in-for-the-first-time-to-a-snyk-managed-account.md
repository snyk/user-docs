---
description: Set up your Organizations and fix first-login problems in a Snyk-managed account
nav_context: agnostic
last_reviewed: 2026-11-19
last_reviewed_source: reviewed
---

# Log in for the first time to a Snyk-managed account

{% hint style="info" %}
This page applies to Snyk-managed accounts. For details, visit [Understand Snyk-managed accounts](snyk-managed-accounts.md).
{% endhint %}

The first time anyone logs in to a Snyk-managed account, Snyk API & Web creates an Organization for every Snyk Organization in your Tenant. A Snyk Tenant Admin must make that first login.

## Log in for the first time

1. As a Tenant Admin, click **Log in with Snyk**.
2. A screen with a progress bar displays the following message: **We are creating your Snyk tenant's organizations here. This might take a while. The app opens once the syncing is complete. Please keep this page open.**
3. Keep the page open. When the sync ends, the app opens.

After this, other users can log in. Snyk API & Web creates each user the first time they log in.

## Later logins

Each login adds the user to their Organizations and creates any Organizations that are missing. While permissions load, the note **Loading your permissions. This may take a few seconds.** appears. You can keep working. Organizations you add, rename, or delete in Snyk can take up to 5 minutes to reach Snyk API & Web.

## Troubleshooting

### A user sees a message asking a Tenant Admin to log in

The full message is: **A Tenant admin has to log in for your Snyk tenant's organizations to be created here. Ask a Tenant admin of your Snyk account to log in to Snyk API & Web, then log in again.**

Nobody has completed the first login yet, and only a Tenant Admin can. Ask a Tenant Admin to log in, then log in again.

### The first sync failed

If the sync failed, click **Retry** to sync again. If you do not retry, your next Tenant Admin login or the next daily update completes the sync. Until then, users cannot browse Groups or Organizations.

### An Organization is missing

A new Organization created in Snyk can take up to 5 minutes to appear. Open the Group again after that. Organizations you are not a member of do not appear. If a user still cannot see an Organization they belong to after 5 minutes, ask them to log in again.

### The Organization list looks out of date

Changes made in Snyk can take up to 5 minutes to reach Snyk API & Web. After that, they show at your next login, or when you next open the Organization. For details, visit [Understand Snyk-managed accounts](snyk-managed-accounts.md).

If a problem persists, contact Snyk Support.

### No permissions error

When you perform an action while viewing all Organizations in a Group, you can get an error message saying that you do not have permission. The exact text depends on the action. This happens when you do not have permission to perform that action, or when permissions have not finished syncing for that Organization. Wait up to 5 minutes before trying again.

If a problem persists, contact Snyk Support.

## Related information

* [Understand Snyk-managed accounts](snyk-managed-accounts.md)
* [Select a Group and an Organization](select-a-group-and-an-organization.md)
