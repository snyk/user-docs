# Remediator role template

This is an Organization-level role, meaning a Remediator can only view certain areas and functions in Snyk and cannot create PRs, Projects, and more.

A Remediator assists developers in issue assessment and prioritization.

## Group-level permissions

This template is for an Organization-level role and has no Group-level permissions.

## Organization-level permissions

To create this role, enable the following permissions in the relevant categories:

### Organization management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization</td><td>true</td></tr><tr><td>Edit Organization</td><td>false</td></tr><tr><td>Remove Organization</td><td>false</td></tr></tbody></table>

### Project management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project</td><td>true</td></tr><tr><td>Add Project</td><td>false</td></tr><tr><td>Edit Project</td><td>true</td></tr><tr><td>Edit Project status</td><td>false</td></tr><tr><td>Test Project</td><td>true</td></tr><tr><td>Move Project</td><td>true</td></tr><tr><td>Remove Project</td><td>false</td></tr><tr><td>View Project history</td><td>true</td></tr><tr><td>Edit Project integrations</td><td>false</td></tr><tr><td>Edit Project attributes</td><td>false</td></tr><tr><td>View Jira issues</td><td>true</td></tr><tr><td>Create Jira issues</td><td>true</td></tr><tr><td>Edit Project Tags</td><td>true</td></tr></tbody></table>

### Project Ignore management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project Ignores</td><td>true</td></tr><tr><td>Create Project Ignores</td><td>false</td></tr><tr><td>Edit Project Ignores</td><td>false</td></tr><tr><td>Remove Project Ignores</td><td>false</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Audit Log management
* Billing management
* Collection management
* Container Image management
* Entitlement management
* Integration management
* Kubernetes Integration management
* Package management
* Project pull request management
* Reports management
* Service account management
* Snyk Apps management
* Snyk Cloud management
* Snyk Preview management
* User management
* Webhook management
