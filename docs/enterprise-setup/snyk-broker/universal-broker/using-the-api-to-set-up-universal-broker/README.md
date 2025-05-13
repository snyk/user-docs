# Using the API to set up Univeral Broker

All flows available in the `snyk-broker-config` CLI tool are built on top of the public REST API. The workflows in the CLI tool implement particular flows to provide for ease of setup and use. The same workflows can be accomplished using  the API, allowing for automation.

The `snyk-broker-config` tool provides non-workflow commands to simplify the automation of particular tasks through simple bash scripting, avoiding the use of a more complicated API-only path.

See the [Universal Broker APIs ](../../../../snyk-api/reference/universal-broker.md)in the API Reference for more details. An example is provided: [Using the API to set up a GitHub connection](using-the-api-to-set-up-a-github-connection.md).

The [workflow diagrams](universal-broker-workflow-diagrams.md) illustrate the main workflows in setting up Universal Broker.

When you change an environment variable, you must restart your Broker.

\
