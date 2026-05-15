---
description: The flags for each SCM script
---

# Flags

| SCM             | Credentials         | Orgs            | Repo        | Exclusion File Path   | Json     | Skip Snyk monitored repos  | Import file folder path | Repo type for import file | Additional flags                                |
| --------------- | ------------------- | --------------- | ----------- | --------------------- | -------- | -------------------------- | ----------------------- | ------------------------- | ----------------------------------------------- |
| **Azure**       | `"token"`           | `"projectKeys"` | `"repo"`    | `"exclusionFilePath"` | `"json"` | `"skipSnykMonitoredRepos"` | `"importConfDir"`       | `"importFileRepoType"`    | `"org" [required]`                              |
| **BB Cloud**    | `"user"/"password"` | `"workspaces"`  | `"repo"`    | `"exclusionFilePath"` | `"json"` | `"skipSnykMonitoredRepos"` | `"importConfDir"`       | `"importFileRepoType"`    |                                                 |
| **BB Server**   | `"token"`           | `"projectKeys"` | `"repo"`    | `"exclusionFilePath"` | `"json"` | `"skipSnykMonitoredRepos"` | `"importConfDir"`       | `"importFileRepoType"`    | `"url" [required]`                              |
| **GitHub**      | `"token"`           | `"orgs"`        | `"repo"`    | `"exclusionFilePath"` | `"json"` |                            |                         |                           |                                                 |
| **GitHub Ent.** | `"token"`           | `"orgs"`        | `"repo"`    | `"exclusionFilePath"` | `"json"` |                            |                         |                           | `"url" [required], "fetchAllOrgs" [optional]**` |
| **GitLab**      | `"token"`           | `"groups"`      | `"project"` | `"exclusionFilePath"` | `"json"` |                            |                         |                           |                                                 |
| **GitLab Ent.** | `"token"`           | `"groups"`      | `"project"` | `"exclusionFilePath"` | `"json"` |                            |                         |                           | `"url" [required]`                              |

{% hint style="info" %}
The flag names correspond to their names in the SCM.\
The "private" SCMs, like Bitbucket server, GitHub Enterprise, and GitLab Enterprise, need a host url to be set as a flag in the command.
{% endhint %}

{% hint style="info" %}
The `"fetchAllOrgs"` flag is unique to GitHub Enterprise and distinguishes between two levels of access to the Orgs in the GHE:\
1\. With the flag - Fetches all the orgs that the provided token has access to.\
2\. Without the flag - Fetches only the orgs to which the **User** whose token was provided has some operation rights, for example, read permissions and so on.
{% endhint %}

{% hint style="info" %}
For further details on how to use the generated import file with the snyk-api-import tool, see [Creating and using the import files](creating-and-using-the-import-file.md).
{% endhint %}
