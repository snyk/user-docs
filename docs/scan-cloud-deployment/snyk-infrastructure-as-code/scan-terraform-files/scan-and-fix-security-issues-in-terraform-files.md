# Scan and fix security issues in Terraform files

Snyk scans your Terraform code for misconfigurations and security issues as well. For configuration files, once scanned, Snyk reports on any misconfigurations based on the settings your administrator has implemented, and makes recommendations for fixing accordingly.

## Prerequisites

* An administrator should integrate your organization with your preferred Git repository and enable detection of configuration files as described [here](https://docs.snyk.io/snyk-infrastructure-as-code/scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-filess).
* You must have a Snyk account and your Terraform files should be in `.tf` format.
* We currently detect AWS, Azure and GCP related security issues.

## Scan and fix your configuration files

* Log in to your account and navigate to the relevant group and organization that you want to manage.
* If you already imported your repositories for testing before the infrastructure as code feature was enabled by your administrator, then you should re-import that repository in order to detect the Terraform code:

![](<../../../.gitbook/assets/screenshot\_2020-07-09\_at\_12.44.03 (1) (1) (3) (3) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1)  (18).png>)

* Every time a repository is scanned:
  * Every Terraform file is imported as a separate project, grouped together per repository, similar to this example: (If you re-imported the repository in order to import the Terraform files, then Snyk imports and re-tests the already imported application manifest files - displaying the test time as "now".)

![](<../../../.gitbook/assets/image (165) (1) (1) (1) (1).png>)

* Click the project link you're interested in, to view the scan results and to help correct your Terraform code:

![](<../../../.gitbook/assets/image (340) (1) (1) (1).png>)
