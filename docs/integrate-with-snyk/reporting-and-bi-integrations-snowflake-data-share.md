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

Snyk Data Share for Snowflake is included within the Enterprise Plus package. Follow the steps below to join Snyk Enterprise Plus package and request a Snowflake Data Share access:

1. Contact your Snyk Account Executive about the Snyk Enterprise Plus package
2. To enable Snowflake Data Share, provide the following Snowflake account details to your Snyk contact person (see [Finding the Organization and Account Name for an Account in the Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#finding-the-organization-and-account-name-for-an-account)):
   * Account Name.
   * Organization Name.
   * If you prefer to limit the data share to a specific set of Snyk Groups or Organizations, mention the relevant Group IDs or Organization IDs (the ID is available in the Snyk Group and Organization Settings).
3. &#x20;The Snowflake account details are shared; our team will prepare the Data Share. You should expect to see your data within 24 hours.

The data share will be provided as a **Privately Shared Listing**. If it is the first time that you consume Snowflake data shares, proceed with the steps below; otherwise, proceed to [Create a database from Snyk Data Share](https://docs.snowflake.com/en/user-guide/data-share-consumers#creating-a-database-from-a-share):

1. Your Snyk Organization admin (requires ORGADMIN role) should [accept Snowflake Provider and Consumer Terms of Service](https://other-docs.snowflake.com/en/collaboration/consumer-becoming#accept-the-snowflake-provider-and-consumer-terms-of-service).
2. [Set up the required privileges](https://other-docs.snowflake.com/en/collaboration/consumer-becoming#set-up-required-privileges) (requires ACCOUNTADMIN role or another role with CREATE DATABASE and IMPORT SHARE privileges).

To get access and be able to query the data share, follow the steps below to create a database from the data share:

1. A user with ACCOUNTADMIN role (or a role granted the IMPORT SHARE global permission) should [view the available shares](https://docs.snowflake.com/en/user-guide/data-share-consumers#viewing-available-shares).
2. [Create a database from Snyk data share](https://docs.snowflake.com/en/user-guide/data-share-consumers#creating-a-database-from-a-share).\
   **Note:** it can take around ten minutes for the data to be provisioned and ready to use, depending on your cloud region
3. [Grant privileges to the shared database](https://docs.snowflake.com/en/user-guide/data-share-consumers#granting-privileges-on-a-shared-database).

The current version of Snyk Data Share focuses on the issues data. There are over 40 fields to browse, including **Issue details:** Issue Type, Issue Severity, Issue Status, CVE, CWE, First/Last Introduced, Computed Fixability, Score, Exploit Maturity, and so on.

The data available within the share refreshes within approximately two hours.
