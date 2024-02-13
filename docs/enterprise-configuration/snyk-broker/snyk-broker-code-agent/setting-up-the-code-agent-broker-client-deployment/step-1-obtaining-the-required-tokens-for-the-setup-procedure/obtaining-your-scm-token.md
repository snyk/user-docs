# Obtaining your SCM token

Your integrated SCM token is required for [setup of the Broker Client component](../step-5-setting-up-the-broker-client/step-5.2a-running-the-broker-client-without-the-code-snippet-display.md), and it is used in the `-e <SCM>_TOKEN`, for example, `-e GITHUB_TOKEN=xxx…`. This token is required for accessing the SCM with certain permissions, which are needed for the operation of the Broker and Snyk Code.

**To obtain your SCM token,** follow the instructions provided by the SCM you want to integrate with the Snyk Broker, and create a token with the required permissions.

The following SCM tokens are required for the different SCMs:

**GitHub and GitHub Enterprise**:

`GITHUB_TOKEN=` - a GitHub personal access token. Scopes: **`repo, read:org`** and **`admin:repo_hook`**.

See GitHub documentation - [_Creating a personal access token_](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)\_\_

**Gitlab**:

**`GITLAB_TOKEN=`** - a GitLab personal access token. Gitlab account with **`Maintainer`** permissions. Scope: **`api`**.

See Gitlab documentation - [_Personal access tokens_](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html)\_\_

**Azure Repos**:

**`AZURE_REPOS_TOKEN=`** - an Azure Repos personal access token. Scopes: **`Custom defined`, \*\* `Code:` \*\* `Read & write`**_._

See Azure Repos documentation - [_Use personal access tokens_](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=Windows)\_\_

**Bitbucket Server/Data Center**:

**`BITBUCKET_USERNAME=`**, **`BITBUCKET_PASSWORD=`** – the Bitbucket Server username and password or a Bitbucket Server personal access token. Scope: **`Repository admin`**.

See Bitbucket Server documentation - [_Personal access tokens_](https://confluence.atlassian.com/bitbucketserver/http-access-tokens-939515499.html)\\
