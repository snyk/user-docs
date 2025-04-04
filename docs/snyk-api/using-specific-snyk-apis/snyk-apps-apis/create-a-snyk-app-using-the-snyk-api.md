# Create a Snyk App using the Snyk API

When you have an API token and `orgId`, you can use the Snyk API to create your Snyk App by sending a `POST` request to the endpoint [Create a new Snyk App for an organization](https://docs.snyk.io/snyk-api/reference/apps#orgs-org_id-apps-creations):

```
https://api.snyk.io/rest/orgs/{orgId}/apps/creations?version={version}
```

Example CURL request to create a Snyk App:

```
curl -XPOST -H"Content-Type: application/vnd.api+json" \
 -H"Authorization: token <REPLACE_WITH_API_TOKEN>" \
 -d '{"data": { "attributes": {"name": "My Awesome Snyk App", "redirect_uris": ["https://example.com/callback"], "scopes": ["org.read"], "context": "tenant"}, "type": "app"}}' \
 "https://api.snyk.io/rest/orgs/<REPLACE_WITH_YOUR_ORGID>/apps/creations?version=2024-10-11"
```

The request body should contain the details for your new App, including the `name`, `redirectURIs`, and [`scopes`](scopes-to-request.md).

The response includes details necessary to complete the integration: `clientId` and `clientSecret`. Use these values with Snyk API endpoints within your App; consider storing them as part of the configuration of your App.

{% hint style="info" %}
Never share the `clientSecret` publicly, as it is used to authenticate your App. This is the only time you will see the `clientSecret`, so keep it secure and private. If you lose it or if the secret is leaked, you can [rotate your App's clientSecret](manage-app-details.md#rotate-app-clientsecret).
{% endhint %}
