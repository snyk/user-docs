# List and review the Universal Broker configuration resources

The `snyk-broker-config` CLI tool provides resources to guide you in configuring your Universal Broker. The following types of resources are available:

* Connections
* Credentials
* Deployments
* Integrations

To list the workflows available, run `snyk-broker-config-workflows`.

```
> snyk-broker-config workflows
Interactive workflows for Deployments, Credentials, Connections and 
Integrations management

USAGE
     $ snyk-broker-config workflows COMMAND

TOPICS
      workflows connections   Universal Broker - Create Connection Workflow
      workflows credentials   Universal Broker - Create Credentials Workflow
      workflows deployments   Universal Broker - Create Deployment Workflow
      workflows integration   Universal Broker - Get Connection Integration Workflow
```

Run `snyk-broker-config workflows <RESOURCE>` to list the available resources to get, create, delete, disconnect, integrate, and migrate.

Run the available workflows and follow the prompts to use the specific resources.

```
> snyk-broker-config-workflows connections
Universal Broker - Create Connection Workflow

USAGE
  > snyk-broker-config workflows connections COMMAND
  
COMMANDS
 workflows connections create     Universal Broker - Create Connection Workflow
 workflows connections delete     Universal Broker - Delete Connection Workflow
 workflows connections disconnect Universal Broker - Connection Disconnect 
Integration(s) Workflow
 workflows connections get        Universal Broker - Get Connection Workflow
 workflows connections integrate  Universal Broker - Connection Create Integration(s) 
Workflow
 workflows connections migrate    Universal Broker - Existing Connection Migration 
Workflow
```

For a complete list of the available workflows and commands, see the list of [`snyk-broker-config` commands](https://github.com/snyk/snyk-broker-config?tab=readme-ov-file#commands) in the repository.
