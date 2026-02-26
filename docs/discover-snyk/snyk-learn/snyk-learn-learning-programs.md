# Snyk Learn learning programs

{% hint style="info" %}
**Feature availability**

This feature is available in Early Access as part of the Snyk Learning Management add-on offering.

Send feedback to your Snyk account team, or email [support@snyk.io](mailto:support@snyk.io).
{% endhint %}

Learning programs allow Snyk administrators to curate specific paths of security education and Snyk product training, and assign them to groups of developers or Snyk users. By grouping lessons into programs and setting deadlines, organizations can automate security onboarding, meet compliance requirements, and drive targeted remediation.

## Prerequisites

To create and manage learning programs, you must must be a [Tenant Admin](../../snyk-platform-administration/groups-and-organizations/tenant/) or have a custom role with the `tenant.learning_program.edit` and `tenant.learning_program.read` permissions.

### Example: Creating a custom role

Use the following API call to create a role called Tenant Training Manager specifically for managing learning programs. Replace `{tenant_id}` with your tenant ID from the Snyk app (`https://app.snyk.io/tenant/{tenant_id}`) and `{snyk_api_token}` with your [API token](../../snyk-api/authentication-for-api/snyk-api-token-permissions-users-can-control.md). Also, update the Snyk API URL to your correct [regional API URL](../../developer-tools/snyk-ci-cd-integrations/azure-pipelines-integration/regional-api-endpoints.md).

```bash
curl --location --globoff --request POST "https://api.snyk.io/rest/tenants/{tenant_id}/roles?version=2024-10-15" \
  --header "Authorization: token {snyk_api_token}" \
  --header "Content-Type: application/vnd.api+json" \
  --data '{
    "data": {
      "type": "tenant_role",
      "attributes": {
        "name": "Tenant Training Manager",
        "description": "This role allows you to create and manage learning programs on Snyk Learn",
        "permissions": [
          "tenant.read",
          "tenant.feature.read",
          "tenant.group.list",
          "tenant.org.list",
          "tenant.pat.create",
          "tenant.membership.read",
          "tenant.user.read",
          "tenant.sso.read",
          "tenant.report.read",
          "tenant.billing.read",
          "tenant.roles.read",
          "tenant.support.case.create",
          "tenant.learning_program.edit",
          "tenant.learning_program.read"
        ]
      }
    }
  }'
```

## Creating a new learning program

{% hint style="info" %}
Programs cannot be edited once they have started. Ensure your lesson list and participant list are final before the start date of the program.

During Early Access, there is a limit to 300 participants per program.
{% endhint %}

Follow these steps to build and launch a new training initiative:

1. Navigate to [Learning Programs](https://learn.snyk.io/admin/management/).
2. Click **Create new learning program**. Enter a unique name (for example, `New Joiners 2026`) and a description.
3. Browse the available Snyk Learn catalog and select the modules you wish to include in the learning program.
4. Download the participants `.csv` template provided in the UI, add the email addresses of your learners, and upload the file.
5. Set the start date and duration for your learning program, and choose whether to send reminders to your users.
6. If needed, you can reset progress for lessons to have your users retake those lessons.
7. Click **Launch program**.

The program appears on the [Learning Progress](https://learn.snyk.io/user/learning-progress/) page for each user in the program, where they can see their progress and the modules required to complete the program.

## Cloning an existing learning program

You can clone an existing learning program to create a new learning program with the same content.

1. Navigate to [Learning Programs](https://learn.snyk.io/admin/management/).
2. Select the learning program you want to clone.
3. Click **Clone**.
4. Enter a unique name (for example, `New Joiners 2026`) and a description.
5. Edit the participants list.
6. Set the start date and duration for your learning program, and choose whether to send reminders to your users.
7. If needed, you can reset progress for lessons to have your users retake those lessons.
8. Click **Launch program**.

## Tracking progress

You can monitor the completion rates of your learning programs.

1. Navigate to [Learning Programs](https://learn.snyk.io/admin/management/).
2. Select the learning program you want to monitor.

The progress state for each user is displayed:

* **Not Started**: The user has been assigned but has not opened any modules.
* **In Progress**: The user has started at least one lesson.
* **Completed**: The user has finished all assigned modules in the program.

## Completing a learning program

Learning programs complete automatically once the scheduled end date is reached.&#x20;

If needed, you can click **Mark as Completed** on the program details page to end the program early.
