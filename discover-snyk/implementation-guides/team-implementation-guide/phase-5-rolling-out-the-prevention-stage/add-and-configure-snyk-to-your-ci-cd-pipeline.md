---
nav_context: agnostic
description: >-
  How to add Snyk to your CI/CD pipeline as a gatekeeper that fails builds on
  new vulnerabilities
---

# Add and configure Snyk to your CI/CD pipeline

Using Snyk as a gatekeeper in your build pipeline prevents the introduction of new vulnerabilities, based on the "fail" criteria you set.

After your teams understand the vulnerabilities in their applications and develop a process for fixing them early in the development cycle, you can configure Snyk to fail your builds if Snyk detects vulnerabilities, to prevent introducing vulnerabilities into your applications.

## No import requirement

A benefit of adding tests to your pipeline is that you do not need to import the repository to Snyk using the source control integration (which is required for Snyk PR Checks). You can also use pipeline tests as an additional gate, even if you are testing PRs, to further decrease the chance of new vulnerabilities entering your production builds.

## Pipeline options

When adding Snyk to a build pipeline, there are common options:

* Using the specific [pipeline integration](https://docs.snyk.io/developer-tools/integrations/snyk-ci-cd-integrations) for your tool.
* Using the [Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli) and running the specific commands directly.

Each option has benefits. An existing pipeline integration can be faster and easier to configure, but the Snyk CLI gives you a greater range of options and flexibility in your "fail" criteria.

## Pipeline test filters

When you run a test in your pipeline, filters determine what results in a test passing or failing. The most common is the severity threshold, which you can use to fail a build only if there are High or Critical severity vulnerabilities.

## CLI supporting tools

If you use the Snyk CLI in your pipeline, a range of supporting [Snyk Tools](https://docs.snyk.io/developer-tools/snyk-apps/snyk-tools) provide additional functionality, including [`snyk-filter`](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter), which you can use for more complex "fail" criteria, such as "fail if more than three High severity vulnerabilities are found."

## Related information

The [CI/CD Best Practices](https://www.youtube.com/watch?v=6QS9gRQ0WVU) webinar covers CI/CD checks in more detail and includes an example of how to introduce this feature gradually.
