# Authentication for API

This section provides information about how to [Authenticate for the API](authenticate-for-the-api.md), including obtaining your API token and using it in the authentication header, and [Snyk API token permissions users can control](snyk-api-token-permissions-users-can-control.md).

For instructions on obtaining a new API token, see [Revoke and regenerate a Snyk API token](revoke-and-regenerate-a-snyk-api-token.md).

The following explains **when to use an API token** and **when to use a service account token**.

Your Snyk API token is a personal token available under your user profile. The Snyk API token is associated with your Snyk Account and not with a specific Organization.

Free and Team plan and trial users have access to only this personal token. The personal token can be used to authenticate with the Snyk CLI running on a local or a build machine and an IDE when you are setting a token manually. Use a personal token with caution if you are authenticating with the API or for CI/CD.

Enterprise users have access to a personal token under their profile and to service account tokens. For details, see [Service accounts](../../../enterprise-setup/service-accounts/).

* **Enterprise users should use a service account** to authenticate for any kind of automation. This includes, but is not limited to, CI/CD scanning with the CLI or build system plugins and any automation, including automation with the API.
* **Enterprise users should use the personal token** under their user profile for:
  * Running the CLI locally on their machine
  * Authenticating with the IDE manually
  * Running API calls one time, for example, to test something

For more information on the personal Snyk API token, see [Authenticate for the API](authenticate-for-the-api.md) and  [Authenticate to use the CLI](../../../snyk-cli/authenticate-to-use-the-cli.md).

