# Developer role template

This Organization-level role enables review of scan results, fixing issues, and initiating Project tests. Users with this role can view Organizations and Projects.

Often, when deploying Snyk, developers may have the ability to override Snyk PR checks, but this permission can be revoked after developers are comfortable using the Snyk IDE extensions and start fixing issues earlier in the SDLC. Similarly, you may start by allowing developers to add Projects and then limit that permission to a Team Lead.

## Group-level permissions

This template is for an Organization-level role and has no Group-level permissions.

## Organization-level permissions

To create this role, enable the following permissions in the relevant categories:

### Organization management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization</td><td>true</td></tr><tr><td>Edit Organization</td><td>false</td></tr><tr><td>Remove Organization</td><td>false</td></tr></tbody></table>

### Project management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project</td><td>true</td></tr><tr><td>Add Project</td><td>true</td></tr><tr><td>Edit Project</td><td>true</td></tr><tr><td>Edit Project status</td><td>true</td></tr><tr><td>Test Project</td><td>true</td></tr><tr><td>Move Project</td><td>true</td></tr><tr><td>Remove Project</td><td>true</td></tr><tr><td>View Project history</td><td>true</td></tr><tr><td>Edit Project integrations</td><td>false</td></tr><tr><td>Edit Project attributes</td><td>true</td></tr><tr><td>View Jira issues</td><td>true</td></tr><tr><td>Create Jira issues</td><td>true</td></tr><tr><td>Edit Project Tags</td><td>true</td></tr></tbody></table>

### Project Ignore management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project Ignores</td><td>true</td></tr><tr><td>Create Project Ignores</td><td>true</td></tr><tr><td>Edit Project Ignores</td><td>true</td></tr><tr><td>Remove Project Ignores</td><td>true</td></tr></tbody></table>

### Project pull request management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>Create pull requests</td><td>true</td></tr><tr><td>Mark pull request checks as successful</td><td>true</td></tr></tbody></table>

### Snyk Cloud management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View environments</td><td>false</td></tr><tr><td>Create environments</td><td>false</td></tr><tr><td>Delete environments</td><td>false</td></tr><tr><td>Update environments</td><td>false</td></tr><tr><td>View scans</td><td>true</td></tr><tr><td>Create scans</td><td>true</td></tr><tr><td>View resources</td><td>false</td></tr><tr><td>View artifacts</td><td>false</td></tr><tr><td>Create artifacts</td><td>false</td></tr><tr><td>View Custom Rules</td><td>false</td></tr><tr><td>Create Custom Rules</td><td>false</td></tr><tr><td>Edit Custom Rules</td><td>false</td></tr><tr><td>Delete Custom Rules</td><td>false</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Audit Log management
* Billing management
* Collection management
* Container Image management
* Entitlement management
* Integration management
* Kubernetes Integration management
* Package management
* Reports management
* Service account management
* Snyk Apps management
* Snyk Preview management
* User management
* Webhook management
