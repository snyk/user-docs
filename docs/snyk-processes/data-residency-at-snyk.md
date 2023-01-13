# Regional hosting and data residency

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Snyk can host your data in a number of available regions. During your system's initial onboarding, you can work with your account team to select a region.

{% hint style="info" %}
By default, customers are hosted in the US region.
{% endhint %}

## **What is data residency?**

Data residency allows you to control what region Snyk hosts a selected subset of your data in (see [regional and global data](data-residency-at-snyk.md#how-it-works)).

Data residency is available for [Snyk Open Source](../products/snyk-open-source/), [Snyk Code](../products/snyk-code/), [Snyk Container](../products/snyk-container/), and [Snyk Infrastructure as Code (IaC)](../products/snyk-infrastructure-as-code/).

### **How does it work?**

During your system onboarding, you can work with your account team to select a hosting region.

{% hint style="warning" %}
After you select a region, that data cannot be migrated to another region. Moving to a new region requires a complete re-onboarding.
{% endhint %}

### **What regions are available?**

Snyk offers data residency for the following regions:

|          Region         |           URL          |
| :---------------------: | :--------------------: |
|      USA (default)      |   https://app.snyk.io  |
| EU (Germany, Frankfurt) | https://app.eu.snyk.io |
|     AUS (Australia)     | https://app.au.snyk.io |

### Regional and global data

Snyk leverages a number of sub-processors to provide you with a high quality service. As such, not all data types can be stored within your region of choice. Our list of sub-processors is available [here](https://snyk.io/policies/sub-processors/).

See the [How Snyk handles your data](https://docs.snyk.io/more-info/how-snyk-handles-your-data) page for product-specific examples on how this data is handled.

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

#### Features not yet available

* Summary and Issues tabs in Reporting
* C/C++ language support in Snyk Open Source
* Snyk Cloud

### EU and AU datacenter account creation

EU and AU datacenter Snyk accounts are only available with the purchase of an [Enterprise plan](https://snyk.io/plans/).

### URLs and endpoints

| Resource           | URL                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Login / Web UI     | <p><strong>EU</strong><br><a href="https://app.eu.snyk.io/">https://app.eu.snyk.io/</a><br><strong>AU</strong><br><a href="https://app.au.snyk.io/">https://app.au.snyk.io/</a></p><p><strong>Note:</strong> using Snyk.io or app.snyk.io will not redirect you.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| API                | <p>Follow the docs as normal, but base URLs are:<br><strong>EU</strong><br>V1: https://api.eu.snyk.io/v1/</p><p>V3: https://api.eu.snyk.io/rest/<br><strong>AU</strong><br>V1: https://api.au.snyk.io/v1/</p><p>V3: https://api.au.snyk.io/rest/</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| CLI / CI pipelines | <p>Both CLI<strong>,</strong> and CI running CLI, need to be configured to run against your instance by running:</p><p></p><p><strong>EU</strong></p><p><code>snyk config set endpoint=https://app.eu.snyk.io/api</code><br><br><strong>AU</strong><br><strong></strong><code>snyk config set endpoint=https://app.au.snyk.io/api</code><br><strong></strong><br><strong></strong>Alternatively, have an environment variable on your machine or CI tool:<br><br><strong>EU</strong><br><strong></strong><code>SNYK_API=https://app.eu.snyk.io/api</code><br><strong>AU</strong><br><strong></strong><code>SNYK_API=https://app.au.snyk.io/api</code><br><br><br>When running Jenkins, please use the additional arguement:<br><br>E<strong>U</strong><br><strong></strong><code>--API=https://app.eu.snyk.io/api</code><br><strong>AU</strong><br><strong></strong><code>--API=https://app.au.snyk.io/api</code></p> |
| IDEs               | <p>IDEs are all running CLI under the hood, so use the settings to set the endpoint accordingly. In your IDE settings you will have a "Custom Endpoint" parameter for the custom value.<br><br>For Snyk Code, ensure the latest version of IDE plugins are in use. The following table specifies the minimum version required:<br><br>VSCode - 1.2.18<br>Visual Studio - 1.1.21<br>IntelliJ - 2.4.32</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| Broker             | <p>Follow <a href="https://github.com/snyk/broker">github.com/snyk/broker</a> docs as normal BUT add an extra environment variable in the container:<br><br><strong>EU</strong></p><p><code>-e BROKER_SERVER_URL=https://broker.eu.snyk.io</code><br><strong>AU</strong></p><p><code>-e BROKER_SERVER_URL=https://broker.au.snyk.io</code></p><p></p><p>For Broker deployed by Helm chart, follow <a href="https://github.com/snyk/snyk-broker-helm">https://github.com/snyk/snyk-broker-helm</a> docs, and add the following variable: </p><p><br><strong>EU</strong></p><p><code>--set brokerServerUrl=https://broker.eu.snyk.io</code><br><strong>AU</strong></p><p><code>--set brokerServerUrl=https://broker.au.snyk.io</code><br></p>                                                                                                                                                                           |
| Code Agent         | <p>Follow <a href="https://docs.snyk.io/integrations/snyk-broker/snyk-broker-code-agent">Snyk Broker - Code Agent</a> instructions BUT add an extra environment variable in the Code Agent container:<br><br><strong>EU</strong></p><p><code>-e UPSTREAM_URL=https://deeproxy.eu.snyk.io</code><br><strong>AU</strong></p><p><code>-e UPSTREAM_URL=https://deeproxy.au.snyk.io</code></p><p></p><p>For Broker with Code Agent deployed by Helm chart, follow <a href="https://github.com/snyk/snyk-broker-helm">https://github.com/snyk/snyk-broker-helm</a> docs, and add the following variable in the Code Agent chart: </p><p><br><strong>EU</strong></p><p><code>--set upstreamUrlCodeAgent=https://deeproxy.eu.snyk.io</code><br><strong>AU</strong></p><p><code>--set upstreamUrlCodeAgent=https://deeproxy.au.snyk.io</code></p>                                                                              |

## **How we maintain GDPR compliance**

We take privacy seriously, and operate a global privacy program to meet the requirements of the GDPR, CCPA, and other applicable privacy laws. We treat all user data the same way and use industry-standard technical and organizational measures to secure the information we store. Our Privacy Program is tailored to meet both legal requirements as well as your needs.

Learn more at the[ ](https://www.atlassian.com/trust/privacy)[Snyk Privacy Policy](https://snyk.io/policies/privacy/).
