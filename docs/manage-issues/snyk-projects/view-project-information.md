# View Project information

## View application vulnerabilities

To view your application vulnerabilities and licensing issues, go to [your Projects listing](https://app.snyk.io/projects), and open the Target application to list Projects:

<figure><img src="../../.gitbook/assets/Project-list.png" alt="Projects listed in the Target application snyk-goof"><figcaption><p>Projects listed in the Target application <strong>snyk-goof</strong></p></figcaption></figure>

Click on the Project entry to view detailed information for that Project:

<figure><img src="../../.gitbook/assets/project-header (1).png" alt="Detailed information for the Package.json Project"><figcaption><p>Detailed information for the <strong>Package.json</strong> Project</p></figcaption></figure>

The following information is available:

* **Header**: shows [project summary information](view-project-information.md#project-summary-information).
* **Issue cards**: show summaries of issues found. See [Issue card information](issue-card-information.md).
* **Views** (links at top right):
  * **Overview**: shows [Project issues, fixes, and dependencies](view-project-issues-fixes-and-dependencies.md).
  * **History**: shows historical snapshots of recent tests. See [View Project history](view-project-history.md).
  * **Settings**: shows [Project settings](view-project-settings.md).

### Project summary information

<figure><img src="../../.gitbook/assets/Project-header-new.png" alt="Header information for package.json project"><figcaption><p>Header information for <strong>package.json</strong> project</p></figcaption></figure>

The summary information shows:

* File and history details:
  * The name (plus link) of the monitored repository
  * The monitored branch name
  * A direct link to the Project file in the SCM
  * The time when the Project was first imported to Snyk
  * The time when an up-to-date snapshot of the file was fetched from the SCM and tested
* Project import information:
  * **Imported by**: The user who imported the project.
  * **Project owner**: Click **Add a project owner** to add an owner for this project, from a list of everyone who has access to that project (everyone who is a member of that Organization). This information can also be viewed in the [Projects endpoint](https://snyk.docs.apiary.io/#reference/projects), in the API.
* Predefined [Project attributes](project-attributes.md) and any additional [Project tag](project-tags.md) metadata.
