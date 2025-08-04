# Choose rollout integrations

## **SDLC integration points**

Snyk offers many integrations that seamlessly integrate into every stage of the SDLC.&#x20;

Many businesses typically roll out automated solutions first, then slowly introduce tools to enable the developers. In addition, gating features are gradually turned on over a period of time to minimize disruption.

{% hint style="info" %}
As using multiple integrations can result in duplicate reporting of issues, you do not initially need to implement more than one integration type. For example, you can start by importing everything with Git repositories, then later use the CI/CD view for fine-grained detail (potentially removing the source control integration if both views are not desired).
{% endhint %}

## Integration types

Below are typical early integrations.

### Source Code Management (SCM) integrations

Integrations with popular version control platforms like GitHub, GitLab, Azure Repos, and Bitbucket seamlessly integrate Snyk security checks into the code review process. This ensures that potential vulnerabilities are identified and addressed before the code is merged into the main branch. Important features include:

* Daily testing/monitoring of a specified branch (typically "development" branch),&#x20;
* (optional) Pull Request/Merge Request checks against any branch of the repository.
* (optional) Automated dependency upgrades and automated security fix upgrades using pull requests.

Advantages:

* Visibility into repository security posture
* Automatic Scan on code change
* Immediate feedback on issues for the developer
* Onboarding of repositories can be configured using the UI
* Supports Cloud Repositories on the Team plan

For more details, see [Git repositories (SCMs)](../../../developer-tools/scm-integrations/organization-level-integrations/).

{% hint style="info" %}
If you have a non-cloud-facing or your own instance of a Git SCM:

* Consider deploying a [Snyk Broker](../../enterprise-setup/snyk-broker/) for Snyk to communicate with your repositories, which would also require a Snyk Enterprise Plan.&#x20;
* Enterprise customers can enable and manage Snyk Broker using the API.&#x20;

[Paid services](../../../snyk-data-and-governance/snyk-terms-of-support-and-services-glossary/) can be engaged to assist in Broker deployments.
{% endhint %}

### Continuous Integration/Continuous Deployment (CI/CD) pipeline integrations

Integrating Snyk into CI/CD pipelines, such as Jenkins, Travis CI, or CircleCI, automates security checks during the build and deployment process. This ensures that vulnerabilities are detected early in the software development lifecycle and prevents their propagation into production. Typical features include:

* (Optional) Ability to passively monitor results during build and view results in Snyk
* (Optional) Ability to test and potentially break the build based on criteria you specified
* Integration can be achieved with specific Marketplace plugins or more generally, with the CLI as part of your pipeline script.

Advantages:

* Assess local code vulnerabilities
* Full control over testing (which tests to run, where in the build script)
* Can automate using CI/CD

For more details, see [Snyk CI/CD integrations](../../../developer-tools/snyk-ci-cd-integrations/).

### Integrated Development Environment (IDE) integrations

IDE integrations like Visual Studio Code, IntelliJ IDEA, and Eclipse allow developers to access Snyk's security features directly within their coding environment. This enables real-time scanning and issue remediation as developers write code at the earliest possible stages.&#x20;

For more details, see [Use Snyk in your IDE](../../../developer-tools/snyk-ide-plugins-and-extensions/).

## Considerations for import strategies&#x20;

<table><thead><tr><th width="200">Project Import Strategy</th><th>Considerations</th><th>Advantages</th><th>Disadvantages</th></tr></thead><tbody><tr><td>CLI (automated CI/CD)</td><td>Must be configured for each application within CI/CD.</td><td><ul><li>Can select what to test and when (i.e. which package managers, where in the process, which language to analyze.</li><li>May need development effort for integration.</li></ul></td><td>It requires configuration per application.</td></tr><tr><td>CLI (Run locally by user)</td><td>User can use CLI to perform testing locally while working on an application, very configurable per scan type.</td><td>Local use case</td><td>Not meant for visibility or automation. Can require buildable code or dependencies to be installed (For example Gradle without Lockfile, Scala).</td></tr><tr><td>Git Code Repository Integration</td><td>Onboarding and daily monitoring: rapid vulnerability assessment across application portfolio.</td><td><p></p><ul><li>Continuous monitoring of repositories (even when you are not working on it).</li><li> Centralized visibility for teams.</li><li>Monitors specified branch</li><li>Code does not need to be built.</li></ul></td><td><ul><li>Can be initiated via UI</li><li>Some languages/package managers have better resolution utilizing the CLI (Gradle without lockfile, Scala).</li></ul></td></tr><tr><td></td><td>Pull request (PR)/merge request (MR)  scanning</td><td>Immediate feedback on introduced issues on the PR/MR against any branch on repository.</td><td>Configurable rules for pass/fail</td></tr></tbody></table>

## Additional considerations

### Infrastructure as Code

For Snyk Infrastructure as Code, it is common that your Terraform or YAML configuration files are held in your SCM, but they may be in a separate area or repository. As a result, consider if there are other areas you need to import. You may also want to integrate with Terraform Cloud (if applicable) to enable Snyk tests as part of your "Terraform run" processes.

For complex environments, modules, and highly templated implementations, utilizing the CLI on your Terraform Plan file may provide the best results.

### CR (Container Registries)

Snyk also integrates with various [Container Registries](../../../scan-with-snyk/snyk-container/container-registry-integrations/) to enable you to import and monitor your containers for vulnerabilities. Snyk tests the containers you have imported for any known security vulnerabilities found at a frequency you control.
