# CrowdStrike Falcon Next-Gen SIEM

{% hint style="info" %}
**Release status**

The CrowdStrike Falcon Next-Gen SIEM integration is in [Early Access](../../discover-snyk/getting-started/snyk-release-process.md#early-access-features), and is available only with Snyk Enterprise plans. To learn more, visit [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

Integrate Snyk vulnerability data into the [CrowdStrike Falcon® platform](https://www.crowdstrike.com/en-us/platform/next-gen-siem/) for unified security visibility across application and endpoint domains. For example, you can detect when a critical vulnerability is introduced in a deployed container image due to a newly discovered CVE in an open-source dependency.

The process for setting up this integration consists of:

1. [Configure the CrowdStrike Data Connector](crowdstrike-falcon-next-gen-siem.md#configure-the-crowdstrike-data-connector)
2. [Configure the Snyk Issue Forwarder](crowdstrike-falcon-next-gen-siem.md#configure-snyk-issue-forwarder)
3. [Verify the integration connection](crowdstrike-falcon-next-gen-siem.md#verify-the-integration-connection)

## **Prerequisites**

* A CrowdStrike subscription for Falcon Next-Gen SIEM or Falcon Next-Gen SIEM 10GB.
* Connector Manager access to the Falcon console for the relevant CID.
* A Snyk user account with [permissions](../../snyk-platform-administration/user-roles/user-role-management.md) to edit and view Group integrations.

## Configure the CrowdStrike Data Connector

To use the CrowdStrike NG-SIEM destination, you need to set up a Snyk Data Connector in your CrowdStrike NG-SIEM environment. To learn more, visit the [Snyk](https://falcon.us-2.crowdstrike.com/documentation/page/c7182a06/snyk) page in the CrowdStrike documentation. Select the **snyk-platform (Snyk Platform)** parser while configuring the data connector.

When setting up the data connector, you receive an API key and URL, which you will use later to configure the Snyk Issue Forwarder.

## Configure Snyk Issue Forwarder

This section configures Snyk to send issue data to the CrowdStrike connector you just created.

* Open the **Integrations** menu.
* Select the **Add integration** option.
* Select the **Issue Forwardin**g tag and search for CrowdStrike Issue Forwarding.
* Click the **Add** button.
* Add the profile name for this integration.
* Add the **API URL** you copied earlier from your CrowdStrike Issue Forwarding account.
* Add the **API key** you copied earlier from your CrowdStrike Issue Forwarding account.
* Click the **Done** button.
* When the connection is established, the status of the CrowdStrike Issue Forwarding integration is changed to **Connected**.

## Verify the integration connection

After the integration is set up, you can verify that Snyk data is being forwarded to CrowdStrike. For more details, visit the [Step 3: Verify successful data ingestion](https://falcon.us-2.crowdstrike.com/documentation/page/c7182a06/snyk#s4fea4f4) CrowdStrike page for instructions.

If you need to run a manual search, use this query in Advanced Event Search:

```
#Vendor = snyk and @error != "true"
```
