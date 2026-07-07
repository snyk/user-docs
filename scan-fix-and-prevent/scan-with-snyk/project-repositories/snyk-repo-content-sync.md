# Snyk Repo Content Sync

{% hint style="info" %}
**Release status**

Snyk Repo Content Sync is in Early Access and available only with Enterprise plans. To enable the feature, visit [Snyk Preview](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-hierarchy/snyk-preview).
{% endhint %}

Repo Content Sync provides native, automatic synchronization between the Org-level SCM Projects and Snyk. This feature ensures that your Snyk Projects accurately reflect the current security posture of your repositories without the need for manual re-imports.

## Supported environments

* Products: the import process supports all Project types, ecosystems, and product lines (Open Source, Container, IaC, and Code).
* SCMs: Snyk supports GitHub (all versions), GitLab, Azure Repos, and Bitbucket (all versions).
* Infrastructure: Snyk supports environments using Snyk Broker.

## Key capabilities

Repo Content Sync automatically manages your Projects based on changes in your repositories that have been scanned by Snyk:

* Automatic Project creation: Snyk automatically creates and monitors new Projects when you add new manifest, Docker, or configuration files to your scanned repos.
* Automatic deactivation: Snyk automatically deactivates Projects when you delete their associated manifest, Docker, or configuration files in the scanned repos.
* File renames and path changes: If you rename a file or change its path, Snyk creates a new Project for the new location and deactivates the old Project.

## How synchronization works

Push events trigger synchronization using webhooks. Snyk creates a webhook when you initially import a repository.

* Manifest, Docker, and IaC configuration files: adding, deleting, or renaming these files triggers an automatic update. You can view details of these actions in your Snyk import logs.
* Exclusions: Snyk respects folder and file exclusions as follows:
  * Snyk Open Source, Container, and IaC: Use the Exclude Folders field in the Organization level repository import window, or configure exclusions using Repo Monitor Configuration in the Inventory. Repo Content Sync will respect these settings on each subsequent sync.
  * Snyk Code: Use a `.snyk` file in the root of your repository to exclude specific directories or files from import. For details, visit [Exclude directories and files from Project import](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects/import-project-repository/exclude-directories-and-files-from-project-import).

## Considerations for Early Access

* Ignore history: For file renames, path changes, or .Net Framework upgrades, Snyk treats the change as a delete and create action. Snyk does not carry over the Project history and previous ignores to the new Project.
* Manual deactivations: When you enable this feature, previously deactivated Projects remain inactive. To reactivate a Project, navigate to the relevant Snyk Project and click **Activate**.
* PR checks: Snyk detects new Projects only when you merge them into the monitored branch. Snyk does not detect them during pull request checks.
