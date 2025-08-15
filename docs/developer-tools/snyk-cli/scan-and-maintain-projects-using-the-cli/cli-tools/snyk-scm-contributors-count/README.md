# snyk-scm-contributors-count

## What does this tool do?

This tool counts and prints a summary of the contributors count for the last 90 days for any of the following SCMs:

* Azure Devops
* Bitbucket Cloud
* Bitbucket Server
* GitHub
* GitHub Enterprise
* GitLab
* GitLab Server

{% hint style="info" %}
There are some minor differences between the SCMs as to the naming convention. For example: "Orgs" in GitHub can be "Projects" in Azure and "Workspaces" in Bitbucket. These differences are reflected in the commands that the tool accepts for each SCM.
{% endhint %}

## **How the SCM-Contributors-Count tool works**

The tool works in two modes: **Scoping usage prior to onboarding** and **Snyk License Consumption.**

* **Scoping usage prior to onboarding:** For users who want to onboard to Snyk and would like to get an estimate of the developer count across their SCMs.\
  In this mode, the tool fetches all the information directly from the SCMs, using the credentials provided by the user.
* **Snyk License Consumption (valid only for Bitbucket and Azure):** For users with an existing Snyk account, who want some clarity and details about their license consumption (number of contributors, names, email, and so on).\
  In this mode, the tool fetches the SCM-related projects monitored by Snyk and then matches those to the repos on the SCM and counts the contributors only for those repos/projects.

{% hint style="info" %}
Note that currently the tool **does not** count emails that end with "noreply.github.com".
{% endhint %}

## Downloading the tool

Run the following:

```
npm i -g snyk-scm-contributors-count
```

You can also get the binaries from the [release page](https://github.com/snyk-tech-services/snyk-scm-contributors-count/releases).
