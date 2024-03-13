# Request scopes

Scopes define the permissions your Snyk App has to perform actions in a user’s account. When a user authorizes your Snyk App to access their Snyk account, they see the list of scopes you are requesting and then decide whether or not they approve the connection.

When deciding which scopes your Snyk App will need, consider the actions your App will be performing. It may seem better to request every available scope, but users may refuse to install an App that asks for more permissions than required. Also, a user installing your App will not be able to complete the authorization process if they don’t have all the permissions matching the scopes you request.

The following lists the **available scopes**.

{% hint style="info" %}
`org.read` is a mandatory scope and should always be included.
{% endhint %}

| Scope                         | Description                                                                 |
| ----------------------------- | --------------------------------------------------------------------------- |
| org.read                      | View organization information and settings                                  |
| org.edit                      | Edit organization information and settings                                  |
| org.report.read               | View reports in your organization                                           |
| org.project.create            | Add new projects                                                            |
| org.project.read              | View project information and settings and view organization targets         |
| org.project.edit              | Edit project information                                                    |
| org.project.delete            | Permanently remove projects and permanently remove organization targets     |
| org.project.status            | Activate and deactivate projects                                            |
| org.project.test              | Test projects                                                               |
| org.project.ignore.create     | Create new project ignores                                                  |
| org.project.ignore.read       | View project ignore information                                             |
| org.project.ignore.edit       | Configure project ignores                                                   |
| org.project.ignore.delete     | Permanently remove project ignores                                          |
| org.project.tag.edit          | Create, apply and remove project tags                                       |
| org.project.pr.create         | Create fix pull requests for projects                                       |
| org.project.pr.skip           | Skip failed security tests on pull requests by marking checks as successful |
| org.project.jira.issue.read   | View Jira issue information                                                 |
| org.project.jira.issue.create | Create new Jira issues                                                      |
| org.package.test              | Test packages in our supported ecosystems                                   |

{% hint style="info" %}
You cannot currently update scopes for a Snyk App after it has been created. If you change your mind about which scopes you need during the App development process, create a new Snyk App with a new list of scopes, and replace the clientId and clientSecret in your App’s configuration. If users have installed the Snyk App already, the users will need to authorize the new App with their Snyk account.
{% endhint %}
