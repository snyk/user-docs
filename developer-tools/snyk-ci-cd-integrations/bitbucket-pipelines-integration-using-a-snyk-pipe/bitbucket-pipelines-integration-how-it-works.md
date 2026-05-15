# Bitbucket Pipelines integration: how it works

After you have added the Snyk pipe to the pipeline, each time the pipeline executes (by any trigger type), the Snyk pipe performs the following actions:

## **Scan**

1. Snyk scans app dependencies or container images for vulnerabilities and open-source license issues, and lists the vulnerabilities and issues.
2. If Snyk finds vulnerabilities, it does one of the following (based on your configuration):
   * Fails the build
   * Lets the build complete

## **Monitor**

Optionally, if the build completes successfully and **MONITOR** is set to **True** in the Snyk step, then Snyk saves a snapshot of the Project dependencies from the Snyk Web UI. From the Snyk Web UI, you can view the dependency tree displaying all of the issues and receive alerts for new issues found in the existing app version.
