---
description: >-
  This page provides a step-by-step guide on how to manage Snyk Learn access
  control.
---

# Snyk Learn access controls

{% hint style="info" %}
**Feature availability**

Assignment permissions are available in the Snyk Learning Management add-on offering. Report permissions are available to Snyk Enterprise plan customers. For more information, contact your Snyk account team.
{% endhint %}

**Snyk Learn Access Controls** functionality enables you to manage user permissions for Snyk Learn using the same [access control model](../../snyk-platform-administration/user-roles/user-role-management.md) as the Snyk Platform. A Group Admin can create custom roles with a set of permissions related to Snyk Learn Assignments. These roles can reflect users and functions in the Organization. For example, you can limit permissions for education and training management functionality for Snyk Learn to the user who is the education and training manager at your organization.

{% hint style="info" %}
Group and Organization Admins have access to Snyk Learn Reports and Assignments by default without extra permissions.
{% endhint %}

#### Assignments

To grant permissions for access to Assignments beyond the default Organization and Group Admin roles, you must create a [custom role](../../snyk-platform-administration/user-roles/custom-role-templates/snyk-learn-learning-admin.md) with a set of permissions for assignments. The custom role must have the View Users permission in order for managing assignments to work. The following lists the permissions for managing assignments:

* View Organization assignments&#x20;
* Edit Organization assignments
* Create Organization assignments
* Delete Organization assignments

After the role is created, a Group or Organization administrator can give this role to users of Snyk.&#x20;

For more information, see [Snyk Learn Assignments](snyk-learn-assignments.md).

#### Reports

To grant permissions for access to Reports beyond the default Organization and Group Admin roles, you must create a [custom role](../../snyk-platform-administration/user-roles/custom-role-templates/snyk-learn-learning-admin.md) with a set of permissions for reports. The custom role must have the View Reports permission for viewing and exporting reports to work. The following lists the permissions for viewing and exporting reports:

* View Organization reports

After the role is created, a Group or Organization administrator can give this role to users of Snyk.&#x20;

For more information, see [Snyk Learn Reports](snyk-learn-reports.md).
