# Reporting and BI Integrations: Snowflake Data Share

{% hint style="warning" %}
**Release status** \
Snowflake Data Share is in [Closed Beta](https://docs.snyk.io/getting-started/snyk-release-process#closed-beta) and available only for customers:

* Under the Enterprise plan. For more information, see [Plans and pricing](https://snyk.io/plans).
* With the Enterprise Plus package. For more information, contact your Snyk Account Executive.
* With an active Snowflake account.

Customers can try the integration for a 60-day beta period starting from their initial access date. For access beyond 60 days, the customer is expected to be on the Enterprise Plus SKU. During beta, the data share will be limited to a capacity of 10 million records per share.&#x20;
{% endhint %}

With the new Snowflake Data Share integration, your data science, BI and AppSec teams can securely access the same underlying data available in Snyk Reporting, but within your own Snowflake account, unlocking powerful new analytical tools to better understand and visualize Snyk data.

Use this integration to enable teams to rapidly build exploratory and custom analytics using the ecosystem of tools that Snowflake supports. Customers can connect Snyk data to BI tools like PowerBI, Tableau, and Looker data studio or build custom Streamlit apps.

Having Snyk datasets directly in your Snowflake account opens doors for combining Snyk data with additional data sources, which contributes a great deal to the organization-level holistic AppSec view.

## What is Snowflake Data Share? <a href="#what-is-snowflake-data-share" id="what-is-snowflake-data-share"></a>

[Snowflake Data Share](https://docs.snowflake.com/en/user-guide/data-sharing-intro.html) allows companies to provide data to their customers in a secure, simple and quick manner. With Snowflake data shares, the data is not exchanged between the accounts; instead, the data consumer is granted with read-only access to the shared database through their own Snowflake account.

### Main use-cases <a href="#main-use-cases" id="main-use-cases"></a>

Snyk Snowflake Data Share can be used for various use-cases and can answer countless security and business-related questions. Some of the these use-cases include:

* Enhance the AppSec posture visibility for the CISO and management team
  * Streamline Snyk’s data to your BI platforms and existing security dashboards
  * Reflect performance metrics and KPIs, such as: MTTR, SLA compliance, remediation trend, etc.
* Answer specific questions or surface unique insights
  * Better understand risk **exposure** trends:
    * Example: Track total issues above a specific risk score only affecting certain project collections or tags across all Snyk groups, while filtering for only main/master branches
  * Measure performance of **fix** behavior against SLA
    * Example: Enter custom SLA targets and track towards those goals
  * Build custom **prevention** reporting to understand shift left impact:
    * Example: View trends in preventable Open Source vulnerability across all Snyk groups filtered by specific severities and risk scores

## Getting Started <a href="#getting-started" id="getting-started"></a>

### Request a Snowflake Data Share access <a href="#request-a-snowflake-data-share-access" id="request-a-snowflake-data-share-access"></a>

Snyk Data Share for Snowflake is included within the Enterprise Plus package. Follow the steps below to join Snyk Enterprise Plus package and request a Snowflake Data Share access:

1. Contact your Snyk Account Executive about the Snyk Enterprise Plus package
2. To enable Snowflake Data Share, provide the following Snowflake account details to your Snyk contact person (see [Finding the Organization and Account Name for an Account in the Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#finding-the-organization-and-account-name-for-an-account)):
   * Account Name.
   * Organization Name.
   * If you prefer to limit the data share to a specific set of Snyk Groups or Organizations, mention the relevant Group IDs or Organization IDs (the ID is available in the Snyk Group and Organization Settings).
3. &#x20;The Snowflake account details are shared; our team will prepare the Data Share. You should expect to see your data within 24 hours.

### Prepare to consume Snowflake Data Shares <a href="#prepare-to-consume-snowflake-data-shares" id="prepare-to-consume-snowflake-data-shares"></a>

The data share will be provided as a **Privately Shared Listing**. If it is the first time that you consume Snowflake data shares, proceed with the steps below; otherwise, proceed to [Create a database from Snyk Data Share](https://docs.snowflake.com/en/user-guide/data-share-consumers#creating-a-database-from-a-share):

1. Your Snyk Organization admin (requires ORGADMIN role) should [accept Snowflake Provider and Consumer Terms of Service](https://other-docs.snowflake.com/en/collaboration/consumer-becoming#accept-the-snowflake-provider-and-consumer-terms-of-service).
2. [Set up the required privileges](https://other-docs.snowflake.com/en/collaboration/consumer-becoming#set-up-required-privileges) (requires ACCOUNTADMIN role or another role with CREATE DATABASE and IMPORT SHARE privileges).

### Create a Database from the Data Share <a href="#create-a-database-from-the-data-share" id="create-a-database-from-the-data-share"></a>

To get access and be able to query the data share, follow the steps below to create a database from the data share:

1. A user with ACCOUNTADMIN role (or a role granted the IMPORT SHARE global permission) should [view the available shares](https://docs.snowflake.com/en/user-guide/data-share-consumers#viewing-available-shares).
2. [Create a database from Snyk data share](https://docs.snowflake.com/en/user-guide/data-share-consumers#creating-a-database-from-a-share).\
   **Note:** it can take around ten minutes for the data to be provisioned and ready to use, depending on your cloud region
3. [Grant privileges to the shared database](https://docs.snowflake.com/en/user-guide/data-share-consumers#granting-privileges-on-a-shared-database).

## Data Policy <a href="#data-policy" id="data-policy"></a>

### Supported Data <a href="#supported-data" id="supported-data"></a>

The current version of Snyk Data Share focuses on the issues data. There are over 40 fields to browse, including:&#x20;

* **Issue details:** Issue Type, Issue Severity, Issue Status, CVE, CWE, First/Last Introduced, Computed Fixability, Score, Exploit Maturity, and so on.
* **Context details:** Org Name, Project Target, Project Name, Project Origin, etc.

### Data freshness <a href="#data-freshness" id="data-freshness"></a>

The data available within the share refreshes within approximately two hours.
