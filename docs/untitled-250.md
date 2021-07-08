# How to install and configure your Snyk Broker client

#### [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

### [ ](untitled-250.md) <a id="category-name"></a>

#### Snyk Broker

* [ Broker introduction](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015367178-Broker-introduction/README.md)
* [ Set up Snyk Broker](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015296637-Set-up-Snyk-Broker/README.md)
* [ Prepare Snyk Broker for deployment](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015296717-Prepare-Snyk-Broker-for-deployment/README.md)
* [ Enable permissions for Snyk Broker from your third-party tool](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015296737-Enable-permissions-for-Snyk-Broker-from-your-third-party-tool/README.md)
* [ Retrieve a unique Broker client token](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015367278-Retrieve-a-unique-Broker-client-token/README.md)
* [ How to install and configure your Snyk Broker client](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015296937-How-to-install-and-configure-your-Snyk-Broker-client/README.md)
* [ Upgrade your Snyk Broker client](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015367458-Upgrade-your-Snyk-Broker-client/README.md)
* [Docs Library \| Snyk](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/README.md)
* [Integrations](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/categories/360000598398-Integrations/README.md)
* [Snyk Broker](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/sections/360001138138-Snyk-Broker/README.md)

## How to install and configure your Snyk Broker client

Snyk’s recommended and supported method of running the Snyk Broker client is using Docker.

To install and configure your client:

1. Visit [the broker repository](https://github.com/snyk/broker) and follow the instructions for the relevant SCM to run one of our provided images, or derive your own.
2. For the environment variables required to run the Broker client, you must:
   1. 1. Retrieve your [unique broker client token,](https://support.snyk.io/hc/en-us/articles/360015367278-Retrieve-a-unique-Broker-client-token)
      2. Ensure that you have the necessary permissions granted to your API credentials for the integration you are using by following [one of our integration guides.](https://support.snyk.io/hc/en-us/categories/360000598398-Integrations)
3. After the broker is running, **identifying Broker server** appears towards the end of the log results, showing that a connection to Snyk has been established
4. Visit the settings for your brokered integration in [the Snyk application](https://app.snyk.io/) to see a message like the following: 
5. You can now use Snyk as normal.
6. 
