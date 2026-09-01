---
description: How Snyk Repo Content Sync keeps repository content current
nav_context: classic
---

# Snyk Repo Content Sync

{% hint style="info" %}
**Release status**

Snyk Repo Content Sync is generally available (GA) and enabled by default for Enterprise plans.
{% endhint %}

Repo Content Sync provides native, automatic synchronization between the Org-level SCM Projects and Snyk. This feature ensures that your Snyk Projects accurately reflect the current security posture of your repositories without the need for manual re-imports.

## Supported environments

* Products: the import process supports Snyk Code, Open Source, Secrets, IaC, and Container (Dockerfiles only).
* SCMs: Snyk supports GitHub (all versions), GitLab, Azure Repos, and Bitbucket (all versions).
* Infrastructure: Snyk supports environments using Snyk Broker.

## Key capabilities

Repo Content Sync automatically manages your Projects based on changes in your repositories that have been scanned by Snyk:

* Automatic Project creation: Snyk automatically creates and monitors new Projects when you add new manifest, Dockerfile, or configuration files to your scanned repos.
* Automatic deactivation: Snyk automatically deactivates Projects when you delete their associated manifest, Dockerfile, or configuration files in the scanned repos.
* File renames and path changes: If you rename a file or change its path, Snyk creates a new Project for the new location and deactivates the old Project.

## How synchronization works

Push events trigger synchronization using webhooks. Snyk creates a webhook when you initially import a repository.

* Manifest, Dockerfile, and IaC configuration files: adding, deleting, or renaming these files triggers an automatic update. You can view details of these actions in your Snyk import logs.
* Exclusions: Snyk respects folder and file exclusions as follows:
  * Snyk Open Source, Container, and IaC: Use the Exclude Folders field in the Organization level repository import window, or configure exclusions using Repo Monitor Configuration in the Inventory. Repo Content Sync respects these settings on each subsequent sync.
  * Snyk Code: Use a `.snyk` file in the root of your repository to exclude specific directories or files from import. For details, visit [Exclude directories and files from Project import](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects/import-project-repository/exclude-directories-and-files-from-project-import).
  * Snyk Secrets: Use a `.snyk` file in the root of your repository, with the global or secrets exclude sections. Snyk Secrets reads only the root file and does not apply `.snyk` files in subdirectories. For details, visit [Secrets scanning in the SCM integration](https://docs.snyk.io/developer-tools/integrations/scm-integrations/secrets-scanning-in-the-scm).

## Delete Projects and exclude them from future scans

When you delete one or more Projects, Snyk displays a confirmation dialog with an **Exclude these projects from future scans** check box, which is cleared by default. These test exclusions are separate from the import exclusions you configure in a `.snyk` file or the **Exclude folders** field.

* Select the check box: Snyk adds the Projects to the test exclusions for the repository and branch, and does not recreate them during future scans.
* Clear the check box: Snyk deletes the Projects but recreates them on the next scan if the underlying files still exist.

To reverse an exclusion, re-import the repository from the Organization or Group level. Re-importing clears the test exclusions and recreates the Projects, but each one is a new Project that does not carry over the issue history of the original.

{% hint style="warning" %}
A repository supports up to 100 test exclusions. If you select **Exclude these projects from future scans** when the list is already at the limit, Snyk deletes the Projects but does not add the exclusion, so the Projects reappear at the next sync.

* Bulk deletion through the API: the response includes a `meta.failed` entry with the reason `exclusion_limit_reached`.
* Deletion through the Web UI: Snyk recreates the Projects with no error message.

Snyk does not provide a view of the test exclusions or the deletion history for a repository.
{% endhint %}

## Considerations

* File renames and history: For file renames, path changes, or .NET Framework upgrades, Snyk treats the change as a delete and create action. Snyk does not carry over the Project history and previous ignores to the new Project.
* Manual deactivations: Snyk does not reactivate manually deactivated Projects during sync. To reactivate a Project, navigate to the relevant Snyk Project and click **Activate**.
* Organization-wide deletion: Deleting a Project removes it for the entire Organization, along with its history and findings.
* PR checks: Snyk detects new Projects only when you merge them into the monitored branch. Snyk does not detect them during pull request checks.
* Test batching: Snyk batches repeated changes to the same repository and branch into 10-minute windows instead of testing on every commit, so a change can take up to 10 minutes to appear.
* Exclusion limit: A repository supports up to 100 exclusions. When you select **Exclude these projects from future scans** and the exclusion list is already at the limit, Snyk deletes the Projects but does not add the exclusion, so the Projects reappear at the next repo sync.
  * Bulk deletion through the API: the response includes a `meta.failed` entry with the reason `exclusion_limit_reached`.
  * Deletion through the UI: Snyk does not surface an error, so the Projects reappear silently.
* Exclusion visibility: There is no in-product view of a repository's exclusions or deletion history.
