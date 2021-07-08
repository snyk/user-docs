# Retrieve a unique Broker client token

* [ Broker introduction](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015367178-Broker-introduction/README.md)
* [ Set up Snyk Broker](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015296637-Set-up-Snyk-Broker/README.md)
* [ Prepare Snyk Broker for deployment](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015296717-Prepare-Snyk-Broker-for-deployment/README.md)
* [ Enable permissions for Snyk Broker from your third-party tool](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015296737-Enable-permissions-for-Snyk-Broker-from-your-third-party-tool/README.md)
* [ Retrieve a unique Broker client token](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015367278-Retrieve-a-unique-Broker-client-token/README.md)
* [ How to install and configure your Snyk Broker client](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015296937-How-to-install-and-configure-your-Snyk-Broker-client/README.md)
* [ Upgrade your Snyk Broker client](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015367458-Upgrade-your-Snyk-Broker-client/README.md)

## Retrieve a unique Broker client token

To retrieve the token enabling you to install the Broker:

1. Enable Broker functionality. For code repository \(SCM\) integrations, Snyk Broker can be enabled via the Integrations API or by contacting support. For [Artifactory](https://support.snyk.io/hc/en-us/articles/360007537418) or Jira integrations, you can enable using the Snyk UI, or Snyk support can assist you with enabling this feature.
2. Click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt; **Integrations** for that specific integration update to see the Broker token. 
3. From the row for the integration you’re setting up, click **Edit settings**. \(As you have not yet installed and configured the client, the notification from this screen correctly displays “Could not connect to…”.\)
4. Copy and paste the credentials / token in a file on the desktop to use in the environment variables of the command line argument when you install the client.

### Enabling across multiple organizations

You can use the same Git across multiple organizations in Snyk, using the same Broker token. To do this, create the token for an organization, then create a new organization based on the original; this clones the token and you can now enable the Broker for it.

