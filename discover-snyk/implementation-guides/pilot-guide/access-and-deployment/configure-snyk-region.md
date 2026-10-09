---
nav_context: agnostic
description: >-
  How to confirm your Snyk region and sign in, using SNYK-US-02 as the example
  setup
---

# Configure Snyk region

{% include "../../../.gitbook/includes/pilot-guide-navigation.md" %}

Confirm with your Snyk account team the region where your Snyk account is located. The guide below shows the setup for SNYK-US-02, but the process is similar for [other regions](https://docs.snyk.io/snyk-data-and-governance/regional-hosting-and-data-residency#available-snyk-regions).

## Snyk Web UI

Passwordless usernames with email authentication have been created for you. Use the following URL to log in to Snyk: [https://app.us.snyk.io/login/passwordless](https://app.us.snyk.io/login/passwordless).

## Snyk CLI (and any CI/CD tools that use the Snyk CLI)

Ensure you set the environment variable `SNYK_API` to point to `api.us.snyk.io` before trying to authenticate the CLI as described on the [Configure Snyk CLI to connect to Snyk API ](https://docs.snyk.io/developer-tools/snyk-cli/configure-the-snyk-cli/configure-snyk-cli-to-connect-to-snyk-api)page.

When running the CLI in a CI/CD pipeline, ensure that the `SNYK_API` variable is set before running `snyk auth` . For example:

```bash
export SNYK_API=api.us.snyk.io
export SNYK_TOKEN=<TOKEN>
snyk auth $SNYK_TOKEN
snyk test
```

See [authenticating the CLI](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/authenticate-to-use-the-cli) for more details.

## Snyk API

Ensure the correct base URL is being used for your region. You can use environment variables to set the Organization ID and the Snyk API token, and then make an API request like:

```
curl --request GET \
    --url "https://api.us.snyk.io/rest/orgs/$ORG_ID/projects?version=2024-06-10" \
    --header "Content-Type: application/vnd.api+json" \ 
    --header "Authorization: token $API_TOKEN"
```

## Snyk IDE extension

Set the Custom Endpoint to https://api.us.snyk.io in the IDE extension settings before authenticating.
