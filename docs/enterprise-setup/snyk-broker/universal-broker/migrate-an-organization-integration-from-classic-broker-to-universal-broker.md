# Migrate an Organization integration from Classic Broker to Universal Broker

Universal Broker operators declare their desired deployment model before running any Broker client, specifying what Broker connections to support. Thus the Classic Broker approach of `org->integrations->broker connections` is evolving to be b`roker bonnections -> integration/org`.

To migrate existing Classic Broker integrations to Universal Broker, follow these steps for one Organization at a time.&#x20;

1. Create the Universal Broker connection you want to use. For details, see [Basic steps to install and configure Universal Broker](basic-steps-to-install-and-configure-universal-broker.md).
2. Run the Universal Broker client. For details, see [Running your Universal Broker client](running-your-universal-broker-client.md).
3. Use a test Organization to test the Universal Broker connection. Run `snyk-broker-config workflows connections integrate`  and enter the test Organization ID, using the same credentials as were used for the Classic Broker connection. The objective is to have the same access as the Classic Broker
4. When you are ready, run `snyk-broker-config workflows connections migrate`.
   1. Select the deployment and connection you want to migrate to.
   2. Enter the Organization ID for the connection you want to migrate to the Universal Broker.
   3. Confirm your choice.\
      When the command completes, the connection to the Organization has been migrated.
5. Repeat the procedure for all the relevant connections integrated with each Organization.
