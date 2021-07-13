# Snyk runtime monitoring: disable the Snyk agent

* [ Configure your integration to find security issues in your CloudFormation files](/hc/en-us/articles/4402937668241-Configure-your-integration-to-find-security-issues-in-your-CloudFormation-files)
* [ Scan and fix security issues in your CloudFormation files](/hc/en-us/articles/4402971349009-Scan-and-fix-security-issues-in-your-CloudFormation-files)
* [ Detecting CloudFormation configuration files using a broker](/hc/en-us/articles/4402964063377-Detecting-CloudFormation-configuration-files-using-a-broker)

##  Scan and fix security issues in your CloudFormation files

Snyk scans CloudFormation code for misconfigurations and security issues. For configuration files, once scanned, Snyk reports on any misconfigurations based on the settings administrators implement and makes recommendations for fixes accordingly.

### Scan and fix configuration files

1. 2. Re-import repositories if testing occurred before the infrastructure as code feature was enabled in order to detect the CloudFormation code:
3. Every time a repository is scanned:
   * Every CloudFormation file is imported as a separate project, grouped together per repository, similar to this example:
   * If the repository was re-imported: in order to then import the CloudFormation files, Snyk imports and re-tests the existing application manifest files--displaying the test time as "now".
4. Click a project link to view the scan results and to help view details on the CloudFormation code:

