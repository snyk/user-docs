---
description: How to choose which Snyk integrations to roll out across the SDLC, from automated scans to developer tools
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Choose rollout integrations

## SDLC integration points

Snyk offers many integrations that seamlessly integrate into every stage of the software development lifecycle (SDLC).

Many businesses roll out automated solutions first, then slowly introduce developer tools. They also turn on gating features gradually to minimize disruption.

{% hint style="info" %}
Using multiple integrations can result in duplicate reporting of issues, so you do not need to implement more than one integration type initially. For example, you can start by importing everything with an SCM integration, and later use the CI/CD view for fine-grained detail. If you do not want both views, you can remove the source control integration.
{% endhint %}

## Integration types

The following are typical early integrations.

### Source code management (SCM) integrations

Integrations with popular version control platforms like GitHub, GitLab, Azure Repos, and Bitbucket seamlessly integrate Snyk security checks into the code review process. This helps you identify and address potential vulnerabilities before the code is merged into the main branch. Important features include:

* Daily testing and monitoring of a specified branch (typically the development branch)
* (Optional) Pull request (PR) and merge request (MR) checks against any branch of the repository
* (Optional) Automated dependency upgrades and automated security fix upgrades using PRs

Advantages:

* Visibility into repository security posture
* Automatic scan on code change
* Immediate feedback on issues for the developer
* Repository onboarding that you can configure in the UI
* Support for cloud repositories on the Team plan

For more details, visit [SCM integrations](https://docs.snyk.io/developer-tools/integrations/scm-integrations/organization-level-integrations).

{% hint style="info" %}
If you have a non-cloud-facing or self-hosted Git SCM:

* Consider deploying [Snyk Broker](https://docs.snyk.io/platform-administration/snyk-broker/snyk-broker) so that Snyk can communicate with your repositories. Snyk Broker requires a Snyk Enterprise plan.
* Enterprise customers can enable and manage Snyk Broker using the API.

You can engage [paid services](https://docs.snyk.io/snyk-data-and-governance/snyk-terms-of-support-and-services-glossary) to help with Broker deployments.
{% endhint %}

### Continuous integration/continuous deployment (CI/CD) pipeline integrations

Integrating Snyk into CI/CD pipelines, such as Jenkins, Travis CI, or CircleCI, automates security checks during the build and deployment process. This helps detect vulnerabilities early in the SDLC and prevents them from reaching production. Typical features include:

* (Optional) Ability to passively monitor results during build and view results in Snyk
* (Optional) Ability to test and potentially break the build based on criteria you specify
* Integration through specific marketplace plugins or, more generally, through the CLI as part of your pipeline script

Advantages:

* Assessment of local code vulnerabilities
* Full control over testing (which tests to run, where in the build script)
* Automation using CI/CD

For more details, visit [Snyk CI/CD integrations](https://docs.snyk.io/developer-tools/integrations/snyk-ci-cd-integrations).

### Integrated development environment (IDE) integrations

IDE integrations for Visual Studio Code, IntelliJ IDEA, and Eclipse let developers access Snyk security features directly in their coding environment. Developers can scan and remediate issues in real time as they write code, at the earliest possible stage.

For more details, visit [Use Snyk in your IDE](https://docs.snyk.io/developer-tools/integrations/snyk-ide-plugins-and-extensions).

## Considerations for import strategies

<table><thead><tr><th width="200">Project Import Strategy</th><th>Considerations</th><th>Advantages</th><th>Disadvantages</th></tr></thead><tbody><tr><td>CLI (automated CI/CD)</td><td>Must be configured for each application within CI/CD.</td><td><ul><li>Can select what to test and when (that is, which package managers, where in the process, which language to analyze).</li><li>May need development effort for integration.</li></ul></td><td>It requires configuration per application.</td></tr><tr><td>CLI (Run locally by user)</td><td>User can use CLI to perform testing locally while working on an application, very configurable per scan type.</td><td>Local use case</td><td>Not meant for visibility or automation. Can require buildable code or dependencies to be installed (For example Gradle without Lockfile, Scala).</td></tr><tr><td>Git Code Repository Integration</td><td>Onboarding and daily monitoring: rapid vulnerability assessment across application portfolio.</td><td><ul><li>Continuous monitoring of repositories (even when you are not working on it).</li><li>Centralized visibility for teams.</li><li>Monitors specified branch</li><li>Code does not need to be built.</li></ul></td><td><ul><li>Can be initiated through the UI</li><li>Some languages/package managers have better resolution using the CLI (Gradle without lockfile, Scala).</li></ul></td></tr><tr><td></td><td>Pull request (PR)/merge request (MR) scanning</td><td>Immediate feedback on introduced issues on the PR/MR against any branch on repository.</td><td>Configurable rules for pass/fail</td></tr></tbody></table>

## Additional considerations

### Infrastructure as code

For Snyk IaC, your Terraform or YAML configuration files are commonly stored in your Git repositories, but they can be in a separate area or repository. As a result, consider whether there are other areas you need to import. If applicable, you can also integrate with Terraform Cloud to run Snyk tests as part of your Terraform run processes.

For complex environments, modules, and highly templated implementations, using the CLI on your Terraform plan file can provide the best results.

### Container registries

Snyk also integrates with various [container registries](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-container/container-registry-integrations) so that you can import and monitor your containers for vulnerabilities. Snyk tests the containers you import for known security vulnerabilities at a frequency you control.
