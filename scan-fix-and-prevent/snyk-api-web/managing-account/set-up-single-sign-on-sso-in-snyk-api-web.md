# Set up Single Sign-On (SSO) in Snyk API & Web

Learn how to configure your company's Single Sign-On to log in to the Snyk API & Web app.

Snyk API & Web uses the standard Security Assertion Markup Language (SAML) to exchange authentication and authorization information with an Identity Provider (IdP) to enable Single Sign-On (SSO) to the Snyk API & Web app. This means that if you have your company's SSO in place for your users to log in to your applications, you can also enable SSO to access the Snyk API & Web app.

When you do this, users from your account can choose the following option from the login screen to access Snyk API & Web:

<figure><img src="../../../.gitbook/assets/managing-account-set-up-single-sign-on-sso-login-option.png" alt="Log in with SSO button on login screen"></figure>

Note:

* If you already use SSO to log in to Snyk, you can log in to Snyk API & Web with your existing Snyk account, using the "Log in with Snyk" button:

<figure><img src="../../../.gitbook/assets/managing-account-set-up-single-sign-on-sso-login-with-snyk.png" alt="Log in with Snyk button"></figure>

* If you already have a Snyk account but you do not use SSO, you can [set up SSO with Snyk](https://docs.snyk.io/enterprise-setup/single-sign-on-sso-for-authentication-to-snyk).

Learn more about [Log in to Snyk API & Web](log-in-to-snyk-api-web.md).

---

This configuration of Snyk API & Web's specific SSO involves two steps:

1. Configure Snyk API & Web in your Identity Provider.
1. Configure SSO in Snyk API & Web.

This article describes these steps in detail.

Once you complete this set up, you can choose the following option from the login screen to access your account:

<figure><img src="../../../.gitbook/assets/managing-account-set-up-single-sign-on-sso-company-login.png" alt="Log in with company SSO option"></figure>

## Step 1: Configure Snyk API & Web in your Identity Provider

In this first step, go to your Identity Provider and create an entry for Snyk API & Web using the following information:

* **Entity Identifier** - The URL that identifies Snyk API & Web as the issuer of SAML requests, responses, or assertions: `https://probely.com`.
* **Assertion Consumer Service** - The Snyk API & Web endpoint to do the SAML authentication and authorization: `https://sso.plus.probely.app/sso/<organization-id>/complete/`

In the endpoint, replace `<organization-id>` with a string that identifies your organization (with lowercase letters and hyphens only). For example, the company name, but if you need any help, we can suggest it for you.

* **Certificate** - The SAML certificate for Snyk API & Web:

```
MIIFjzCCA3egAwIBAgIUBjrMlHlE8dKutYm0cz0JXFIjMMQwDQYJKoZIhvcNAQELBQAwVzELMAkGA1UEBhMCUFQxDzANBgNVBAgMBkxpc2JvbjEPMA0GA1UEBwwGTGlzYm9uMRAwDgYDVQQKDAdQcm9iZWx5MRQwEgYDVQQDDAtwcm9iZWx5LmNvbTAeFw0yMTAxMTExNTEyMDBaFw0zMTAxMTExNTEyMDBaMFcxCzAJBgNVBAYTAlBUMQ8wDQYDVQQIDAZMaXNib24xDzANBgNVBAcMBkxpc2JvbjEQMA4GA1UECgwHUHJvYmVseTEUMBIGA1UEAwwLcHJvYmVseS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDBIuwIGTIDaKPvv8lRKNlQVRD/lZgUeeSifEGl81JywWIDKWqvrXv07dupDoEAodwX9mfl68RfrBi4SfB1FYijEl7axURJgUrijCYolF76zxhcSVwadlzR9uOQZu/y//mwEZZtuYk62s1XM6o94pomR9Iv+ODoD9QgIARJEXTo/c0a55pt1Ps5yepUAHrr2JQaKyzgqrYeh3wXjKUf74MIWo3HNeRK16q5GSEdWs5bzjhdEynm8PCsAJ8kLqn/JawLgbhnOKVtgHf/DvywgVR+r3ZdFlqzKQOIiTD0y0n+Niy8jVI50xPT0Sf4H+jzW1VUhKuU2hbuPjUbPAoIiCDm0KquUfGJNkmPO8wfFwSN9HJLGBS8J56IvDf6yib5u111ddVqY6nfq/lMydrRlahV1ifouiep9vNsG9pbDy63biXVyyZD+0M8vom3AoeV5Vi6WM6OitrmzjevPy2XdPzCmmPFsUiQaAMNvgrYlE5OmeAyRQiBAlBlzDWBdRQBHmzVfcwB4TIdFcwBA2hkCDOkX/StrPb66YHMd4l6BIfGNgpGXyUDeDw8sV0R9Z8UXAHZ7C/CkgCwiLXGXpKMYcXLW8+51TC+SFksBGq6xSv4rBTTIP/+BewWyrG4F9E08VY0wEwa7fS19Ub/sC5s0TZxc9udLODlPhXKqvUC2qfGBQIDAQABo1MwUTAdBgNVHQ4EFgQUZAuqKmA9OaaqrLYoDtKVPtCXA3MwHwYDVR0jBBgwFoAUZAuqKmA9OaaqrLYoDtKVPtCXA3MwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAgEAZsze5LAdBkRBLmEK57E1Wp+3rhLIPjF49/riQcQZPhGV2ED4elMWsGbBTzPAeInq1m8FqMPCuAAU9Sbv+WmjzcJeDY2Re0jelmkSpiGejpRMdYhmoH3WOvHo1SfqZrlTowKKJ9qVp/nqMRMMxxjwvKvcjRjzdANl1oxitT6MjF7ISIQSHMi/sAXvYKprgcCdr+10bY+hQuK/r952Kd+YetTbfvS3B5UvhtJwCl1g3SUZCQ8n0Erxsg5rME2d10Ykva6UNlTeEEBsVWkvgkZTYzVv06T+BBIo8XX/FuwTnFj2Goej+CGxCntuFR+JMwuSX3eZKSVts3S7/ytGDFE3DKREN/xexARTICbl4+lCpvrdYIZHUR4YXTCAuw9OpYilWF/A/KlJb2aZWHI/RCDG2I7mwvA+Q9q+W5WNhBhYgeqvmF5bj0hDdVG0CpoGQEe34xasd0+EgoNlxJ565+fgXYdPdOCwXau1epRSYNNdWgjQFVs5oGhA0dbtfYQTjMN7WthvaXJNFofN4BvrGsEufIJwHTYbeHizbGu6dUxLCCLtVYVOHoF+EmL17djSU5PIT3bshDg2qmC/uJS/HTdr5NRybCyy3F8c5P7oDeVTbcqBRMY2e0ZxjJMMrwehU2q+gZovNRVTybplQXuvGzg3r0yuhRjQ1dWhKAZkEcV7B/Y=
```

## Step 2: Configure SSO in Snyk API & Web

With Snyk API & Web configured in your Identity Provider, the second part of the SSO configuration is on the Snyk API & Web side. For that, you need to provide the following information:

* Your Entity Identifier.
* Your Certificate.
* The URL of your SSO or the URL with the metadata.
* The SAML claims with information about:
  * The first name.
  * The last name.
  * The email.

You can also map your SAML Groups to [Snyk API & Web Roles](roles-and-permissions.md) and specific scopes (global to the account, a [team](get-started-with-teams.md), or a target). In this case, you must provide Snyk API & Web with the SAML claim with the information about the SAML Groups and tell Snyk API & Web how the mapping is done. Here is an example:

| SAML Group | Snyk API & Web Role | Snyk API & Web Scope |
| --- | --- | --- |
| probely\_admin | Admin | Global (account) |
| teamX\_admin | Admin | Team X |
| teamX\_developers | Developer | Team X |
| portal\_developers | Developer | Portal Target |

This mapping would produce the following results:

* Users belonging to **probely\_admin** would be given **Admin** permissions to the whole Snyk API & Web account (global scope). They could view and take action on any target of your account.
* Users belonging to groups **teamX\_admin** and **teamX\_developers** would only perform actions on targets of **Team X**, with permissions to do what the respective **Admin** and **Developer** roles allow.
* Users belonging to **portal\_developers** would have the permissions given by the **Developer** role and would only perform actions on a single target: the **Portal Target**.
