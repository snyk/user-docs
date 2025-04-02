# Getting started with Universal Broker

The basic steps in implementing the Universal Broker are the same regardless of the configuration method you use. The Getting started pages explain how to install and use the `snyk-broker-config` CLI tool, which walks you through the following basic steps.

1. Create your first connection. This returns an install ID, a client ID, and a client secret, which are all needed to interact with the Snyk platform. The Snyk Broker Admin Organization ID is required to create the connection.
2. Create your first deployment. This includes creating credentials references needed for your connections. Snyk recommends that you set environment variables at this time.
3. Integrate your connection with an Organization.
4. Run the Broker Client to verify that your connection is configured.
5. Integrate your connection with additional Organizations that should use the connection.

For details concerning each step, see [Basic steps to install and configure Universal Broker](basic-steps-to-install-and-configure-universal-broker.md). Before you follow the steps, be sure you have met the [prerequisites for Universal Broker](prerequisites-for-universal-broker.md).
