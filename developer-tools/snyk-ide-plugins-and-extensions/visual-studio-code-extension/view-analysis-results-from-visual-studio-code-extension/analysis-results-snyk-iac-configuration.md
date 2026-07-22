---
description: How to view Snyk IaC configuration results for Terraform, Kubernetes, and CloudFormation in VS Code
---

# Analysis results: Snyk IaC configuration

Snyk IaC configuration analysis shows issues in your Terraform, Kubernetes, AWS CloudFormation, and Azure Resource Manager (ARM) code with every scan. Based on the Snyk CLI, the scan is fast and friendly for local development. The scan runs in the background and is enabled by default.

## Snyk IaC configuration issues window

The configuration issues window shows information about issues. By clicking on an issue, you can learn more about it:

<figure><img src="../../../.gitbook/assets/snyk-iac-configuration-issues-window.png" alt="Snyk IaC configuration issues window"><figcaption><p>Snyk IaC configuration issues window</p></figcaption></figure>

The following information is shown:

* Issue description
* Issue impact
* Issue path
* Remediation details
* Links to references

In the **Problems** tab of the Visual Studio Code configuration issues screen, you can see all configuration issues found in your Project.

<figure><img src="../../../.gitbook/assets/problems-tab.png" alt="Problems tab" width="563"><figcaption><p>Problems tab</p></figcaption></figure>

## Snyk IaC configuration editor window

The issues are visible within the editor, with the detailed information available on hover.

<figure><img src="../../../.gitbook/assets/snyk-iac-configuration-issue.png" alt="Snyk IaC configuration issue" width="563"><figcaption><p>Snyk IaC configuration issue</p></figcaption></figure>

Choose **Quick Fix** to open the details panel for an issue using Code Action.

<figure><img src="../../../.gitbook/assets/quick-fix.png" alt="Quick Fix" width="563"><figcaption><p>Quick Fix</p></figcaption></figure>

The details panel shows the issue name with the **Description**, **Impact** statement, **Path** by which the issue was introduced, and suggested **Remediation**.

<figure><img src="../../../.gitbook/assets/details-panel-snyk-iac-configuration-issue.png" alt="Details panel for a Snyk IaC configuration issue" width="375"><figcaption><p>Details panel for a Snyk IaC configuration issue</p></figcaption></figure>
