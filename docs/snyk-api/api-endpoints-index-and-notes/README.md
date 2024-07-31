# API endpoints index and notes

This list includes the categories and names of REST GA and beta and V1 API endpoints, with the URL in the reference docs for each endpoint, and links to related information where available. REST is the default, and GA is the status unless beta is noted. V1 API is specified where applicable. This listing is a work in progress; additional information is being added continually. For additional information, see [Solutions for specific use cases](solutions-for-specific-use-cases.md) and [Scenarios for using Snyk API](scenarios-for-using-snyk-api.md).

## AccessRequests (beta)

### [Get access requests](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#get-/self/access\_requests)

## Apps

**More information:** [Snyk Apps](../../snyk-api-info/snyk-apps/)

### [Get a list of apps that act on your behalf](../reference/apps.md#self-apps)

### [Revoke an app](../reference/apps.md#self-apps-app\_id)

### [Get a list of active OAuth sessions for the app](../reference/apps.md#self-apps-app\_id-sessions)

### [Revoke an active user app session](../reference/apps.md#self-apps-app\_id-sessions-session\_id)

### [Get a list of apps installed for a user](../reference/apps.md#self-apps-installs)

### [Revoke access for an app by install ID](../reference/apps.md#self-apps-installs-install\_id)

**Replaces:** DEPRECATED Revoke app bot authorization

### DEPRECATED [Create a new app for an organization](../reference/apps.md#orgs-org\_id-apps)

**Replaced by:** Create a new Snyk App for an organization

**More information:** [Create a Snyk App using the Snyk API](../snyk-apps/create-a-snyk-app-using-the-snyk-api.md)

### [Get a list of apps created by an organization](../reference/apps.md#orgs-org\_id-apps-1)

**Replaces:** DEPRECATED Get a list of apps created by an organization

### DEPRECATED [Update app attributes that are name, redirect URIs, and access token time to live](../reference/apps.md#orgs-org\_id-apps-client\_id)

**Replaced by:** Update app creation attributes such as name, redirect URIs, and access token time to live using the App ID

### DEPRECATED [Get an app by client id](../reference/apps.md#orgs-org\_id-apps-client\_id-1)

**Replaced by:** Get a Snyk App by its App ID

### DEPRECATED [Delete an app](../reference/apps.md#orgs-org\_id-apps-client\_id-2)

**Replaced by:** Delete an app by its App ID

### DEPRECATED [Manage client secrets for an app](../reference/apps.md#orgs-org\_id-apps-client\_id-secrets)

**Replaced by:** Manage client secret for non-interactive Snyk App installations

### [Install a Snyk App to this organization](../reference/apps.md#orgs-org\_id-apps-installs)

### [Get a list of apps installed for an organization](../reference/apps.md#orgs-org\_id-apps-installs-1)

**Replaces:** DEPRECATED Get a list of app bots authorized to an organization

### [Revoke app authorization for a Snyk organization](../reference/apps.md#orgs-org\_id-apps-installs-install\_id)

**See also:** Revoke app authorization for a Snyk Group with install ID

### [Manage client secret for non-interactive Snyk App installations](../reference/apps.md#orgs-org\_id-apps-installs-install\_id-secrets)

### [Create a new Snyk App for an organization](../reference/apps.md#orgs-org\_id-apps-creations)

**Replaces:** DEPRECATED Create a new app for an organization

### DEPRECATED [Get a list of apps created by an organization](../reference/apps.md#orgs-org\_id-apps-creations-1)

**Replaced by:** Get a list of apps created by an organization

**More information:** [Manage App details](../snyk-apps/manage-app-details.md)

### [Update app creation attributes such as name, redirect URIs, and access token time to live using the App ID](../reference/apps.md#orgs-org\_id-apps-creations-app\_id)

**Replaces:** DEPRECATED Update App attributes that are name, redirect URIs, and access token time to live

**More information:** [Manage App details](../snyk-apps/manage-app-details.md)

### [Get a Snyk APP by its App ID](../reference/apps.md#orgs-org\_id-apps-creations-app\_id)

**Replaces:** DEPRECATED Get an app by client id

### [Delete a Snyk App by its App ID](../reference/apps.md#orgs-org\_id-apps-creations-app\_id-2)

**Replaces:** DEPRECATED Delete an app

**More information:** [Manage App details](../snyk-apps/manage-app-details.md)

### [Manage client secret for the Snyk App](../reference/apps.md#orgs-org\_id-apps-creations-app\_id-secrets)

**More information:** [Manage App details](../snyk-apps/manage-app-details.md)

### DEPRECATED [Get a list of app bots authorized to an organization](../reference/apps.md#orgs-org\_id-app\_bots)

**Replaced by:** [Get a list of apps installed for an organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/apps/installs)

### DEPRECATED [Revoke app bot authorization](../reference/apps.md#orgs-org\_id-app\_bots-bot\_id)

**Replaced by:** Revoke app authorization for a Snyk Group with install ID

**See also:** [Revoke access for an app by install](https://apidocs.snyk.io/?#delete-/self/apps/installs/-install\_id-)

### [Install a Snyk App to this group](../reference/apps.md#groups-group\_id-apps-installs)

### [Get a list of apps installed for a group](../reference/apps.md#groups-group\_id-apps-installs-1)

### [Revoke app authorization for a Snyk Group with install ID](../reference/apps.md#groups-group\_id-apps-installs-install\_id)

### [Manage client secret for non-interactive Snyk App installations](../reference/apps.md#groups-group\_id-apps-installs-install\_id-secrets)

**Replaces:** DEPRECATED Manage client secrets for an app

## Audit Logs

### [Search Organization audit log](../reference/audit-logs.md#orgs-org\_id-audit\_logs-search)s

**More information:** [Retrieve audit logs of user-initiated activity by API for an Org or Group](../../snyk-admin/user-management-with-the-snyk-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)

### [Search Group audit logs](../reference/audit-logs.md#groups-group\_id-audit\_logs-search)

**More information:** [Filter through your audit logs more efficiently with the new GA REST version of the audit logs API, and api.access is now opt-in](https://updates.snyk.io/filter-through-your-audit-logs-more-efficiently-with-the-new-ga-rest-version-of-the-audit-logs-api-and-api-access-is-now-opt-in-291850)

[Retrieve audit logs of user-initiated activity by API for an Org or Group](../../snyk-admin/user-management-with-the-snyk-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)

## Audit logs (v1)

### Group level audit logs

Use [Search Group audit log](../reference/audit-logs.md#groups-group\_id-audit\_logs-search)s

### Organization level audit logs

Use [Search Organization audit logs](../reference/audit-logs.md#orgs-org\_id-audit\_logs-search)

## Cloud (beta)

### [List Environments](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/cloud/environments)

### [Create New Environment](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#post-/orgs/-org\_id-/cloud/environments)

### [Delete Environment](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#delete-/orgs/-org\_id-/cloud/environments/-environment\_id-)

### [Update Environment](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#patch-/orgs/-org\_id-/cloud/environments/-environment\_id-)

### [Generate Cloud Provider Permissions](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#post-/orgs/-org\_id-/cloud/permissions)

### [List Resources](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/cloud/resources)

### [List Scans](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/cloud/scans)

### [Create Scan](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#post-/orgs/-org\_id-/cloud/scans)

### [Get scan](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/cloud/scans/-scan\_id-)

## Collection

### [Create a collection](../reference/collection.md#orgs-org\_id-collections)

### [Get collections](../reference/collection.md#orgs-org\_id-collections-1)

### [Edit a collection](../reference/collection.md#orgs-org\_id-collections-collection\_id)

### [Get a collection](../reference/collection.md#orgs-org\_id-collections-collection\_id-1)

### [Delete a collection](../reference/collection.md#orgs-org\_id-collections-collection\_id-2)

### [Add projects to a collection](../reference/collection.md#orgs-org\_id-collections-collection\_id-relationships-projects)

### [Get projects from the specified collection](../reference/collection.md#orgs-org\_id-collections-collection\_id-relationships-projects-1)

### [Remove projects from a collection](../reference/collection.md#orgs-org\_id-collections-collection\_id-relationships-projects-2)

## ContainerImage

### [List instances of container image](../reference/containerimage.md#orgs-org\_id-container\_images)

### [Get instance of container image](../reference/containerimage.md#orgs-org\_id-container\_images-image\_id)

### [List instances of image target references for a container image](../reference/containerimage.md#orgs-org\_id-container\_images-image\_id-relationships-image\_target\_refs)

## Custom Base Images

**More information:** [Use Custom Base Image Recommendations](../../scan-using-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations/)

### [Create a Custom Base Image from an existing container project](../reference/custom-base-images.md#custom\_base\_images)

**More information:** [Use Custom Base Image Recommendations: Mark the created Project as a custom base image](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations#mark-the-created-project-as-a-custom-base-image)

### [Get a custom base image collection](../reference/custom-base-images.md#custom\_base\_images-1)

### [Update a custom base image](../reference/custom-base-images.md#custom\_base\_images-custombaseimage\_id)

### [Get a custom base image](../reference/custom-base-images.md#custom\_base\_images-custombaseimage\_id-1)

### [Delete a custom base image](../reference/custom-base-images.md#custom\_base\_images-custombaseimage\_id-2)

## Dependencies (v1)

### [List all dependencies](../reference/dependencies-v1.md)

## Entitlements (v1)

### [List all entitlements](../reference/entitlements-v1.md#org-orgid-entitlements)

### [Get an organization's entitlement value](../reference/entitlements-v1.md#org-orgid-entitlement-entitlementkey)

## Groups (beta)

### [Get all groups](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/groups)

### [Get a group](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/groups/-group\_id-)

**More information:** [Org and group identification for Projects](undefined.md)

### [Get all SSO connections for a group](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/groups/-group\_id-/sso\_connections)

### [Get all users using a given SSO connection](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/groups/-group\_id-/sso\_connections/-sso\_id-/users)

### [Delete a user from a Group SSO connection](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#delete-/groups/-group\_id-/sso\_connections/-sso\_id-/users/-user\_id-)

**More information:** [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-snyk-api/remove-members-from-groups-and-orgs-using-the-api.md) and [Retrieve audit logs of user-initiated activity by API for an Org or Group](../../snyk-admin/user-management-with-the-snyk-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md).

## Groups (v1)

### [List all tags in a group](../reference/groups-v1.md#group-groupid-tags)

### [Delete tag from group](../reference/groups-v1.md#group-groupid-tags-delete)

### [Update group settings](../reference/groups-v1.md#group-groupid-settings)

### [View group settings](../reference/groups-v1.md#group-groupid-settings-1)

### [List all roles in a group](../reference/groups-v1.md#group-groupid-roles)

**More information :**\
[Update member roles using the V1 API](../../snyk-admin/user-management-with-the-snyk-api/update-member-roles-using-the-v1-api.md).

### [List all organizations in a group](../reference/groups-v1.md#group-groupid-orgs)

**More information:**\
[Org and group identification for Projects](undefined.md)

### [Add a member to an organization within a group](../reference/groups-v1.md#group-groupid-org-orgid-members)

### [List all members in a group](../reference/groups-v1.md#group-groupid-members)

**More information:** [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-snyk-api/remove-members-from-groups-and-orgs-using-the-api.md).

## IacSettings

### [Update the Infrastructure as Code Settings for an org](../reference/groups-v1.md#group-groupid-org-orgid-members)

**More information:** [Use a remote IaC custom rules bundle](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md), [Use a remote IaC custom rules bundle](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md)

### [Get the Infrastructure as Code Settings for a group](../reference/iacsettings.md#orgs-org\_id-settings-iac-1)

### [Update the Infrastructure as Code Settings for a group](../reference/iacsettings.md#groups-group\_id-settings-iac)

**More information:** [Use a remote IaC custom rules bundle](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md), [IaC custom rules within a pipeline](../../scan-using-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/iac-custom-rules-within-a-pipeline.md), [Use a remote IaC custom rules bundle](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md)

### [Get the Infrastructure as Code Settings for a group](../reference/iacsettings.md#groups-group\_id-settings-iac-1)

## Import Projects (v1)

### [Import targets](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)

If this fails, use [Get import job details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details) to help determine why. There are two types of failures:

* The repository was rejected for processing, that is, HTTP status code 201 was not returned. This happens if there is an issue Snyk can see quickly for example:
  * The repository does not exist.
  * The repository is unreachable by Snyk because the token is invalid or does not have sufficient permissions; there is no default branch.
* The repository was accepted for processing, that is, the user got back HTTP status code 201 and a url to poll, but no projects were detected or some failed. This may occur because:
  * There are no Snyk-supported manifests in this repository.
  * The repository is archived and the Snyk API calls to fetch files fail.
  * The individual project or manifest had issues during processing. In this case Snyk returns success: false with a message in the log.

The poll results return a message per manifest processed, either `success: true` or `success: false.`

### [Get import job details](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import-jobid)

## Integrations (v1)

### [Add new integration](../reference/integrations-v1.md#org-orgid-integrations)

### [List](../reference/integrations-v1.md#org-orgid-integrations-1)

### [Get existing integration by type](../reference/integrations-v1.md#org-orgid-integrations-type)

### [Update existing integration](../reference/integrations-v1.md#org-orgid-integrations-integrationid)

### [Update](../reference/integrations-v1.md#org-orgid-integrations-integrationid-settings)

### [Retrieve](../reference/integrations-v1.md#org-orgid-integrations-integrationid-settings-1)

### [Clone an integration (with settings and credentials)](../reference/integrations-v1.md#org-orgid-integrations-integrationid-clone)

### [Delete credentials](../reference/integrations-v1.md#org-orgid-integrations-integrationid-authentication)

### [Switch between broker tokens](../reference/integrations-v1.md#org-orgid-integrations-integrationid-authentication-switch-token)

### [Provision new broker token](../reference/integrations-v1.md#org-orgid-integrations-integrationid-authentication-provision-token)

## Invites

See also [Invite users](../reference/organizations-v1.md#org-orgid-invite).

### [Invite a user to an organization](../reference/invites.md#orgs-org\_id-invites)

### [List pending user invitation to an organization](../reference/invites.md#orgs-org\_id-invites-1)

### [Cancel a pending user invitations to an organization](../reference/invites.md#orgs-org\_id-invites-invite\_id)

## Issues

### [List issues for a package](../reference/issues.md#orgs-org\_id-packages-purl-issues)

### [List issues for a given set of packages](../reference/issues.md#orgs-org\_id-packages-issues) (Currently not available to all customers)

### [Get issues by org ID](../reference/issues.md#orgs-org\_id-issues)

### [Get an issue](../reference/issues.md#orgs-org\_id-issues-issue\_id) (Org)

### [Get issues by group ID](../reference/issues.md#orgs-org\_id-issues-issue\_id)

Note: Remedies are not included in the response.

### [Get an issue](../reference/issues.md#groups-group\_id-issues-issue\_id) (Group)

## Licenses (v1)

### [List all licenses](../reference/licenses-v1.md)

## Monitor (v1)

### [Dep Graph](../reference/monitor-v1.md)

**More information:** [Dep Graph API](../../supported-languages-package-managers-and-frameworks/bazel-tool/dep-graph-api.md)

## Organizations (v1)

### The Snyk organization for a request

#### List all the organizations a user belongs to

[Migrated List all the organizations a user belongs to](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to)

**More information:** [Org and group identification for Projects](undefined.md)

### Create organization

#### Create a new organization

[Migrated Create a new organization](https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization)

### Notification settings

#### Get organization notification settings

[Migrated Get organization notification settings](https://snyk.docs.apiary.io/#reference/organizations/notification-settings/get-organization-notification-settings)

#### Set notification settings

[Migrated Set notification settings](https://snyk.docs.apiary.io/#reference/organizations/notification-settings/set-notification-settings)

### User invitation to organization

#### Invite users

[Migrated Invite users](https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users)

### Members in organization

#### List members

[Migrated List members](https://snyk.docs.apiary.io/#reference/organizations/members-in-organization/list-members)

**More information:** [Update member roles using the V1 API](../../snyk-admin/user-management-with-the-snyk-api/update-member-roles-using-the-v1-api.md) and [Remove members from Groups and Orgs using the API.](../../snyk-admin/user-management-with-the-snyk-api/remove-members-from-groups-and-orgs-using-the-api.md)

### Organization settings

#### View organization settings

[Migrated View organization settings](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/view-organization-settings)

#### Update organization settings

[Migrated Update organization settings](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-organization-settings)

### Manage roles in organization

#### Update a member in the organization

[Migrated Update a member in the organization](https://snyk.docs.apiary.io/#reference/organizations/manage-roles-in-organization/update-a-member-in-the-organization)

#### Remove a member from the organization

[Migrated Remove a member from the organization](https://snyk.docs.apiary.io/#reference/organizations/manage-roles-in-organization/remove-a-member-from-the-organization)

**More information:** [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-snyk-api/remove-members-from-groups-and-orgs-using-the-api.md).

### Update member roles in your organization

#### Update a member's role in the organization

[Migrated Update a member's role in the organization](https://snyk.docs.apiary.io/#reference/organizations/update-member-roles-in-your-organization/update-a-member's-role-in-the-organization)

**More information:** [Update member roles using the V1 API](../../snyk-admin/user-management-with-the-snyk-api/update-member-roles-using-the-v1-api.md).

### Manage organization

#### Remove organization

[Migrated Remove organization](https://snyk.docs.apiary.io/#reference/organizations/manage-organization/remove-organization)

### Provision user

**More information:** [Provision users to Organizations using the V1 API](../../snyk-admin/user-management-with-the-snyk-api/provision-users-to-organizations-using-the-v1-api.md).

#### Provision a user to the organization

[Migrated Provision a user to the organization](https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization)

#### List pending user provisions

[Migrated List pending user provisions](https://snyk.docs.apiary.io/#reference/organizations/provision-user/list-pending-user-provisions)

#### Delete pending user provision

[Migrated Delete pending user provision](https://snyk.docs.apiary.io/#reference/organizations/provision-user/delete-pending-user-provision)

## Orgs (GA and beta)

### [List all organizations in a group](https://apidocs.snyk.io/?#get-/groups/-group\_id-/orgs)

### [List accessible organizations](https://apidocs.snyk.io/?#get-/orgs)

### [Get an ORG](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/orgs/-org\_id-) (beta)

### [Get organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-)

**More information:** [Org and group identification for Projects](undefined.md)

### [Update organization](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-)

## Projects

### [List all Projects for an Org with the given Org ID](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects)

The query-string parameter types is optional. The endpoint does not enforce specific project types and will return no matching projects if you enter a string that does not match a project type.

### [Get project by project ID](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects/-project\_id-)

### [Delete project by project ID](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/projects/-project\_id-)

### [Updates project by project ID](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/projects/-project\_id-)

## Projects (v1)

**More information:** [Project type responses from API](project-type-responses-from-api.md)

### Individual project

#### Retrieve a single project

[Migrated Retrieve a single project](https://snyk.docs.apiary.io/#reference/projects/individual-project/retrieve-a-single-project)

#### Update a project

[Migrated Update a project](https://snyk.docs.apiary.io/#reference/projects/individual-project/update-a-project)

#### Delete a project

[Migrated Delete a project](https://snyk.docs.apiary.io/#reference/projects/individual-project/delete-a-project)

### Deactivate an individual project

#### Deactivate

[Migrated Deactivate an individual project](https://snyk.docs.apiary.io/#reference/projects/deactivate-an-individual-project/deactivate)

### Activate an individual project

#### Activate

[Migrated Activate an individual project](https://snyk.docs.apiary.io/#reference/projects/activate-an-individual-project/activate)

### Aggregated Project issues

#### List all Aggregated issues

[Migrated List  all aggregated issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues)

The Snyk V1 API endpoint [List all aggregated issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues) returns an array of `ignoreReasons` for each vulnerability. This happens because ignores implemented using the CLI and API are path-based and thus potentially could have different `ignoreReasons` for different paths. Because List all aggregated issues returns only one issue for all paths, the entire set of reasons is returned. Snyk groups issues together by their identifier, so one response for the List all aggregated issues endpoint could correspond to the same issue across multiple paths. Thus the `ignoredReason` is across all issues that are aggregated and applies to that single grouped issue.

### Project Issue Paths

**More information:** [V1 API Project issue paths endpoints](project-issue-paths-v1-api-endpoints.md)

#### List all project issue paths

[Migrated List all project issue paths](https://snyk.docs.apiary.io/#reference/projects/project-issue-paths/list-all-project-issue-paths)

### Project History

#### List all project snapshots

[Migrated List all project snapshots](https://snyk.docs.apiary.io/#reference/projects/project-history/list-all-project-snapshots)

### Aggregated Project Snapshot Issues

#### List all project snapshot aggregated issues

[Migrated List all project snapshot aggregated issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-snapshot-issues/list-all-project-snapshot-aggregated-issues)

### Project Snapshot Issue Paths

#### List all project snapshot issue paths

[Migrated List all project snapshot issue paths](https://snyk.docs.apiary.io/#reference/projects/project-snapshot-issue-paths/list-all-project-snapshot-issue-paths)

### Project dependency graph

#### Get Project dependency graph

[Migrated Get Project dependency graph](https://snyk.docs.apiary.io/#reference/projects/project-dependency-graph/get-project-dependency-graph)

### Project ignores

#### List all ignores

[Migrated List all ignores](https://snyk.docs.apiary.io/#reference/projects/project-ignores/list-all-ignores)

### Ignored issues

#### Retrieve ignore

[Migrated Retrieve ignore](https://snyk.docs.apiary.io/#reference/projects/ignored-issues/retrieve-ignore)

#### Add ignore

[Migrated Add ignore](https://snyk.docs.apiary.io/#reference/projects/ignored-issues/add-ignore)

#### Replace ignores

[Migrated Replace ignores](https://snyk.docs.apiary.io/#reference/projects/ignored-issues/replace-ignores)

#### Delete ignores

[Migrated Delete ignores](https://snyk.docs.apiary.io/#reference/projects/ignored-issues/delete-ignores)

### Project jira issues

#### List all jira issues

[Migrated List all jira issues](https://snyk.docs.apiary.io/#reference/projects/project-jira-issues/list-all-jira-issues)

#### Create jira issue

[Migrated Create Jira issue](https://snyk.docs.apiary.io/#reference/projects/project-jira-issues/create-jira-issue)

### Project settings

#### List project settings

[Migrated List project settings](https://snyk.docs.apiary.io/#reference/projects/project-settings/list-project-settings)

#### Update project settings

[Migrated Update project settings](https://snyk.docs.apiary.io/#reference/projects/project-settings/update-project-settings)

#### Delete project settings

[Migrated Delete project settings](https://snyk.docs.apiary.io/#reference/projects/project-settings/delete-project-settings)

### Move project

#### Move project to a different organization

[Migrated Move project to a different organization](https://snyk.docs.apiary.io/#reference/projects/move-project/move-project-to-a-different-organization)

### Project tags

#### Add a tag to a project

[Migrated Add a tag to a project](https://snyk.docs.apiary.io/#reference/projects/project-tags/add-a-tag-to-a-project)

### Remove project tag

#### Remove a tag from a project

[Migrated Remove a tag from a project](https://snyk.docs.apiary.io/#reference/projects/remove-project-tag/remove-a-tag-from-a-project)

### Project Attributes

#### Applying attributes

By using the Snyk API v1 endpoint [Applying attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes) you can set attributes for Snyk Projects including business criticality, lifecycle stage, and environment once the project has been created . To do so:

* Import the project using the Snyk API v1 endpoint [Import targets](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets).
* Get the status API ID from Import targets.
* Poll using [Import job details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details) until all imports have completed.
* Parse the project IDs from the projectURL field.
* Use the [Applying attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes) endpoint to set the project attributes.

[Migrated Applying attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes)

## Pull request templates

**More information:** [Create and manage a custom PR template using the API](../../scan-with-snyk/pull-requests/snyk-fix-pull-or-merge-requests/customize-pr-templates/apply-a-custom-pr-template.md#create-and-manage-a-custom-pr-template-using-the-api).

### [Get pull request template for group](https://apidocs.snyk.io/?#get-/groups/-group\_id-/settings/pull\_request\_template)

### [Create or update pull request template for group](https://apidocs.snyk.io/?version=2024-05-23#delete-/groups/-group\_id-/settings/pull\_request\_template)

### [Delete pull request template for group](https://apidocs.snyk.io/?version=2024-05-23#delete-/groups/-group\_id-/settings/pull\_request\_template)

## Reporting API (v1)

### Latest issues

#### Get list of latest issues

To list all projects that have a vulnerability linked to a CVE use the capability to filter on strings with the [Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues) and [Get List of issues](https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues) reporting endpoints. Filter by the identifier attribute.

To get a list of issues that have been fixed: Use [Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues) and filter by `“isFixed”: true` in the request body. This endpoint also provides a [list of all IaC issues](../../scan-with-snyk/snyk-iac/view-snyk-iac-issue-reports.md#api-access-to-iac-issues).

[Migrated Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues)

### Issues

#### Get list of issues

To list all projects that have a vulnerability linked to a CVE use the capability to filter on strings with the [Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues) and [Get List of issues](https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues) (reporting) endpoints. Filter by the identifier attribute.

[Migrated Get list of issues](https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues)

### Latest issue counts

#### Get latest issue counts

[Migrated latest issue counts](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issue-counts/get-latest-issue-counts)

### Issue counts over time

#### Get issue counts

[Migrated Get Issue counts](https://snyk.docs.apiary.io/#reference/reporting-api/issue-counts-over-time/get-issue-counts)

### Latest project counts

#### Get latest project counts

[Migrated Get latest project counts](https://snyk.docs.apiary.io/#reference/reporting-api/latest-project-counts/get-latest-project-counts)

### Project counts over time&#x20;

#### Get project counts

[Migrated Get project counts](https://snyk.docs.apiary.io/#reference/reporting-api/project-counts-over-time/get-project-counts)

### Test counts

#### Get test counts

[Migrated test counts](https://snyk.docs.apiary.io/#reference/reporting-api/test-counts/get-test-counts)

## SBOM (GA and beta)

### [Get a project’s SBOM document](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/projects/-project\_id-/sbom)

### [Create an SBOM test run](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#post-/orgs/-org\_id-/sbom\_tests) (beta)

### [Gets an SBOM test run status](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-) (beta)

### [Gets an SBOM test run result](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-/results) (beta)

## SastSettings

### [Retrieves the SAST settings for an org](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/settings/sast)

### [Enable/Disable the Snyk Code settings for an org](https://apidocs.snyk.io/?#patch-/orgs/-org\_id-/settings/sast)

## ServiceAccounts

**More information:** [Manage service accounts using the Snyk API](../../enterprise-configuration/service-accounts/manage-service-accounts-using-the-snyk-api.md); [Choose a service account type to use with Snyk APIs](../../enterprise-configuration/service-accounts/choose-a-service-account-type-to-use-with-snyk-apis.md)

### [Get a list of group service accounts](https://apidocs.snyk.io/?#get-/groups/-group\_id-/service\_accounts)

### [Create a service account for a group](https://apidocs.snyk.io/?#post-/groups/-group\_id-/service\_accounts)

### [Get a group service account](https://apidocs.snyk.io/?#get-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

### [Delete a group service account](https://apidocs.snyk.io/?#delete-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

### [Update a group service account](https://apidocs.snyk.io/?#patch-/groups/-group\_id-/service\_accounts/-serviceaccount\_id-)

### [Manage a group service account’s client secret](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/service\_accounts)

### [Get a list of organization service accounts](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/service\_accounts)

### [Create a service account for an organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

### [Get an organization service account](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

### [Delete a service account in an organization](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

### [Update an organization service account](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

### [Manage an organization service account’s client secret](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-/secrets)

## Slack

### [Get a list of Slack channels](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/slack\_app/-tenant\_id-/channels)

### [Get Slack Channel name by Slack Channel ID](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/slack\_app/-tenant\_id-/channels/-channel\_id-)

## SlackSettings

### [Get Slack integration default notification settings](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/slack\_app/-bot\_id-)

### [Create new Slack notification default settings](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/slack\_app/-bot\_id-)

### [Remove the given Slack App integration](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/slack\_app/-bot\_id-)

### [Slack notification settings overrides for projects](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/slack\_app/-bot\_id-/projects)

### [Create a new Slack settings override for a given project](https://apidocs.snyk.io/?#post-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-)

### [Remove Slack settings override for a project](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-)

### [Update Slack notification settings for a project](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/slack\_app/-bot\_id-/projects/-project\_id-)

## Targets

### [Get targets by org ID](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/targets)&#x20;

**More information:** [Target definition on the Projects page](../../snyk-admin/snyk-projects/#target)

### [Get target by target ID](https://apidocs.snyk.io/?#get-/orgs/-org\_id-/targets/-target\_id-)

### [Delete target by target ID](https://apidocs.snyk.io/?#delete-/orgs/-org\_id-/targets/-target\_id-)

## Test (v1)

### Maven

#### Test for issues in a public package by group id, artifact id and version

[Migrated Test for issues in a public package by group id, artfact id and version](https://snyk.docs.apiary.io/#reference/test/maven/test-for-issues-in-a-public-package-by-group-id,-artifact-id-and-version)

#### Test maven file

[Migrated Test maven file](https://snyk.docs.apiary.io/#reference/test/maven/test-maven-file)

### npm

#### Test for issues in a public package by name and version

[Migrated Test for issues in a public package by name and version](https://snyk.docs.apiary.io/#reference/test/npm/test-for-issues-in-a-public-package-by-name-and-version)

#### Test package.json & package-lock.json File

[Migrated Test package.json & package-lock.json file](https://snyk.docs.apiary.io/#reference/test/npm/test-package.json-&-package-lock.json-file)

### dep

#### Test Gopkg.toml & Gopkg.lock File

[Migrated Test Gopkg.toml & Gopkg.lock file](https://snyk.docs.apiary.io/#reference/test/dep/test-gopkg.toml-&-gopkg.lock-file)

### vendor

#### Test vendor.json file

[Migrated Test vendor.json file](https://snyk.docs.apiary.io/#reference/test/vendor/test-vendor.json-file)

### yarn

#### Test package.json & yarn.lock file

[Migrated Test package.json & yarn.lock file](https://snyk.docs.apiary.io/#reference/test/yarn/test-package.json-&-yarn.lock-file)

### rubygems

#### Test for issues in a public gem by name and version

[Migrated Test for issues in a public gem by name and version](https://snyk.docs.apiary.io/#reference/test/rubygems/test-for-issues-in-a-public-gem-by-name-and-version)

#### Test gemfile.lock file

[Migrated Test gemfile.lock file](https://snyk.docs.apiary.io/#reference/test/rubygems/test-gemfile.lock-file)

### Gradle

#### Test for issues in a public package by group, name and version

[Migrated Test for issues in a public package by group, name and version](https://snyk.docs.apiary.io/#reference/test/gradle/test-for-issues-in-a-public-package-by-group,-name-and-version)

#### Test gradle file

[Migrated test gradle file](https://snyk.docs.apiary.io/#reference/test/gradle/test-gradle-file)

### sbt

#### Test for issues in a public package by group id, artifact id and version

[Migrated Test for issues in a public package by gorup id, artifact id and version](https://snyk.docs.apiary.io/#reference/test/sbt/test-for-issues-in-a-public-package-by-group-id,-artifact-id-and-version)

#### Test sbt file

[Migrated Test sbt file](https://snyk.docs.apiary.io/#reference/test/sbt/test-sbt-file)

### pip

#### Test for issues in a public package by name and version

[Migrated Test for issues in a public package by name and version](https://snyk.docs.apiary.io/#reference/test/pip/test-for-issues-in-a-public-package-by-name-and-version)

#### Test requirements.txt file

[Migrated Test requirements.txt file](https://snyk.docs.apiary.io/#reference/test/pip/test-requirements.txt-file)

### composer

#### Test composer.json & composer.lock file

[Migrated Test composer.json & composer.lock file](https://snyk.docs.apiary.io/#reference/test/composer/test-composer.json-&-composer.lock-file)

### Dep Graph

**More information:** [Dep Graph API](../../supported-languages-package-managers-and-frameworks/bazel-tool/dep-graph-api.md)

#### Test Dep Graph

[Migrated Test Dep Graph](https://snyk.docs.apiary.io/#reference/test/dep-graph/test-dep-graph)

## Users

### [Update a user’s role in a group](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#patch-/groups/-group\_id-/users/-id-) (beta)

Note: Use this endpoint to remove users from a group.

**More information:** [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-snyk-api/remove-members-from-groups-and-orgs-using-the-api.md).

### [Get user by ID](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/users/-id-) (beta)

### [My User Details](https://apidocs.snyk.io/?#get-/self)

## Users (v1)

### User Details

#### Get user details

[Migrated Get user details](https://snyk.docs.apiary.io/#reference/users/user-details/get-user-details)

### My User Details

#### Get My Details

[Migrated Get My Details](https://snyk.docs.apiary.io/#reference/users/my-user-details/get-organization-notification-settings)

### User organization notification settings

#### Get organization notification settings

[Migrated Get organization notification settings](https://snyk.docs.apiary.io/#reference/users/user-organization-notification-settings/get-organization-notification-settings)

#### Modify organization notification settings

[Migrated Modify organization notification settings](https://snyk.docs.apiary.io/#reference/users/user-organization-notification-settings/modify-organization-notification-settings)

### User project notification settings

#### Get project notification settings

[Migrated Get project notification settings](https://snyk.docs.apiary.io/#reference/users/user-project-notification-settings/get-project-notification-settings)

#### Modify project notification settings

[Migrated Modify project notification settings](https://snyk.docs.apiary.io/#reference/users/user-project-notification-settings/modify-project-notification-settings)

## Webhooks (v1)

### Webhook Collection

#### Create a webhook

[Migrated Create a webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection)

#### List webhooks

[Migrated List webhooks](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/create-a-webhook)

### Webhook

#### Retrieve a webhook

[Migrated Retrieve a webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook/retrieve-a-webhook)

#### Delete a webhook

[Migrated Delete a webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook/delete-a-webhook)

### Ping

#### Ping a webhook

[Migrated Ping a webhook](https://snyk.docs.apiary.io/#reference/webhooks/ping/ping-a-webhook)



\
