# Snyk Broker - commit signing

{% hint style="info" %}
**Availability**

Snyk Broker commit signing is available only with Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans/).
{% endhint %}

As of version v4.151.0, the Snyk Broker Client supports commit signing for GitHub integration. With a brokered setup, you can sign GitHub commits for fix PRs with your GPG key and a dedicated user you have configured.

## Requirements for commit signing

* Broker Client version v4.151.0 or higher
* A GitHub account configured to sign commits with a GPG key properly configured under the **Access** > **SSH and GPG keys** section

## Configuration of commit signing

1. To use commit signing, provide the following environment variables for the Broker Client:
   * `GPG_PRIVATE_KEY`: GPG private key exported as an ASCII armored version. Note that the value must start with `-----BEGIN PGP PRIVATE KEY BLOCK-----` and end with `-----END PGP PRIVATE KEY BLOCK-----`.
   * `GPG_PASSPHRASE`: Passphrase of the GPG private key.
   * `GIT_COMMITTER_NAME`: will be used to set a committer name.
   * `GIT_COMMITTER_EMAIL`: will be used to set a committer email.
2. Enable **Broker Client Commit Signing** in the Broker settings.

## Troubleshooting commit signing

If commits are shown as `Unverified` in GitHub, you must do the following:

* Verify that the GPG public key is imported to the correct GitHub user and that the email address is the same in GitHub as in the environment variables.
* To ensure that the commit signing feature has been activated for the Broker Client in the Snyk Organization, check the logs and verify that the following message appears when the Broker Client starts up: ​​`loading commit signing rules (enabled=true, rulesCount=5)`
