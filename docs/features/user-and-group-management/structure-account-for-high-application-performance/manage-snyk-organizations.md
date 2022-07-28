# Manage Snyk organizations

In the **Manage organization** section, you can:

* see how many private, public and inactive projects are in the organization.
* see and manage team members.
* leave this organization.
* delete this organization (administrator users only).

### Creating a new Snyk organization

You can have an unlimited number of organizations on Snyk. Each organization can be on a different pricing plan.

{% hint style="info" %}
Once you've fine-tuned the settings and integrations for an organization, you can copy its settings as you create additional organizations.
{% endhint %}

**To create a new organization:**

1\. On the Snyk Web UI, open the Organizations drop-down list from the top menu, and click the **Create a new organization** option:

![](<../../../.gitbook/assets/Org Settings - Creating New Organization option.png>)

2\. On the **Create a new organization** page, enter a name for the new organization. Consider using a structured naming convention to identify your organizations.

**Note**: It is highly recommended to enter a unique name for the new organization.

3\. Select an organization to use as a model for settings and integrations from the drop down list.&#x20;

4\. Select **Create organization**.&#x20;

![](../../../.gitbook/assets/2022-06-27\_17-28-16.png)

A new organization is created and is added to your Organizations list.

By default, each organization has a unique Snyk ID and internal name, which you can find in the **Settings** page of the organization.

You'll want to fine-tune the settings and integrations for the new organization as needed before importing projects to the new organization.

### Deleting an organization

Org Admins can delete organizations so long as there are no groups involved.

If groups are being used in the organization, only a Group Admin can delete the organization.

**To delete an organization:**

1\. On the Snyk Web UI, open the Organization drop-down list on the top menu, and select the organization you want to delete:

![](<../../../.gitbook/assets/Org Settings - Selecting an Org.png>)

2\. After the selected organization appears, click the **Org Settings** button<img src="../../../.gitbook/assets/Org Settings button - Icon (1) (1) (1) (1).png" alt="" data-size="line">on the top menu:

![](<../../../.gitbook/assets/Org Settings - Button.png>)

3\. On the **Settings** page, select **Genera**l on the left menu:

![](<../../../.gitbook/assets/Org Settings - General tab (1).png>)

3\. Scroll down to the **Delete organization** section, and click the **Delete organization** button:

![](<../../../.gitbook/assets/Org Settings - Delete organization.png>)

4\. On the Confirmation dialog box, enter the name of the organization you want to delete to confirm its deletion. Then, click **OK**:

![](<../../../.gitbook/assets/Org Settings - Delete organization - Confirmation (1).png>)

The selected organization is deleted from your Snyk Account.

### Setting your Preferred Organization

If you have several organizations, one of these organizations is set by default as your **Preferred Organization** in your Snyk Account. A Preferred Organization determines the following:

* On the Snyk Web UI - which organization will be displayed by default when you log in to your Snyk Account.
* On the Snyk CLI - which organization will be used by default for the test count when running tests via the CLI.\
  **Note**: To change the organization that will be used for the test count via the CLI, use the\
  `--org=<ORG_ID>` command. For more information, see [Options for the code test subcommand](https://docs.snyk.io/snyk-cli/commands/code).

You can change the Preferred Organization in your Snyk Account via your Account Settings on the Web UI.

**To change your Preferred Organization:**

1\. On the Snyk Web UI, click your Account icon on the top right corner of the screen. Then, select the **Account settings** option:

![](<../../../.gitbook/assets/Account Settings - Opening.png>)

2\. On the **Account Settings** page – **Preferred Organization** section, open the Organization drop-down list, and select the organization you want to set as your Preferred Organization:

**Note**: The Organization drop-down list displays your existing organizations.

![](<../../../.gitbook/assets/Account Settings - Preferred Org.png>)

3\. Click the **Update Preferred Org** button to save your new setting.

The organization you selected as your **Preferred Organization** will be displayed when you will log into your Snyk Account, and it will be used by default for the test count when running tests via the CLI.
