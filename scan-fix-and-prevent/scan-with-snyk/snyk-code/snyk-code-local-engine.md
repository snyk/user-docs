---
description: Snyk Code Local Engine is deprecated; reference information for existing deployments
nav_context: classic
---

# Snyk Code Local Engine

{% include "../../.gitbook/includes/release-status-snyk-code-local-engine.md" %}

Snyk Code Local Engine (SCLE) is a fully contained version of the Snyk Code Engine that allows you to avoid uploading your code to the internet. When you use the Local Engine, only the scan is performed locally. Your scan results are uploaded to Snyk so you can view them on the Snyk Web UI.

Snyk is not onboarding new Local Engine deployments. If you are evaluating Snyk Code, see [Configure Snyk Code](configure-snyk-code.md) for the supported deployment options.

## Support for existing deployments

If you are running the Local Engine today, your Snyk account team is the route for all configuration, upgrade, and migration questions. The instructions to configure and deploy the Local Engine in your environment are in the Readme file in the Local Engine installation package.

To use the Snyk CLI and IDEs with the Local Engine, provide your Snyk account team with the URL of the Local Engine running on your premises. Once your account team has configured that URL for your Organization, you can view it under **Settings** > **Snyk Code**.

## Regional endpoints

Broker deployments that route through the Local Engine need the Broker Server URL and the `deeproxy` verification endpoint for your Snyk region. Both are listed in [Broker with Snyk Code Local Engine (SCLE)](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/regional-hosting-and-data-residency#broker-with-snyk-code-local-engine-scle).
