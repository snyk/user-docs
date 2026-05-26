# Prerequisites for Universal Broker

{% hint style="info" %}
Using Snyk Broker on Windows is not supported. Snyk recommends that Windows users deploy Broker using Linux.

Using the snyk-broker-config CLI tool is supported on Windows.
{% endhint %}

Before installing the Universal Broker `snyk-broker-config` CLI tool, be sure you have met the following prerequisites. If you need help, contact your Snyk account team.

* Minimum client machine system requirements: 1 CPU and 512 MB RAM.\
  Note: If you're performing a high-load import, we would recommend more RAM be allocated. This could be decreased after traffic is lowered.
* Network access that is allowed by any firewalls installed on your network: an outbound TLS (443) to `https://broker.snyk.io` AND `https://api.snyk.io` or your [regional Broker URL](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/regional-hosting-and-data-residency#broker-server-urls).
* A Snyk account and your personal Snyk API token; you cannot use a service account.
* Snyk Tenant admin permissions. If you are not a Tenant admin, you can reach out to your team's Tenant admin [to add you](../../../../snyk-platform-administration/groups-and-organizations/tenant/manage-users-in-a-tenant.md). Otherwise, reach out to your support team member or [raise a support case](https://support.snyk.io/s/).
* A new/dedicated Snyk Organization. This will be used to administrate your Broker configurations, and a dedicated Organization will help prevent accidental removal. See [Create an Organization](../../../../snyk-platform-administration/groups-and-organizations/organizations/create-and-delete-organizations.md#create-an-organization) for details.
* An SCM token or password. Snyk Broker does not support authentication with the mTLS method.
* Node 20 or higher installed.
* Docker Compose or Kubernetes installed and configured to pull images from Docker Hub.

Snyk recommends that you export SNYK\_TOKEN and TENANT\_ID environment variables to avoid having to paste the values in for every command. Use the following commands to set these environment variables:

Linux/Mac

* `export SNYK_TOKEN=<your_snyk_token>`
* `export TENANT_ID=<your_tenant_id>`

Windows

* `set SNYK_TOKEN=<your_snyk_token>`
* `set TENANT_ID=<your_tenant_id>`

When running the Broker container, you may need to override the Broker server URL to the region where your data is hosted if you are using a region other than SNYK US-01. For the commands and URLs to use, see [Broker URLs](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/regional-hosting-and-data-residency#broker-server-urls).\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`.

If you are using a different region than SNYK US-01 you should also export the relevant `SNYK_API_HOSTNAME` value into your terminal's environment variables after you install the `snyk-broker-config` CLI tool and before you create your first connection:

* `export SNYK_API_HOSTNAME=https://api.REGION.snyk.io` (Linux/Mac)
* `set SNYK_API_HOSTNAME=https://api.REGION.snyk.io` (Windows)
