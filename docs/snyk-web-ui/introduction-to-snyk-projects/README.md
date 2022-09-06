# Snyk Projects

### Targets

Targets represent an external resource Snyk has scanned: a code repository, a Kubernetes workload, or other scannable resources external to Snyk.&#x20;

All [Snyk Projects](./#projects) are associated to a parent Target. One Target may relate to many Projects. The structure of the Target depends on the [origin](./#origin).\
\
Targets appear on the **Projects** menu on the Snyk dashboard:

![](<../../.gitbook/assets/image (65) (1) (2) (1).png>)

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

![](<../../.gitbook/assets/image (71) (3).png>)

### Snyk Projects

Snyk Projects define the items Snyk scans at a given Target. A Project includes:

* A scannable item external to Snyk.
* Configuration to define how to run that scan.

Projects appear on the **Projects** menu on the Snyk dashboard, and on the [Snyk API](https://apidocs.snyk.io/?version=2022-02-16%7Ebeta#tag--Projects):

![](<../../.gitbook/assets/image (12) (1) (2) (1) (1) (1) (1).png>)

## Targetfile

The specific item to scan in a target, such as a **pom.xml** file in a GitHub repo.

{% hint style="info" %}
[Snyk Code](https://docs.snyk.io/snyk-code) scans do not use targetfiles.
{% endhint %}

## Type

The scanning method to use for this project, such as Static Application Security Testing ([SAST](https://snyk.io/learn/application-security/sast-vs-dast/)) for scanning using Snyk Code, or Maven for a Maven project using Snyk Open Source. Part of the configuration for scanning.
