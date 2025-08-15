# Add and configure Snyk in your CI/CD pipeline

Using Snyk as a gatekeeper in your build pipeline prevents the introduction of new vulnerabilities. based on the fail criteria you set.

After your teams understand the vulnerabilities in their applications and develop a process for fixing them early in the development cycle, you can configure Snyk to fail your builds if vulnerabilities are detected, to prevent introducing vulnerabilities into your applications.

## No import requirement

A benefit of adding tests to your pipeline is that you do not need to import the repository to Snyk using source control integration, which is required for Snyk PR Checks. Adding tests to your pipeline can also be used as an additional gate, even if you are testing PRs, to further decrease the chance of new vulnerabilities entering your production builds.

## Pipeline options

When adding Snyk to a build pipeline, you have these common options:&#x20;

* Using the specific [pipeline integration](../../../developer-tools/snyk-ci-cd-integrations/) for your tool.
* Using the [Snyk CLI](../../../developer-tools/snyk-cli/) and running the specific commands directly.&#x20;

Each option has benefits; using an existing pipeline integration may be faster and easier to configure, but using the Snyk CLI will give you a greater range of options and flexibility in your fail criteria.&#x20;

## Pipeline test filters

When you run a test in your pipeline, you can use filters to determine what would result in a test's passing or failing. The most common of these is severity threshold, where you can specify to fail a build only if there are High or Critical severity vulnerabilities.

## CLI supporting tools

When you use the Snyk CLI in your pipeline, you can use a range of supporting [CLI Tools](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/) that provide additional functionality, including:

* [snyk-delta,](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-delta.md) which can be used to compare two sets of results and identify new vulnerabilities, similar to how the PR Checks feature tests for new vulnerabilities only
* [snyk-filter](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md), which can be used for more complex fail criteria, such as `fail if more than three High severity vulnerabilities are found`.

## Additional information

[CI/CD Best Practices](https://www.youtube.com/watch?v=6QS9gRQ0WVU) is a webinar that covers CI/CD checks in more detail and includes an example of how you can gradually introduce this feature.
