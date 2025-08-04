# Enable the Kubernetes integration

{% hint style="info" %}
**Feature availability**&#x20;

Kubernetes integration is available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

## Set up the Kubernetes integration for your Organization

To integrate your Snyk Organization with Kubernetes:

1. On the main page of the Snyk Web UI, navigate to **Integrations**.
2. Select **Container orchestrators** and then select **Kubernetes.**
3. Click **Connect**. On the page that appears, you can find the **Integration ID** which is created for the Kubernetes integration. The **Integration ID** is required during the Snyk Controller installation.
4. Create a [Group or Organization service account token](../../../../implementation-and-setup/enterprise-setup/service-accounts/). This service account token is required during the Snyk Controller installation. For more information, including the permissions required, see [Prerequisites for installing the Snyk Controller](../install-the-snyk-controller/#prerequisites-for-installing-the-snyk-controller).

## View and manage your Organization integration settings

To view and manage the Organization integration settings:

1. On the main page of the Snyk Web UI, navigate to **Settings**.
2. On the **ORGANIZATION SETTINGS** page, navigate to **Integrations**.
3. From the list of integrations, navigate to **Kubernetes** and click **Edit** **settings**. The Kubernetes settings page appears.

In this window, the following sections are available:

| Section                        | Description                                                                                                                                                                                                                                                                                                                                                                                                               |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Connected to Kubernetes**    | <p>This section appears after you have successfully connected, allowing you to add your workloads from the <strong>Dashboard</strong>, from the <strong>Projects</strong> page.</p><p></p><p>When the Kubernetes integration is not yet set up, this is the only section that appears in this window, together with the <strong>Connect</strong> button. See <a href="./">Overview of the Kubernetes integration</a>.</p> |
| **Integration ID**             | <p>The <strong>Integration ID</strong> is a UUID and appears in the format <code>abcd1234-abcd-1234-abcd-1234abcd1234</code>.</p><ul><li>If you have not set up the connection yet, click <strong>Connect</strong> and use this ID to set up your configuration.</li><li>If you have set up configuration, this information is visible on the page. </li></ul>                                                            |
| **Snyk controller versions**   | Displays the version of the Snyk controller you have installed on your clusters.                                                                                                                                                                                                                                                                                                                                          |
| **Disconnect from Kubernetes** | Allows you to remove this integration from this Organization. To do this, click **Disconnect**.                                                                                                                                                                                                                                                                                                                           |
