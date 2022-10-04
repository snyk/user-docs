# Identity Provider (IdP) migration

When migrating from a legacy IdP to a new IdP, new IdP metadata information must be submitted to Snyk.

To migrate identity providers:

1. Download the correct Worksheet from the [SSO resources](set-up-snyk-single-sign-on-sso.md#resources).
2. Fill out Worksheet with the IdP metadata information
3. Submit Worksheet ([contact our Support team](https://support.snyk.io/hc/en-us/requests/new) to raise a Support ticket)

To prevent new users from being created within Snyk, you’ll need to maintain your SAML protocol and use both the same Entity ID and ACS URL. If you are changing SAML protocols, [contact our Support team](https://support.snyk.io/hc/en-us/requests/new).

After this is done, the Support team will contact you and confirm the updated metadata has granted access through the new IdP.
