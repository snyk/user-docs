# How can I delete my account?

* [ Introduction to Snyk Single Sign-On \(SSO\)](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360018025297-Introduction-to-Snyk-Single-Sign-On-SSO-/README.md)
* [ Choose a provisioning option](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360019607318-Choose-a-provisioning-option/README.md)
* [ Set up Snyk Single Sign-On \(SSO\)](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360017753618-Set-up-Snyk-Single-Sign-On-SSO-/README.md)
* [ Identity Provider \(IdP\) migration](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/4402292397969-Identity-Provider-IdP-migration/README.md)

## Introduction to Snyk Single Sign-On \(SSO\)

**Feature availability**  
This feature is available with Business and Enterprise customers. See [Pricing plans](https://snyk.io/plans/) for more details.

You can take advantage of your company's existing identity management systems, and have employees sign in to Snyk using their corporate identity. This makes provisioning users to Snyk easier. It also allows for deeper integration for group and organization membership, role-based access, and more.

Snyk can integrate with any SAML-based and OpenID Connect \(OIDC\)-based SSO, as well as ADFS. You can also use your Enterprise Identity Provider for SSO, including [Azure AD](https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-whatis) and [Google G Suite](https://community.snowflake.com/s/article/configuring-g-suite-as-an-identity-provider). Read more about SAML in [the Auth0 documentation](https://auth0.com/docs/protocols/saml).

### User authentication and provisioning

With SSO configured, users are provisioned with a new Snyk account when they first sign on through SSO, even if they previously created their own account.

The sign on process includes these steps:

1. When a user selects SSO from Snyk.io to log in, they are redirected to \(and authenticated\) by the identity provider you requested.
2. The identity provider communicates this authentication to Snyk servers, sending relevant data to Snyk in order to create each user.
3. Snyk checks the directory for that user.
4. If the user is already configured, Snyk enables the appropriate access. For a new user, Snyk adds the user to the directory, and then redirects them to Snyk.io with the appropriate access.

Read more about the next step, [choose a provisioning option](https://github.com/snyk/user-docs/tree/caef522cc2da817b75170d43049a1e6dd9d856fb/hc/en-us/articles/360019607318/README.md).

