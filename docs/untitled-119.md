# Clone an integration across your Snyk orgs

* [ Service accounts](/hc/en-us/articles/360004037597-Service-accounts)
* [ Managing integrations](/hc/en-us/articles/360004002498-Managing-integrations)
* [ Clone an integration across your Snyk orgs](/hc/en-us/articles/360004008298-Clone-an-integration-across-your-Snyk-orgs)
* [ Disable a git integration](/hc/en-us/articles/360004008318-Disable-a-git-integration)

##  Clone an integration across your Snyk orgs

You can choose to use the same brokered Git integration across multiple organizations in Snyk by simply copying and duplicating the organization you've already configured. For example, you can integrate Snyk organizations X, Y and Z with your single Git repo X.

**Prerequisites**: in order to clone organization configurations, teams and groups must first be enabled.

1. From the **organization** dropdown, navigate to any organization within the group that you are working with.
2. Now, from the same **organization** dropdown, navigate to click **Create a new organization**:
3. From the page that loads, enter a name for the new organization that you are creating.
4. From the **Copy across all settings and integrations?** area, select the organization that you've already configured for the Broker token and then click **Create organization**.
5. The browser navigates to the **Dashboard** for the organization that you just created. The Broker integration is duplicated and setup, and the Broker token is identical to the token for the original organization.
6. To double-check your cloned configuration, click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Integrations.**
7. From the row for the integration you’re setting up, click **Edit settings** to see the cloned Broker integration.

