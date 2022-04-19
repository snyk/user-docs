# Snyk Apps

## Introduction

{% hint style="info" %}
Currently, Snyk Apps are in beta and only an App’s developer (Snyk Organization) can install the Snyk App they created -- you. We want to learn more about what the community needs before we launch the Snyk App marketplace.
{% endhint %}

Snyk Apps are integrations that extend the functionality of the Snyk platform, allowing you to create a Snyk experience to suit your specific needs. An example Snyk App might automate Snyk’s [application security testing](https://snyk.io/learn/application-security/testing/) as part of a build tool. Another example Snyk App might stream Snyk’s security testing results into an incident management tool.

The easiest way to start building a new Snyk App is to clone the [Demo Snyk Apps ](https://github.com/snyk/snyk-apps-demo)GitHub repository. You can either modify the demo to suit your App design or use it as a base for creating your own Snyk App.

See the [Snyk Apps blog post](https://snyk.io/blog/snyk-apps-beta-build-custom-apps-extend-snyk-security-into-workflows/) for more details.

This document describes how to create a new Snyk App in the language of your choice, and can be used along with the [API documentation](https://apidocs.snyk.io/#tag--Apps) and [Snyk OAuth2 API documentation](https://snykoauth2.docs.apiary.io).

Snyk Apps are based on the [Snyk API](../../snyk-api-info/), so that your integrations are inherently stable and secure.

### Integrating Apps with Snyk

The Snyk Apps platform uses an OAuth 2.0 authorization flow. This allows your Snyk App to get an access token to act on a users’ behalf, depending on the scopes you request. There are many OAuth 2.0 libraries available that will greatly simplify the integration. The [Snyk Apps Demo](https://github.com/snyk/snyk-apps-demo) uses the popular JavaScript library [passportjs](http://www.passportjs.org/packages/passport-oauth2/).

See the Snyk [OAuth2 API documentation](https://snykoauth2.docs.apiary.io) for details.

### How do Apps connect?

When a user authorizes a Snyk App, the user selects how to connect - to either an Organization (Org) or a Group (with access to either all or one of the Group's Orgs). The resulting connection is no longer tied to that specific user, so it is not bound by the user’s access or lifecycle but is linked to the Org’s lifecycle. You can read more about Orgs in [What’s a Snyk organization](https://docs.snyk.io/user-and-group-management/managing-groups-and-organizations/whats-a-snyk-organization) and Groups in [What's a Snyk Group](https://docs.snyk.io/features/user-and-group-management/managing-groups-and-organizations/whats-a-snyk-group).

The `scopes` of an App determine which actions an App can perform while connected to a Snyk account. Currently, `scopes` are related to actions performed through Snyk Orgs, which determine the API endpoints that can be used in an App. Endpoints authorized through a Snyk Group are not yet supported. You can read more about `scopes` in [Snyk Apps Scopes](snyk-apps-scopes.md).

The final step of the integration process is to [get the ID of the Snyk organization the user has authorized](getting-started-with-snyk-apps/retrieve-the-app-org-ids.md). This ID is needed to use most API endpoints, so the App must associate it with your user's account/workspace for future use.

### Quickstart: clone the Snyk demo app

The easiest way to start building a new Snyk App is to clone the [Demo Snyk Apps](https://github.com/snyk/snyk-apps-demo) Github repository. You can either modify the demo to suit your App design or use it as a base for creating your own Snyk App.

See [Getting started with Snyk Apps](getting-started-with-snyk-apps/) for more details.



