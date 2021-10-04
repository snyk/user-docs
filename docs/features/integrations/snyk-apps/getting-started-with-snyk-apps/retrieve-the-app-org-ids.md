# Retrieve the App Org IDs

To use the Snyk API in your App, you must know the Snyk Org ID that the user has authorized the App for.

To retrieve this information, send a GET request to the `orgs` endpoint:

```text
https://api.snyk.io/v3/apps/{clientId}/orgs?version={version}
```

The current **version** can be found in our [API documentation](https://snyk.docs.apiary.io/#reference/apps/app-org-access/getappaccessorgs).

We recommend you store this value and associate it with the user's details.

{% hint style="info" %}
For the beta release of Snyk Apps, users may only install an App on a single Organization per authorization request.
{% endhint %}

