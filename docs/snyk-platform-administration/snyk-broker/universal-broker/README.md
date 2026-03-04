# Universal Broker

The Universal Broker improves the management of Broker deployments and connections by supporting many connections of any type with a single running client (container) or with multiple client replicas. Credentials remain within your network and are never stored by or transmitted to Snyk. The Snyk platform uses credentials references to instruct the Universal Broker how to authenticate with your private resources.

Each client or set of client replicas is called a Broker deployment. The diagram below shows two deployments, Universal Broker A and Universal Broker B.

A deployment can support multiple connections of any type, as shown in the diagram examples: GitHub, GitLab, Artifactory, Jira, and Container Registry Agent. Connections are configured to communicate with specific private resources: SCMs, JIRA, and others.

<figure><img src="../../../.gitbook/assets/image 5 (6).png" alt=""><figcaption><p>Universal Broker deployment example</p></figcaption></figure>

Connections are integrated with Organizations to provide access to your private resources for the appropriate Snyk Organization(s). These Organizations can be in the same or different Snyk Groups.

In the diagram, Group 1 includes Organizations A through D, and Group 2 includes Organizations E and F.

Organizations A, B, C, D, and E are integrated with the Universal Broker A deployment, and thus have access to all of the resources except the Container Registry Agent.

Organization F is integrated with Universal Broker B, and thus has access to Jira and the Container Registry Agent.

{% hint style="info" %}
You can learn more about Universal Broker in the dedicated Snyk Learn course. Explore the advantages, configuration, architecture, and much more with [Snyk Learn: Universal Broker](https://learn.snyk.io/lesson/universal-broker/?ecosystem=general).
{% endhint %}

To get started with using Universal Broker, be sure you have met the [prerequisites](prerequisites-for-universal-broker.md) and follow the [basic steps to install and configure Universal Broker](basic-steps-to-install-and-configure-universal-broker.md).
