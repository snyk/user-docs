# API endpoints index and notes



{% hint style="info" %}
This index and notes section of the documentation provides, in addition to this index, [solutions for specific use cases](solutions-for-specific-use-cases.md), [scenarios for using Snyk APIs](scenarios-for-using-snyk-api.md), and pages with detailed information about using Snyk API endpoints:

* [Organization and Group identification for Projects using the API](undefined.md)
* [Project issue paths V1 API endpoints](project-issue-paths-api-endpoints.md)
* [Project type responses from the API](project-type-responses-from-api.md)

See also the following sections on specific APIs:

* [How to use Snyk Apps APIs](../how-to-use-snyk-apps-apis/)
* [How to use Snyk SBOM and List issues APIs](../how-to-use-snyk-sbom-and-list-issues-apis/)
* [How to use Snyk webhooks APIs](../how-to-use-snyk-webhooks-apis/)

For more information about using the Snyk API, see the [API support articles](https://support.snyk.io/hc/en-us/sections/360001344097-API).
{% endhint %}

This index includes the categories and names of REST GA and beta and V1 API endpoints, with the URL in the reference docs for each endpoint, and links to related information where available. REST is the default, and GA is the status unless beta is noted. V1 API is specified where applicable. This index is a work in progress; additional information is being added continually.

## AccessRequests (beta)

### [Get access requests](https://apidocs.snyk.io/?beta=\&version=2024-07-10%7Ebeta#get-/self/access\_requests)

## Apps

**More information:** [Snyk Apps](../how-to-use-snyk-apps-apis/)

### [Get a list of apps that act on your behalf](../reference/apps.md#self-apps)

### [Revoke an app](../reference/apps.md#self-apps-app\_id)

### [Get a list of active OAuth sessions for the app](../reference/apps.md#self-apps-app\_id-sessions)

### [Revoke an active user app session](../reference/apps.md#self-apps-app\_id-sessions-session\_id)

### [Get a list of apps installed for a user](../reference/apps.md#self-apps-installs)

### [Revoke access for an app by install ID](../reference/apps.md#self-apps-installs-install\_id)

**Replaces:** DEPRECATED Revoke app bot authorization

### DEPRECATED [Create a new app for an organization](../reference/apps.md#orgs-org\_id-apps)

**Replaced by:** Create a new Snyk App for an organization

**More information:** [Create a Snyk App using the Snyk API](../how-to-use-snyk-apps-apis/create-a-snyk-app-using-the-snyk-api.md)

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

**More information:** [Slack app](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (Find the Slack App Bot ID)

### [Revoke app authorization for a Snyk organization](../reference/apps.md#orgs-org\_id-apps-installs-install\_id)

**See also:** Revoke app authorization for a Snyk Group with install ID

### [Manage client secret for non-interactive Snyk App installations](../reference/apps.md#orgs-org\_id-apps-installs-install\_id-secrets)

### [Create a new Snyk App for an organization](../reference/apps.md#orgs-org\_id-apps-creations)

**Replaces:** DEPRECATED Create a new app for an organization

### DEPRECATED [Get a list of apps created by an organization](../reference/apps.md#orgs-org\_id-apps-creations-1)

**Replaced by:** Get a list of apps created by an organization

**More information:** [Manage App details](../how-to-use-snyk-apps-apis/manage-app-details.md)

### [Update app creation attributes such as name, redirect URIs, and access token time to live using the App ID](../reference/apps.md#orgs-org\_id-apps-creations-app\_id)

**Replaces:** DEPRECATED Update App attributes that are name, redirect URIs, and access token time to live

**More information:** [Manage App details](../how-to-use-snyk-apps-apis/manage-app-details.md)

### [Get a Snyk APP by its App ID](../reference/apps.md#orgs-org\_id-apps-creations-app\_id)

**Replaces:** DEPRECATED Get an app by client id

### [Delete a Snyk App by its App ID](../reference/apps.md#orgs-org\_id-apps-creations-app\_id-2)

**Replaces:** DEPRECATED Delete an app

**More information:** [Manage App details](../how-to-use-snyk-apps-apis/manage-app-details.md)

### [Manage client secret for the Snyk App](../reference/apps.md#orgs-org\_id-apps-creations-app\_id-secrets)

**More information:** [Manage App details](../how-to-use-snyk-apps-apis/manage-app-details.md)

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

**More information**: [Retrieve audit logs of user-initiated activity by API for an Org or Group](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)

### [Search Organization audit logs](../reference/audit-logs.md#orgs-org\_id-audit\_logs-search)

**More information:** [Retrieve audit logs of user-initiated activity by API for an Org or Group](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md), [AWS CloudTrail Lake](../../integrate-with-snyk/event-forwarding/aws-cloudtrail-lake.md)

### [Search Group audit logs](../reference/audit-logs.md#groups-group\_id-audit\_logs-search)

**More information:** [Filter through your audit logs more efficiently with the new GA REST version of the audit logs API](https://updates.snyk.io/filter-through-your-audit-logs-more-efficiently-with-the-new-ga-rest-version-of-the-audit-logs-api-and-api-access-is-now-opt-in-291850) (product update); [Retrieve audit logs of user-initiated activity by API for an Org or Group](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)

## Audit logs (v1)

### Group level audit logs

Use [Search Group audit logs](../reference/audit-logs.md#groups-group\_id-audit\_logs-search)

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

**More information:** [Use Custom Base Image Recommendations](../../scan-with-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations/)

### [Create a Custom Base Image from an existing container project](../reference/custom-base-images.md#custom\_base\_images)

**More information:** [Mark the created Project as a custom base image](../../scan-with-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations/#mark-the-created-project-as-a-custom-base-image) (Use Custom Base Image Recommendations)

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

**More information:** [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md); [Retrieve audit logs of user-initiated activity by API for an Org or Group](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)

## Groups (v1)

### [List all tags in a group](../reference/groups-v1.md#group-groupid-tags)

### [Delete tag from group](../reference/groups-v1.md#group-groupid-tags-delete)

### [Update group settings](../reference/groups-v1.md#group-groupid-settings)

### [View group settings](../reference/groups-v1.md#group-groupid-settings-1)

### [List all roles in a group](../reference/groups-v1.md#group-groupid-roles)

**More information:** [Update member roles using the API](../../snyk-admin/user-management-with-the-api/update-member-roles-using-the-api.md); [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [List all organizations in a group](../reference/groups-v1.md#group-groupid-orgs)

**More information:** [Org and group identification for Projects](undefined.md); [Legacy custom mapping](../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/custom-mapping/legacy-custom-mapping.md)

### [Add a member to an organization within a group](../reference/groups-v1.md#group-groupid-org-orgid-members)

### [List all members in a group](../reference/groups-v1.md#group-groupid-members)

**More information:** [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md).

## IacSettings

### [Update the Infrastructure as Code Settings for an org](../reference/groups-v1.md#group-groupid-org-orgid-members)

**More information:** [Use a remote IaC custom rules bundle](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md); [Use a remote IaC custom rules bundle](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md)

### [Get the Infrastructure as Code Settings for a group](../reference/iacsettings.md#orgs-org\_id-settings-iac-1)

### [Update the Infrastructure as Code Settings for a group](../reference/iacsettings.md#groups-group\_id-settings-iac)

**More information:** [Use a remote IaC custom rules bundle](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md), [IaC custom rules within a pipeline](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/iac-custom-rules-within-a-pipeline.md);[Use a remote IaC custom rules bundle](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md)

### [Get the Infrastructure as Code Settings for a group](../reference/iacsettings.md#groups-group\_id-settings-iac-1)

## Ignores (v1)

**More information:** [Snyk test and snyk monitor in CI/CD integration](../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md)

### [List all ignores](../reference/ignores-v1.md#org-orgid-project-projectid-ignores)

### [Replace ignores](../reference/ignores-v1.md#org-orgid-project-projectid-ignore-issueid)

### [Add ignore](../reference/ignores-v1.md#org-orgid-project-projectid-ignore-issueid-1)

### [Retrieve ignore](../reference/ignores-v1.md#org-orgid-project-projectid-ignore-issueid-1)

### [Delete ignores](../reference/ignores-v1.md#org-orgid-project-projectid-ignore-issueid-3)

## Import Projects (v1)

### [Import targets](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)

For information on when and how you can use this endpoint, see this page on i[mporting targets](../../implement-snyk/enterprise-implementation-guide/phase-3-gain-visibility/import-projects.md#git-integration).

If this fails, use [Get import job details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details) to help determine why. There are two types of failures:

* The repository was rejected for processing, that is, HTTP status code 201 was not returned. This happens if there is an issue Snyk can see quickly for example:
  * The repository does not exist.
  * The repository is unreachable by Snyk because the token is invalid or does not have sufficient permissions; there is no default branch.
* The repository was accepted for processing, that is, the user got back HTTP status code 201 and a url to poll, but no projects were detected or some failed. This may occur because:
  * There are no Snyk-supported manifests in this repository.
  * The repository is archived and the Snyk API calls to fetch files fail.
  * The individual project or manifest had issues during processing. In this case Snyk returns success: false with a message in the log.

The poll results return a message per manifest processed, either `success: true` or `success: false.`

**More information:** [Snyk Broker Code Agent](../../enterprise-setup/snyk-broker/snyk-broker-code-agent/), [Configure integrations](../../implement-snyk/team-implementation-guide/phase-2-configure-your-organization/configure-integrations.md) (Enterprise implementation guide, Phase 2), [Import Projects](../../implement-snyk/team-implementation-guide/phase-3-gain-visibility/import-projects.md) (Enterprise implementation guide, Phase 3)

### [Get import job details](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import-jobid)

## Integrations (v1)

### [Add new integration](../reference/integrations-v1.md#org-orgid-integrations)

### [List](../reference/integrations-v1.md#org-orgid-integrations-1)

### [Get existing integration by type](../reference/integrations-v1.md#org-orgid-integrations-type)

### [Update existing integration](../reference/integrations-v1.md#org-orgid-integrations-integrationid)

**More information:** [Obtain the required tokens for setup](../../enterprise-setup/snyk-broker/snyk-broker-code-agent/install-snyk-broker-code-agent-using-docker/obtain-the-required-tokens-for-setup.md) (Snyk Broker Code Agent)

### [Update](../reference/integrations-v1.md#org-orgid-integrations-integrationid-settings)

### [Retrieve](../reference/integrations-v1.md#org-orgid-integrations-integrationid-settings-1)

### [Clone an integration (with settings and credentials)](../reference/integrations-v1.md#org-orgid-integrations-integrationid-clone)

**More information:** [Obtain the required tokens for setup](../../enterprise-setup/snyk-broker/snyk-broker-code-agent/install-snyk-broker-code-agent-using-docker/obtain-the-required-tokens-for-setup.md) (Snyk Broker Code Agent)

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

**More information:** [Dart and Flutter](../../supported-languages-package-managers-and-frameworks/dart-and-flutter.md), [Rust](../../supported-languages-package-managers-and-frameworks/rust.md), [Guidance for Snyk for C++ page, Alternate testing options section](../../supported-languages-package-managers-and-frameworks/c-c++/guidance-for-snyk-for-c-c++.md#alternate-testing-options), [Guidance for Java and Kotlin](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/guidance-for-java-and-kotlin.md), [Guidance for JavaScript and Node.js, Unmanaged JavaScript section](../../supported-languages-package-managers-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#unmanaged-javascript), [List issues for a package page](../how-to-use-snyk-sbom-and-list-issues-apis/list-issues-for-a-package.md)

### [List issues for a given set of packages](../reference/issues.md#orgs-org\_id-packages-issues) (Currently not available to all customers)

### [Get issues by org ID](../reference/issues.md#orgs-org\_id-issues)

### [Get an issue](../reference/issues.md#orgs-org\_id-issues-issue\_id) (Org)

### [Get issues by group ID](../reference/issues.md#orgs-org\_id-issues-issue\_id)

**Note:** Remedies are not included in the response.

More information: [ Reachability](../../manage-risk/prioritize-issues-for-fixing/reachability-analysis.md)

### [Get an issue](../reference/issues.md#groups-group\_id-issues-issue\_id) (Group)

## Licenses (v1)

### [List all licenses](../reference/licenses-v1.md)

## Monitor (v1)

### [Dep Graph](../reference/monitor-v1.md)

**More information:** [Dep Graph API](../../supported-languages-package-managers-and-frameworks/bazel/dep-graph-api.md)

## Organizations (v1)

### [List all the organizations a user belongs to](../reference/organizations-v1.md#orgs)

**More information:** [Org and group identification for Projects](undefined.md)

### [Create a new organization](../reference/organizations-v1.md#org)

**More information:** [Set visibility and configure an Organization template](../../implement-snyk/enterprise-implementation-guide/phase-2-configure-account/set-visibility-and-configure-an-organization-template/) (Enterprise implementation guide Phase 2)

### [Remove organization](../reference/organizations-v1.md#org-orgid)

### [Update organization settings](../reference/organizations-v1.md#org-orgid-settings)

### [View organization settings](../reference/organizations-v1.md#org-orgid-settings-1)

### [Provision a user to the organization](../reference/organizations-v1.md#org-orgid-provision)

**More information:** [Provision users to Organizations using the API](../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md); [Configure SSO](../../implement-snyk/enterprise-implementation-guide/phase-2-configure-account/configure-sso.md)

### [List pending user provisions](../reference/organizations-v1.md#org-orgid-provision-1)

**More information:** [Provision users to Organizations using the API](../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md)

### [Delete pending user provision](../reference/organizations-v1.md#org-orgid-provision-2)

**More information:** [Provision users to Organizations using the API](../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md)

### [Set notification settings](../reference/organizations-v1.md#org-orgid-notification-settings)

### [Get organization notification settings](../reference/organizations-v1.md#org-orgid-notification-settings-1)

### [List members](../reference/organizations-v1.md#org-orgid-members)

**More information:** [Update member roles using the API](../../snyk-admin/user-management-with-the-api/update-member-roles-using-the-api.md); [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md)

### [Update a member in the organization](../reference/organizations-v1.md#org-orgid-members-userid)

**More information:** [User role management](../../snyk-admin/user-roles/user-role-management.md)

### [Remove a member from the organization](../reference/organizations-v1.md#org-orgid-members-userid-1)

**More information:** [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md); [User role management](../../snyk-admin/user-roles/user-role-management.md)

### [Update a member's role in the organization](../reference/organizations-v1.md#org-orgid-members-update-userid)

**More information:** [User role management](../../snyk-admin/user-roles/user-role-management.md); [Update member roles using the API](../../snyk-admin/user-management-with-the-api/update-member-roles-using-the-api.md)

### [Invite users](../reference/organizations-v1.md#org-orgid-invite)

**More information:** [Update member roles using the API](../../snyk-admin/user-management-with-the-api/update-member-roles-using-the-api.md)

## Orgs (GA and beta)

### [List accessible organizations](../reference/orgs.md#orgs)

### [Update organization](../reference/orgs.md#orgs-org\_id)

### [Get organization](../reference/orgs.md#orgs-org\_id-1)

### [List all organizations in a group](../reference/orgs.md#groups-group\_id-orgs)

### [Get an ORG](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/orgs/-org\_id-) (beta)

**More information:** [Org and group identification for Projects](undefined.md)

## Projects (v1)

**More information:** [Project type responses from API](project-type-responses-from-api.md)

### [Update a project](../reference/projects-v1.md#org-orgid-project-projectid)

### [Retrieve a single project](../reference/projects-v1.md#org-orgid-project-projectid-1)

### [Delete a project](../reference/projects-v1.md#org-orgid-project-projectid-2)

### [Add a tag to a project](../reference/projects-v1.md#org-orgid-project-projectid-tags)

### [Remove a tag from a project](../reference/projects-v1.md#org-orgid-project-projectid-tags-remove)

### [Update project settings](../reference/projects-v1.md#org-orgid-project-projectid-settings)

### [List project settings](../reference/projects-v1.md#org-orgid-project-projectid-settings-1)

### [Delete project settings](../reference/projects-v1.md#org-orgid-project-projectid-settings-2)

### [Move project to a different organization](../reference/projects-v1.md#org-orgid-project-projectid-move)

### [List all jira issues](../reference/projects-v1.md#org-orgid-project-projectid-jira-issues)

**More information:** [Jira integration](../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md); [Snyk test and snyk monitor in CI/CD integration](../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md)

### [List all project issue paths](../reference/projects-v1.md#org-orgid-project-projectid-issue-issueid-paths)

**More information:** [Project issue paths API endpoints](project-issue-paths-api-endpoints.md)

### [Create jira issue](../reference/projects-v1.md#org-orgid-project-projectid-issue-issueid-jira-issue)

**More information:** [Jira integration](../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md); [Snyk test and snyk monitor in CI/CD integration](../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md)

### [List all (project) ignores](../reference/projects-v1.md#org-orgid-project-projectid-ignores)

### [Replace ignores](../reference/projects-v1.md#org-orgid-project-projectid-ignore-issueid)

### [Add ignore](../reference/projects-v1.md#org-orgid-project-projectid-ignore-issueid-1)

### [Retrieve ignore](../reference/projects-v1.md#org-orgid-project-projectid-ignore-issueid-2)

### [Delete ignores](../reference/projects-v1.md#org-orgid-project-projectid-ignore-issueid-2)

### [List all project snapshots](../reference/projects-v1.md#org-orgid-project-projectid-history)

### [List all project snapshot issue paths](../reference/projects-v1.md#org-orgid-project-projectid-history-snapshotid-issue-issueid-paths)

### [List all project snapshot aggregated issues](../reference/projects-v1.md#org-orgid-project-projectid-history-snapshotid-aggregated-issues)

### [Get Project dependency graph](../reference/projects-v1.md#org-orgid-project-projectid-dep-graph)

### [Deactivate](../reference/projects-v1.md#org-orgid-project-projectid-deactivate) (a project)

### [Applying (project) attributes](../reference/projects-v1.md#org-orgid-project-projectid-attributes)

By using the Snyk API v1 endpoint [Applying attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes) you can set attributes for Snyk Projects including business criticality, lifecycle stage, and environment once the project has been created . To do so:

* Import the project using the Snyk API v1 endpoint [Import targets](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets).
* Get the status API ID from Import targets.
* Poll using [Import job details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details) until all imports have completed.
* Parse the project IDs from the projectURL field.
* Use the [Applying attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes) endpoint to set the project attributes.

### [List all Aggregated (Project) issues](../reference/projects-v1.md#org-orgid-project-projectid-aggregated-issues)

The Snyk V1 API endpoint [List all aggregated issues](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues) returns an array of `ignoreReasons` for each vulnerability. This happens because ignores implemented using the CLI and API are path-based and thus potentially could have different `ignoreReasons` for different paths. Because List all aggregated issues returns only one issue for all paths, the entire set of reasons is returned. Snyk groups issues together by their identifier, so one response for the List all aggregated issues endpoint could correspond to the same issue across multiple paths. Thus the `ignoredReason` is across all issues that are aggregated and applies to that single grouped issue.

### [Activate](../reference/projects-v1.md#org-orgid-project-projectid-activate) (a project)

## Projects

### [List all Projects for an Org with the given Org ID](../reference/projects.md#orgs-org\_id-projects)

The query-string parameter types is optional. The endpoint does not enforce specific project types and will return no matching projects if you enter a string that does not match a project type.

**More information:** [Slack apps](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (Find your Project ID)

### [Updates project by project ID](../reference/projects.md#orgs-org\_id-projects-project\_id)

### [Get project by project ID](../reference/projects.md#orgs-org\_id-projects-project\_id-1)

### [Delete project by project ID](../reference/projects.md#orgs-org\_id-projects-project\_id-2)

## Pull request templates

### [Create or update pull request template for group](../reference/pull-request-templates.md#groups-group\_id-settings-pull\_request\_template)

**More information:** [Create and manage a custom PR template using the API](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/customize-pr-templates/apply-a-custom-pr-template.md#create-and-manage-a-custom-pr-template-using-the-api)

### [Get pull request template for group](../reference/pull-request-templates.md#groups-group\_id-settings-pull\_request\_template-1)

### [Delete pull request template for group](../reference/pull-request-templates.md#groups-group\_id-settings-pull\_request\_template-2)

## Reporting API (v1)

### [Get list of latest issues](../reference/reporting-api-v1.md#reporting-issues-latest)

To list all Projects that have a vulnerability linked to a CVE, use the capability to filter on strings with the reporting endpoints [Get list of latest issues](../reference/reporting-api-v1.md#reporting-issues-latest) and [Get List of issues](../reference/reporting-api-v1.md#reporting-issues). Filter by the identifier attribute.

To get a list of issues that have been fixed, use the endpoint [Get list of latest issues](../reference/reporting-api-v1.md#reporting-issues-latest) and filter by `“isFixed”: true` in the request body. This endpoint also provides a [list of all IaC issues](../../scan-with-snyk/snyk-iac/view-snyk-iac-issue-reports.md#api-access-to-iac-issues).

### [Get list of issues](../reference/reporting-api-v1.md#reporting-issues)

See notes for [Get list of latest issues](./#get-list-of-latest-issues).

### [Get test counts](../reference/reporting-api-v1.md#reporting-counts-tests)

### [Get project counts](../reference/reporting-api-v1.md#reporting-counts-projects)

### [Get latest project counts](../reference/reporting-api-v1.md#reporting-counts-projects-latest)

### [Get issue counts](../reference/reporting-api-v1.md#reporting-counts-issues)

### [Get latest issue counts](../reference/reporting-api-v1.md#reporting-counts-issues-latest)

## SBOM (GA and beta)

**More information:** [Rust](../../supported-languages-package-managers-and-frameworks/rust.md), [SBOM test endpoints](../how-to-use-snyk-sbom-and-list-issues-apis/test-an-sbom-document-for-vulnerabilities.md)

### [Get a project’s SBOM document](../reference/sbom.md)

### [Create an SBOM test run](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#post-/orgs/-org\_id-/sbom\_tests) (beta)

### [Gets an SBOM test run status](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-) (beta)

### [Gets an SBOM test run result](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/sbom\_tests/-job\_id-/results) (beta)

## SastSettings

### [Enable/Disable the Snyk Code settings for an org](../reference/sastsettings.md#orgs-org\_id-settings-sast)

**More information:** [Enable Snyk Code](../../implement-snyk/enterprise-implementation-guide/phase-2-configure-account/set-visibility-and-configure-an-organization-template/enable-snyk-code.md) (Enterprise implementation guide, Phase 2)

### [Retrieves the SAST settings for an org](../reference/sastsettings.md#orgs-org\_id-settings-sast-1)

## ServiceAccounts

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md); [Choose a service account type to use with Snyk APIs](../../enterprise-setup/service-accounts/choose-a-service-account-type-to-use-with-snyk-apis.md)

### [Create a service account for an organization](../reference/serviceaccounts.md#orgs-org\_id-service\_accounts)

**More information:** [Service accounts using OAuth 2.0](../../enterprise-setup/service-accounts/service-accounts-using-oauth-2.0.md), [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Get a list of organization service accounts](../reference/serviceaccounts.md#orgs-org\_id-service\_accounts-1)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Update an organization service account](../reference/serviceaccounts.md#orgs-org\_id-service\_accounts-serviceaccount\_id)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Get an organization service account](../reference/serviceaccounts.md#orgs-org\_id-service\_accounts-serviceaccount\_id-1)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Delete a service account in an organization](../reference/serviceaccounts.md#orgs-org\_id-service\_accounts-serviceaccount\_id-2)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Manage an organization service account’s client secret](../reference/serviceaccounts.md#orgs-org\_id-service\_accounts-serviceaccount\_id-secrets)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Create a service account for a group](../reference/serviceaccounts.md#groups-group\_id-service\_accounts)

**More information:** [Service accounts using OAuth 2.0](../../enterprise-setup/service-accounts/service-accounts-using-oauth-2.0.md), [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Get a list of group service accounts](../reference/serviceaccounts.md#groups-group\_id-service\_accounts-1)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Update a group service account](../reference/serviceaccounts.md#groups-group\_id-service\_accounts-serviceaccount\_id)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Get a group service account](../reference/serviceaccounts.md#groups-group\_id-service\_accounts-serviceaccount\_id-1)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Delete a group service account](../reference/serviceaccounts.md#groups-group\_id-service\_accounts-serviceaccount\_id)

**More information:** [Manage service accounts using the Snyk API](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [Manage a group service account’s client secret](../reference/serviceaccounts.md#groups-group\_id-service\_accounts-serviceaccount\_id-secrets)

## SlackSettings

### [Create new Slack notification default settings](../reference/slacksettings.md#orgs-org\_id-slack\_app-bot\_id)

### [Get Slack integration default notification settings](../reference/slacksettings.md#orgs-org\_id-slack\_app-bot\_id-1)

### [Remove the given Slack App integration](../reference/slacksettings.md#orgs-org\_id-slack\_app-bot\_id-2)

### [Slack notification settings override for projects](../reference/slacksettings.md#orgs-org\_id-slack\_app-bot\_id-projects)

**More information:** [Slack apps](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (List all Slack notification customizations for a project)

### [Create a new Slack settings override for a given project](../reference/slacksettings.md#orgs-org\_id-slack\_app-bot\_id-projects-project\_id)

**More information:** [Slack apps](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (Create a Slack notification customization for a Project)

### [Update Slack notification settings for a project](../reference/slacksettings.md#orgs-org\_id-slack\_app-bot\_id-projects-project\_id-1)

**More information:** [Slack apps](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (Update a Slack notification customization for a Project)

### [Remove Slack settings override for a project](../reference/slacksettings.md#orgs-org\_id-slack\_app-bot\_id-projects-project\_id-2)

**More information:** [Slack apps](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (Delete a Slack notification customization for a Project)

## Slack

### [Get a list of Slack channels](../reference/slack.md#orgs-org\_id-slack\_app-tenant\_id-channels)

### [Get Slack Channel name by Slack Channel ID](../reference/slack.md#orgs-org\_id-slack\_app-tenant\_id-channels-channel\_id)

## Targets

### [Get targets by org ID](../reference/targets.md#orgs-org\_id-targets)&#x20;

**More information:** [Target definition on the Projects page](../../snyk-admin/snyk-projects/#target)

### [Get target by target ID](../reference/targets.md#orgs-org\_id-targets-target\_id)

### [Delete target by target ID](../reference/targets.md#orgs-org\_id-targets-target\_id-1)

## Test (v1)

**More information:** [Guidance for Java and Kotlin](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/guidance-for-java-and-kotlin.md)

### [Test package.json & yarn-lock file](../reference/test-v1.md#test-yarn)

### [Test sbt file](../reference/test-v1.md#test-sbt)

### [sbt\_Test for issues in a public package by group id, artifact id and version](../reference/test-v1.md#test-sbt-groupid-artifactid-version)

### [Test gemfile.lock file](../reference/test-v1.md#test-rubygems)

### [Test for issues in a public gem by name and version](../reference/test-v1.md#test-rubygems-gemname-version)

### [Test (pip) requirements.txt file](../reference/test-v1.md#test-pip)

### [Pip\_Test for issues in a public (pip) package by name and version](../reference/test-v1.md#test-pip-packagename-version)

### [Test package.json & package-lock.json file](../reference/test-v1.md#test-npm)

### [Test for issues in a public (npm) package by name and version](../reference/test-v1.md#test-npm-packagename-version)

**More information:** [Guidance for JavaScript and Node.js, Unmanaged JavaScript section](../../supported-languages-package-managers-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#unmanaged-javascript)

### [Test maven file](../reference/test-v1.md#test-maven)

### [Test  for issues in a (maven) public package by group id, artifact id and version](../reference/test-v1.md#test-maven-groupid-artifactid-version)

**More information:** [Guidance for Java and Kotlin](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/guidance-for-java-and-kotlin.md)

### [Test gradle file](../reference/test-v1.md#test-gradle)

### [Test for issues in a public (Gradle) package by group, name and version](../reference/test-v1.md#test-gradle-group-name-version)

### [Test vendor.json file](../reference/test-v1.md#test-govendor)

### [Test Gopkg.toml & Gopkg.lock File](../reference/test-v1.md#test-golangdep)

### [Test Dep Graph](../reference/test-v1.md#test-dep-graph)

**More information:** [Dep Graph API](../../supported-languages-package-managers-and-frameworks/bazel/dep-graph-api.md);[ Unmanaged JavaScript](../../supported-languages-package-managers-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#unmanaged-javascript) (Guidance for JavaScript and Node.js)

### [Test composer.json & composer.lock file](../reference/test-v1.md#test-composer)

## Users (v1)

### [Get user details](../reference/users-v1.md#user-userid)

### [Get My Details](../reference/users-v1.md#user-me)

### [Modify organization notification settings](../reference/users-v1.md#user-me-notification-settings-org-orgid)

### [Get organization notification settings](../reference/users-v1.md#user-me-notification-settings-org-orgid-1)

### [Modify project notification settings](../reference/users-v1.md#user-me-notification-settings-org-orgid-project-projectid)

### [Get project notification settings](../reference/users-v1.md#user-me-notification-settings-org-orgid-project-projectid-1)

## Users

### [My User Details](../reference/users.md)

### [Update a user’s role in a group](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#patch-/groups/-group\_id-/users/-id-) (beta)

Note: Use this endpoint to remove users from a group.

**More information:** [Remove members from Groups and Orgs using the API](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md)

### [Get user by ID](https://apidocs.snyk.io/?version=2024-07-10%7Ebeta#get-/orgs/-org\_id-/users/-id-) (beta)

## Webhooks (v1)

### [Create a webhook](../reference/webhooks-v1.md#org-orgid-webhooks)

### [List webhooks](../reference/webhooks-v1.md#org-orgid-webhooks-1)

### [Retrieve a webhook](../reference/webhooks-v1.md#org-orgid-webhooks-webhookid)

### [Delete a webhook](../reference/webhooks-v1.md#org-orgid-webhooks-webhookid-1)

### [Ping a webhook](../reference/webhooks-v1.md#org-orgid-webhooks-webhookid-ping)



\
