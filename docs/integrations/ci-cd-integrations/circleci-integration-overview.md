# CircleCI integration overview

Snyk integrates with [CircleCI](https://circleci.com/) using a Snyk **Orb**, seamlessly scanning your application dependencies and Docker images for open source security vulnerabilities as part of the continuous integration/continuous delivery \(CI/CD\) workflow.

CircleCI enables users to easily create CI/CD workflows using a group of ready-to-use commands \([Orbs](https://circleci.com/orbs/)\) that can be added to your configuration file.

With the Snyk Orb, you can quickly add Snyk scanning to your CI/CD in order to test and monitor for open source vulnerabilities, based on your configurations. Results are then displayed from the CircleCI output view and can also be monitored from [Snyk.io](http://app.snyk.io/).

Getting started with CircleCI from scratch to a green build with Snyk is simple! You can read all about the Snyk Orb here:

* [Our Circle CI README](https://circleci.com/orbs/registry/orb/snyk/snyk) - the page includes all the info that you need in order to set your CI/CD with Snyk including a list of parameters and samples.
* [A Circle CI blog](https://circleci.com/blog/adding-application-and-image-scanning-to-your-cicd-pipeline/) - discussing how to set up a secure pipeline with the Snyk orb.

## Configure your CircleCI integration

Once the user adds a project to CircleCI and adds the Snyk Orb to the configuration file, every time that a build will run, the Snyk Orb will be used as well.

1. Create a Snyk account and retrieve the Snyk API token from your Account settings.
2. Import the relevant repo into CircleCI.
3. Go to `Settings -> Security -> Orb security settings` and make sure you allow to opt-in to third party Orbs.
4. Make sure your configuration \(`config.yml`\) file follows version 2.1.
5. Add the required variables to CircleCI \(e.g. Snyk API token as `API_TOKEN`\)

## Getting Snyk Orb details from the CircleCI registry

1. From the [Orbs registry](https://circleci.com/orbs/registry/), CircleCI displays a list of available Orbs customized for you directly, similar to the following image:

   ![image1.png](../../.gitbook/assets/uuid-10d3ba7f-799b-45a9-5c8e-b2abe9aab955-en.png)

2. From this list, find and click the relevant \#Snyk line to view the Orb's information with examples, parameters, and values:

   ![image2.png](../../.gitbook/assets/uuid-ce212e67-b7ac-3cf7-4772-c84f6897aed9-en.png)

