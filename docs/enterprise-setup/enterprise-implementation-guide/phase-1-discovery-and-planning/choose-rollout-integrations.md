# Choose rollout integrations

### **SDLC Integration points**

Snyk offers many integrations that seamlessly integrate into every stage of SDLC.&#x20;

Many businesses typically roll out automated solutions first, then slowly introduce tools to enable the developers. In addition, gating features are gradually turned on over a period of time, to minimize disruption.

{% hint style="info" %}
As using multiple integrations can result in duplicate reporting of issues, you do not initially need to implement more than one integration type. For example, you can start by importing everything with Git repositories, then later use the CI/CD view for fine-grained detail (potentially removing the source control integration if both views are not desired).
{% endhint %}

## Integration types

Here are typical early integrations

### Source Code Management (SCM) Integrations

Integrations with popular version control platforms like GitHub, GitLab, Azure Repos, and Bitbucket seamlessly integrate Snyk's security checks into the code review process. This ensures that potential vulnerabilities are identified and addressed before code is merged into the main branch. Important features include,&#x20;

* Daily testing/monitoring of a specified branch (typically 'development' branch),&#x20;
* (optional) Pull Request/Merge Request checks against any branch of the repository.
* (optional) Automated dependency upgrades, automated security fix upgrades via pullrequest

Advantages:

* Visibility into repository security posture
* Automatic Scan on code change
* Immediate feedback on issues for the developer
* Onboarding of repositories can be configured via UI or [API/API Import Tool](https://docs.snyk.io/snyk-api-info/other-tools/tool-snyk-api-import)
* Supports Cloud and Private Code Repositories on the Snyk Enterprise plan

See [Git repositories (SCMs)](../../../integrations/git-repository-scm-integrations/) for more details.

If you have an on-premise git repository, you will need to consider deploying a [Snyk Broker](https://docs.snyk.io/snyk-admin/snyk-broker) for Snyk to communicate with your repositories.

{% hint style="info" %}
Enterprise customers can enable and manage [Snyk Broker](https://docs.snyk.io/enterprise-setup/snyk-broker) via API.&#x20;

[Paid services](https://docs.snyk.io/more-info/snyk-terms-of-support-and-services-glossary) can be engaged to assist in broker deployments.
{% endhint %}

### Continuous Integration/Continuous Deployment (CI/CD) Pipeline Integrations

Integrating Snyk into CI/CD pipelines, such as Jenkins, Travis CI, or CircleCI, automates security checks during the build and deployment process. This ensures that vulnerabilities are detected early in the software development lifecycle and prevents their propagation into production. Typical features include

* (Optional) Ability to passively monitor results during build and view results in Snyk
* (Optional) Ability to test and potentially break the build if results found based on criteria you specified

Advantages:

* Assess local code vulnerabilities
* Full control over testing (which tests to run, where in the builld script)
* Can automate via CI/CD

See [Snyk CI/CD integrations](../../../integrations/snyk-ci-cd-integrations/) for more details.

### IDE Integrations

Integrated Development Environment (IDE) integrations like Visual Studio Code, IntelliJ IDEA, and Eclipse allow developers to access Snyk's security features directly within their coding environment. This enables real-time scanning and issue remediation as developers write code.&#x20;

See [Use Snyk in your IDE](../../../integrations/ide-tools/) for more details.

## Considerations for import strategies&#x20;

<table><thead><tr><th width="200">Project Import Strategy</th><th>Considerations</th><th>Advantages</th><th>Disadvantages</th></tr></thead><tbody><tr><td>CLI (automated CI/CD)</td><td>Has to be configured for each application within CI/CD</td><td><ul><li>Can select what to test and when (i.e. which package managers, where in the process, which language to analyze.</li><li>May need development effort for integration</li></ul></td><td><ul><li>It requires configuration per application.</li></ul></td></tr><tr><td>CLI (Run locally by user)</td><td>User can use CLI to perform testing locally while working on an application, very configurable per scan type.</td><td><ul><li>Local use case</li></ul></td><td><ul><li>Not meant for visibility or automation. Can require buildable code or dependencies to be installed (For example Gradle without Lockfile, Scala)</li></ul></td></tr><tr><td>API</td><td><ul><li>Typically for advanced use cases.</li><li>Integration into CI/CD workflows or custom tooling. </li></ul></td><td><ul><li>Automated integration into CI/CD pipelines</li></ul></td><td><ul><li>Requires API familiarity, access and Enterprise plan.</li></ul></td></tr><tr><td>Git Code Repository Integration</td><td><ul><li>Onboarding and daily monitoring: rapid vulnerability assessment across application portfolio</li></ul></td><td><p></p><ul><li>Continuous monitoring of repositories (Even when you are not working on it)</li><li> Centralized visibility for teams</li><li>Monitors specified branch</li><li>Code doesn't need to be built</li></ul></td><td><ul><li>Can be initiated via UI, however larger portfolios should use API to initiate onboarding of repositories using the <a href="https://docs.snyk.io/snyk-api/other-tools/tool-snyk-api-import">Api Import Tool</a></li><li>Some languages/package managers have better resolution utilizing the CLI (Gradle without lockfile, Scala)</li></ul></td></tr><tr><td></td><td><ul><li>Pull request (PR)/merge request (MR)  scanning</li></ul></td><td><ul><li>Immediate feedback on introduced issues on the PR/MR against any branch on repository</li></ul></td><td>Configurable rules for pass/fail</td></tr></tbody></table>

## Additional Considerations

### Infrastructure as Code

For Snyk Infrastructure as code, it is common that your Terraform or yaml configuration files are held in your SCM, but they may be in a separate area or repository, so it’s worth considering if there are other areas you need to import. You may also want to integrate with Terraform Cloud (if applicable) to enable Snyk tests as part of your ‘terraform run’ processes.

For complex environments, modules, and highly templated implementations, utilizing the CLI on your Terraform Plan file may provide the best results.

### CR (Container Registries)

Snyk also integrates with various [Container Registries](../../../scan-containers/snyk-container-integrations/) to enable you to import and monitor your containers for vulnerabilities. Snyk tests the containers you have imported for any known security vulnerabilities found at a frequency you control.

### Kubernetes

Snyk can be configured to monitor workloads deployed to Kubernetes. See [Overview of the Kubernetes integration](https://docs.snyk.io/scan-applications/snyk-container/kubernetes-integration/overview-of-the-kubernetes-integration) for more information on how to configure the controller.
