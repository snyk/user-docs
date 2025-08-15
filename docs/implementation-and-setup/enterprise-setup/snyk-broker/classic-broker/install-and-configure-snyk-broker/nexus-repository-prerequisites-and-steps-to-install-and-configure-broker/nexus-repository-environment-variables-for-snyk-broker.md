# Nexus Repository - environment variables for Snyk Broker

## Environment variables for Nexus 3 configuration

The following environment variables are needed to customize the Broker client for Nexus 3:

`BROKER_TOKEN`\
The Snyk Broker token, obtained from your Nexus integration settings (**Integrations > Nexus**).

`BROKER_SERVER_URL` - the URL of the Broker server for the region in which your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).

`BASE_NEXUS_URL`\
The URL of your Nexus 3 deployment.\
Example:\
`BASE_NEXUS_URL=https://[<username_or_token><password_or_token>]@<your.nexus.hostname>`\
Must not end with a forward slash.\
The following field is optional:\
`Auth`: Omit if no auth required.\
Can either be plain text or a two-part token (`Nexus Pro`).\
URL encode username, password, and tokens to avoid errors that may prevent authentication.\
Minimal example: `acme.com`\
Complex example: `https://alice:mypassword@acme.com`

`BROKER_CLIENT_VALIDATION_URL`\
Nexus validation URL, checked by Broker Client `systemcheck` endpoint.\
If Nexus user requires `auth`, use `$BASE_NEXUS_URL/service/rest/v1/status/check`\
Example:\
`https://<user>:<pass>@<your.nexus.hostname>/service/rest/v1/status/check`)\
Otherwise, use `$BASE_NEXUS_URL/service/rest/v1/status`\
Example:\
`https://<your.nexus.hostname>/service/rest/v1/status`).

Optional. `RES_BODY_URL_SUB`\
This URL substitution is required for npm and Yarn integration and is the same as the URL of the Nexus without credentials appended with `/repository`\
Example:\
`https://<your.nexus.hostname>/repository`. Must not end with a forward slash.

## Environment variables for Nexus 2 configuration

The following environment variables are needed to customize the Broker client for Nexus 2:

`BROKER_TOKEN` - the Snyk Broker token, obtained from your Nexus integration settings (**Integrations > Nexus**).

`BROKER_SERVER_URL` - the URL of the Broker server for the region in which your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).

`BASE_NEXUS_URL`- the URL of your Nexus 2 deployment.\
Example:\
`BASE_NEXUS_URL=https://[username_or_token:password_or_token]@acme.com`\
Must not end with a forward slash.\
The following fields are optional:\
`Auth`: Omit if no auth required.\
Can be either plain text or a two-part token (`Nexus Pro`).\
URL encode username, password, and tokens to avoid errors that may prevent authentication.\
Minimal example: `https://acme.com`\
Complex example: `https://alice:mypassword@acme.com: 8000`

`NEXUS_URL`: Used as an override if your repositories do not live under /nexus/content.\
The URL that points to the base of your repositories within Nexus. By default, the broker assumes `BASE_NEXUS_URL`/nexus/content/ as the value.

`RES_BODY_URL_SUB`\
The URL of the Nexus instance, including `https://` and `/nexus/content` without basic auth credentials. Required for npm and Yarn integrations onl&#x79;**.** Must not end with a forward slash.

Examples:\
`https://acme.com/nexus/content/groups`\
`https://acme.com/nexus/content/repositories`
