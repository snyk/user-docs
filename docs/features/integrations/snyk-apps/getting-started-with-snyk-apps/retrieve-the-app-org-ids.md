# Retrieve the App Org IDs

To retrieve the App Org ID that's used by your App, send a GET request to the `orgs` endpoint:

```
https://api.snyk.io/rest/orgs?version={version}
```

{% hint style="danger" %}
Snyk deprecated the `/rest/apps/{clientId}/orgs?version={version}` endpoint as of 2022-03-02, and will sunset it after 2022-04-03.
{% endhint %}

The current **version** can be found in our [API documentation](https://apidocs.snyk.io).

We recommend you store this value and associate it with the user's details.

{% hint style="info" %}
For the beta release of Snyk Apps, users may only install an App on a single Organization per authorization request.
{% endhint %}
