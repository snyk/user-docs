# Data Considerations for Snyk AppRisk

## Repository Assets

Snyk AppRisk Essentials generates asset data from existing Snyk targets and projects. Repository assets can be generated from SCM and CLI targets from Snyk, depending on the Snyk product:

| Snyk Product               | Git integration (SCM) | Snyk CLI                                                        |
| -------------------------- | --------------------- | --------------------------------------------------------------- |
| **Open Source**            | **✔︎**                | Monitor - only if .git folder is detectable                     |
| **Container**              | dockerfiles only      | Monitor - for dockerfiles and only if .git folder is detectable |
| **Infrastructure as Code** | **✔︎**                | Report - for IaC and only if .git folder is detectable.         |
| **Code**                   | **✔︎**                | Not supported                                                   |

## Packages (first party)

Snyk AppRisk Essentials generates first party packages from Snyk Open Source data, and represents them on asset inventory pages. Currently, Snyk AppRisk Essentials does not support dependencies (third party packages).

## Data Freshness

Snyk handles data synchronization as follows:

|                  | Scheduled                                                                                                                                            | Manual                                                              |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| Policies         | Policies are generally executed within one hour of policy creation. You can use the Run button to manually run a policy.                             | Run policies in the policy editor.                                  |
| SCM Integrations | Snyk endeavors to refresh data from SCM integrations at least once within a 24 hour period.                                                          | Integrations can be refreshed manually using the Integrations page. |
| Snyk Products    | Data from the Snyk organizations in the associated Snyk Group for a Snyk AppRisk tenant are synced, generally at least once within a 24 hour period. | N/A                                                                 |

