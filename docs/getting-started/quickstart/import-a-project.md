# Import a Project

{% hint style="info" %}
**Recap**\
You have [created a Snyk account](create-or-log-in-to-a-snyk-account.md) and [integrated with your source code (Git) repository](set-up-an-integration.md) to allow access to your code for scanning. You can now run scans.
{% endhint %}

**Snyk Projects** are items that Snyk scans for issues, for example, a manifest file listing your open-source dependencies.

When you import a Project, Snyk scans that imported Project, and displays the results for you to review.

The following video shows how to import a Snyk Project:

{% embed url="https://thoughtindustries-1.wistia.com/medias/9hwr0bnvko" %}
Video demonstration of importing projects via the Snyk Web UI
{% endembed %}

The steps to import a Project are as follows:

* Select **Projects** > **Add Project,** and select where to import the Project from, for example, select **GitHub** to import from your GitHub repository, or select **CLI** to use the [Snyk CLI](../../snyk-cli/) locally:

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-10-20 at 15.23.55.png" alt="Add Project choices"><figcaption><p>Add Project choices</p></figcaption></figure>

</div>

* Select the repositories to use; then click **Add selected repositories** to import the selected repositories into your Project.\
  You can choose optional Settings for the Project you are importing: **Add custom file location** and **Exclude folders**, supported only for Snyk Open Source and Snyk Container. For details, see [Adding custom file locations and excluding folders](https://docs.snyk.io/integrations/git-repository-scm-integrations/snyk-azure-repositories-tfs-integration#adding-custom-file-locations-and-excluding-folders) in the documentation for Azure repositories integrations,

<figure><img src="../../.gitbook/assets/Screenshot 2023-10-20 at 15.20.49.png" alt="Select GitHub repositories to import"><figcaption><p>Select GitHub repositories to import</p></figcaption></figure>

{% hint style="info" %}
Projects you select to import are shown with a ![Check mark in check box](<../../.gitbook/assets/image (7) (2).png>).\
Projects previously imported are marked by a ✔.\
Private repositories are marked with a![Padlock](<../../.gitbook/assets/Screenshot 2023-05-11 at 23.05.30.png>).\
Forked repositories are marked with a ![Fork symbol](<../../.gitbook/assets/Screenshot 2023-05-11 at 23.15.46.png>)
{% endhint %}

## Project import settings

In **Settings**, optionally choose to:

* **Add custom file location** to add any additional dependencies from custom paths.
* **Exclude folders** to list up to ten folders to exclude from scanning during the import, for example, to shorten scanning time. Each folder name must not exceed 100 characters. This feature is supported for Snyk Open Source and Snyk Container.

## Import progress

A progress bar appears during import. Select **View last import log** to see log results.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-23 at 13.23.59.png" alt="Import Projects progress and option to view import log"><figcaption><p><strong>Import Projects progress and option to view import log</strong></p></figcaption></figure>

During the import, Snyk starts scanning the selected repos for relevant files, such as `package.json` files listing dependencies, in the entire directory tree and imports these files as Snyk Projects.

## Import results

Project import completes, with a status message:

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-23 at 13.24.35.png" alt="Project import success message"><figcaption><p>Project import success message</p></figcaption></figure>

You have now successfully imported and scanned the selected Project.

{% hint style="success" %}
If you see any errors during import, see [Project import errors](https://support.snyk.io/hc/en-us/articles/360001373118).
{% endhint %}

## Additional benefits of importing a Project

Importing a Project also does the following:

* Sets Snyk to run a regular scan on that Project for issues; see [Test frequency settings](../../snyk-admin/manage-settings/usage-settings.md#test-frequency-settings) for the defaults.
* Initiates some automation, especially default Snyk tests on pull and merge requests, which help prevent vulnerabilities from being added to the Project. This automation fails builds according to your conditions and can be disabled or customized in your [integration settings](../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/).

{% hint style="info" %}
For training on best practices in using automation, visit the Snyk course [Source Code Manager Configurations](https://learn.dev.snyk.io/lesson/configure-snyk-scm/).
{% endhint %}

## What's next?

You can now [view Snyk scan results](view-snyk-scan-results.md).
