# Setting up and integrating your Universal Broker connections

The following diagram illustrates installing the Snyk Broker App, which facilitates the secure connection and communication with the Snyk platform through OAuth, and creating a Universal Broker deployment with connections to GitHub and Jira. The process includes creating credentials references for both GitHub and Jira and creating Broker connections for both.

In the diagram, GitHub connection 123 is integrated with Organization Y using integration X. Jira connection 456 is integrated with Organization E using integration D.

<figure><img src="../../../../.gitbook/assets/Universal-Broker-deployment-connections.png" alt=""><figcaption><p>Universal Broker connections</p></figcaption></figure>

## Create deployments and connections

After you have installed the `snyk-broker-config` CLI tool:

* Run `snyk-broker-config workflows connections create`.
* Create a deployment if none exists, or select a deployment if more than one exists.\
  A single deployment setup will be selected automatically.

If you want to create a new deployment:

* Run `snyk-broker-config workflows deployments create`.
* Then run `snyk-broker-config workflows connections create` again, and select the new deployment.
* In `snyk-broker-config workflows connections create`, select the type of connection you want to create from the list presented.

Options include SCM connection types like GitHub and variants, Bitbucket server and variants, GitLab, and Azure, as well as package registry connections, Jira, container registry integrations, and more.

For container registry-type Broker connections, specify a CR\_AGENT\_URL to point to a Container Registry Agent. You must configure and run both the Universal Broker and a separate Container Registry Agent. Follow the instructions for configuring and running a Container Registry Agent.

* In response to the prompts in the workflow, provide the required fields for the configuration.

When you see the messages `Connection created` and `Ready to configure integrations to use this connection`, you can integrate the connection with the Organization.

Note that you can start the Broker client before the connection is integrated with any Organization. The Organization integration steps can be done in parallel with the Broker client container running.

If you are adding a connection to a running Broker deployment, the Broker client might take up to two minutes to pick up the new connection.

{% hint style="info" %}
You can learn more about importing a code repository from an SCM into Snyk and creating connections, and more, with the [Universal Broker Snyk Learn ](https://learn.snyk.io/lesson/universal-broker/#3f799f7f-58a9-4225-53fb-9cc5b6913920)course.
{% endhint %}

## Integrate a connection with an Organization

After you have created at least one deployment and connection:

* Run `snyk-broker-config workflows connections integrate`.
* Select the deployment and connection to use.\
  Selection is automatic if only one deployment and one connection exist.
* Enter the Organization ID you want to integrate the connection with.

The connection is now integrated with the Organization and ready to use.
