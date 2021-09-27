# Snyk customer onboarding - team plan

Use this information to onboard your colleagues with the Snyk on the [Team plan](https://snyk.io/plans/), allowing your team to make the best use of Snyk to find and fix project vulnerabilities by:

* Developing rollout strategies for best adoption of Snyk to multiple users in your company, 
* Making best use of reports, prioritizations and remediations to get quick results from Snyk adoption.

This documentation assumes you have used Snyk, and are familiar with core Snyk functions.

## More details

* [Introduction to the Snyk Team plan](https://support.snyk.io/hc/en-us/articles/360018365737)
* [Getting started with Snyk products](https://docs.snyk.io/getting-started/getting-started-snyk-products) 
* [Quick start guide and general overview](https://www.youtube.com/watch?v=PCculVmSPtg&list=PLkgGOmXHS2S3txqFVxiVNVt2AYIXmaH6c&index=6) \(video\)

## Rollout strategy: good practices

These good practice suggestions help you design your rollout plan to align to your team’s success criteria and targets.

You may use different approaches, depending on your situation and the languages of your projects that you want to test.

## Work with early adopters

Consider initially working with early-adopter teams and users, testing and monitoring their projects using the CLI. This should give you some quick results; you can then repeat the process with other teams and users.

Starting with a few users avoids issues you may encounter if you try to roll Snyk out all at once across your company.

**More details**

* Video: [How to find vulnerabilities using your CLI](https://www.youtube.com/watch?v=h9-pP6nOldo&list=PLkgGOmXHS2S3txqFVxiVNVt2AYIXmaH6c&index=2). 
* Documentation: [Snyk CLI](https://docs.snyk.io/snyk-cli) 

## Integrate with an SCM integration

Assuming these integrations are valid for your company, we recommend setting up early adopters with a Source Code Management \(SCM\) integration such as GitHub.

Starting with this integration makes it easy for developers to find and fix vulnerabilities, and to get alerted when new vulnerabilities have been discovered. After this is successful, these users can be used as internal advocates and can demonstrate the benefits of using Snyk.

**NOTE**

This approach works for projects using Node, Java, Python and Ruby. If support is needed for Go, PHP or .Net, the CLI should be used as the primary method of testing and monitoring projects.

\*\* For Scala & Gradle \(only\) we recommend scanning the projects only in the CLI or CI/CD, to get accurate results.

**More details**

* [Snyk SCM integration: good practices](https://support.snyk.io/hc/en-us/articles/360018010597)
* [How to add your Git Repos to Snyk](https://www.youtube.com/watch?v=Krs8IOGy87Q&list=PLkgGOmXHS2S3txqFVxiVNVt2AYIXmaH6c&index=2&t=4s) \(video\)

## Integrate into your CI/CD pipeline

A further step in rollout is to integrate Snyk into your CI/CD pipeline across all teams and projects, to add Snyk to your build system.

This approach allows you to get coverage of your codebase quickly, as CI/CD pipelines are often common and usually maintained by relatively few people. This will give you visibility into the state of your Open Source Vulnerability exposure.

You can also decide if you want to simply report on the vulnerabilities, or break the build when a vulnerability is discovered. A typical approach is to start with reporting, and then once you have visibility, you can roll out the next phase of addressing the vulnerabilities.

**More details**

* [Snyk CI/CD Integration: good practices](https://docs.snyk.io/getting-started/snyk-billing-plan-onboarding/snyk-cicd-integration-good-practices)
* [Snyk CI/CD integration examples on GitHub](https://github.com/snyk-labs/snyk-cicd-integration-examples) 

Snyk automatically notifies you when new issues are found in the projects you're monitoring, to help make you aware of new risks:

You can customize the emails all your organization’s members receive, and individual users can set defaults in their own account settings.

## More details

* [Notification management](https://docs.snyk.io/user-and-group-management/notifications/notification-management)
* [How to configure your Snyk notification settings](https://www.youtube.com/watch?v=MyLgmcHUrL4&list=PLkgGOmXHS2S3txqFVxiVNVt2AYIXmaH6c&index=5) \(video\)



