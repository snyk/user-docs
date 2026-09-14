---
description: How to configure Snyk GitHub Actions to use a regional Snyk API endpoint, and troubleshoot related authentication errors
nav_context: agnostic
---

# Regional API endpoints

For information about using GitHub Actions with Snyk, visit [GitHub Actions for Snyk setup and checking for vulnerabilities](./).

By default, Snyk GitHub Actions use the `https://api.snyk.io` endpoint. To configure Snyk to use a different endpoint, set a `SNYK_API` environment variable in your workflow, for example, `https://api.eu.snyk.io`.

For more information about environment configuration, visit [Configure the Snyk CLI](../../snyk-cli/configure-the-snyk-cli/). For the list of available regions, visit [Regional hosting and data residency](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#available-snyk-regions).

An example follows of how you can modify a Snyk GitHub Actions workflow to use an alternate endpoint:

```yaml
name: Example workflow using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/node@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
          SNYK_API: https://api.us.snyk.io
```

## Troubleshooting

If a workflow fails to authenticate, the token and the API endpoint likely belong to different regions.

### Why this happens

Each Snyk region issues tokens that only work against that region's API endpoint. If `SNYK_API` points to the default endpoint, `https://api.snyk.io`, but your Snyk Organization is provisioned in a different region, Snyk rejects the request with an authentication error, even though the token itself is valid.

### Identify your region

Check your Snyk Organization's URL in the Snyk web UI, or contact your account team to confirm your region. New Enterprise and Pilot accounts provisioned in the United States through Automated Provisioning use SNYK-US-02 (`https://app.us.snyk.io`), not the default SNYK-US-01.

### Common mistakes

* Using the default `api.snyk.io` endpoint for an account provisioned outside SNYK-US-01
* Setting `SNYK_API` to the web UI URL, for example `https://app.us.snyk.io`, instead of the API URL, `https://api.us.snyk.io`
* Assuming SNYK-US-01 and SNYK-US-02 are interchangeable — both are US-hosted, but they are separate regions with separate endpoints
