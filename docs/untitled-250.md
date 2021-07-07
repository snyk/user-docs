# How to install and configure your Snyk Broker client

####  [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

###  [ ]() <a id="category-name"></a>

#### Snyk Broker

* [ Broker introduction](/hc/en-us/articles/360015367178-Broker-introduction)
* [ Set up Snyk Broker](/hc/en-us/articles/360015296637-Set-up-Snyk-Broker)
* [ Prepare Snyk Broker for deployment](/hc/en-us/articles/360015296717-Prepare-Snyk-Broker-for-deployment)
* [ Enable permissions for Snyk Broker from your third-party tool](/hc/en-us/articles/360015296737-Enable-permissions-for-Snyk-Broker-from-your-third-party-tool)
* [ Retrieve a unique Broker client token](/hc/en-us/articles/360015367278-Retrieve-a-unique-Broker-client-token)
* [ How to install and configure your Snyk Broker client](/hc/en-us/articles/360015296937-How-to-install-and-configure-your-Snyk-Broker-client)
* [ Upgrade your Snyk Broker client](/hc/en-us/articles/360015367458-Upgrade-your-Snyk-Broker-client)

1.  [Docs Library \| Snyk](/hc/en-us)
2.  [Integrations](/hc/en-us/categories/360000598398-Integrations)
3.  [Snyk Broker](/hc/en-us/sections/360001138138-Snyk-Broker)

##  How to install and configure your Snyk Broker client

Snyk’s recommended and supported method of running the Snyk Broker client is using Docker.

To install and configure your client:

1. Visit [the broker repository](https://github.com/snyk/broker) and follow the instructions for the relevant SCM to run one of our provided images, or derive your own.
2. For the environment variables required to run the Broker client, you must:
   1. 1. Retrieve your [unique broker client token,](https://support.snyk.io/hc/en-us/articles/360015367278-Retrieve-a-unique-Broker-client-token)
      2. Ensure that you have the necessary permissions granted to your API credentials for the integration you are using by following [one of our integration guides.](https://support.snyk.io/hc/en-us/categories/360000598398-Integrations)
3. After the broker is running, **identifying Broker server** appears towards the end of the log results, showing that a connection to Snyk has been established
4. Visit the settings for your brokered integration in [the Snyk application](https://app.snyk.io/) to see a message like the following: 
5. You can now use Snyk as normal.

* 
