# Configure your integration to find security issues in your Terraform files

Snyk tests and monitors your Terraform files from your source code repositories, guiding you with advice on how you can better secure your cloud environment--catching misconfigurations before you push to production and helping you to fix them.

## Supported Git repositories and Terraform file formats

Snyk scans Terraform (`.tf`) files when they are imported from an integrated Git repository. You can scan a Terraform module repository by importing the repo that holds the module from an SCM or by scanning the directory itself using `snyk iac test` CLI command.

Scanning Terraform files gives you security feedback on everything that is statically configured in the module. To benefit from recurring and scheduled testing, follow best practices and import custom modules directly from an SCM.

See the Snyk blog post about the ability to interpolate variables: [Snyk IaC public beta introduces Terraform plan analysis](https://snyk.io/blog/snyk-iac-public-beta-introduces-terraform-plan-analysis/). This allows scanning of the Terraform Plan output with the CLI, enabling scanning of the entire Terraform deployment to include the output of the modules used to create the deployment.

## Configure Snyk to scan your Terraform configuration files

### **Prerequisites for scanning Terraform files**

* You must be an administrator for the Organization you are configuring in Snyk.
* Ensure you have already integrated your Git repository. For details, see [Git repository (SCM) integrations](../../../../scm-integrations/organization-level-integrations/).

### **Configure Snyk to scan Terraform files**

* Log in to your account and navigate to the relevant Group and Organization that you want to manage.\
  Integrations are managed per Organization.
* Toggle the setting to enable Snyk to detect Infrastructure as code files as shown:

<figure><img src="../../../../.gitbook/assets/snyk-iac-enable.png" alt="Enable detecting infrastructure as code configuration files"><figcaption><p>Enable detecting infrastructure as code configuration files</p></figcaption></figure>

* If needed, review and adjust the **Infrastructure as code** **Severity settings** on the AWS tab in the example.\
  Check to select the file types to scan, CloudFormation, Terraform, or both, and from the pulldown selection, choose the severity level for each API Gateway.

<figure><img src="../../../../.gitbook/assets/image (105) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1) (2).png" alt="Select Severity settings for IaC scans"><figcaption><p>Select Severity settings for IaC scans</p></figcaption></figure>
