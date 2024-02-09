# CI/CD setup

## Prerequisites for CI/CD

To configure Snyk to run in a pipeline, retrieve key configuration inputs from your Snyk account.

### Target Organization

When you run Snyk in your CI/CD platform, you typically want to post the test results to Snyk for review and ongoing monitoring.

**If you do not define a target Organization**, Snyk uses the default Organization for your authentication token:

* For user accounts, this is the user's preferred Organization, configurable in the user's settings.
* For Organization service accounts, this is the Organization in which the account was created.

You can **define the target organization** in the Snyk CLI by using the  `--org` CLI option and either the `orgslugnam`e or Organization ID:

* You can define the target Organization using its `orgslugname` as displayed in the address bar of the browser in the Snyk UI.
* Alternatively, you can define the target Organization using its ORG\_ID on the Organization settings page.

<figure><img src="../../../.gitbook/assets/image1.png" alt="Organization ID"><figcaption><p>Organization ID</p></figcaption></figure>

For more information, see [How to select the Organization to use in the CLI.](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli.md)

### Snyk authentication token

To run `snyk test`, you must authenticate using a token with access to the target Organization. For more information, see [How to obtain and authenticate with your Snyk API token](../../../getting-started/how-to-obtain-and-authenticate-with-your-snyk-api-token.md).

Snyk Free and Team plan users must use their **personal Snyk API token**. See [API token configuration](ci-cd-setup.md#api-token-configuration) on this page for information on how to do this.

Enterprise users should use a **service account** for any type of automation, including CI/CD setup. For more information, see [Service accounts](../../../enterprise-setup/service-accounts/).

## Setting up Snyk to run in a pipeline

Snyk supports the following approaches to add tests to a build pipeline:

* **Snyk integration plugins**: Snyk provides pre-built plugins for several CI servers, including [Jenkins](https://docs.snyk.io/integrations/ci-cd-integrations/jenkins-integration-overview), [Team City](https://docs.snyk.io/integrations/ci-cd-integrations/teamcity-integration-overview)[, Bitbucket Pipelines](https://docs.snyk.io/integrations/ci-cd-integrations/bitbucket-pipelines-integration-overview), and [Azure Pipelines](https://docs.snyk.io/integrations/ci-cd-integrations/azure-pipelines-integration).
* **Snyk CLI:** Teams with more complex workflows, or using a build system without a Snyk pre-built plugin, can use the Snyk CLI during CI/CD setups. See [Setting up using Snyk CLI](ci-cd-setup.md#setting-up-using-snyk-cli) for details.
* **Snyk API**: For teams with complex requirements, Snyk provides a REST API, which you can use for functions including initiating scans, onboarding new Projects, and testing arbitrary libraries. See the [Snyk API documentation](../../../snyk-api/) for details.

## Setting up CI/CD using Snyk CLI

The Snyk CLI is a NodeJS application that can be scripted directly by developers for easy integration into most CI/CD environments. The Snyk CLI is available as an npm application, pre-packaged binary, or container image. For more information, see [Install or update the Snyk CLI](../../../snyk-cli/install-or-update-the-snyk-cli/).

The Snyk CLI can be configured to:

* Return non-zero error codes only when certain criteria are met, for example, exit with an error code only if vulnerabilities of high severity are present.
* Output all of its data into JSON for more flexibility.

## Configure your continuous integration

To continuously avoid known vulnerabilities in your dependencies, integrate Snyk into your continuous integration (build) system. In addition to this documentation, see the [integration configuration examples](https://github.com/snyk-labs/snyk-cicd-integration-examples) in the Snyk Labs GitHub repository.

### Set up automatic monitoring

If you monitor a Project with Snyk, you will be notified if the dependencies in your Project are affected by newly disclosed vulnerabilities. To ensure the list of dependencies Snyk has for your Open Source Project is up to date, refresh it continuously by running s`nyk monitor` in your deployment process. Configure your environment to include the `SNYK_TOKEN` environment variable. You can [find your API token ](../../../snyk-api-info/authentication-for-api.md)in your Snyk account settings.

### API token configuration

Ensure you do not check your personal Snyk API token into source control, to avoid exposing it to others. Instead, use CI environment variables to configure your token.

See guidance for how to do this in the following documentation:

* [Travis](https://docs.travis-ci.com/user/environment-variables/)
* [Circle CI](https://circleci.com/docs/set-environment-variable/)
* [Codeship Basic](https://docs.cloudbees.com/docs/cloudbees-codeship/latest/basic-builds-and-configuration/set-environment-variables), [Codeship Pro](https://docs.cloudbees.com/docs/cloudbees-codeship/latest/pro-builds-and-configuration/environment-variables)

You can find additional documentation through a web search for `setting up env variables in CI`.
