# Snyk for Government (US)

[Snyk for Government (US)](https://snyk.io/government-security-solution/) enables US federal agencies to develop fast and securely. By integrating with tools and workflows that developers already use, Snyk for Government (US) allows these agencies to shift left in their Software Development Lifecycle, enabling secure development from the start.

Because Snyk for Government (US) adheres to the [FedRAMP](https://www.fedramp.gov/) and [NIST](https://www.nist.gov/) security control requirements, federal agencies can be assured that the product complies with the security standards set forth by the US Government.

Snyk for Government (US) has differences from standard Snyk products that allow Snyk to be deployed to the US federal government. Adhering to FedRAMP and NIST control requirements means that some aspects of standard Snyk products are not supported in the FedRAMP environment.

This list identifies all the areas with differences in functionality in the Snyk for Government (US) product.

## Core products limitations on availability

* Snyk Code is available except:
  * Does not include Code Search
  * Does not include Snyk Agent Fix
* Snyk Container is available except for [Kubernetes Integration](../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/).
* Snyk Open Source is available except:
  * Does not include Unmanaged C++
  * Does not include the npm packages `@snyk/protect` and `@snyk/fix.`
* Snyk AppRisk is **not available**.

## API keys not available

API keys are not available.

This means that attempts to create Service Accounts either through the UI or using the API with an `auth_type` of `api_key` are not accepted. You must ensure that the OAuth protocol is used instead for all scenarios where API keys would typically be used. See [Service accounts using OAuth 2.0](../enterprise-setup/service-accounts/service-accounts-using-oauth-2.0.md) for details. If you need help, contact [Snyk support](https://support.snyk.io).

In addition, the CLI must be used in OAuth mode, not with token-driven authentication.

## Single Sign-On limitations on availability

[Single Sign-On (SSO)](../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/) is available except for [Self-Serve Single Sign-On (SSO)](../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/configure-self-serve-single-sign-on-sso/). All SSO setups are managed by Snyk. Some setup steps for Single Sign-On are slightly different:

* The service provider is Okta rather than Auth0.
* The ACS URL and Entity ID and certificate will be different per connection and thus will not match the Snyk Single Sign-On documentation.
* To get the ACS Url, Entity ID, and cert, Snyk will need to part-provision the connection in Okta.

See [Single Sign-On (SSO) for authentication to Snyk](../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/) for details.

## Integrations not available

[Gatekeeper plugins](../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/) are **not available** as they support OAuth authentication.

## Reporting and data not available

* [Legacy reporting](../manage-issues/reporting/legacy-reports/)
* [Issues Analytics](../manage-risk/analytics/issues-analytics.md)
* [Insights](../manage-risk/prioritize-issues-for-fixing/using-the-issues-ui-with-snyk-apprisk/)

## Platform features not available

* [Bitbucket Cloud App](../developer-tools/scms/organization-level-integrations/bitbucket-cloud-app.md). Note that [Bitbucket Data Center/Server](../developer-tools/scms/organization-level-integrations/bitbucket-data-center-server.md) integration is available.
* [Slack App](../integrate-with-snyk/jira-and-slack-integrations/slack-app.md)
* [Jira App](../integrate-with-snyk/jira-and-slack-integrations/snyk-security-in-jira-cloud-integration.md). Note that [Jira Integration](../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md) is available.
* [Snyk Advisor](https://snyk.io/advisor/)
* [Snyk Learn](https://learn.snyk.io/?)
* Social logins: Google, GitHub, and so on as identity provider
* SSO into [Snyk support](https://support.snyk.io)
* [Status page](https://status.snyk.io)
* Outbound webhooks
* Session Concurrency is limited to three (3) sessions per user.
* Session lockout: After sessions expire, the signed-in user loses access to all data present in existing session windows.
* Session timeout: The default session timeout is shorter (15 minutes). See [Configure session length for a Snyk Group](../snyk-admin/groups-and-organizations/groups/configure-session-length-for-a-snyk-group.md) for details.
* [Snyk CLI docker images](../cli-ide-and-ci-cd-integrations/snyk-cli/install-or-update-the-snyk-cli/#snyk-cli-in-a-docker-image). These do not support FIPS-validated cryptography and should only be used when this can be accepted.
