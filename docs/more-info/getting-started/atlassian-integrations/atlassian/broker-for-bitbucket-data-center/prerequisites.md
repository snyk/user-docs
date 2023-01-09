---
description: What you need for Snyk Broker for Bitbucket Data Center
---

# Prerequisites

This tutorial assumes you have the following:

* Bitbucket Data Center running on your infrastructure or in the cloud
* A Snyk account

{% hint style="info" %}
The base requirements for Snyk Broker are described in the [snyk-broker](../../../../../integrations/snyk-broker/ "mention") docs.\
We highlight some of the requirements necessary.
{% endhint %}

You have several options to install Snyk Broker on your infrastructure, including a container or an AWS Quick Start.

To get started, you'll need:

* A Broker Token for the specific type of integration you are using, provided by the support team at Snyk. For this Tutorial, we'll use the Bitbucket Server connection. See the [list ](../../../../../tutorials/atlassian/broker-for-bitbucket-data-center/broken-reference/)on our documentation site for examples of other options.
* Administrative permissions on your Bitbucket Data Center instance.
* Permissions on snyk.io to import projects.
* A successful installation of Snyk Broker. You would have provided the following information the configuration of that instance as outlined in the [Snyk documentation](../../../../../integrations/git-repository-scm-integrations/bitbucket-data-center-server-integration.md):
  * A Bitbucket Username
  * A Bitbucket Password
  * A Bitbucket Hostname
  * A Bitbucket API endpoint
