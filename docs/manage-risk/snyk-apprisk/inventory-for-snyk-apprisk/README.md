# Inventory for Snyk AppRisk

Snyk AppRisk inventory layouts are organizing your repository assets in meaningful ways, enabling you to:

* Gain full repository asset visibility from your SCM tools, including details about configured teams and repository code committers.
* Track controls coverage for Snyk products.
* Prioritize coverage mitigation efforts according to business impact.

Each line in the inventory represents either a repository asset or a scanned artifact from Snyk that is likely a repository but lacks some identifying information. Scanned artifacts are not supported through Policies.

## Inventory Layouts <a href="#inventory-layouts" id="inventory-layouts"></a>

To get better context and clarity over your asset inventory, Snyk AppRisk allows flexible structuring with inventory layouts. Snyk AppRisk includes four inventory layouts and groups assets by different contexts. You can find all inventory layouts under the Inventory menu option at the Group level:&#x20;

* **Code Assets**: A hierarchical list of all assets.&#x20;
* **Organization Teams**: SCM repository assets grouped by teams. Note that only SCM organizations with teams, and repositories assigned to a team, appear on this layout.
* **Technology**: SCM repository assets grouped by technology, as detected and tagged by Snyk AppRisk.
* **Type**: All the discovered assets, grouped by their type.&#x20;

Each inventory layout may include different counts of assets and scanned artifacts, depending on the grouping context. Otherwise, all columns and data manipulation features are the same on each inventory layout.

You can filter the information for all the inventory layouts and can use the following filters:

* **Asset ID** - the unique identifier of the asset.
* **Asset name** - the name of the asset.
* **Asset type** - repository, package or scanned artifact.
* **Attribute** - asset attributes retrieved from the data source.
* **Class** - specify the class of the asset.
* **Coverage** - specify the product or products used to scan the asset. The Coverage filter identifies if at least one scan has been run by the specified product.
* **Coverage gap** - specify the products for which the asset has not been scanned and do not meet the Set Coverage Control Policy requirements. The coverage gap applies only if you previously defined the coverage requirements of an asset and the asset has never been scanned, or the last scan is older than the default scanning frequency.
* **Developers** - specify the developer or developers who contributed to the asset.
* **Discovered** - specify the period when the asset was discovered.
* **Issue severity** - specify the severity of the issue: critical, high, medium, low.
* **Issue source** - specify where the issue was identified: SCM or third-party integrations. A source category will only be visible if there is at least one source present.
* **Locked attributes** - specify if the attribute value is locked.
* **Last seen** - specify the repository freshness status.
* **SCM Repository freshness** - the status of the repository and the date of the last commit.
  * **Active**: Had commits in the last 3 months.
  * **Inactive**: The last commits were made in the last 3 - 6 months.
  * **Dormant**: No commits in the last 6 months.
* **Source** - specify the asset source.
* **Tags** - information about the detected languages and repository update status.

<figure><img src="../../../.gitbook/assets/Inventory view-new UI.png" alt="Snyk AppRisk - Inventory layouts"><figcaption><p>Snyk AppRisk - Inventory layouts</p></figcaption></figure>

## Assets and their attributes

Every item listed in the inventory is considered an individual asset. Most assets are actual components of the application (code repositories, domains, endpoints, and so on), but an asset can also be a representation of a Group such as the asset type (repository), a group (certain business unit), or even a product.&#x20;

{% hint style="warning" %}
The Issues column is available in a Closed Beta state and is applicable only for the Snyk AppRisk Pro version.  Please contact your salesperson if you are interested in Snyk AppRisk Pro.
{% endhint %}

Assets in the inventory are presented with key attributes in the following columns:

* **Asset** - The name of the repository asset, scanned artifact, and the Git remote URL, if available. Scanned artifacts are missing Git remote URLs.
* **Issue** - The number of issue counts on open assets aggregated across all relevant tools of the same severity of the asset itself and its child assets or packages. The severity level is classified into **C** (critical), **H** (high), **M** (medium), and **L** (low).
* **Controls** - A report detailing all products detected by the Snyk AppRisk on a specific repository asset and all products that should be covered but are not yet by the Snyk AppRisk.
* **Tags** -  Snyk AppRisk automatically tags repository assets with information about the used technologies (Python, Terraform, and so on) in the repository, and repository latest updates. You can also use policies to tag repository assets.
* **Developers** - Includes the SCM profile details for code committers to the repository asset.
* **Class** - Reflects the business criticality of the asset from A (most critical) to D (least critical), as defined by the user in the Policies view. The class can be manually changed by you or automatically changed by a policy. You can lock the value you have manually set for a Class to prevent policies from overriding it.
* **Source** - Reflects the source of the asset, which can come from Snyk, an SCM, or a third-party integration.
* **SCM Repository freshness** - Reflects the status of the repository and the date of the last commit.

### **Asset Sources, Types, and Scanned Artifacts**

Snyk AppRisk derives assets from Snyk automatically, and also from any SCM tools that are onboarded using the Snyk AppRisk Integration Hub. SCM tools from Snyk AppRisk Integration Hub may add additional repositories that aren’t scanned by Snyk, as well as additional context such as teams and code committers.

### Repository assets and scanned artifacts

#### Repository assets

Currently, Snyk AppRisk supports repository assets as an asset type. In future releases, Snyk will add support for additional code-based assets, such as software packages. Repository assets are visible in all inventory layouts and are supported by Policies. To avoid duplication, assets are identified using a unique identifier, which is the git remote URL for repository assets.

{% hint style="info" %}
If you archive or delete repositories, they are not displayed in the asset inventory and, therefore, not considered to be shown in the dashboard widgets.
{% endhint %}

#### Scanned artifacts

Snyk AppRisk also includes the concept of scanned artifacts. A scanned artifact is an entity detected by Snyk and cannot be identified as a repository asset because it does not include identifying information, such as a Git remote URL.&#x20;

Scanned artifacts provide users with visibility into what Snyk AppRisk detects from scans but require additional troubleshooting.&#x20;

You can see the scanned artifacts in the Inventory Type view. The scanned artifacts are not supported by Policies. Furthermore, scanned artifacts may include duplicates, as identifying information is missing.

