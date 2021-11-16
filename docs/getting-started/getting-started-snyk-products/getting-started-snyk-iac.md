# Getting started with Snyk Infrastructure as Code (IaC)

Get started with Snyk IaC to inspect, find and fix issues in configuration files for Terraform or Kubernetes (including Helm) environments. For more information, see [Scan Kubernetes configuration files](../../products/snyk-infrastructure-as-code/scan-kubernetes-configuration-files/) and [Scan Terraform files](../../products/snyk-infrastructure-as-code/scan-terraform-files/).

{% hint style="info" %}
This article describes a process using the Snyk.io UI. For details of using IaC with the Snyk CLI, see [Snyk CLI for Infrastructure as Code](../../products/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/).
{% endhint %}

## **Prerequisites**

Ensure you have:

* A Snyk account (go to [https://snyk.io/](https://snyk.io) and sign up - see [Create a Snyk account](https://docs.snyk.io/getting-started/getting-started-snyk-products) for details).
* An existing Kubernetes or Terraform environment to work in.
* Integrated your Git repository as for other Snyk products - see [Git repository (SCM) integrations](../../features/integrations/git-repository-scm-integrations/) for more details.

For more details, see:

* [Configure your integration to find security issues in your Kubernetes configuration files](https://docs.snyk.io/snyk-infrastructure-as-code/scan-kubernetes-configuration-files/configure-integration-for-security-issues-in-kubernetes-configuration-files)
* [Configure your integration to find security issues in your Terraform files](https://docs.snyk.io/snyk-infrastructure-as-code/scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-filess).

## Stage 1: Import projects

Import projects to test with Snyk, by choosing repositories for Snyk to test and monitor.

1. Select **Projects** from Snyk.io.
2. Select the tool to add the project from (for example GitHub).
3. In **Personal and Organization repositories**, select the repositories to use.&#x20;
4. Click **Add selected repositories** to import the selected repositories into your projects.&#x20;
5. A progress bar appears: click **View log** to see import log results (you can scan both Kubernetes and Terraform files simultaneously)
6. Project import completes.

## Stage 2: View configuration file issues

View results for configuration files in imported projects.

Select **Projects**, then click on the imported project entry, to see information for scanned configuration files, including the number of high, medium and low severity issues found. For example:

![](../../.gitbook/assets/iac\_-\_issues\_list.png)

(Issues are sorted into project types: Helm, Kubernetes and Terraform.)

Click on a project to see more information and details of the issues in a configuration file:

![](../../.gitbook/assets/iac\_-\_select\_config\_file.png)

{% hint style="info" %}
If you encounter any errors during import, see [Importing projects](https://support.snyk.io/hc/en-us/sections/360000923478-Importing-projects) FAQs.
{% endhint %}

## Stage 3: View and fix config files

Act on the recommendations produced by Snyk IaC.

1. IaC results appear as direct issues in the relevant scanned configuration files.
2. Click on an issue to see the details for that issue, and specific recommendations from Snyk IaC.&#x20;
3. Edit the configuration file to fix the issue identified, based on the recommendations, then commit the change.
4. Snyk automatically rescans the changed file, and you can see the change reflected in the issue display.

## For more information

See [Infrastructure as Code](https://docs.snyk.io/snyk-infrastructure-as-code).
