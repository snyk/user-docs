# Developer Environment and Snyk

## The developer experience

This section of the workshop demonstrates a common developer practice of scanning your application for vulnerabilities and license compliance using the Snyk CLI. Snyk offers many ways to execute a vulnerability and license scan. Developers can choose to make scanning a part of every build using a Snyk plugin, like Maven or Gradle for Java, or using the Snyk CLI to perform ad-hoc scans. In this section, we will perform an ad-hoc scan, review the results, and configure Maven to use the Snyk plugin.

### Workshop exercises

We will complete the following steps. 

1. Clone our forked version of the repo to our lab VM
2. Authenticate with the Snyk UI using a personal token.  
3. Perform Synk test using the CLI to review vulnerabilities and license issues. We will also use the Snyk CLI to scan container images for vulnerabilities. 
4. Configure Maven to use the Snyk plugin for every build.
5. Validate our IaC Kubernetes files and deploy SPC to Kubernetes.

