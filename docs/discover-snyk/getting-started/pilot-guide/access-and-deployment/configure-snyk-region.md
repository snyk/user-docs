# Configure Snyk region

{% include "../../../../.gitbook/includes/pilot-guide-toc.md" %}

Confirm with your Snyk account team the region where your Snyk account is located. The guide below shows the setup for SNYK-US-02, but the process is similar for [other regions](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#available-snyk-regions).

## Snyk Web UI

Passwordless usernames with email authentication have been created for you. Use the following URL to log in to Snyk: [https://app.us.snyk.io/login/passwordless](https://app.us.snyk.io/login/passwordless).

## Snyk CLI (and any CI/CD tools that use the Snyk CLI)

Ensure you set the environment variable `SNYK_API` to point to `api.us.snyk.io` before trying to authenticate the CLI as described on the [Configure Snyk CLI to connect to Snyk API ](../../../../cli-ide-and-ci-cd-integrations/snyk-cli/configure-the-snyk-cli/configure-snyk-cli-to-connect-to-snyk-api.md)page.

When running the CLI in a CI/CD pipeline, ensure that the `SNYK_API` variable is set before running `snyk auth` . For example:

```bash
export SNYK_API=api.us.snyk.io
export SNYK_TOKEN=<TOKEN>
snyk auth $SNYK_TOKEN
snyk test
```

See [authenticating the CLI](../../../../developer-tools/snyk-cli/authenticate-to-use-the-cli.md) for more details.

## Snyk API

Ensure the correct base URL is being used for your region. You can use environment variables to set the Organization ID and the Snyk API token, and then make an API request like:

```
curl --request GET \
    --url "https://api.us.snyk.io/rest/orgs/$ORG_ID/projects?version=2024-06-10" \
    --header "Content-Type: application/vnd.api+json" \ 
    --header "Authorization: token $API_TOKEN"
```

## Snyk IDE extension

Set the Custom Endpoint to https://api.us.snyk.io in the IDE extension settings before authenticating. For example, in VS Code looks like this:

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXdpXh6Us6U_mzXGtnWRqfCVIsao_BQrMyi9Y2rbWhvbcEJKZMl3497yDG6GPUe9zrXDLgoChd-KafUhoLID5gqsy50X8PHD34kJ5XALvZZb5xG6UZUmxrxXCkr5cslSk63Msi_L?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>
