# Assets inventory features

{% hint style="info" %}
**Release status** \
Risk factors for Snyk AppRisk are in Early Access and are available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.

**Runtime discovered** and **Runtime last seen** filters take their release status from the&#x20;
{% endhint %}

Snyk Essentials provides powerful search and filtering capabilities to help you narrow down assets for investigation and mitigation.

## Search capabilities

Use the search bar to search for specific keywords across various asset properties. Results can also include data retrieved from the Attributes tab of an asset, not only by the asset name.&#x20;

## Filters capabilities

### Quick filters

Quick filters are pre-defined filters to help you quickly focus on important assets. Here are the available quick filters:

* Assets with risk factors `Deployed` and `Public facing` and `No Coverage`: displays only the assets that have both the `Deployed` and `Public facing` risk factors available, and with a coverage gap on any of the selected Snyk products (available only with Snyk AppRisk).
* Assets with risk factor `Deployed` and `No Coverage`: displays only the assets that have both the `Deployed` risk factor available, and with a coverage gap on any of the selected Snyk products  (available only with Snyk AppRisk).
* Assets with Repository freshness Active and No Coverage: displays only the assets from active repositories and with a coverage gap on any of the selected Snyk products.
* Assets with `Asset Class A` and `No Coverage:`displays only the assets that are Class A and with a coverage gap on any of the selected Snyk products.

{% hint style="info" %}
After applying a quick filter, you can change or add additional filters by clicking on **Advanced Filters**.&#x20;
{% endhint %}

If you had a quick filter selected and then added an advanced filter, the quick filter will be deselected. In the **Advanced Filters** section, you can see the filters used to create the **Quick filter** you selected and can add additional filters on top of those. &#x20;

### Advanced filters

With this feature, you can [define filters](../manage-risk/policies/assets-policies/create-policies.md#define-filters) and filter assets based on highly specific criteria. For example, repository assets that have `AWS` in the name, are classified as either **A** or **B**, and do not have Snyk IaC as a control executed. This can be useful for finding repositories that have infrastructure as code that Snyk IaC does not scan.

Click Filters from the top left side of the screen. A pop-up is displayed allowing you to add quick filters or advanced filters.  If you select the Advanced filters, you can specify one or more sets of criteria as follows:

* **Property** of an asset (such as asset name, class, control executed).
* **Condition** depends on the asset selected (such as `contains` or `does not contain` for `asset name`).
* **Value** depends on the Property and Condition.

You can add as many filters as needed. To add another filter click **Add Filter,** set the condition as **And** or **Or** and customize the **Property**, **Condition**, and **Value** fields.&#x20;

{% hint style="info" %}
If you use Snyk Essentials for the first time, start with the **Coverage** filter to determine where you have Snyk implemented.
{% endhint %}

You can filter the information for all the inventory layouts and can use the following filters available under the **Advanced filters** section:

* **Application\*** - the list of the applications for which you have configured the application context catalog in Snyk Essentials.
* **Asset ID** - the unique identifier of the asset.
* **Asset name** - the name of the asset.
* **Asset type** - repository, package, or scanned artifact.
* **Attribute** - asset attributes retrieved from the data source.
* **Catalog name\*** - the name of your application context catalog.
* **Category**\* - the category of a repository asset. For example, service or library.
* **Class** - specify the class of the asset.
* **Clusters** - specify the cluster names where the asset is deployed (an asset can be deployed in more than one cluster).
* **Coverage** - specify the product or products used to scan the asset. The Coverage filter identifies if at least one scan has been run by the specified product.
* **Coverage gap** - specify the products for which the asset has not been scanned and do not meet the Set Coverage Control Policy requirements. The coverage gap applies only if you previously defined the coverage requirements of an asset and the asset has never been scanned, or the last scan is older than the default scanning frequency.
* **Developers** - specify the developer or developers who contributed to the asset.
* **Discovered** - specify the period when the asset was discovered.
* **Issue severity** - specify the severity of the issue: critical, high, medium, low.
* **Issue source** - specify where the issue was identified: SCM or third-party integrations. A source category will only be visible if there is at least one source present.
* **Last seen** - specifies the most recent time this asset was detected by Snyk in any of the sources (for example, Snyk, SCM, and so on).
* **Lifecycle\*** - represents the lifecycle state of the application context catalog component, for example `production`, `experimental`, `deprecated`.
* **Locked attributes** - specify if the attribute value is locked.
* **Organizations** - lists all the Snyk Organizations that are mapped to an asset.
* **Owner\*** - represents the team owning the repository for which the application context catalog was configured.
* **Risk factors** - The list of available risk factors. Risk factors refer to assets that can be vulnerable to security threats based on their exposure, sensitivity, compliance with security standards, and vulnerability history.
* **Runtime discovered** - specify the period when the runtime image asset was discovered.
* **Runtime last seen** - specifies the most recent time this image asset was detected in runtime.
* **Visibility** - the visibility status of the repositories.
  * **Public**: Repositories that are publicly accessible.
  * **Private**: Restricted repositories.
  * **Internal**: Internal repositories specific to GitHub and GitLab.
* **SCM Project** -  Specify the Project in the Azure DevOps or Bitbucket SCM integrations, where this asset is located.
* **SCM Organization** - Specify the SCM Organization or Workspace where this asset is located.&#x20;
* **SCM Repository freshness** - the status of the repository and the date of the last commit.
  * **Active**: Had commits in the last 3 months.
  * **Inactive**: The last commits were made in the last 3 - 6 months.
  * **Dormant**: No commits in the last 6 months.
  * **N/A**: There are no commits detected by Snyk Essentials.

{% hint style="info" %}
N/A indicates that the repository was detected through a Snyk scan but not directly from the SCM. To detect SCM repositories, you must set up SCM integration at the Group level.
{% endhint %}

* **Source** - specify the asset source.
* **Tags** - information about the detected languages and repository update status.
* **Title\*** - represents the name of the component for which the application context catalog was configured.

**\***&#x41;ll filters marked with `*` are visible only to the users who configured the [application context](../developer-tools/scm-integrations/application-context-for-scm-integrations/) catalog for their SCM integrations.

After applying the filters, the asset list will display only the assets that match the specified conditions, without a hierarchical structure.

{% hint style="info" %}
**Coverage** and **Coverage gap**  filter differences

* Use the Coverage filter to identify the assets scanned by the products at least once.
* Use the Coverage gap filter for assets that do not meet the requirements defined in the Set coverage control policy.&#x20;

The Coverage gap filter identifies assets that fall 'out of policy' and do not satisfy the coverage criteria you have specified, due to infrequent scanning or no scanning at all. On the other hand, the Coverage filter allows you to locate assets that have or have not been scanned, irrespective of any coverage requirements.
{% endhint %}



### Filtering using the unenriched repository banner

With this feature, you can filter the repositories that are not being enriched by the Group-level integration. These assets are discovered by the Organization-level integration or by the Snyk targets. To filter by the unenriched repositories, select the **\[the number of] repositories** part of the banner.

<figure><img src="../.gitbook/assets/image (512).png" alt=""><figcaption></figcaption></figure>

#### Unenriched assets with Group SCM Integration

As a user with Group-level integration, you can see the banner in Inventory about assets not being discovered through the group-level SCM integration. You already have a Group-level integration, but assets are not being discovered. Below is a playbook to better understand where each asset is coming from and why it is not discovered through the Group-level SCM integration.

\
Potential reasons for this include the following:

* **Permission issues** - Permission issues can happen when the Organization-level integration has wider permissions than the Group-level integration.
* **Old deleted repository** - a repository that Snyk tested (has a target in Snyk) and was deleted in the SCM integration before the group-level integration was set up.
* **Snyk CLI Projects**  - Snyk CLI projects that have a different repo URL than the SCM repository URL.
* **Missing SCM Organization** - SCM Organization is not being explicitly written and is missing in the Group-level integration.
* **Discovered via a different vendor** - The asset was discovered by a vendor that is not SCM integration, such as GitGuardian.\


There are three sections with step-by-step instructions:

1. **Source is only Snyk:** This section covers assets that are not discovered by Group-level integration, or the Organization-level integration.&#x20;
2. **Source is something other than Snyk (SCM and Snyk or just SCM):** This section covers assets that are not discovered by Group-level integration and are discovered by Organization-level integration.
3. **Additional tools:** This section provides a Looker dashboard to do a deep dive into all the assets and where they are discovered.

#### Source is only Snyk

This section covers assets that are not discovered by Group-level integration or the Organization-level integration, meaning they are only discovered ‌through the Snyk targets. Issues here are Snyk CLI Projects or old deleted repositories.

1. Log in to the Snyk Web UI and select your [Group](https://docs.snyk.io/snyk-platform-administration/groups-and-organizations).
2. Navigate to **Inventory**.
3. Select the repositories that are not being discovered through Group-level SCM Integration.

<figure><img src="../.gitbook/assets/image (513).png" alt=""><figcaption></figcaption></figure>

2.  Filter out the SCM sources first to tackle all assets not discovered through Group or Organization-level integrations. **Source is not one of**.

    <figure><img src="../.gitbook/assets/image (514).png" alt=""><figcaption></figcaption></figure>
3.  If the asset comes from a Snyk CLI project, check the repo URL . When a repo URL from the Snyk CLI Project does not match the repo URL from the SCM repository, duplicated assets can appear. Ensure the local repo URL matches the remote repo URL.

    <figure><img src="../.gitbook/assets/image (516).png" alt=""><figcaption></figcaption></figure>
4. If the project is not a Snyk CLI project, it’s possible that these repositories were deleted long before setting up the group-level integration. Double-check to ensure the repositories exist in their SCMs. If they are deleted, remove those assets by deleting the Snyk targets of those repositories.&#x20;
5.  If it’s not a Snyk CLI project or a deleted Project, it’s possible the repository was imported, but the integration connection is no longer working. This means it cannot be picked up through Auto Repo Discovery, and there are not enough permissions in the Group-level integration to discover it. In this case, double-check the Organization-level integrations and the Group-level integration’s permissions.



#### Source field is something other than Snyk (SCM and Snyk or just SCM)

This section covers assets that the Group-level integration does not discover and are discovered by the Organization-level integration. Issues here are either missing SCM Organization in the integration, permission issues, and/or assets being discovered through a different vendor.

1. Log in to the Snyk Web UI and select your [Group](https://docs.snyk.io/snyk-platform-administration/groups-and-organizations).
2. Navigate to **Inventory**.
3. Select the repositories that are not being discovered through Group-level SCM Integration

<figure><img src="../.gitbook/assets/image (517).png" alt=""><figcaption></figcaption></figure>

1.  **Filter** **Source** **is one of the SCMs**, which will exclude assets that have a source that is only Snyk, and select Apply.\


    <figure><img src="../.gitbook/assets/image (518).png" alt=""><figcaption></figcaption></figure>
2. Open one of the assets and look for the SCM Organization in the URL . The following screenshots show you how it will look like based on your SCM integration
   1.  GitHub

       <figure><img src="../.gitbook/assets/image (519).png" alt=""><figcaption></figcaption></figure>


   2.  GitLab

       <figure><img src="../.gitbook/assets/image (520).png" alt=""><figcaption></figcaption></figure>


   3.  Azure DevOps

       <figure><img src="../.gitbook/assets/image (521).png" alt=""><figcaption></figcaption></figure>


   4.  Bitbucket

       <figure><img src="../.gitbook/assets/image (522).png" alt=""><figcaption></figcaption></figure>
3.  On GitHub and Azure, you can see if ‌Organizations were added to the profiles. Go to the Group-level integration page and open the SCM profile to check if the SCM Organization has been added to it.&#x20;

    <figure><img src="../.gitbook/assets/image (523).png" alt=""><figcaption></figcaption></figure>
4. If it has not been added, add the SCM Organization and select **Done**.
5. If it has been added, check the token's permissions and ensure it has access to the unenriched repositories.
6. For GitLab and Bitbucket, ensure that the Group-level tokens have access to and the right permissions for the organization from the repo URL.
7. Another reason this can happen is if the assets are being discovered by a different vendor. Ensure the repo URLs are the same to avoid creating duplicate assets.

