# Regional hosting and data residency



{% hint style="info" %}
**Feature availability**

Regional hosting and data residency are available only for Enterprise plans. For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Data residency allows you to control the region in which Snyk hosts a selected subset of your data. For information about GDPR, see [Privacy compliance](how-snyk-handles-your-data.md#privacy-compliance).

Data residency is available for [Snyk Open Source](../scan-with-snyk/snyk-open-source/), [Snyk Code](../scan-with-snyk/snyk-code/), [Snyk Container](../scan-with-snyk/snyk-container/), and [Snyk IaC](../scan-with-snyk/snyk-iac/). Snyk can host your data in a number of regions.

This page provides information on:

* [Regionally stored data types](regional-hosting-and-data-residency.md#regionally-stored-data-types)
* [Globally stored data types](regional-hosting-and-data-residency.md#globally-stored-data-types)
* [Available Snyk Regions](regional-hosting-and-data-residency.md#available-snyk-regions)
* [Regional multi- and single-tenant hosting](regional-hosting-and-data-residency.md#regional-multi-and-single-tenant-hosting)

A list of [regional URLs](regional-hosting-and-data-residency.md#regional-urls) is also provided.

## Regional and global data

Snyk uses subprocessors to provide a high-quality service. Thus, not all data types can be stored within your region of choice. The list of subprocessors is available [on the Snyk website](https://snyk.io/policies/subprocessors/).

See [How Snyk handles your data](how-snyk-handles-your-data.md) for product-specific examples of how data is handled.

### Regionally stored data types

* Vulnerability data
* Vulnerability source
* Audit logs
* Integration-related data
* Customer source code

### Globally stored data types

* Billing data
* Customer relationship management data
* Operational logs and metrics
* Product analytics
* Support Tickets
* User authentication data

## Available Snyk Regions

{% hint style="warning" %}
After you select a region, the data in that region cannot be migrated to another region. Moving to a new region requires a complete re-onboarding.
{% endhint %}

During the initial onboarding of your system, you can work with your account team to select a multi-tenant hosting region. For single-tenant availability (Snyk Private Cloud), reach out to your account team in advance of onboarding. When using Snyk features, you will use specific URLs that differ from the SNYK-US-01 URL. See [Regional URLs](regional-hosting-and-data-residency.md#regional-urls) for the list of URLs.

You must configure your environment to set your region before you authenticate. This does not apply when you are using the SNYK-US-01 URL. For details, see the [snyk config environment CLI help](../developer-tools/snyk-cli/commands/config-environment.md).

Snyk offers data residency for the following regions:

|                 Region                 |           URL          |
| :------------------------------------: | :--------------------: |
|             SNYK-US-01 (US)            |   https://app.snyk.io  |
|             SNYK-US-02 (US)            | https://app.us.snyk.io |
|     SNYK-EU-01 (Germany, Frankfurt)    | https://app.eu.snyk.io |
|         SNYK-AU-01 (Australia)         | https://app.au.snyk.io |
| SNYK-GOV-01 (Snyk for Government (US)) | https://app.snykgov.io |

Single-tenant deployments may support more regions than the ones listed here, subject to validation of architectural service supportability by Snyk engineering.

{% hint style="info" %}
Snyk Essentials and Snyk AppRisk work across all regions. There is no need to set your region and use a specific URL.
{% endhint %}

## Regional multi- and single-tenant hosting

{% hint style="warning" %}
The SNYK-US-01 region is available to customers who joined before September 2024 and to Free plan users.
{% endhint %}

For Free or Team plan users in the SNYK-US-01 region, no configuration is required for URLs. For hosting in other regions, you must configure your environment for that region.

SNYK-US-02, EU, and AU data center Snyk accounts are available only with the purchase of an [Enterprise plan](https://snyk.io/plans/).

Snyk offers almost all the same features, support, and performance in the regional multi- and single-tenant regions as in SNYK-US-01. For an up-to-date overview of feature parity across regions, contact your account team.

All Snyk Essentials and Snyk AppRisk features are supported in all multi-tenant environments, while for the single-tenant environment you need to check the availability with your account team.

[Automated provisioning](../implementation-and-setup/enterprise-setup/auto-provisioning-guide.md) is only possible for multi-tenant environments and Pilot or Enterprise plan users.

## Integration considerations

There are special considerations when you set up specific integrations in your Snyk ecosystem.

* If you choose to install the Snyk Runtime Sensor using a Helm chart, you must provide the Snyk API base URL. You can find all the necessary details under step 5 of the [Using a Helm chart](../integrations/snyk-runtime-sensor.md#using-a-helm-chart) section of the Snyk Runtime Sensor documentation page.
* If you choose to set up third-party integrations, verify whether you need to specify the base API URL of the third-party

## Regional URLs

### Login and Web UI URLs

**SNYK-US-01:** [https://app.snyk.io/](https://app.snyk.io/)

**SNYK-US-02:** [https://app.us.snyk.io/](https://app.us.snyk.io/)

**SNYK-EU-01:** [https://app.eu.snyk.io/](https://app.eu.snyk.io/)

**SNYK-AU-01:** [https://app.au.snyk.io/](https://app.au.snyk.io/)

### Support portal links

To view and submit tickets, use the link for your region.

**SNYK-US-01**: [https://support.snyk.io/](https://support.snyk.io/)

**SNYK-US-02**: [https://support.snyk.io/services/auth/sso/MT\_US](https://support.snyk.io/services/auth/sso/MT_US)

**SNYK-EU-01:**[https://support.snyk.io/services/auth/sso/MT\_EU](https://support.snyk.io/services/auth/sso/MT_EU)

**SNYK-AU-01:**[https://support.snyk.io/services/auth/sso/MT\_AU](https://support.snyk.io/services/auth/sso/MT_AU)

### API URLs

Use the base URLs for your region:

**SNYK-US-01:** **API v1** : https://api.snyk.io/v1/ and **REST** API: https://api.snyk.io/rest/

**SNYK-US-02**: **API v1:** https://api.us.snyk.io/v1/ and **REST** API: https://api.us.snyk.io/rest/

**SNYK-EU-01: API v1**: https://api.eu.snyk.io/v1/ and **REST** API: https://api.eu.snyk.io/rest/

**SNYK-AU-01: API v1**: https://api.au.snyk.io/v1/ and **REST** API: https://api.au.snyk.io/rest/

### CLI and CI pipeline URLs

Both the CLI and CI running CLI must be configured to run against your instance.

To do this, for [CLI v1.1293.0](https://updates.snyk.io/announcing-snyk-cli-v1-1293-0-299452) and later versions, use the [snyk config environment command](../developer-tools/snyk-cli/commands/config-environment.md), for example:

`snyk config environment SNYK-US-02`

The [Supported environment URL mappings](../developer-tools/snyk-cli/commands/config-environment.md#supported-environment-urls-mappings) are listed in the `snyk config environment` help.

### IDEs URLs

{% hint style="warning" %}
Ensure you are using the latest version of the IDE plugins. The following specifies the minimum versions required:\
VSCode - 1.2.18\
Visual Studio - 1.1.21\
IntelliJ - 2.4.32
{% endhint %}

Snyk IDE extensions have modifiable options similar to the CLI and must be configured to use the appropriate endpoint. In the extension settings for Snyk in your IDE, set **Custom Endpoint** to the appropriate value for SNYK-US-02, SNYK-EU-01, and SNYK-AU-01 as follows.

**SNYK-US-01:** `https://api.snyk.io` (no configuration required)

**SNYK-US-02:** `https://api.us.snyk.io`

**SNYK-EU-01:** `https://api.eu.snyk.io`

**SNYK-AU-01 :** `https://api.au.snyk.io`

Multi-tenant users who do not belong to the `SNYK-US-01` region will be automatically redirected to the correct domain for the email with which the user authenticated. The redirect will not occur for cases where the users are expected to use a custom URL, such as companies with single-tenant setups.

### Broker client URLs

**SNYK-US-01:** `https://api.snyk.io` (no configuration required)

**SNYK-US-02:** `https://api.us.snyk.io`

**SNYK-EU-01:** `https://api.eu.snyk.io`

**SNYK-AU-01 :** `https://api.au.snyk.io`

### Broker server URLs

Use [github.com/snyk/broker](https://github.com/snyk/broker) and add an extra environment variable in the container:

**SNYK-US-01**: `https://broker.snyk.io` (no configuration required)

**SNYK-US-02**: `-e BROKER_SERVER_URL=https://broker.us.snyk.io`

**SNYK-EU-01:** `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`

**SNYK-AU-01:** `-e BROKER_SERVER_URL=https://broker.au.snyk.io`

For Broker deployed by Helm chart, use [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm) and add the following variable:

**SNYK-US-02:** `--set brokerServerUrl=https://broker.us.snyk.io`

**SNYK-EU-01:** `--set brokerServerUrl=https://broker.eu.snyk.io`

**SNYK-AU-01:** `--set brokerServerUrl=https://broker.au.snyk.io`

### Broker with high availability (HA) mode URLs

Follow the [High availability mode](../implementation-and-setup/enterprise-setup/snyk-broker/high-availability-mode.md) instructions, BUT use the following details for BROKER\_DISPATCHER\_BASE\_URL:

**SNYK-US-02**: `-e BROKER_DISPATCHER_BASE_URL=https://api.us.snyk.io`

**SNYK-EU-01:** `-e BROKER_DISPATCHER_BASE_URL=https://api.eu.snyk.io`

**SNYK-AU-01 :** `-e BROKER_DISPATCHER_BASE_URL=https://api.au.snyk.io`

For Broker deployed by Helm chart, edit the `values.yaml` file to include the relevant details in brokerDispatcherUrl.

### Broker with Snyk Code Local Engine (SCLE)

{% include "../.gitbook/includes/release-status-snyk-code-local-engine.md" %}

Set up your `values-customer-settings.yml` with the correct Broker Server URL for your region based on the values found in the [Broker URLs](regional-hosting-and-data-residency.md#broker-server-urls) section.

Then add an extra variable to the `values-customer-settings.yml`:

**SNYK-US-02**\
`deeproxy:`\
`verificationEndpoint: "https://api.us.snyk.io/v1/validate/token/snyk-to-deepcode-proxy-validation"`

**SNYK-EU-01**\
`deeproxy:`\
`verificationEndpoint: "https://api.eu.snyk.io/v1/validate/token/snyk-to-deepcode-proxy-validation"`

**SNYK-AU-01**\
`deeproxy:`\
`verificationEndpoint: "https://api.au.snyk.io/v1/validate/token/snyk-to-deepcode-proxy-validation"`

### Kubernetes Integration URLs

Follow the instructions to install the [Snyk Controller with Helm](../scan-with-snyk/snyk-container/kubernetes-integration/install-the-snyk-controller/install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md) with an extra variable added to the `helm upgrade`command:

**SNYK-US-01**: (no configuration required)

**SNYK-US-02**: `--set integrationApi=https://api.us.snyk.io/v2/kubernetes-upstream`

**SNYK-EU-01:** `--set integrationApi=https://api.eu.snyk.io/v2/kubernetes-upstream`

**SNYK-AU-01:** `--set integrationApi=https://api.au.snyk.io/v2/kubernetes-upstream`
