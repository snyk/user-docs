# CI/CD integrations overview

This guide provides everything you need to start integrating dynamic application security testing (DAST) into your CI/CD pipeline.

## Why integrate DAST into CI/CD

Integrating DAST into your CI/CD pipeline is the most scalable way to automate security. It serves as a central point to guarantee security coverage across all your applications, provides a readily available version of your app for testing, and fits directly into your developers' existing workflows.

Whether your goal is to monitor applications at scale, fail vulnerable deployments before they go live, or provide early security feedback to developers, the Snyk API & Web integration is flexible enough for any environment.

To begin, find your provider in the following guides or use the Snyk CLI for any other tool.

## CI/CD provider guides

* [GitHub Actions](integrate-snyk-api-web-with-github-actions.md)
* [Jenkins](integrate-snyk-api-web-with-jenkins.md)
* [GitLab CI/CD](integrate-snyk-api-web-with-gitlab-cicd.md)
* [Bitbucket Pipelines](integrate-snyk-api-web-with-bitbucket-pipelines.md)

## Connect to any CI/CD tool

Detailed guides are provided for popular platforms, but you can integrate Snyk into any CI/CD provider using the Snyk CLI. This command-line tool is simple to script and versatile for integration into various development and automation environments.

Visit the [CLI key features](../cli-key-features.md) article to learn more about the CLI capabilities.

## Best practices

### Choose your integration strategy: blocking or non-blocking

The first step is to decide on your goal. The CI/CD integration supports two primary use cases:

**Non-blocking scans (for monitoring)**

This approach uses the Snyk CLI to monitor your pipeline for security vulnerabilities. It is ideal for AppSec teams who want to continuously monitor new changes introduced by developers for compliance and visibility.

Snyk triggers scans in the pipeline but does not block the pipeline if it finds vulnerabilities. You can track test results for each application version over time without failing the deployments.

**Blocking scans (for prevention)**

This approach uses the Snyk CLI to prevent security vulnerabilities from reaching production.

The CLI provides no direct support for blocking, but you can use the code examples in this guide to configure blocking directly in your own CI/CD pipeline script. For example, you can add a condition that blocks the pipeline if the scan discovers any high or critical severity vulnerabilities. This gives you full control to define your security standards as code and reduce risk.

After the pipeline is blocked, you can check the details of findings in the Snyk dashboard, or in the issue tracker if it is integrated with Snyk.

You can perform the same actions using the Snyk API.

### Use incremental scans in CI/CD

To balance speed with effective coverage, especially for blocking scans, use incremental scans. This feature shortens scan times by focusing on new URLs, updated URLs, and URLs with existing findings.

### Scan the right application branch

The solution integrates seamlessly into your branching strategy, whether you use feature branches, release branches, or trunk-based development.

Consider which branches represent the most logical points for security feedback and monitoring. You can also scan an application in an ephemeral environment.

## Frequently asked questions

For answers to common questions about CI/CD integrations, visit the [CI/CD integrations FAQ](cicd-integrations-faq.md).
