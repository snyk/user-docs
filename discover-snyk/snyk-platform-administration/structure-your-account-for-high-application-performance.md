# Structure your account for high application performance

To ensure the best experience using Snyk, consider these guidelines when making decisions about your Snyk user accounts, Groups, Organizations, and Projects.

## Structure of accounts for users

You can have a large number of users in your Organizations and Groups.

Snyk recommends structuring your Organizations so that there are no more than 2,000 users in each Organization.

If you have more than 2,000 users in an Organization, you begin to risk performance issues. When the application must load a large number of users, performance is slowed for the dashboard and the Group members management page.

If users have a large number of memberships in a given Group, all requests--in the Snyk Web UI and through the CLI and the API, are slowed because, on most requests, calculations and queries occur to check access and permissions.

## Structure of Groups

A small number of Snyk customers have more than one Group, for example, to keep different business units completely separate. However, anyone considering multiple Groups must understand the limitations of setting up their account with multiple groups.

Specifically, each Group is a standalone entity. This has the following consequences:

* The functionality for Groups is not tied together.
* There is no cross-Group reporting.
* Users, Projects, and Organizations cannot be shared between Groups.
* SSO is more difficult to manage across multiple Groups.
* Service accounts cannot span multiple Groups.
* Getting data for multiple Groups through the API requires multiple calls.

If you believe your business case requires multiple Groups, consult with your Snyk account team for assistance.

## Structure of Organizations

Using either the Snyk Web UI or the Snyk API, you can create a large number of Organizations in your Group. However, if you have more than 2,000 Organizations in your Group, you begin to risk performance issues.

When Snyk loads a large number of Organizations, these consequences follow:

* Performance is slowed for Group administrators and Group-level notifications.
* Group-level service account creation may fail.

## Structure of Projects

You can import a large number of Projects into your Organizations.

Snyk recommends limiting each Organization to **no more than 10,000 Project**s, and does not allow more than 25,000 Projects in each Organization.

If you need more than 10,000 Projects, consider how a large number of Projects affects the experience with slower performance for listing Projects, sending notifications, displaying the dashboard, and displaying the Usage page. Also, it is difficult to delete Organizations with a large number of Projects.

While there is no limit to the overall number of Projects across all Organizations in a Group, Group-level reporting and displaying dependencies and license issues are slowed when there are several hundred thousand Projects.
