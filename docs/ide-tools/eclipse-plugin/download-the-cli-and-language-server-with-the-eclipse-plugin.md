# Download the CLI and language server with the Eclipse plugin

The Snyk Eclipse plugin works with an underlying language server for an optimal Eclipse experience. After restart, on opening a file that Snyk supports, the Eclipse plugin attempts to start a workspace scan:

![Eclipse plugin starting a scan](<../../.gitbook/assets/Screenshot 2022-05-13 at 09.28.30 (1).png>)

Once the plugin is installed, all of the plugin's prerequisites are triggered on opening a file that Snyk supports (Snyk hooks into this action). The prerequisites include downloading the Snyk CLI, the language server, and asking you to authenticate. All of these steps are shown in the following sections, in the order they happen.

### Snyk CLI and language server download

The process downloads the Snyk CLI and the language server and uses the CLI for authentication.

![Download the Snyk CLI](<../../.gitbook/assets/Screenshot 2022-05-13 at 11.27.00.png>)

### Authentication

Once the CLI is downloaded you will be redirected to the browser to authenticate and then connect the IDE with your account. Steps are as follows:

* You get a notification that a browser window will open.

![Notification, browser window opening for authentication](<../../.gitbook/assets/Screenshot 2022-05-13 at 11.29.37.png>)

* Once you are redirected to the browser for authentication, click Authenticate.

![Authenticate](<../../.gitbook/assets/Screenshot 2022-05-13 at 11.30.02.png>)

* You should see a successful message saying you've been authenticated.

![Confirmation of authentication](<../../.gitbook/assets/Screenshot 2022-05-13 at 11.30.30.png>)

* Going back to the IDE, you should see a confirmation that the IDE has been successfully connected and the API token has been securely stored.

![Confirmation of connection](<../../.gitbook/assets/Screenshot 2022-05-13 at 11.30.54.png>)

To verify that the Eclipse plugin is ready to start scanning, be sure the Snyk preferences show the following after the downloads and the authentication are done:

* Language server path that the Eclipse plugin will use
* Snyk CLI path that the Eclipse plugin will use
* Snyk API token, securely stored through the Eclipse's secure storage mechanism

![Snyk preferences after downloads and authentication](<../../.gitbook/assets/Screenshot 2022-05-17 at 16.36.07.png>)

### Configure the API token manually

You can provide the API token by copying it from your [account settings](https://app.snyk.io/account) and paste it into the Eclipse preferences Snyk API Token field. Click **Apply and Close** once the token has been set\*\*.\*\*

![Providing the API token manually](<../../.gitbook/assets/Screenshot 2022-05-17 at 16.36.07.png>)
