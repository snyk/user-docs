# Snyk Apps

## Introduction

{% hint style="info" %}
Currently, Snyk Apps are in beta and only an App’s developer (Snyk Organization) can install the Snyk App they created.
{% endhint %}

Snyk Apps are integrations that extend the functionality of the Snyk platform, allowing you to create a Snyk experience to suit your specific needs. An example Snyk App might automate Snyk’s [application security testing](https://snyk.io/learn/application-security/testing/) as part of a build tool. Another example Snyk App might stream Snyk’s security testing results into an incident management tool.

See our [Snyk Apps blog post](https://snyk.io/blog/snyk-apps-beta-build-custom-apps-extend-snyk-security-into-workflows/) for more details.

Snyk Apps are based on the [Snyk API](../../snyk-api-info/), so that your integrations are inherently stable and secure.

### Integrating Apps with Snyk

The Snyk Apps platform uses an OAuth 2.0 authorization flow. This allows your Snyk App to get an access token to act on a users’ behalf, depending on the scopes you request. There are many OAuth 2.0 libraries available that will greatly simplify the integration. Our [Snyk App Demo](https://github.com/snyk/snyk-apps-demo) uses the popular JavaScript library [passportjs](http://www.passportjs.org/packages/passport-oauth2/).

See our [OAuth2 API documentation](https://snykoauth2.docs.apiary.io) for details.

#### How do Apps connect?

When a Snyk App [connects to a user’s account](getting-started-with-snyk-apps/set-up-to-authorize-users.md), the user selects which Organization (Org) to authorize with the App. The resulting connection is no longer tied to that specific user, so it is not bound by the user’s access or lifecycle but is linked to the Org’s lifecycle. You can read more on Orgs in [What’s a Snyk organization](https://docs.snyk.io/user-and-group-management/managing-groups-and-organizations/whats-a-snyk-organization).

{% hint style="info" %}
As an App is connected and authorized through a Snyk organization, this controls which API endpoints can be used in an App, and which actions can be performed.
{% endhint %}

The final step of the integration process is to [get the ID of the Snyk organization the user has authorized](getting-started-with-snyk-apps/retrieve-the-app-org-ids.md). This ID is needed to use most API endpoints and should be associated with the account/user/workspace on the App side. For example, endpoints that are authorized through a Snyk Group cannot be authorized at the Snyk Org level, and so are not supported in Apps.

### Quickstart: clone our demo app

The easiest way to start building a new Snyk App is to clone our [Demo Snyk App](https://github.com/snyk/snyk-apps-demo) Github repository. You can either modify the demo to suit your App design or use it as a base for creating your own Snyk App.

See [getting-started-with-snyk-apps](getting-started-with-snyk-apps/ "mention")for more details.



