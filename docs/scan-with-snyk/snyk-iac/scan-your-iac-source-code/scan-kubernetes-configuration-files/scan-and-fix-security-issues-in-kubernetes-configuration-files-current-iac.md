# Scan and fix security issues in Kubernetes configuration files

Snyk Infrastructure as Code scans your manifest files for security vulnerabilities and scans your Kubernetes configuration files for misconfigurations and security issues as well. After configuration files are scanned, Snyk reports on any misconfigurations based on the settings your administrator has implemented and makes recommendations for fixing accordingly.

## Prerequisites for scanning and fixing issues in Kubernetes files

* An administrator should [integrate your Organization](../scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-files-current-iac.md) with your preferred Git repository and enable the detection of configuration files.
* You must have a Snyk account, and your configuration files should be in either JSON or YAML format.

Snyk Infrastructure as Code supports:

* Deployments, Pods, and Services.
* CronJobs, Jobs, StatefulSet, ReplicaSet, DaemonSet, and ReplicationController.

## Scan and fix your Kubernetes configuration files

* Log in to your account and navigate to the relevant Group and Organization that you want to manage.
* If you imported your repositories for testing before cloud configuration file detection was enabled by your administrator, re-import that repository to import the relevant JSON or YAML configuration files.

Every time a repository is scanned, every supported manifest file and every supported configuration file is imported as a separate Project, grouped together per repository.

When you re-import the repository in order to import the cloud configuration files, Snyk imports and tests the configuration files and re-tests the already imported application manifest files, displaying the test time as "now".

* Click the link for the Project of interest to you to view the scan results and to correct your configuration files accordingly:

<figure><img src="../../../../.gitbook/assets/image (343) (1) (1).png" alt="Kubernetes Proejct detail"><figcaption><p>Kubernetes Proejct detail</p></figcaption></figure>
