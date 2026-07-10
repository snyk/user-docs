---
description: How to integrate Snyk API and Web with Jira Server
---

# Integrate with Jira Server

By connecting Snyk API & Web to your Jira Server, you can synchronize target scan results with a Jira project of your choice. Snyk can do this synchronization automatically or manually, finding by finding.

The synchronization is bi-directional. Snyk sends a reported finding to Jira, and as soon as the finding is closed, Snyk triggers a retest. If the finding is fixed, the Jira issue remains closed. Otherwise, Snyk reopens it.

Connecting Snyk to your Jira Server takes no more than five minutes if you follow these instructions.

## Generate an RSA public/private key pair

Jira validates the identity of the Snyk server by requiring it to use a certificate. You can use any RSA public/private key pair, so you can skip the generation if you want to use another pair.

When prompted for the certificate details, you can enter anything, including the default values, by pressing **Enter**. The exception is the Common Name, where you must enter something.

```bash
openssl genrsa -out jira_privatekey.pem 2048

openssl req -newkey rsa:2048 -x509 -key jira_privatekey.pem -out jira_publickey.cer -days 1825

openssl pkcs8 -topk8 -nocrypt -in jira_privatekey.pem -out jira_privatekey.pcks8

openssl x509 -pubkey -noout -in jira_publickey.cer > jira_publickey.pem
```

## Create a new Jira application link

1. In Jira Server, in the top-right corner, navigate to **gear icon > Applications** and then click **Application Links** under **Integrations**.
1. Enter `https://plus.probely.app/jira-server/` in the input field and then click **Create new link**.

Ignore the warning "No response was received from the URL you entered" and click **Continue**.

1. On the next dialog, enter the following:
   * **Application Name**: Probely
   * **Application Type**: Generic Application
   * **Service Provider Name**: Probely
   * **Consumer key**: The value does not matter, but you need it later
   * **Shared secret**: probely
   * **Request Token URL**: `https://plus.probely.app/jira-server/`
   * **Access Token URL**: `https://plus.probely.app/jira-server/`
   * **Authorize URL**: `https://plus.probely.app/jira-server/`
1. Click **Continue**. You now see Probely on your application links.
1. Edit the Probely application link (pencil icon on the right) and fill out the Incoming Authentication form as follows:
   * **Consumer Key**: Same key as in the previous form
   * **Consumer Name**: Probely
   * **Public Key**: Public key created at the start (in the jira\_publickey.pem file)
   * **Callback URL**: `https://plus.probely.app/jira-server/callback/`
1. Click **Save**.

## Connect Snyk API & Web

1. In your Snyk account, open the **Settings** dropdown menu in the bottom-left corner and click **Integrations**.
1. Fill out the **Jira Server** form as follows:
   * **Server URL**: URL for your Jira Server instance
   * **Consumer Key**: Same as previous steps
   * **Consumer Secret**: Private key created at the start (in the jira\_privatekey.pem file)
   * **Verify TLS**: On (do not turn this option off without a very good reason)
1. Click **Authorize**.
1. Click **Allow** to allow Snyk to access your Jira Server.

Your Jira Server is now connected to your Snyk account.

The next step is configuring the Snyk targets you want to synchronize. For each target whose findings you want to synchronize, navigate to its **Target Settings > Integrations > Jira Server** and configure how they synchronize.

Visit [Configure Jira synchronization settings](configure-jira-synchronization-settings.md) for more information about the configuration options.
