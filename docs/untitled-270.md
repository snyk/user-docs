# View project settings

* [ Introduction to projects](/hc/en-us/articles/360019058297-Introduction-to-projects)
* [ View project information](/hc/en-us/articles/360011450838-View-project-information)
* [ View project issues, remediation and dependencies](/hc/en-us/articles/360016910877-View-project-issues-remediation-and-dependencies)
* [ View project settings](/hc/en-us/articles/360017002718-View-project-settings)
* [ View project history](/hc/en-us/articles/360016910977-View-project-history)
* [ Issue card information](/hc/en-us/articles/360018049037-Issue-card-information)
* [ Project Tags](/hc/en-us/articles/360013865038-Project-Tags)
* [ Project Attributes](/hc/en-us/articles/360012703537-Project-Attributes)

##  View project settings

Click **Settings** from your project, to view and edit project settings:

Click the **GitHub integration** section to edit SCM-specific settings \(applicable also to other SCMs\).

You can edit settings including:

* Notification settings
* Update test frequency
* Retrieve the unique project ID

### Deactivate a project

Deactivating a project will:

* Remove the webhook from the GitHub repository.
* Disable pull request tests for new vulnerabilities.
* Disable Fix pull request from being opened for newly disclosed vulnerabilities.
* Disable recurring tests - email alerts about newly disclosed vulnerabilities will be turned off

### Delete a project

Deleting a project will:

* Delete the project and all historical snapshot data from Snyk.
* Remove the webhook from the GitHub repository.

