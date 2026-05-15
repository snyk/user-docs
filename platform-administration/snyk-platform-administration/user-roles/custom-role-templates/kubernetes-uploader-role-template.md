# Kubernetes Uploader role template

This Organization-level role can publish Kubernetes Monitor and Insights Collector data to Snyk and is often tied to a [**Service Account**](../../../implementation-and-setup/enterprise-setup/service-accounts/).

## Group-level permissions

This template is for an Organization-level role and has no Group-level permissions.

## Organization-level permissions

To create this role, enable the following permissions in the relevant categories:

### Kubernetes integration management

<table><thead><tr><th>Permission</th><th data-type="checkbox">Enabled?</th></tr></thead><tbody><tr><td>Publish Kubernetes resources</td><td>true</td></tr></tbody></table>

The remaining categories of permissions listed below should have all permissions within them set to disabled:

* Organization management
* Audit Log management
* Billing management
* Collection management
* Container Image management
* Entitlement management
* Integration management
* Package management
* Project management
* Project Ignore management
* Project pull request management
* Reports management
* Service account management
* Snyk Apps management
* Snyk Cloud management
* Snyk Preview management
* User management
* Webhook management
