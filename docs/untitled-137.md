# Authentication for API

####  [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

###  [ ]() <a id="category-name"></a>

#### Getting started with the API

* [ API documentation](/hc/en-us/articles/360007584578-API-documentation)
* [ Authentication for API](/hc/en-us/articles/360004037557-Authentication-for-API)
* [ Revoking and regenerating Snyk API tokens](/hc/en-us/articles/360004008278-Revoking-and-regenerating-Snyk-API-tokens)

1.  [Docs Library \| Snyk](/hc/en-us)
2.  [Snyk API](/hc/en-us/categories/360000665657-Snyk-API)
3.  [Getting started with the API](/hc/en-us/sections/360002233417-Getting-started-with-the-API)

##  Authentication for API

To use our API, you must first authenticate your Snyk account with an API token.

1. Visit [your Snyk account](https://app.snyk.io/account).
2. Navigate to **General Account Settings** and copy your token.
3. From the token field, click **click to show** and then select and copy your API token.
4. To use our API, use an Authorization header, preceded by the token:

   ```text
   Authorization: token API_KEY
   ```

* 
