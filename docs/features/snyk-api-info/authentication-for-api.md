# Authentication for API

To use our API, you must first authenticate your Snyk account with an API token.

1. Visit [your Snyk account](https://app.snyk.io/account).
2. Navigate to **General Account Settings** and copy your token.
3. From the token field, click **click to show** and then select and copy your API token.
4. To use our API, use an Authorization header, preceded by the token:

![api token screen; revoke; regenerate; click to show](<../../.gitbook/assets/uuid-8d94edf8-b42b-e5b3-ada1-e157d18ff884-en (1) (2) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (15).png>)

```
   Authorization: token API_KEY
```
