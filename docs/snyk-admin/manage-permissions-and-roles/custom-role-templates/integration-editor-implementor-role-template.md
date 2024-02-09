# Integration Editor/Implementor role template

This is a Group-level role with integration-related permissions to enable and process the integration of multiple third-party tools.

As an Integration Editor/Implementor, you should define whether this user implements the API. This template includes API implementation using the four Organization-level **Service Accounts** permissions. If this is not required for your custom role, do not include the **Service Accounts** permissions.

## Organization-level permissions

This table details the Organization-level permissions that apply to the **Integration Editor/Implementor** custom role.

| Permission                             | Enabled? |
| -------------------------------------- | :------: |
| View Organization                      |     x    |
| Edit Organization                      |          |
| Remove Organization                    |          |
| View Organization Reports              |          |
| View Project                           |     x    |
| Add Project                            |     x    |
| Edit Project                           |     x    |
| Project Status                         |     x    |
| Test Project                           |     x    |
| Move Project                           |     x    |
| Remove Project                         |     x    |
| View Project History                   |          |
| Edit Project Integrations              |     x    |
| Edit Project Attributes                |          |
| View Jira Issues                       |          |
| Create Jira Issues                     |          |
| Edit Project Tags                      |          |
| View Project Ignores                   |          |
| Create Project Ignores                 |          |
| Edit Project Ignores                   |          |
| Remove Project Ignores                 |          |
| Create Pull Requests                   |     x    |
| Mark Pull Request checks as successful |     x    |
| View Collections                       |          |
| Create Collections                     |          |
| Edit Collections                       |          |
| Delete Collections                     |          |
| View Service Accounts                  |     x    |
| Create Service Accounts                |     x    |
| Edit Service Accounts                  |     x    |
| Remove Service Accounts                |     x    |
| View Users                             |          |
| Invite Users                           |          |
| Manage Users                           |          |
| Add Users                              |          |
| Provision Users                        |          |
| User Leave                             |          |
| User Remove                            |          |
| View Integrations                      |     x    |
| Edit Integrations                      |     x    |
| Test Packages                          |          |
| View Billing                           |          |
| Edit Billing                           |          |
| View Entitlements                      |          |
| View Preview Features                  |          |
| Edit Preview Features                  |          |
| View Audit Logs                        |          |
| View Outbound Webhooks                 |     x    |
| Create Outbound Webhooks               |     x    |
| Remove Outbound Webhooks               |     x    |
| View Apps                              |     x    |
| Install Apps                           |     x    |
| Create Apps                            |     x    |
| Edit Apps                              |     x    |
| Delete Apps                            |     x    |
| View Environments                      |          |
| Create Environments                    |          |
| Delete Environments                    |          |
| Update Environments                    |          |
| View Scans                             |          |
| Create Scans                           |          |
| View Resources                         |          |
| View Artifacts                         |          |
| Create Artifacts                       |          |
| View Custom Rules                      |          |
| Create Custom Rules                    |          |
| Edit Custom Rules                      |          |
| Delete Custom Rules                    |          |
| View Container Image                   |          |
| Create Container Image                 |          |
| Edit Container Image                   |          |
| Publish Kubernetes Resources           |     x    |

## Group-level permissions

This table details the Group-level permissions that apply to the **Integration Editor/Implementor** custom role.

| Permission                       | Enabled? |
| -------------------------------- | :------: |
| View groups                      |          |
| Edit group details               |          |
| View group settings              |          |
| Edit settings                    |          |
| View group notification settings |          |
| Edit group notification settings |          |
| View orgs                        |          |
| Add orgs                         |          |
| Remove orgs                      |          |
| Read roles                       |          |
| Create roles                     |          |
| Edit roles                       |          |
| Remove roles                     |          |
| View users                       |          |
| Add users to the group           |          |
| Edit users in the group          |          |
| Remove users                     |          |
| Delete users                     |          |
| Provision users                  |          |
| Assign and unassign roles        |          |
| View service accounts            |          |
| Create service accounts          |          |
| Edit service accounts            |          |
| Remove service accounts          |          |
| View audit logs                  |          |
| View policies                    |          |
| Create policies                  |          |
| Edit policies                    |          |
| Delete policies                  |          |
| View reports                     |          |
| View tags                        |          |
| View IaC settings                |          |
| Edit IaC settings                |          |
| View feature flags               |          |
| Edit feature flags               |          |
| View request access settings     |          |
| Edit request access settings     |          |
| View SSO settings                |          |
| Edit SSO settings                |          |
| View Apps                        |     x    |
| Install Apps                     |     x    |
| Edit Apps                        |     x    |
| View AppRisk                     |          |
| Edit AppRisk                     |          |
| Access Insights                  |          |
