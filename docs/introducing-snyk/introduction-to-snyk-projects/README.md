# Introduction to Snyk targets and projects

## Targets

Targets represent an external resource Snyk has scanned through an integration, the CLI, UI or API.

Targets may represent a SCM repository, a Kubernetes workload, or other scannable resources external to Snyk. All [Projects](./#project) Snyk creates, are associate to a parent Target. One target may relate to many projects. The structure of the target depends on the [origin](./#origin).\
\
Targets appear on the **Projects** menu on the Snyk dashboard:

![](<../../.gitbook/assets/image (65) (1) (2) (1).png>)

{% hint style="info" %}
Targets also appear on the [Snyk API](https://apidocs.snyk.io/?version=2022-02-16%7Ebeta#tag--Targets).
{% endhint %}

## Origin

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

## Projects

Projects define the items Snyk scans at a given Target. A project includes:

* A scannable item external to Snyk.
* Configuration to define how to run that scan.

Projects appear on the **Projects** menu on the Snyk dashboard, and on the [Snyk API](https://apidocs.snyk.io/?version=2022-02-16%7Ebeta#tag--Projects):

![](<../../.gitbook/assets/image (76) (1) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png>)

## Targetfile

The specific item to scan in a target, such as a pom file in a GitHub repo.

[Snyk Code](https://docs.snyk.io/snyk-code) scans do not use targetfiles.

## Type

The scanning method to use for this project, such as static application security testing , [SAST](https://snyk.io/learn/application-security/sast-vs-dast/) for scanning using Snyk Code or Maven for a Maven project using Snyk Open Source). Part of the configuration for scanning.

## Introducing the Snyk Dashboard

Use the [Dashboard](https://docs.snyk.io/snyk-web-ui/getting-started-with-the-snyk-web-ui#dashboard) to gain a different perspective on project work that needs to be handled and to get an overview of the most vulnerable projects in your organization.
