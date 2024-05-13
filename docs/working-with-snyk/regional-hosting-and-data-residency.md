# Regional hosting and data residency

{% hint style="info" %}
**Feature availability**

Regional hosting and data residency are available only for Enterprise plans.

For more information, see [Plans and pricing](https://snyk.io/plans).
{% endhint %}

Snyk can host your data in a number of regions. During initial onboarding of your system, you can work with your account team to select a multi-tenant region. For single-tenant availability (Snyk Private Cloud), reach out to your account team in advance of onboarding.

{% hint style="info" %}
By default, customers are hosted in the US region.
{% endhint %}

## **What is data residency?**

Data residency allows you to control the region in which Snyk hosts a selected subset of your data. For more information, see [Regional and global data](regional-hosting-and-data-residency.md#regional-and-global-data).

Data residency is available for [Snyk Open Source](../scan-with-snyk/snyk-open-source/), [Snyk Code](../scan-with-snyk/snyk-code/), [Snyk Container](../scan-with-snyk/snyk-container/), and [Snyk IaC](../scan-with-snyk/snyk-iac/).

## **How does data residency work?**

During your system onboarding, you can work with your account team to select a hosting region. When using Snyk features, you will use specific URLs that differ from the default US region URLs.

{% hint style="warning" %}
After you select a region, the data in that region cannot be migrated to another region. Moving to a new region requires a complete re-onboarding.
{% endhint %}

## **What regions are available?**

Snyk offers data residency for the following regions:

|          Region         |           URL          |
| :---------------------: | :--------------------: |
|      USA (default)      |   https://app.snyk.io  |
| EU (Germany, Frankfurt) | https://app.eu.snyk.io |
|     AUS (Australia)     | https://app.au.snyk.io |

Single-tenant deployments may support more regions than the ones listed here, subject to validation of architectural service supportability by Snyk engineering.&#x20;

## Regional and global data

Snyk leverages a number of sub-processors to provide you with a high-quality service. Thus not all data types can be stored within your region of choice. The list of sub-processors is available [on the Snyk website](https://snyk.io/policies/subprocessors/).

See [How Snyk handles your data](how-snyk-handles-your-data.md) for product-specific examples of how this data is handled.

{% tabs %}
{% tab title="Product" %}
* Snyk Open Source
* Snyk Code
* Snyk Container
* Snyk Infrastructure as Code (IaC)
{% endtab %}

{% tab title="Regionally stored data" %}
* Vulnerability data
* Vulnerability source
* Audit logs
* Integration-related data
* Customer source code
{% endtab %}

{% tab title="Globally stored data" %}
* Billing data
* Customer relationship management data
* Operational logs and metrics
* Product analytics
* Support Tickets
* User authentication data
{% endtab %}
{% endtabs %}

## US region URLs

The following are the default URLs for Snyk features in the US region. No configuration is required to use these URLs. For hosting in non-default regions, configuration is required. See [Login and Web UI URLs](regional-hosting-and-data-residency.md#login-and-web-ui-urls).

US login and Web UI\
https://app.snyk.io

US Support portal tickets\
[US Tenant Support Portal](https://support.snyk.io)

US API v1 URL\
https://api.snyk.io/v1/\
US REST API URL\
https://api.snyk.io/rest/

US CLI and CI pipelines URL\
https://app.snyk.io/api

US IDEs URL\
https://app.snyk.io/api

US Broker server URL\
https://broker.snyk.io

## Regional multi- and single-tenant hosting notes

Snyk offers almost all the same features, support, and performance in the regional multi- and single-tenant regions as in the US region. For an up-to-date overview of feature parity across regions, reach out to your account team.

### EU and AU data center account creation

EU and AU data center Snyk accounts are available only with the purchase of an [Enterprise plan](https://snyk.io/plans/). The resources and URLs are as follows.

### Login and Web UI URLs

#### **EU**

[https://app.eu.snyk.io/](https://app.eu.snyk.io/)

#### **AU**

[https://app.au.snyk.io/](https://app.au.snyk.io/)

{% hint style="warning" %}
Using Snyk.io or app.snyk.io will not redirect you to these URLs.
{% endhint %}

### Support Portal Tickets

#### **EU**

[EU Tenant Support Portal](https://snyk-mt-eu-prod-1.eu.auth0.com/samlp/xU6rUSC7zvEco2ndKemfJNT6oKtc13Qw)

#### **AU**

[AU Tenant Support Portal](https://snyk-mt-au-prod-1.au.auth0.com/samlp/HnGYPWKxM9JegYL2aq0OAdBAwGJDz0vQ)

### API URLs

Follow the docs as usual, but base URLs are:

#### **EU**

API v1: https://api.eu.snyk.io/v1/\
REST API: https://api.eu.snyk.io/rest/

#### **AU**

API v1: https://api.au.snyk.io/v1/\
REST API: https://api.au.snyk.io/rest/

### CLI and CI pipelines URLs

Both CLI and CI running CLI need to be configured to run against your instance. To do this, run the following:

#### **EU**

`snyk config set endpoint=https://app.eu.snyk.io/api`

#### **AU**

`snyk config set endpoint=https://app.au.snyk.io/api`

Alternatively, have an environment variable on your machine or CI tool:

#### **EU**

`SNYK_API=https://app.eu.snyk.io/api`

#### **AU**

`SNYK_API=https://app.au.snyk.io/api`

When running Jenkins, use the additional argument:

#### **EU**

`--API=https://app.eu.snyk.io/api`

#### **AU**

`--API=https://app.au.snyk.io/api`

### IDEs URLs

Snyk IDE extensions have modifiable options similar to the CLI and need to be configured to use the appropriate endpoint. In the extensions settings for Snyk in your IDE, set the **Custom Endpoint** parameter to the appropriate value for EU and AU as follows.

{% hint style="warning" %}
Ensure you are using the latest version of IDE plugins. The following specifies the minimum version required:\
VSCode - 1.2.18\
Visual Studio - 1.1.21\
IntelliJ - 2.4.32
{% endhint %}

#### EU

`https://api.eu.snyk.io`

#### **AU**

`https://api.au.snyk.io`

<figure><img src="../.gitbook/assets/Screenshot 2024-05-02 at 10.42.52.png" alt="Configure the AU endpoint in Visual Studio Code"><figcaption><p>Configure the AU endpoint in Visual Studio Code</p></figcaption></figure>

### Broker URLs

Follow [github.com/snyk/broker](https://github.com/snyk/broker) docs as usual BUT add an extra environment variable in the container:

#### **EU**

`-e BROKER_SERVER_URL=https://broker.eu.snyk.io`

#### **AU**

`-e BROKER_SERVER_URL=https://broker.au.snyk.io`

For Broker deployed by Helm chart, follow [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm) docs, and add the following variable:

#### **EU**

`--set brokerServerUrl=https://broker.eu.snyk.io`

#### **AU**

`--set brokerServerUrl=https://broker.au.snyk.io`

### Broker with high availability (HA) mode URLs.

Follow the [High availability mode](../enterprise-configuration/snyk-broker/high-availability-mode.md) instructions BUT use the following details for BROKER\_DISPATCHER\_BASE\_URL:

#### **EU**

`-e BROKER_DISPATCHER_BASE_URL=https://api.eu.snyk.io`

#### **AU**

`-e BROKER_DISPATCHER_BASE_URL=https://api.au.snyk.io`

For Broker deployed by Helm chart, edit the values.yaml file to include the relevant details in brokerDispatcherUrl.

### Broker with Code Agent URLs

Follow the [Snyk Broker - Code Agent](https://docs.snyk.io/integrations/snyk-broker/snyk-broker-code-agent) instructions BUT add an extra environment variable in the Code Agent container:

#### **EU**

`-e UPSTREAM_URL=https://deeproxy.eu.snyk.io`

#### **AU**

`-e UPSTREAM_URL=https://deeproxy.au.snyk.io`

For Broker with Code Agent deployed by Helm chart, follow the [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm) instructions, and add the following variable in the Code Agent chart:

#### **EU**

`--set upstreamUrlCodeAgent=https://deeproxy.eu.snyk.io`

#### **AU**

`--set upstreamUrlCodeAgent=https://deeproxy.au.snyk.io`

### Snyk Code Local Engine (SCLE)

Set up your `values-customer-settings.yml` with the correct Broker Server URL for your region, based on the values found in the [Broker URLs](regional-hosting-and-data-residency.md#broker-urls) instructions.&#x20;

Then add an extra  variable in the `values-customer-settings.yml`:

#### **EU**

`deeproxy:` \
&#x20; `verificationEndpoint: "https://api.eu.snyk.io/v1/validate/token/snyk-to-deepcode-proxy-validation"`

#### **AU**

`deeproxy:` \
&#x20; `verificationEndpoint: "https://api.au.snyk.io/v1/validate/token/snyk-to-deepcode-proxy-validation"`

## **How Snyk maintains GDPR compliance**

Snyk takes privacy seriously and operates a global privacy program to meet the requirements of the GDPR, CCPA, and other applicable privacy laws. Snyk treats all user data the same way and uses industry-standard technical and organizational measures to secure the information Snyk stores. Snyk's Privacy Program is tailored to meet both legal requirements and your needs.

Learn more at the[ ](https://www.atlassian.com/trust/privacy)[Snyk Privacy Policy](https://snyk.io/policies/privacy/).
