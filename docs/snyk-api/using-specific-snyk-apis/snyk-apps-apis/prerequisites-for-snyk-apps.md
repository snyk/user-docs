# Prerequisites for Snyk Apps

To create a Snyk App, you must have access to the Snyk API. To get started, follow the instructions to [authenticate for the API](../../authentication-for-api/).

You must also retrieve the ID of the Snyk Organization you intend the App to be owned by (your `orgId`). You can get the Organization ID from the Organization settings in the Snyk Web UI or by using the [List accessible organizations](https://docs.snyk.io/snyk-api/reference/orgs#orgs) endpoint (`https://api.snyk.io/rest/orgs` ) with the Snyk API token in the Authorization header.

{% hint style="info" %}
Snyk Apps have first-class access to the API, regardless of whether users installing the App have paid for access or not. To take advantage of this feature, Apps must use API endpoints with the domain https://api.snyk.io/ rather than the deprecated `https://snyk.io/api/` to call the API within the App.
{% endhint %}
