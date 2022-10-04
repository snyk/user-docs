# Scenarios for using Snyk API

The Snyk API scenarios identify procedures you can use to accomplish tasks with Snyk applications using the APIs.

The scenarios are listed on this page under the endpoints they use. They are provided in a [repository](https://github.com/snyk-playground/cx-tools/tree/main/scripts) or on the user docs site (links included).

If you have issues when using these procedures, contact your Technical Success Manager or Solutions Engineer, or [submit a ticket](https://support.snyk.io/hc/en-us/requests/new) to Snyk support.

## API v1 Groups

SEE also:\
API v1 Organizations: [Assign all users in a given list to all the organizations a company has (all organizations in a group)](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#assign-all-users-in-a-given-list-to-all-the-organizations-a-company-has-all-organization-in-a-group)

API v1 Issues (Reporting API): [Find all projects affected by a vulnerability](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#find-all-projects-affected-by-a-vulnerability)

### Retrieve a project snapshot for every project in a given group

Scenario: [Retrieve-project-snapshots](https://github.com/snyk-playground/cx-tools/blob/main/scripts/retrieve-projects-snapshots.md) (complete procedure)

**Endpoints used:**\
****API v1 [List all organizations in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)

API v1 (Reporting API) [Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues)

## API v1 Import projects

SEE also:\
API v1 Projects: [Import fresh Container images](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#import-fresh-container-images)

API v1 Projects: [Detect and import new projects in a repository into a target](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#detect-and-import-new-projects-in-a-repository-into-a-target)

API v1 REST API Targets (beta): [Detect new projects (files) in repositories and import them into a target in Snyk on a regular basis](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#detect-new-projects-files-in-repositories-and-import-them-into-a-target-in-snyk-on-a-regular-basis)

### Identify and import new repositories only

Scenario: [Identify-and-import-new-repos](https://github.com/snyk-playground/cx-tools/blob/main/scripts/Identify-and-import-new-repos.md) (complete procedure)

**Endpoints used:**\
****REST API [Get targets by org ID (use latest version)](https://apidocs.snyk.io/?version=2022-08-12%7Ebeta#tag--Targets)

API v1 [Import targets](https://snyk.docs.apiary.io/#reference/import-projects/import-targets)

## API v1 Integrations

SEE also:\
API v1 Organizations: [Create multiple new organizations that all have the same settings in a given group](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#create-multiple-new-organizations-that-all-have-the-same-settings-in-a-given-group)

API v1 Projects: [Detect and import new projects in a repository into a target](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#detect-and-import-new-projects-in-a-repository-into-a-target)

### Rotate or change your broker token for any reason

Scenario: [Broker-token-rotation](https://github.com/snyk-playground/cx-tools/blob/main/scripts/broker-token-rotation.md) (complete procedure)

**Endpoints used:**\
****API v1 [List all the organizations a user belongs to](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to) (group admin only)

API v1 [Add new integration](https://snyk.docs.apiary.io/#reference/integrations/integrations/add-new-integration)

API v1 [Update existing integration](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration) (to enable broker)

### For a specific event or time, disable all interactions (pull requests, tests) from Snyk to the code base (source control management repository)

Scenario: [disable-all-interaction-from-snyk](https://github.com/snyk-playground/cx-tools/blob/main/scripts/disable-all-interaction-from-snyk.md) (complete procedure)

**Endpoints used alternative 1: Get integrations from different organizations and then update the settings for each integration**\
****API v1 (Integrations) [List](https://snyk.docs.apiary.io/#reference/integrations/integrations/list)

API v1 (Integration settings) [Update](https://snyk.docs.apiary.io/#reference/integrations/update)

API v1 [Update existing integration](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration)

**Endpoints used alternative 2:** **Webhooks approach: remove the Snyk webhook by getting the Webhook Id and using it to delete the webhook**\
****API v1 [List webhooks](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/list-webhooks)

API v1 [Delete a webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook/delete-a-webhook)

API v1 [Create a webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/create-a-webhook)

## API v1 Issues (Reporting API)

SEE also:\
API v1 Groups: [Retrieve a project snapshot for every project in a given group](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#retrieve-a-project-snapshot-for-every-project-in-a-given-group)

### Find all projects affected by a vulnerability

Scenario: [find-all-projects-affected-by-a-vuln.md](https://github.com/snyk-playground/cx-tools/blob/main/scripts/find-all-projects-affected-by-a-vuln.md) (complete procedure)

**Endpoints used:**\
****API v1 Issues (Reporting API) [Get list of issues](https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues)

API v1 Issues (Reporting API) [List all organizations in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)

API v1 Projects [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects)

### Bulk ignore issues

Scenario: [bulk-ignore-issues](https://github.com/snyk-playground/cx-tools/blob/main/scripts/bulk-ignore-issues.md) (complete procedure)

**Endpoints used:**\
****REST API beta [Get projects by org ID](https://apidocs.snyk.io/?version=2022-08-12%7Ebeta#get-/orgs/-org\_id-/projects)

API v1 [Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues) (To get all issues but Code)

REST experimental [Get a summary of issues within an org](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org\_id-/issues) (To get all Code issues)

### REST API Issues (experimental)

SEE:\
API v1 Issues: [Bulk ignore issues](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#bulk-ignore-issues)

API v1 Projects: [List all issues including Snyk Code issues in all the projects in a group](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#list-all-issues-including-snyk-code-issues-in-all-the-projects-in-an-organization)

### API v1 Organizations

SEE also:\
API v1 Integrations: [Rotate or change your broker token for any reason](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#rotate-or-change-your-broker-token-for-any-reason)

### Create multiple new organizations that all have the same settings in a given group

Scenario: [create-multiple-orgs-and-copy-settings](https://github.com/snyk-playground/cx-tools/blob/main/scripts/create-multiple-orgs-and-copy-settings.md) (complete procedure)

**Endpoints used:**\
****API v1 [Create a new organization](https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization)

API v1 [View organization settings](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/view-organization-settings)

API v1 [Update organization settings](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-organization-settings)

API v1 [Clone an integration with settings and credentials](https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-switching/clone-an-integration-\(with-settings-and-credentials\))

### Assign all users in a given list to all the organizations a company has (all organization in a group)

Scenario: [assign-users-to-all-orgs](https://github.com/snyk-playground/cx-tools/blob/main/scripts/assign-users-to-all-orgs.md) (complete procedure)

**Endpoints used:**\
****API v1 [List all members in a group](https://snyk.docs.apiary.io/#reference/users/user-project-notification-settings/list-all-members-in-a-group)

API v1 [Invite users](https://snyk.docs.apiary.io/#reference/groups/list-all-roles-in-a-group/invite-users)

### Add users to organizations at scale ahead of the first login

Scenario: [Provision users to Orgs via API](../features/user-and-group-management/managing-users-and-permissions/provision-users-to-orgs-via-api.md)

**Endpoint used:**\
API v1 [Provision a user to the organization](https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization)

## API v1 Projects

SEE also:\
API v1 Issues (Reporting API): [Find all projects affected by a vulnerability](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#find-all-projects-affected-by-a-vulnerability)

API v1 Issues (Reporting API): [Bulk ignore issues](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#bulk-ignore-issues)

### List all issues including Snyk Code issues in all the projects in an organization

Scenario: [list-all-issues-for-a-snyk-org](https://github.com/snyk-playground/cx-tools/blob/main/scripts/list-all-issues-for-a-snyk-org.md) (complete procedure)

**Endpoints used:**\
****API v1 Projects [List all projects](https://snyk.docs.apiary.io/reference/projects/all-projects)

API v1 Projects [List all aggregated issues](https://snyk.docs.apiary.io/reference/projects/aggregated-project-issues/list-all-aggregated-issues) (no Code)

REST (experimental) [Get summary of issues within an org](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org\_id-/issues)

REST (experimental) [Get a Snyk Code issue by its ID](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org\_id-/issues/detail/code/-issue\_id-)

API v1 Projects (Ignored issues) [Retrieve ignore](https://snyk.docs.apiary.io/reference/projects/project-ignores-by-issue/retrieve-ignore) (scroll down)

### Tag all projects in Snyk

Scenario: [Tag projects in Snyk](https://github.com/snyk-playground/cx-tools/blob/main/scripts/tag-snyk-projects.md) (complete procedure)

**Endpoints used:**\
****API v1 [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects)

### Import fresh Container images

Scenario: [import-new-container-images](https://github.com/snyk-playground/cx-tools/blob/main/scripts/import-new-container-images.md) (complete procedure)

**Endpoints used:**\
****API v1 [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects)

API v1 [Import targets](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets)

API v1 [Get import job details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details)

API v1 [Delete a project](https://snyk.docs.apiary.io/#reference/projects/individual-project/delete-a-project)

### Detect and import new projects in a repository into a target

Scenario: [detect-and-import-new-projects](https://github.com/snyk-playground/cx-tools/blob/main/scripts/detect-and-import-new-projects.md) (complete procedure)

**Endpoints used:**\
****API v1 [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects)

API v1 [Import targets](https://snyk.docs.apiary.io/#reference/import-projects)

## REST API Targets (Beta)

### Detect new projects (files) in repositories and import them into a target in Snyk on a regular basis

Scenario: [Identify-and-import-new-repos](https://github.com/snyk-playground/cx-tools/blob/main/scripts/Identify-and-import-new-repos.md) (complete procedure)

**Endpoint used:**\
****REST API [Get targets by org ID](https://apidocs.snyk.io/?version=2022-07-08%7Ebeta#get-/orgs/-org\_id-/targets)

API v1 [Import targets](https://snyk.docs.apiary.io/#reference/import-projects)

### API v1 Webhooks

SEE:\
API v1 Integrations: [Rotate or change your broker token for any reason](https://docs.snyk.io/snyk-api-info/scenarios-for-using-snyk-api#rotate-or-change-your-broker-token-for-any-reason)
