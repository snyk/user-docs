# Scan and fix security issues in Kubernetes configuration files

Snyk Snyk Infrastructure as Code scans your manifest files for security vulnerabilities, and scans your Kubernetes configuration files for misconfigurations and security issues as well. For configuration files, once scanned, Snyk reports on any misconfigurations based on the settings your administrator has implemented, and makes recommendations for fixing accordingly.

## Prerequisites

* An administrator should integrate your organization with your preferred Git repository and enable detection of configuration files as described [here](https://support.snyk.io/hc/articles/360006402818#UUID-c1919782-6bfa-b84b-a638-3913cee39fc5).
* You must have a Snyk account and your configuration files should be in either JSON or YAML format.

Snyk Infrastructure as Code supports:

* Deployments, Pods and Services.
* CronJobs, Jobs, StatefulSet, ReplicaSet, DaemonSet, and ReplicationController.

## Scan and fix your configuration files

1. 1. If you already imported your repositories for testing before cloud configuration file detection was enabled by your administrator, then you should re-import that repository again in order to import the relevant JSON or YAML configuration files:

   ![2.1.png](https://support.snyk.io/hc/article_attachments/4402311164561/2.1.png)

   ![2.2.png](https://support.snyk.io/hc/article_attachments/360010830638/2.2.png)

2. Every time a repository is scanned:
   * Every supported manifest file and every supported configuration file is imported as a separate project, grouped together per repository, similar to this example:

     ![3.png](https://support.snyk.io/hc/article_attachments/360010830578/3.png)

   * If you re-imported the repository in order to import the cloud configuration files, then Snyk imports and tests the configuration files and also re-tests the already imported application manifest files - displaying the test time as "now".
3. Click the project link you're interested in, to view the scan results and to correct your configuration files accordingly:

   ![4.png](https://support.snyk.io/hc/article_attachments/360010755677/4.png)

