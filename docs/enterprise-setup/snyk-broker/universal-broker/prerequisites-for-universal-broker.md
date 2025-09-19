# Prerequisites for Universal Broker

{% hint style="info" %}
Using Snyk Broker on Windows is not supported. Snyk recommends that Windows users deploy Broker using Linux.

Using the snyk-broker-config CLI tool is supported on Windows.
{% endhint %}

Before installing the Universal Broker `snyk-broker-config` CLI tool, be sure you have met the following prerequisites. If you need help, contact your Snyk account team.

* Client machine system requirements: one CPU and 256 MB of RAM.
* Network access that is allowed by any firewalls installed on your network: outbound TLS (443) to `https://broker.snyk.io` AND `https://api.snyk.io` or your [regional Broker URL](../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-server-urls).
* A Snyk account and your personal Snyk API token; you cannot use a service account.
  * See [Obtain and use your Snyk API token](../../../getting-started/#obtain-and-use-your-snyk-api-token) for instructions on finding your SNYK\_TOKEN.
* Snyk Tenant admin permissions. If you are not a Tenant admin, you can reach out to your team's Tenant admin [to add you](../../../snyk-admin/tenant-groups-and-organizations/tenant/manage-users-in-a-tenant.md). Otherwise, reach out to your support team member or [raise a support case](https://support.snyk.io/s/).
* A new Snyk Broker Admin Organization created in the Group of your choice in your Tenant, not a personal Organization. See [Create an Organization](../../../snyk-admin/groups-and-organizations/organizations/create-and-delete-organizations.md#create-an-organization) for details.
* The TENANT\_ID of the Snyk Tenant you wish to use.
  * To find your TENANT\_ID:
    1. Log into the Snyk UI.
    2. Click your **Tenant name**. If you are a member of multiple Tenants, navigate to the Tenant where the Universal Broker will be configured.
    3. Copy the Tenant ID from the URL. The format is `https://app.snyk.io/tenant/1234abcd-1234-abcd-5678-1234abcd5678`.
* An SCM token or password. Snyk Broker does not support authentication with the mTLS method.
* Node 20 or higher installed.
* Docker installed and configured to pull images from Docker Hub.

Snyk recommends that you export SNYK\_TOKEN and TENANT\_ID in your terminal session environment variables now using the following commands:

(Linux/Mac)
* `export SNYK_TOKEN=xxxx`
* `export TENANT_ID=yyyy`
(Windows)
* `set SNYK_TOKEN=xxxx`
* `set TENANT_ID=yyyy`

When running the Broker container, you may need to override the Broker server URL to the region where your data is hosted if you are using a region other than SNYK US-01. For the commands and URLs to use, see [Broker URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-server-urls).\
Example: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`.

If you are using a different region than SNYK US-01 you should also export the relevant `SNYK_API_HOSTNAME` value into your terminal's environment variables after you install the `snyk-broker-config` CLI tool and before you create your first connection:\
`export SNYK_API_HOSTNAME=https://api.REGION.snyk.io` (Linux and Mac).\
`set SNYK_API_HOSTNAME=https://api.REGION.snyk.io` (Windows)

\
