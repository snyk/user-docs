---
description: Considerations for groups, organizations, projects, and users
---

# Structure account for high application performance

To ensure the best experience using Snyk, there are several guidelines to consider when making decisions about your Snyk Groups, Organizations, Projects, and users.

### Groups

A small number of Snyk customers have more than one Group, when required for specific reasons (like wanting to keep different business units completely separate). However, anyone considering multiple Groups needs to understand the limitations of setting up their account in that way.

#### Multiple Groups - limitations

Each Group is a stand-alone entity. This means:

* The functionality for Groups is not tied together
* There is no cross-Group reporting
* Users, Projects, and Organizations cannot be shared between Groups
* SSO is more difficult to manage across multiple Groups
* Service accounts cannot span multiple Groups
* Getting data for multiple Groups via the API requires multiple calls.

If your business case calls for multiple groups, work with your Snyk account team.

### Organizations

Using either the Snyk web app or the API, you can create a large number of Organizations in your Group. However, if you have more than 2,000 organizations in your Group, you begin to risk performance issues.

#### Large number of Organizations - limitations

When Snyk loads a high number of entities, this means:

* Performance is slowed for Group administrators and Group-level notifications
* Group-level service account creation may fail

### Projects

You can import a large number of projects to your organizations.

We recommend limiting each organization to no more than 10,000 projects, and we do not allow more than 25,000 projects per organization.

#### Large numbers of Projects - limitations

If you’ll need more than 10,000 Projects, consider how a large number of Projects affects the experience with slower performance for listing projects, notifications, the Dashboard and the Usage page. Currently, it is difficult to delete Organizations with a large number of Projects.

While there is no limit to the overall number of Projects across all Organizations in a Group, Group-level reporting (especially the Dependencies and Licenses tabs) is slower when there are several hundred thousand projects.

### Users

You can have a large number of users in your organizations and groups.

We recommend structuring your organizations so that there are not more than 2,000 users per organization.

#### Large number of users - limitations

If you have more than 2,000 users in an organization, you begin to risk performance issues. When the application must load a high number of users, this means performance is slowed for the dashboard and the group members management page.

If users have a number of memberships in a given Group, all requests (in the Snyk web app, via the API, or in the CLI) are slowed as calculations and queries occur on most requests to check access and permissions.
