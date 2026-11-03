---
nav_context: agnostic
description: >-
  How to confirm your Snyk region and log in, using SNYK-US-02 as the example
  setup
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Configure Snyk region

{% include "../../../.gitbook/includes/pilot-guide-navigation.md" %}

Confirm with your Snyk account team the region where your Snyk account is located. This page shows the setup for SNYK-US-02. The process is similar for [other regions](https://docs.snyk.io/snyk-data-and-governance/regional-hosting-and-data-residency#available-snyk-regions).

## Snyk Web UI

Snyk creates passwordless usernames with email authentication for you. Use the following URL to log in to Snyk: [https://app.us.snyk.io/login/passwordless](https://app.us.snyk.io/login/passwordless).

## Snyk CLI and CI/CD tools that use the Snyk CLI

Before you authenticate the CLI, set the `SNYK_API` environment variable to `api.us.snyk.io`, as described on the [Configure Snyk CLI to connect to Snyk API](https://docs.snyk.io/developer-tools/snyk-cli/configure-the-snyk-cli/configure-snyk-cli-to-connect-to-snyk-api) page.

When you run the CLI in a CI/CD pipeline, ensure that the `SNYK_API` variable is set before you run `snyk auth`. For example:

```bash
export SNYK_API=api.us.snyk.io
export SNYK_TOKEN=<TOKEN>
snyk auth $SNYK_TOKEN
snyk test
```

For more details, visit [Authenticate to use the CLI](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/authenticate-to-use-the-cli).

## Snyk API

Ensure that you use the correct base URL for your region. You can use environment variables to set the Organization ID and the Snyk API token, and then make an API request such as the following:

```
curl --request GET \
    --url "https://api.us.snyk.io/rest/orgs/$ORG_ID/projects?version=2024-06-10" \
    --header "Content-Type: application/vnd.api+json" \ 
    --header "Authorization: token $API_TOKEN"
```

## Snyk IDE extension

Before you authenticate, set **Custom Endpoint** to https://api.us.snyk.io in the IDE extension settings.
