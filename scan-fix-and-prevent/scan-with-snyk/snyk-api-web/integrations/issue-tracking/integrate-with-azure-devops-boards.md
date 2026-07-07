# Integrate with Azure DevOps Boards

By connecting Snyk API & Web to Microsoft Azure DevOps Boards, you can synchronize target scan results with an Azure Boards organization and project of your choice.

The synchronization is bi-directional. Snyk sends a reported finding to your Azure Boards, and as soon as the finding is set to **Done**, Snyk triggers a retest. If the finding is fixed, the Azure Boards work item remains closed. Otherwise, Snyk reopens it.

If Snyk detects a change, it updates the Azure Boards work item. For example, if the underlying vulnerability is fixed, Snyk detects it, sets the finding as fixed, and updates the Azure Boards work item to fixed as well.

Snyk synchronizes comments in both directions so that you always have all the information about a finding in both places.

You enable the Azure DevOps service connection at the account level and then on demand for each target. You set configurations such as which product and work item to use in the target settings.

## Authorize Snyk API & Web to access your Azure DevOps account

To start, navigate to the Snyk Integrations page at `https://plus.probely.app/integrations` and find the Azure DevOps section.

Snyk needs permission to connect to your Azure DevOps account to synchronize findings. Snyk asks for the minimum set of permissions that let it list organizations and projects and create work items from Snyk findings.

Click the link to begin. A browser tab opens at Azure, where you log in if you are not already authenticated. You then accept the permissions Snyk requests.

Click **Accept** to continue.

If you already have an organization and at least one project created at Azure DevOps, Azure redirects you to Snyk right away. Otherwise, Azure asks you to create an organization and a project, both required for this integration to work.

Choose the organization you want to use and click **Save**.

You might need to update your organization settings at Azure. Access your **Organization**, navigate to **Organization Settings**, and under **Security**, click **Policies**. Make sure that **Third-party application access via OAuth** is set to **On**.

## Choose your synchronization settings

Back at Snyk, choose which targets to synchronize and how. To configure a target to use Azure DevOps Boards, navigate to the **Integrations** tab from that target's **Settings** and locate the **Azure DevOps Boards** module.

You see the following screen:

### Project

Choose which project to sync with.

### Work Item Type

Choose which work item type to sync with. This depends on the selected project. Snyk supports only two types of work items: **Issue** or **Task**.

### Automatically sync all findings

If selected, Snyk syncs all findings that are not fixed to Azure, as well as any new findings.

Alternatively, you can enable per-finding synchronization. To do so, select **Sync finding** in the details of the finding, as shown here:

### Delete

Removes the configuration for this target. Azure keeps findings that are already synchronized.

To finish the configuration, select the project and work item and click **Save**. If the **Automatically sync all findings** box is selected, synchronization starts immediately and takes a few seconds.

Snyk adds two tags to each work item created at Azure:

* One indicating the severity, with the following values: `Low severity`, `Medium severity`, or `High severity`.
* One with `Probely`, identifying which work items are synced. It also gives you a way to filter those coming from Snyk. **Do not remove this tag**. Otherwise, the synchronization stops working for that work item.

## Integrate with Azure DevOps CI/CD pipelines

To foster automation between systems, integrate with Azure DevOps to run operations in Snyk triggered from Azure pipelines using the Snyk API.

The integration involves two steps:

1. Get the integration information from Snyk API & Web.
2. Configure Azure to integrate with Snyk API & Web.

### Get integration information from Snyk API & Web

Before configuring the integration in Azure, get the necessary information from Snyk:

1. Get the **Target Identifier** (Target ID):
   1. Navigate to the **Targets** list in the Snyk application.
   2. Click the target and obtain the target ID from the URL.
2. Generate the **API Key** and save it so Azure can perform actions in Snyk. Learn how to generate an API key in the Snyk documentation.

### Configure Azure to integrate with Snyk API & Web

With the information from Snyk, do the configuration in Azure:

1. Log in to the Azure DevOps account at `https://dev.azure.com` and navigate to **Pipelines**.
2. Click the pipeline to select it and then click **Edit** in the top-right corner of the screen.
3. In the list of Agent Jobs, click the plus (**+**) button to add a new task.
4. Select **Command Line** from the list and click **Add**.
5. Select the newly added Command Line Script.
6. Fill out the form with the command line configuration:
7. Set the **Display name** with the name of the script.
8. Set the **Script** field to:

```bash
curl -k -X POST 'https://api.probely.com/targets/<TARGET_ID>/scan_now/' -H 'Authorization: JWT <API_KEY>' -H 'Content-type: application/json' -d '{"scan_profile": "normal"}'
```

**Notes:**

1. In this example, the command triggers a target scan using the Snyk API endpoint. Explore the API for other operations to trigger from your Azure pipelines.
2. In the curl command:
   1. Replace `<TARGET_ID>` and `<API_KEY>` with the corresponding values obtained in the previous step.
   2. A parameter also defines the scan profile to use in the scan: `-d '{"scan_profile": "normal"}'`. You can remove it, and the scan profile is the one defined in the target settings.
3. Click the **Save & queue** dropdown menu and select **Save & queue** from the list.
4. A **Run pipeline** dialog appears. Click **Save and run** to run the pipeline manually and test the integration.

From now on, every time this Azure pipeline runs, it triggers the scan of the target in Snyk.
