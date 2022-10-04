# Data residency at Snyk

## **What is data residency?**

Data residency allows you to control what region Snyk hosts a selected subset of your data in (see [regional and global data](data-residency-at-snyk.md#how-it-works)).

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Data residency is available for [Snyk Open Source](../scan-application-code/snyk-open-source/), [Snyk Code](../scan-application-code/snyk-code/), [Snyk Container](../snyk-container/), and [Snyk Infrastructure as Code (IaC)](../scan-cloud-deployment/snyk-infrastructure-as-code/).

### **How does it work?**

During your system onboarding, you can work with your account team to select a hosting region.

{% hint style="warning" %}
After you select a region, that data cannot be migrated to another region. Moving to a new region requires a complete re-onboarding.
{% endhint %}

### **What regions are available?**

Snyk offers data residency for the following regions:

|          Region         |                           URL                          |
| :---------------------: | :----------------------------------------------------: |
|      USA (default)      |        [http://app.snyk.io](http://app.snyk.io)        |
| EU (Germany, Frankfurt) | [http://app.eu.snyk.io](http://app.eu.snyk.io)\*\*\*\* |

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

## EU region hosting notes

Snyk offers almost all the same features, support, and performance in the EU region as in the US region.

#### Features not yet available in the EU

* Summary and Issues tabs in Reporting
* C/C++ language support in Snyk Open Source
* Snyk Cloud

### EU Datacenter account creation

EU Datacenter Snyk accounts are only available with the purchase of an [Enterprise plan](https://snyk.io/plans/).

### URLs and endpoints to the EU datacenter

| Resource       | URL                                                                                                                                                                                                                                                                                                                                                                                                                                |
| -------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Login / Web UI | <p><a href="https://app.eu.snyk.io/">https://app.eu.snyk.io/</a><br></p><p><strong>Note:</strong> using Snyk.io or app.snyk.io will not redirect you.</p>                                                                                                                                                                                                                                                                          |
| API            | <p>Follow the docs as normal, but base URLs are</p><p>V1: https://api.eu.snyk.io/v1/</p><p>V3: https://api.eu.snyk.io/rest/</p>                                                                                                                                                                                                                                                                                                    |
| CLI            | <p>Both CLI<strong>,</strong> and CI running CLI, need to be configured to run against your instance by running:</p><pre><code>snyk config set endpoint=https://app.eu.snyk.io/api </code></pre><p>Alternatively, have an environment variable on your machine or CI tool</p><pre><code>SNYK_API=https://app.eu.snyk.io/api</code></pre><p>Jenkins additional argument</p><pre><code>--API=https://app.eu.snyk.io/api</code></pre> |
| IDEs           | <p>IDEs are all running CLI under the hood, so use the settings to set the endpoint accordingly.<br><br>For Snyk Code, ensure the latest version of IDE plugins are in use. The following table specifies the minimum version required:<br><br>VSCode - 1.2.18<br>Visual Studio - 1.1.21<br>IntelliJ - 2.4.32<br><br></p>                                                                                                          |
| Broker         | <p>Follow github.com/snyk/broker docs as normal BUT add an extra environment variable in the container:<br></p><p><code>-e BROKER_SERVER_URL=https://broker.eu.snyk.io</code></p>                                                                                                                                                                                                                                                  |

## **How we maintain GDPR compliance**

We take privacy seriously, and operate a global privacy program to meet the requirements of the GDPR, CCPA, and other applicable privacy laws. We treat all user data the same way and use industry-standard technical and organizational measures to secure the information we store. Our Privacy Program is tailored to meet both legal requirements as well as your needs.

Learn more at the[ ](https://www.atlassian.com/trust/privacy)[Snyk Privacy Policy](https://snyk.io/policies/privacy/).
