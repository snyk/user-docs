# CI/CD integrations FAQ

Find answers to the most common questions about integrating Snyk API & Web into your CI/CD pipeline, all at a glance.

## Which CI/CD providers do you support?

You can integrate Snyk API & Web with any CI/CD provider. Two primary methods are supported:

* **Snyk API & Web CLI:** This is the recommended method. The CLI is a convenient wrapper for the API that is easy to use in shell scripts for tasks like starting a scan or adding a target.
* **Direct API:** For more complex integrations or to access the full range of features with fine-grained control, you can interact with the Snyk API & Web API directly.

## Is it possible to block the pipeline if vulnerabilities are detected?

Yes. Blocking mode code examples are provided for [GitLab](integrate-snyk-api-web-with-gitlab-cicd.md), [Bitbucket](integrate-snyk-api-web-with-bitbucket-pipelines.md), and [GitHub Actions](integrate-snyk-api-web-with-github-actions.md) in their respective guides and the GitHub repository. These examples contain a script that checks the scan results for high-severity findings and blocks the pipeline (or causes the deployment to fail) if any are found.

If you use a different CI/CD tool, you can adapt the logic from these examples to fit your pipeline.

## My scan is taking too long. What can I do?

If your pipeline scans are too slow, you have several options to optimize their speed. A balanced approach is recommended: use less comprehensive scans earlier in the development cycle (on feature branches) and more comprehensive scans in pre-production.

Here are a few ways to reduce scan time:

* **Adjust the scan profile:** You can increase scan speed by choosing a less comprehensive scan profile.
* **Enable incremental scans:** This focuses the scan only on new or updated parts of your application, which is significantly faster than a full scan.
* **Leverage partial scopes:** You can configure the scan to check only the specific endpoints or URLs that have been impacted by recent code changes.
* **Optimize tests:** You can create a custom scan profile and select only a specific list of vulnerabilities to test for. This is faster than running the default comprehensive profile, as it avoids running tests for vulnerability types that may not be a priority for a specific pipeline.
* **Vary scan strategies by environment:** A scan should be faster the closer it is to the developer. Reduce the comprehensiveness of scans in dev or feature-branch environments to provide fast feedback. Then run more comprehensive scans in your pre-production environment. You can also set different severity policies for each environment. For example, you might block the pipeline only for critical findings in dev, but block the pipeline for both critical and high findings in QA.

## How do I specify a scan profile in my CI/CD pipeline?

You can specify a scan profile in two ways:

* **To set a default:** You can update the target's settings in the Snyk API & Web UI or via an API call. This profile is used for all scans on that target unless overridden.
* **To override for one scan:** For a one-time scan in a CI/CD pipeline, you can use the Snyk API & Web CLI to specify the profile in your YAML configuration file or make a direct API call to the /scan-now/ endpoint.

For technical details, see the Snyk API & Web API documentation.

## What is the difference between running SAST and DAST in a CI/CD pipeline?

You typically run a SAST (Static Application Security Testing) scan on your source code before the application is built.

A DAST (Dynamic Application Security Testing) scan, like Snyk API & Web, requires a running application to test. Therefore, you must run DAST scans at a later stage in your pipeline, after your application has been deployed to a test or staging environment. DAST scans also typically take more time than SAST scans.

## How can I dynamically manage target settings in my CI/CD pipeline?

Use the Snyk API & Web CLI for all CI/CD configurations. It is flexible and allows you to script any action, such as creating, updating, or deleting targets, and initiating scans.

For detailed commands, see the Snyk API & Web CLI Targets Reference and CLI Scans Reference.

## How can I label findings with metadata like a commit hash or repo name?

The recommended way to add CI/CD metadata is to set labels on specific findings after a scan completes. This allows you to trace vulnerabilities back to their source, such as the exact commit, branch, or repository.

The workflow from your pipeline script is:

1. Start the scan and wait for it to finish.
1. Get the list of new findings via the Snyk API & Web API.
1. Loop through the new findings and make an API call to add your desired labels (for example, `repo:my-project`, `commit:a1b2c3d4`) to each one.

You can find the specific endpoints for this in the Finding Labels Reference.
