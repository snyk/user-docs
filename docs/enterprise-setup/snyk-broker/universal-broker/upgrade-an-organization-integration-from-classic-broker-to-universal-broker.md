# Upgrade an Organization integration from Classic Broker to Universal Broker

Universal Broker operators declare their desired deployment model before running any Broker client, specifying what Broker connections to support. Thus the Classic Broker approach of `org->integrations->broker connections` is evolving to be `broker connections -> integration/org`.

To upgrade existing Classic Broker integrations to Universal Broker, follow these steps for one Organization at a time.&#x20;

1. Create the Universal Broker connection you want to use. For details, see [Basic steps to install and configure Universal Broker](basic-steps-to-install-and-configure-universal-broker.md).
2. Run the Universal Broker client. For details, see [Running your Universal Broker client](running-your-universal-broker-client.md).
3. Use a test Organization to test the Universal Broker connection. Run `snyk-broker-config workflows connections integrate`  and enter the test Organization ID, using the same credentials as were used for the Classic Broker connection. The objective is to have the same access as the Classic Broker.
4. When you are ready, run `snyk-broker-config workflows connections migrate`.
   1. Select the deployment and connection you want to upgrade to.
   2. Enter the Organization ID for the connection you want to upgrade to the Universal Broker.
   3. Confirm your choice.\
      When the command completes, the connection to the Organization has been upgraded.
5. Repeat the procedure for all the relevant connections integrated with each Organization.
6. When you have migrated all of the integrations to the Universal Broker client, delete or remove the Classic Broker container.

If you run into issues, you can roll back to the Classic Broker client as long as you have not terminated the Classic Broker client. To roll back:

1. Disconnect the Universal Broker connection that you migrated to previously.
2. IF any other Classic Broker Organization is left with the same Broker token, after you disconnect the Universal Broker token, use the API endpoint [Clone an integration (with settings and credentials) ](../../../snyk-api/reference/integrations-v1.md#post-org-orgid-integrations-integrationid-clone)to copy the integration settings from another Organization and reuse the previously used Classic Broker token through the API.
3. IF there is no Classic Broker Organization left with the old Broker token, then after you disconnect the Universal Broker connection, set up a new Classic Brokered connection. Copy the Broker token from the new Brokered connection as a parameter and restart the Broker container.
