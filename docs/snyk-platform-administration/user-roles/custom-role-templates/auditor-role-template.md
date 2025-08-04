# Auditor role template

This is a Group-level read-only role, meaning an Auditor can only view certain areas and functions in Snyk and cannot create PRs, Projects, and more.

This role can view issues, results of scans, and reports. An Auditor often verifies that there is a scan snapshot for a particular resource or Snyk Project. The Auditor may be external to the company.

## Group-level permissions

To create this role, enable the following permissions in the relevant categories:

### Group Management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Groups</td><td>true</td></tr><tr><td>Edit Group details</td><td>false</td></tr><tr><td>View Group settings</td><td>false</td></tr><tr><td>Edit settings</td><td>false</td></tr><tr><td>View Group notification settings</td><td>false</td></tr><tr><td>Edit Group notification settings</td><td>false</td></tr></tbody></table>

### Organization management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organizations</td><td>true</td></tr><tr><td>Edit Organizations</td><td>false</td></tr><tr><td>Remove Organizations</td><td>false</td></tr></tbody></table>

### Snyk Essentials management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Snyk Essentials</td><td>true</td></tr><tr><td>Edit Snyk Essentials</td><td>false</td></tr></tbody></table>

### Audit Log management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Audit Logs</td><td>true</td></tr></tbody></table>

### Insights management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>Access Insights</td><td>true</td></tr></tbody></table>

### Reports management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View reports</td><td>true</td></tr></tbody></table>

### Security and License Policies

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Policies</td><td>true</td></tr><tr><td>Create Policies</td><td>false</td></tr><tr><td>Edit Policies</td><td>false</td></tr><tr><td>Delete Policies</td><td>false</td></tr></tbody></table>

### User management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View users</td><td>true</td></tr><tr><td>Invite users</td><td>false</td></tr><tr><td>Manage users</td><td>false</td></tr><tr><td>Add users</td><td>false</td></tr><tr><td>Provision users</td><td>false</td></tr><tr><td>User Leave</td><td>false</td></tr><tr><td>User Remove</td><td>false</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* IaC settings management
* Issue management
* Request access management
* Role management
* Service account management
* Snyk Apps management
* Snyk Preview management
* SSO settings management
* Tags management

## Organization-level permissions

To create this role, enable the following permissions in the relevant categories:

### Organization management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization</td><td>true</td></tr><tr><td>Edit Organization</td><td>false</td></tr><tr><td>Remove Organization</td><td>false</td></tr></tbody></table>

### Audit Log management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View audit logs</td><td>true</td></tr></tbody></table>

### Collection management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Collections</td><td>true</td></tr><tr><td>Create Collection</td><td>false</td></tr><tr><td>Edit Collections</td><td>false</td></tr><tr><td>Delete Collections</td><td>false</td></tr></tbody></table>

### Container Image management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View container image</td><td>true</td></tr><tr><td>Create container image</td><td>false</td></tr><tr><td>Edit container image</td><td>false</td></tr></tbody></table>

### Integration management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View integrations</td><td>true</td></tr><tr><td>Edit integrations</td><td>false</td></tr></tbody></table>

### Project management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project</td><td>true</td></tr><tr><td>Add Project</td><td>false</td></tr><tr><td>Edit Project</td><td>false</td></tr><tr><td>Edit Project status</td><td>false</td></tr><tr><td>Test Project</td><td>false</td></tr><tr><td>Move Project</td><td>false</td></tr><tr><td>Remove Project</td><td>false</td></tr><tr><td>View Project history</td><td>true</td></tr><tr><td>Edit Project integrations</td><td>false</td></tr><tr><td>Edit Project attributes</td><td>false</td></tr><tr><td>View Jira issues</td><td>true</td></tr><tr><td>Create Jira issues</td><td>false</td></tr><tr><td>Edit Project Tags</td><td>false</td></tr></tbody></table>

### Project Ignore management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project Ignores</td><td>true</td></tr><tr><td>Create Project Ignores</td><td>false</td></tr><tr><td>Edit Project Ignores</td><td>false</td></tr><tr><td>Remove Project Ignores</td><td>false</td></tr></tbody></table>

### Reports management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization reports</td><td>true</td></tr></tbody></table>

### Snyk Cloud management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View environments</td><td>false</td></tr><tr><td>Create environments</td><td>false</td></tr><tr><td>Delete environments</td><td>false</td></tr><tr><td>Update environments</td><td>false</td></tr><tr><td>View scans</td><td>true</td></tr><tr><td>Create scans</td><td>false</td></tr><tr><td>View resources</td><td>true</td></tr><tr><td>View artifacts</td><td>true</td></tr><tr><td>Create artifacts</td><td>false</td></tr><tr><td>View Custom Rules</td><td>false</td></tr><tr><td>Create Custom Rules</td><td>false</td></tr><tr><td>Edit Custom Rules</td><td>false</td></tr><tr><td>Delete Custom Rules</td><td>false</td></tr></tbody></table>

### Webhook management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Outbound Webhooks</td><td>true</td></tr><tr><td>Create Outbound Webhooks</td><td>false</td></tr><tr><td>Remove Outbound Webhooks</td><td>false</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Billing management
* Entitlement management
* Kubernetes Integration management
* Package management
* Project pull request management
* Service account management
* Snyk Apps management
* Snyk Preview management
* User management
