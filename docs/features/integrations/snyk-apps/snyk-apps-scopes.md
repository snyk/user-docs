# Snyk Apps Scopes

The following is a list of scopes that are currently in active development and as such not available publicly. Please check the **Requesting Scopes** section on the [Create an App](getting-started-with-snyk-apps/create-an-app-via-the-api.md#generally-available) page for generally available scopes.

{% hint style="danger" %}
The following list of scopes is available behind a feature flag (`appsGranularPermissions`). Being behind a feature flag means they are still in active development. If you would like to test them out as they give a look into how Snyk Apps scopes will look in the future, contact our [support team](https://snyk.io/contact-us/).
{% endhint %}

Scopes behind a feature flag (**in active development**):

| Scope                             | Description                                    |
| --------------------------------- | ---------------------------------------------- |
| org.read                          | List organizations                             |
| org.test                          | Test Snyk organization test                    |
| org.settings.read                 | Read Snyk organization's settings              |
| org.sast.settings.read            | Read Snyk organization's SAS settings          |
| org.notification\_settings.read   | Read Snyk organization's notification settings |
| org.report.read                   | Read Snyk organization's report                |
| org.project.add                   | Add projects under Snyk organization           |
| org.project.read                  | Read Snyk organization's projects              |
| org.project.edit                  | Edit Snyk organization's project               |
| org.project.remove                | Remove Snyk organization's project             |
| org.project.activate              | Activate Snyk organization's project           |
| org.project.deactivate            | Deactivate Snyk organization's project         |
| org.project.test                  | Test Snyk organization's project               |
| org.project.monitor               | Monitor project under Snyk organization        |
| org.project.import.read           | Read project imported for Snyk organization    |
| org.project.ignore.create         | Create project ignores                         |
| org.project.ignore.read           | Read project ignores                           |
| org.project.ignore.edit           | Edit project ignores                           |
| org.project.ignore.delete         | Delete project ignores                         |
| org.project.settings.read         | Read project settings                          |
| org.project.settings.edit         | Edit project settings                          |
| org.project.tag.apply             | Apply project tags                             |
| org.project.tag.remove            | Remove project tags                            |
| org.project.pr.create             | Create PR                                      |
| org.project.jira.issue.read       | Read project JIRA issue                        |
| org.project.jira.issue.create     | Create a project JIRA issue                    |
| org.package.test                  | Test package                                   |
| org.target.read                   | Read Snyk organization target                  |
| org.target.delete                 | Delete Organization target                     |
| org.integration.read              | Read Snyk organization's integrations          |
| org.user.read                     | Read Snyk organization's users                 |
| org.user.leave                    | Read Snyk organization user leave              |
| org.entitlements.read             | Read Snyk organization entitlements            |
| org.audit\_log.read               | Read Snyk organization audit logs              |
| org.flags.read                    | Read Snyk organization's feature flags         |
| org.settings.request\_access.read | Read request access settings                   |
