# Authenticate the CLI with your account

To associate your Snyk account with the CLI, you must first authenticate your machine \(except for using the `snyk protect` command\). No repository permissions are needed at this stage, only your email address.

You can authenticate:

* Through your browser
* Using your API token

### Authenticate using your API token

1. Visit [your Snyk account](https://app.snyk.io/account) \(**Account Settings &gt; API Token** section\).
2. In the **KEY** field, click **click to show**, then select and copy your API token.
3. In the CLI, run `snyk config set api=XXXXXXXX`

