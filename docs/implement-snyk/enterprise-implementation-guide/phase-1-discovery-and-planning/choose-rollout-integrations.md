# Choose rollout integrations

## **SDLC integration points**

Snyk offers many integrations to work seamlessly with Snyk in every stage of the SDLC.&#x20;

Many businesses roll out automated solutions first, and then slowly introduce tools to enable the developers. In addition, gating features are gradually turned on over a period of time, to minimize disruption.

{% hint style="info" %}
As using multiple integrations can result in duplicate reporting of issues, you do not initially need to implement more than one integration type. For example, you can start by importing everything with Git repositories, then later use the CI/CD view for fine-grained detail. You can remove the source control integration if both views are not desired.
{% endhint %}

## Integration types

The following are typical early integrations.

### Source Code Management (SCM) integrations

Integrations with popular version control platforms like GitHub, GitLab, Azure Repos, and Bitbucket seamlessly integrate Snyk security checks into the code review process. This ensures that potential vulnerabilities are identified and addressed before code is merged into the main branch. Important features include:

* Daily testing and monitoring of a specified branch, typically the development branch
* (optional) Pull Request/Merge Request checks against any branch of the repository
* (optional) Automated dependency upgrades and automated security fix upgrades through pull requests

The advantages of SCM integrations are:

* Visibility into repository security posture
* Automatic Scan on code change
* Immediate feedback on issues for the developer
* Onboarding repositories can be configured through the UI or [API/API Import Tool](../../../scan-with-snyk/snyk-tools/tool-snyk-api-import/)
* Support for Cloud and Private Code Repositories on the Snyk Enterprise plan

See [Snyk SCM integrations](../../../scm-integrations/organization-level-integrations/) for more details.

If you have an on-premise Git repository, you must consider deploying [Snyk Broker](../../../enterprise-setup/snyk-broker/) for Snyk to communicate with your repositories.

{% hint style="info" %}
Enterprise customers can enable and manage [Snyk Broker](../../../enterprise-setup/snyk-broker/) using the API.&#x20;

[Paid services](../../../working-with-snyk/snyk-terms-of-support-and-services-glossary/) can be engaged to assist in Snyk Broker deployments.
{% endhint %}

### Continuous Integration/Continuous Deployment (CI/CD) pipeline integrations

Integrating Snyk into CI/CD pipelines, such as Jenkins, Travis CI, or CircleCI, automates security checks during the build and deployment process. This ensures that vulnerabilities are detected early in the software development lifecycle and prevents their propagation into production. Typical features include:

* (Optional) Ability to passively monitor results during build and view results in Snyk
* (Optional) Ability to test and potentially break the build if results are found based on the criteria you specified

The advantages of CI/CD integrations are:

* Assess local code vulnerabilities
* Full control over testing: which tests to run and where in the build script
* Automation by CI/CD if desired

See [Snyk CI/CD integrations](../../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/) for more details.

### IDE Integrations

Integrated Development Environment (IDE) integrations like Visual Studio Code, IntelliJ IDEA, and Eclipse allow developers to access Snyk security features directly within their coding environment. This enables real-time scanning and issue remediation as developers write code.&#x20;

See [Snyk IDE plugins and extensions](../../../cli-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/) for more details.

## Considerations for import strategies&#x20;

<table><thead><tr><th width="200">Project Import Strategy</th><th>Considerations</th><th>Advantages</th><th>Disadvantages</th></tr></thead><tbody><tr><td>CLI (automated with CI/CD)</td><td>Has to be configured for each application within CI/CD</td><td><ul><li>Can select what to test and when: which package managers, where in the process, which language to analyze</li><li>May need development effort for integration</li></ul></td><td>Requires configuration per application.</td></tr><tr><td>CLI (run locally by user)</td><td>Can be used to perform testing locally while the developer is working on an application, very configurable per scan type.</td><td>Local use case</td><td>Not meant for visibility or automation. Can require buildable code or dependencies to be installed, for example, Gradle without lockfile, Scala</td></tr><tr><td>API</td><td><ul><li>Typically for advanced use cases.</li><li>Integration into CI/CD workflows or custom tooling. </li></ul></td><td>Automated integration into CI/CD pipelines</td><td>Requires API familiarity, access through the  Enterprise plan.</td></tr><tr><td>Git code repository integration</td><td>Used for onboarding and daily monitoring: rapid vulnerability assessment across application portfolio</td><td><p></p><ul><li>Continuous monitoring of repositories, even when you are not working on it</li><li>Centralized visibility for teams</li><li>Monitors specified branch</li><li>Code does not need to be built</li></ul></td><td><ul><li>Can be initiated through the UI, however larger portfolios should use API to initiate onboarding of repositories with the <a href="https://docs.snyk.io/snyk-api/other-tools/tool-snyk-api-import">Api Import Tool</a></li><li>Some languages/package managers have better resolution through use of  the CLI: Gradle without lockfile, Scala</li></ul></td></tr><tr><td></td><td><ul><li>Pull request (PR)/merge request (MR)  scanning</li></ul></td><td><ul><li>Immediate feedback on introduced issues on the PR/MR against any branch on repository</li></ul></td><td>Configurable rules for pass/fail</td></tr></tbody></table>

## Additional considerations for integrations

### Infrastructure as Code integration

For Snyk Infrastructure as Code, it is common that your Terraform or YAML configuration files are held in your SCM, but they may be in a separate area or repository. Thus, consider whether there are other areas you need to import. You may also want to integrate with Terraform Cloud if applicable, to enable Snyk tests as part of your `terraform run` processes.

For complex environments, modules, and highly templated implementations, using the CLI on your Terraform Plan file may provide the best results.

### Container registry (CR) integrations

Snyk also integrates with various [container registries](../../../scan-with-snyk/snyk-container/container-registry-integrations/) to enable you to import and monitor your containers for vulnerabilities. Snyk tests the containers you have imported for any known security vulnerabilities found, at a frequency you control.

### Kubernetes

Snyk can be configured to monitor workloads deployed to Kubernetes. See [Overview of Kubernetes integration](../../../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/) for more information on how to configure the controller.
