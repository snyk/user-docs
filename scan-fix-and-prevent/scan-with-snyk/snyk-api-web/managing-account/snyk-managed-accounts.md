---
description: How access works for Snyk-managed accounts in Snyk API & Web
nav_context: agnostic
last_reviewed: 2026-11-19
last_reviewed_source: reviewed
---

# Understand Snyk-managed accounts

If your Snyk API & Web account is managed in Snyk, your users, roles, and Organizations come from Snyk instead of being administered in Snyk API & Web. This page explains how the two line up and where each setting is managed.

## Check whether your account is managed in Snyk

Your account is managed in Snyk if a bar at the top of your pages shows a Tenant, a Group, and an Organization. If you do not see these, your account is not managed in Snyk, and the other pages in this section apply as written.

## How your Snyk structure maps to Snyk API & Web

| In Snyk      | In Snyk API & Web                                                         |
| ------------ | ------------------------------------------------------------------------- |
| Tenant       | Account. Shown in the bar, read-only.                                     |
| Group        | Group. Read from Snyk and used to choose Organizations.                   |
| Organization | Organization. Targets belong to an Organization (previously called Team). |

The first time a Tenant Admin logs in, Snyk API & Web creates an Organization for each Snyk Organization in your Tenant, named after it. After that, new, renamed, and deleted Organizations in Snyk can take up to 5 minutes to reach Snyk API & Web. If you delete an Organization in Snyk while it still has targets in Snyk API & Web, you can no longer access those targets. You do not create, rename, or delete Organizations in Snyk API & Web. Snyk API & Web ignores Organizations outside the Tenant linked to your account.

## Log in

Users log in with **Log in with Snyk**. They have no separate Snyk API & Web password and no two-factor authentication in Snyk API & Web, because Snyk handles authentication.

You do not add users in advance. Snyk API & Web creates a user the first time that person logs in. The first person to log in to the account must be a Snyk Tenant Admin. For details, visit [Log in for the first time to a Snyk-managed account](log-in-for-the-first-time-to-a-snyk-managed-account.md).

## How permissions are applied

Snyk API & Web reads the roles of each user from Snyk:

* Tenant role: at each login.
* Organization role: every time the user opens that Organization. A Group role applies to every Organization in the Group. If a user has both, the role that grants more applies.

A change made in Snyk can take up to 5 minutes to reach Snyk API & Web. After that, a role change shows up at the next login of the user, or the next time they open the Organization. The same applies to Groups and Organizations. Removing a user in Snyk does not end a session already in progress.

To see which Snyk role grants which permission, visit [Map Snyk roles to Snyk API & Web permissions](snyk-role-permission-mapping.md).

Because Snyk API & Web reads roles per Organization, you work in one Organization at a time. For details, visit [Select a Group and an Organization](select-a-group-and-an-organization.md).

## What you manage in Snyk

For Snyk-managed accounts, you manage the following in Snyk:

* Users
* Roles
* Organizations
* Authentication

## What is unchanged

You continue to manage the following in Snyk API & Web:

* Targets. Every target belongs to an Organization, so **Access Scope** is required when you add or edit a target.
* API keys. A key can hold roles in one or several Organizations of a Group. Roles come from Snyk, and you cannot scope a key to a single target. For details, visit [Generate API key](../integrations/generate-api-key.md).
* API key policies, which apply to the whole account. Find them on the **API Keys** page, in **Settings** > **API Keys**.
* Scanning agents, including restricting an agent to specific Organizations.

Accounts that are not managed in Snyk are unaffected.

## Related information

* [Log in to Snyk API & Web](log-in-to-snyk-api-web.md)
* [Select a Group and an Organization](select-a-group-and-an-organization.md)
* [Log in for the first time to a Snyk-managed account](log-in-for-the-first-time-to-a-snyk-managed-account.md)
* [Understanding permissions](understanding-permissions.md)
* [Get started with teams](get-started-with-teams.md)
