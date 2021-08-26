# Scan and fix security issues in your CloudFormation files

Snyk scans CloudFormation code for misconfigurations and security issues. For configuration files, once scanned, Snyk reports on any misconfigurations based on the settings administrators implement and makes recommendations for fixes accordingly.

### **Prerequisites**

* An administrator integrates the organization with a preferred Git repository and enables the detection of configuration files as described in [Configure your integration to find security issues in your CloudFormation files](https://docs.snyk.io/snyk-infrastructure-as-code/scan-cloudformation-files/configure-your-integration-to-find-security-issues-in-your-cloudformation-files).
* The Snyk account and CloudFormation files will be in `JSON` and `YAML` formats.

## Scan and fix configuration files

* Log in to the account and navigate to the relevant group and organization.

![](../../.gitbook/assets/screenshot-2020-07-09-at-12.43.02-2-%20%283%29%20%284%29%20%284%29%20%284%29%20%283%29.png)

* Re-import repositories if testing occurred before the infrastructure as code feature was enabled in order to detect the CloudFormation code:

 

![Screenshot\_2020-07-09\_at\_12.44.03.png](../../.gitbook/assets/screenshot_2020-07-09_at_12.44.03%20%281%29%20%281%29%20%283%29.png)

* Every time a repository is scanned:
  * Every CloudFormation file is imported as a separate project, grouped together per repository, similar to this example:

   

![Screen\_Shot\_2021-06-23\_at\_10.16.38.png](../../.gitbook/assets/screen_shot_2021-06-23_at_10.16.38.png)

  * If the repository was re-imported: in order to then import the CloudFormation files, Snyk imports and re-tests the existing application manifest files--displaying the test time as "now".
* Click a project link to view the scan results and to help view details on the CloudFormation code:![Screen\_Shot\_2021-06-23\_at\_10.18.49.png](../../.gitbook/assets/screen_shot_2021-06-23_at_10.18.49.png)

