# Using the Integrations API to share Broker tokens between Snyk organizations

* [ Using the Integrations API to share Broker tokens between Snyk organizations](/hc/en-us/articles/360016692397-Using-the-Integrations-API-to-share-Broker-tokens-between-Snyk-organizations)
* [ API - Maven test using a different repository](/hc/en-us/articles/360004661298-API-Maven-test-using-a-different-repository)
* [ Apiary - API server took too long to respond](/hc/en-us/articles/360004661218-Apiary-API-server-took-too-long-to-respond)

##  Using the Integrations API to share Broker tokens between Snyk organizations

A common [Snyk Broker](https://support.snyk.io/hc/en-us/sections/360001138138-Snyk-Broker) architecture is to use the same single Broker client for multiple Snyk organizations:

This architecture reduces the number of Broker client instances you need to monitor and maintain, resulting in saving operational costs.

This articles describes how to build such an architecture, sharing the same Broker token between multiple integrations.

### Step 1: Create an integration with a shared Broker token

To create an integration with a shared Broker token:

1. Add a new integration that uses Broker, using the [add-new-integration](https://snyk.docs.apiary.io/#reference/integrations/integrations/add-new-integration) API.  
   Make the following call:

   ```text
   POST https://snyk.io/api/v1/org/orgId/integrations
   ```

    with the following body:

   ```text
   {
      "type": "github-enterprise",
      "broker": {
          "enabled": true
      }
   }
   ```

    The call returns the following response, that contains the new integration id and Broker token:   


   ```text
   {
       "id": "9a3e5d90-b782-468a-a042-9a2073736f0b",
       "brokerToken": "4a18d42f-0706-4ad0-b127-24078731fbed"
   }  
   ```

2. Clone the new integration to the other organizations that you want to share the token with, using the [clone-an-integration](https://snyk.docs.apiary.io/#reference/integrations/integration-cloning/clone-an-integration-%28with-settings-and-credentials%29) API.
3. Use the previously returned integration id to make the following call:

   ```text
   POST https://snyk.io/api/v1/org/orgId/integrations/integrationId/clone
   ```

    The call returns the following response, that contains the new integration id:  


   ```text
   {
       "newIntegrationId": "9a3e5d90-b782-468a-a042-9a2073736f0b"
   }
   ```

4. Run the Broker client within your infrastructure, using the newly generated Broker token that was returned in the first step.

### Step 2: Rotate a shared Broker token

To rotate a shared Broker token without causing downtime:

1. Issue a new and unique provisional token for the brokered integration, using the [provision-new-broker-token](https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-provisioning/provision-new-broker-token) API.  
   Make the following call:

   ```text
   POST https://snyk.io/api/v1/org/orgId/integrations/integrationId/authentication/provision-token
   ```

    The call returns the following response, that contains the newly provisioned token:

   ```text
   {
       "id": "9a3e5d90-b782-468a-a042-9a2073736f0b",
       "provisionalBrokerToken": "4a18d42f-0706-4ad0-b127-24078731fbed"
   }
   ```

    Make sure there is no other token that is currently provisioned within the group, otherwise, this operation will fail.

2. Rerun your Broker client with the newly provisioned token.
3. Switch the existing shared broker tokens with the provisioned token, using the [switch-between-broker-tokens](https://snyk.docs.apiary.io/#reference/integrations/integration-broker-token-switching/switch-between-broker-tokens) API.  
   Make the following call:

   ```text
   POST https://snyk.io/api/v1/org/orgId/integrations/integrationId/authentication/switch-token
   ```

    This action switches the Broker token for the specified integration, and for all integration within the group that share the same Broker token.  
   **Note**: this action requires a provisioned token for this integration within the group.

