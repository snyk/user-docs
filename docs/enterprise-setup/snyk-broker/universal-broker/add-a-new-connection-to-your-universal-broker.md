# Add a new connection to your Universal Broker

To add a new connection, run  `snyk-broker-config workflows connections create`. If prompted, select the desired deployment.

Then follow the prompts to create the connection of the desired type. For details, see [Setting up and integrating your Universal Broker connections](setting-up-and-integrating-your-universal-broker-connections.md).

Note the following information about Broker deployments.

A Broker deployment supports numerous connections of any type, including the same type. This means, for example, that a Broker deployment can support multiple GitLab connections. There is a limit of 25 connections for one deployment.

While the Universal Broker supports multiple connections of the same type, each Organization can be integrated with only one connection for a given type. For example, an Organization cannot be integrated with two GitLab connections.

You can share credentials references across connections by selecting a credentials reference when you create a connection. However, you can share credentials references only across connections of the same type AND running on the same Broker deployment.

After the connection is created, run `snyk-broker-config workflows connections integrate` and select the desired deployment and connection. Then enter the Organization ID for the Organization with which you want to integrate. For details, see [Integrate a connection with an Organization](setting-up-and-integrating-your-universal-broker-connections.md#integrate-a-connection-with-an-organization).

If you integrate a connection already integrated with an Organization, the previous integration will be lost in favor of the new one. Be sure to paste the correct Organization ID.

\
