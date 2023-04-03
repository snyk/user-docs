# Snyk Broker - Client integration setups with Helm

After you add the [Snyk Broker Helm Char](how-to-install-and-configure-your-snyk-broker-client.md)t, run the following commands based on the repository type.

The following are the allowed values for `scmType`:

Github.com: `github-com`\
Github Enterprise: `github-enterprise`\
Bitbucket: `bitbucket-server`\
Gitlab: `gitlab`\
Azure Repos: `azure-repos`\
Artifactory: `artifactory`\
Jira: `jira`\
Container Registry Agent: `container-registry-agent`\\

The following examples create a namespace called `snyk-broker`. To deploy into an existing namespace, adjust the `-n` parameter and delete the `--create-namespace` parameter. See alao [Deploying multiple Brokers in the same namespace](set-up-snyk-broker/how-to-install-and-configure-your-snyk-broker-client/deploying-multiple-brokers-in-the-same-namespace.md).

## Github.com Helm install

For more details including definitions of the environment variables see [Set up Snyk Broker with GitHub](snyk-broker-set-up-examples/broker-example-set-up-snyk-broker-with-github.md).

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

## Github Enterprise Helm install

For more details including definitions of the environment variables see [Set up Snyk Broker with GitHub Enterprise](snyk-broker-set-up-examples/setup-broker-with-github-enterprise.md).

Note: for `github`, `githubApi` and `githubGraphQl` values do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-enterprise \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set github=<ENTER_GHE_ADDRESS> \
             --set githubApi=<ENTER_GHE_API_ADDRESS> \
             --set githubGraphQl=<ENTER_GRAPHQL_ADDRESS> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

## Bitbucket Helm install

For more details including definitions of the environment variables see [Set up Snyk Broker with Bitbucket Server/Data Center](snyk-broker-set-up-examples/data-center.md).

&#x20;in the Docker setup instructions.

Note: for `bitbucket` and `bitbucketApi` values do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=bitbucket-server \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set bitbucketUsername=<ENTER_USERNAME> \
             --set bitbucketPassword=<ENTER_PASSWORD> \
             --set bitbucket=<ENTER_BITBUCKET_URL> \
             --set bitbucketApi=<ENTER_BITBUCKET_API_URL> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

## Gitlab Helm install

For more details including definitions of the environment variables see [Set up Snyk Broker with GitLab](snyk-broker-set-up-examples/setup-broker-with-gitlab.md).

Note: for `gitlab` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=gitlab \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set gitlab=<ENTER_GITLAB_URL> \
             --set scmToken=<ENTER_GITLAB_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

## Azure Repos Helm install

For more details including definitions of the environment variables see [Set up Snyk Broker with Azure Repos](snyk-broker-set-up-examples/setup-broker-with-azure-repos.md).

Note: for `azureReposHost` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=azure-repos \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set azureReposToken=<ENTER_REPO_TOKEN> \
             --set azureReposOrg=<ENTER_REPO_ORG> \
             --set azureReposHost=<ENTER_REPO_HOST> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

## Artifactory Repository Helm install

For more details including definitions of the environment variables see [Set up Snyk Broker with Artifactory Repository](snyk-broker-set-up-examples/set-up-snyk-broker-with-artifactory-repository.md).

Note: for `artifactoryUrl` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=artifactory \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set artifactoryUrl=<ENTER_ARTIFACTORY_URL> \
             -n snyk-broker --create-namespace
```

## Jira notifications Helm install

For more details including definitions of the environment variables see [Set up Snyk Broker with Jira](snyk-broker-set-up-examples/setup-broker-with-jira.md).

Note: for `jiraHostname` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=jira \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set jiraUsername=<ENTER_JIRA_USERNAME> \
             --set jiraPassword=<ENTER_JIRA_PASSWORD>  \
             --set jiraHostname=<ENTER_JIRA_HOSTNAME>  \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```
