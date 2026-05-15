# CLI Tester role template

This Organization-level role can use `snyk test` and `snyk monitor`.

## Group-level permissions

This template is for an Organization-level role and has no Group-level permissions.

## Organization-level permissions

To create this role, enable the following permissions in the relevant categories:

### Organization management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization</td><td>true</td></tr><tr><td>Edit Organization</td><td>false</td></tr><tr><td>Remove Organization</td><td>false</td></tr></tbody></table>

### Package management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>Test packages</td><td>true</td></tr></tbody></table>

### Project management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project</td><td>true</td></tr><tr><td>Add Project</td><td>true</td></tr><tr><td>Edit Project</td><td>false</td></tr><tr><td>Edit Project status</td><td>false</td></tr><tr><td>Test Project</td><td>true</td></tr><tr><td>Move Project</td><td>false</td></tr><tr><td>Remove Project</td><td>false</td></tr><tr><td>View Project history</td><td>false</td></tr><tr><td>Edit Project integrations</td><td>false</td></tr><tr><td>Edit Project attributes</td><td>false</td></tr><tr><td>View Jira issues</td><td>false</td></tr><tr><td>Create Jira issues</td><td>false</td></tr><tr><td>Edit Project Tags</td><td>false</td></tr></tbody></table>

### Project Ignore management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project Ignores</td><td>true</td></tr><tr><td>Create Project Ignores</td><td>false</td></tr><tr><td>Edit Project Ignores</td><td>false</td></tr><tr><td>Remove Project Ignores</td><td>false</td></tr></tbody></table>

### Snyk Preview management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Snyk Preview features</td><td>true</td></tr><tr><td>Edit Snyk Preview features</td><td>false</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Audit Log management
* Billing management
* Collection management
* Container Image management
* Entitlement management
* Integration management
* Kubernetes integration management
* Project pull request management
* Reports management
* Service account management
* Snyk Apps management
* Snyk Cloud management
* User management
* Webhook management
