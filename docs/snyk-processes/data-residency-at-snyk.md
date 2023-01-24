# Regional hosting and data residency

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can host your data in a number of regions. During your system's initial onboarding, you can work with your account team to select a region.

{% hint style="info" %}
By default, customers are hosted in the US region.
{% endhint %}

## **What is data residency?**

Data residency allows you to control what region Snyk hosts a selected subset of your data in. For more information. see [Regional and global data](data-residency-at-snyk.md#regional-and-global-data) on this page.

Data residency is available for [Snyk Open Source](../products/snyk-open-source/), [Snyk Code](../products/snyk-code/), [Snyk Container](../scan-containers/), and [Snyk Infrastructure as Code (IaC)](../scan-cloud-deployment/snyk-infrastructure-as-code/).

## **How does data residency work?**

During your system onboarding, you can work with your account team to select a hosting region.

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

## Regional and global data

Snyk leverages a number of sub-processors to provide you with a high quality service. Thus not all data types can be stored within your region of choice. The list of sub-processors is available [on the Snyk website](https://snyk.io/policies/sub-processors/).

See [How Snyk handles your data](https://docs.snyk.io/more-info/how-snyk-handles-your-data) for product-specific examples on how this data is handled.

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

## EU and AU region hosting notes

Snyk offers almost all the same features, support, and performance in the EU region as in the US region.

The following **features are not yet available**:

* Summary and Issues tabs in Reporting
* C/C++ language support in Snyk Open Source
* Snyk Cloud

## EU and AU datacenter account creation

EU and AU datacenter Snyk accounts are available only with the purchase of an [Enterprise plan](https://snyk.io/plans/). The resources and URLs are as follows.

### Login and Web UI URLs

**EU**

[https://app.eu.snyk.io/](https://app.eu.snyk.io/)

**AU**

[https://app.au.snyk.io/](https://app.au.snyk.io/)

**Note:** using Snyk.io or app.snyk.io will not redirect you.

### API URLs

Follow the docs as usual, but base URLs are:

**EU**

v1: https://api.eu.snyk.io/v1/\
v3: https://api.eu.snyk.io/rest/

**AU**

v1: https://api.au.snyk.io/v1/\
v3: https://api.au.snyk.io/rest/

### CLI and CI pipelines URLs

Both CLI and CI running CLI need to be configured to run against your instance. To do this, run the following:

**EU**

`snyk config set endpoint=https://app.eu.snyk.io/api`

**AU**

`snyk config set endpoint=https://app.au.snyk.io/api`

Alternatively, have an environment variable on your machine or CI tool:

**EU**

`SNYK_API=https://app.eu.snyk.io/api`

**AU**

`SNYK_API=https://app.au.snyk.io/api`

When running Jenkins, use the additional argument:

**EU**

`--API=https://app.eu.snyk.io/api`

**AU**

`--API=https://app.au.snyk.io/api`

### IDEs URLs

IDEs are all running CLI under the hood, so use the settings to set the endpoint accordingly. In your IDE settings you will have a "Custom Endpoint" parameter for the custom value.

For Snyk Code, ensure the latest version of IDE plugins are in use. The following specifies the minimum version required:

VSCode - 1.2.18\
Visual Studio - 1.1.21\
IntelliJ - 2.4.32

### Broker URLs

Follow [github.com/snyk/broker](https://github.com/snyk/broker) docs as usual BUT add an extra environment variable in the container:

**EU**

`-e BROKER_SERVER_URL=https://broker.eu.snyk.io`

**AU**

`-e BROKER_SERVER_URL=https://broker.au.snyk.io`

For Broker deployed by Helm chart, follow [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm) docs, and add the following variable:&#x20;

**EU**

`--set brokerServerUrl=https://broker.eu.snyk.io`

**AU**

`--set brokerServerUrl=https://broker.au.snyk.io`

### Broker with Code Agent URLs

Follow the [Snyk Broker - Code Agent](https://docs.snyk.io/integrations/snyk-broker/snyk-broker-code-agent) instructions BUT add an extra environment variable in the Code Agent container:

**EU**

`-e UPSTREAM_URL=https://deeproxy.eu.snyk.io`

**AU**

`-e UPSTREAM_URL=https://deeproxy.au.snyk.io`

For Broker with Code Agent deployed by Helm chart, follow the [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm) instructions, and add the following variable in the Code Agent chart:

**EU**

`--set upstreamUrlCodeAgent=https://deeproxy.eu.snyk.io`

**AU**

`--set upstreamUrlCodeAgent=https://deeproxy.au.snyk.io`

## **How Snyk maintains GDPR compliance**

Snyk takes privacy seriously, and operates a global privacy program to meet the requirements of the GDPR, CCPA, and other applicable privacy laws. Snyk treats all user data the same way and uses industry-standard technical and organizational measures to secure the information Snyk stores. Snyk's Privacy Program is tailored to meet both legal requirements and your needs.

Learn more at the[ ](https://www.atlassian.com/trust/privacy)[Snyk Privacy Policy](https://snyk.io/policies/privacy/).
