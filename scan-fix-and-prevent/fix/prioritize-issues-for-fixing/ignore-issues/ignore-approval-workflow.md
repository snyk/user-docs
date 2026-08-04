# Ignore Approval Workflow

Ignore Approval Workflow creates a layer of governance around risk acceptance within Snyk. You can perform the following:

* Request ignores for specific findings from developer workflows (CLI, IDE, and API).
* Approve or reject ignore requests before ignores are applied.
* Apply access controls and audit user actions.
* Notify users about new ignore requests pending approval and changes in request status.

{% hint style="info" %}
Ignore Approval Workflow supports Snyk Code. Snyk does not support consistent\
ignores for Snyk OS, Snyk Container, and Snyk IaC.
{% endhint %}

## Prerequisites <a href="#prerequisites" id="prerequisites"></a>

* ​[Enable Snyk Code Consistent Ignores](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code#enable-snyk-code-consistent-ignores) as the Ignore Approval Workflow does not process other types of ignores. Any existing ignores persist as approved ignores if the Ignore Approval Workflow is enabled.
* ​Disable DeepCode inline Code ignores. Existing inline Code does not persist when the Ignore Approval Workflow is enabled. DeepCode inline ignores, which are behind a feature flag and not publicly documented, need to be disabled.
* Ignore Approval Workflow is supported on the following versions of Snyk CLI and IDE:
  * Snyk CLI v1.299.0
  * Snyk IDE
    * Visual Studio Code v2.24.0
    * JetBrains v2.16.0
    * Eclipse v3.4.0
    * Visual Studio v2.4.0

## Enabling Ignore Approval Workflow

You can enable Ignore Approval Workflow at the Organization-level settings as follows:

1. Log in to the Snyk Web UI and select your [Group and Organization](https://docs.snyk.io/snyk-platform-administration/groups-and-organizations).
2. Navigate to **Organization** > **Settings**.
3. Enable **Ignore Approval Workflow for Snyk Code**.

Ensure Code Consistent Ignores is enabled for this Snyk Organization as well, since it is a prerequisite for Ignore Approval Workflow.

{% hint style="info" %}
Ignore Approval Workflow cannot be enabled at the Group level. An\
Organization-level API is available to enable it, so Groups with many\
Organizations can enable the feature at scale.
{% endhint %}

{% hint style="info" %}
When Ignore Approval Workflow is disabled, any pending ignore requests are canceled. Resubmit these ignores.
{% endhint %}

## Ignore Approval Workflow features <a href="#ignore-approval-workflow-features" id="ignore-approval-workflow-features"></a>

### Create and review ignore requests <a href="#create-and-review-ignore-requests" id="create-and-review-ignore-requests"></a>

Users can create and view ignore requests through the CLI, IDE, and API. Users with the reviewer permission can review and ignore requests through the UI on a new Ignore request page. Ignores are no longer automatically created, but instead a user can create an ignore with the status of `Pending` that can then be approved or rejected by reviewers. Requesters must provide a reason for their ignore request submission.

### Notifications <a href="#notifications" id="notifications"></a>

Reviewers get an email notification when someone creates a new request.\
Requesters get notified when a reviewer triages their request. Each of these\
has its own setting, and you can only enable or disable them through the API; neither is available on the account notifications page in the UI.

Snyk batches these emails and sends them every three hours if there is new\
activity. You cannot change this frequency. To receive a notification, you\
must be a member of the Organization.

#### Auto-approve ignores <a href="#auto-approve-ignores" id="auto-approve-ignores"></a>

To ensure AppSec teams have the option to move as quickly as they do today, there is a setting to allow users with the reviewer permissions to automatically approve ignores they themselves create, which bypasses the need for triage.

#### Audit <a href="#audit" id="audit"></a>

You can export all ignore request data for your Organization from the past 90 days\
through the audit logs API only, in JSON format. Most customers use this\
to automate exports into their reporting tools and dashboards. After 90 days,\
Snyk no longer makes the ignore request data available.

### Review Ignore Request in the Snyk UI

Reviewers get email notifications about new requests pending their approval. You can review the request in the **Ignore Requests** page in Snyk Web UI.

1. Log in to the Snyk Web UI and select your [Group and Organization](https://docs.snyk.io/snyk-platform-administration/groups-and-organizations).
2. Navigate to **Organization** > **Ignore requests**.

You can list all ignored requests in a **Pending**, **Approved**, or **Rejected** state on the **Ignore** **Requests** page only. Because you create ignore requests for findings, they do not appear on the **Issues** page until the associated finding becomes an issue.

3. Choose an Ignore request and then select **Manage** to view the details.
4. Investigate the potential risk and decide whether you need to approve or reject the request. You can change the ignore type and expiration prior to approving.
5. Add a comment to explain the decision. Requesters and reviewers cannot go\
   back and forth in comments; use this field to summarize any offline\
   discussion before you approve or reject.

{% hint style="info" %}
Snyk does not support assigning specific reviewers to ignore triage requests. Any user with reviewer permission can pick up any pending request
{% endhint %}

The developer who submitted the request receives an email notification with the decision. The request is no longer pending, and new tests will reflect the decision: either a finding that needs to be fixed if the request was rejected, or an ignored finding that is not shown by default.
