# Prerequisites for Universal Broker

{% hint style="info" %}
Using Snyk Broker on Windows is not supported. Snyk recommends that Windows users deploy Broker using Linux.

Using the snyk-broker-config CLI tool is supported on Windows.
{% endhint %}

Before installing the Universal Broker `snyk-broker-config` CLI tool, be sure you have met the following prerequisites.

* Client machine system requirements: one CPU and 256 MB of RAM.
* Network access that is allowed by any firewalls installed on your network: an outbound TLS (443) to <kbd>https://broker.snyk.io</kbd> AND `https://api.snyk.io`or your [regional Broker URL](../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-urls).
* A Snyk account and your personal Snyk token or your PAT; you cannot use a service account.
* An SCM token or password. See [Obtain the tokens required to set up Snyk Broker](../classic-broker/prepare-snyk-broker-for-deployment/obtain-the-tokens-required-to-set-up-snyk-broker.md). Snyk Broker does not support authentication with the mTLS method.
* Tenant admin permissions.
* Node 18 or higher installed.
* Docker configured to pull images from Docker Hub in order to install with Docker.
* A new Snyk Broker Admin Organization created in the Group of your choice in your Tenant, not a personal Organization.

When installing, you must add a command in your script to set the Broker server URL for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-urls).\
Example:  `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`.

If you are using a region other than SNYK US-01, run the following as applicable after you install the CLI tool and before you create your first connection:\
`- export SNYK_API_HOSTNAME=https://api.REGION.snyk.io` (Linux and Mac).\
`- set SNYK_API_HOSTNAME=https://api.REGION.snyk.io` (Windows)

\
