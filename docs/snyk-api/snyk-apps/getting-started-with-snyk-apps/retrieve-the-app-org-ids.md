# Retrieve the App Org IDs

Users may connect with a single organization., or a single group. Most of Snyk's API endpoints require an orgid in the path, which is used for authorizing the action being performed. For more information see [Integrating Apps with Snyk](https://docs.snyk.io/integrations/snyk-apps#integrating-apps-with-snyk).

To retrieve the orgid that is used by your App, send a GET request to the `orgs` endpoint:

```
https://api.snyk.io/rest/orgs?version={version}
```

{% hint style="danger" %}
Snyk deprecated the `/rest/apps/{clientId}/orgs?version={version}` endpoint as of 2022-03-02, and sunset it after 2022-04-03.
{% endhint %}

The current **version** can be found in the Snyk [API documentation](https://apidocs.snyk.io/#get-/orgs).

Snyk recommends you store this value and associate it with the user's details.
