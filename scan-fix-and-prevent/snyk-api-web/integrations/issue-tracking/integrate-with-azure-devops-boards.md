# Integrate with Azure DevOps Boards

By connecting Snyk API & Web to Microsoft Azure DevOps Boards, you can synchronize target scan results with an Azure Boards organization and project of your choice.

The synchronization is bi-directional, meaning that a finding reported by Snyk API & Web is sent to your Azure Boards, and as soon as it is set as Done, Snyk API & Web triggers a retest. If the finding is fixed, the Azure Boards work item remains closed. Otherwise, it is reopened.

If Snyk API & Web detects a change, it updates the Azure Boards work item. For instance, if the underlying vulnerability was fixed, Snyk API & Web detects it, sets the finding as fixed, and updates the Azure Boards work item to fixed as well.

Comments are also synchronized in both directions to ensure you always have all the information about that finding in both places.

The Azure DevOps service connection is enabled at the account level and then enabled on demand for each target. Configurations such as which product and work item to use are set in the target settings.

## Authorize Snyk API & Web to access your Azure DevOps account

To start, go to the Snyk API & Web Integrations page at `https://plus.probely.app/integrations` and find the Azure DevOps section.

Snyk API & Web needs permission to connect to your Azure DevOps account to synchronize findings. Snyk API & Web asks for the minimum set of permissions that allow it to list organizations and projects and create work items from Snyk API & Web findings.

<figure><img src="../../../../../.gitbook/assets/integrate-with-azure-devops-boards-authorization.png" alt="Azure DevOps integration section showing the authorization button"></figure>

Click the link to begin. A browser tab opens at Azure, where you are asked to log in if you are not already authenticated. Then you are asked to accept the permissions Snyk API & Web is requesting.

<figure><img src="../../../../../.gitbook/assets/integrate-with-azure-devops-boards-accept-permissions.png" alt="Azure DevOps permissions request dialog showing the permissions Snyk API & Web requires"></figure>

Click **Accept** to continue.

If you already have an organization and at least one project created at Azure DevOps, you are redirected to Snyk API & Web right away. If that is not the case, Azure asks you to create an organization and a project, both required for this integration to work.

Choose the organization you want to use and click **Save**.

<figure><img src="../../../../../.gitbook/assets/integrate-with-azure-devops-boards-select-organization.png" alt="Organization selection dropdown in Snyk API & Web"></figure>

You may need to update your organization settings at Azure. Access your **Organization**, go to **Organization Settings**, and under **Security**, click the **Policies** entry. There you need to make sure that the "Third-party application access via OAuth" is set to On.

<figure><img src="../../../../../.gitbook/assets/integrate-with-azure-devops-boards-oauth-policy.png" alt="Azure DevOps organization policies showing the Third-party application access via OAuth setting"></figure>

## Choose your synchronization settings

Back at Snyk API & Web, you need to choose which targets to synchronize and how. To configure a target to use Azure DevOps Boards, go to the **Integrations** tab from that target's **Settings** and locate the **Azure DevOps Boards** module.

You see the following screen:

<figure><img src="../../../../../.gitbook/assets/integrate-with-azure-devops-boards-settings.png" alt="Azure DevOps Boards target settings form showing project, work item type, and sync options"></figure>

### Project

Choose which project to sync with.

### Work Item Type

Choose which work item type to sync with. This depends on the project selected. Only two types of work items are currently supported: **Issue** or **Task**.

### Automatically sync all findings

If checked, all findings that are not fixed are synced to Azure, as well as any new findings.

Alternatively, you can enable per-finding synchronization. To do so, check **Sync finding** in the details of the finding, as shown here:

<figure><img src="../../../../../.gitbook/assets/integrate-with-azure-devops-boards-sync-finding.png" alt="Finding details page showing the Sync finding checkbox"></figure>

### Delete

Removes the configuration for this target. Findings already synchronized are kept at Azure.

To finish the configuration, select the project and work item and click **Save**. If the **Automatically sync all findings** box is checked, synchronization starts immediately and takes just a few seconds.

Snyk API & Web adds two tags to each work item created at Azure:

* One indicating the severity, with the following values: `Low severity`, `Medium severity`, or `High severity`.
* One with `Probely`, identifying which work items are being synced. It also gives you a way to easily filter those coming from Snyk API & Web. **Do not remove this tag**. Otherwise, the synchronization stops working for that work item.

<figure><img src="../../../../../.gitbook/assets/integrate-with-azure-devops-boards-tags.png" alt="Azure DevOps work item showing the Probely and severity tags"></figure>

## Integrate with Azure DevOps CI/CD pipelines

To foster automation between systems, integrate with Azure DevOps to execute operations in Snyk API & Web triggered from Azure pipelines using the Snyk API & Web API.

The integration involves two steps:

1. Get the integration information from Snyk API & Web.
1. Configure Azure to integrate with Snyk API & Web.

### Get integration information from Snyk API & Web

Before configuring the integration in Azure, get the necessary information from Snyk API & Web:

1. Get the **Target Identifier** (Target ID):
   1. Go to the **Targets** list in the Snyk API & Web application.
   1. Click the target and obtain the target ID from the URL.

<figure><img src="../../../../../.gitbook/assets/integrate-snyk-api-web-with-azure-devops-target-id.png" alt="Browser URL bar showing the target ID in the Snyk API & Web URL"></figure>

1. Generate the **API Key** and save it so Azure is able to perform actions in Snyk API & Web. Learn how to generate an API key in the Snyk API & Web documentation.

### Configure Azure to integrate with Snyk API & Web

With the information from Snyk API & Web, it is time to do the configuration in Azure:

1. Log in to the Azure DevOps account at `https://dev.azure.com` and go to **Pipelines**.

<figure><img src="../../../../../.gitbook/assets/integrate-snyk-api-web-with-azure-devops-pipelines.png" alt="Azure DevOps navigation showing the Pipelines menu item"></figure>

1. Click the pipeline to select it and then click the **Edit** button on the top-right corner of the screen.
1. In the list of Agent Jobs, click the plus (**+**) button to add a new task.

<figure><img src="../../../../../.gitbook/assets/integrate-snyk-api-web-with-azure-devops-add-task.png" alt="Azure DevOps agent jobs list showing the add task button"></figure>

1. Select **Command Line** from the list and click **Add**.

<figure><img src="../../../../../.gitbook/assets/integrate-snyk-api-web-with-azure-devops-command-line.png" alt="Azure DevOps task catalog showing the Command Line task"></figure>

1. Select the newly added Command Line Script.

<figure><img src="../../../../../.gitbook/assets/integrate-snyk-api-web-with-azure-devops-select-script.png" alt="Azure DevOps pipeline showing the newly added Command Line Script task"></figure>

1. Fill out the form with the command line configuration:

<figure><img src="../../../../../.gitbook/assets/integrate-snyk-api-web-with-azure-devops-configure-script.png" alt="Command Line Script configuration form showing display name and script fields"></figure>

1. Set the **Display name** with the name of the script.
1. Set the **Script** field to:

```bash
curl -k -X POST 'https://api.probely.com/targets/<TARGET_ID>/scan_now/' -H 'Authorization: JWT <API_KEY>' -H 'Content-type: application/json' -d '{"scan_profile": "normal"}'
```

**Notes:**

1. In this example, the command triggers a target scan using the Snyk API & Web API endpoint. Explore the API for other operations to trigger from your Azure pipelines.
1. In the curl command:
   1. Replace `<TARGET_ID>` and `<API_KEY>` with the corresponding values obtained in the previous step.
   1. There is also a parameter defining the scan profile to be used in the scan: `-d '{"scan_profile": "normal"}'`. You can remove it, and the scan profile will be the one defined in the target settings.

1. Click the **Save & queue** dropdown menu and select **Save & queue** from the list.

<figure><img src="../../../../../.gitbook/assets/integrate-snyk-api-web-with-azure-devops-save-queue.png" alt="Azure DevOps pipeline editor showing the Save & queue dropdown menu"></figure>

1. A **Run pipeline** dialog is displayed. Click **Save and run** to manually run the pipeline and test the integration.

From now on, every time this Azure pipeline runs, it triggers the scan of the target in Snyk API & Web.
