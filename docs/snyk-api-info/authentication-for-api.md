# Authentication for API

To use our API, you must first authenticate your Snyk account with an API token.

1. Visit [your Snyk account](https://app.snyk.io/account).
2. Navigate to **General Account Settings** and copy your token.
3. From the token field, click **click to show** and then select and copy your API token. \[![api token screen; revoke; regenerate; click to show](../.gitbook/assets/uuid-8d94edf8-b42b-e5b3-ada1-e157d18ff884-en.png)

\]\([https://support.snyk.io/hc/article\_attachments/360006930518/uuid-8d94edf8-b42b-e5b3-ada1-e157d18ff884-en.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](https://support.snyk.io/hc/article_attachments/360006930518/uuid-8d94edf8-b42b-e5b3-ada1-e157d18ff884-en.png)\)

1. To use our API, use an Authorization header, preceded by the token:

   ```text
   Authorization: token API_KEY
   ```

