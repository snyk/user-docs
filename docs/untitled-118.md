# Revoking and regenerating Snyk API tokens

* [ API documentation](/hc/en-us/articles/360007584578-API-documentation)
* [ Authentication for API](/hc/en-us/articles/360004037557-Authentication-for-API)
* [ Revoking and regenerating Snyk API tokens](/hc/en-us/articles/360004008278-Revoking-and-regenerating-Snyk-API-tokens)

##  Revoking and regenerating Snyk API tokens

When an API token is revoked, all integrations using that key will immediately stop working. Proceed with caution!

If you suspect an API token has been leaked, it's good practice to revoke it and generate a new one to use in its place.

To revoke your Snyk user API token, navigate to your User Account Settings at [app.snyk.io/account](https://app.snyk.io/account).

Click the **Revoke & Regenerate** button to revoke your API token. A new one will be generated in its place. You can now grab the newly generated API token and update integrations that used the old key.

