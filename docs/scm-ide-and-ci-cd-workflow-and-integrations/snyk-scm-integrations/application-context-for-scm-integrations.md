# Application context for SCM Integrations

These are the available integrations that you can set up for the application context:

* [Backstage file](application-context-for-scm-integrations.md#backstage-file-for-scm-integrations)&#x20;
* [ServiceNow CMDB](application-context-for-scm-integrations.md#servicenow-cmdb-for-scm-integrations)
* [Atlassian Compass](application-context-for-scm-integrations.md#atlassian-compass)
* [Harness](application-context-for-scm-integrations.md#harness)
* [OpsLevel](application-context-for-scm-integrations.md#opslevel)
* [Datadog Org Context ](application-context-for-scm-integrations.md#datadog-org-context-service-catalog)&#x20;

## Backstage file for SCM integrations

Backstage is a service catalog that allows users to add metadata or annotations to their repositories, helping to organize and categorize the available resources for easier navigation and understanding. You can leverage your SCM integration to pull metadata associated with Backstage catalog files into Snyk AppRisk.

You can use the Backstage catalog file for GitHub, GitLab, Azure DevOps, BitBucket Cloud, and BitBucket on-prem SCM integrations.

### Prerequisites

* Configured SCM integration for Snyk AppRisk.
* The `catalog-info.yaml` file is available in the root folder or your Project.&#x20;

### Required Parameters

* A configured SCM integration.&#x20;
* The `catalog-info.yaml` file from your Project.

### Integration Hub setup

1. Open the **Integration Hub** menu.&#x20;
2. Select an SCM integration.&#x20;
3. Click the **Settings** option of the SCM integration.&#x20;
4. Enable the **Add Backstage Catalog** option.
5. Optional - if the Backstage catalog filename in your repository is not `catalog-info.yaml` you can change the default value in the Backstage catalog filename field.
6. Select at least one attribute you want to add to Snyk AppRisk.

{% hint style="info" %}
Snyk AppRisk parses the fields of the detected file using the default field names unless an alternate field name is specified.
{% endhint %}

7. Click the **Done** button.

<figure><img src="../../.gitbook/assets/image (1) (12) (1).png" alt="Integration Hub - Backstage setup"><figcaption><p>Integration Hub - Backstage setup</p></figcaption></figure>

After you finish configuring the Backstage catalog, Snyk AppRisk starts enriching your repository assets with the data found in the backstage catalog .yaml file.

### Backstage catalog in Asset Inventory

Use the Backstage catalog to enrich the repository assets and to define the component entity. For this type of situation, a component is defined as a software component, like a service, repository, a website, library, and so on.&#x20;

Components have several attributes and most of them are optional:

* `spec.type` (mandatory) - represents the classification of the repository.&#x20;
* `spec.owner` (mandatory) - represents the team owning the repository.
* `spec.lifecycle` - represents the lifecycle state of the component, for example `production`, `experimental`, `deprecated`.
* `spec.system` (optional) - represents a group of components that serve the same purpose. This concept is referred to as “Application”.
* `Metadata.name` (mandatory) - represents the name of the component.
* `Metadata.title` (optional) - represents the name of the component.

The Backstage data is dynamic and may change over time:

* If new commits or updates are made on the `catalog-info.yaml` file, then Snyk AppRisk updates the asset attribute for that specific repository asset.
* If the`catalog-info.yaml` file is removed from the repository, then Snyk AppRisk deletes the asset attribute from that specific repository assets.

#### Inventory menu

Depending on the selection you made on the Integration Hub configuration menu, only those selections are displayed in filters from the Inventory menu. For example, if you selected the Category attribute, then it will also be displayed in the filters list.

<figure><img src="../../.gitbook/assets/image (448).png" alt="Backstage selections in the Integration Hub menu"><figcaption><p>Backstage selections in the Integration Hub menu</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (367).png" alt="Backstage selections in the Snyk AppRisk Inventory menu" width="375"><figcaption><p>Backstage selections in the Snyk AppRisk Inventory menu</p></figcaption></figure>

#### Asset Summary Tab

The Asset Summary tab shows the six Backstage attributes that are configured in the Integration Hub only if you choose to integrate with Backstage.

<figure><img src="../../.gitbook/assets/image (368).png" alt="Backstage selections in the Asset Summary menu"><figcaption><p>Backstage selections in the Asset Summary menu</p></figcaption></figure>

#### Asset Attributes Tab

In the Asset Attributes tab only the selected attributes should be added as metadata to the repository asset.

```
{
    name:"spring.goof",
    repositoryURL:"https://github.com/snyk/spring.goof.git",
    context:[
             {
              name: "super-duper-component",
              title: "Super Duper Component",
	      application: "super-duper-app",
	      lifecycle: "production",
	      owner: "super-duper-team",
	      category: "service",
              source: "Backstage"
              }]
}
```

#### Policies Filter

In the policy builder you can find only the attributes you have previously selected when configuring the backstage catalog file.&#x20;

<figure><img src="../../.gitbook/assets/image (369).png" alt="Backstage selections in the Policies menu" width="375"><figcaption><p>Backstage selections in the Policies menu</p></figcaption></figure>

The following list describes all possible backstage attributes that you can choose from when you configure the backstage catalog file.&#x20;

* **Application** - represents a group of components that serve the same purpose.&#x20;
* **Owner** - specifies the team owning the repository.
* **Catalog name** - the metadata name.
* **Title** - a name to display for the entity instead of the property. It is an alternative to the metadata name, when the catalog name is too hard to read.
* **Category** - represents the classification of the repository. The Organization can choose any name or text.
* **Lifecycle** - specifies the lifecycle state of the component, for example production, experimental, deprecated.

The following video provides an overview of the Backstage file option from the Integration Hub and a quick explanation of the available attributes:

{% embed url="https://www.youtube.com/watch?v=gfOUSE0UhHA" %}
Liked the video? Checkout the rest of the course on [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)!
{% endembed %}

## ServiceNow CMDB for SCM integrations

{% hint style="warning" %}
**Release status**

The ServiceNow CMDB integration is in [Early Access](https://docs.snyk.io/getting-started/snyk-release-process#early-access) and available for both Snyk AppRisk Essentials and Snyk AppRisk Pro plans.
{% endhint %}

### Required Parameters <a href="#servicenow-cmdb-required-parameters" id="servicenow-cmdb-required-parameters"></a>

1. Add the **profile name** for your ServiceNow CMDB instance.
2. Setup the **CMDB instance** for the ServiceNow CMDB by following this example `https://<INSTANCE_NAME>.service-now.com`.
3. **Username** and **Password** - Credentials for your ServiceNow CMDB instance.
4. Add the t**able name** for the CMDB configuration item class. The list can be found here: [https://docs.servicenow.com/bundle/washingtondc-servicenow-platform/page/product/configuration-management/reference/cmdb-tables-details.html](https://docs.servicenow.com/bundle/washingtondc-servicenow-platform/page/product/configuration-management/reference/cmdb-tables-details.html)
5. Add the **CMDB field to map Repo URL** - Add the URL of the repository.

{% hint style="info" %}
* This feature is only for the integration with ServiceNow CMDB
* The data gathered by Snyk from ServiceNow CMDB will be correlated with the Repository Assets.
{% endhint %}

### Integration Hub setup <a href="#servicenow-cmdb-integration-hub-setup" id="servicenow-cmdb-integration-hub-setup"></a>

* Open the **Integration Hub** menu.&#x20;
* Select the **App Context** tag and search for ServiceNow CMDB.&#x20;
* Click the **Add** button.
* Add the **Profile name** - this is the name of your ServiceNow CMDB profile.
* Add the **CMDB Instance** - this is your ServiceNow instance, use this format: `https://<INSTANCE_NAME>.service-now.com`
* Add the **Username** and the **Password**- the username and password to access the ServiceNow CMDB instance
* Add the **Table name** - select the configuration item class that Snyk AppRisk should onboard. Use this format `cmdb_ci_<class>`
* Add the **CMDB Field to map Repo URL** - the specific URL that is being referred to in the ServiceNow CMDB record.
* You can select one or more attributes related to repository assets and configure where Snyk AppRisk can take this attribute in ServiceNow CMDB. Example:&#x20;
  * Category: application\_type
  * Owner: business\_unit
* Click the **Done** button.
* When the connection is established, the status of the ServiceNow CMDB integration is changed to **Connected**.

{% hint style="info" %}
The ServiceNow CMDB integration is available for both Snyk AppRisk Essentials and Snyk AppRisk Pro plans.
{% endhint %}

## Atlassian Compass

### Required Parameters

1. Add your Atlassian Compass **Profile name**.
2. Add your Atlassian Compass **Instance URL**. You can use this format type: `https://<YOUR ORGANIZATION>.atlassian.net`.
3. Add your Atlassian Compass **Username**.
4. Add your Atlassian Compass instance **Token**. Refer to this document for creating the Atlassian API token: [https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/](https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/)&#x20;

{% hint style="info" %}
The gathered data from Atlassian Compass will be correlated with the Repository Assets.

This feature is availabale only for the integration with Atlassian Compass.
{% endhint %}

### Integration Hub setup

1. Open the **Integration Hub** menu.&#x20;
2. Select the **App Context** tag and search for Atlassian Compass.&#x20;
3. Click the **Add** button.
4. Add the **Profile name** - this is the name of your Atlassian Compass profile.
5. Add the **Instance URL** - this is the URL of the Atlassian Compass instance. Use this format type: `https://<YOUR ORGANIZATION>.atlassian.net`
6. Add the **Username** - this is the username to access the Atlassian Compass instance.
7. Add the **Token** - this is the API token to access the Atlassian Compass instance.&#x20;
8. You can select one or more attributes related to repository assets that Snyk AppRisk can pull from Atlassian Compass based on the [Component Data](https://developer.atlassian.com/cloud/compass/forge-graphql-toolkit/Interfaces/Component/):&#x20;
   * **Catalog Name** - Matches with name.
   * **Category** - Identified when '`fields.definition.name`' equals tier.
   * **Lifecycle** - Identified when '`fields.definition.name`' equals lifecycle.
   * **Owner** - the `ownerId` (finding owner name from ownerId).
   * **Application** - the `typeId` (all component types, Application, Service, Library, and so on receive an ID).
9. Click the **Done** button.
10. When the connection is established, the status of the Atlassian Compass integration is changed to **Connected**, and Snyk AppRisk will start enriching repository assets with the data found in Atlassian Compass.

## Harness

### Required Parameters

1. Add your Harness **Profile name**.
2. Add the **Host URL** of your Harness account. You can use this format type: `https://<YOUR ORGANIZATION>.harness.io`
3. Add the **API key** for your Harness instance. You can use [this documentation](https://developer.harness.io/docs/platform/automation/api/add-and-manage-api-keys/) to manage your API key.

{% hint style="info" %}
This integration is focused on [Harness’s](https://developer.harness.io/docs/internal-developer-portal/catalog/software-catalog/) service catalog module and it is backed by the Backstage catalog.
{% endhint %}

### Integration Hub setup

1. Open the **Integration Hub** menu.&#x20;
2. Select the **App Context** tag and search for Harness.&#x20;
3. Click the **Add** button.
4. Add the **Profile name** - this is the name of your Harness instance.
5. Add the **Host URL** of your Harness account.&#x20;
6. Add the **API key** of your Harness instance.
7. Select at least one Harness software catalog [metadata](https://developer.harness.io/docs/internal-developer-portal/catalog/software-catalog#component-definition-yaml):
   * Catalog name - If you select this metadata, it is mandatory to add the **Catalog name key**.
   * Title - If you select this metadata, it is mandatory to add the **Title key**.
   * Category - If you select this metadata, it is mandatory to add the **Category key**.
   * Lifecycle - If you select this metadata, it is mandatory to add the **Lifecycle key**.
   * Owner - If you select this metadata, it is mandatory to add the **Owner key**.
   * Application - If you select this metadata, it is mandatory to add the **Application key**.
8. Click the **Done** button.
9. When the connection is established, the status of the Harness integration is changed to **Connected**, and Snyk AppRisk will start enriching repository assets with the data found in Harness.

## OpsLevel

### Required Parameters

1. Add your OpsLevel **Profile name**.
2. Add the **Instance URL** of your OpsLevel account. You can use this format type: `https://<YOUR Organizer>.opslevel.com`
3. Add the **API Token** for your OpsLevel instance. Create the API Token in your OpsLevel account by following the instructions from [here](https://docs.opslevel.com/docs/graphql#1-create-an-api-token).

### Integration Hub setup

1. Open the **Integration Hub** menu.&#x20;
2. Select the **App Context** tag and search for OpsLevel.&#x20;
3. Click the **Add** button.
4. Add the **Profile name** - this is the name of your OpsLevel instance.
5. Add the **Instance URL** of your OpsLevel account.&#x20;
6. Add the **API Token** for your OpsLevel instance.
7. You can select one or more attributes related to repository assets that Snyk AppRisk can pull from OpsLevel with the following mapping:
   * Catalog name - Identified with `name` in OpsLevel.
   * Category - Identified with `tier.name` in OpsLevel.
   * Lifecycle - Identified with `lifecycle.name` in OpsLevel.
   * Owner - Identified with `owner.name` in OpsLevel.
   * Application - Identified with `product` in OpsLevel.
8. Click the **Done** button.
9. When the connection is established, the status of the OpsLevel integration is changed to **Connected**, and Snyk AppRisk will start enriching repository assets with the data found in OpsLevel.

## Datadog Org Context (Service Catalog)

### Required Parameters

1. Add your Datadog **Profile name**.
2. Add the **API key** for the Datadog instance.
3. Add the **Application Key** along with your organization's API key to grant users access to Datadog's programmatic API. [Here](https://docs.datadoghq.com/account\_management/api-app-keys/) you can find details about the Datadog API and Application key.

### Integration Hub setup

1. Open the **Integration Hub** menu.&#x20;
2. Select the **App Context** tag and search for Datadog Org Context.&#x20;
3. Click the **Add** button.
4. Add the **Profile name** - this is the name of your Datadog instance.
5. Add the **API key** for your Datadog instance.
6. Add the **Application key** for your Datadog instance.
7. Add the details of your **Datadog site**.
8. You can select one or more attributes related to repository assets that Snyk AppRisk can pull from OpsLevel with the following mapping:
   * Catalog name - If you select this metadata, it is mandatory to add the **Catalog name key**.
   * Title - If you select this metadata, it is mandatory to add the **Title key**.
   * Category - If you select this metadata, it is mandatory to add the **Category key**.
   * Lifecycle - If you select this metadata, it is mandatory to add the **Lifecycle key**.
   * Owner - If you select this metadata, it is mandatory to add the **Owner key**.
   * Application - If you select this metadata, it is mandatory to add the **Application key**.
9. Click the **Done** button.
10. When the connection is established, the status of the Datadog Org Context integration is changed to **Connected**, and Snyk AppRisk will start enriching repository assets with the data found in Datadog Org Context.

\
