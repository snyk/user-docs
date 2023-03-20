# Getting started with Snyk Infrastructure as Code (IaC)

Get started with Snyk IaC to inspect, find, and fix issues in configuration files for [Terraform](scan-terraform-files/), [AWS CloudFormation](scan-cloudformation-files/), [Kubernetes](scan-kubernetes-configuration-files/) (including Helm), or [Azure Resource Manager (ARM)](scan-arm-configuration-files.md) environments.

{% hint style="info" %}
This article describes a process using the Snyk Web UI. For details of using IaC with the Snyk CLI, see [snyk-cli-for-infrastructure-as-code](snyk-cli-for-infrastructure-as-code/ "mention"). Note that [ARM configuration files](scan-arm-configuration-files.md) can only be scanned via the CLI.
{% endhint %}

## **Prerequisites**

Ensure you have:

* A Snyk account (go to [https://snyk.io/](https://snyk.io) and sign up - see [Create a Snyk account](../../getting-started/quickstart/create-a-snyk-account/) for details).
* An existing Terraform, CloudFormation, Kubernetes, or ARM environment to work in.
* Integrated your Git repository as for other Snyk products - see [Git repository (SCM) integrations](../../integrations/git-repository-scm-integrations/) for more details.

For more details, see:

* [Configure your integration to find security issues in your Terraform files](scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-filess.md)
* [Configure your integration to find security issues in your CloudFormation files](scan-cloudformation-files/configure-your-integration-to-find-security-issues-in-your-cloudformation-files.md)
* [Configure your integration to find security issues in your Kubernetes configuration files](scan-kubernetes-configuration-files/configure-integration-for-security-issues-in-kubernetes-configuration-files.md)

{% hint style="info" %}
ARM configuration files can only be scanned via the Snyk CLI. See [Scan ARM configuration files](scan-arm-configuration-files.md).
{% endhint %}

## Stage 1: Import projects

Import [Projects](../../manage-issues/introduction-to-snyk-projects/#project) to test with Snyk by choosing repositories for Snyk to test and monitor.

1. Select **Projects** from the Snyk Web UI.
2. In the **Add projects** drop-down menu, select the tool to add the Project from (for example, GitHub).
3. In **Personal and Organization repositories**, select the repositories to use.
4. Click **Add selected repositories** to import the selected repositories into your projects.
5. Select **View import Log** to see import log results (you can scan multiple types of configuration files simultaneously).
6. Project import completes.

{% hint style="info" %}
Currently Snyk Infrastructure as Code Projects have a recurring test default interval of **1 week**. The default interval is changed on the **Settings** tab of the Project's page.
{% endhint %}

## Stage 2: View configuration file issues

View results for configuration files in imported Projects by selecting **Projects** from the menu on the left.

* If **Group by targets** is selected: A list of [Targets](../../manage-issues/introduction-to-snyk-projects/#target) is displayed. Select a Target to expand its list of Projects.
* If **Group by none** is selected: A list of all [Projects](../../manage-issues/introduction-to-snyk-projects/#project) is displayed.

Each Project entry shows information for a scanned configuration file, including the number of critical, high, medium, and low severity issues found. For example:

<figure><img src="../../.gitbook/assets/snyk-iac-getting-started-list-of-projects.png" alt="A list of Snyk IaC Projects"><figcaption><p>A list of Snyk IaC Projects</p></figcaption></figure>

Select a Project to see more information, including details of the issues in the configuration file:

<figure><img src="../../.gitbook/assets/snyk-iac-getting-started-project-page.png" alt="An example Snyk IaC Project with a list of issues"><figcaption><p>An example Snyk IaC Project with a list of issues</p></figcaption></figure>

{% hint style="info" %}
If you encounter any errors during import, see [Importing projects](https://support.snyk.io/hc/en-us/sections/360000923478-Importing-projects) FAQs.
{% endhint %}

## Stage 3: View and fix config files

Act on the recommendations produced by Snyk IaC. IaC results appear as issues in each Project.

1. From a Project page, select an issue to see the details for that issue and specific recommendations from Snyk IaC.
2. Edit the configuration file to fix the issue identified, based on the recommendations, then commit the change.
3. Snyk automatically rescans the changed file, and you can see the change reflected in the issue display.

<figure><img src="../../.gitbook/assets/snyk-iac-getting-started-issue-card.png" alt="An example IaC issue within a Project"><figcaption><p>An example IaC issue within a Project</p></figcaption></figure>

## For more information

See [Using Snyk IaC with the Web UI](using-snyk-iac-via-web.md) and [Snyk CLI for Infrastructure as Code](snyk-cli-for-infrastructure-as-code/).
