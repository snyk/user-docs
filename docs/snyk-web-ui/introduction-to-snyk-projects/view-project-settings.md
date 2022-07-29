# View project settings

Click **Settings** from your project, to view and edit project settings:

![](../../.gitbook/assets/screenshot_2021-04-14_at_09.23.38.png)

Click the **GitHub integration** section to edit SCM-specific settings \(applicable also to other SCMs\).

You can edit settings including:

* Notification settings
* Update test frequency
* Retrieve the unique project ID

## Deactivate a project

Deactivating a project will:

* Remove the webhook from the GitHub repository.
* Disable pull request tests for new vulnerabilities.
* Disable Fix pull request from being opened for newly disclosed vulnerabilities.
* Disable recurring tests - email alerts about newly disclosed vulnerabilities will be turned off

## Delete a project

Deleting a project will:

* Delete the project and all historical snapshot data from Snyk.
* Remove the webhook from the GitHub repository.

