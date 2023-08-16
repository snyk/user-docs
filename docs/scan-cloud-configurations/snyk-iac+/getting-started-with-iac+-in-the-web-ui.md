# Getting started with IaC+ in the Web UI

{% hint style="info" %}
This page describes a process using the Snyk Web UI. For details on using Integrated IaC with the Snyk CLI, see [Test your Integrated IaC files with Snyk CLI](test-your-iac+-files-with-the-snyk-cli.md).
{% endhint %}

{% hint style="warning" %}
**Feature availability**\
This feature is in **closed beta**, and requires allowing Snyk to clone an entire Git repository for security analysis. To enable this feature, you must opt-in to this feature in writing with your account team via email or Slack.
{% endhint %}

Use [Snyk Integrated Infrastructure as Code](broken-reference) to inspect, find, and fix issues in cloud configuration files for Terraform and AWS CloudFormation and Azure Resource Manager (ARM) in your Git repositories. Support for Kubernetes is coming soon.

You can test your IaC files in Git repositories found via SCM integrations with [Integrated IaC](broken-reference), much like you would with the current [IaC](../snyk-infrastructure-as-code/). There are some differences, which are summarized in the following table.

<table data-header-hidden><thead><tr><th width="271"></th><th width="261.3333333333333"></th><th></th></tr></thead><tbody><tr><td></td><td><strong>Current IaC support</strong></td><td><strong>Integrated IaC support</strong></td></tr><tr><td><strong>Terraform (single file)</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Terraform (modules)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>Terraform (variables)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>CloudFormation</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Azure Resource Manager</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Kubernetes manifests</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Helm charts</strong></td><td>Yes</td><td>Coming soon</td></tr></tbody></table>

## Prerequisites for Integrated IaC

To start using Integrated IaC, you need the following:

* A Snyk account; go to [https://snyk.io/](https://snyk.io) and sign up. See [Create a Snyk account](../../getting-started/quickstart/create-a-snyk-account/) for details.
* An existing Terraform, CloudFormation, or Azure Resource Manager environment to work in.
* Integration with your Git repository as for other Snyk products; see [Git repository (SCM) integrations](../../integrations/git-repository-scm-integrations/) for more details.

## Stage 1: Import Projects

{% hint style="warning" %}
If you want to add a new Integrated IaC Project from an SCM repository that you have already imported, you must re-import the repository. This will not affect any of your existing Projects
{% endhint %}

Import [Projects](../../manage-issues/snyk-projects/) to test with Snyk by choosing repositories for Snyk to test and monitor.

1. Select **Projects** from the Snyk Web UI.
2. In the **Add projects** drop-down menu, select the SCM to add the Project from, for example, GitHub.
3. In **Personal and Organization repositories**, select the repositories to use.
4. Click **Add selected repositories** to import the selected repositories into your Projects.\
   Project import completes.

## Stage 2: View Integrated IaC Projects

On the [Projects](../../manage-issues/snyk-projects/) page, navigate to the appropriate target (Git repository) that contains IaC files for Snyk to test. You will see a single **Infrastructure as Code issues** Project. Snyk Integrated IaC generates only one Project per repository, unlike the current [IaC](../snyk-infrastructure-as-code/), which generates one Project per IaC file.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-07 at 3.57.30 PM.png" alt="Integrated IaC project for your SCM Git repository"><figcaption><p>Integrated IaC project for your SCM Git repository</p></figcaption></figure>

## Stage 3: View Integrated IaC issues in the Cloud Issues UI

Clicking on the **Infrastructure as Code Issues** link opens a filtered view of the Cloud Issues UI, to include only issues from the corresponding Integrated IaC [environment](key-concepts-in-iac+.md#environments) that aligns with your Project.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-07 at 4.04.13 PM.png" alt="Cloud Issues UI, filtered to the specific environment for your SCM Git repository"><figcaption><p>Cloud Issues UI, filtered to the specific environment for your SCM Git repository</p></figcaption></figure>

Expanding the grouped issue and selecting a given issue opens an Issue Card, that includes information on:

* The **resource** - including the location, cloud platform (such as aws) with a link to the SCM file in question for faster fixes, as well as the input type (such as `tf_hcl` for Terraform HCL).
* The **environment** - providing details on the Integrated IaC environment that corresponds to your Git repository.
* The **rule** that failed - including a link to Snyk's [security rules](https://security.snyk.io/rules/cloud/) documentation for additional information, such as specific remediation steps.
* Why your developer should fix this particular misconfiguration.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-07 at 4.09.40 PM.png" alt="Integrated IaC issue card"><figcaption><p>Integrated IaC issue card</p></figcaption></figure>
