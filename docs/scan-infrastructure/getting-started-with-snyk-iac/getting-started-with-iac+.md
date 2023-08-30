# Getting started with IaC+

{% hint style="warning" %}
,to **Feature availability**\
IaC+ is in closed beta and requires that you allow Snyk to clone an entire Git repository, for security analysis. To use IaC+, you must choose to use the feature in writing by email or a Slack message to your account team.
{% endhint %}

{% hint style="info" %}
This page explains using IaC+ in the Snyk Web UI. For information about using IaC+ the Snyk CLI, see [Test your Integrated IaC files with Snyk CLI](../snyk-cli-for-infrastructure-as-code/test-your-iac-files-with-the-snyk-cli.md).
{% endhint %}

Use IaC+ to find, view, and fix issues in cloud configuration files for Terraform, Kubernetes (except Helm, coming soon), AWS CloudFormation, and Azure Resource Manager (ARM) in your Git repositories.

You can scan your IaC files in Git repositories that are integrated with Snyk, much as you would with the original IaC. There are some differences, which are summarized in the following table.

<table data-header-hidden><thead><tr><th width="271"></th><th width="261.3333333333333"></th><th></th></tr></thead><tbody><tr><td></td><td><strong>Original IaC support</strong></td><td><strong>IaC+ support</strong></td></tr><tr><td><strong>Terraform (single file)</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Terraform (modules)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>Terraform (variables)</strong></td><td>No</td><td>Yes</td></tr><tr><td><strong>CloudFormation</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Azure Resource Manager</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Kubernetes manifests</strong></td><td>Yes</td><td>Yes</td></tr><tr><td><strong>Helm charts</strong></td><td>Yes</td><td>Coming soon</td></tr></tbody></table>

## Prerequisites for IaC+

To start using IaC+ you must have the following:

* A Snyk account. For details, seee [Create a Snyk account](../../getting-started/quickstart/create-a-snyk-account/).
* An existing Terraform, CloudFormation, or Azure Resource Manager environment to work in.
* Integration with your Git repository as for other Snyk products. For details, see [Git repositories (SCMs)](../../integrations/git-repository-scm-integrations/).

## Import IaC+ Projects

{% hint style="warning" %}
If you want to add a new IaC+ Project from a Git repository that you have already imported, you must re-import the repository. This will not affect any of your existing Projects.
{% endhint %}

You will start by importing [Projects](../../manage-issues/snyk-projects/) you want to scan with Snyk. In these steps, you choose repositories for Snyk to test and re-test:

1. Log in to Snyki, and on your dashboard, select **Projects** from the navigation.
2. On the Projects page, from the **Add projects** dropdown, select the SCM from which to add the Projects; for example, select GitHub.
3. From the list of **Personal and Organization repositories**, select the Git repositories you want to use.
4. Click **Add selected repositories** to add the selected repositories to Snyk.\
   The import completes and the Projects page displays the Snyk Projects that have been added.

## View Integrated IaC Projects

On the [Projects](../../manage-issues/snyk-projects/) page, ensure **Group by targets** is selected and navigate to the Target (Git repository) that contains the files for IaC+ to test.

You will see a single **Infrastructure as Code issues** Project. IaC+ generates only one Project in each repository, unlike the original IaC, which generates one Project for each configuration file.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-07 at 3.57.30 PM.png" alt="Integrated IaC Project in your SCM Git repository"><figcaption><p>Integrated IaC Project in your SCM Git repository</p></figcaption></figure>

## View cloud configuration file issues in IaC+

Click on the **Infrastructure as Code Issues** Project link to open a view of the IaC+ **Issues** UI, filtered to include only issues from the IaC+ [environment](../key-concepts-in-iac+/#environments) that corresponds to your Project.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-07 at 4.04.13 PM.png" alt=".IaC+ Issues UI, filtered to show issues from the environment for your repository"><figcaption><p>IaC+ Issues UI, filtered to show issues from the environment for your repository</p></figcaption></figure>

Issues are grouped by rule. Expand the rule and select an issue to open its issue card. Each issue card has information about the following:

* The **resource**, including the location, cloud platform, such as aws, a link to the SCM file for fast fixes, and the input type, such as `tf_hcl` for Terraform HCL.
* The **environment**, providing details on the IaC+ environment that corresponds to your Project.
* The **rule** that failed, including a link to the Snyk [security rules](https://security.snyk.io/rules/cloud/) for additional information, such as specific remediation steps.
* The reason why your developers should fix this misconfiguration.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-07 at 4.09.40 PM.png" alt="IaC+ issue card"><figcaption><p>IaC+ issue card</p></figcaption></figure>
