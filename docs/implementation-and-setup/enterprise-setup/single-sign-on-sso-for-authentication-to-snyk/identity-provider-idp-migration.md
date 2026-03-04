# Identity Provider (IdP) migration

When migrating from a legacy IdP to a new IdP, you must submit new IdP metadata information to Snyk.

To migrate identity providers:

1. Download the correct Worksheet from the [SSO resources](set-up-snyk-single-sign-on-sso.md#resources).
2. Fill out Worksheet with the IdP metadata information.
3. Submit the Worksheet; [contact Snyk Support](https://support.snyk.io) to raise a Support ticket.

To prevent new users from being created within Snyk, you must maintain your SAML protocol and use both the same Entity ID and ACS URL. If you are changing SAML protocols, [contact Snyk Support](https://support.snyk.io).

After you have done this, the Support team will contact you and confirm the updated metadata has granted access through the new IdP.
