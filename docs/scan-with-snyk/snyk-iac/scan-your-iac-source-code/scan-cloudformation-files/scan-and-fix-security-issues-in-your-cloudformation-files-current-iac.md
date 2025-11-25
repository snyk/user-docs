# Scan and fix security issues in your CloudFormation files

Snyk scans CloudFormation code for misconfigurations and security issues. After configuration files are scanned, Snyk reports on any misconfigurations based on the settings that administrators implement and makes recommendations for fixes accordingly.

## **Prerequisites for scanning and fixing issues in CloudFormation files**

* An administrator integrates the Organization with a preferred Git repository and enables the detection of configuration files as described in [Configure your integration to find security issues in your CloudFormation files](configure-your-integration-to-find-security-issues-in-your-cloudformation-files-current-iac.md).
* The Snyk account and CloudFormation files are in `JSON` and `YAML` formats.

## Scan and fix configuration files

* Log in to the account and navigate to the relevant Group and Organization.
* If you imported your repositories for testing before the infrastructure as code feature was enabled by your administrator, from the Add project screen, re-import that repository in order to detect the CloudFormation code:

<figure><img src="broken-reference" alt="Add project screen"><figcaption><p>Add project screen</p></figcaption></figure>

Every time a repository is scanned, every CloudFormation file is imported as a separate Project, grouped together per repository, similar to the example shown.

If you re-imported the repository in order to import the CloudFormation files, then Snyk imports and re-tests the already imported application manifest files, displaying the test time as "now".

<figure><img src="../../../../.gitbook/assets/image (49).png" alt="List of CloudFormation Projects"><figcaption><p>List of CloudFormation Projects</p></figcaption></figure>

* Click a Project link to view the scan results and details for the CloudFormation code:

<figure><img src="broken-reference" alt="CloudFormation Project detail"><figcaption><p>CloudFormation Project detail</p></figcaption></figure>
