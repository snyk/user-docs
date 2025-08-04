# How the Snyk Security Scan task works

After the Snyk Security Scan task is added to a pipeline, each time the pipeline runs, the Snyk task performs the following actions:

## **Test**

1. Snyk tests application dependencies or container images for vulnerabilities and open-source license issues or application source code for security vulnerabilities and lists these vulnerabilities and issues.
2. If Snyk finds vulnerabilities or license issues, it does one of the following (based on your configuration):
   * Fails the pipeline
   * Allows the pipeline to continue

## **Monitor**

After the snyk test is complete on the dependencies for an application or a container image, you can run `snyk monitor`. `snyk monitor` saves a snapshot of the Project dependencies in your [snyk.io](https://snyk.io) account, where you can see the dependency tree with all of the issues and be alerted if and when new issues are found in the dependencies.
