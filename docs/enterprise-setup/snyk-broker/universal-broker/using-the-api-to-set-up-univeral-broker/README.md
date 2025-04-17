# Using the API to set up Univeral Broker

All flows available in the `snyk-broker-config` CLI tool are built on top of the public REST API. The workflows in the CLI tool implement particular flows to provide for ease of set up and use. The same workflows can be accomplished using  the API, allowing for automation.

The `snyk-broker-config` tool provides non-workflow commands to simplify the automation of particular tasks through simple bash scripting, avoiding the use of a more complicated API-only path.

See the [Broker APIs in the API Reference](https://apidocs.snyk.io/experimental?version=2024-10-14%7Eexperimental#get-/orgs/-org_id-/brokers/connections) for more details. An example is provided: [Using the API to set up a GitHub connection](using-the-api-to-set-up-a-github-connection.md).

The workflow diagrams illustrate the main workflows in setting up Universal Broker.

When you change an environment variable, you must restart your Broker.

\
