# Scan and fix security issues in Terraform files

Snyk scans your Terraform code for misconfigurations and security issues as well. For configuration files, once scanned, Snyk reports on any misconfigurations based on the settings your administrator has implemented, and makes recommendations for fixing accordingly.

### Prerequisites

* An administrator should integrate your organization with your preferred Git repository and enable detection of configuration files as described [here](https://docs.snyk.io/snyk-infrastructure-as-code/scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-filess).
* You must have a Snyk account and your Terraform files should be in `.tf` format.
* We currently detect AWS, Azure and GCP related security issues.

### Scan and fix your configuration files

1. 2. If you already imported your repositories for testing before the infrastructure as code feature was enabled by your administrator, then you should re-import that repository in order to detect the Terraform code:

   ![Screenshot\_2020-07-09\_at\_12.44.03.png](https://support.snyk.io/hc/article_attachments/360009907118/Screenshot_2020-07-09_at_12.44.03.png)

3. Every time a repository is scanned:
   * Every Terraform file is imported as a separate project, grouped together per repository, similar to this example:

     ![Screenshot\_2020-07-09\_at\_12.44.48.png](https://support.snyk.io/hc/article_attachments/360009813417/Screenshot_2020-07-09_at_12.44.48.png)

   * If you re-imported the repository in order to import the Terraform files, then Snyk imports and re-tests the already imported application manifest files - displaying the test time as "now".
4. Click the project link you're interested in, to view the scan results and to help correct your Terraform code: ![Screenshot\_2020-07-09\_at\_12.45.26.png](https://support.snyk.io/hc/article_attachments/360009813457/Screenshot_2020-07-09_at_12.45.26.png)

