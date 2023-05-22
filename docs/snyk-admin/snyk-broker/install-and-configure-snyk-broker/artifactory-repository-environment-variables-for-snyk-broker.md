# Artifactory Repository - environment variables for Snyk Broker

The following environment variables are needed to customize the Broker Client for Artifactory Repository:

`BROKER_TOKEN` - the Snyk Broker token, obtained from your Artifactory integration settings (**Integrations > Artifactory**).

`ARTIFACTORY_URL` - the URL of your Artifactory deployment, such as `<yourdomain>.artifactory.com/artifactory`.

The following fields are optional:

* _Port_: Omit if no port number is needed.
* _Basic auth_: Omit if no basic auth required.\
  URL encode both username and password info to avoid errors that may prevent authentication.
* _Protocol_: Defaults to `https://`\
  This should only be specified when no certificate is present and `http://` is required instead for your instance

`ARTIFACTORY_URL` format with optional fields:\
`[http://][username:password@]hostname[:port]/artifactory`\
Example:\
`http://alice:mypassword@acme.com:8080/artifactory`

Optional. `RES_BODY_URL_SUB` - The URL of the Artifactory instance, including http:// and without basic auth credentials. **Required for npm/Yarn integrations only**.\
Example:\
`http://acme.com/artifactory`
