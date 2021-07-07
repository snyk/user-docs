# My project appears as  /file://nothing in Snyk

##  My project appears as /file://nothing in Snyk

## Overview

Sometimes a project appears in Snyk with an odd or unfamiliar Project name, commonly seen as /file://nothing

## Cause

When importing a project, Snyk will look inside the .git/config file for the remote origin URL and use that as the Project name

We have seen instances where a 3rd party Continuous Integration tool is adding the following to the project's .git/config:

```text
[remote "origin"]
        url = file://nothing
```

Resolution

There are two ways to solve this issue:

1\) Update the .git/config file to remove or update the remote origin URL

2\) If using the CLI monitor command, use the `--remote-repo-url` argument to override the Project's group name.

See this article for more details: [Snyk CLI monitored projects are created with IDs in the project name](/hc/en-us/articles/360000910677)

