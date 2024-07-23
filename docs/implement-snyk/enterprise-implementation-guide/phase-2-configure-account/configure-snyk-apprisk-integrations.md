# Configure Snyk AppRisk Integrations

{% hint style="info" %}
Snyk AppRisk Essentials is included in Snyk Enterprise.&#x20;
{% endhint %}

## Prerequisites for Snyk AppRisk Essentials

Ensure that you meet the following prerequisites before setting up Snyk AppRisk Essentials:

* You are a Group Administrator for the Group associated with Snyk AppRisk, or you are assigned a Group level role with permissions to View Group and Edit AppRisk.
* The Group associated with Snyk AppRisk includes organizations that have onboarded Snyk application security products.
* You have the necessary permissions and authority to onboard cloud-based SCM tools (Azure DevOps, GitHub, GitLab, and so on) to Snyk AppRisk for repository asset discovery.

## Configure Snyk AppRisk and setup SCM integrations

Start onboarding AppRisk by identifying all inventory code-based assets and detecting which assets have security controls set up.

## Access Snyk AppRisk

You can access Snyk AppRisk from the [Snyk Web UI.](../../../getting-started/snyk-web-ui.md)

* Access Snyk AppRisk from the Group level of your Snyk Group.
* Ensure you have Group Admin access.
* Access Group that has Snyk AppRisk enabled.

## Setup integrations <a href="#setup-integrations" id="setup-integrations"></a>

After you ensure you can correctly access Snyk AppRisk, you can start to build your asset inventory by setting up the integrations.

{% hint style="info" %}
The scanned information is automatically imported within two hours after enabling all features.&#x20;
{% endhint %}

You can access and configure the integrations from the Integrations view. Select the Integration Hub option to see the list of all available integrations. You can find more details about integration configuration in the [Using the Integration Hub](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#using-the-integration-hub) section.

The default display in the Integrations view includes the configured Snyk integrations. The status of each integration, **Connected** or **Not connected**, depends on the specific content imported into Snyk.

The integrations view can be configured to apply to your needs, meaning that you can [customize an existing integration](../../../getting-started/snyk-web-ui.md#edit-an-integration) or [connect a new SCM integration](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#snyk-apprisk-integrations-ecosystem).

<figure><img src="../../../.gitbook/assets/image (357) (1).png" alt="Snyk AppRisk - Integration Hub option displaying the list of available integrations"><figcaption><p>Snyk AppRisk - Integration Hub option displaying the list of available integrations</p></figcaption></figure>

After you click on the Integration Hub, a list of the available integrations is displayed. For each integration, you can add one or multiple profiles.&#x20;

### SCM integrations

Use the Snyk AppRisk Integrations Hub to configure your SCM integrations.&#x20;

{% hint style="danger" %}
Integration Hub is a distinct integration interface dedicated to Snyk AppRisk, separate from the Organization integrations interface.
{% endhint %}

The integrations at the organization level that facilitate Snyk Code, Snyk Open Source, Snyk Container, and Snyk IaC scanning are typicallly maintained or restricted by or for development. When you set a token at the group level integrations screen, it is intended for security to have a comprehensive view. This token provides broad access to all repositories that developers may not have access to. In a way, you are checking if the security perspective on assets aligns with that of the development teams.

The supported SCM integrations are:

* [GitHub](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md#group-level-snyk-apprisk-integrations)
* [GitLab](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab.md#group-level-snyk-apprisk-integrations)
* [Azure DevOps (Azure Repos)](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/azure-repositories-tfs.md#group-level-snyk-apprisk-integrations)
* [BitBucket](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-cloud.md#group-level-snyk-apprisk-integrations)

Navigate to the [Snyk SCM Integrations](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#snyk-apprisk-integrations-ecosystem) page for more details about the supported SCM integrations.

### Brokered SCM integration <a href="#brokered-scm-integration" id="brokered-scm-integration"></a>

When setting up a Snyk Broker, there are some questions you need to ask regarding either standing up a new broker or updating an existing [Snyk broker connection](https://docs.snyk.io/enterprise-setup/snyk-broker):

* Are you hitting any API Rate Limit issues?
* Do you need to update the SCM token to a user that has access to all relevant SCM repositories?
* Do you have more than 1000 repos?

If you answered Yes to any of the above questions, then you need to deploy a new Snyk Broker to accommodate the Snyk AppRisk SCM connection.&#x20;

{% hint style="info" %}
Snyk recommends creating a new Organization in Snyk specifically for the Snyk AppRisk Broker.
{% endhint %}

Navigate to the [Snyk Broker - AppRisk](../../../enterprise-configuration/snyk-broker/snyk-broker-apprisk.md) page for more details about installing and configuring Snyk AppRisk using Snyk Broker.\


### Third-party integrations

{% hint style="info" %}
The third-party integrations are mostly available only for Snyk AppRisk Pro. As a Snyk AppRisk Essentials user, you can use the third-party integrations marked as available for Snyk AppRisk Essentials, too.
{% endhint %}

To set up your third-party integrations, you can use the [Snyk AppRisk Integrations Hub](../../../getting-started/snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors). \
\
In each Snyk Organization, administrators can give out tokens that provide restricted access to the applications utilized by developers. \
\
With regards to Snyk AppRisk, the purpose of a token is to provide an overview of the current assets as compared to what is imported into Snyk.

In each Snyk Organization, administrators have the capability to provide tokens with limited access to the applications utilized by developers.

The scope of the token used in Snyk AppRisk is to provide an overview of the existing assets compared to what is imported into Snyk.

The supported third-party integrations are:

* Veracode
* Nightfall

### Backstage file for SCM Integrations

Backstage is a service catalog that allows users to add metadata or annotations to their repositories, helping to organize and categorize the available resources for easier navigation and understanding. You can leverage your SCM integration to pull metadata associated with Backstage catalog files into Snyk AppRisk.

Access the [Application context for SCM Integrations](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations.md) docs for more details about how to use this feature.

{% hint style="info" %}
After setting up all the Snyk AppRisk integrations, depending on the number of repositories, results may take up to a day to appear.
{% endhint %}

## Features

The Snyk AppRisk functionality is split across several menu options from the Group level.&#x20;

* [Dashboard](../../../getting-started/snyk-web-ui.md#view-the-assets-dashboard)
* [Inventory](../../../manage-assets/)
* [Policies](../../../manage-risk/policies/assets-policies/)
* [Integrations for SCM](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) and for [third-party](../../../integrate-with-snyk/#integrations-for-snyk-apprisk)
* [Issues](../../../manage-risk/prioritize-issues-for-fixing/)

#### Inventory view

The Inventory feature is structured in four sections, each focused on a specific area:

* **All Assets**: All the discovered assets, grouped by their type.
* **Asset Hierarchy**: Asset Hierarchy layout shows assets in a hierarchical structure. The list of assets is sorted by issue counts, and, where applicable, the package assets are listed underneath the repositories where they are located. Assets hierarchy is visible only when there are no filters applied. Navigate to the [Assets inventory components](../../../manage-assets/assets-inventory-components.md) page for a detailed overview of all options available in the Assets Hierarchy view and to the [Filters capabilities](../../../manage-assets/assets-inventory-features.md#filters-capabilities) page for more details about the filtering options and how to use them.
* **Teams**: SCM repository assets grouped by teams. Note that only SCM organizations with teams, and repositories assigned to a team, appear on this layout.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk AppRisk.

If you are using Snyk AppRisk for the first time, Snyk recommends you to first use the Coverage filter to determine where you currently have Snyk implemented. Then, you can use the Coverage Gap filter to identify the assets that do not meet the coverage requirements set in a **Set coverage control** policy.

You can use the Coverage Gap filter to:

*   &#x20;Find any asset that does not comply with the Set coverage control policy requirements:&#x20;

    <figure><img src="../../../.gitbook/assets/image (1) (10).png" alt="Use the Coverage gap filter to Find any asset that is out of policy"><figcaption><p>Coverage gap - Use case 1</p></figcaption></figure>
*   Find any assets that do not meet the coverage requirements for Snyk Open Source or Snyk Code, or both of them simultaneously:&#x20;

    <figure><img src="../../../.gitbook/assets/image (1) (10) (1).png" alt="Use the Coverage gap filter to Find any asset that is out of policy for Snyk Open Source OR Snyk Code"><figcaption><p>Coverage gap - Use case 2</p></figcaption></figure>

#### Tags <a href="#hardbreak-tags" id="hardbreak-tags"></a>

You can use tags to categorize the assets. You can use tags in multiple ways:

* Automatic tags: Snyk AppRisk automatically tags repository assets with information about the used technologies (Python, Terraform, and so on) in the repository and repository latest updates. You can also use policies to tag repository and package assets. GitHub and GitLab topics can also be pulled from the repository and applied as Asset Tags in Snyk AppRisk.

{% hint style="info" %}
BitBucket cannot automatically detect the language used in the source code from the repositories. In Snyk AppRisk you can see only the language tags that have been manually added for BitBucket. For more information, see the official documentation provided by BitBucket.
{% endhint %}

* User-defined tags: you can set up custom tags through policies to categorize your assets beyond the system-generated tags. See the [Create policies ](../../../manage-risk/policies/assets-policies/create-policies.md)page for more details.&#x20;

#### Dashboard

You can use the dashboard for a quick overview of your application and security controls. Use the default widgets and customize the displayed information as needed, or add new widgets that meet your needs. See the [Assets dashboard ](../../../getting-started/snyk-web-ui.md#view-the-assets-dashboard)page for more details.

Here are the available dashboard widgets:

* **SAST coverage**: check which repositories are being covered or not by Snyk Code and Snyk Infrastructure as Code.&#x20;

{% hint style="info" %}
The SAST coverage widget uses an OR statement, meaning that a repository is covered for SAST if it is also covered for Snyk Code OR Snyk Infrastructure as Code.
{% endhint %}

* **SCA coverage**: check which repositories are being covered or not Snyk Open Source and Snyk Container. You are able to edit the widget if you want to see either Snyk Open Source coverage or Snyk Container coverage.&#x20;

{% hint style="info" %}
The SCA coverage widget uses an OR statement, meaning that a repository is covered for SCA if it is also covered for Snyk Open Source OR Snyk Container.
{% endhint %}

* **Repository breakdown by source**: check which repositories Snyk AppRisk discovered using the SCM integrations (Azure DevOps(Azure Repos), Gitlab, GitHub, BitBucket). The Others categories are repositories that Snyk discovered but did not correlate back to a SCM repository.
* **Technology breakdown**: check the top technology (language) tags of the repositories that Snyk discovered.
* **Asset breakdown by type**: check if the asset is a repository or a package.
* **Repository activity**: check if the repository is active, inactive, or dormant.





