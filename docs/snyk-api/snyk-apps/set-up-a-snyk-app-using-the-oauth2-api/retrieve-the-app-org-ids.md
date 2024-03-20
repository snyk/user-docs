# Retrieve the App Org IDs

Users may connect with a single Organization or a single Group. Most of the Snyk API endpoints require an `orgid` in the path, which is used for authorizing the action being performed. For more information see [Integrating Apps with Snyk](../../../snyk-api-info/snyk-apps/#integrating-apps-with-snyk).

To retrieve the `orgid` that is used by your App, send a GET request to the `orgs` endpoint, [List accessible organizations](https://apidocs.snyk.io/#get-/orgs) at the following URL:

`https://api.snyk.io/rest/orgs?version={version}`

Snyk recommends you store this value and associate it with the user's details.
