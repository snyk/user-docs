---
description: Overview of the tool
---

# Overview

## What does the tool do?

This tool counts and prints a summary of the contributors count for the last 90 days, for any of the following SCMs:

* Azure Devops
* Bitbucket Cloud
* Bitbucket Server
* Github
* Github Enterprise
* Gitlab
* Gitlab Server

{% hint style="info" %}
&#x20;There are some minor differences between the SCMs as to the naming convention. For example : "Orgs" in Github can be "Projects" in Azure and "Workspaces" in Bitbucket. These differences are reflected in the commands that the tool accepts for each SCM.
{% endhint %}

## **How it works**

In two modes : **Scoping usage prior to onboarding** and **Snyk License Consumption**

* **Scoping usage prior to onboarding -  **For users who want to onboard to Snyk and would like to get an estimate of the developer count across their SCMs**.**\
  ****In this mode, the tool fetches all the information from directly from the SCMs, using the credentials provided by the user.
* **Snyk License Consumption - **For users with an existing account at Snyk, who want some clarity and details about their license consumption (number of contributors, names, email and so on).\
  In this mode, the tool fetches the SCM related projects monitored by Snyk, then matches those to the repos on the SCM, and counts the contributors only for those repos/projects.

## Downloading the tool

```
npm i -g snyk-scm-contributors-count
```

Or, you can get the binaries from the [release page](https://github.com/snyk-tech-services/snyk-scm-contributors-count/releases)
