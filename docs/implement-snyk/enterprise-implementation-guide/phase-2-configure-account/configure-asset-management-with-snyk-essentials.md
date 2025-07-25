# Configure Asset Management with Snyk Essentials

## Prerequisites for Snyk Essentials

Ensure that you meet the following prerequisites before setting up Snyk Essentials:

* You are a Group Administrator for the Group, or you are assigned a Group level role with permissions to View Group and Edit Essentials.
* The Group associated with Snyk Essentials includes organizations that have onboarded Snyk application security products.
* You have the necessary permissions and authority to onboard cloud-based SCM tools (Azure DevOps, GitHub, GitLab, and so on) to Snyk Essentials for repository asset discovery.

## Configure Snyk Essentials and setup SCM integrations

Start onboarding Snyk Essentials by identifying all inventory code-based assets and detecting which assets have security controls set up.

## Access Snyk Essentials

Verify you can access Inventory from the [Snyk Web UI.](../../../getting-started/snyk-web-ui.md)

* Access the **Inventory** menu from the Group level of your Snyk Group.
* Ensure you have Group Admin access.

## Setup integrations <a href="#setup-integrations" id="setup-integrations"></a>

Start building your asset inventory by setting up the necessary integrations.

{% hint style="info" %}
The scanned information is automatically imported within two hours after enabling all features.  Note that for large onboardings, it is suggested to wait up to 24hours for all information to be processed.
{% endhint %}

Access and configure the integrations from the **Integrations** view. Select the **Add integration** option to see the list of all available integrations. See the [Adding an integration](../../../scm-integrations/organization-level-integrations/#adding-an-integration).

The default display in the **Integrations** view includes the configured Snyk integrations. The status of each integration, **Connected** or **Not connected**, depends on the specific content imported into Snyk.

You can [customize an existing integration](../../../getting-started/snyk-web-ui.md#edit-an-integration) or [connect a new SCM integration](../../../scm-integrations/organization-level-integrations/#adding-an-integration).

<figure><img src="../../../.gitbook/assets/integration-add-integration.png" alt="Snyk Essentials - The list of available integrations"><figcaption><p>Snyk Essentials - The list of available integrations</p></figcaption></figure>

After you select **Add an integration**, a list of the available integrations is displayed. For each integration, you can add one or multiple profiles.&#x20;

The following video shows how to configure integrations for asset management and discovery.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737657000/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-3_-_v1_-_Snyk_Essentials_and_Snyk_AppRisk_Integrations.mp4" %}
Configure integrations for asset management and discovery&#x20;
{% endembed %}

### SCM integrations

Configure your SCM integrations in the **Integrations** page.

{% hint style="warning" %}
The integration interface is dedicated to asset management. It is available with Snyk Essentials and is separate from the Organization integrations interface, which is used for security and license scanning.
{% endhint %}

When you set a token on the Group Level Integrations screen, it enhances security by providing a comprehensive view of access. This token allows broad access to all repositories, even those not permitted for individual developers, ensuring alignment between security and development teams.

The supported SCM integrations are:

* [GitHub](../../../scm-integrations/group-level-integrations/github-for-snyk-essentials.md)
* [GitLab](../../../scm-integrations/group-level-integrations/gitlab-for-snyk-essentials.md)
* [Azure DevOps (Azure Repos)](../../../scm-integrations/group-level-integrations/azure-devops-for-snyk-essentials.md)
* [BitBucket](../../../scm-integrations/group-level-integrations/bitbucket-for-snyk-essentials.md)

Navigate to the [Group-level integrations](../../../scm-integrations/group-level-integrations/) page for more details about the supported SCM integrations.

### Brokered SCM integration <a href="#brokered-scm-integration" id="brokered-scm-integration"></a>

When setting up a Snyk Broker, there are some questions you need to ask regarding either standing up a new broker or updating an existing [Snyk broker connection](../../../enterprise-setup/snyk-broker/):

* Are you hitting any API Rate Limit issues?
* Do you need to update the SCM token to a user that has access to all relevant SCM repositories?
* Do you have more than 1000 repos?

If you answered Yes to any of the above questions, then you need to deploy a new Snyk Broker to accommodate the Snyk Essentials SCM connection.&#x20;

{% hint style="info" %}
Snyk recommends creating a new Organization in Snyk specifically for the Snyk Essentials Broker.
{% endhint %}

Navigate to the [Snyk Broker](../../../enterprise-setup/snyk-broker/using-snyk-essentials-with-snyk-broker.md) page for more details about installing and configuring Snyk Essentials using Snyk Broker.

### Application context for SCM Integrations

The application context refers to the information surrounding a particular application, which includes the assets that form the application and how these assets are related. Understanding this context is essential for assessing the impact and risk of security issues, as it provides a comprehensive view of the application's structure and its components. This understanding is crucial for effectively evaluating and managing application security risks.

Access the [Application context for SCM Integrations](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/) docs for more details about how to use this feature.

## Features

The Snyk Essentials functionality is split across several menu options from the Group level.&#x20;

* The Asset [Dashboard](../../../getting-started/snyk-web-ui.md#view-the-assets-dashboard) report
* [Inventory](../../../manage-assets/)
* [Policies](../../../manage-risk/policies/assets-policies/)
* [Integrations for SCM](../../../scm-integrations/group-level-integrations/)&#x20;
* [Issues](../../../manage-risk/prioritize-issues-for-fixing/)

#### Inventory view

The Inventory feature is structured in four sections, each focused on a specific area:

* **Overview:** Provides quick insights into the discovered repositories.
* **All Assets**: All the discovered assets, grouped by their type.
* **Asset Hierarchy**: Asset Hierarchy layout shows assets in a hierarchical structure. The list of assets is sorted by issue counts, and, where applicable, the package assets are listed underneath the repositories where they are located. Assets hierarchy is visible only when there are no filters applied. Navigate to the [Assets inventory components](../../../manage-assets/assets-inventory-components.md) page for a detailed overview of all options available in the Assets Hierarchy view and to the [Filters capabilities](../../../manage-assets/assets-inventory-features.md#filters-capabilities) page for more details about the filtering options and how to use them.
* **Teams**: SCM repository assets grouped by teams. Note that only SCM organizations with teams, and repositories assigned to a team, appear on this layout.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk Essentials.

If you are using Snyk Essentials for the first time, Snyk recommends you first use the Coverage filter to determine where you have Snyk implemented. Then, you can use the Coverage Gap filter to identify the assets that do not meet the coverage requirements set in a **Set coverage control** policy.

You can use the Coverage Gap filter to:

*   &#x20;Find any asset that does not comply with the Set coverage control policy requirements:&#x20;

    <figure><img src="../../../.gitbook/assets/image (1) (10).png" alt="Use the Coverage gap filter to Find any asset that is out of policy"><figcaption><p>Coverage gap - Use case 1</p></figcaption></figure>
*   Find any assets that do not meet the coverage requirements for Snyk Open Source or Snyk Code, or both of them simultaneously:&#x20;

    <figure><img src="../../../.gitbook/assets/image (1) (10) (1).png" alt="Use the Coverage gap filter to Find any asset that is out of policy for Snyk Open Source OR Snyk Code"><figcaption><p>Coverage gap - Use case 2</p></figcaption></figure>

#### Tags <a href="#hardbreak-tags" id="hardbreak-tags"></a>

Use tags to categorize the assets. You can use tags in multiple ways:

* Automatic tags: Snyk Essentials automatically tags repository assets with information about the used technologies (Python, Terraform, and so on) in the repository and repository latest updates. You can also use policies to tag repository and package assets. GitHub and GitLab topics can also be pulled from the repository and applied as Asset Tags in Snyk Essentials or Snyk AppRisk.

{% hint style="info" %}
BitBucket cannot automatically detect the language used in the source code from the repositories. In Snyk Essentials or Snyk Essentials, you can only see the language tags that have been manually added for BitBucket. For more information, see the official documentation provided by BitBucket.
{% endhint %}

* User-defined tags: set up custom tags through policies to categorize your assets beyond the system-generated tags. See the [Create policies](../../../manage-risk/policies/assets-policies/create-policies.md) page for more details.&#x20;

#### Asset Dashboard Report

The [Asset Dashboard](../../../manage-issues/reporting/available-snyk-reports.md#asset-dashboard) provides a comprehensive overview of your application and security controls. It displays essential data such as the status and trends of open issues, control coverage, and repository metadata.





