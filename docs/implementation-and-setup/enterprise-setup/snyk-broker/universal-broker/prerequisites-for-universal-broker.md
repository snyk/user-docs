# Prerequisites for Universal Broker

{% hint style="info" %}
Using Snyk Broker on Windows is not supported. Snyk recommends that Windows users deploy Broker using Linux.

Using the snyk-broker-config CLI tool is supported on Windows.
{% endhint %}

Before installing the Universal Broker `snyk-broker-config` CLI tool, be sure you have met the following prerequisites. If you need help, contact your Snyk account team.

* Minimum client machine system requirements: 1 CPU and 256 MB RAM. \
  Note: If you're performing a high-load import, we would recommend more RAM be allocated. This could be decreased after traffic is lowered.
* Network access that is allowed by any firewalls installed on your network: an outbound TLS (443) to <kbd>https://broker.snyk.io</kbd> AND `https://api.snyk.io` or your [regional Broker URL](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
* A Snyk account and your personal Snyk API token; you cannot use a service account.
* Snyk Tenant admin permissions. If you are not a Tenant admin, you can reach out to your team's Tenant admin [to add you](../../../../snyk-platform-administration/groups-and-organizations/tenant/manage-users-in-a-tenant.md). Otherwise, reach out to your support team member or [raise a support case](https://support.snyk.io/s/).
* A new/dedicated Snyk Organization. This will be used to administrate your Broker configuration(s) and a dedicated organization will help prevent accidental removal. See [Create an Organization](../../../../snyk-platform-administration/groups-and-organizations/organizations/create-and-delete-organizations.md#create-an-organization) for details.
* An SCM token or password. Snyk Broker does not support authentication with the mTLS method.
* Node 20 or higher installed.
* Docker configured to pull images from Docker Hub in order to install with Docker.

Snyk recommends that you export SNYK\_TOKEN and TENANT\_ID in your terminal session environment variables now, using the following commands:

Linux/Mac

* `export SNYK_TOKEN=xxxx`
* `export TENANT_ID=yyyy`

Windows

* `set SNYK_TOKEN=xxxx`
* `set TENANT_ID=yyyy`

When running the Broker container, you may need to override the Broker server URL to the region where your data is hosted if you are using a region other than SNYK US-01. For the commands and URLs to use, see [Broker URLs](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`.

If you are using a different region than SNYK US-01 you should also export the relevant `SNYK_API_HOSTNAME` value into your terminal's environment variables after you install the `snyk-broker-config` CLI tool and before you create your first connection:

* `export SNYK_API_HOSTNAME=https://api.REGION.snyk.io` (Linux/Mac)
* `set SNYK_API_HOSTNAME=https://api.REGION.snyk.io` (Windows)
