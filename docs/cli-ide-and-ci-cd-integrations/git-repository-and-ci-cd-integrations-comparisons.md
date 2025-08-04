# Git repositories and CI/CD comparisons

[Git repository](../developer-tools/scms/organization-level-integrations/) and [CI/CD pipeline](../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/) integrations are commonly used.&#x20;

As the following diagram illustrates, you can use a Git repository (SCM) integration to check for issues, including vulnerabilities and license issues, and to prevent pull requests based on policies. You can start by testing and fixing in your development environment, and then test, fix, and monitor using a Git repository integration. Use a Git repository integration to improve application security in your Git repository, preventing vulnerable code from entering your codebase and getting quick visibility for your vulnerabilities.

<figure><img src="../.gitbook/assets/scm-ci-cid.png" alt="Snyk integrations"><figcaption><p>Snyk Integrations</p></figcaption></figure>

As shown in the diagram, you can add Snyk in the build as a step in your pipeline with CI/CD integration to test, fix, and monitor. Use Snyk in your build to keep your applications secure by preventing the deployment of vulnerable applications or components (registries).

You can implement either Git repository or CI/CD integration, or both. Your implementation will depend on your team's flows and organizational processes.

## Git repository considerations

Use Git repository integrations to improve the security of your code and deployed applications.

Git repository integration allows scanning and visibility early in the software development lifecycle through:

* Automatic daily rescanning of all imported Open Source Projects
* Checking all submitted PRs for security issues
* Generating dependency upgrade PRs to deal with technical debt
* Generating fix PRs for existing vulnerabilities that have not been addressed

Git repository integration does not take resources from your CI/CD pipeline.

For more details, see [Deployment recommendations for SCM integrations](../developer-tools/scms/deployment-recommendations.md).

## CI/CD considerations

Use CI/CD integrations to accomplish the following:

* Keep your code and deployed applications secure
* Give visibility to components that are pushed to production by either breaking builds and reporting to Snyk or only reporting to Snyk.

When implementing CI/CD integration, consider the following:

* Some package managers require local context and are better run within your environment, including Scala, Gradle, Go modules, Artifactory, and Nexus.
* CI/CD integrations offer granular options to block the build, providing a strong gatekeeper.
* CI/CD integration is a best practice for container and infrastructure as code scans. For IaC, to get more accurate results, scan a Terraform plan file rather than the untemplated IaC declarations.

For more details, see [Snyk CI/CD Integration deployment and strategies](snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/).
