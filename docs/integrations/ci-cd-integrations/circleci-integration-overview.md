# CircleCI integration

Snyk integrates with [CircleCI](https://circleci.com) using a **Snyk Orb**, seamlessly scanning your application dependencies and Docker images for open source security vulnerabilities as part of the continuous integration/continuous delivery (CI/CD) workflow.

CircleCI enables users to easily create CI/CD workflows using a group of ready-to-use commands ([Orbs](https://circleci.com/orbs/)) that can be added to your configuration file.

With the Snyk Orb, you can quickly add Snyk scanning to your CI/CD in order to test and monitor for open source vulnerabilities, based on your configurations. Results are then displayed from the CircleCI output view and can also be monitored from [Snyk.io](http://app.snyk.io).

## Getting started

Getting started with CircleCI from scratch to a green build with Snyk is simple. See the following for information about the Snyk Orb:

* [Snyk Circle CI README](https://circleci.com/orbs/registry/orb/snyk/snyk) - includes all the information that you need to set up your CI/CD with Snyk including a list of parameters and samples
* [A Circle CI blog post](https://circleci.com/blog/adding-application-and-image-scanning-to-your-cicd-pipeline/) - discusses how to set up a secure pipeline with the Snyk Orb

## Configure your CircleCI integration

Once you add a project to CircleCI and add the Snyk Orb to the configuration file, every time that a build runs, the Snyk Orb is also used and performs the following actions:

### Scan

1. Scans app dependencies or container images for vulnerabilities or licensing issues, and lists the vulnerabilities and issues.
2. If Snyk finds vulnerabilities, it does one of the following (based on your configuration):
   * Fails the build
   * Lets the build complete

### **Monitor**

Optionally, if the build completes successfully and MONITOR is set to True in the Snyk step, then Snyk saves a snapshot of the project dependencies from the Snyk Web UI, where you can view the dependency tree displaying all of the issues, and you can receive alerts for new issues found in the existing app version.

## **Prerequisites**

1. Create a Snyk account and retrieve the Snyk API token from your Account settings.
2. Import the relevant repo into CircleCI.
3. Go to `Settings -> Security -> Orb security settings` and make sure you allow to opt-in to third party Orbs.
4. Make sure your configuration (`config.yml`) file follows version 2.1.
5. Add the required environment variables to CircleCI (including the Snyk API token as `SNYK_TOKEN`).

## Getting Snyk Orb details from the CircleCI registry

From the [Orbs registry](https://circleci.com/orbs/registry/), CircleCI displays a list of available Orbs customized for you directly, similar to the image that follows.

![Snyk Orb for CircleCI](../../.gitbook/assets/uuid-10d3ba7f-799b-45a9-5c8e-b2abe9aab955-en.png)

From this list, find and click the relevant #Snyk line to view the Snyk Orb information with examples, parameters, and values:

![Snyk Orb information](../../.gitbook/assets/uuid-ce212e67-b7ac-3cf7-4772-c84f6897aed9-en.png)
