# Reporting and BI Integrations: Snowflake Data Share

{% hint style="warning" %}
**Release status** \
Snowflake Data Share is in [Early Access](../../getting-started/snyk-release-process.md#early-access) and available only for customers:

* Under the Enterprise plan. For more information, see [Plans and pricing](https://snyk.io/plans).
* With an active Snowflake account.

For more information, contact your Snyk Account Executive.
{% endhint %}

With the new Snowflake Data Share integration, your data science, BI and AppSec teams can securely access the same underlying data available in Snyk Reporting, but within your own Snowflake account, unlocking powerful new analytical tools to better understand and visualize Snyk data.

Use this integration to enable teams to rapidly build exploratory and custom analytics using the ecosystem of tools that Snowflake supports. Customers can connect Snyk data to BI tools like PowerBI, Tableau, and Looker Data Studio or build custom Streamlit apps.

Having Snyk datasets directly in your Snowflake account opens doors for combining Snyk data with additional data sources, which contributes a great deal to the organization-level holistic AppSec view.

## What is Snowflake Data Share? <a href="#what-is-snowflake-data-share" id="what-is-snowflake-data-share"></a>

[Snowflake Data Share](https://docs.snowflake.com/en/user-guide/data-sharing-intro.html) allows companies to provide data to their customers in a secure, simple, and quick manner. With Snowflake data shares, the data is not exchanged between the accounts; instead, the data consumer is granted with read-only access to the shared database through their own Snowflake account.

### Snowflake Cost Impact  <a href="#main-use-cases" id="main-use-cases"></a>

* **Storage**\
  As data is not being transferred during data sharing, there will be **no** additional storage costs attributed to your Snowflake account.
* **Compute Resources**\
  Snowflake will charge for compute resources that are used to query the data share.

## Main use cases <a href="#main-use-cases" id="main-use-cases"></a>

Snyk Snowflake Data Share can be used for various use cases and can answer countless security and business-related questions. Some of these use cases include:

* **Enhance the AppSec posture visibility for the CISO and management team.**\
  Streamline Snyk data to your BI platforms and existing security dashboards and reflect performance metrics and KPI's, for example, MTTR, SLA compliance, remediation trends, and so on.&#x20;
* **Answer specific questions or surface unique insights.**\
  Better understand risk exposure trends, such as tracking total issues above a specific risk score only affect certain Project collections or tags across all Snyk Groups whilst filtering for only main or master branches.\
  Measure the performance of fix behaviour against SLA. For example, enter customer SLA targets and track towards those goals.\
  Build custom prevention reporting to understand shift left impact. For example, view trends in preventable Open Source vulnerabilities across all Snyk Groups, filtered by specific severities and risk scores.

## Getting started <a href="#getting-started" id="getting-started"></a>

### Request a Snowflake Data Share access <a href="#request-a-snowflake-data-share-access" id="request-a-snowflake-data-share-access"></a>

Follow the steps below to request a Snowflake Data Share access:

1. Contact your Snyk Account Executive to request access
2. Provide your Snyk contact person with the following Snowflake account details (find [here](https://docs.snowflake.com/en/user-guide/admin-account-identifier#finding-the-organization-and-account-name-for-an-account) guidelines to trace your credentials):
   * Account Name.
   * Organization Name.
   * If you prefer to limit the data share to a specific set of Snyk Groups, mention the relevant Group IDs (the ID is available in the Snyk Group Settings).
3. &#x20;After Snyk receives the Snowflake account details, the team will prepare the Data Share. You should expect to see your data within 24 hours.

### Prepare to consume Snowflake Data Shares <a href="#prepare-to-consume-snowflake-data-shares" id="prepare-to-consume-snowflake-data-shares"></a>

The data share will be provided as a **Privately Shared Listing**. If it is the first time that you consume Snowflake data shares, proceed with the steps below; otherwise, proceed to [Create a database from Snyk Data Share](https://docs.snowflake.com/en/user-guide/data-share-consumers#creating-a-database-from-a-share):

1. Your Snowflake Organization admin (requires ORGADMIN role) should [accept Snowflake Provider and Consumer Terms of Service](https://other-docs.snowflake.com/en/collaboration/consumer-becoming#accept-the-snowflake-provider-and-consumer-terms-of-service).
2. [Set up the required privileges](https://other-docs.snowflake.com/en/collaboration/consumer-becoming#set-up-required-privileges) (requires ACCOUNTADMIN role or another role with CREATE DATABASE and IMPORT SHARE privileges).

### Create a Database from the Data Share <a href="#create-a-database-from-the-data-share" id="create-a-database-from-the-data-share"></a>

To get access and be able to query the data share, follow the steps below to create a database from the data share:

1. A Snowflake user with ACCOUNTADMIN role (or a role with CREATE DATABASE and IMPORT SHARE privileges) should [view the available shares](https://docs.snowflake.com/en/user-guide/data-share-consumers#viewing-available-shares).
2. [Create a database from Snyk data share](https://docs.snowflake.com/en/user-guide/data-share-consumers#creating-a-database-from-a-share).\
   **Note:** it can take about ten minutes for the data to be provisioned and ready to use, depending on your cloud region
3. [Grant privileges to the shared database](https://docs.snowflake.com/en/user-guide/data-share-consumers#granting-privileges-on-a-shared-database).

## Data policy <a href="#data-policy" id="data-policy"></a>

### Data scope and accessibility <a href="#data-freshness" id="data-freshness"></a>

* Snyk Snowflake Data Share is scoped to the data of a requested set of Snyk Groups. A customer can request access to all Snyk groups or to specific ones.&#x20;
* Snyk shares all the data that is available for the requested Snyk Groups as exists in the Snyk database, with no additional limitations in the data share itself.
* The data share itself is provided as a read-only database and is accessible according to Snowflake standard role-based access control.

### Supported data <a href="#supported-data" id="supported-data"></a>

Snyk Data Share focuses on issues data. \
There are over 50 fields available to query, including:&#x20;

* **Issue details:** Issue Type, Issue Severity, Issue Status, CVE, CWE, First/Last Introduced, Computed Fixability, Score, Exploit Maturity, and so on.
* **Context details:** Organization Name, Project Target, Project Name, Project Origin, etc.

Review the [issue columns dictionary](issue-columns-dictionary.md) to learn more about the Snowflake data share fields that are also available in Snyk Reports.&#x20;

{% hint style="info" %}
To allow greater flexibility, the data share includes issues that are unavailable in Snyk reports:

* Deleted issues
* Issues of Projects that are not monitored

To reach the same figures as presented in Snyk reports, select only issues with:

* DELETED\_AT is NULL
* IS\_PROJECT\_MONITORED = TRUE
{% endhint %}

### Data freshness <a href="#data-freshness" id="data-freshness"></a>

The data available within the share refreshes within approximately two hours. This process occurs automatically and doesn't require any action from the customer's end.
