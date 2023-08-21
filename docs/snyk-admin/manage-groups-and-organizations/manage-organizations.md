# Manage Organizations

You can find instructions on this page to do the following:

* [Create an Organization](manage-organizations.md#create-an-organization)
* [Delete an Organization](manage-organizations.md#delete-an-organization)
* [Set your preferred Organization](manage-organizations.md#set-your-preferred-organization)

## Create an Organization

You can have an unlimited number of Organizations on Snyk. Each Organization can be on a different pricing plan.

{% hint style="info" %}
When you are satisfied with the settings and integrations for an Organization, you can copy its settings as you create additional Organizations.
{% endhint %}

**To create a new Organization:**

1\. On the Snyk Web UI, open the Organizations drop-down list from the top menu, and click the **Create new Organization** option:

<figure><img src="../../.gitbook/assets/snyk-org-switcher (1).png" alt="Create an Organization"><figcaption><p>Create an Organization</p></figcaption></figure>

2\. On the page to **create a new Organization**, enter a name for the new Organization. Consider using a structured naming convention to identify your Organizations.

{% hint style="info" %}
It is highly recommended to enter a unique name for the new Organization.
{% endhint %}

3\. From the drop-down list, select an Organization to use as a model for settings and integrations.

4\. Select **Create organization**.

<figure><img src="../../.gitbook/assets/2022-06-27_17-28-16.png" alt="Create an Organization"><figcaption><p>Create an Organization</p></figcaption></figure>

A new Organization is created and is added to your Organizations list.

By default, each Organization has a unique Snyk ID and internal name, which you can find in the **Settings** page of the Organization.

Review the settings and integrations for the new Organization and adjust them to meet your requirements before importing Projects to the new Organization.

## Delete an Organization

Organization Admins can delete Organizations when there are no Groups involved . If Groups are being used in the Organization, only a Group Admin can delete the Organization.

**To delete an Organization:**

1\. On the Snyk Web UI, open the Organization drop-down list on the top menu, and select the Organization you want to delete:

2\. After the selected Organization appears, click the **Org Settings** button in the left-hand Group menu:

3\. On the **Settings** page, select **Genera**l on the left menu.

4\. Scroll down to the **Delete organization** section, and click the **Delete organization** button:

<figure><img src="../../.gitbook/assets/Org Settings - Delete organization.png" alt="Delete Organization"><figcaption><p>Delete Organization</p></figcaption></figure>

4\. In the Confirmation dialog box, enter the name of the Organization you want to delete to confirm its deletion. Then click **OK**:

<figure><img src="../../.gitbook/assets/Org Settings - Delete organization - Confirmation.png" alt="Confirm delete Organization"><figcaption><p>Confirm delete Organization</p></figcaption></figure>

The selected Organization is deleted from your Snyk Account.

## Set your preferred Organization

If you have several Organizations, one of these Organizations is set by default as your **Preferred Organization** in your Snyk Account. A Preferred Organization determines the following:

* On the Snyk Web UI - which Organization will be displayed by default when you log in to your Snyk Account
* On the Snyk CLI - which Organization will be used by default for the test count when running tests via the CLI.\
  **Note**: To change the Organization used for the test count via the CLI, use the\
  `--org=<ORG_ID>` command. For more information, see the CLI help for the test commands: [Snyk test](../../snyk-cli/commands/test.md), [Snyk Code test](../../snyk-cli/commands/code-test.md), [Snyk Container test](../../snyk-cli/commands/container-test.md), or [Snyk IaC test](../../snyk-cli/commands/iac-test.md).

**To change your Preferred Organization:**

1\. On the Snyk Web UI, click your Account icon on the bottom left corner of the screen. Then select the **Account settings** option:

<figure><img src="../../.gitbook/assets/snyk-account-settings.png" alt="Account settings"><figcaption><p>Account settingd</p></figcaption></figure>

2\. On the **Account Settings** page – **Preferred Organization** section, open the Organization drop-down list, and select the Organization you want to set as your Preferred Organization:

{% hint style="info" %}
The Organization drop-down list displays your existing Organzations.
{% endhint %}

<figure><img src="../../.gitbook/assets/image (20) (1) (1).png" alt="Change preferred Organization"><figcaption><p>Change preferred Organization</p></figcaption></figure>

3\. Click the **Update Preferred Org** button to save your new setting.

The Organization you selected as your **Preferred Organization** is displayed when you log in to your Snyk Account, and used by default for the test count when you run tests using the CLI.
