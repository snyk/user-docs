# Snyk Apps



{% hint style="info" %}
Currently, Snyk Apps are in beta and only an App’s developer (Snyk Organization) can install the Snyk App they created - you. We want to learn more about what the community needs before we launch our App marketplace.
{% endhint %}

Snyk Apps are integrations that extend the functionality of the Snyk platform. They provide you with an opportunity to mold your Snyk experience to suit your specific needs.

The easiest way to start building a new Snyk App is to clone our [Demo Snyk App](https://github.com/snyk/snyk-apps-demo) Github repository. You can either modify the demo to suit your App design or use it as a base for creating your own Snyk App.

This document describes how to create a new Snyk App in the language of your choice, and can be used along with the [API documentation](https://snykv3.docs.apiary.io/#reference/apps) and our [OAuth2 API documentation](https://snykoauth2.docs.apiary.io).

## Integrating Apps with Snyk

The Snyk Apps platform uses an OAuth 2.0 authorization flow. This allows your Snyk App to get an access token to act on a users’ behalf, depending on the scopes you request. There are many OAuth 2.0 libraries available that will greatly simplify the integration. Our [Snyk App Demo](https://github.com/snyk/snyk-apps-demo) uses the popular JavaScript library [passportjs](http://www.passportjs.org/packages/passport-oauth2/).

### How do Apps connect?

When a Snyk App [connects to a user’s account](getting-started-with-snyk-apps/set-up-to-authorize-users.md), the user selects which Organization (Org) to authorize with the App. The resulting connection is no longer tied to that specific user, so it is not bound by the user’s access or lifecycle but is linked to the Org’s lifecycle. You can read more on Orgs in [What’s a Snyk organization](https://docs.snyk.io/user-and-group-management/managing-groups-and-organizations/whats-a-snyk-organization).

{% hint style="info" %}
As an App is connected and authorized through a Snyk organization, this controls which API endpoints can be used in an App, and which actions can be performed.
{% endhint %}

The final step of the integration process is to [get the ID of the Snyk organization the user has authorized](getting-started-with-snyk-apps/retrieve-the-app-org-ids.md). This ID is needed to use most API endpoints and should be associated with the account/user/workspace on the App side. For example, endpoints that are authorized through a Snyk Group cannot be authorized at the Snyk Org level, and so are not supported in Apps.
