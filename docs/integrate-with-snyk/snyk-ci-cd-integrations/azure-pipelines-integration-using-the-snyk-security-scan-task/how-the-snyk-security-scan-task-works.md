# How the Snyk Security Scan task works

After the Snyk Security Scan task is added to a pipeline, each time the pipeline runs, the Snyk task performs the following actions:

## **Test**

1. Snyk tests the application dependencies or container images for vulnerabilities and open source license issues and lists the vulnerabilities and issues.
2. If Snyk finds vulnerabilities or license issues, it does one of the following (based on your configuration):
   * Fails the pipeline
   * Lets the pipeline continue

## **Monitor**

After the snyk test is complete, you have the option of doing `snyk monitor.` `snyk monitor` saves a snapshot of the project dependencies in your [snyk.io](https://snyk.io) account, where you can see the dependency tree with all of the issues and be alerted if and when new issues are found in the dependencies.
