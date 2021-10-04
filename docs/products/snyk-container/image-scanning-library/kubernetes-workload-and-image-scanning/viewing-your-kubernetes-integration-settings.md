# Viewing your Kubernetes integration settings

## In the relevant organization:

1. Click on settings ![](../../../../.gitbook/assets/cog_icon.png) &gt; **Integrations**. 
2. Navigate to Kubernetes and click **Edit** **Settings**. 
3. Navigate to the **Integration ID** and other integration settings:

![](../../../../.gitbook/assets/uuid-03a03790-d87e-6260-4ffc-dc474ce014fa-en.gif)

From this window, you can access and work with the following:

| Part | Description |
| :--- | :--- |
| Connected to Kubernetes | When the integration is not yet set up, this is the only section that appears in this window, with the Connect button \(Integrate with these instructions: \([Kubernetes integration overview](kubernetes-integration-overview.md)\). If you just clicked the Connect button, and when you're connected successfully, this area should appear as in the image. |
| Integration ID | The Integration ID is a UUID and appears similar to this format: `abcd1234-abcd-1234-abcd-1234abcd1234`. Click Connect and use this ID to set up your configuration if you haven't already. If you've already set up configuration, this and the rest of the data in the window should now all appear immediately when you navigate to and land on this page. You can now add your workloads from here, from the Dashboard, from the Projects page or set up automatic importing from your cluster. [Adding Kubernetes workloads for security scanning](adding-kubernetes-workloads-for-security-scanning.md) |
| Snyk controller versions | Check the version of the Snyk controller that you've installed on your clusters from this area. |
| Disconnect from Kubernetes | To remove this integration from this organization, click Disconnect. |

