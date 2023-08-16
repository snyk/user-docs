# Snyk for C/C++ Developers

## Introduction

Use this guide to understand the best way to apply Snyk in your workflow and to be aware of key considerations for your chosen technology stack.&#x20;

This guide focuses on scanning your application code, which is specific to your open source usage and programming language. [Snyk Container](../scan-containers/) and [Snyk Infrastructure as Code](../scan-cloud-configurations/snyk-infrastructure-as-code/) (IaC) also support your container and Infrastructure as Code needs.&#x20;

{% hint style="info" %}
[Snyk Code](../scan-application-code/snyk-code/) for C/C++ is currently in Beta, it can be enabled using **Settings > Snyk Preview.**
{% endhint %}

### Developer-first approach

Snyk takes a developer-first approach to securing your application. This approach focuses less on overhead-heavy work (such as putting in hard QA gates), and more on building a secure application, providing visibility in a developer’s workflow, and providing actionable insights. The benefits of this approach are to engage developers in security practices as part of their development, without seeing security as a costly overhead.

So we recommend you use Snyk to focus on earlier enablement, not later enforcement. Do not wait until you commit the code to find an issue requiring changes. Snyk provides tools that allow you to test while working on a project to minimize rework. Add a package: test it before writing the code that interfaces with it. Similarly, after writing major sections of code, test it before moving on.

### Snyk overview

Before we get started, let's introduce Snyk, as some features mentioned in this document may not be available depending on your Snyk plan or product. Each Snyk product provides key capabilities for the ecosystems you are working in. Snyk [Pricing plans](https://snyk.io/plans) determine what features are available.

{% hint style="info" %}
Note that there is a monthly limit to the number of tests performed if a particular product is not purchased. You can find the limits outlined [here](https://snyk.io/plans/) in the first square.
{% endhint %}

#### Snyk products

* [Snyk Code](../scan-application-code/snyk-code/): scan your own code for security vulnerabilities using source code analysis.
* [Snyk Open Source](../scan-application-code/snyk-open-source/)\*
  * Open Source vulnerability testing and monitoring (All plans)
  * License Compliance (paid plans)&#x20;
* [Snyk Container](../scan-containers/)
  * Scan for issues with container images if you are building containers
* [Snyk Infrastructure as Code](../scan-cloud-configurations/snyk-infrastructure-as-code/)
  * Scan for configuration issues when you deploy your applications using AWS Cloudformation templates, Kubernetes deployment files, Terraform, or Azure Resource Manager.
* [Snyk Cloud](../scan-cloud-configurations/snyk-iac+/): Security from code to cloud and back
  * Scan for runtime misconfiguration issues in your cloud and containers, detect infrastructure drift, and fix issues at their source.

{% hint style="info" %}
\*Some capabilities may be limited for some languages and package managers
{% endhint %}

## Product features to enhance your workflow

As you start planning and your code progresses through the pipeline, Snyk can provide different capabilities at each stage, to help you find and fix security issues.

### Designing and planning

The following public resources are available for all users:

* [Snyk Learn](https://learn.snyk.io/): Assists you in learning to code securely.
* [Snyk Training](https://training.snyk.io/): Provides training on how to use Snyk.

### Coding

{% hint style="info" %}
These capabilities focus on enablement, not enforcement. Use Snyk tools to test while working on a project, rather than testing after you commit the code and then discovering issues requiring code rework. Add a package and test it, before writing the code that interfaces with it. Similarly, after writing major sections of code, test the change before you continue.
{% endhint %}

The following capabilities are available for all Snyk users.

* [Snyk CLI](../snyk-cli/):  A powerful terminal program that allows you to test locally on your machine. Very useful in testing containers and more complex IaC files that are templated with variables (such as Terraform plan files), as well as scanning open source and your own code.
* [IDE Plugins](../integrations/ide-tools/): for VS Code, Visual Studio and others: Test your open source packages and first party code as you develop. Additionally test infrastructure as code (IaC) Kubernetes deployment files you create.

### Validating, monitoring, alerting and gating

The following is available for all Snyk users.

**With Git integrations**

* Snyk can monitor for vulnerabilities in source code directly via Git integration.
  * Pull request checks ([PR Checks](../scan-application-code/run-pr-checks/)) for Snyk Code is currently in Closed Beta.

{% hint style="info" %}
For the Git integration, Open Source for C/C++ is only supported via CLI for automation purposes, as dependencies are typically managed outside of Git in the C/C++ ecosystem.&#x20;
{% endhint %}

#### With CI/CD integrations&#x20;

Snyk can passively monitor using **snyk monitor**, and provide a QA gate by failing build checks during testing for policy violations using **snyk test** command. See [Snyk test and snyk monitor in CI/CD integration](../integrations/ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md).

Snyk provides flexible capabilities, including:

* CLI can be used in most CI/CD systems ([examples](https://github.com/snyk-labs/snyk-cicd-integration-examples))
  * Remember to add the **--unmanaged** commands for Snyk Open Source.
  * Fail the build based on criteria using options or the [snyk-filter](../snyk-api-info/other-tools/tool-snyk-filter.md) tool
  * There are [containerized](https://hub.docker.com/r/snyk/snyk) versions available
* Partner Platforms - Azure, Bitbucket, and AWS have built-in pipes/components for use with Snyk or use the CLI.
* Using [Github Actions](https://snyk.io/blog/building-a-secure-pipeline-with-github-actions/)

#### Production monitoring

* Where a production integration does not exist, use the [snyk monitor](../snyk-cli/commands/monitor.md) CLI command to take a snapshot and monitor what Open Source is being pushed to production.&#x20;

## Language/Open Source - specific facts for using Snyk

### C/C++ with Snyk Code

{% hint style="info" %}
**Snyk Code for C/C++ is currently in Beta.** \
The initial Beta release focuses mostly on desktop and server applications that run on terminal on Linux. There are rules planned to span application types. \
\
The rules expansion, GUI framework support and much more are planned for the GA release.
{% endhint %}

* Snyk does not compile or require a build to perform analysis.
* Snyk Code analyzes source code directly
* If you precompile components, make the source available during the scan

### C/C++ with Snyk Open Source

Snyk Open Source can be used to test open-source packages for known vulnerabilities, license usage and license compliance issues. Snyk can also alert to new vulnerabilities discovered since the last snapshot.

In the case of package managers like npm or maven, it is traditionally using the managed open source capabilities of "snyk test" and "snyk monitor". In the case of C/C++, Snyk supports unmanaged dependencies by adding **--unmanaged**.

{% hint style="info" %}
Snyk does not hook into a build, nor rely on a build to perform scanning. Snyk performs analysis from source code.
{% endhint %}

1. Open Source source code must be present
2. Snyk will fingerprint files and compare to the Snyk database to identify packages, versions, licenses, and vulnerabilities.

## Snyk Integrations and common usage patterns

### IDE

#### With Snyk Code

No additional options required. The Snyk plugin has views within the IDE for displaying results.

**With Snyk Open Source**&#x20;

Under **Additional Parameters** in the IDE settings, enter the **--unmanaged** option to scan for C/C++ open source dependencies.

<div align="left">

<figure><img src="https://lh6.googleusercontent.com/1j-2sJjuVejBJ6nARpaAx2uhdhqT7G3XyNCGZqFxBXJV9ujqRHBYiwInr_mFT7SH-fnhG6iUysKxzYKluPG1f3xUKyb2q-JycA_0QevtaS3hdm4I7-QT7M5benqzWkIe5N-7L3czV-F84_xUR5yl7k0" alt="Scan for dependencies"><figcaption><p>Scan for dependencies</p></figcaption></figure>

</div>

### CLI Tips and Tricks

#### With Snyk Code

Snyk does not rely on a build; only source code is required:

```
snyk code test
```

{% hint style="info" %}
If you precompile components, the source code should still be present to get the best resuts and coverage.
{% endhint %}

For reporting, you can generate reports using the [snyk-to-html](../scan-application-code/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature/) plugin to generate reporting artifacts. Additionally, there are JSON and SARIF export capabilities for programmatic access to results, using **--json** and **--sarif**, respectively. See [Exporting the test results to a JSON or SARIF file](../scan-application-code/snyk-code/cli-for-snyk-code/working-with-the-snyk-code-cli-results/exporting-the-test-results-to-a-json-or-sarif-file.md).

#### **With Snyk Open Source**

For C/C++ open source, use the **--unmanaged** option to analyze license compliance issues and known security issues associated with open source. See [Snyk for C/C++ (Open Source) ](../scan-application-code/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-c-c++.md)for details.

* To test, make sure the open-source source code is present, it may be placed in a vendored folder. This should not be an issue.
* If you precompile open source, the open source code must still be present.

```
snyk test --unmanaged
```

Similarly, for monitoring and sharing reporting:

```
snyk monitor --unmanaged --org=<org-id>
```

Where **org-id** is found under your org settings in the Snyk web interface. Although the Organization id is not required, it's strongly suggested. Similar to Snyk Code, you can generate reports using the [snyk-to-html](../scan-application-code/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature/) plugin to generate reporting artifacts.&#x20;

* For individual or personal scans, use the CLI or IDE and use the **snyk monitor --unmanaged** command to upload results, but we recommend you send these results to your personal folder and disable the scheduled scanning in the Project settings to ensure an individual scan does not cause noise. This will provide license/policy information in a viewable state.
* For automated scans, such as CI/CD, use **snyk monitor --unmanaged** and send results to the organization of your choice. This will provide license/policy information in a viewable state.

**Dependency lists**

A very useful feature. **--print-deps**, exists when performing open source scans to determine how much coverage, or what is detected, during a scan.&#x20;

In C/C++ this has the additional benefit of identifying the confidence level of a given match. If there is a significant drop (< 90% confidence) it's likely the file has been modified and may not be the original source. You should consider investigating.

```
snyk test --unmanaged --print-deps
```

The list will be printed before the issues list, as shown below:

<figure><img src="https://lh5.googleusercontent.com/x4y1uIQ2fCFX956f1eP4664i6VKEgK6eOOddlAZ4p4WnQWJu1t_ugSOpL394KEnuzSIPRs08gNAsmjvPa-GAV0C-975esRdy0EPDY7WImG1-SXSOFO0TIAVfh_Jp2DLYc6bm7iZu55UbE3Boh4TNk_I" alt="View dependency lists"><figcaption><p>View dependency lists</p></figcaption></figure>

**License policy text during the Beta phase**

{% hint style="info" %}
[License Compliance](../scan-application-code/snyk-open-source/licenses/) allows a company to create a license policy for your Open Source indicating what licenses are not approved for use. To get access to [License Compliance](../scan-application-code/snyk-open-source/licenses/), you must be on a Snyk Team or Enterprise [plan](https://snyk.io/plans). Snyk will detect and alert when a match is found. The alert contains the name of the license and license policy text.&#x20;

**License policy text** is the text associated to the issue by your administrators that provides custom direction on what to do, or why it's contrary to the policy, if it's found in your application.
{% endhint %}

For the Beta, while the issues are displayed, the license policy text associated to the issues are not currently displayed within the IDE or CLI. This is planned to be addressed by GA. To see the policy text, use the Snyk Web UI.&#x20;

Below, you can see the license policy text example at the bottom of the screen giving the developer direction on what to do if the license is found.

<div align="left">

<figure><img src="https://lh4.googleusercontent.com/lIn5JFEyaZaTNMVenBoeGIgTpC6YHxpmAjK947z5ISPlHV1rlOvPNCLyzXxsGNj65AAlGn6ff9dF4lHVsVFYMaKXWC939tasD91k98xcDv_Ske6Dz7goMXl5lByyqg6ptvvqaK0UEqLSdzUU9GKrW4U" alt=""><figcaption><p>License policy text example</p></figcaption></figure>

</div>

**Alternate Testing Options**

Sometimes customers develop advanced dependency management strategies and may not necessarily use the standard/well-traveled package managers. For that reason Snyk has provided test apis. In the case of C++, if you know the OS packages/versions that are included in the application but don't have the source code, the purl API can be used to do this analysis.

* PURL See [Purl issues](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org\_id-/packages/-purl-/issues)&#x20;

#### Helpful Options/Plugins

* See [snyk-to-html](../scan-application-code/snyk-code/cli-for-snyk-code/displaying-the-cli-results-in-an-html-format-using-the-snyk-to-html-feature/) plugin to help generate reports locally or at build time
* See **--json** and **--sarif** options for generating output that can be programmatically accessed
* See [snyk-filter](../snyk-api-info/other-tools/tool-snyk-filter.md) for advanced filtering options and [other tools](../snyk-api-info/other-tools/)&#x20;

## Deployment and rollout recommendations

Startups, small teams, individuals, and open-source maintainers can typically onboard their applications in minutes and start to address issues almost immediately. Small teams have the benefit of being nimble and determining what works best for their workflow.

With large businesses, and potentially hundreds of applications, a slower approach is recommended to get developer buy-in/adoption and to ensure a positive experience. Passive monitoring can be performed to start, with perhaps one or two key projects having gating enabled early on to familiarize everyone with the process. Gating using blocking builds on a wider scope typically starts after the first 30 days.

If you onboard a large number of legacy applications, we suggest you use [priority score](../manage-issues/prioritizing-and-managing-issues/priority-score.md) (typically 700 as a starting place) or criteria such as “known exploit” or “fix available”, to define a starting point to quickly engage developers to start fixing vulnerabilities on key applications.

## How to fix issues using Snyk

This discussion has two sides: what to fix and how to fix it.

#### What to fix - prioritization factors

Some tools only use the single factor of severity to prioritize issues, but this can still result in thousands of results, with no clear starting point to fix these issues.

Snyk provides more factors to help you prioritize issues, such as having a known exploit, is it fixable, social trending, and others. This can be done both at the project level when looking at a specific project, or Enterprise customers can prioritize across all your projects.

To see prioritization in action, see the [Prioritize issues in the Snyk Web UI](https://training.snyk.io/learn/video/prioritize-ui) training course.

#### How to fix - addressing issues

Snyk offers capabilities in this ecosystem to help address issues, both reactively and proactively:

1. **Being proactive**:\
   Use the CLI and IDE plugins to test while developing.\
   Add a package, make sure it’s installed, and scan for security before writing your code.
2. **Remediation advice**:\
   Snyk supplies advice across integrations on the results screens to indicate what version an issue has been fixed in.
