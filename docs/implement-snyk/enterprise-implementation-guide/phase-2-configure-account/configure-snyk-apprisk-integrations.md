# Configure Asset Management with Snyk AppRisk

## Prerequisites for Snyk AppRisk

Ensure that you meet the following prerequisites before setting up Snyk AppRisk:

* You are a Group Administrator for the Group associated with Snyk AppRisk, or you are assigned a Group level role with permissions to View Group and Edit Snyk Essentials.
* The Group associated with Snyk AppRisk includes organizations that have onboarded Snyk application security products.
* You have the necessary permissions and authority to onboard cloud-based SCM tools (Azure DevOps, GitHub, GitLab, and so on) to Snyk Snyk Essentials for repository asset discovery.

## Configure Snyk AppRisk and setup SCM integrations

Start onboarding Snyk AppRisk by identifying all inventory code-based assets and detecting which assets have security controls set up. See the [Configure Asset Management with Snyk Essentials](configure-asset-management-with-snyk-essentials.md#configure-snyk-essentials-and-setup-scm-integrations) page for more details on configuring and setting up SCM integrations.

## Access Snyk AppRisk

See the [Configure Asset Management with Snyk Essentials](configure-snyk-apprisk-integrations.md#access-snyk-apprisk) page for more details on accessing Snyk AppRisk.

## Setup integrations <a href="#setup-integrations" id="setup-integrations"></a>

Start building your asset inventory by setting up the necessary integrations.

{% hint style="info" %}
The scanned information is automatically imported within two hours after enabling all features.  Note that for large onboardings, it is suggested to wait up to 24hours for all information to be processed.
{% endhint %}

Access and configure the integrations from the **Integrations** view. Select the **Integration Hub** option to see the list of all available integrations. See the [Using the Integration Hub](../../../developer-tools/scms/group-level-integrations/#using-the-integration-hub) section for more details.

The default display in the **Integrations** view includes the configured Snyk integrations. The status of each integration, **Connected** or **Not connected**, depends on the specific content imported into Snyk.

You can [customize an existing integration](../../../getting-started/snyk-web-ui.md#edit-an-integration) or [connect a new SCM integration](../../../developer-tools/scms/organization-level-integrations/#adding-an-integration).

### SCM integrations

Navigate to the [Group-level integrations](../../../developer-tools/scms/group-level-integrations/) page for more details about the supported SCM integrations.

### Brokered SCM integration <a href="#brokered-scm-integration" id="brokered-scm-integration"></a>

Navigate to the [Snyk Broker - SCM integrations](../../../enterprise-setup/snyk-broker/using-snyk-essentials-with-snyk-broker.md) page for more details about installing and configuring Snyk AppRisk using Snyk Broker.

### Third-party integrations

Set up your third-party integrations in the [Snyk AppRisk Integrations Hub](../../../getting-started/snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors). \
\
In each Snyk Organization, administrators can distribute tokens that provide restricted access to the applications utilized by developers. \
\
For Group-level integrations in Snyk AppRisk, a token provides an overview of current assets, regardless of what is imported for security testing or monitoring.

### Application context for SCM Integrations

See the [Application context for SCM Integrations](../../../developer-tools/scms/application-context-for-scm-integrations/) docs for more details about how to use this feature.

## Features

The Snyk AppRisk functionality is split across several menu options from the Group level.&#x20;

* [Inventory](../../../manage-assets/overview.md)
* [Policies](../../../manage-risk/policies/assets-policies/)
* [Integrations for SCM](../../../developer-tools/scms/group-level-integrations/) and for [third-party](../../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md)
* [Issues](../../../manage-risk/prioritize-issues-for-fixing/)
* [Analytics](../../../manage-risk/analytics/application-analytics.md)

#### Inventory view

The Inventory feature is structured in four sections, each focused on a specific area:

* **Overview:** Provides quick insights into the discovered repositories.
* **All Assets**: All the discovered assets, grouped by their type.
* **Asset Hierarchy**: Asset Hierarchy layout shows assets in a hierarchical structure. The list of assets is sorted by issue counts, and, where applicable, the package assets are listed underneath the repositories where they are located. Assets hierarchy is visible only when there are no filters applied. Navigate to the [Assets inventory components](../../../manage-assets/assets-inventory-components.md) page for a detailed overview of all options available in the Assets Hierarchy view and to the [Filters capabilities](../../../manage-assets/assets-inventory-features.md#filters-capabilities) page for more details about the filtering options and how to use them.
* **Teams**: SCM repository assets grouped by teams. Note that only SCM organizations with teams, and repositories assigned to a team, appear on this layout.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk AppRisk.

If you are using Snyk AppRisk for the first time, Snyk recommends you to first use the Coverage filter to determine where you have Snyk implemented. Then, you can use the Coverage Gap filter to identify the assets that do not meet the coverage requirements set in a **Set coverage control** policy.
