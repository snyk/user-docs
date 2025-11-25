# Scan and fix security issues in Terraform files

Snyk scans your Terraform code for misconfigurations and security issues as well. After scanning configuration files, Snyk reports on any misconfigurations based on the settings your administrator implemented, and makes recommendations for fixing accordingly.

## Prerequisites for scanning and fixing issues in Terraform files in SCM repositories

* An administrator should integrate your Organization with your preferred Git repository and enable detection of configuration files as described on [Configure your integration to find security issues in your Terraform files](configure-your-integration-to-find-security-issues-in-your-terraform-files-current-iac.md)
* You must have a Snyk account, and your Terraform files should be in `.tf` format.
* Snyk detects AWS, Azure, and Google Cloud-related security issues.

## Scan and fix your Terraform configuration files

* Log in to your account and navigate to the relevant Group and Organization that you want to manage.
* If you imported your repositories for testing before the infrastructure as code feature was enabled by your administrator, from the **Add project** screen, re-import that repository in order to detect the Terraform code:

<figure><img src="broken-reference" alt="Add project screen"><figcaption><p>Add project screen</p></figcaption></figure>

Every time a repository is scanned, every Terraform file is imported as a separate Project, grouped together per repository, similar to the example shown.

If you re-imported the repository in order to import the Terraform files, then Snyk imports and re-tests the already imported application manifest files, displaying the test time as "now".

<figure><img src="../../../../.gitbook/assets/image (63).png" alt="List of Terraform Projects"><figcaption><p>List of Terraform Projects</p></figcaption></figure>

* Click the link for the Project of interest to you to view the scan results and to help correct your Terraform code:

<figure><img src="broken-reference" alt="Terraform Project detail"><figcaption><p>Terraform Project detail</p></figcaption></figure>
