# Snyk changed the "resolved" URL's in my Lock file

* [ Introduction to projects](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360019058297-Introduction-to-projects/README.md)
* [ View project information](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360011450838-View-project-information/README.md)
* [ View project issues, remediation and dependencies](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360016910877-View-project-issues-remediation-and-dependencies/README.md)
* [ View project settings](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360017002718-View-project-settings/README.md)
* [ View project history](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360016910977-View-project-history/README.md)
* [ Issue card information](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360018049037-Issue-card-information/README.md)
* [ Project Tags](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360013865038-Project-Tags/README.md)
* [ Project Attributes](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360012703537-Project-Attributes/README.md)

## Introduction to projects

Projects define the items Snyk scans.

A project includes:

* A scannable item external to Snyk.
* Configuration to define how to run that scan.

Projects appear on the **Projects** menu on the Snyk dashboard:

### Target

The address of the item to scan \(external to Snyk\), such as a Kubernetes cluster or a GitHub repo. One target may relate to many projects. The structure of the target depends on the origin.

### Origin

The project ecosystem, such as CLI, API, or Kubernetes.

### Targetfile

The specific item to scan in a target, such as a pom file in a GitHub repo.

[Snyk Code](https://support.snyk.io/hc/en-us/categories/360003257537-Snyk-Code) scans do not use targetfiles.

### Type

The scanning method to use for this project, such as static application security testing \(SAST, for scanning using Snyk Code\) or maven for a maven project using Snyk Open Source\). Part of the configuration for scanning.

