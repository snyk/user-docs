# Snyk Apps

{% hint style="info" %}
Only an App’s developer (Snyk Organization) can install the Snyk App they created. Snyk wants to learn more about what the community needs before launching the Snyk App marketplace.
{% endhint %}

Snyk Apps are integrations that extend the functionality of the Snyk platform, allowing you to create a Snyk experience to suit your specific needs. For example, a Snyk App might automate Snyk’s [application security testing](https://snyk.io/learn/application-security/testing/) as part of a build tool. Another Snyk App might stream Snyk’s security testing results into an incident management tool.

The easiest way to start building a new Snyk App is to clone the [Demo Snyk Apps ](https://github.com/snyk/snyk-apps-demo)GitHub repository. See [Quickstart: clone the Snyk demo app](./#quickstart-clone-the-snyk-demo-app).

See the [Snyk Apps blog post](https://snyk.io/blog/snyk-apps-beta-build-custom-apps-extend-snyk-security-into-workflows/) for more details.

This document describes how to create a new Snyk App in the language of your choice, and can be used along with the [REST API documentation](https://apidocs.snyk.io) and [Snyk OAuth2 API documentation](https://snykoauth2.docs.apiary.io/).

Snyk Apps are based on the [Snyk API](../../snyk-api/) so that your integrations are inherently stable and secure.

## Integrating Apps with Snyk

The Snyk Apps platform uses an OAuth 2.0 authorization flow. This allows your Snyk App to get an access token to act on a users’ behalf, depending on the scopes you request. There are many OAuth 2.0 libraries available that will greatly simplify the integration. The [Snyk Apps Demo](https://github.com/snyk/snyk-apps-demo) uses the popular JavaScript library [passport.js](http://www.passportjs.org/packages/passport-oauth2/).

See the Snyk [OAuth2 API documentation](https://snykoauth2.docs.apiary.io) for details.

## How do Apps connect?

A Snyk App is set up to use a specific 'context' when creating the Snyk App, either 'tenant' or 'user'. This can be set with the `context` field if creating a Snyk App with the API or the `--context` flag if using the CLI. If not specified, a Snyk App will default to using the 'tenant' context, which should be preferred unless your Snyk App specifically needs to perform operations as an individual user.

### User context

Authorizing a Snyk App that uses the user context will grant the Snyk App access to perform actions on behalf of the user. The Snyk App will have access to the same set of Organizations and Groups as the installing user, as well as access to any new Organizations and Groups the user is added to. If at any point the installing user deactivates their account, any installed apps using the 'user' context that user has installed will not be able to request new access tokens.

### Tenant context

When a user authorizes a Snyk App that uses the tenant context, the user selects the way to connect to either an Organization (Org) or a Group, with access to either all or one of the Orgs in the Group. The resulting connection is no longer tied to that specific user, so it is not bound by the user’s access or lifecycle but is linked to the Org’s lifecycle of the Org. For more information about Orgs, see [Introduction to Organizations](../../snyk-admin/manage-groups-and-organizations/whats-a-snyk-organization.md), and for Groups, see [Introduction to Snyk Groups](../../snyk-admin/manage-groups-and-organizations/introduction-to-groups.md).

The `scopes` of an App determine which actions an App can perform while connected to a Snyk account. Currently, `scopes` are related to actions performed through Snyk Orgs, which determine the API endpoints that can be used in an App. Endpoints authorized through a Snyk Group are not yet supported. For more information about scopes, see [Requesting scopes](https://docs.snyk.io/snyk-apps/getting-started-with-snyk-apps/create-an-app-via-the-api#requesting-scopes).

The final step of the integration process is to [get the ID of the Snyk organization the user has authorized](getting-started-with-snyk-apps/retrieve-the-app-org-ids.md). This ID is needed to use most API endpoints, so the App must associate it with your user's account or workspace for future use.

## Quickstart: clone the Snyk demo app

The easiest way to start building a new Snyk App is to clone the [Demo Snyk Apps](https://github.com/snyk/snyk-apps-demo) Github repository. You can either modify the demo to suit your App design or use it as a base for creating your own Snyk App.

See [Getting started with Snyk Apps](getting-started-with-snyk-apps/) for more details.
