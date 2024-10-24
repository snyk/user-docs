# Scenarios for using the Snyk API

The Snyk API scenarios identify procedures you can use to accomplish tasks with Snyk applications using the API.

The scenarios are listed on this page under the endpoints they use. They are provided in a [repository](https://github.com/snyk-playground/cx-tools/tree/main/scripts) or on the user docs site. Links are included.

If you have issues when using these procedures, contact your Technical Success Manager or Solutions Engineer, or [submit a ticket](https://support.snyk.io/hc/en-us/requests/new) to Snyk support.

## API v1 Groups

**Task: Retrieve a Project snapshot for every Project in a given Group**

Scenario: [Retrieve-project-snapshots](https://github.com/snyk-playground/cx-tools/blob/main/scripts/retrieve-projects-snapshots.md) (complete procedure)

**Endpoints used:**\
API v1 [List all organizations in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)\
API v1 (Reporting API) [Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues)

## API v1 Import projects

**Task: Identify and import new repositories only**

Scenario: [Identify-and-import-new-repos](https://github.com/snyk-playground/cx-tools/blob/main/scripts/Identify-and-import-new-repos.md) (complete procedure)

**Endpoints used:**\
REST API [Get targets by org ID (use latest version)](https://apidocs.snyk.io/?version=2022-08-12%7Ebeta#tag--Targets)\
API v1 [Import targets](https://snyk.docs.apiary.io/#reference/import-projects/import-targets)

## API v1 Integrations

### Task: Rotate or change your broker token for any reason

Scenario: [Broker-token-rotation](https://github.com/snyk-playground/cx-tools/blob/main/scripts/broker-token-rotation.md) (complete procedure)

**Endpoints used:**\
API v1 [List all the organizations a user belongs to](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to) (group admin only)\
API v1 [Add new integration](https://snyk.docs.apiary.io/#reference/integrations/integrations/add-new-integration)\
API v1 [Update existing integration](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration) (to enable broker)

### Task: For a specific event or time, disable all interactions (pull requests, tests) from Snyk to the code base (source control management repository)

Scenario: [disable-all-interaction-from-snyk](https://github.com/snyk-playground/cx-tools/blob/main/scripts/disable-all-interaction-from-snyk.md) (complete procedure)

**Endpoints used alternative 1: Get integrations from different organizations and then update the settings for each integration**\
API v1 (Integrations) [List](https://snyk.docs.apiary.io/#reference/integrations/integrations/list)\
API v1 (Integration settings) [Update](https://snyk.docs.apiary.io/#reference/integrations/update)\
API v1 [Update existing integration](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration)

**Endpoints used alternative 2:** **Webhooks approach: remove the Snyk webhook by getting the Webhook Id and using it to delete the webhook**\
API v1 [List webhooks](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/list-webhooks)\
API v1 [Delete a webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook/delete-a-webhook)\
API v1 [Create a webhook](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/create-a-webhook)

## API v1 Issues (Reporting API)

### Task: Find all Projects affected by a vulnerability

Scenario: [find-all-projects-affected-by-a-vuln.md](https://github.com/snyk-playground/cx-tools/blob/main/scripts/find-all-projects-affected-by-a-vuln.md) (complete procedure)

**Endpoints used:**\
API v1 Issues (Reporting API) [Get list of issues](https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues)\
API v1 Issues (Reporting API) [List all organizations in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)\
API v1 Projects [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects)

### Task: Bulk ignore issues

Scenario: [bulk-ignore-issues](https://github.com/snyk-playground/cx-tools/blob/main/scripts/bulk-ignore-issues.md) (complete procedure)

**Endpoints used:**\
REST API beta [Get projects by org ID](https://apidocs.snyk.io/?version=2022-08-12%7Ebeta#get-/orgs/-org\_id-/projects)\
API v1 [Get list of latest issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues) (To get all issues but Code)\
REST experimental [Get a summary of issues within an org](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org\_id-/issues) (To get all Code issues)

## API v1 Organizations

### Task: Create multiple new Organizations that all have the same settings in a given Group

Scenario: [create-multiple-orgs-and-copy-settings](https://github.com/snyk-playground/cx-tools/blob/main/scripts/create-multiple-orgs-and-copy-settings.md) (complete procedure)

**Endpoints used:**\
API v1 [Create a new organization](https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization)\
API v1 [View organization settings](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/view-organization-settings)\
API v1 [Update organization settings](https://snyk.docs.apiary.io/#reference/organizations/organization-settings/update-organization-settings)\
API v1 [Clone an integration with settings and credentials](https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-switching/clone-an-integration-\(with-settings-and-credentials\))

### Task: Assign all users in a given list to all the Organizations a company has (all Organizations in a Group)

Scenario: [assign-users-to-all-orgs](https://github.com/snyk-playground/cx-tools/blob/main/scripts/assign-users-to-all-orgs.md) (complete procedure)

**Endpoints used:**\
API v1 [List all members in a group](https://snyk.docs.apiary.io/#reference/users/user-project-notification-settings/list-all-members-in-a-group)\
API v1 [Invite users](https://snyk.docs.apiary.io/#reference/groups/list-all-roles-in-a-group/invite-users)

### Task: Add users to organizations at scale ahead of the first login

Scenario: [Provision users to Orgs via API](../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md)

**Endpoint used:**\
API v1 [Provision a user to the organization](https://snyk.docs.apiary.io/#reference/organizations/provision-user/provision-a-user-to-the-organization)

## API v1 Projects

### Task: List all issues including Snyk Code issues in all the Projects in an Organization

Scenario: [list-all-issues-for-a-snyk-org](https://github.com/snyk-playground/cx-tools/blob/main/scripts/list-all-issues-for-a-snyk-org.md) (complete procedure)

**Endpoints used:**\
API v1 Projects [List all projects](https://snyk.docs.apiary.io/reference/projects/all-projects)\
API v1 Projects [List all aggregated issues](https://snyk.docs.apiary.io/reference/projects/aggregated-project-issues/list-all-aggregated-issues) (no Code)\
REST (experimental) [Get summary of issues within an org](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org\_id-/issues)\
REST (experimental) [Get a Snyk Code issue by its ID](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org\_id-/issues/detail/code/-issue\_id-)\
API v1 Projects (Ignored issues) [Retrieve ignore](https://snyk.docs.apiary.io/reference/projects/project-ignores-by-issue/retrieve-ignore) (scroll down)

### Task: Tag all Projects in Snyk

Scenario: [Tag projects in Snyk](https://github.com/snyk-playground/cx-tools/blob/main/scripts/tag-snyk-projects.md) (complete procedure)

**Endpoints used:**\
API v1 [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects) now REST API [List all Projects for an Org with the given Org ID](https://apidocs.snyk.io/?version=2023-06-19#get-/orgs/-org\_id-/projects)

### Task: Import fresh container images

Scenario: [import-new-container-images](https://github.com/snyk-playground/cx-tools/blob/main/scripts/import-new-container-images.md) (complete procedure)

**Endpoints used:**\
API v1 [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects)\
API v1 [Import targets](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets)\
API v1 [Get import job details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details)\
API v1 [Delete a project](https://snyk.docs.apiary.io/#reference/projects/individual-project/delete-a-project)

### Task: Detect and import new Projects in a repository into a target

Scenario: [detect-and-import-new-projects](https://github.com/snyk-playground/cx-tools/blob/main/scripts/detect-and-import-new-projects.md) (complete procedure)

**Endpoints used:**\
API v1 [List all projects](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects)\
API v1 [Import targets](https://snyk.docs.apiary.io/#reference/import-projects)

### Task: Move projects from one Organization to another

Scenario: [Move projects between organisations](https://github.com/snyk-playground/cx-tools/blob/main/scripts/move-projects.md) (complete procedure)

{% hint style="info" %}
THe API token used must have Group Admin access. \
If you are moving between Organizations in different Groups, you must use a personal API token with Group Admin permissions in both Groups. Service Accounts cannot move projects between Organizations in different Groups.&#x20;

Historical data for reporting will be lost.
{% endhint %}

**Endpoints used:**\
API v1 [Move Projects](https://snyk.docs.apiary.io/#reference/projects/move-project/move-project-to-a-different-organization)

## REST API Targets (Beta)

**Task: Detect new Projects (files) in repositories and import them into a Target in Snyk on a regular basis**

Scenario: [Identify-and-import-new-repos](https://github.com/snyk-playground/cx-tools/blob/main/scripts/Identify-and-import-new-repos.md) (complete procedure)

**Endpoint used:**\
REST API [Get targets by org ID](https://apidocs.snyk.io/?version=2022-07-08%7Ebeta#get-/orgs/-org\_id-/targets)\
API v1 [Import targets](https://snyk.docs.apiary.io/#reference/import-projects)
