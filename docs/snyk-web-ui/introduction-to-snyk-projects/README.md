# Snyk Projects

### Targets

Targets represent an external resource Snyk has scanned: a code repository, a Kubernetes workload, or other scannable resources external to Snyk.&#x20;

All [Snyk Projects](./#projects) are associated to a parent Target. One Target may relate to many Projects. The structure of the Target depends on the [origin](./#origin).

Targets appear in the **Projects** menu on the Snyk dashboard.

<figure><img src="../../.gitbook/assets/Targets 1.png" alt="Screenshot showing projects&#x27; target in the Snyk UI"><figcaption></figcaption></figure>

Snyk provides pagination to improve the page loading time for Projects page requests and filtering, which is particularly relevant if you have hundreds of thousands of Projects to scan.&#x20;

Use **Group by** to collect your Projects by target or leave them ungrouped.

Use **Sort by** to list your Projects by severity, by how recently they were imported, or in alphabetical order.

&#x20;![](../../.gitbook/assets/projects\_group-sort\_20sept2022.png)

{% hint style="info" %}
Targets also appear on the [Snyk API](https://apidocs.snyk.io/?version=2022-02-16%7Ebeta#tag--Targets).
{% endhint %}

### Origin

The Target ecosystem, such as CLI, GitHub, or Kubernetes.

Possible values are:

* acr
* api
* artifactory-cr
* aws-config
* aws-lamba
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

Origins are a property of [Targets](./#targets) and appear in the Projects menu, as an icon by the target name.

<figure><img src="../../.gitbook/assets/Targets 2.png" alt="Screenshot showing projects&#x27; origin in the Snyk UI"><figcaption></figcaption></figure>

### Snyk Projects

Snyk Projects define the items Snyk scans at a given Target. A Project includes:

* A scannable item external to Snyk.
* Configuration to define how to run that scan.

Projects appear on the **Projects** menu on the Snyk dashboard, and on the [Snyk API](https://apidocs.snyk.io/?version=2022-02-16%7Ebeta#tag--Projects):

<figure><img src="../../.gitbook/assets/Targets 3.png" alt="Screenshot showing projects listed under a target in Snyk UI"><figcaption></figcaption></figure>

## Targetfile

The specific item to scan in a target, such as a **pom.xml** file in a GitHub repo.

{% hint style="info" %}
[Snyk Code](https://docs.snyk.io/snyk-code) scans do not use targetfiles.
{% endhint %}

## Type

The scanning method to use for this project, such as Static Application Security Testing ([SAST](https://snyk.io/learn/application-security/sast-vs-dast/)) for scanning using Snyk Code, or Maven for a Maven project using Snyk Open Source. Part of the configuration for scanning.
