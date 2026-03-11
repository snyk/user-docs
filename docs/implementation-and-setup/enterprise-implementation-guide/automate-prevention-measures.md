# Automate prevention measures

Once you have visibility into your existing security posture, implement prevention and gating systems to stop new vulnerabilities from entering your applications. By automating these checks, you empower developers to take responsibility for the security of their specific changes without manually triaging every issue.

Implementing prevention measures involves the following key stages:

1. [Define prevention methods:](automate-prevention-measures.md#define-prevention-methods) Choose between Pull Request (PR) checks and CI/CD pipeline gating.
2. [Establish exception processes:](automate-prevention-measures.md#establish-exception-processes) Create clear workflows for handling blocked builds or PRs.
3. [Configure PR/MR checks:](automate-prevention-measures.md#configure-pr-mr-checks) Set up automated scanning for code changes in your source control manager (SCM).
4. [Integrate with CI/CD pipelines:](automate-prevention-measures.md#integrate-with-ci-cd-pipelines) Add Snyk tests to your build process as a final gate.
5. [Secure custom images and IaC:](automate-prevention-measures.md#secure-custom-images-and-iac) Apply prevention to containers and infrastructure as code.
6. [Announce changes:](automate-prevention-measures.md#announce-prevention-measures) Use templates to inform developers about upcoming gating.

## Define prevention methods

{% hint style="success" %}
**Key decision**: Choose a primary prevention point. Use PR checks for early developer feedback or CI/CD gating for a final check before deployment.
{% endhint %}

You can prevent new issues using two main functions:

* **PR/MR checks**: Available for Snyk Open Source and Snyk Code. These test code changes are immediately upon submission.
* **CI/CD pipelines**: Integrate Snyk into your build pipeline to gate Open Source, Code, IaC, and Container vulnerabilities.

## Establish exception processes

{% hint style="success" %}
**Key decision**: Define who can override security gates. Clear authority prevents development bottlenecks during urgent releases.:&#x20;
{% endhint %}

Ensure teams understand how to address blocked PRs or failed builds:

* Identify who can override a PR check if a pass is mandatory.
* Determine if an issue can be ignored or if a script can bypass a specific step.

## Configure PR/MR checks

{% hint style="success" %}
**Key decision**: Use a phased rollout. Start with optional checks that show results, then move to blocking checks once developers are familiar with the workflow.
{% endhint %}

PR checks prevent issues from entering the codebase.

* Configure tests to fail only under specific criteria, such as High or Critical severity issues.
* For Snyk Open Source, you can set tests to fail only when a fix is available.

## Integrate with CI/CD pipelines

{% hint style="success" %}
**Key decision**: Choose between native integrations or the Snyk CLI. The CLI offers more flexibility for complex fail criteria, such as using `snyk-filter` to set specific thresholds.
{% endhint %}

Adding Snyk to your pipeline acts as a gatekeeper:

* **No import is required**: Unlike PR checks, pipeline tests do not require repositories to be imported via SCM integration.
* **Use CLI tools**: Use `snyk-delta` to identify only the new vulnerabilities introduced in a specific build.

## Secure custom images and IaC

{% hint style="success" %}
**Key decision**: Move security upstream. Test base images before developers use them to ensure all derived containers start from a secure foundation.
{% endhint %}

* **Container registry**: Run Snyk container tests when creating custom base images. Snyk Container (**Detect Dockerfiles**) is enabled for Organizations by default. To disable it, navigate to the Dockerfile tile under your Org-level SCM integration **Settings**.
* **IaC**: Integrate with workflows like Terraform Cloud to scan configuration files before deployment. IaC is enabled for Organizations by default. To disable it, navigate to Organization **Settings** > **Snyk IaC**.

## Announce prevention measures

{% hint style="success" %}
**Key decision**: Move security upstream. Test base images before developers use them to ensure all derived containers start from a secure foundation.
{% endhint %}

### Email template

|                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p><em>To: Developers</em></p><p><em>Subject: Introducing Snyk tests to PRs [Company name]</em></p><p><em>Hi all,</em></p><p><em>As part of our ongoing aim to improve our application security at [Company name], we are preparing to start running Snyk tests against all new pull requests for any repository that has been imported into Snyk.</em></p><p><em>[optional: add personalized video, if desired]</em></p><p><em>These checks will identify any new High or Critical severity issues that are part of the PR, with the aim of preventing any new significant issues from entering our repositories. At first, these checks will be optional, meaning you are not blocked from merging a PR if one of these vulnerabilities is detected.</em></p><p><em>In the future, this will be changing to a blocking check, so we would recommend you start remediating any new High or Critical issues that are detected in your PRs, so that you aren’t affected when the test is no longer optional.</em></p><p><em>This change will make a huge difference in improving our application security, and by gradually introducing this feature, we hope to avoid any interruptions to your workflow.</em></p><p><em>More info can be found at [hyperlink to your internal resource page/wiki with more info].</em></p><p><em>Regards,</em></p><p><em>_____ [Sender]</em></p> |

### Instant message template

|                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| _Snyk tests being introduced to our PRs: From \[date] we’ll be enabling a feature in Snyk so that all new PRs on repositories that have been imported to Snyk will be tested for new vulnerabilities. You’ll see the test will fail if any new High or Critical severity issues are found. Please fix these before merging if possible! For now, the tests are optional, so you can merge the PR even if the test fails, but in the future, we’ll be setting this to be a required check. Get in touch if you have any questions!_ |
