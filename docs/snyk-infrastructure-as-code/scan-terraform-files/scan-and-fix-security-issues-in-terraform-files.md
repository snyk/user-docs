# Scan and fix security issues in Terraform  files

Snyk scans your Terraform code for misconfigurations and security issues as well. For configuration files, once scanned, Snyk reports on any misconfigurations based on the settings your administrator has implemented, and makes recommendations for fixing accordingly.

## Prerequisites

* An administrator should integrate your organization with your preferred Git repository and enable detection of configuration files as described [here](https://support.snyk.io/hc/en-us/articles/360011018938-Configure-your-integration-to-find-security-issues-in-your-Terraform-files).
* You must have a Snyk account and your Terraform files should be in `.tf` format.
* We currently detect AWS, Azure and GCP related security issues.

## Scan and fix your configuration files

* Log in to your account and navigate to the relevant group and organization that you want to manage. 

![](../../.gitbook/assets/screenshot-2020-07-09-at-12.43.02-2-%20%283%29%20%282%29.png)

* If you already imported your repositories for testing before the infrastructure as code feature was enabled by your administrator, then you should re-import that repository in order to detect the Terraform code:
* ![Screenshot\_2020-07-09\_at\_12.44.03.png](../../.gitbook/assets/screenshot_2020-07-09_at_12.44.03.png)
* Every time a repository is scanned:
  * Every Terraform file is imported as a separate project, grouped together per repository, similar to this example:

    ![Screenshot\_2020-07-09\_at\_12.44.48.png](../../.gitbook/assets/screenshot_2020-07-09_at_12.44.48.png)

  * If you re-imported the repository in order to import the Terraform files, then Snyk imports and re-tests the already imported application manifest files - displaying the test time as "now".
* Click the project link you're interested in, to view the scan results and to help correct your Terraform code: ![Screenshot\_2020-07-09\_at\_12.45.26.png](../../.gitbook/assets/screenshot_2020-07-09_at_12.45.26.png)

