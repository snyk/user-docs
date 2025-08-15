# Clone an integration across your Snyk Organizations

You can choose to use the same brokered Git integration across multiple Organizations in Snyk by copying and duplicating the Organization you have already configured.

For example, you can integrate Snyk Organizations X, Y, and Z with your single Git repo X.

**Prerequisites**: to clone Organization configurations, you must have teams and groups enabled.

1. From the **Organization** list, choose an Organization in the Group that you are working with.\
   <img src="../../../../.gitbook/assets/switch_org_02oct2022.png" alt="Choose Organization" data-size="original">
2.  From the same **Organization** dropdown, click **+Create new Organization.**

    <figure><img src="../../../../.gitbook/assets/clone-organization1_02oct2022.png" alt="Select +Create new Organization"><figcaption><p>Select +Create new Organization</p></figcaption></figure>
3. In the next window, enter a name for your new Organization.
4. In the **Copy settings from an existing org** section, choose an Organization that you have already configured for the Broker token.
5.  Review the summary of what will be copied across to the new Organization and click **Create organization** to confirm.

    <figure><img src="../../../../.gitbook/assets/clone-org-3screens_02oct2022.png" alt="Review summary of what will be copied to the new Organization and Create"><figcaption><p>Review summary of what will be copied to the new Organization and Create</p></figcaption></figure>

The **Dashboard** for the Organization that you just created opens. The Broker integration is duplicated and set up, and the Broker token is identical to the token for the original Organization.
