# Project fix advice

## Fix vulnerabilities

Snyk knowledge of the transitive dependencies in your project makes it easy for Snyk to offer fix advice. Snyk can fix vulnerabilities in two ways. Snyk can upgrade the direct dependencies to a vulnerability free version, or it can patch the vulnerability.

Snyk supports the following workflows to help developers fix their vulnerabilities.

1. Snyk generates automatic git pull requests (PRs) using GitHub, GitLab, or Bitbucket PR workflows. Snyk executes this as a nightly workflow.
2. Snyk users use the _**Open a fix PR**_ feature in the Snyk UI.
3. Snyk users use the _**fix this vulnerability**_ feature on a specific issue card in the Snyk UI.

{% hint style="info" %}
In this workshop, we will perform the **fix this vulnerability** option.
{% endhint %}

## Generating automatic pull requests

Snyk generates automatic pull requests for your projects using the source control integration associated with your organization. To check the setting, select the integrations tab at the top of the Snyk UI and select your source control repository. Click the gear icon for your source control and review the settings. This setting can be configured on our per-project basis or for an entire organization.

{% hint style="info" %}
This setting should be set from our initial configuration of GitHub
{% endhint %}

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/auto\_pr\_setting.png)

Snyk will issues PRs for your projects when they are ready. You can view the PR in GitHub (GitLab and Bitbucket have similar views) by select the Pull Request tab at the top of the git repository. Keep in mind, Synk generates PRs overnight, and multiple pull requests exist for different dependencies.

{% hint style="info" %}
SPC will generate a PR given enough time.
{% endhint %}

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/github\_pull\_req\_auto.png)

## Open a fix PR

The Snyk UI offers the ability to create Pull Requests directly from the project. To generate a PR find your SPC project from the list of projects and select the pom.xml file.

Select the Open a fix PR and Snyk will start with the workflow to generate a PR for SPC. Prior to sending the PR to GitHub, the Snyk UI will show you what will be fixed by the PR, any partial fixes, and anything that will not be fixed. Submit the PR to GitHub and visit your GitHub repository Pull Request tab to view the Pull Request.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/open\_pr.png)

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/open\_fix\_pr\_top\_half.png)

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/open\_fix\_pr\_bottom.png)

## Fix this vulnerability

Snyk's _**fix this vulnerability**_ workflow is the same as the _**open a PR fix**_ workflow. When generating a PR, only the specified vulnerability issue is selected in the PR, as shown below. Find the Cross-site Scripting (XSS) vulnerability and start the PR workflow.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/screen-shot-2020-08-22-at-12.32.44-pm.png)

Verify the issue is selected and complete the PR request.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/screen-shot-2020-08-22-at-12.40.36-pm.png)
