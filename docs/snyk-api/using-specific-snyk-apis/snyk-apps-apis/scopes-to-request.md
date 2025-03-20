# Scopes to request

Scopes define the permissions your Snyk App has to perform actions in a user’s account. When a user authorizes your Snyk App to access their Snyk account, they see the list of scopes the App is requesting and then decide whether or not they approve the connection.

When deciding which scopes your Snyk App will request, consider the actions your App will be performing. It may seem better to request every available scope, but users may refuse to install an App that asks for more permissions than needed. Also, a user installing your App will not be able to complete the authorization process if they do not have all the permissions matching the scopes the App is requesting.

The following lists the **available scopes**.

{% hint style="info" %}
`org.read` is a mandatory scope and should always be included.
{% endhint %}

| Scope                           | Description                                                                                     |
| ------------------------------- | ----------------------------------------------------------------------------------------------- |
| `org.read`                      | View Organization information and settings                                                      |
| `org.edit`                      | Edit Organization information and settings                                                      |
| `org.report.read`               | View reports in your Organization                                                               |
| `org.project.create`            | Add new Projects                                                                                |
| `org.project.read`              | View Project information and settings and view Organization targets                             |
| `org.project.edit`              | Edit Project information                                                                        |
| `org.project.delete`            | Permanently remove Projects and permanently remove Organization targets                         |
| `org.project.status`            | Activate and deactivate Projects                                                                |
| `org.project.test`              | Test Projects                                                                                   |
| `org.project.ignore.create`     | Create new Project ignores                                                                      |
| `org.project.ignore.read`       | View Project ignore information                                                                 |
| `org.project.ignore.edit`       | Configure Project ignores                                                                       |
| `org.project.ignore.delete`     | Permanently remove Project ignores                                                              |
| `org.project.attributes.edit`   | Apply and remove project attributes                                                             |
| `org.project.tag.edit`          | Create, apply and remove Project tags                                                           |
| `org.project.pr.create`         | Create fix pull requests for Projects                                                           |
| `org.project.pr.skip`           | Skip failed security tests on pull requests by marking checks as successful                     |
| `org.project.jira.issue.read`   | View Jira issue information                                                                     |
| `org.project.jira.issue.create` | Create new Jira issues                                                                          |
| `org.project.snapshot.read`     | View project dependencies, vulnerabilities, and other information obtained by scanning Projects |
| `org.package.test`              | Test packages in ecosystems supported by Snyk                                                   |
| `org.container_image.read`      | View container images                                                                           |
| `org.collection.create`         | Create a collection of Projects                                                                 |
| `org.collection.read`           | View Project collections                                                                        |
| `org.collection.edit`           | Add and remove Projects from collections                                                        |
| `org.collection.delete`         | Delete Project collections                                                                      |

{% hint style="info" %}
You cannot update scopes for a Snyk App after it has been created. If you change your mind about which scopes you need during the App development process, create a new Snyk App with a new list of scopes, and replace the `clientId` and `clientSecret` in the configuration of your App. If users have already installed the Snyk App, the users must authorize the new App with their Snyk account.
{% endhint %}
