# Assets inventory layouts

Snyk defines an asset as a meaningful, real-world component in an application’s SDLC, where meaningful means either carries a risk or aggregates risk of other components (for example, repositories that contain packages), and real-world means that the concept exists outside of Snyk, for example, repository (which is a generally applicable term). In most cases, assets carry a risk or aggregate risk of other components, for example, repositories that contain packages.

Snyk AppRisk inventory layouts are organizing your repository assets in meaningful ways, enabling you to:

* Gain full repository asset visibility from your SCM tools, including details about configured teams and repository code committers.
* Track controls coverage for Snyk products.
* Prioritize coverage mitigation efforts according to business impact.

{% hint style="info" %}
Each line in the inventory represents an asset.
{% endhint %}

## Inventory Layouts <a href="#inventory-layouts" id="inventory-layouts"></a>

To get better context and clarity over your asset inventory, Snyk AppRisk allows flexible structuring with inventory layouts. Snyk AppRisk includes four inventory layouts and groups assets by different contexts. You can find all inventory layouts under the Inventory menu option at the Group level:&#x20;

* **All Assets**: All the discovered assets, grouped by their type.&#x20;
* **Asset Hierarchy**: Asset Hierarchy layout shows assets in a hierarchical structure. The list of assets is sorted by issue counts, and, where applicable, the package assets are listed underneath the repositories where they are located. Assets hierarchy is visible only when there are no filters applied.
* **Teams**: SCM repository assets grouped by teams. Note that only SCM organizations with teams, and repositories assigned to a team, appear on this layout.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk AppRisk.

Each inventory layout may include different counts of assets and scanned artifacts, depending on the grouping context. Otherwise, all columns and data manipulation features are the same on each inventory layout.

The following video presents an overview of the Inventory view from the Snyk Web UI.

{% embed url="https://www.youtube.com/watch?v=NwF1a08dwjk" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)!
{% endembed %}

{% hint style="info" %}
**Release status** \
**Risk factors** for Snyk AppRisk Pro are in Early Access and are available only for Snyk Enterprise plans with Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.

**Runtime discovered** and **Runtime last seen** filters take their release status from the [runtime integration](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md) for which they provide data and are available only for Snyk AppRisk Pro.
{% endhint %}

You can filter the information for all the inventory layouts and use any of the available filters listed on the [Assets inventory features](assets-inventory-features.md#available-filters) page.

## Assets and their attributes

Every item listed in the inventory is considered an individual asset. Most assets are actual components of the application (code repositories, domains, endpoints, and so on), but an asset can also represent a Group, such as the asset type (repository), a group (certain business unit), or even a product.&#x20;

Assets in the inventory are presented with key attributes in the following columns:

* **Asset** - The name of the repository asset, scanned artifact, and the Git remote URL, if available. Scanned artifacts are missing Git remote URLs.
* **Issue** - The number of issue counts on open assets aggregated across all relevant tools of the same severity of the asset itself and its child assets or packages. The severity level is classified into **C** (critical), **H** (high), **M** (medium), and **L** (low).
* **Controls** - A report detailing all products detected by the Snyk AppRisk on a specific repository asset and all products that should be but are not covered by the Snyk AppRisk.
* **Tags** -  Snyk AppRisk automatically tags repository assets with information about the used technologies (Python, Terraform, and so on) in the repository, and repository latest updates. You can also use policies to tag repository assets.
* **Developers** - Includes the SCM profile details for code committers to the repository asset.
* **Class** - Reflects the business criticality of the asset from A (most critical) to D (least critical), as defined by the user in the Policies view. You can manually change the class or automatically change it by applying a policy. You can lock the value you have manually set for a Class to prevent policies from overriding it.
* **Source** - Reflects the source of the asset, which can come from Snyk, an SCM, or a third-party integration.
* **SCM Repository freshness** - Reflects the status of the repository and the date of the last commit.
* **Clusters** - Provides a list of all the cluster names where the image asset is deployed.&#x20;

{% hint style="warning" %}
The Clusters column is hidden by default. To enable it, click Columns, select Clusters from the dropdown list, then click Apply to save the changes.&#x20;
{% endhint %}

### **Asset Sources, Types, and Scanned Artifacts**

Snyk AppRisk derives assets from Snyk automatically, and also from any SCM tools that are onboarded using the Snyk AppRisk Integration Hub. SCM tools from Snyk AppRisk Integration Hub may add additional repositories that aren’t scanned by Snyk, as well as additional context such as teams and code committers.

### Repository assets, scanned artifacts and packages

#### Repository assets

Snyk AppRisk supports repository assets (from main branches) as an asset type. Repository assets are visible in all inventory layouts and are supported by Policies. To avoid duplication, assets are identified using a unique identifier, which is the git remote URL for repository assets.

{% hint style="info" %}
For Snyk AppRisk SCM imported repositories, archived or deleted repositories will not be displayed in the asset inventory and will not be shown in the dashboard widgets.
{% endhint %}

#### Scanned artifacts

Snyk AppRisk also includes the concept of scanned artifacts. A scanned artifact is an entity detected by Snyk that cannot be identified as a repository asset because it does not include identifying information, such as a Git remote URL.&#x20;

Scanned artifacts provide users with visibility into what Snyk AppRisk detects from scans but require additional troubleshooting.&#x20;

You can see the scanned artifacts in the Inventory Type view. The scanned artifacts are not supported by Policies. Furthermore, scanned artifacts may include duplicates, as identifying information is missing.

#### Packages

Packages are defined as software or libraries that are managed by package management systems.

Package assets are created when you scan the dependencies of a Project through package management systems or by using the Snyk CLI. This enables Snyk AppRisk to identify and analyze the security vulnerabilities of the packages used within a Project, offering insights into possible risk exposures and providing recommendations for mitigation.

