# Configure Jira synchronization settings

You can connect Snyk API & Web with either Jira Cloud or with your own Jira Server instance. This enables you to have two-way synchronization of your findings with Jira, by fully integrating Snyk API & Web with your existing bug tracker or task manager.

Regardless of your Jira setup, you need to choose which targets to synchronize and how. The options available are the same for both Jira Cloud and Jira Server.

For information on how to connect Snyk API & Web with your Jira Server instance, visit [Integrate with Jira Server](integrate-with-jira-server.md).

## Access the Jira integration settings

To set up the configuration, access the **Integrations** tab of your **Target Settings** and locate the **Jira Server** or **Jira Cloud** module, depending on the integration you want to configure.

If the integration is available to be set up, you see the following screen:

<figure><img src="../../../../../.gitbook/assets/jira-synchronization-settings.png" alt="Jira synchronization settings configuration form showing project, issue type, sync options, status mapping, and priority mapping fields"></figure>

## Configure synchronization options

### Project

Choose which Jira project to sync with.

### Issue Type

Choose the type of issue for the findings.

The options available depend on the project you chose before. The same goes for the status and priority mapping that follows.

### Automatically sync all findings

If enabled, all findings, existing and future, are automatically synchronized with Jira. If disabled, you need to choose individual findings to sync at each finding's details.

### Status mapping

Maps Snyk API & Web status to Jira status:

* **Not Fixed**: The initial state of the finding, right after being reported.
* **Invalid** (optional): Manually changed by the user to report a false positive.
* **Accepted Risk** (optional): Manually changed by the user, who acknowledges the finding but accepts its risk and will not fix it.
* **Fixed**: Snyk API & Web confirms that the finding is fixed.

Each Jira status can only be used once for each project.

### Priority mapping

Maps Snyk API & Web severity to Jira priority:

* **Critical**: A finding that represents a critical risk for the application and requires immediate remediation.
* **High**: A finding that represents a high risk for the application, whose exploitation can relatively easily cause damage to the application.
* **Medium**: A finding that alone is unlikely to cause damage to the application. However, if combined with another one, or in rare situations by itself, it can cause damage.
* **Low**: Alone, a low risk finding will not compromise your application, except in extreme situations. The attacker will normally require another higher risk finding to be able to take advantage of this one.

Once you are done, do not forget to click **Save** so that your changes are not lost.
