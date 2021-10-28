---
description: The flags for each SCM script
---

# Flags



| **SCM/Flag**    | **Credentials**     | **Orgs**        | **Repo**    | **Exclusion File Path** | **Json** | **Skip Snyk monitored repos** | **Import file folder path** | **Repo type for import file** | **Additional flags**                            |
| --------------- | ------------------- | --------------- | ----------- | ----------------------- | -------- | ----------------------------- | --------------------------- | ----------------------------- | ----------------------------------------------- |
| **Azure**       | `"token"`           | `"projectKeys"` | `"repo"`    | `"exclusionFilePath"`   | `"json"` | `"skipSnykMonitoredRepos"`    | `"importConfDir"`           | `"importFileRepoType"`        | `"org" [required]`                              |
| **BB Cloud**    | `"user"/"paasword"` | `"workspaces"`  | `"repo"`    | `"exclusionFilePath"`   | `"json"` | `"skipSnykMonitoredRepos"`    | `"importConfDir"`           | `"importFileRepoType"`        |                                                 |
| **BB Server**   | `"token"`           | `"projectKeys"` | `"repo"`    | `"exclusionFilePath"`   | `"json"` | `"skipSnykMonitoredRepos"`    | `"importConfDir"`           | `"importFileRepoType"`        | `"url" [required]`                              |
| **Github**      | `"token"`           | `"orgs"`        | `"repo"`    | `"exclusionFilePath"`   | `"json"` | `"skipSnykMonitoredRepos"`    | `"importConfDir"`           | `"importFileRepoType"`        |                                                 |
| **Github Ent.** | `"token"`           | `"orgs"`        | `"repo"`    | `"exclusionFilePath"`   | `"json"` | `"skipSnykMonitoredRepos"`    | `"importConfDir"`           | `"importFileRepoType"`        | `"url" [required], "fetchAllOrgs" [optional]**` |
| **Gitlab**      | `"token"`           | `"groups"`      | `"project"` | `"exclusionFilePath"`   | `"json"` | `"skipSnykMonitoredRepos"`    | `"importConfDir"`           | `"importFileRepoType"`        | ``                                              |
| **Gitlab Ent.** | `"token"`           | `"groups"`      | `"project"` | `"exclusionFilePath"`   | `"json"` | `"skipSnykMonitoredRepos"`    | `"importConfDir"`           | `"importFileRepoType"`        | `"url" [required]`                              |

{% hint style="info" %}
The differences in the flags names are according to their naming at the SCM. \
The "private" SCMs, like BB server and Github enterprise and Gitlab enterprise, need a Host url to be set as a flag in the command for obvious reasons
{% endhint %}

{% hint style="info" %}
The `"fetchAllOrgs"` flag is unique fo Github Enterprise and distinguishes between two levels of access to the Orgs in the GHE:\
1\. With the flag - Will fetch all the orgs that the provided token has access to.\
2\. Without the flag - Will fetch only the orgs that the **\*User** has some operation rights to (e.g. read permissions etc).

**\*User** = The User that the provided token belongs to
{% endhint %}

{% hint style="info" %}
For further details on how to use the generated import file with the snyk-api-import tool, see [Creating and using the import files](creating-and-using-the-import-files.md).
{% endhint %}
