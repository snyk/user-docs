# Retrieve a unique Broker client token

In order to use the Broker, the Broker token is required and must be generated.&#x20;



For code repository (SCM) integrations, a Broker token can be generated via API , or by contacting Snyk Support

1. Go to our API documentation and follow the example under "Set up a broker for an existing integration" within the [Integrations API](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration) or by contacting support.&#x20;
2. Verify the Broker token is generated in the Snyk UI under the specified SCM integration. by clicking on settings ![cog\_icon.png](../../../.gitbook/assets/cog\_icon.png) > **Integrations** for that specific integration update to see the Broker token. &#x20;

For [Artifactory](https://docs.snyk.io/integrations/private-registry-integrations/artifactory-registry-for-npm) or [Jira](https://docs.snyk.io/features/integrations/notifications-ticketing-system-integrations/jira) integrations, a Broker token can be generated within the Snyk UI, or by contacting Snyk support

1. Click on settings ![cog\_icon.png](../../../.gitbook/assets/cog\_icon.png) > **Integrations** for that specific integration to generate the Broker token.&#x20;
2. Once the broker token is generated, under the integration, you will see the notification from this screen correctly displays “Could not connect to…”.) as you have not yet installed and configured the client,&#x20;
3. Copy and paste the broker token from the UI to use it when you install the client.

## Enabling across multiple organizations

You can use the same Git across multiple organizations in Snyk, using the same Broker token. To do this, create the token for an organization, then create a new organization based on the original; this clones the token and you can now enable the Broker for it.
