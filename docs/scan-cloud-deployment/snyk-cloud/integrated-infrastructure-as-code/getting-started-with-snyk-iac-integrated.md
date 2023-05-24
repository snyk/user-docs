# Getting started with Snyk IaC in the Web UI (Integrated IaC)

{% hint style="info" %}
This article describes a process using the Snyk Web UI. For details of using Integrated IaC with the Snyk CLI, see [Test your IaC files with Snyk CLI](test-your-iac-files-with-snyk-cli-integrated-iac.md).
{% endhint %}

{% hint style="warning" %}
**Feature availability**\
This feature is in **closed beta**, and requires allowing Snyk to clone an entire git repository for security analysis. To enable this feature, you will need to opt-in to this feature in writing with your account team via email or Slack.
{% endhint %}

Use [Snyk Integrated Infrastructure as Code](./) to inspect, find, and fix issues in cloud configuration files for Terraform and AWS CloudFormation and Azure Resource Manager (ARM) in your Git repositories, with support for Kubernetes coming soon.

You can test your IaC files in Git repositories found via SCM integrations with [Integrated IaC](./), much like you would with [Current IaC](../../snyk-infrastructure-as-code/). There are some differences, which are summarized in the following table.

<table data-header-hidden><thead><tr><th width="271"></th><th width="261.3333333333333"></th><th></th></tr></thead><tbody><tr><td></td><td><strong>Current IaC support</strong></td><td><strong>Integrated IaC support</strong></td></tr><tr><td><strong>Terraform (single file)</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Terraform (modules)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>Terraform (variables)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>CloudFormation</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Azure Resource Manager</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Kubernetes manifests</strong></td><td>Yes</td><td>Coming soon</td></tr><tr><td><strong>Helm charts</strong></td><td>Yes</td><td>Coming soon</td></tr></tbody></table>

## Prerequisites

To start using Integrated IaC, you need the following:

* A Snyk account (go to [https://snyk.io/](https://snyk.io) and sign up - see [Create a Snyk account](https://docs.snyk.io/getting-started/getting-started-snyk-products) for details).
* An existing Terraform, CloudFormation, or Azure Resource Manager environment to work in.
* Integration with your Git repository as for other Snyk products; see [Git repository (SCM) integrations](../../../integrations/git-repository-scm-integrations/) for more details.

## Stage 1: Import projects

{% hint style="warning" %}
If you want to add a new Integrated IaC project from an SCM repository that you've already imported, you will need to re-import the repository. This will not affect any existing projects that you have.
{% endhint %}

Import [Projects](../../../manage-issues/introduction-to-snyk-projects/#project) to test with Snyk by choosing repositories for Snyk to test and monitor.

1. Select **Projects** from the Snyk Web UI.
2. In the **Add projects** drop-down menu, select the tool to add the Project from (for example, GitHub).
3. In **Personal and Organization repositories**, select the repositories to use.
4. Click **Add selected repositories** to import the selected repositories into your projects.
5. Project import completes.

## Stage 2: View Integrated IaC projects

On the [Projects](../../../manage-issues/introduction-to-snyk-projects/) page, navigate to the appropriate target (git repository) that contains IaC files for Snyk to test. You will see a single "Infrastructure as Code issues" project - Snyk Integrated IaC will only generate one project per repository, unlike [Current IaC](../../snyk-infrastructure-as-code/), which will generate one project per IaC file.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-07 at 3.57.30 PM.png" alt=""><figcaption><p>Integrated IaC project for your SCM Git repository</p></figcaption></figure>

## Stage 3: View Integrated IaC issues on the Cloud Issues UI

Clicking on the "Infrastructure as Code Issues" link will take you to a filtered view of the Cloud Issues UI, to include only issues from the corresponding Integrated IaC [Environment](../snyk-cloud-concepts.md#environments) that aligns with your Project.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-07 at 4.04.13 PM.png" alt=""><figcaption><p>Cloud Issues UI, filtered to the specific environment for your SCM Git repository</p></figcaption></figure>

Expanding the grouped issues, and selecting a given issue will open an Issue Card, that includes information on:

* The **resource** - including the location, cloud platform (such as "aws") with a link to the SCM file in question for faster fixes, as well as the input type (such as "tf\_hcl" for Terraform HCL).
* The **environment** - providing details on the Integrated IaC environment that corresponds to your Git repository.
* The **rule** that failed - including a link to Snyk's [security rules](https://snyk.io/security-rules/cloud/) documentation for additional information, such as specific remediation steps.
* Why your developer should fix this particular misconfiguration.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-07 at 4.09.40 PM.png" alt=""><figcaption><p>Integrated IaC issue card</p></figcaption></figure>
