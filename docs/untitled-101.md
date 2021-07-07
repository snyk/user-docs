# CircleCI integration overview

* [ Configure your Continuous Integration](/hc/en-us/articles/360004002258-Configure-your-Continuous-Integration)
* [ Continuous Integration: language support](/hc/en-us/articles/360004032157-Continuous-Integration-language-support)
* [ AWS CodePipeline integration](/hc/en-us/articles/4402158184081-AWS-CodePipeline-integration)
* [ Azure Pipelines integration](/hc/en-us/articles/360004127677-Azure-Pipelines-integration)
* [ Bitbucket Pipelines integration overview](/hc/en-us/articles/360004032177-Bitbucket-Pipelines-integration-overview)
* [ CircleCI integration overview](/hc/en-us/articles/360004002278-CircleCI-integration-overview)
* [ Configure your CircleCI integration](/hc/en-us/articles/360004002298-Configure-your-CircleCI-integration)
* [ Getting Snyk Orb details from the CircleCI registry](/hc/en-us/articles/360004032197-Getting-Snyk-Orb-details-from-the-CircleCI-registry)
* [ GitHub Actions integration](/hc/en-us/articles/360019718618-GitHub-Actions-integration)
* [ Bitbucket Pipelines integration](/hc/en-us/articles/360000921778-Bitbucket-Pipelines-integration)

 [See more](/hc/en-us/sections/360001152577-CI-CD-integrations)

##  CircleCI integration overview

Snyk integrates with [CircleCI](https://circleci.com/) using a Snyk **Orb**, seamlessly scanning your application dependencies and Docker images for open source security vulnerabilities as part of the continuous integration/continuous delivery \(CI/CD\) workflow.

CircleCI enables users to easily create CI/CD workflows using a group of ready-to-use commands \([Orbs](https://circleci.com/orbs/)\) that can be added to your configuration file.

With the Snyk Orb, you can quickly add Snyk scanning to your CI/CD in order to test and monitor for open source vulnerabilities, based on your configurations. Results are then displayed from the CircleCI output view and can also be monitored from [Snyk.io](http://app.snyk.io/).

Getting started with CircleCI from scratch to a green build with Snyk is simple! You can read all about the Snyk Orb here:

* [Our Circle CI README](https://circleci.com/orbs/registry/orb/snyk/snyk) - the page includes all the info that you need in order to set your CI/CD with Snyk including a list of parameters and samples.
* [A Circle CI blog](https://circleci.com/blog/adding-application-and-image-scanning-to-your-cicd-pipeline/) - discussing how to set up a secure pipeline with the Snyk orb.

###  Configure your CircleCI integration

Once the user adds a project to CircleCI and adds the Snyk Orb to the configuration file, every time that a build will run, the Snyk Orb will be used as well.

1. Create a Snyk account and retrieve the Snyk API token from your Account settings.
2. Import the relevant repo into CircleCI.
3. Go to `Settings -> Security -> Orb security settings` and make sure you allow to opt-in to third party Orbs.
4. Make sure your configuration \(`config.yml`\) file follows version 2.1.
5. Add the required variables to CircleCI \(e.g. Snyk API token as `API_TOKEN`\)

###  Getting Snyk Orb details from the CircleCI registry

1. From the [Orbs registry](https://circleci.com/orbs/registry/), CircleCI displays a list of available Orbs customized for you directly, similar to the following image:
2. From this list, find and click the relevant \#Snyk line to view the Orb's information with examples, parameters, and values:

