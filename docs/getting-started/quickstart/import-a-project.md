# Import a Project

{% hint style="info" %}
**Recap**\
You have [created a Snyk account](create-a-snyk-account.md) and [integrated with your source code (Git) repository](set-up-an-integration.md) to allow access to your code for scanning. You can now run scans.
{% endhint %}

**Snyk Projects** are items that Snyk scans for issues, for example, a manifest file listing your open-source dependencies.

When you import a Project, Snyk scans that imported Project, and displays the results for you to review.

The following short video shows how to import a Snyk Project:

{% embed url="https://thoughtindustries-1.wistia.com/medias/9hwr0bnvko" %}
Video demonstration of importing projects via the Snyk Web UI
{% endembed %}

Alternatively, follow these instructions:

* Select **Projects** > **Add Project,** and select where to import the Project from. For example, click GitHub to import from your GitHub repo, or click **CLI** to use the [Snyk CLI](../../snyk-cli/) locally:

<figure><img src="../../.gitbook/assets/Screenshot 2022-07-26 at 10.06.54.png" alt="Add Project choices"><figcaption><p>Add Project choices</p></figcaption></figure>

* Select the repositories to use, then click **Add selected repositories** to import the selected repositories into your project:

<figure><img src="../../.gitbook/assets/image (4).png" alt=""><figcaption><p>Add selected repositories</p></figcaption></figure>

{% hint style="info" %}
Projects you select to import are shown with a ![](<../../.gitbook/assets/image (7).png>).\
Projects previously imported are marked by a ✔.
{% endhint %}

#### Project import settings

In **Settings**, optionally choose to:

* **Add custom file location** to add any additional dependencies from custom paths.
* **Exclude folders** to list up to 10 folders to exclude from scanning during the import; for example, to shorten scanning time.

### Import progress

A progress bar appears during import. Click **View last import log** to see log results.

<figure><img src="../../.gitbook/assets/Screenshot 2022-07-26 at 10.23.09.png" alt="View log"><figcaption><p>View log</p></figcaption></figure>

During the import, Snyk starts scanning the selected repos for relevant files (such as  **package.json** files listing dependencies) in the entire directory tree, and imports these files as Snyk Projects.

### Import results

Project import completes, with a status message:

<figure><img src="../../.gitbook/assets/Screenshot 2022-06-13 at 11.38.00.png" alt="Project import success message"><figcaption><p>Project import success message</p></figcaption></figure>

You have now successfully imported and scanned the selected Project.

{% hint style="success" %}
If you see any errors during import, see [Project import errors](https://support.snyk.io/hc/en-us/articles/360001373118).
{% endhint %}

#### Additional benefits

Importing a Project also:

* Sets Snyk to run a regular scan on that Project for issues ([daily by default](../../snyk-admin/managing-settings/usage-page-details.md#projects)).
* Initiates some automations, especially default Snyk tests on pull/merge requests, which help prevent vulnerabilities from being added to the Project. This automation fails builds according to your conditions and can be disabled or customized in your [integration settings](../../integrations/git-repository-scm-integrations/).

{% hint style="info" %}
For training on best practices using automations, visit the Snyk Training course: [Source Code Manager Configurations](https://training.snyk.io/courses/source-code-manager-configurations).
{% endhint %}

## What's next?

You can now [view Snyk scan results](view-snyk-scan-results.md).
