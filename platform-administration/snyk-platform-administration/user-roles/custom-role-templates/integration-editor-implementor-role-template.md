# Integration Editor/Implementor role template

This is a Group-level role with integration-related permissions to enable and process the integration of multiple third-party tools.

When you create an Integration Editor/Implementor, you define whether this user implements the API. This template includes API implementation using the four Organization-level **Service Accounts** permissions. If this is not required for your custom role, do not include the **Service Accounts** permissions.

## Group-level permissions

To create this role, enable the following permissions in the relevant categories:

### Snyk Apps management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Apps</td><td>true</td></tr><tr><td>Install Apps</td><td>true</td></tr><tr><td>Edit Apps</td><td>true</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Group management
* Organization management
* Snyk Essentials management
* Audit Log management
* IaC settings management
* Insights management
* Issue management
* Reports management
* Request access management
* Role management
* Security and licence policies
* Service account management
* Snyk Preview management
* SSO settings management
* Tags management
* User Management

## Organization-level permissions

To create this role, enable the following permissions in the relevant categories:

### Organization management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization</td><td>true</td></tr><tr><td>Edit Organization</td><td>false</td></tr><tr><td>Remove Organization</td><td>false</td></tr></tbody></table>

### Integration management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View integrations</td><td>true</td></tr><tr><td>Edit integrations</td><td>true</td></tr></tbody></table>

### Kubernetes integration management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>Publish Kubernetes resources</td><td>true</td></tr></tbody></table>

### Project management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Project</td><td>true</td></tr><tr><td>Add Project</td><td>true</td></tr><tr><td>Edit Project</td><td>true</td></tr><tr><td>Edit Project status</td><td>true</td></tr><tr><td>Test Project</td><td>true</td></tr><tr><td>Move Project</td><td>true</td></tr><tr><td>Remove Project</td><td>true</td></tr><tr><td>View Project history</td><td>false</td></tr><tr><td>Edit Project integrations</td><td>true</td></tr><tr><td>Edit Project attributes</td><td>false</td></tr><tr><td>View Jira issues</td><td>false</td></tr><tr><td>Create Jira issues</td><td>false</td></tr><tr><td>Edit Project Tags</td><td>false</td></tr></tbody></table>

### Project pull request management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>Create pull requests</td><td>true</td></tr><tr><td>Mark pull request checks as successful</td><td>true</td></tr></tbody></table>

### Reports management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Organization reports</td><td>true</td></tr></tbody></table>

### Service account management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View service accounts</td><td>true</td></tr><tr><td>Create service accounts</td><td>true</td></tr><tr><td>Edit service accounts</td><td>true</td></tr><tr><td>Remove service accounts</td><td>true</td></tr></tbody></table>

### Snyk Apps management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Apps</td><td>true</td></tr><tr><td>Install Apps</td><td>true</td></tr><tr><td>Create Apps</td><td>true</td></tr><tr><td>Edit Apps</td><td>true</td></tr><tr><td>Delete Apps</td><td>true</td></tr></tbody></table>

### Webhook management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>View Outbound Webhooks</td><td>true</td></tr><tr><td>Create Outbound Webhooks</td><td>true</td></tr><tr><td>Remove Outbound Webhooks</td><td>true</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Audit Log management
* Billing management
* Collection management
* Container Image management
* Entitlement management
* Integration management
* Package management
* Project Ignore management
* Snyk Cloud management
* Snyk Preview management
* User management
