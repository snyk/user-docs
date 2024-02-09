# TeamCity integration: install the Snyk plugin

Install or upgrade the Snyk Security plugin by following these steps. When the installation is complete, you can add a Snyk step to your projects.

**Note:** Before you begin, sign up for a Snyk account.

1. Log in to your TeamCity instance to install the Snyk Security plugin. Configure the **Plugins list** to **Periodically check for plugin updates**, in order to ensure regular automatic upgrades in the background.
2. Navigate to the [JetBrains Plugins Repository](https://plugins.jetbrains.com/plugin/12227-snyk-security), search for Snyk, and from the **Get** dropdown list, select the plugin for your TeamCity installation.
3. In response to the prompt, click **Install**.
4. When the installation ends, and the **Administration Plugins List** loads notifying you that the plugin has been uploaded, ensure the plugin is enabled.

![Install plugin from the JetBrains Plugins Repository](../../../.gitbook/assets/uuid-fe65f4bc-9578-016c-00dd-6ddb97d2ead7-en.png)

To configure the integration, see [TeamCity configuration parameters](teamcity-configuration-parameters.md). For information on how to configure your build with a Snyk step, see [Team City integration: use Snyk in your build](teamcity-integration-use-snyk-in-your-build.md).
