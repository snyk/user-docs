# Assets inventory tabs

Snyk defines an asset as a meaningful, real-world component in an application’s SDLC, where meaningful means either carries a risk or aggregates risk of other components (for example, repositories that contain packages), and real-world means that the concept exists outside of Snyk, for example, repository (which is a generally applicable term). In most cases, assets carry a risk or aggregate risk of other components, such as repositories that contain packages.

Snyk Essentials inventory tabs are organizing your repository assets in meaningful ways, enabling you to:

* Gain full repository asset visibility from your SCM tools, including details about configured teams and repository code committers.
* Track controls coverage for Snyk products.
* Prioritize coverage mitigation efforts according to business impact.
* Use automatic repository discovery to surface repositories that have not yet been imported into Snyk to identify coverage gaps.

{% hint style="info" %}
Each line in the inventory represents an asset.
{% endhint %}

## Inventory tabs <a href="#inventory-layouts" id="inventory-layouts"></a>

To get better context and clarity over your asset inventory, Snyk Essentials allows flexible structuring with inventory tabs. Snyk Essentials includes five inventory tabs and groups assets by different contexts. You can find all inventory tabs under the Inventory menu option at the Group level:&#x20;

* **Overview:** Provides quick insights into discovered repositories, enabling AppSec teams to effectively operationalize their program using Snyk.&#x20;
* **All Assets:** All the discovered assets are grouped by their type.&#x20;
* **Asset Hierarchy**: Shows assets in a hierarchical structure. The list of assets is sorted by issue counts, and, where applicable, the package assets are listed underneath the repositories where they are located. The Assets Hierarchy is visible only when no filters are applied.
* **Teams**: SCM repository assets are grouped by teams. On this tab, you can only see SCM Organizations with teams and repositories assigned to a team.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk Essentials.

Each inventory tab may include different counts of assets and scanned artifacts, depending on the grouping context. Otherwise, all columns and data manipulation features are the same on each inventory tab.

{% hint style="info" %}
**Release status** \
Risk factors for Snyk AppRisk are in Early Access and are available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.

**Runtime discovered** and **Runtime last seen** filters take their release status from the [runtime integration](../integrations/connect-a-third-party-integration.md) for which they provide data and are available only for Snyk AppRisk.
{% endhint %}

You can filter the information for all the inventory tabs and use any of the available filters listed on the [Assets inventory features](assets-inventory-filters.md#filters-capabilities) page.

### Inventory Overview&#x20;

The Overview tab in Snyk Inventory provides insights into the discovered repositories, highlighting key features and characteristics such as the total number of discovered repositories and the distribution of tested and not tested repositories, the number of dormant repositories or coverage details based on the asset policies.

Provides quick insights into discovered repositories, enabling AppSec teams to effectively operationalize their program using Snyk. This helps reduce coverage gaps, organize and leverage asset context, and ensure compliance with coverage policies.

#### Repositories tested

Use this widget to get an overview of all repositories discovered and the number of repositories that Snyk has not yet tested. Click the **Not Tested** section of the widget to see the full list of not-tested repositories. You can import all not-tested repositories into the correct Snyk Organization so that they can be tested.

#### Control coverage gaps

Use this widget to get a clear overview of all discovered repositories and see how many have at least one control coverage gap, as defined by an asset policy. A repository with a coverage gap is a repository that does not meet the coverage requirements set in the asset policy. The coverage gaps are automatically highlighted using a default policy applied to new Groups, helping you reduce application risk.&#x20;

Follow the next steps to remediate the coverage gaps:

1. Click "Coverage gap" to see all affected repositories.
2. Determine the reasons for the policy non-compliance.
3. Remediate and bring repositories into compliance.
4. Set up an asset policy.

#### Dormant repositories

Use this widget to see all dormant repositories with critical and high-risk issues. A dormant repository is one that has not had any commits in the past six months. Based on this information, you can decide whether to decommission or fix stale repos.

#### Languages with most issues

Use this widget to identify the programming languages that often present issues within your codebase. If you hover over any of the listed languages, you can see and access the Snyk Learn training focused on setting up, integrating, and customizing the selected language.&#x20;

#### Class A repositories with most high and critical issues

Use this widget to see a maximum of top ten high-risk Class A repositories with the biggest impact on the business (class A). This tool helps your development team identify and prioritize remediation efforts with asset context. By addressing high-risk areas promptly, you improve the stability and security of your Project, ultimately enhancing software quality.

### All Assets

The All Assets tab under the Inventory menu provides a central view of all your assets, offering a comprehensive overview of your security posture. You can access a list of your assets and customize the view to meet your needs. Select the columns that you want to be visible, use filters to refine the information, and export the details to share them with others. &#x20;

This unified view allows you to efficiently monitor assets and prioritize remediation for stronger application security.

### Asset Hierarchy

The Asset Hierarchy in Snyk Inventory organizes all assets in a structured, hierarchical format.\
Assets are sorted by issue counts, and where applicable, package assets are listed underneath the repositories where they are located.

The Asset Hierarchy is visible only when no filters are applied, allowing you to see a clear, unfiltered view of your assets and their relationships.

This layout helps in understanding the relationship between different assets and their associated issues, providing a comprehensive view of the asset landscape within your Organization.

### Teams

The Teams tab in Snyk Inventory organizes assets from SCM repositories by team. Assets are grouped here according to the teams assigned to them within the SCM organizations.

Only SCM organizations that have teams and repositories assigned to a team will appear in this layout. This helps in visualizing and managing repository assets according to team structures, making it easier to track and prioritize security efforts based on team responsibilities.

### Technology

The Technology tab in Snyk Inventory groups SCM repository assets by the technology they use, such as programming languages and frameworks. This categorization is detected and tagged by Snyk Essentials, allowing you to easily identify and manage assets based on the used technologies.&#x20;

This feature helps in understanding the technological landscape of your repositories and can be useful for prioritizing security efforts and managing risks associated with different technologies.

## Assets and their attributes

Every item listed in the inventory is considered an individual asset. Most assets are actual components of the application (code repositories, domains, endpoints, and so on), but an asset can also represent a Group (certain business unit) or even a product.&#x20;

Assets in the inventory are presented with key attributes in the following columns:

* **Asset** - The name of the repository asset, scanned artifact, and the Git remote URL, if available. Scanned artifacts are missing Git remote URLs.
* **Issue** - The number of issue counts on open assets aggregated across all relevant tools of the same severity of the asset itself and its child assets or packages. The severity level is classified into **C** (critical), **H** (high), **M** (medium), and **L** (low).
* **Controls** - A report detailing all products detected by the Snyk Essentials on a specific repository asset and all products that should be but are not covered by Snyk Essentials.
* **Tags** - You will be able to add a unique key-value tag to provide a more powerful and granular context for your assets. This attribute lets you attach specific, unique metadata to your assets, which enables  precise filtering, robust policy creation, and alignment with your internal systems.
* **Labels** -  Snyk Essentials automatically labels repository assets with information about the used technologies (Python, Terraform, and so on) in the repository, and repository latest updates. You can also use policies to label repository assets.
* **Developers** - Includes the SCM profile details for code committers to the repository asset.
* **Class** - Reflects the business criticality of the asset from A (most critical) to D (least critical), as defined by the user in the Policies view. You can manually change the class or automatically change it by applying a policy. You can lock the value you have manually set for a Class to prevent policies from overriding it.
* **Risk factors** - Lists the potential vulnerabilities and security threats associated with each asset and helps users identify specific risks, enabling them to prioritize and address issues more effectively.&#x20;
* **Source** - Reflects the source of the asset, which can come from Snyk, an SCM, or a third-party integration.
* **SCM Integrations** - Shows how each SCM was integrated at the Group or Organization level. By understanding the source of the SCM integration, you can determine if you require a Group-level integration to unlock full asset context.
* **SCM Repository freshness** - Reflects the status of the repository and the date of the last commit.
* **Clusters** - Provides a list of all the cluster names where the image asset is deployed.&#x20;
* **Organizations** - Provides a list of the Snyk Organizations that are mapped to the asset.
* **Actions** - Provides a workflow to set up an SCM integration, enriching the asset context with information such as labels, developers, and repository freshness. This use case is available when a Group-level integration is not configured.

{% hint style="info" %}
The Clusters column is hidden by default. To enable it, click Columns, select Clusters from the dropdown list, then click Apply to save the changes.&#x20;
{% endhint %}

### **Asset Sources, Types, and Scanned Artifacts**

Snyk Essentials automatically derives assets from Snyk and any SCM tools that are onboarded using the Snyk Essentials Integration. SCM tools from the Snyk Essentials Integration may add additional repositories that are not scanned by Snyk and additional contexts, such as teams and code committers.

### Repository assets, scanned artifacts and packages

#### Repository assets

Snyk Essentials supports repository assets (from main branches) as an asset type. Repository assets are visible in all inventory layouts and are supported by Policies. To avoid duplication, assets are identified using a unique identifier, which is the git remote URL for repository assets.

{% hint style="info" %}
For Snyk Essentials SCM imported repositories, archived or deleted repositories will not be displayed in the asset inventory and will not be shown in the dashboard widgets.
{% endhint %}

#### Scanned artifacts

Snyk Essentials also includes the concept of scanned artifacts. A scanned artifact is an entity detected by Snyk that cannot be identified as a repository asset because it does not include identifying information, such as a Git remote URL.&#x20;

Scanned artifacts provide users with visibility into what Snyk Essentials detects from scans but require additional troubleshooting.&#x20;

You can see the scanned artifacts in the Inventory Type view. The scanned artifacts are not supported by Policies. Furthermore, scanned artifacts may include duplicates, as identifying information is missing.

#### Packages

Packages are defined as software or libraries that are managed by package management systems.

Package assets are created when you scan the dependencies of a Project through package management systems or by using the Snyk CLI. This enables Snyk Essentials to identify and analyze the security vulnerabilities of the packages used within a Project, offering insights into possible risk exposures and providing recommendations for mitigation.
