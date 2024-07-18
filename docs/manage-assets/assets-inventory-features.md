# Assets inventory features

Snyk AppRisk provides powerful search and filtering capabilities to help you narrow in on assets for investigation and mitigation.

## Search capabilities

Use the search bar to search for specific keywords across various asset properties. Results can also include data retrieved from the Attributes tab of an asset, not only by the asset name.&#x20;

## Filters capabilities

With this feature, you can [define filters](../manage-risk/policies/assets-policies/create-policies.md#define-filters) and filter assets based on highly specific criteria. For example, repository assets that have `AWS` in the name, are classified as either **A** or **B**, and do not have Snyk IaC as a control executed. This can be useful for finding repositories that have infrastructure as code that Snyk IaC is not currently scanning.

Click Filters from the top left side of the screen. A pop-up is displayed allowing you to add new filters. The filter capability enables you to specify one or more sets of criteria as follows:

* **Property** of an asset (such as asset name, class, control executed).
* **Condition** depends on the asset selected (such as `contains` or `does not contain` for `asset name`).
* **Value** depends on the Property and Condition.

<figure><img src="../.gitbook/assets/image (502).png" alt="Snyk AppRisk - Filters"><figcaption><p>Snyk AppRisk - Filters</p></figcaption></figure>

You can add as many filters as needed. To add another filter click **Add Filter,** set the condition as **And** or **Or** and customize the **Property**, **Condition**, and **Value** fields.&#x20;

There are several filtering options:

* **Asset ID -** filter assets by their unique identifier.
* **Asset Type** - filter assets by type.
* **Asset Name** - filter assets by name.
* **Tags** - filter assets by certain categories, for example, tags.\
  Assets have default tags for various use cases, such as tracing active/inactive assets, filtering by specific technology, and more. On top of it, users can create policy rules to tag assets according to their own logic.
* **Discovered** -  filter assets according to duration, since they were discovered.
* **Locked attributes** -  filter assets with locked attributes to identify the assets that are not affected by your policies.
* **Coverage and Coverage gap -** mostly used to answer scan coverage questions.&#x20;
  * **Coverage** means that an asset has been tested by this product at some point in the past.
  * **Coverage gap** means the asset does not meet the coverage requirements as set by the **Set coverage control** policy.

{% hint style="info" %}
If you use Snyk AppRisk for the first time, start with the **Coverage** filter to determine where you currently have Snyk Implemented.
{% endhint %}

## Available filters

{% hint style="warning" %}
**Release status** \
**Risk factors** on assets, **Runtime discovered** and **Runtime last seen** filters are currently in Closed Beta and available only with Snyk AppRisk Pro.&#x20;
{% endhint %}

Each filter component requires you to specify an asset property. Available properties for asset policies include:

* **Application\*** - the list of the applications for which you have configured the Backstage catalog in Snyk AppRisk.
* **Asset ID** - the unique identifier of the asset.
* **Asset name** - the name of the asset.
* **Asset type** - repository, package, or scanned artifact.
* **Attribute** - asset attributes retrieved from the data source.
* **Catalog name\*** - the name of your backstage catalog.
* **Category**\* - the category of a repository asset. For example, service or library.
* **Class** - specify the class of the asset.
* **Coverage** - specify the product or products used to scan the asset. The Coverage filter identifies if at least one scan has been run by the specified product.
* **Coverage gap** - specify the products for which the asset has not been scanned and do not meet the Set Coverage Control Policy requirements. The coverage gap applies only if you previously defined the coverage requirements of an asset and the asset has never been scanned, or the last scan is older than the default scanning frequency.
* **Developers** - specify the developer or developers who contributed to the asset.
* **Discovered** - specify the period when the asset was discovered.
* **Issue severity** - specify the severity of the issue: critical, high, medium, low.
* **Issue source** - specify where the issue was identified: SCM or third-party integrations. A source category will only be visible if there is at least one source present.
* **Last seen** - specify the repository freshness status.
* **Lifecycle\*** - represents the lifecycle state of the backstage catalog component, for example `production`, `experimental`, `deprecated`.
* **Locked attributes** - specify if the attribute value is locked.
* **Owner\*** - represents the team owning the repository for which the backstage catalog was configured.
* **Risk factors** - The list of available risk factors. Risk factors refer to assets that can be vulnerable to security threats based on their exposure, sensitivity, compliance with security standards, and vulnerability history.
* **Runtime discovered** - specify the period when the runtime image asset was discovered.
* **Runtime last seen** - specify the freshness status for the runtime image asset.
* **SCM Repository freshness** - the status of the repository and the date of the last commit.
  * **Active**: Had commits in the last 3 months.
  * **Inactive**: The last commits were made in the last 3 - 6 months.
  * **Dormant**: No commits in the last 6 months.
* **Source** - specify the asset source.
* **Tags** - information about the detected languages and repository update status.
* **Title\*** - represents the name of the component for which the backstage catalog was configured.

**\***All filters marked with `*` are visible only to the users who configured the Backstage catalog for their SCM integrations.

After applying the filters, the asset list will display only the assets that match the specified conditions, without a hierarchical structure.

