# How to find information about API endpoints

This list includes the categories and names of REST GA and beta and V1 API endpoints, with the URL in the reference docs for each endpoint, and links to related information where available. REST is the default, and GA is the status unless beta is noted. V1 API is specified where applicable. This listing is a work in progress; additional information is being added continually.

## AccessRequests (beta)

### [Get access requests](https://apidocs.snyk.io/?beta=\&version=2024-05-23%7Ebeta#get-/self/access\_requests)

## Apps

More information: [Snyk Apps](../snyk-api-info/snyk-apps/)

### [Get a list of apps installed for a group](https://apidocs.snyk.io/?#get-/groups/-group\_id-/apps/installs)

### [Install a Snyk App to this group](https://apidocs.snyk.io/?#post-/groups/-group\_id-/apps/installs)

### [Revoke app authorization for a Snyk Group with install ID](https://apidocs.snyk.io/?#delete-/groups/-group\_id-/apps/installs/-install\_id-)

### [Manage client secret for non-interactive Snyk App installations](https://apidocs.snyk.io/?#post-/groups/-group\_id-/apps/installs/-install\_id-/secrets)

### DEPRECATED [Get a list of app bots authorized to an organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/app\_bots)

**Replaced by:** [Get a list of apps installed for an organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps/installs)

### DEPRECATED [Revoke app bot authorization](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/app\_bots/-bot\_id-)

**Replaced by:** [Revoke app authorization for a Snyk Group with install ID](https://apidocs.snyk.io/?#delete-/groups/-group\_id-/apps/installs/-install\_id-)

**See also:** [Revoke access for an app by install ID](https://apidocs.snyk.io/?#delete-/self/apps/installs/-install\_id-)

### DEPRECATED [Get a list of apps created by an organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps)

**Replaced by:** [Get a list of apps created by an organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps/creations)

More information: [Manage App details](snyk-apps/manage-app-details.md)

### DEPRECATED [Create a new app for an organization](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/apps)

**Replaced by:** [Create a new Snyk App for an organization](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/apps/creations)

**More information:** [Create a Snyk App using the Snyk API](snyk-apps/create-a-snyk-app-using-the-snyk-api.md)

### [Get a list of apps created by an organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps/creations)

Replaces: DEPRECATED Get a list of apps created by an organization

### [Create a new Snyk App for an organization](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/apps/creations)

**Replaces:** DEPRECATED Create a new app for an organization

### [Get a Snyk APP by its App ID](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps/creations/-app\_id-)

**Replaces:** DEPRECATED Get an app by client id

### [Delete an app by its App ID](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/apps/creations/-app\_id-)

**Replaces:** DEPRECATED Delete an app

**More information:** [Manage App details](snyk-apps/manage-app-details.md)

### [Update app creation attributes such as name, redirect URIs, and access token time to live using the App ID](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/apps/creations/-app\_id-)

**Replaces:** DEPRECATED Update App attributes that are name, redirect URIs, and access token time to live

**More information:** [Manage App details](snyk-apps/manage-app-details.md)

### [Manage client secret for the Snyk App](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/apps/creations/-app\_id-/secrets)

**More information:** [Manage App details](snyk-apps/manage-app-details.md)

### [Get a list of apps installed for an organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps/installs)

**Replaces:** DEPRECATED Get a list of app bots authorized to an organization

### [Install a Snyk App to this organization](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/apps/installs)

### [Revoke app authorization for a Snyk Organization with install ID](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/apps/installs/-install\_id-)

**See also:** Revoke app authorization for a Snyk Group with install ID

### [Manage client secret for non-interactive Snyk App installations](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/apps/installs/-install\_id-/secrets)

**Replaces:** DEPRECATED Manage client secrets for an app

### DEPRECATED [Get an app by client id](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps/-client\_id-)

**Replaced by:** [Get a Snyk App by its App ID](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps/creations/-app\_id-)

### DEPRECATED [Delete an app](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/apps/-client\_id-)

**Replaced by:** [Delete an app by its App ID](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/apps/creations/-app\_id-)

### DEPRECATED [Update app attributes that are name, redirect URIs, and access token time to live](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/apps/-client\_id-)

**Replaced by:** [Update app creation attributes such as name, redirect URIs, and access token time to live using the App ID](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/apps/creations/-app\_id-)

### DEPRECATED [Manage client secrets for an app](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/apps/-client\_id-/secrets)

**Replaced by:** [Revoke an app](https://apidocs.snyk.io/?#delete-/self/apps/-app\_id-)

### [Get a list of apps that can act on your behalf](https://apidocs.snyk.io/?version=2024-04-22#get-/self/apps)

### [Get a list of apps installed for a user](https://apidocs.snyk.io/?#get-/self/apps/installs)

### [Revoke access for an app by install ID](https://apidocs.snyk.io/?#delete-/self/apps/installs/-install\_id-)

**Replaces:** DEPRECATED Revoke app bot authorization

### [Revoke an app](https://apidocs.snyk.io/?#delete-/self/apps/-app\_id-)

### [Get a list of active OAuth sessions for the app](https://apidocs.snyk.io/?#get-/self/apps/-app\_id-/sessions)

### [Revoke an active user app session](https://apidocs.snyk.io/?#delete-/self/apps/-app\_id-/sessions/-session\_id-)

## Audit Logs

### [Search Group audit logs](https://apidocs.snyk.io/?#get-/groups/-group\_id-/audit\_logs/search)

**More information**

[Filter through your audit logs more efficiently with the new GA REST version of the audit logs API, and api.access is now opt-in](https://updates.snyk.io/filter-through-your-audit-logs-more-efficiently-with-the-new-ga-rest-version-of-the-audit-logs-api-and-api-access-is-now-opt-in-291850)

### [Search Organization audit logs](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/audit\_logs/search)

## Audit logs (v1)

### Group level audit logs

#### [Get group level audit logs](reference/audit-logs-v1.md)

[Migrated Get group level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/group-level-audit-logs/get-group-level-audit-logs)

### Organization level audit logs

#### [Get organization level audit logs](reference/audit-logs-v1.md)

[Migrated Get organization level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/organization-level-audit-logs/get-organization-level-audit-logs)

## Cloud (beta)

### [List Environments](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#get-/orgs/-org\_id-/cloud/environments)

### [Create New Environment](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#post-/orgs/-org\_id-/cloud/environments)

### [Delete Environment](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#delete-/orgs/-org\_id-/cloud/environments/-environment\_id-)

### [Update Environment](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#patch-/orgs/-org\_id-/cloud/environments/-environment\_id-)

### [Generate Cloud Provider Permissions](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#post-/orgs/-org\_id-/cloud/permissions)

### [List Resources](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#get-/orgs/-org\_id-/cloud/resources)

### [List Scans](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#get-/orgs/-org\_id-/cloud/scans)

### [Create Scan](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#post-/orgs/-org\_id-/cloud/scans)

### [Get scan](https://apidocs.snyk.io/?version=2024-05-23%7Ebeta#get-/orgs/-org\_id-/cloud/scans/-scan\_id-)

## Collection

### Get collections

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/collections](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/collections)

### Create a collection

[https://apidocs.snyk.io/?#post-/orgs/-org\_id-/collections](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/collections)

### Get a collection

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/collections/-collection\_id-](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/collections/-collection\_id-)

### Delete a collection

[https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/collections/-collection\_id-](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/collections/-collection\_id-)

### Edit a collection

[https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/collections/-collection\_id-](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/collections/-collection\_id-)

### Get projects from the specified collection

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/collections/-collection\_id-/relationships/projects](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/collections/-collection\_id-/relationships/projects)

### Add projects to a collection

[https://apidocs.snyk.io/?#post-/orgs/-org\_id-/collections/-collection\_id-/relationships/projects](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/collections/-collection\_id-/relationships/projects)

### Remove projects from a collection

[https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/collections/-collection\_id-/relationships/projects](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/collections/-collection\_id-/relationships/projects)

## ContainerImage

### List instances of container image

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/container\_images](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/container\_images)

### Get instance of container image

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/container\_images/-image\_id-](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/container\_images/-image\_id-)

### List instances of image target references for a container image

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/container\_images/-image\_id-/relationships/image\_target\_refs](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/container\_images/-image\_id-/relationships/image\_target\_refs)

## Custom Base Images

### Get a custom base image collection

[https://apidocs.snyk.io/?#get-/custom\_base\_images](https://apidocs.snyk.io/?#get-/custom\_base\_images)

### Create a Custom Base Image from an existing container project

[https://apidocs.snyk.io/?#post-/custom\_base\_images](https://apidocs.snyk.io/?#post-/custom\_base\_images)

### Get a custom base image

[https://apidocs.snyk.io/?#get-/custom\_base\_images/-custombaseimage\_id-](https://apidocs.snyk.io/?#get-/custom\_base\_images/-custombaseimage\_id-)

### Delete a custom base image

[https://apidocs.snyk.io/?#delete-/custom\_base\_images/-custombaseimage\_id-](https://apidocs.snyk.io/?#delete-/custom\_base\_images/-custombaseimage\_id-)

### Update a custom base image

[https://apidocs.snyk.io/?#patch-/custom\_base\_images/-custombaseimage\_id-](https://apidocs.snyk.io/?#patch-/custom\_base\_images/-custombaseimage\_id-)

## Dependencies (v1)

### Dependencies by organization

#### List all dependencies

[Migrated List all dependencies](how-to-find-information-about-api-endpoints.md#list-all-dependencies)

## Entitlements (v1)

### Entitlements by organization

#### List all entitlements

[Migrated list all entitlements](https://snyk.docs.apiary.io/#reference/entitlements/entitlements-by-organization/list-all-entitlements)

### A specific entitlement by organization

#### Get an organization's entitlement value

[Migrated Get an organization's entitlement value](https://snyk.docs.apiary.io/#reference/entitlements/a-specific-entitlement-by-organization/get-an-organization's-entitlement-value)

## Groups (beta)

### Get all groups

https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#get-/groups

### Get a group

https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#get-/groups/-group\_id-/sso\_connections

### Get all SSO connections for a group

https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#get-/groups/-group\_id-/sso\_connections/-sso\_id-/users

### Get all users using a given SSO connection

[https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#get-/groups/-group\_id-/sso\_connections/-sso\_id-/users](https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#get-/groups/-group\_id-/sso\_connections/-sso\_id-/users)

### Delete a user from a Group SSO connection

https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#delete-/groups/-group\_id-/sso\_connections/-sso\_id-/users/-user\_id-

## Groups (v1)

### Group settings

#### View group settings

[Migrated View group settings](https://snyk.docs.apiary.io/#reference/groups/group-settings/view-group-settings)

#### Update group settings

[Migrated Update group settings](https://snyk.docs.apiary.io/#reference/groups/group-settings/update-group-settings)

### List members in a group

#### List all members in a group

[https://snyk.docs.apiary.io/#reference/groups/list-members-in-a-group/list-all-members-in-a-group](https://snyk.docs.apiary.io/#reference/groups/list-members-in-a-group/list-all-members-in-a-group)

### Members in an organization of a group

#### Add a member to an organization within a group

[https://snyk.docs.apiary.io/#reference/groups/members-in-an-organization-of-a-group/add-a-member-to-an-organization-within-a-group](https://snyk.docs.apiary.io/#reference/groups/members-in-an-organization-of-a-group/add-a-member-to-an-organization-within-a-group)

### List all tags in a group

#### List all tags in a group

[https://snyk.docs.apiary.io/#reference/groups/list-all-tags-in-a-group/list-all-tags-in-a-group](https://snyk.docs.apiary.io/#reference/groups/list-all-tags-in-a-group/list-all-tags-in-a-group)

### Delete Tag From Group

#### Delete tag from group

[https://snyk.docs.apiary.io/#reference/groups/delete-tag-from-group/delete-tag-from-group](https://snyk.docs.apiary.io/#reference/groups/delete-tag-from-group/delete-tag-from-group)

### List all organizations in a group

#### List all organizations in a group

[https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)

### List all roles in a group

#### List all roles in a group

[https://snyk.docs.apiary.io/#reference/groups/list-all-roles-in-a-group/list-all-roles-in-a-group](https://snyk.docs.apiary.io/#reference/groups/list-all-roles-in-a-group/list-all-roles-in-a-group)

## IacSettings

### Get the Infrastructure as Code Settings for a group

[https://apidocs.snyk.io/?#get-/groups/-group\_id-/settings/iac](https://apidocs.snyk.io/?#get-/groups/-group\_id-/settings/iac)

### Update the Infrastructure as Code Settings for a group

[https://apidocs.snyk.io/?#patch-/groups/-group\_id-/settings/iac](https://apidocs.snyk.io/?#patch-/groups/-group\_id-/settings/iac)

[https://apidocs.snyk.io/?#get-/groups/-group\_id-/settings/iac](https://apidocs.snyk.io/?#get-/groups/-group\_id-/settings/iac)

### Get the Infrastructure as Code Settings for an org

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/settings/iac](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/settings/iac)

### Update the Infrastructure as Code Settings for an org

[https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/settings/iac](https://apidocs.snyk.io/?version=2024-04-22#patch-/orgs/-org\_id-/settings/iac)

## Import Projects (v1)

### Import

#### Import targets

[https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets)

### Import job

#### Get import job details

[https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details)

## Integrations (v1)

### Integrations

#### List

[https://snyk.docs.apiary.io/#reference/integrations/integrations/list](https://snyk.docs.apiary.io/#reference/integrations/integrations/list)

#### Add new integration

[https://snyk.docs.apiary.io/#reference/integrations/integrations/add-new-integration](https://snyk.docs.apiary.io/#reference/integrations/integrations/add-new-integration)

### Integration

#### Update existing integration

[https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration)

### Integration authentication

#### Delete credentials

[https://snyk.docs.apiary.io/#reference/integrations/integration-authentication/delete-credentials](https://snyk.docs.apiary.io/#reference/integrations/integration-authentication/delete-credentials)

### Integration broker token provisioning

#### Provision new broker token

[https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-provisioning/provision-new-broker-token](https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-provisioning/provision-new-broker-token)

### Integration broker token switching

#### Switch between broker tokens

[https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-switching/switch-between-broker-tokens](https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-switching/switch-between-broker-tokens)

### Integration cloning

#### Clone  an integration (with settings and credentials)

[https://snyk.docs.apiary.io/#reference/integrations/integration-cloning/clone-an-integration-(with-settings-and-credentials)](https://snyk.docs.apiary.io/#reference/integrations/integration-cloning/clone-an-integration-\(with-settings-and-credentials\))

### Integration by type

#### Get existing integration by type

[https://snyk.docs.apiary.io/#reference/integrations/integration-by-type/get-existing-integration-by-type](https://snyk.docs.apiary.io/#reference/integrations/integration-by-type/get-existing-integration-by-type)

### Integration settings

#### Retrieve

[https://snyk.docs.apiary.io/#reference/integrations/integration-settings/retrieve](https://snyk.docs.apiary.io/#reference/integrations/integration-settings/retrieve)

#### Update

[https://snyk.docs.apiary.io/#reference/integrations/integration-settings/update](https://snyk.docs.apiary.io/#reference/integrations/integration-settings/update)

## Invites

#### See also Invite users

[https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users](https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users)

### List pending user invitations to an organization

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/invites](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/invites)

### Invite a user to an organization

[https://apidocs.snyk.io/?#post-/orgs/-org\_id-/invites](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/invites)

### Cancel a pending user invitations to an organization

[https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/invites/-invite\_id-](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/invites/-invite\_id-)

## Issues

### Get issues by group ID

[https://apidocs.snyk.io/?#get-/groups/-group\_id-/issues](https://apidocs.snyk.io/?#get-/groups/-group\_id-/issues)

### Get an issue

[https://apidocs.snyk.io/?#get-/groups/-group\_id-/issues/-issue\_id-](https://apidocs.snyk.io/?#get-/groups/-group\_id-/issues/-issue\_id-)

### Get issues by org ID

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/issues](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/issues)

### Get an issue

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/issues/-issue\_id-](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/issues/-issue\_id-)

### List issues for a given set of packages (Currently not available to all customers)

[https://apidocs.snyk.io/?#post-/orgs/-org\_id-/packages/issues](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/packages/issues)

### List issues for a package

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/packages/-purl-/issues](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/packages/-purl-/issues)

## Licenses (v1)

### Licenses by organization

#### List all licenses

[https://snyk.docs.apiary.io/#reference/licenses/licenses-by-organization/list-all-licenses](https://snyk.docs.apiary.io/#reference/licenses/licenses-by-organization/list-all-licenses)

## Monitor (v1)

### Dep Graph

#### Monitor Dep Graph

[https://snyk.docs.apiary.io/#reference/monitor/depgraph/monitor-dep-graph](https://snyk.docs.apiary.io/#reference/monitor/depgraph/monitor-dep-graph)

## OpenAPI

### List available versions of OpenAPI specification

[https://apidocs.snyk.io/?#get-/openapi](https://apidocs.snyk.io/?#get-/openapi)

### Get OpenAPI specification effective at version

[https://apidocs.snyk.io/?#get-/openapi/-version-](https://apidocs.snyk.io/?#get-/openapi/-version-)

## Organizations (v1)

### The Snyk organization for a request

#### List all the organizations a user belongs to

[https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to)

### Create organization

#### Create a new organization

[https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization](https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization)

### Notification settings

#### Get organization notification settings

[https://snyk.docs.apiary.io/#reference/organizations/notification-settings/get-organization-notification-settings](https://snyk.docs.apiary.io/#reference/organizations/notification-settings/get-organization-notification-settings)

#### Set notification settings

[https://snyk.docs.apiary.io/#reference/organizations/notification-settings/set-notification-settings](https://snyk.docs.apiary.io/#reference/organizations/notification-settings/set-notification-settings)

### User invitation to organization

#### Invite users

[https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users](https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users)

### Members in organization

#### List members

[https://snyk.docs.apiary.io/#reference/organizations/members-in-organization/list-members](https://snyk.docs.apiary.io/#reference/organizations/members-in-organization/list-members)

### Organization settings

#### View organization settings

[https://snyk.docs.apiary.io/#reference/organizations/organization-settings/view-organization-settings](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/view-organization-settings)

#### Update organization settings

[https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-organization-settings](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-organization-settings)

### Manage roles in organization

#### Update a member in the organization

[https://snyk.docs.apiary.io/#reference/organizations/manage-roles-in-organization/update-a-member-in-the-organization](https://snyk.docs.apiary.io/#reference/organizations/manage-roles-in-organization/update-a-member-in-the-organization)

#### Remove a member from the organization

[https://snyk.docs.apiary.io/#reference/organizations/manage-roles-in-organization/remove-a-member-from-the-organization](https://snyk.docs.apiary.io/#reference/organizations/manage-roles-in-organization/remove-a-member-from-the-organization)

### Update member roles in your organization

#### Update a member's role in the organization

[https://snyk.docs.apiary.io/#reference/organizations/update-member-roles-in-your-organization/update-a-member's-role-in-the-organization](https://snyk.docs.apiary.io/#reference/organizations/update-member-roles-in-your-organization/update-a-member's-role-in-the-organization)

### Manage organization

#### Remove organization

[https://snyk.docs.apiary.io/#reference/organizations/manage-organization/remove-organization](https://snyk.docs.apiary.io/#reference/organizations/manage-organization/remove-organization)

### Provision user

#### Provision a user to the organization

[https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization](https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization)

#### List pending user provisions

[https://snyk.docs.apiary.io/#reference/organizations/provision-user/list-pending-user-provisions](https://snyk.docs.apiary.io/#reference/organizations/provision-user/list-pending-user-provisions)

#### Delete pending user provision

[https://snyk.docs.apiary.io/#reference/organizations/provision-user/delete-pending-user-provision](https://snyk.docs.apiary.io/#reference/organizations/provision-user/delete-pending-user-provision)

## Orgs (GA and beta)

### List all organizations in a group

[https://apidocs.snyk.io/?#get-/groups/-group\_id-/orgs](https://apidocs.snyk.io/?#get-/groups/-group\_id-/orgs)

### List accessible organizations

[https://apidocs.snyk.io/?#get-/orgs](https://apidocs.snyk.io/?#get-/orgs)

### Get an ORG (beta)

[https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#get-/orgs/-org\_id-](https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#get-/orgs/-org\_id-)

### Get organization

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-](https://apidocs.snyk.io/?#get-/orgs/-org\_id-)

### Update organization

[https://apidocs.snyk.io/?#patch-/orgs/-org\_id-](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-)

## Projects

### List all Projects for an Org with the given Org ID

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects)

### Get project by project ID

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects/-project\_id-](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects/-project\_id-)

### Delete project by project ID

[https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/projects/-project\_id-](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/projects/-project\_id-)

### Updates project by project ID

[https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/projects/-project\_id-](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/projects/-project\_id-)

## Projects (v1)

### Individual project

#### Retrieve a single project

[https://snyk.docs.apiary.io/#reference/projects/individual-project/retrieve-a-single-project](https://snyk.docs.apiary.io/#reference/projects/individual-project/retrieve-a-single-project)

#### Update a project

[https://snyk.docs.apiary.io/#reference/projects/individual-project/update-a-project](https://snyk.docs.apiary.io/#reference/projects/individual-project/update-a-project)

#### Delete a project

[https://snyk.docs.apiary.io/#reference/projects/individual-project/delete-a-project](https://snyk.docs.apiary.io/#reference/projects/individual-project/delete-a-project)

### Deactivate an individual project

#### Deactivate

[https://snyk.docs.apiary.io/#reference/projects/deactivate-an-individual-project/deactivate](https://snyk.docs.apiary.io/#reference/projects/deactivate-an-individual-project/deactivate)

### Activate an individual project

#### Activate

[https://snyk.docs.apiary.io/#reference/projects/activate-an-individual-project/activate](https://snyk.docs.apiary.io/#reference/projects/activate-an-individual-project/activate)

### Aggregated Project issues

#### List all Aggregated issues

[https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues)

### Project Issue Paths

#### List all project issue paths

[https://snyk.docs.apiary.io/#reference/projects/project-issue-paths/list-all-project-issue-paths](https://snyk.docs.apiary.io/#reference/projects/project-issue-paths/list-all-project-issue-paths)

### Project History

#### List all project snapshots

[https://snyk.docs.apiary.io/#reference/projects/project-history/list-all-project-snapshots](https://snyk.docs.apiary.io/#reference/projects/project-history/list-all-project-snapshots)

### Aggregated Project Snapshot Issues

#### List all project snapshot aggregated issues

[https://snyk.docs.apiary.io/#reference/projects/aggregated-project-snapshot-issues/list-all-project-snapshot-aggregated-issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-snapshot-issues/list-all-project-snapshot-aggregated-issues)

### Project Snapshot Issue Paths

#### List all project snapshot issue paths

[https://snyk.docs.apiary.io/#reference/projects/project-snapshot-issue-paths/list-all-project-snapshot-issue-paths](https://snyk.docs.apiary.io/#reference/projects/project-snapshot-issue-paths/list-all-project-snapshot-issue-paths)

### Project dependency graph

#### Get Project dependency graph

[https://snyk.docs.apiary.io/#reference/projects/project-dependency-graph/get-project-dependency-graph](https://snyk.docs.apiary.io/#reference/projects/project-dependency-graph/get-project-dependency-graph)

### Project ignores

#### List all ignores

[https://snyk.docs.apiary.io/#reference/projects/project-ignores/list-all-ignores](https://snyk.docs.apiary.io/#reference/projects/project-ignores/list-all-ignores)

### Ignored issues

#### Retrieve ignore

[https://snyk.docs.apiary.io/#reference/projects/ignored-issues/retrieve-ignore](https://snyk.docs.apiary.io/#reference/projects/ignored-issues/retrieve-ignore)

#### Add ignore

[https://snyk.docs.apiary.io/#reference/projects/ignored-issues/add-ignore](https://snyk.docs.apiary.io/#reference/projects/ignored-issues/add-ignore)

#### Replace ignores

[https://snyk.docs.apiary.io/#reference/projects/ignored-issues/replace-ignores](https://snyk.docs.apiary.io/#reference/projects/ignored-issues/replace-ignores)

#### Delete ignores

[https://snyk.docs.apiary.io/#reference/projects/ignored-issues/delete-ignores](https://snyk.docs.apiary.io/#reference/projects/ignored-issues/delete-ignores)&#x20;

### Project jira issues

#### List all jira issues

[https://snyk.docs.apiary.io/#reference/projects/project-jira-issues/list-all-jira-issues](https://snyk.docs.apiary.io/#reference/projects/project-jira-issues/list-all-jira-issues)&#x20;

#### Create jira issue

[https://snyk.docs.apiary.io/#reference/projects/project-jira-issues/create-jira-issue](https://snyk.docs.apiary.io/#reference/projects/project-jira-issues/create-jira-issue)

### Project settings

#### List project settings

[https://snyk.docs.apiary.io/#reference/projects/project-settings/list-project-settings](https://snyk.docs.apiary.io/#reference/projects/project-settings/list-project-settings)

#### Update project settings

[https://snyk.docs.apiary.io/#reference/projects/project-settings/update-project-settings](https://snyk.docs.apiary.io/#reference/projects/project-settings/update-project-settings)

#### Delete project settings

[https://snyk.docs.apiary.io/#reference/projects/project-settings/delete-project-settings](https://snyk.docs.apiary.io/#reference/projects/project-settings/delete-project-settings)

### Move project

#### Move project to a different organization

[https://snyk.docs.apiary.io/#reference/projects/move-project/move-project-to-a-different-organization](https://snyk.docs.apiary.io/#reference/projects/move-project/move-project-to-a-different-organization)

### Project tags

#### Add a tag to a project

[https://snyk.docs.apiary.io/#reference/projects/project-tags/add-a-tag-to-a-project](https://snyk.docs.apiary.io/#reference/projects/project-tags/add-a-tag-to-a-project)

### Remove project tag

#### Remove a tag from a project

[https://snyk.docs.apiary.io/#reference/projects/remove-project-tag/remove-a-tag-from-a-project](https://snyk.docs.apiary.io/#reference/projects/remove-project-tag/remove-a-tag-from-a-project)

### Project Attributes

#### Applying attributes

[https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes)

## Pull request templates

For information on how to use these endpoints, see [Create and manage a custom PR template using the API](../scan-with-snyk/pull-requests/snyk-fix-pull-or-merge-requests/customize-pr-templates/apply-a-custom-pr-template.md#create-and-manage-a-custom-pr-template-using-the-api).

### Get pull request template for group

[https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#get-/groups/-group\_id-/settings/pull\_request\_template](https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#get-/groups/-group\_id-/settings/pull\_request\_template)

### Create or update pull request template for group

[https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#post-/groups/-group\_id-/settings/pull\_request\_template](https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#post-/groups/-group\_id-/settings/pull\_request\_template)

### Delete pull request template for group

[https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#delete-/groups/-group\_id-/settings/pull\_request\_template](https://apidocs.snyk.io/?version=2024-04-29%7Ebeta#delete-/groups/-group\_id-/settings/pull\_request\_template)

## Reporting API (v1)

### Latest issues

#### Get list of latest issues

[https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues)

### Issues

#### Get list of latest issues

[https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues](https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues)

### Latest issue counts

#### Get latest issue counts

[https://snyk.docs.apiary.io/#reference/reporting-api/latest-issue-counts/get-latest-issue-counts](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issue-counts/get-latest-issue-counts)

### Issue counts over time

#### Get issue counts

[https://snyk.docs.apiary.io/#reference/reporting-api/issue-counts-over-time/get-issue-counts](https://snyk.docs.apiary.io/#reference/reporting-api/issue-counts-over-time/get-issue-counts)

### Latest project counts

#### Get latest project counts

[https://snyk.docs.apiary.io/#reference/reporting-api/latest-project-counts/get-latest-project-counts](https://snyk.docs.apiary.io/#reference/reporting-api/latest-project-counts/get-latest-project-counts)

### Project counts over time

#### Get project counts

[https://snyk.docs.apiary.io/#reference/reporting-api/project-counts-over-time/get-project-counts](https://snyk.docs.apiary.io/#reference/reporting-api/project-counts-over-time/get-project-counts)

### Test counts

#### Get test counts

[https://snyk.docs.apiary.io/#reference/reporting-api/test-counts/get-test-counts](https://snyk.docs.apiary.io/#reference/reporting-api/test-counts/get-test-counts)

## SBOM (GA and beta)

### Get a project’s SBOM document

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects/-project\_id-/sbom](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects/-project\_id-/sbom)

### Create an SBOM test run (beta)

[https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#post-/orgs/-org\_id-/sbom\_tests](https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#post-/orgs/-org\_id-/sbom\_tests)

### Gets an SBOM test run status (beta)

[https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-](https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-)

### Gets an SBOM test run result (beta)

[https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-/results](https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-/results)

## SastSettings

### Retrieves the SAST settings for an org

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/settings/sast](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/settings/sast)

### Enable/Disable the Snyk Code settings for an org

[https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/settings/sast](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/settings/sast)

## ServiceAccounts

### Get a list of group service accounts

[https://apidocs.snyk.io/?#get-/groups/-group\_id-/service\_accounts](https://apidocs.snyk.io/?#get-/groups/-group\_id-/service\_accounts)

### Create a service account for a group

[https://apidocs.snyk.io/?#post-/groups/-group\_id-/service\_accounts](https://apidocs.snyk.io/?#post-/groups/-group\_id-/service\_accounts)

### Get a group service account

[https://apidocs.snyk.io/?#get-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?#get-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

### Delete a group service account

[https://apidocs.snyk.io/?#delete-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?version=2024-04-22#delete-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

### Update a group service account

[https://apidocs.snyk.io/?#patch-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?#patch-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

### Manage a group service account’s client secret

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/service\_accounts](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/service\_accounts)

### Get a list of organization service accounts

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/service\_accounts](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/service\_accounts)

### Create a service account for an organization

[https://apidocs.snyk.io/?#get-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

### Get an organization service account

[https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

### Delete a service account in an organization

[https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

### Update an organization service account

[https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

### Manage an organization service account’s client secret

[https://apidocs.snyk.io/?#post-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-/secrets](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-/secrets)

## Slack

### Get a list of Slack channels

[https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/slack\_app/-tenant\_id-/channels](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/slack\_app/-tenant\_id-/channels)

### Get Slack Channel name by Slack Channel ID

[https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/slack\_app/-tenant\_id-/channels/-channel\_id-](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/slack\_app/-tenant\_id-/channels/-channel\_id-)

## SlackSettings

### Get Slack integration default notification settings

[https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/slack\_app/-bot\_id-](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/slack\_app/-bot\_id-)

### Create new Slack notification default settings

[https://apidocs.snyk.io/?version=2024-04-22#post-/orgs/-org\_id-/slack\_app/-bot\_id-](https://apidocs.snyk.io/?version=2024-04-22#post-/orgs/-org\_id-/slack\_app/-bot\_id-)

### Remove the given Slack App integration

[https://apidocs.snyk.io/?version=2024-04-22#delete-/orgs/-org\_id-/slack\_app/-bot\_id-](https://apidocs.snyk.io/?version=2024-04-22#delete-/orgs/-org\_id-/slack\_app/-bot\_id-)

### Slack notification settings overrides for projects

[https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/slack\_app/-bot\_id-/projects](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/slack\_app/-bot\_id-/projects)

### Create a mew Slack settings override for a given project

[https://apidocs.snyk.io/?version=2024-04-22#post-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-](https://apidocs.snyk.io/?version=2024-04-22#post-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-)

### Remove Slack settings override for a project

https://apidocs.snyk.io/?version=2024-04-22#delete-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-

### Update Slack notification settings for a project

https://apidocs.snyk.io/?version=2024-04-22#delete-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-

## Targets

### Get targets by org ID

[https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/targets](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/targets)

### Get target by target ID

[https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/targets/-target\_id-](https://apidocs.snyk.io/?version=2024-04-22#get-/orgs/-org\_id-/targets/-target\_id-)

### Delete target by target ID

[https://apidocs.snyk.io/?version=2024-04-22#delete-/orgs/-org\_id-/targets/-target\_id-](https://apidocs.snyk.io/?version=2024-04-22#delete-/orgs/-org\_id-/targets/-target\_id-)

## Test (v1)

### Maven

#### Test for issues in a public package by group is, artifact id and version

[https://snyk.docs.apiary.io/#reference/test/maven/test-for-issues-in-a-public-package-by-group-id,-artifact-id-and-version](https://snyk.docs.apiary.io/#reference/test/maven/test-for-issues-in-a-public-package-by-group-id,-artifact-id-and-version)

#### Test maven file

[https://snyk.docs.apiary.io/#reference/test/maven/test-maven-file](https://snyk.docs.apiary.io/#reference/test/maven/test-maven-file)

### npm

#### Test for issues in a public package by name and version

[https://snyk.docs.apiary.io/#reference/test/npm/test-for-issues-in-a-public-package-by-name-and-version](https://snyk.docs.apiary.io/#reference/test/npm/test-for-issues-in-a-public-package-by-name-and-version)

#### Test package.json & package-lock.json File

[https://snyk.docs.apiary.io/#reference/test/npm/test-package.json-&-package-lock.json-file](https://snyk.docs.apiary.io/#reference/test/npm/test-package.json-&-package-lock.json-file)

### dep

#### Test Gopkg.toml & Gopkg.lock File

[https://snyk.docs.apiary.io/#reference/test/dep/test-gopkg.toml-&-gopkg.lock-file](https://snyk.docs.apiary.io/#reference/test/dep/test-gopkg.toml-&-gopkg.lock-file)

### vendor

#### Test vendor.json file

[https://snyk.docs.apiary.io/#reference/test/vendor/test-vendor.json-file](https://snyk.docs.apiary.io/#reference/test/vendor/test-vendor.json-file)

### yarn

#### Test package.json & yarn.lock file

[https://snyk.docs.apiary.io/#reference/test/yarn/test-package.json-&-yarn.lock-file](https://snyk.docs.apiary.io/#reference/test/yarn/test-package.json-&-yarn.lock-file)

### rubygems

#### Test for issues in a public gem by name and version

[https://snyk.docs.apiary.io/#reference/test/rubygems/test-for-issues-in-a-public-gem-by-name-and-version](https://snyk.docs.apiary.io/#reference/test/rubygems/test-for-issues-in-a-public-gem-by-name-and-version)

#### Test gemfile.lock file

[https://snyk.docs.apiary.io/#reference/test/rubygems/test-gemfile.lock-file](https://snyk.docs.apiary.io/#reference/test/rubygems/test-gemfile.lock-file)

### Gradle

#### Test for issues in a public package by group, name and version

[https://snyk.docs.apiary.io/#reference/test/gradle/test-for-issues-in-a-public-package-by-group,-name-and-version](https://snyk.docs.apiary.io/#reference/test/gradle/test-for-issues-in-a-public-package-by-group,-name-and-version)

#### Test gradle file

[https://snyk.docs.apiary.io/#reference/test/gradle/test-gradle-file](https://snyk.docs.apiary.io/#reference/test/gradle/test-gradle-file)

### sbt

#### Test for issues in a public package by group id, artifact id and version

[https://snyk.docs.apiary.io/#reference/test/sbt/test-for-issues-in-a-public-package-by-group-id,-artifact-id-and-version](https://snyk.docs.apiary.io/#reference/test/sbt/test-for-issues-in-a-public-package-by-group-id,-artifact-id-and-version)

#### Test sbt file

[https://snyk.docs.apiary.io/#reference/test/sbt/test-sbt-file](https://snyk.docs.apiary.io/#reference/test/sbt/test-sbt-file)

### pip

#### Test for issues in a public package by name and version

[https://snyk.docs.apiary.io/#reference/test/pip/test-for-issues-in-a-public-package-by-name-and-version](https://snyk.docs.apiary.io/#reference/test/pip/test-for-issues-in-a-public-package-by-name-and-version)

#### Test requirements.txt file

[https://snyk.docs.apiary.io/#reference/test/pip/test-requirements.txt-file](https://snyk.docs.apiary.io/#reference/test/pip/test-requirements.txt-file)

### composer

#### Test composer.json & composer.lock file

[https://snyk.docs.apiary.io/#reference/test/composer/test-composer.json-&-composer.lock-file](https://snyk.docs.apiary.io/#reference/test/composer/test-composer.json-&-composer.lock-file)

### Dep Graph

#### Test Dep Graph

[https://snyk.docs.apiary.io/#reference/test/dep-graph/test-dep-graph](https://snyk.docs.apiary.io/#reference/test/dep-graph/test-dep-graph)

## Users

### Update a user’s role in a group (beta)

[https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#patch-/groups/-group\_id-/users/-id-](https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#patch-/groups/-group\_id-/users/-id-)

### Get user by ID (beta)

[https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#get-/orgs/-org\_id-/users/-id-](https://apidocs.snyk.io/?version=2024-04-22%7Ebeta#get-/orgs/-org\_id-/users/-id-)

### My User Details

[https://apidocs.snyk.io/?version=2024-04-22#get-/self](https://apidocs.snyk.io/?version=2024-04-22#get-/self)

## Users(v1)

### User Details

#### Get user details

[https://snyk.docs.apiary.io/#reference/users/user-details/get-user-details](https://snyk.docs.apiary.io/#reference/users/user-details/get-user-details)

### My User Details

#### Get My Details

[https://snyk.docs.apiary.io/#reference/users/my-user-details/get-organization-notification-settings](https://snyk.docs.apiary.io/#reference/users/my-user-details/get-organization-notification-settings)

### User organization notification settings

#### Get organization notification settings

[https://snyk.docs.apiary.io/#reference/users/user-organization-notification-settings/get-organization-notification-settings](https://snyk.docs.apiary.io/#reference/users/user-organization-notification-settings/get-organization-notification-settings)

#### Modify organization notification settings

[https://snyk.docs.apiary.io/#reference/users/user-organization-notification-settings/modify-organization-notification-settings](https://snyk.docs.apiary.io/#reference/users/user-organization-notification-settings/modify-organization-notification-settings)

### User project notification settings

#### Get project notification settings

[https://snyk.docs.apiary.io/#reference/users/user-project-notification-settings/get-project-notification-settings](https://snyk.docs.apiary.io/#reference/users/user-project-notification-settings/get-project-notification-settings)

#### Modify project notification settings

[https://snyk.docs.apiary.io/#reference/users/user-project-notification-settings/modify-project-notification-settings](https://snyk.docs.apiary.io/#reference/users/user-project-notification-settings/modify-project-notification-settings)

## Webhooks (v1)

### Webhook Collection

#### Create a webhook

#### List webhooks

[https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/create-a-webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/create-a-webhook)

### Webhook

#### Retrieve a webhook

[https://snyk.docs.apiary.io/#reference/webhooks/webhook/retrieve-a-webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook/retrieve-a-webhook)

#### Delete a webhook

[https://snyk.docs.apiary.io/#reference/webhooks/webhook/delete-a-webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook/delete-a-webhook)

### Ping

#### Ping a webhook

[https://snyk.docs.apiary.io/#reference/webhooks/ping/ping-a-webhook](https://snyk.docs.apiary.io/#reference/webhooks/ping/ping-a-webhook)

\
