# Configure authentication

Authentication tells agents who to log in as. Most of an application's interesting attack surface sits behind a login, and the most damaging vulnerabilities only become visible when agents can compare what one user can reach against what another user should be able to reach.

## Configure more than one user

You can add any number of users to a target. Add at least two, with different roles.

Agents use multiple identities to test whether the application enforces its own authorization rules. With a single user there is nothing to compare against, so this entire class of vulnerability goes untested:

| Vulnerability class                          | Why multiple users are required                                                                              |
| -------------------------------------------- | ------------------------------------------------------------------------------------------------------------ |
| **BOLA** (Broken Object Level Authorization) | Agents need user A's object identifiers to attempt reaching them as user B                                   |
| **Privilege escalation**                     | Agents need a low-privilege identity to test whether it can perform high-privilege actions                   |
| **Broken authorization**                     | Agents need to know what a role should not be able to do, which requires seeing what a different role can do |

{% hint style="info" %}
A useful minimum is one low-privilege user and one administrator. If your application has tenants, add users in different tenants as well, so agents can test for cross-tenant data access.
{% endhint %}

Give each user a label describing its role, for example `Standard user` or `Admin user`. Labels are how you tell configured users apart later, and they give agents a hint about the privilege level each one represents.

Use dedicated test accounts. Agents will attempt real attacks with these credentials, including actions that modify data.

## The login URL

You can specify the URL of your login page. If you omit it, it is detected automatically.

Set it explicitly when your login page is somewhere agents are unlikely to find on their own, or when several login pages exist and you want a specific one used.

## Authentication types

Each user is configured with one of two types.

### Credentials login

Use this when the login is a form. You provide any number of key-value pairs, so the configuration adapts to whatever fields your form actually has rather than assuming a fixed username and password.

Typical fields:

| Field                  | Notes                                                                                                                |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------- |
| Username or email      | Whatever your login form uses as the identifier                                                                      |
| Password               | The account password                                                                                                 |
| Additional form fields | Any other input the form requires, such as a tenant identifier, an organization slug, a domain, or a region selector |
| API key                | Where a user authenticates with a key rather than, or in addition to, a password                                     |

Values holding secrets are marked as sensitive, which keeps them out of logs and output.

**Two-factor authentication**

For accounts protected by two-factor authentication, provide the TOTP secret rather than a code. Codes expire; the secret lets agents generate a valid code whenever one is needed, including mid-scan if the session has to be re-established.

Use the secret your authenticator app was originally enrolled with, which is the value behind the QR code you scanned during enrollment.

**Adding hints to a credentials login**

A credentials login can also carry free-text instructions. Use these for small obstacles around an otherwise ordinary form, without moving to a fully custom login:

> Accept the cookie banner on first visit. Select the EU region if prompted.

This is the right tool when the form itself is standard but something in front of it needs handling.

### Custom login

Use this when the login flow cannot be expressed as a set of form fields. Instead of naming fields, you describe the flow in natural language and the agents follow your description.

Custom logins suit flows such as:

* Single sign-on through an external identity provider.
* Multi-page logins where the identifier and the password are on separate screens.
* Logins requiring a tenant selection or an interstitial before the credential prompt.
* Flows completed with an out-of-band step, such as accepting a push notification.
* Any flow where the correct sequence of actions matters more than the values entered.

Write the instructions as you would explain the login to a new colleague. Be specific about the order of steps, and name elements directly where you can:

> Navigate to /login, click "Sign in with SSO", enter the email in #sso-email, click Continue, complete the identity provider login, then accept the MFA push notification.

{% hint style="info" %}
If your login uses a third-party identity provider, the provider's domain must be in the target scope. Otherwise the login redirect is rejected by the proxy and authentication fails, whichever type you chose. See Define the target scope.
{% endhint %}

## Choosing between the two

| Situation                                             | Type                                                          |
| ----------------------------------------------------- | ------------------------------------------------------------- |
| Single-page form, fields you can name                 | Credentials login                                             |
| Form with unusual extra fields                        | Credentials login, adding the extra fields as key-value pairs |
| Two-factor authentication on a standard form          | Credentials login, with the TOTP secret                       |
| Standard form behind a cookie banner or region prompt | Credentials login, with instructions as hints                 |
| Single sign-on                                        | Custom login                                                  |
| Multi-step or multi-page login                        | Custom login                                                  |
| Flow requiring interaction beyond filling fields      | Custom login                                                  |

Prefer a credentials login where one works, because supplying field values is more deterministic than having agents interpret a description. Move to a custom login when the flow needs judgment rather than values.

## When authentication is not needed

If the application has no login, or you only want to test the unauthenticated surface, you can leave authentication unconfigured. The assessment then covers only what an anonymous visitor can reach, which for most applications is a small fraction of the attack surface.

## When authentication fails

Authentication failure is caught at scan start: a scan will not begin if the target is unreachable or authentication fails. That is a useful early signal, but it only detects failures visible before the scan runs.

A subtler failure is authentication that succeeds while reaching less than you expect. Check the recon numbers on the first scan: if the discovered surface is much smaller than your application, the likely causes in order are the identity provider not being in scope, the rest of the application not being in scope, or a login that completes without reaching the authenticated area.

See [Monitor and manage scans](../scans/monitor-and-manage-scans.md).
