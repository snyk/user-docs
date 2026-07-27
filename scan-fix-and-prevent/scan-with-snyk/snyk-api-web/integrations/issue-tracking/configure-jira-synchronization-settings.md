---
description: How to configure Jira synchronization settings for Snyk API and Web
nav_context: classic
---

# Configure Jira synchronization settings

You can connect Snyk API & Web with either Jira Cloud or your own Jira Server instance. This gives you two-way synchronization of your findings with Jira by integrating Snyk with your existing bug tracker or task manager.

Regardless of your Jira setup, you must choose which targets to synchronize and how. The available options are the same for both Jira Cloud and Jira Server.

For information about connecting Snyk with your Jira Server instance, visit [Integrate with Jira Server](integrate-with-jira-server.md).

## Access the Jira integration settings

To set up the configuration, access the **Integrations** tab of your **Target Settings** and locate the **Jira Server** or **Jira Cloud** module, depending on the integration you want to configure.

If the integration is available to set up, you see the following screen:

## Configure synchronization options

### Project

Choose which Jira project to sync with.

### Issue Type

Choose the type of issue for the findings.

The available options depend on the project you chose. The same applies to the status and priority mapping that follows.

### Automatically sync all findings

If enabled, Snyk automatically synchronizes all findings, existing and future, with Jira. If disabled, you must choose individual findings to sync in each finding's details.

### Status mapping

Maps Snyk API & Web status to Jira status:

* **Not Fixed**: The initial state of the finding, right after Snyk reports it.
* **Invalid** (optional): The user changes the finding manually to report a false positive.
* **Accepted Risk** (optional): The user changes the finding manually to acknowledge it but accept its risk and not fix it.
* **Fixed**: Snyk confirms that the finding is fixed.

You can use each Jira status only once for each project.

### Priority mapping

Maps Snyk API & Web severity to Jira priority:

* **Critical**: A finding that represents a critical risk for the application and requires immediate remediation.
* **High**: A finding that represents a high risk for the application, whose exploitation can relatively easily cause damage to the application.
* **Medium**: A finding that alone is unlikely to cause damage to the application. However, if combined with another one, or in rare situations by itself, it can cause damage.
* **Low**: Alone, a low-risk finding does not compromise your application, except in extreme situations. The attacker normally requires another higher-risk finding to take advantage of this one.

After you finish, click **Save** so that you do not lose your changes.
