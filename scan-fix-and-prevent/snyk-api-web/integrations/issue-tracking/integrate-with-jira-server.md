# Integrate with Jira Server

By connecting Snyk API & Web to your Jira Server, you can synchronize target scan results with a Jira project of your choice. This synchronization can be done automatically or manually, finding by finding.

The synchronization is bi-directional. A finding reported by Snyk API & Web is sent to Jira, and as soon as it is closed, Snyk API & Web triggers a retest. If the finding is fixed, the Jira issue remains closed. Otherwise, it is reopened.

Connecting Snyk API & Web to your Jira Server takes no more than five minutes by following the instructions below.

## Generate an RSA public/private key pair

Jira validates the identity of the Snyk API & Web server by requiring it to use a certificate. You can use any RSA public/private key pair, so you can skip the generation if you want to use another pair.

When prompted for the certificate details, you can enter whatever you want, including using the default values by pressing Enter, except for the Common Name, where you must enter something (anything will do).

```bash
openssl genrsa -out jira_privatekey.pem 2048

openssl req -newkey rsa:2048 -x509 -key jira_privatekey.pem -out jira_publickey.cer -days 1825

openssl pkcs8 -topk8 -nocrypt -in jira_privatekey.pem -out jira_privatekey.pcks8

openssl x509 -pubkey -noout -in jira_publickey.cer > jira_publickey.pem
```

## Create a new Jira application link

1. In Jira Server, at the top right corner, go to **gear icon > Applications** and then click **Application Links** under **Integrations**.
1. Enter `https://plus.probely.app/jira-server/` in the input field and then click **Create new link**.

Ignore the warning "No response was received from the URL you entered" and click **Continue**.

1. On the next dialog, input the following:
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
   * **Public Key**: Public key created in the beginning (in the jira\_publickey.pem file)
   * **Callback URL**: `https://plus.probely.app/jira-server/callback/`
1. Click **Save**.

## Connect Snyk API & Web

1. In your Snyk API & Web account, open the **Settings** dropdown menu on the bottom-left corner and click **Integrations**.
1. Fill out the **Jira Server** form as follows:
   * **Server URL**: URL for your Jira Server instance
   * **Consumer Key**: Same as previous steps
   * **Consumer Secret**: Private key created at the beginning (in the jira\_privatekey.pem file)
   * **Verify TLS**: On (do not turn this option off without a very good reason)
1. Click **Authorize**.
1. Click **Allow** to allow Snyk API & Web to access your Jira Server.

Your Jira Server is now connected to your Snyk API & Web account.

The next step is configuring the Snyk API & Web targets you want to synchronize. For each target you want to synchronize the findings, go to its **Target Settings > Integrations > Jira Server** and configure how they synchronize.

Visit [Configure Jira synchronization settings](configure-jira-synchronization-settings.md) for more information about the configuration options.
