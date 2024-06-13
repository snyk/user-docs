# Authentication for API

To use the Snyk API, you must first get your API token from Snyk.

1. Go to the **General Account Settings** in [your Snyk account](https://app.snyk.io/account).
2. In the KEY field, **click to show** and then select and copy your API token.
3. To use the Snyk API, supply the token in an `Authorization` header, preceded by `Token`:\
   `Authorization: Token API_KEY`

For more details and for information about using personal tokens versus service account tokens, see [How to obtain and authenticate with your Snyk API token](../../../getting-started/how-to-obtain-and-authenticate-with-your-snyk-api-token.md).

For information about using the token, refer to [About authentication](https://docs.snyk.io/snyk-api/snyk-rest-api-overview#about-authentication) in the REST API overview. For additional information, see [Revoking and regenerating Snyk API tokens](revoking-and-regenerating-snyk-api-tokens.md).
