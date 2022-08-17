# Import a Project

{% hint style="info" %}
**Recap**\
You have [created a Snyk account](create-a-snyk-account.md), and [integrated with your source code (Git) repository](set-up-an-integration.md) to allow access to your code for scanning. You can now run the scans.
{% endhint %}

### **Introduction**

**Snyk Projects** are items that Snyk can scan for issues; for example, a manifest file listing your open source dependencies.

When you import a Project, Snyk scans that imported project and displays the results for you to review.

### Import a Project

{% embed url="https://thoughtindustries-1.wistia.com/medias/9hwr0bnvko" %}
Video demonstration of importing projects via the Snyk Web UI
{% endembed %}

To import a Projects to scan, from the Snyk Web UI:

* Select **Projects** >  **Add Project,** and select where to import the Project from (for example GitHub), or click **CLI** to use the [Snyk CLI](../snyk-cli/) tool locally:

![](<../.gitbook/assets/Screenshot 2022-07-26 at 10.06.54.png>)

* Select the repositories to use, then click **Add selected repositories** to import the selected repositories into your projects:

![](<../.gitbook/assets/Screenshot 2022-06-13 at 10.57.25.png>)

{% hint style="info" %}
Imported projects are indicated by a ✔ next to the repo name.
{% endhint %}

#### Project import settings

In **Settings**, optionally choose to:

* **Add custom file location** to add any additional dependencies from custom paths.
* **Exclude folders** to list up to 10 folders to exclude from scanning during the import; for example, to shorten scanning time.

#### Import progress

A progress bar appears during import: click **View last import log** to see log results.

![](<../.gitbook/assets/Screenshot 2022-07-26 at 10.23.09.png>)

During the import, Snyk starts scanning the selected repos for relevant files (for example, **package.json** files listing dependencies) in the entire directory tree, and imports these files as Snyk Projects.&#x20;

### Import results

Project import completes, with a status message:

![](<../.gitbook/assets/Screenshot 2022-06-13 at 11.38.00.png>)

You have now successfully imported and scanned the selected Project.

{% hint style="success" %}
If you see any errors during import, see [Project import errors](https://support.snyk.io/hc/en-us/articles/360001373118).
{% endhint %}

Importing a Project will also:

* Set Snyk to run a regular scan on that Project for issues ([daily by default](../features/user-and-group-management/managing-settings/usage-page-details.md#projects)).
* Initiate some automations, especially default Snyk tests on pull/merge requests, which help prevent vulnerabilities from being added to the Project. This automation fails builds according to your conditions and can be disabled or customized in your [integration settings](https://docs.snyk.io/integrations/git-repository-scm-integrations).

{% hint style="info" %}
For training on best practices using automations, visit the Snyk Training course: [Source Code Manager Configurations](https://training.snyk.io/courses/source-code-manager-configurations).
{% endhint %}

### What's next?

You can now [view Snyk scan results](view-snyk-scan-results.md).
