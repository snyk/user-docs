# Authenticate for the API

To use the Snyk API, you must get your API token from Snyk. You can find your token in your personal [General Account Settings](https://app.snyk.io/account) after you register with Snyk and log in. In the **key** field, **Click to show**. Then, highlight and copy the API key.

If you want a new API token, select **Revoke & Regenerate.** This will make the previous API token invalid. For details, see [Revoke and regenerate a Snyk API token](revoke-and-regenerate-a-snyk-api-token.md).

When using the API directly, provide the API token in an `Authorization` header, as in the following example request, replacing `API_TOKEN` with your API Token

```bash
curl --request GET \
--url "https://api.snyk.io/rest/self?version=2024-06-10" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: token API_TOKEN"
```

If you are using the API through [Snyk Apps](../../how-to-use-snyk-apps-apis/), provide the `access_token` in an `Authorization` header preceded by `bearer` as follows:

```
Authorization: bearer ACCESS_TOKEN
```

Otherwise, a `401 Unauthorized` the response will be returned:

```http
HTTP/1.1 401 Unauthorized

{
    "status": "401",
    "code": "Unauthorized"
}
```

For information on when to use an API token and when to use a service account token, see [Authentication for API](./).
