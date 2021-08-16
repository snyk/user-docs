# Getting started with Snyk Infrastructure as Code \(IaC\)

Get started with Snyk IaC to inspect, find and fix issues in configuration files for Terraform or Kubernetes \(including Helm\) environments. For more information, see [Scan your Kubernetes configuration files](https://support.snyk.io/hc/en-us/sections/360001881957-Scan-your-Kubernetes-configuration-files) and [Scan your Terraform files](https://support.snyk.io/hc/en-us/sections/360003156537-Scan-your-Terraform-files).

#### Stage 2: View configuration file issues

View results for configuration files in imported projects.

1. Select **Projects**, then click on the imported project entry, to see information for scanned configuration files, including the number of high, medium and low severity issues found. For example: ![IaC\_-\_issues\_list.png](https://support.snyk.io/hc/article_attachments/360012553158/IaC_-_issues_list.png) \(Issues are sorted into project types: Helm, Kubernetes and Terraform.\)
2. Click on a project to see more information and details of the issues in a configuration file: ![IaC\_-\_select\_config\_file.png](https://support.snyk.io/hc/article_attachments/360012553198/IaC_-_select_config_file.png)

If you encounter any errors during import, see the [Importing projects](https://support.snyk.io/hc/en-us/sections/360000923478-Importing-projects) information.

#### Stage 3: View and fix config files

Act on the recommendations produced by Snyk IaC.

1. IaC results appear as direct issues in the relevant scanned configuration files. For example:

   ![IaC\_-\_issue\_details.png](https://support.snyk.io/hc/article_attachments/360012478437/IaC_-_issue_details.png)

2. Click on an issue to see the details for that issue, and specific recommendations from Snyk IaC. For example:

   ![IaC\_-\_corrected.png](https://support.snyk.io/hc/article_attachments/360012806977/IaC_-_corrected.png)

3. Edit the configuration file to fix the issue identified, based on the recommendations, then commit the change.
4. Snyk automatically rescans the changed file, and you can see the change reflected in the issue display.

### For more information

See [Infrastructure as Code](https://docs.snyk.io/snyk-infrastructure-as-code).

