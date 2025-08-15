# Legal Advisor role template

This Organization-level role can manage security and license policies for the Group and view and export reports, but not manage the Group or Organizations or view individual Projects.

## Group-level permissions

This template is for an Organization-level role and has no Group-level permissions.

## Organization-level permissions

To create this role, enable the following permissions in the relevant categories:

### Organization management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization</td><td>true</td></tr><tr><td>Edit Organization</td><td>false</td></tr><tr><td>Remove Organization</td><td>false</td></tr></tbody></table>

### Audit log management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View audit logs</td><td>true</td></tr></tbody></table>

### Billing management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View billing</td><td>true</td></tr><tr><td>Edit billing</td><td>false</td></tr></tbody></table>

### Project management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project</td><td>false</td></tr><tr><td>Add Project</td><td>false</td></tr><tr><td>Edit Project</td><td>false</td></tr><tr><td>Edit Project status</td><td>false</td></tr><tr><td>Test Project</td><td>false</td></tr><tr><td>Move Project</td><td>false</td></tr><tr><td>Remove Project</td><td>false</td></tr><tr><td>View Project history</td><td>true</td></tr><tr><td>Edit Project integrations</td><td>false</td></tr><tr><td>Edit Project attributes</td><td>false</td></tr><tr><td>View Jira issues</td><td>false</td></tr><tr><td>Create Jira issues</td><td>false</td></tr><tr><td>Edit Project Tags</td><td>false</td></tr></tbody></table>

### Project Ignore management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project Ignores</td><td>true</td></tr><tr><td>Create Project Ignores</td><td>false</td></tr><tr><td>Edit Project Ignores</td><td>false</td></tr><tr><td>Remove Project Ignores</td><td>false</td></tr></tbody></table>

### Reports management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization reports</td><td>true</td></tr></tbody></table>

### User management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View users</td><td>true</td></tr><tr><td>Invite users</td><td>false</td></tr><tr><td>Manage users</td><td>false</td></tr><tr><td>Add users</td><td>false</td></tr><tr><td>Provision users</td><td>false</td></tr><tr><td>User leave</td><td>false</td></tr><tr><td>User remove</td><td>false</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Collection management
* Container Image management
* Entitlement management
* Integration management
* Kubernetes integration management
* Package management
* Project pull request management
* Service account management
* Snyk Apps management
* Snyk Cloud management
* Snyk Preview management
* Webhook management
