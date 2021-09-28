# Provision Azure services

### Background

The objective for these exercises is to demonstrate how Snyk secures your workloads in the context of [source control systems](https://docs.microsoft.com/en-us/azure/devops/user-guide/source-control?view=azure-devops). We will provide basic patterns intended for use in learning environments. For a deeper dive and learning more about Azure DevOps, we suggest referencing Microsoft's [getting started documentation](https://docs.microsoft.com/en-us/azure/devops/get-started/?view=azure-devops).

### Create your Azure DevOps Project

We will organize our development work using a project which will contain a repository. In the following steps, we will invoke the [`az devops project create`](https://docs.microsoft.com/en-us/cli/azure/ext/azure-devops/devops/project?view=azure-cli-latest#ext-azure-devops-az-devops-project-create) command to create our team project. Let's get started.

First, let's create the project by running the following command:

```bash
az devops project create --name MySnykProject --org https://dev.azure.com/myOrganizationName --source-control git --visibility private
```

{% hint style="info" %}
Note that we are passing a few additional parameters to set our `source control` to `git` and our repository to `private`.
{% endhint %}

Upon successful completion, you will see output similar to the following:

```javascript
{
  "abbreviation": null,
  "capabilities": {
    "processTemplate": {
      "templateName": "Basic",
      "templateTypeId": "<guid>"
    },
    "versioncontrol": {
      "gitEnabled": "True",
      "sourceControlType": "Git",
      "tfvcEnabled": "False"
    }
  },
  "defaultTeam": {
    "id": "<guid>",
    "name": "MySnykProject Team",
    "url": "https://dev.azure.com/myOrganizationName/_apis/projects/<guid>/teams/<guid>"
  },
  "defaultTeamImageUrl": null,
  "description": null,
  "id": "<guid>",
  "name": "MySnykProject",
  "revision": 21,
  "state": "wellFormed",
  "url": "https://dev.azure.com/myOrganizationName/_apis/projects/<guid>",
  "visibility": "private"
}
```

Alternatively, you can pass the `--open` parameter to the above CLI command to open the project in the default web browser which should look similar to the following figure:

![](../../../.gitbook/assets/azure_devops_02.png)

