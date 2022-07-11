---
description: Considerations for groups, organizations, projects, and users
---

# Structure account for high application performance

To ensure the best experience using Snyk, there are several guidelines to consider when making decisions about your groups, organizations, projects, and users.

Snyk is actively working on improving these performance issues.

### Groups

Groups can hold many organizations and group members. We recommend limiting your account to 1 group.

A small number of Snyk customers have more than 1 group, when required for specific reasons (like wanting to keep different business units completely separate). However, anyone considering multiple groups needs to understand the limitations of setting up their account in that way.

Each group is a stand alone entity. This means:

* the functionality for groups is not tied together at this time
* there is no cross-group reporting
* users, projects, and organizations cannot be shared between groups&#x20;
* SSO is more difficult to manage across multiple groups
* service accounts cannot span multiple groups

Getting data for multiple groups via the API requires multiple calls.

If your business case calls for multiple groups, work with your Customer Success Manager.

### Organizations

Using either the Snyk web app or the API, you can create a large number of organizations within your group.

However, if you have more than 2000 organizations in your group, you begin to risk performance issues. When the application must load a high number of entities, this means:

* performance is slowed for group administrators and group-level notifications&#x20;
* group-level service account creation may fail

### Projects

You can import a large number of projects to your organizations.

We recommend limiting each organization to no more than 10,000 projects, and we do not allow more than 25,000 projects per organization.

If you’ll need more than 10,000 projects, consider how a large number of projects affects the experience with slower performance for listing projects, notifications, the Dashboard and the Usage page. Currently, it is difficult or even impossible to delete organizations with a large number of projects.

While there is no limit to the overall number of projects across all organizations in a group, we are aware that group-level reporting (especially the Dependencies and Licenses tabs) is slower when there are several hundred thousand projects.

### Users

You can have a large number of users in your organizations and groups.

We recommend structuring your organizations so that there are not more than 2000 users per organization.

If you have more than 2000 users in an organization, you begin to risk performance issues. When the application must load a high number of users, this means performance is slowed for the dashboard and the group members management page.

If users have a number of memberships in a given group, all requests (in the Snyk web app, via the API, or in the CLI) are slowed as calculations and queries occur on most requests to check access and permissions.
