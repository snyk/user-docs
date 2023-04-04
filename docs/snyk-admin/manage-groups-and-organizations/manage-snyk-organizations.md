# Manage Snyk Organizations

You can:

* [Create an Organization](manage-snyk-organizations.md#create-an-organization)
* [Delete an Organization](manage-snyk-organizations.md#delete-an-organization)
* [Set your preferred Organization](manage-snyk-organizations.md#set-your-preferred-organization)

### Create an Organization

You can have an unlimited number of Organizations on Snyk. Each organization can be on a different pricing plan.

{% hint style="info" %}
After you fine-tune the settings and integrations for an Organization, you can copy its settings as you create additional Organizations.
{% endhint %}

**To create a new organization:**

1\. On the Snyk Web UI, open the Organizations drop-down list from the top menu, and click the **Create a new organization** option:

<figure><img src="../../.gitbook/assets/snyk-org-switcher (1).png" alt="Create an Organization"><figcaption><p>Create an Organization</p></figcaption></figure>

2\. On the **Create a new organization** page, enter a name for the new Organization. Consider using a structured naming convention to identify your Organizations.

**Note**: It is highly recommended to enter a unique name for the new Organization.

3\. Select an Organization to use as a model for settings and integrations from the drop-down list.

4\. Select **Create organization**.

<figure><img src="../../.gitbook/assets/2022-06-27_17-28-16.png" alt="Create an Organization"><figcaption><p>Create an Organization</p></figcaption></figure>

A new Organization is created and is added to your Organizations list.

By default, each Organization has a unique Snyk ID and internal name, which you can find in the **Settings** page of the Organization.

You'll want to fine-tune the settings and integrations for the new Organization before importing Projects to the new Organization.

## Delete an Organization

Organization Admins can delete Organizations so long as there are no Groups involved (if Groups are being used in the organization, only a Group Admin can delete the Organization).

**To delete an organization:**

1\. On the Snyk Web UI, open the Organization drop-down list on the top menu, and select the organization you want to delete:

2\. After the selected organization appears, click the **Org Settings** button in the left-hand group menu:

3\. On the **Settings** page, select **Genera**l on the left menu:

4\. Scroll down to the **Delete organization** section, and click the **Delete organization** button:

<figure><img src="../../.gitbook/assets/Org Settings - Delete organization.png" alt="Delete Organization"><figcaption><p>Delete Organization</p></figcaption></figure>

4\. On the Confirmation dialog box, enter the name of the organization you want to delete to confirm its deletion. Then, click **OK**:

<figure><img src="../../.gitbook/assets/Org Settings - Delete organization - Confirmation (1).png" alt="Confirm delete Organization"><figcaption><p>Confirm delete Organization</p></figcaption></figure>

The selected organization is deleted from your Snyk Account.

## Set your preferred Organization

If you have several Organizations, one of these organizations is set by default as your **Preferred Organization** in your Snyk Account. A Preferred Organization determines the following:

* On the Snyk Web UI - which organization will be displayed by default when you log in to your Snyk Account.
* On the Snyk CLI - which organization will be used by default for the test count when running tests via the CLI.\
  **Note**: To change the organization used for the test count via the CLI, use the\
  `--org=<ORG_ID>` command. For more information, see [Options for the code test subcommand](https://docs.snyk.io/snyk-cli/commands/code).

**To change your Preferred Organization:**

1\. On the Snyk Web UI, click your Account icon on the bottom left corner of the screen. Then, select the **Account settings** option:

<figure><img src="../../.gitbook/assets/snyk-account-settings.png" alt="Account settings"><figcaption><p>Account settingd</p></figcaption></figure>

2\. On the **Account Settings** page – **Preferred Organization** section, open the Organization drop-down list, and select the organization you want to set as your Preferred Organization:

**Note**: The Organization drop-down list displays your existing organizations.

<figure><img src="../../.gitbook/assets/image (313).png" alt="Change preferred Organization"><figcaption><p>Change preferred Organization</p></figcaption></figure>

3\. Click the **Update Preferred Org** button to save your new setting.

The Organization you selected as your **Preferred Organization** will be displayed when you will log into your Snyk Account, and it will be used by default for the test count when running tests via the CLI.
