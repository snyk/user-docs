# Authentication for the Eclipse plugin

After the CLI is downloaded, you are redirected to the browser to authenticate and then connect the IDE with your account.

## Steps in authentication

You get a notification that a browser window will open (lower right).

<figure><img src="../../../.gitbook/assets/Screenshot 2022-05-13 at 11.29.37.png" alt="Notification browser window will open for authentication"><figcaption><p>Notification, browser window opening for authentication</p></figcaption></figure>

* When you are redirected to the browser for authentication, click **Authenticate**.

<figure><img src="../../../.gitbook/assets/authenticate-23-06.png" alt="Authenticate"><figcaption><p>Authenticate</p></figcaption></figure>

You should see a confirmation message that authentication is complete and you can now use Snyk.

<figure><img src="../../../.gitbook/assets/authentication-complete-23-06.png" alt="Authenticated"><figcaption><p>Authenticated</p></figcaption></figure>

Return to the IDE. You should see a confirmation that the IDE has been successfully connected and the API token has been securely stored (lower right).

<figure><img src="../../../.gitbook/assets/Screenshot 2022-05-13 at 11.30.54.png" alt="Confirmation of connection"><figcaption><p>Confirmation of connection</p></figcaption></figure>

To verify that the Eclipse plugin is ready to start scanning, be sure the Snyk preferences show the following after the downloads and the authentication are done:

* Snyk CLI path that the Eclipse plugin will use
* Snyk API token, securely stored through the Eclipse's secure storage mechanism

<figure><img src="../../../.gitbook/assets/image (347).png" alt="CLI path and API token in preferences"><figcaption><p>CLI path and API toekn in preferences</p></figcaption></figure>

## Configure the API token manually

You can provide the API token by copying it from your [account settings](https://app.snyk.io/account) and paste it into the Eclipse preferences Snyk API Token field. Click **Apply and Close** once the token has been set.

<figure><img src="../../../.gitbook/assets/image (1) (6).png" alt="Set the API token"><figcaption><p>Set the API token</p></figcaption></figure>

## Next steps

Change the [configuration settings](https://docs.snyk.io/ide-tools/eclipse-plugin/configuration-of-the-eclipse-plugin) as needed and set the [environment variables](https://docs.snyk.io/ide-tools/eclipse-plugin/environment-variables-for-the-eclipse-plugin).
