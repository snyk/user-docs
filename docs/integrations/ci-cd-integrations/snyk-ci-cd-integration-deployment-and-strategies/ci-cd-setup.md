# CI/CD setup

## Prerequisites

To configure Snyk to run in a pipeline, retrieve key configuration inputs from your Snyk account.

### Target organization

When you run Snyk in your CI/CD platform, you typically want to post the test results to Snyk for review and ongoing monitoring.

**If you do not define a target organization**, Snyk uses the default organization for the authentication token you use:

* For user accounts, this is the user's preferred organization (configurable in the user's settings).
* For organization service accounts, this is the organization in which the account was created.

You can **define the target organization** in the Snyk CLI, by either URL orgslugname or organization ID, using the `--org` CLI option:

* You can define the target organization using its URL slug (orgslugname), as displayed in the address bar of the browser in the Snyk UI.
* Alternatively you can define the target organization using its `ORG_ID` in the organization's settings page.

<figure><img src="../../../.gitbook/assets/image1.png" alt="Organization ID"><figcaption><p>Organization ID</p></figcaption></figure>

For more information see see [How to select the organization to use in the CLI.](../../../snyk-cli/test-for-vulnerabilities/how-to-select-the-organization-to-use-in-the-cli.md)

### Snyk authentication token

To run `snyk test`, you need an authentication token with access to the target organization . While you can use any valid authentication token, using a service account is recommended. For more details, see the [`snyk auth` command help](../../../snyk-cli/commands/auth.md) and [Service accounts](https://docs.snyk.io/integrations/managing-integrations/service-accounts).

## Setting up Snyk to run in a pipeline

Snyk supports the following approaches to add tests to a build pipeline:

* **Snyk integration plugins**: Snyk provides pre-built plugins for several CI servers, including [Jenkins](https://docs.snyk.io/integrations/ci-cd-integrations/jenkins-integration-overview), [Team City](https://docs.snyk.io/integrations/ci-cd-integrations/teamcity-integration-overview)[, Bitbucket Pipelines](https://docs.snyk.io/integrations/ci-cd-integrations/bitbucket-pipelines-integration-overview) and [Azure Pipelines](https://docs.snyk.io/integrations/ci-cd-integrations/azure-pipelines-integration).
* **Snyk CLI:** Teams with more complex workflows, or using a build system without a Snyk pre-built plugin, can use the Snyk CLI tool during CI/CD setups. See [Setting up using Snyk CLI](ci-cd-setup.md#setting-up-using-snyk-cli) for details.
* **Snyk API**: For teams with complex requirements Snyk provides a REST API, which you can use for functions including initiating scans, onboarding new projects, and testing arbitrary libraries. See the [Snyk API documentation](../../../snyk-api-info/) for more details.

## Setting up using Snyk CLI

Snyk CLI is a NodeJS application that can be scripted directly by developers for easy integration into most CI/CD environments, and is available as an npm application, pre-packaged binary, or container image. For more information see [Install the Snyk CLI](https://docs.snyk.io/snyk-cli/install-the-snyk-cli).

Snyk CLI can be configured to:

* Return non-zero error codes only when certain criteria are met, for example, exit with an error code only if vulnerabilities of high severity are present.
* Output all of its data into JSON for more flexibility.

## Configure your continuous integration

To continuously avoid known vulnerabilities in your dependencies, integrate Snyk into your continuous integration (build) system. In addition to this documentation, see the [integration configuration examples](https://github.com/snyk-labs/snyk-cicd-integration-examples) in the Snyk Labs GitHub repository.

### Set up automatic monitoring

If you monitor a project with Snyk, you’ll be notified if the dependencies in your project are affected by newly disclosed vulnerabilities. To make sure the list of dependencies Snyk has for your project is up to date, refresh it continuously by running Snyk monitor in your deployment process. Configure your environment to include the SNYK\_TOKEN environment variable. You can find your API token on the dashboard after logging in.

### API token configuration

Make sure you do not check your API token into source control, to avoid exposing it to others. Instead, use your CI environment variables to configure it.

See guidance for how to do this on:

* [Travis](https://docs.travis-ci.com/user/environment-variables/)
* [Circle CI](https://circleci.com/docs/set-environment-variable/)
* [Codeship Basic](https://docs.cloudbees.com/docs/cloudbees-codeship/latest/basic-builds-and-configuration/set-environment-variables), [Codeship Pro](https://docs.cloudbees.com/docs/cloudbees-codeship/latest/pro-builds-and-configuration/environment-variables)

You can find others through a [Google search](https://www.google.co.uk/search?q=setting+up+env+variables+in+CI).
