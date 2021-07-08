# Authentication for API

#### [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

### [ ](untitled-137.md) <a id="category-name"></a>

#### Getting started with the API

* [ API documentation](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360007584578-API-documentation/README.md)
* [ Authentication for API](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360004037557-Authentication-for-API/README.md)
* [ Revoking and regenerating Snyk API tokens](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360004008278-Revoking-and-regenerating-Snyk-API-tokens/README.md)
* [Docs Library \| Snyk](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/README.md)
* [Snyk API](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/categories/360000665657-Snyk-API/README.md)
* [Getting started with the API](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/sections/360002233417-Getting-started-with-the-API/README.md)

## Authentication for API

To use our API, you must first authenticate your Snyk account with an API token.

1. Visit [your Snyk account](https://app.snyk.io/account).
2. Navigate to **General Account Settings** and copy your token.
3. From the token field, click **click to show** and then select and copy your API token.
4. To use our API, use an Authorization header, preceded by the token:

   ```text
   Authorization: token API_KEY
   ```

5. 
