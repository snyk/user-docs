# Git repositories and CI/CD comparisons

The most popular integrations for Snyk implementation are [Git Repository](git-repositories-scms-integrations-with-snyk/) and [CI/CD pipeline](snyk-ci-cd-integrations/).

<figure><img src="../.gitbook/assets/scm-ci-cid.png" alt="Snyk integrations"><figcaption><p>Snyk Integrations</p></figcaption></figure>

* **Git repository**: Improve application security in your Git repository, preventing vulnerable code from entering your codebase and getting quick visibility for your vulnerabilities.
* **CI/CD**: Keep your applications secure by preventing deployment of vulnerable applications or components (registries), adding Snyk in the build as a step of the pipeline.

## Which integrations to choose?

You can decide to implement either Git repository or CI/CD, or both. Both have advantages and disadvantages; your choice will depend on your team's flows and organizational processes.

The following considerations explain the benefits of each type of integration.

### Git repository considerations

Use Git repository integrations to Improve the security of your code and deployed applications.

**Considerations:**

* Easier to set up and maintain.
* Allows scanning and visibility earlier in the software development lifecycle through:
  * Automatic daily rescanning of all imported Projects.
  * Checking all submitted PRs for security issues.
  * Generating dependency upgrade PRs to deal with technical debt.
  * Generating fix PRs for existing vulnerabilities that have not been addressed.
* More friendly experience for developers.
* Does not take resources from your CI/CD pipeline.

For more details, see [Snyk Git repository integration: deployment recommendations](git-repositories-scms-integrations-with-snyk/introduction-to-git-repository-integrations/snyk-scm-integration-good-practices.md).

### CI/CD

Use CI/CD integrations to accomplish the following:

* Keep your code and deployed applications secure
* Give visibility to components that are pushed to production by either breaking builds and reporting to Snyk or only reporting to Snyk.

**Considerations:**

* Some package managers require local context and are better run within your environment, including Scala, Gradle, Go modules, Artifactory, and Nexus.
* More granular options to block.
* Strong gatekeeper.
* Best practice for container and Infrastructure as Code scans. For IaC, get more accurate results by scanning a Terraform plan file rather than the untemplated IaC declarations.

For more details, see [Snyk CI/CD Integration deployment and strategies](snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/).
