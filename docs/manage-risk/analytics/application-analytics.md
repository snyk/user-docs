# Application analytics

{% hint style="info" %}
**Release status**&#x20;

Application Analytics for Snyk AppRisk is in Early Access and available only with Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

The Analytics menu is available at the tenant level, under the Application Analytics tab (available only with Snyk AppRisk). Application Analytics is designed to provide an in-depth analysis of your AppSec program, highlighting areas for improvement, emerging risks, and previously overlooked vulnerabilities to support AppSec managers and engineering teams.&#x20;

The dashboard displays essential data such as the status and trends of open issues, control coverage, and repository metadata. It also shows the state of imported assets. It provides a comprehensive and at-a-glance review of information from different viewpoints, such as asset class, application, or team, with a global filter bar to enhance your experience.

## Overview

Application Analytics enables you to review and explore your AppSec program status and results from a top-down approach. You can start the exploration from a high, general level over applications, teams (owners), or asset classes, and then narrow it down to the asset level.\
You can enhance the security of your application by identifying areas for improvement, recognizing developing risks, and addressing blind spots. The Application Analytics retrieves the data from all the Groups available for the tenant.

{% hint style="info" %}
If you are using Snyk Essentials, navigate to the asset dashboard page to learn more about your assets or remain on the [Analytics](issues-analytics.md) page to explore the detected issues.
{% endhint %}

Harnessing Application Analytics provides answers to questions such as:

* Which sensitive assets are being publicly exposed and not tested according to the coverage policy?
* Which applications and code owners bear the most risk in terms of accumulated critical and high issues, and how do they compare to others?
* How many repositories exist without a clear association to an application or a code owner? And are new assets being associated as expected?

<figure><img src="../../.gitbook/assets/image (220).png" alt=""><figcaption><p>Application Analytics Overview</p></figcaption></figure>

## Filters and views <a href="#filters-and-views" id="filters-and-views"></a>

You can customize your data by using the available filters, dimension views, and specific timeframes.&#x20;

Filters are applied at the tenant level, and after being customized, they have an impact on all the reports and statistics presented on the Application Analytics page.

You can refine the data even more by using the View by options. This focuses the data on specific dimensions: Asset Class, Application, or Owner.

### Filters

The filters are located at the top left of your Application Analytics page, and you can customize them based on your needs.

The following are the available filters:

| Filter                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Groups                                                                                                                                                               | <p>Provides a list of all the available Groups that exist for the selected tenant.<br>You can customize the selection and focus only on specific Groups. The default setting is providing information for all available groups.</p>                                                                                                                                                                                                                                                                                                                    |
| Issue severity                                                                                                                                                       | Provides a list with all available types of severity for an issue. The default setting provides information about issues with Critical and High severity.                                                                                                                                                                                                                                                                                                                                                                                              |
| <p>Add filter</p><ul><li>Asset type</li><li>Asset classes</li><li>Assets application</li><li>Assets owner</li><li>Asset risk factors</li><li>Issues source</li></ul> | <p>You can add filters for a more customized data analysis. </p><ul><li>Asset type - filter by the asset type (Container image, Repository)</li><li>Asset classes - filter by the asset class (A, B, C, D)</li><li>Assets application - filter by the application for which you want to see the assets</li><li>Assets owner - filter by the repository owner of the analyzed assets</li><li>Asset risk factors - filter by specific risk factors of the analyzed assets </li><li>Issues source - filter by the source of the analyzed issues</li></ul> |
| Reset filters                                                                                                                                                        | Resets the filters to the default state                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Views

Managing an AppSec program can be challenging, especially ensuring complete visibility of assets and issues. Identifying assets that require protection and monitoring them with all applicable ASTs can be difficult. This task becomes more complex with new assets and misconfigurations in ASTs, leading to incomplete coverage and critical visibility.\
\
AppSec teams must maintain a comprehensive understanding of risks and vulnerabilities linked to applications and their respective owners.

Viewing metrics for an application or owner is much more meaningful and helps to:

* Clearly communicate the status and developing trends to all levels and groups across the enterprise.
* Identify who should be aware of the situation and who should take action.
* Make comparisons and conclude where more attention is needed.

Application analytics provides you with new levels of visibility over your important assets, applications, and code owners (teams) and helps you to identify and take action on risk and coverage.

Collaboration across Research and Development teams is necessary for achieving optimal visibility and requires attention from the AppSec team.

You can display the analytics view by:

* Asset Class
* Application&#x20;
* Owner

By selecting a View by dimension, all exposed widgets will be affected, enabling you to compare data points based on the selected dimension.

{% hint style="info" %}
The widgets display the top five applications or code owners based on the context. For instance, in the "open issues by control" widget, the top five applications or owners are chosen based on the total number of issues after applying the display views.\
You can also compare specific applications or owners by adding the application or owner display views.
{% endhint %}

Assets and applications vary in importance and sensitivity. Some repositories are internal and used for testing only, while others are public-facing and used in key services.

The dashboard default view compares assets and issues metrics by asset class. Display the view by dimension to see a comparison between applications or code owners throughout the dashboard.

#### Asset Class view <a href="#asset-class-view" id="asset-class-view"></a>

[Asset class](../../manage-assets/assets-inventory-components.md#class) reflects the business criticality of the asset from A (most critical) to D (least critical).\
By having this level of visibility, you can prioritize the most crucial assets in your inventory, applications, or code owners.\
To associate assets with asset class, you can either change the asset class manually in the inventory screen or, preferably, define a [classification policy](../policies/assets-policies/use-cases-for-policies/classification-policy.md) that will automatically assign an asset class to your assets.

#### Applications and Owner view <a href="#applications-and-owners-view" id="applications-and-owners-view"></a>

You can filter the data from your Application Analytics dashboard based on application or code owner. To proceed, it is necessary to have the appropriate metadata available for the repositories. The metadata can be pulled directly from the Snyk SCM integration. You can find details of how to set this up on the [Backstage catalog in Asset Inventory](../../developer-tools/scm-integrations/application-context-for-scm-integrations/#backstage-file-for-scm-integrations) page. \
To determine if this metadata is available in your repositories, check the completeness widget for repository metadata. Snyk recommends verifying that all class A assets are properly configured by using the asset class filter from the dashboard.&#x20;

### Analytics timeframe

You can select a specific date range for the assets analyzed data by adding the **Asset Introduction Date** filter. Applying that filter will impact all non-trend widgets, narrowing them down from showing all available data to data for assets introduced in the selected date range. The trend widgets are configured to show a fixed timeframe of the last three (3) months.

{% hint style="info" %}
The data in Application Analytics is updated on an daily basis.
{% endhint %}

The following video presents an overview of the Application Analytics filters and views from the Snyk Web UI.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1739216878/snyk-learn/product-training-videos/Snyk_AppRisk_-10a_-_v3_-_App_Analytics_Navigation_and_filters.mp4" %}
Application Analytics - filters and views
{% endembed %}

## Data categories <a href="#data-categories" id="data-categories"></a>

The Application Analytics dashboard focuses on three main data categories:

* Coverage - provides the coverage status and the trends for the analyzed assets.
* Issues - provides the status of the open issues.
* Assets - provides the coverage status of the repository metadata and the status and trends for the imported assets.

### Coverage

One of the leading missions of an AppSec team is ensuring appropriate scan coverage across the asset inventory. A [covered asset](../../discover-snyk/getting-started/glossary.md#coverage-snyk-essentials) is simply an asset that has been scanned by a certain application security testing (AST) product. Having uncovered assets expose the company to unknown risks, that is why it is essential to verify that business critical assets (based on asset class or strategic applications), are being properly scanned.

In the Coverage section, you have information about the assets coverage.&#x20;

* Coverage overview - provides information, in percentages, about the scanned assets, distributed by the scan category (SAST, SCA, Container, and Secrets).
* Coverage trend - allows reviewing the coverage trend for the last three (3) months. A growing trend will indicate that a larger portion of your asset inventory was scanned.

The Coverage section is based on the scan category and is not impacted by the selected view ( Asset Class, Application, or Owner).

<figure><img src="../../.gitbook/assets/image (223).png" alt=""><figcaption><p>The Coverage Section</p></figcaption></figure>

The following video presents an overview of the Application Analytics Coverage view from the Snyk Web UI.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1739216985/snyk-learn/product-training-videos/Snyk_AppRisk_-10b_-_v1_-_App_Analytics_Coverage.mp4" %}
Application Analytics Coverage view
{% endembed %}

### Issues

In the Issues section, you have information about the analyzed open issues.

* Open issues by category - This graphic provides a clear overview of the number of issues distributed by the issue source category (SAST, SCA, Container, and Secrets) and by the selected view (allowing to compare between asset classes, applications and owners).
* Open issues breakdown - This graphic provides information about the backlog of your open issues. The desired trend is a negative one, especially for higher asset classes or strategic applications. The selected view allows comparing asset classes, applications and owners.

You can choose to view the issues based on Asset Class, Application, or Owner. The focus of the presented information is changed based on your View by selection. When viewing by application or owner, only the top five (5) applications or owners with the most issues are displayed.

You can see more details about each graphic by hovering over the presented data. Extra controls are available on the right side of each graphic, allowing you to download it as an image.

The following video presents an overview of the Application Analytics Issues view from the Snyk Web UI.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1739217894/snyk-learn/product-training-videos/Snyk_AppRisk_-10c_-_v1_-_App_Analytics_Issues.mp4" %}
Application Analytics Issues view
{% endembed %}

### Assets

In the Assets section, you have information about the analyzed assets.

* Risk factors breakdown **-** a funnel that shows the progression of risk factors on code repositories and container images. Each layer is divided according to the selected view, asset class, application, or owner. When viewing by application or owner, only the top five (5) applications or owners will be displayed according to the number of assets with risk factors.
* New assets introduced - allows tracking the inventory size over time. The trend only counts repositories and container images. When viewing by application or owner, only the top five (5) applications or owners with the most assets will be displayed.

You can choose to view the Assets section based on Asset Class, Application, or Owner. The focus of the presented information is changed based on your View by selection.

You can see more details about each graphic by hovering over the presented data. Extra controls are available on the right side of each graphic, allowing you to download it as an image.

The following video presents an overview of the Application Analytics Assets view from the Snyk Web UI.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1739218078/snyk-learn/product-training-videos/Snyk_AppRisk_-10d_-_v1_-_App_Analytics_Assets.mp4" %}
Application Analytics Assets view
{% endembed %}

### Metadata completeness

The metadata completeness section provides information on the completeness of metadata from application context sources for your repositories.

* Repo metadata completeness - displays the availability of application context metadata across code repositories. For more information about context metadata, see [Application context for SCM Integrations](../../developer-tools/scm-integrations/application-context-for-scm-integrations/).
* Repository source distribution - provides information about the repositories distributed by the type of integration (SCM integrations, third-party integrations). When viewing by application or owner, only the top five (5)  applications or owners with the most assets will be displayed.

The following video presents an overview of the Application Analytics Repository metadata completeness and source distribution view from the Snyk Web UI.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1739219742/snyk-learn/product-training-videos/Snyk_AppRisk_-10e_-_v1_-_App_Analytics_Repo_Metadata_Completeness_and_Source_Distribution.mp4" %}
Analytics Repository metadata completeness and source distribution view
{% endembed %}
