# Assets inventory components

Snyk presents each inventory layout in a table format, detailing the available key attributes:

* [Assets](assets-inventory-components.md#asset)
* [Issues](assets-inventory-components.md#issues)
* [Coverage Controls](assets-inventory-components.md#coverage-controls)
* [Tags](assets-inventory-components.md#tags)
* [Labels](assets-inventory-components.md#labels)
* [Developers](assets-inventory-components.md#developers)
* [Class](assets-inventory-components.md#class)
* [Risk factors](assets-inventory-components.md#risk-factors)
* [Source](assets-inventory-components.md#source)
* [SCM Repository freshness](assets-inventory-components.md#scm-repository-freshness)
* [Clusters](assets-inventory-components.md#clusters)
* [Organizations](assets-inventory-components.md#organizations)
* [Visibility](assets-inventory-components.md#visibility)

## Asset

Assets in Snyk Essentials are meaningful, real-world components in an application’s SDLC. The following asset types are available:

* Repository assets [**`Billable`**](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/how-does-snyk-count-assets#billable-assets)
* Container images [**`Billable`**](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/how-does-snyk-count-assets#billable-assets)
* Packages
* Scanned artifacts

An asset can be the parent of multiple items. For example, a repository asset usually contains one or more package assets.

The Asset column incorporates the name of the repository asset, package, scanned artifact, or container images. Click the arrow next to the parent asset name to expand the list of all contained items.

You can copy the name of an asset or browse the repository. Each asset has a menu at the end of the row. Click the menu, then select **Copy** to copy the URL or **Browse** to navigate to the asset repository.

### Repository assets

Repository assets represent SCM repositories. A repository asset is created by discovering the repositories directly in the SCM, when such integration is configured. Alternatively, scanning a repository can create a repository asset (by Snyk or third-party tools) as long as the scanned code is identified with a specific repository (in Snyk, this means filling in the `gitRemoteURL` parameter).

If you scan the code locally using CLI, with no association to a repository, then Snyk does not create a repository asset. For more details about CLI commands, see [Scanning methods](../scan-with-snyk/snyk-essentials.md#scanning-methods).

### Container Image assets

You can identify a container image based on the Image ID. If multiple container images have the same Image ID, then Snyk generates only one image asset for that Image ID, enriched with information from all the identified container images for that ID.

Snyk Essentials retrieves all image assets from Snyk Container. Reimport the images to ensure you scan the latest image. If you run a new scan on a Project that contains image assets, it rescans the same image for new vulnerabilities. To identify new image assets, you need to first reimport, and then scan the Project. Check the [Detect application vulnerabilities in container images](../scan-with-snyk/snyk-container/use-snyk-container/detect-application-vulnerabilities-in-container-images.md) page for more details.

### Packages

Packages in Snyk Essentials are defined as software or libraries that are managed by package management systems.

Snyk creates package assets when you scan the dependencies of a Project through package management systems or by using the Snyk CLI. This lets Snyk Essentials identify and analyze the security vulnerabilities of the packages used in a Project, offering insights into possible risk exposures and providing recommendations for mitigation.

### Scanned artifacts

A scanned artifact in Snyk Essentials is an entity detected by Snyk that cannot be identified as a repository asset because it does not include identifying information, such as a Git remote URL.

Scanned artifacts provide users with visibility into what Snyk Essentials detects from scans but require additional troubleshooting.

You can find scanned artifacts in the Inventory Type view, but they are not supported by Policies. Scanned artifacts may include duplicates due to missing identifying information.

## Asset tabs

The asset information is divided into the following tabs:

* [Summary](assets-inventory-components.md#summary)
* [Related Assets](assets-inventory-components.md#related-assets)
* [Related Projects](assets-inventory-components.md#related-projects)
* [Attributes](assets-inventory-components.md#attributes)

### Summary

The Summary tab is a concentrated view of asset properties. The Summary screen presents you with the following information:

* **Info**
  * Class - specifies the business criticality of the asset.
  * Source - specifies the origin of the asset.
  * Visibility - lists the visibility status of the repositories.
  * Risk factors - provides the list of active risk factors.
  * SCM Repository freshness - provides the current status of your repositories, including the date of the last commit.
* **Organization** - specifies the Organizations associated with the asset.
* **Labels** - provides the list of all labels available for that asset.
* **Tags** - provides a key-value pair that lets you attach structured metadata to your assets.
* **Issues** - categorizes the identified types of open issues.
* **App Context**\* - asset metadata from App Context integrations, such as Backstage catalog or ServiceNow CMDB, can include the following attributes: catalog name, category, application, owner, and so on.

\*App Context information is visible only when the asset is part of a Project for which the application context was configured.

{% hint style="info" %}
After you apply the filters, the assets list displays only the assets that directly match the filter conditions, and, if available, a list of children assets related to the selected one, with the information shown in a table format, with a focus on the following topics: Asset (name), Issues, Controls, Class.
{% endhint %}

### Related Assets

The Related assets tab provides a detailed view of assets related to the selected one. Use this tab to assess scanning coverage or asset ownership. You can see the details of a related asset by clicking one of them. Usually, these are Package assets. When looking at Related Assets, you can notice a link to the parent repository. If you click the parent asset link, you revert to the initial view of the parent asset.

### Related Projects

The Related Projects tab provides a collection of Snyk Projects that are associated with a specific asset in the platform. Snyk arranges these projects in a table format, letting you view relevant information that assists in managing and assessing vulnerabilities related to the asset. Each Project is displayed with the following details:

* **Projects by Target**: A list of Projects grouped by targets. You can see both the Project name and the target name under which the Project is grouped.
* **Target Reference**: An optional identifier that may not always be available.
* **Test Surface**: Indicates the source of the Project scan, specifying whether it originated from SCM or the CLI.
* **Issues**: Provides insight into the number and severity of identified issues in the Project.
* **Organization**: Displays the Snyk Organization to which the Project belongs.
* **Tested**: Shows the relative time since the last scan (for example, "3 hours ago") along with a tooltip that reveals the full date and time upon hovering.

‌Projects are sorted by Target, Target Reference, and Tested date. This makes it easy to find related Projects to monitor and fix.

### Attributes

The Attributes tab shows miscellaneous attributes, like the Asset ID or Asset Type, that Snyk fetches from the data source but that do not have a dedicated column. The benefit of having this information is not only in presenting it but mostly in making it searchable. You can search for an attribute using either the inventory search bar or the filters.

## Issues

The Issues column presents a comprehensive list of issues identified in your assets. These findings result from scans that Snyk and internal tools you have deployed perform. This detailed list helps you understand the security posture of your assets and prioritize remediation efforts based on the severity and impact of each issue. With visibility into these issues, you can take proactive steps toward improving the overall security of your applications and infrastructure.

Most of the issues are mapped to an asset. However, some of the issues are associated with an asset but not directly linked to it. This is the case with image assets.

You can see these assets under the **Inventory** view, and they also appear under the **Asset and Source Code** column from the **Insights UI** view.

The **Issues** column from the Asset view presents an aggregated count of open issues. Snyk categorizes these counts based on the severity level of the issues found in assets, their children's assets, or associated packages. The severity is divided into four distinct levels:

* **C** (Critical): Issues that represent a serious threat and should be addressed immediately to prevent potential exploits or major disruptions.
* **H** (High): These are significant issues that, while not immediately dangerous, could potentially lead to critical vulnerabilities if not resolved in a timely manner.
* **M** (Medium): Issues of medium severity might not pose an immediate threat but are still important to fix as part of regular maintenance to improve overall security and functionality.
* **L** (Low): These are considered minor issues that have a low impact on the security of the system and operation but should still be addressed to maintain optimal performance and prevent future vulnerabilities.

This classification streamlines prioritization, helping you focus on critical areas and optimize remediation.

## Coverage Controls

The Controls column displays all of the Snyk products that ran on a specific repository asset. This column displays, in circles, a logo for each Snyk product. The logo icon itself indicates the highest severity of issues from this source. For example, if the highest severity issue is **C** (critical), you can see a red dot on the control icon.

The Controls logos can have one of the following states:

| Logo                                                                       | Description                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| <img src="../.gitbook/assets/image (132).png" alt="" data-size="original"> | The Snyk product was executed.                                                   |
| <img src="../.gitbook/assets/image (133).png" alt="" data-size="original"> | The Snyk product was executed but with issues.                                   |
| <img src="../.gitbook/assets/image (134).png" alt="" data-size="original"> | The Snyk product should have been executed but was not executed.                 |
| <img src="../.gitbook/assets/image (135).png" alt="" data-size="original"> | The Snyk product was executed and failed.                                        |
| <img src="../.gitbook/assets/image (136).png" alt="" data-size="original"> | The Snyk product was executed and failed with issues.                            |
| <img src="../.gitbook/assets/image (137).png" alt="" data-size="original"> | The Snyk product was executed and failed due to not being covered by the policy. |

Click a Controls logo to see **Last test** details and the **Issues** count, split by severity. This reflects the most recent time that the asset was scanned by a specific product.

## Tags <a href="#tags" id="tags"></a>

Provides a key-value tagging capability that lets you attach specific, structured metadata to your assets. This feature lets you use granular filtering, robust policy creation, and better alignment with your internal systems.

**Example**: A structured tag provides both a key and a value, such as `platform:aws` or `region:eu-central-1`.

## How to filter assets by tags <a href="#how-to-filter-the-projects-listing-by-tags" id="how-to-filter-the-projects-listing-by-tags"></a>

In Snyk Web UI, you can filter assets by their tags using **Advanced filters**. You can define filters based on specific criteria, such as a property of an asset, a condition, and a value.

* Filter by `Tags`: The new `Tags` filter is a key-value pair filter. This filter lets you select a specific tag key such as `department` and then choose a corresponding value such as `finance` to narrow down the asset list.

## Labels

Asset labels are metadata that is applied to repository assets and build artifacts. You can use asset labels to tag based on predefined values, manage and apply security policies, and group assets based on common characteristics. The following asset types are available:

* **GitHub custom properties** - lists the GitHub custom properties associated with your GitHub repository as a label.
* **User-defined labels** are customizable, as you can define their logic through Assets Policies. For example, you can set labels to represent a repository that comes from a specific source, such as GitHub. Labels associated with assets are identified in the UI with the **Asset policy label's** name.
* **System labels** are automatically assigned by Snyk based on asset names or detected keywords (for example, `codeowners`).

You can add a repository asset label through Policies, or Snyk Essentials can system-generate it to provide more context. Click a labels field to view all labels.

{% hint style="info" %}
BitBucket Cloud cannot automatically detect the language used in the source code from the repositories. In Snyk Essentials, you can only see the language labels that have been manually added for BitBucket Cloud.\
Language data is not available for BitBucket Server.\
For more information, you can refer to the official documentation provided by BitBucket.
{% endhint %}

A system-generated label includes the following information:

* **Technology** - The languages detected by Snyk Essentials in the source code in a repository asset.
* **SCM Topic** - The topics found in the integrated SCM repositories. Snyk Essentials supports topics from GitHub and GitLab.
* **Asset type label** - The label explaining the type of the asset. For example, ‌Snyk assigns container assets an image asset label.
* **SCM Repository freshness** - The status of the repository and the date of the last commit.
  * **Active**: Had commits in the last 3 months.
  * **Inactive**: The last commits were made in the last 3 - 6 months.
  * **Dormant**: No commits in the last 6 months.
  * **N/A**: There are no commits detected by Snyk Essentials.

### Labels rules overview

Labels are organized into three main categories:

* GitHub custom properties
* Asset policy labels
* System labels

System labels are automatically generated from the SCM repositories. System labels can be classified into three main categories:

* Languages:
  * This applies to GitHub, GitLab, Azure DevOps, and BitBucket as long as the data is available in the repository.
  * GitHub, GitLab, and Azure DevOps have automated language detection. Instead, BitBucket requires users to set up the language in their repositories.
* SCM Topic:
  * This applies to GitHub and GitLab.
* Multiple different rules based on the words Snyk found in the repositories:
  * This applies to GitHub, GitLab, Azure DevOps, and BitBucket.

### Labeling policy

You can use pre-defined system labels and asset labels to mark the repositories that meet your filter criteria. Check the following [Labeling policy](../manage-risk/policies/assets-policies/use-cases-for-policies/tagging-policy.md) use case.

### Labeling rules related to metadata

| Rule                                                                  | Label             |
| --------------------------------------------------------------------- | ----------------- |
| Snyk Essentials found technologies in use.                            | `< technologies>` |
| Snyk Essentials found languages from the SCM.                         | `<languages>`     |
| Snyk Essentials detected a new repository created in the last 7 days. | `new repository`  |
| Snyk Essentials found the code Project with the code owner.           | `codeowners`      |

In the Snyk web interface, you can filter assets by their labels using the **Advanced Filters** option. You can define filters based on highly specific criteria, such as a property of an asset, a condition, and a value.

* Filter by `labels`: This filter lets you select a specific label.

## Developers

You can see the list of all the developers that worked on that specific asset. The details list includes the SCM profile details for code committers to the repository asset.

## Class

Reflects the business criticality of the asset from A (most critical) to D (least critical), as you defined it in the Policies view.

You can manually change the business criticality of an asset. To update it, click the criticality level and select another one from the list.

After manually setting the value of a class, you can lock the value to prevent a policy that has Set Asset Class as an action from overriding it. You can lock the value from the general or summary views of an asset. You can unlock the class value at any time by clicking the lock icon. A pop-up appears, asking you to confirm unlocking the value.

The Asset Class column is available on the Insights UI for risk-based prioritization, and it has the same functionality as it does here. At the moment, the Asset Class column is available only for repository assets, and applicable only for Snyk Code.

{% hint style="info" %}
Synchronization between the Asset Class and the Insights UI can take up to 3 hours.
{% endhint %}

Policies can auto-generate the class value. Create a policy that has **Set Asset Class** as an action.

## Risk factors

The Risk Factors column lists the potential vulnerabilities and security threats associated with each asset. These risk factors help you identify specific risks, letting you prioritize and address issues more effectively. By understanding the particular risks tied to your assets, you can take more informed remedial actions.

Here is a list of the available risk factors:

* [Deployed](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors/risk-factor-deployed.md)
* [OS Condition](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors/risk-factor-os-condition.md)
* [Public facing](../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors/risk-factor-public-facing.md)

## Source

The Source column in Snyk Essentials helps you identify the origin of your assets, which can be directly from Snyk, through SCM systems, or using third-party integrations. This feature simplifies asset management and risk prioritization by providing clear visibility into the origin of each asset, and it supports more effective security strategies and remediation efforts.

## SCM Integrations

The SCM Integrations column indicates how each SCM was integrated into Snyk at the Group or Organization level. Full context enrichment is available at the Group level, while testing is available at the Organization level.

The column is hidden by default, and you can enable it in the **Columns** section. It shows the following value:

* **Snyk Org**: The Snyk Organization-level integration is used for import and testing.
* **Snyk Essentials**: The Snyk Group-level integration is used for discovery and asset enrichment.

## SCM Repository freshness

The SCM Repository freshness column provides you with an immediate understanding of the current status of your repositories, including the date of the last commit. This helps you quickly identify active and dormant Projects and make decisions regarding maintenance, security patching, and resource allocation.

The repository freshness displays the repository status according to the last commit date:

* **Active**: Had commits in the last 3 months.
* **Inactive**: The last commits were made in the last 3 - 6 months.
* **Dormant**: No commits in the last 6 months.
* **N/A**: Commits data is unavailable.

## Clusters

The Clusters column lists all cluster names where an image is deployed and uses the runtime integrations as the source of the information. When you remove an image from a cluster, the cluster name is also deleted from the collection. Clusters are also available under Filters and let you filter assets in the Inventory view or create policies in the Policies view.

{% hint style="info" %}
The Cluster column is populated only when you use the Snyk Runtime Sensor.
{% endhint %}

## Organizations

The Organizations column lists all Snyk Organizations associated with each asset. This includes the names of Organizations that contain Projects linked to the asset, letting you filter and organize your asset inventory based on your organizational structures. Organizations are also available under Filters and let you filter assets in the Inventory view or create policies in the Policies view.

## Visibility

The Visibility column lists the visibility status of the repositories as follows:

* **Public**: Repositories that are publicly accessible.
* **Private**: Restricted repositories.
* **Internal**: Internal repositories specific to GitHub and GitLab.
* **N/A**

Use this metadata to prioritize risk and apply visibility-based coverage controls. The column is unavailable for image assets and is excluded from [report filters](../manage-risk/analytics/reports-tab/#snyk-reporting-filters).

## Actions

The Actions column provides a workflow to set up the SCM integration at the Group level to access full context enrichment. To identify the type of integration, check the [SCM Integrations column](assets-inventory-components.md#scm-integrations). This use case applies where a Group-level integration is not configured.

If a Group-level integration is not set up, repositories discovered at the Organization level display **Set up integration** under the **Actions** column. If you set up the integration at the Group level, this option becomes unavailable.

To add context enrichment, find an asset and select **Set up integration**. For configuration details, see [Snyk SCM Integrations](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/integrations/scm-integrations/organization-level-integrations).
