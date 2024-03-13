# Snyk Apps

Snyk Apps are integrations that extend the functionality of the Snyk platform, allowing you to create a Snyk experience to suit your specific needs.

For example, a Snyk App might automate Snyk’s [application security testing](https://snyk.io/learn/application-security/testing/) as part of a build tool. Another Snyk App might stream Snyk’s security testing results into an incident management tool.  You can create your own reporting, metrics, or even a user management App. You can move your automated Snyk API scripts and transform them into a Snyk App, which will add a better UX/UI experience for your Snyk Organization users.

The easiest way to start building a new Snyk App is to clone the [Demo Snyk Apps ](https://github.com/snyk/snyk-apps-demo)GitHub repository. You can either modify the demo to suit your App design or use it as a base for creating your own Snyk App. See the [Snyk Apps blog post](https://snyk.io/blog/snyk-apps-beta-build-custom-apps-extend-snyk-security-into-workflows/) for more details.

This document describes how to create a new Snyk App in the language of your choice and can be used along with the [REST API documentation](https://apidocs.snyk.io/#tag--Apps) and [Snyk OAuth2 API documentation](https://snykoauth2.docs.apiary.io/).

Snyk Apps are based on the [Snyk API](../../snyk-api/) so that your integrations are inherently stable and secure.

## Integrating Apps with Snyk

The Snyk Apps platform uses an OAuth 2.0 authorization flow. This allows your Snyk App to get an access token to act on behalf of the user, depending on the scopes you request. There are many OAuth 2.0 libraries available that will greatly simplify the integration. The [Snyk Apps Demo](https://github.com/snyk/snyk-apps-demo) uses the popular JavaScript library [passport.js](http://www.passportjs.org/packages/passport-oauth2/).

See the Snyk [OAuth2 API documentation](https://snykoauth2.docs.apiary.io) for details.

## How do Apps connect?

When a Snyk App is created, it is set up to use `a specific context`, either `tenant` or `user`. Set this with the `context` field if you are creating a Snyk App with the API or the `--context` flag if you are using the CLI. If the `context` is not specified, a Snyk App will default to the `tenant` context, which is preferred unless your Snyk App specifically needs to perform operations as an individual user.

### User context

Authorizing a Snyk App that has the `user` context grants the Snyk App access to perform actions on behalf of the user. The Snyk App will have access to the same set of Organizations and Groups as the installing user, as well as access to any new Snyk Organizations and Groups the user is added to. If the installing user deactivates their account, any installed apps having the `user` context that user has installed will not be able to request new access tokens.

### Tenant context

When a user authorizes a Snyk App that has the `tenant` context, the user selects the way to connect to either a Snyk Organization or a Group, with access to either all or one of the Organizations in the Group. The resulting connection is no longer tied to that specific user, so it is not bound by the user’s access or lifecycle but is linked to the lifecycle of the Organization. For more information, see [Organizations](../../snyk-admin/groups-and-organizations/organizations/) and  [Groups](../../snyk-admin/groups-and-organizations/groups/).

The `scopes` of an App determine which actions an App can perform while connected to a Snyk account. The`scopes` are related to actions performed through Snyk Organizations, which determine the API endpoints that can be used in an App. Endpoints authorized through a Snyk Group are not supported. For more information about scopes, see [Requesting scopes](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/-MdwVZ6HOZriajCf5nXH/\~/changes/6758/snyk-api/snyk-apps/create-an-app-via-the-api#requesting-scopes).

### Snyk Organization ID

The final step for the integration process is to [get the ID of the Snyk Organization the user has authorized](../../snyk-api/snyk-apps/set-up-and-use-a-snyk-app/retrieve-the-app-org-ids.md). This ID is needed to use most API endpoints, so the App must associate it with the user's account or workspace for future use.
