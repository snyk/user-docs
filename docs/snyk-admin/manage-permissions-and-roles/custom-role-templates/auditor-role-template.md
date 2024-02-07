# Auditor role template

This is a Group-level read-only role, meaning an Auditor can only view certain areas and functions in Snyk and cannot create PRs, Projects, and more.

This role can view issues, scan results, and reports. An Auditor often verifies that there is a scan snapshot for a particular resource or Snyk Project. The Auditor may be external to the company.

## Organization-level permissions

This table details the Organization-level permissions that apply to the **Auditor** custom role.

| Permission                             | Enabled? |
| -------------------------------------- | :------: |
| View Organization                      |     x    |
| Edit Organization                      |          |
| Remove Organization                    |          |
| View Organization Reports              |     x    |
| View Project                           |     x    |
| Add Project                            |          |
| Edit Project                           |          |
| Project Status                         |          |
| Test Project                           |          |
| Move Project                           |          |
| Remove Project                         |          |
| View Project History                   |     x    |
| Edit Project Integrations              |          |
| Edit Project Attributes                |          |
| View Jira Issues                       |     x    |
| Create Jira Issues                     |          |
| Edit Project Tags                      |          |
| View Project Ignores                   |     x    |
| Create Project Ignores                 |          |
| Edit Project Ignores                   |          |
| Remove Project Ignores                 |          |
| Create Pull Requests                   |          |
| Mark Pull Request checks as successful |          |
| View Collections                       |     x    |
| Create Collections                     |          |
| Edit Collections                       |          |
| Delete Collections                     |          |
| View Service Accounts                  |          |
| Create Service Accounts                |          |
| Edit Service Accounts                  |          |
| Remove Service Accounts                |          |
| View Users                             |          |
| Invite Users                           |          |
| Manage Users                           |          |
| Add Users                              |          |
| Provision Users                        |          |
| User Leave                             |          |
| User Remove                            |          |
| View Integrations                      |     x    |
| Edit Integrations                      |          |
| Test Packages                          |          |
| View Billing                           |          |
| Edit Billing                           |          |
| View Entitlements                      |          |
| View Preview Features                  |          |
| Edit Preview Features                  |          |
| View Audit Logs                        |     x    |
| View Outbound Webhooks                 |     x    |
| Create Outbound Webhooks               |          |
| Remove Outbound Webhooks               |          |
| View Apps                              |          |
| Install Apps                           |          |
| Create Apps                            |          |
| Edit Apps                              |          |
| Delete Apps                            |          |
| View Environments                      |          |
| Create Environments                    |          |
| Delete Environments                    |          |
| Update Environments                    |          |
| View Scans                             |     x    |
| Create Scans                           |          |
| View Resources                         |     x    |
| View Artifacts                         |     x    |
| Create Artifacts                       |          |
| View Custom Rules                      |          |
| Create Custom Rules                    |          |
| Edit Custom Rules                      |          |
| Delete Custom Rules                    |          |
| View Container Image                   |     x    |
| Create Container Image                 |          |
| Edit Container Image                   |          |
| Publish Kubernetes Resources           |          |

## Group-level permissions

This table details the Group-level permissions that apply to the **Auditor** custom role.

| Permission                       | Enabled? |
| -------------------------------- | :------: |
| View groups                      |     x    |
| Edit group details               |          |
| View group settings              |          |
| Edit settings                    |          |
| View group notification settings |          |
| Edit group notification settings |          |
| View orgs                        |     x    |
| Add orgs                         |          |
| Remove orgs                      |          |
| Read roles                       |          |
| Create roles                     |          |
| Edit roles                       |          |
| Remove roles                     |          |
| View users                       |     x    |
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
| View audit logs                  |     x    |
| View policies                    |     x    |
| Create policies                  |          |
| Edit policies                    |          |
| Delete policies                  |          |
| View reports                     |     x    |
| View tags                        |          |
| View IaC settings                |          |
| Edit IaC settings                |          |
| View feature flags               |          |
| Edit feature flags               |          |
| View request access settings     |          |
| Edit request access settings     |          |
| View SSO settings                |          |
| Edit SSO settings                |          |
| View Apps                        |          |
| Install Apps                     |          |
| Edit Apps                        |          |
| View AppRisk                     |     x    |
| Edit AppRisk                     |          |
| Access Insights                  |     x    |
