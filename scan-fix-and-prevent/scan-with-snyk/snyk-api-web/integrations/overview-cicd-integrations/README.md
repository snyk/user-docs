# CI/CD integrations overview

This guide provides everything you need to get started integrating Dynamic Application Security Testing (DAST) into your CI/CD pipeline.

## Why integrate DAST into CI/CD

Integrating DAST into your CI/CD pipeline is the most scalable way to automate security. It serves as a central point to guarantee security coverage across all your applications, provides a readily available version of your app for testing, and fits directly into your developers' existing workflows.

Whether your goal is to monitor applications at scale, cause vulnerable deployments to fail before they go live, or provide early security feedback to developers, the Snyk API & Web integration is designed to be flexible enough for any environment.

To begin, find your specific provider in the guides below or use the Snyk API & Web CLI for any other tool.

## CI/CD provider guides

* [GitHub Actions](integrate-snyk-api-web-with-github-actions.md)
* [Jenkins](integrate-snyk-api-web-with-jenkins.md)
* [GitLab CI/CD](integrate-snyk-api-web-with-gitlab-cicd.md)
* [Bitbucket Pipelines](integrate-snyk-api-web-with-bitbucket-pipelines.md)

## Connect to any CI/CD tool

While detailed guides are provided for popular platforms, you can integrate Snyk API & Web into any CI/CD provider using the Snyk API & Web CLI. This command-line tool is simple to script and a versatile tool for integration into various development and automation environments.

Visit the [CLI key features](../cli-key-features.md) article to learn more about the CLI capabilities.

## Best practices

### Choose your integration strategy: blocking vs. non-blocking

The first step is to decide on your goal. The CI/CD integration is flexible to support two primary use cases:

**Non-blocking scans (for monitoring)**

This approach uses the Snyk API & Web CLI to actively monitor your pipeline for security vulnerabilities. This approach is ideal for AppSec teams who want to continuously monitor new changes introduced by developers for compliance and visibility.

Scans are triggered in the pipeline but do not block the pipeline if vulnerabilities are found. This allows you to track test results for each application version over time without causing the deployments to fail.

**Blocking scans (for prevention)**

This approach uses the Snyk API & Web CLI to actively prevent security vulnerabilities from reaching production.

While no direct support using the CLI is available, you can use the code examples in this guide to configure blocking directly in your own CI/CD pipeline script. For example, you can add a condition that blocks the pipeline if the scan discovers any high or critical severity issues. This gives you full control to define your security standards as code and reduce risk.

Once the pipeline is blocked, you can check details of findings in the Snyk API & Web dashboard, or in the issue tracker if it is integrated with Snyk API & Web.

You can perform the same actions using the Snyk API & Web API.

### Leverage incremental scans in CI/CD

To balance speed with effective coverage, especially for blocking scans, use incremental scans. This feature shortens scan times by focusing on new URLs, updated URLs, and URLs with existing findings.

### Scan the right application branch

The solution is designed to integrate seamlessly into your branching strategy, whether you use feature branches, release branches, or trunk-based development.

Consider which branches represent the most logical points for security feedback and monitoring. You can also scan an application in an ephemeral environment.

## Frequently asked questions

For answers to common questions about CI/CD integrations, visit the [CI/CD integrations FAQ](cicd-integrations-faq.md).
