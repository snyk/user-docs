# Install Broker for SCMs using Helm

After you add the [Snyk Broker Helm Char](how-to-install-and-configure-your-snyk-broker-client/)t, run the following commands based on the repository type.

The following are the allowed values for `scmType`:

Github.com: `github-com`\
Github Enterprise: `github-enterprise`\
Bitbucket: `bitbucket-server`\
Gitlab: `gitlab`\
Azure Repos: `azure-repos`\
Artifactory: `artifactory`\
Jira: `jira`\
Container Registry Agent: `container-registry-agent`\


The following examples create a namespace called `snyk-broker`. To deploy into an existing namespace, adjust the `-n` parameter and delete the `--create-namespace` parameter.

## Github.com helm install

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

## Github Enterprise helm install

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

## Bitbucket helm install

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

## Gitlab helm install

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

## Azure Repos helm install

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

## Artifactory Respository helm install

Note: for `artifactoryUrl` value do not include `https://`

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=artifactory \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set artifactoryUrl=<ENTER_ARTIFACTORY_URL> \
             -n snyk-broker --create-namespace
```

## Jira notifications helm install

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

