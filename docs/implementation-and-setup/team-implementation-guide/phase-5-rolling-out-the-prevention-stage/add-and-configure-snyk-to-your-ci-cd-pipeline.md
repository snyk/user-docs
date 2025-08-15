# Add and configure Snyk to your CI/CD pipeline

Using Snyk as a gatekeeper in your build pipeline prevents the introduction of new vulnerabilities, based on the "fail" criteria you set.

After your teams understand the vulnerabilities in their applications and develop a process for fixing them early in the development cycle, you can configure Snyk to fail your builds if vulnerabilities are detected, to prevent introducing vulnerabilities into your applications.

## No import requirement

A benefit of adding tests to your pipeline is that you do not need to import the repository to Snyk using the source control integration (which is required for Snyk PR Checks). It can also be used as an additional gate, even if you are testing PRs, to further decrease the chance of new vulnerabilities entering your production builds.

## Pipeline options

When adding Snyk to a build pipeline, there are common options:&#x20;

* Using the specific [pipeline integration](../../../developer-tools/snyk-ci-cd-integrations/) for your tool.
* Using the [Snyk CLI](../../../developer-tools/snyk-cli/) and running the specific commands directly.&#x20;

Each option has benefits - using an existing pipeline integration may be faster and easier to configure, but using the Snyk CLI will give you a greater range of options and flexibility in your "fail" criteria.&#x20;

## Pipeline test filters

When running a test in your pipeline, there are filters available to determine what would result in a test passing or failing. The most common of these is "severity threshold", where you can specify to only fail a build if there are High or Critical severity vulnerabilities.

## CLI supporting tools

If you use the Snyk CLI in your pipeline, a range of supporting [Snyk Tools](../../../scan-with-snyk/snyk-tools/) provide additional functionality, including [`snyk-filter`](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md), which can be used for more complex "fail" criteria, such as “fail if more than three High severity vulnerabilities are found”.

## Also see

Here is a link to a webinar that covers CI/CD checks in more detail and includes an example of how you can gradually introduce this feature: [CI/CD Best Practices](https://www.youtube.com/watch?v=6QS9gRQ0WVU).
