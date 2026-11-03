---
description: Guidelines for structuring Snyk accounts, Groups, Organizations, and Projects to maintain high performance at scale
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Structure your account for high application performance

For the best experience with Snyk, follow these guidelines when you make decisions about your Snyk user accounts, Groups, Organizations, and Projects.

## Structure of accounts for users

You can have a large number of users in your Organizations and Groups.

Snyk recommends structuring your Organizations so that there are no more than 2,000 users in each Organization.

If you have more than 2,000 users in an Organization, you begin to risk performance issues. When the application must load a large number of users, the dashboard and the Group members management page slow down.

If users have a large number of memberships in a given Group, all requests — in the Snyk Web UI and through the CLI and the API — slow down. On most requests, Snyk runs calculations and queries to check access and permissions.

## Structure of Groups

A small number of Snyk customers have more than one Group, for example, to keep different business units completely separate. However, before you set up your account with multiple Groups, you must understand the limitations.

Specifically, each Group is a standalone entity. This has the following consequences:

* The functionality for Groups is not tied together.
* There is no cross-Group reporting.
* You cannot share users, Projects, or Organizations between Groups.
* SSO is more difficult to manage across multiple Groups.
* Service accounts cannot span multiple Groups.
* Getting data for multiple Groups through the API requires multiple calls.

If you believe your business case requires multiple Groups, contact your Snyk account team for assistance.

## Structure of Organizations

Using either the Snyk Web UI or the Snyk API, you can create a large number of Organizations in your Group. However, if you have more than 2,000 Organizations in your Group, you begin to risk performance issues.

When Snyk loads a large number of Organizations, these consequences follow:

* Performance slows down for Group administrators and Group-level notifications.
* Group-level service account creation can fail.

## Structure of Projects

You can import a large number of Projects into your Organizations.

Snyk recommends limiting each Organization to no more than 10,000 Projects, and does not allow more than 25,000 Projects in each Organization.

If you need more than 10,000 Projects, consider the effect on performance. A large number of Projects slows down listing Projects, sending notifications, displaying the dashboard, and displaying the Usage page. Also, it is difficult to delete Organizations with a large number of Projects.

There is no limit to the overall number of Projects across all Organizations in a Group. However, with several hundred thousand Projects, Group-level reporting and the display of dependencies and license issues slow down.
