# Retrieve the App Org IDs

Users may connect with a single Organization or a single Group. Most of the Snyk API endpoints require an `orgid` in the path, which is used for authorizing the action being performed.

To retrieve the `orgid` that is used by your App, send a GET request to the `orgs` endpoint, [List accessible organizations](../../../reference/orgs.md#get-orgs) at the following URL:

`https://api.snyk.io/rest/orgs?version={version}`

Snyk recommends you store this value and associate it with the user's details.
