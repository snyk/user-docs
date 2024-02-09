# Snyk Projects

Snyk Project information appears in the **Projects** listing, which you can display from the menu on the Snyk dashboard. The filters you can add depend on the **Group by** option you choose from the pulldown on the right. To filter by Origin or source, use an Integrations filter.

{% hint style="info" %}
After filters have been applied to the Project listing page, you can bookmark the URL and share it with other users in the Organization. This allows all users to see the same view of the page.
{% endhint %}

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-24 at 09.09.25.png" alt="Snyk Projects listing grouped by Target"><figcaption><p>Snyk Projects listing grouped by Target</p></figcaption></figure>

**Snyk Projects concepts** include the following:

* [Target](./#target)
* [Origin](./#origin)
* [Project](./#project)
* [Targetfile](./#targetfile)
* [Type](./#type)

## Target

Projects are held in a Target. A Target represents an external resource Snyk has scanned: a code repository, a Kubernetes workload, or other scannable resource external to Snyk. Snyk creates a Target when you request to import a Project or scan using the CLI. If the import fails or finds nothing, the Target will be empty.

When you select **Group by target**, Snyk Targets appear in the **Projects** listing. You can also find Targets using the Snyk REST API endpoint [Get targets by org ID](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#get-/orgs/-org\_id-/targets).

Click the **dots** at the far right of each Target on the list to see the **available actions**; these include **deleting** a target.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-24 at 08.59.20.png" alt="Snyk Target and Projects in that Target"><figcaption><p>Snyk Target and Projects in that Target</p></figcaption></figure>

Each Snyk Project is associated with a parent Target. One Target may include many Projects. The structure of the Target depends on the Origin.

The grouping option controls whether the filtering attributes are applied at the Target or at the Project level. **Group by none** (ungrouped) lets you apply [tags](../introduction-to-snyk-projects/project-tags.md) and [filtering attributes at the Project level](project-attributes.md) to the individual Projects.

Snyk provides both pagination to improve the page loading time for Projects page requests and filtering, which is particularly helpful if you have hundreds of thousands of Projects to scan.

Use **Sort by** (pull down on the far right) to sort the **Projects** listing by severity, by how recently the Projects were imported, or in alphabetical order.

<figure><img src="../../.gitbook/assets/image (2) (5).png" alt="Sorting attributes available when grouping by Target"><figcaption><p>Sorting attributes available when grouping by Target</p></figcaption></figure>

## Origin or source

The Origin defines the Target ecosystem, such as CLI, GitHub, or Kubernetes. Origins are a property of Targets (see preceding section) and appear in the **Projects** listing as an icon next to the Target name.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-24 at 08.59.07.png" alt="Origin icon next to the Target name"><figcaption><p>Origin icon next to the Target name</p></figcaption></figure>

Possible Origin values are:

* acr
* api
* artifactory-cr
* aws-config
* aws-lambda
* azure-functions
* azure-repos
* bitbucket-cloud
* bitbucket-server
* cli
* cloud-foundry
* digitalocean-cr
* docker-hub
* ecr
* gcr
* github
* github-cr
* github-enterprise
* gitlab
* gitlab-cr
* google-artifact-cr
* harbor-cr
* heroku
* ibm-cloud
* kubernetes
* nexus-cr
* pivotal
* quay-cr
* terraform-cloud

## Project

A Snyk Project defines the items, such as manifest files, that Snyk scans for a given Target, with configuration information defining how to run that scan.

Projects appear in the **Projects** listing. You can also find Projects using the Snyk REST API endpoint [Get projects by org ID](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#get-/orgs/-org\_id-/projects).

Use **Group by none** (ungrouped) for better Project visibility and to apply [filtering attributes at the Project level](project-attributes.md).

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-23 at 18.07.46 (1) (1) (1) (1) (1) (1).png" alt="Filtering attributes applied at the Project level"><figcaption><p>Filtering attributes applied at the Project level</p></figcaption></figure>

## Targetfile

The Targetfile is the specific item to scan in a Target, such as a `pom.xml` file in a GitHub repo.

{% hint style="info" %}
[Snyk Code](../../scan-with-snyk/snyk-code/) scans do not use Targetfiles.
{% endhint %}

## Type

The Type is the scanning method to use for a particular Project, such as Static Application Security Testing ([SAST](https://snyk.io/learn/application-security/sast-vs-dast/)) for scanning using Snyk Code, or Maven for a Maven project using Snyk Open Source. This is part of the configuration for scanning.

## Project actions on the Project Listing Page

### Delete, Activate, or Deactivate

To perform an action at bulk on your Projects, first select the Projects and then choose to **Delete**, **Activate,** or **Deactivate** the Projects.&#x20;

<figure><img src="../../.gitbook/assets/Screenshot 2023-11-15 at 10.42.37.png" alt=""><figcaption><p>Deleting Projects in bulk</p></figcaption></figure>

### Test frequency settings

You can set the frequency of testing for each Project.&#x20;

For each entry, you can select the frequency of testing for that Project (`never`, `daily`, or `weekly`) as applicable to the type of Project, Open Source, Code analysis, Container, or IaC:

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-11-15 at 10.42.46.png" alt="Select test frequency"><figcaption><p>Select test frequency</p></figcaption></figure>

</div>

{% hint style="info" %}
The default test frequency and limitations are as follows:

* Open Source: The default is daily.
* Code analysis Projects: The default is weekly; daily is not available. To test your code daily, submit a request to [Snyk Support.](https://support.snyk.io/hc/en-us/requests)
* Container: The default is daily.
* IaC: The default is weekly.
{% endhint %}

Click **Deactivate** to never test, and also remove webhooks and stop showing the Project’s results in reporting.
