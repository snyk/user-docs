# Universal Broker workflow diagrams

The following workflow diagrams illustrate the steps that are implemented in the `snyk broker config` tool when you use the commands to automate. The same workflows are implemented when you use the API.

## Create a deployment workflow

You start the workflow and choose to use an existing deployment or create a new one. The deployment is then created.

<figure><img src="../../../../../.gitbook/assets/Create-deployment-workflow.png" alt="" width="375"><figcaption><p>Use or create a new deployment</p></figcaption></figure>

## Create connection workflow

You start the workflow and select a deployment. You create a new connection for that deployment, selecting a connection type and entering the required parameters. You can use an existing or create a new credentials reference. The connection is then created.

<figure><img src="../../../../../.gitbook/assets/Create-connecton-workflow.png" alt="" width="563"><figcaption><p>Create a connection</p></figcaption></figure>

## Integrate connection workflow

You start the workflow and select a deployment. You select an existing connection to integrate with an Organization and enter the Org ID for that Organization. The connection is then integrated with the Organization by Org ID. The connection is ready.

<figure><img src="../../../../../.gitbook/assets/Integrate-connection-workflow.png" alt="" width="375"><figcaption><p>Integrate a connection with an Organizaton</p></figcaption></figure>
