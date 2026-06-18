# Artifactory Repository - environment variables for Snyk Broker

The following environment variables are needed to customize the Broker Client for Artifactory Repository:

`BROKER_TOKEN` - the Snyk Broker token, obtained from your Artifactory integration settings (**Integrations** > **Artifactory**).

`BROKER_SERVER_URL` - the URL of the Broker server for the region in which your data is hosted. For the commands and URLs to use, see [Broker URLs](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/regional-hosting-and-data-residency#broker-server-urls).

`ARTIFACTORY_URL` - the URL of your Artifactory deployment, such as `<yourdomain>.artifactory.com/artifactory`.

The following fields are optional:

* `Port`: Omit if no port number is needed.
* `Basic auth`: Omit if no basic auth required.\
  URL encode both username and password info to avoid errors that may prevent authentication.
* `Protocol`: Defaults to `https://`\
  Specify this only when no certificate is present and your instance requires `http://` instead.

`ARTIFACTORY_URL` format with optional fields:\
`[http://][username:password@]hostname[:port]/artifactory`\
Example:\
`http://alice:mypassword@acme.com:8080/artifactory`

Optional. `RES_BODY_URL_SUB` - The URL of the Artifactory instance, including http:// and without basic auth credentials. Required for npm and Yarn integrations only.\
Example:\
`http://acme.com/artifactory`
