---
nav_context: agnostic
last_reviewed: 2026-10-09T00:00:00.000Z
last_reviewed_source: bulk-sync
description: >-
  How to manage Snyk Learn access controls for assignment and report permissions
  across your Organization
---

# Snyk Learn access controls

{% include "../../.gitbook/includes/deprecatedsnyk-is-deprecati....md" %}

{% hint style="info" %}
**Feature availability**

Assignment permissions are available in the Snyk Learning Management add-on offering. Report permissions are available to Snyk Enterprise plan customers. For more information, contact your Snyk account team.
{% endhint %}

With Snyk Learn access controls, you can manage user permissions for Snyk Learn using the same [access control model](https://docs.snyk.io/platform-administration/user-management/user-role-management) as the Snyk platform. A Group Admin can create custom roles with a set of permissions related to Snyk Learn assignments. These roles can reflect users and functions in the Organization. For example, you can limit permissions for education and training management functionality for Snyk Learn to the user who is the education and training manager at your company.

{% hint style="info" %}
Group Admins and Organization Admins have access to Snyk Learn reports and assignments by default, without extra permissions.
{% endhint %}

## Assignments

To grant permissions for access to Assignments beyond the default Organization and Group Admin roles, you must create a [custom role](https://docs.snyk.io/platform-administration/user-management/custom-role-templates/snyk-learn-learning-admin) with a set of permissions for assignments. For assignment management to work, the custom role must have the View Users permission. The following lists the permissions for managing assignments:

* View Organization assignments
* Edit assignments
* Create assignments
* Delete assignments

After you create the role, a Group Admin or Organization Admin can assign it to Snyk users.

For more information, visit [Snyk Learn assignments](snyk-learn-assignments.md).

## Reports

To grant permissions for access to Reports beyond the default Organization and Group Admin roles, you must create a [custom role](https://docs.snyk.io/platform-administration/user-management/custom-role-templates/snyk-learn-learning-admin) with a set of permissions for reports. The custom role must have the View Reports permission for viewing and exporting reports to work. The following lists the permissions for viewing and exporting reports:

* View Organization reports

After you create the role, a Group Admin or Organization Admin can assign it to Snyk users.

For more information, visit [Snyk Learn reporting](snyk-learn-reports/).
