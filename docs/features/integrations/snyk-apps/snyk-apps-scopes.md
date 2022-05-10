# Snyk Apps Scopes

The following is a list of scopes that are currently in active development and as such _not available_ publicly. See [Requesting scopes](getting-started-with-snyk-apps/create-an-app-via-the-api.md#requesting-scopes) for generally available scopes.

{% hint style="danger" %}
The following list of scopes is available behind a feature flag (`appsGranularPermissions`). Being behind a feature flag means they are still in active development. If you would like to test them out and see how Snyk Apps scopes will look in the future, contact Snyk [support](https://snyk.io/contact-us/).
{% endhint %}

Scopes behind a feature flag (**in active development**):

| Scope                         | Description                                |
| ----------------------------- | ------------------------------------------ |
| org.read                      | View organization information and settings |
| org.report.read               | View reports in your organization          |
| org.project.read              | View project information and settings      |
| org.project.add               | Add new projects                           |
| org.project.edit              | Edit project information                   |
| org.project.remove            | Permanently remove projects                |
| org.project.status            | Activate and deactivate projects           |
| org.project.test              | Test projects                              |
| org.project.ignore.read       | View project ignore information            |
| org.project.ignore.create     | Create new project ignores                 |
| org.project.ignore.edit       | Configure project ignores                  |
| org.project.ignore.delete     | Permanently remove project ignores         |
| org.project.tag.edit          | Create, apply and remove project tags      |
| org.project.pr.create         | Create fix pull requests for projects      |
| org.project.jira.issue.read   | View Jira issue information                |
| org.project.jira.issue.create | Create new Jira issues                     |
| org.package.test              | Test packages in our supported ecosystems  |
| org.target.read               | View organization targets                  |
| org.target.delete             | Permanently remove organization targets    |
| org.integration.read          | View organization integration information  |
| org.user.read                 | View organization member information       |
| org.user.leave                | Leave the organization                     |
| org.entitlements.read         | View organization entitlements             |
| org.audit\_log.read           | View audit logs for the organization       |
| org.flags.read                | View upcoming and experimental features    |
