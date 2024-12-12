# Examples for the Update existing integration endpoint

Examples follow for the [Update existing integration](../reference/integrations-v1.md#org-orgid-integrations-integrationid) endpoint in the [Integrations (v1)](../reference/integrations-v1.md) API.

## Set up Broker for an existing integration

### Command

```
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"github\",
    \"broker\": { \"enabled\": true }
}" \
'https://api.snyk.io/v1/org/{orgId}/integrations/{integrationId}'
```

### Response

```
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b",
  "brokerToken": "4a18d42f-0706-4ad0-b127-24078731fbed"
}
```

### Possible values for Type (integration type, enum)

`acr`, `artifactory-cr`, `azure-repos`, `bitbucket-cloud`, `bitbucket-server`, `digitalocean-cr`, `docker-hub`, `ecr,` `gcr`, `github`, `github-cr`, `github-enterprise`, `gitlab`, `gitlab-cr`, `google-artifact-cr`, `harbor-cr`, `nexus-cr`, `quay-cr`

### Credentials needed for the integration you are updating

AcrCredentials: object\
username: required. string\
password: required, string\
registryBase: required, string, for example, name.azurecr.io

OR ArtifactoryCrCredentials: object\
username: required, string\
password: required, string \
registryBase: required, string, for example, name.jfrog.io

OR AzureReposCredentials: object\
username: required, string\
url: required, string

OR BitbucketCloudCredentials: object\
username: required, string\
password: required, string

OR BitbucketServerCredentials: object\
username: required, string\
password: required, string\
url: required, string

OR DigitalOceanCrCredentials: object\
token: required, string, Personal Access Token

OR DockerHubCredentials: object\
username: required, string\
password: required, string, Access Token

OR EcrCredentials: object\
region: required, string, for example, eu-west-3\
roleArn: required, string, for example, arn:aws:iam::\<account-id>:role/\<newRole>

OR GcrCredentials: object\
password: required, string, JSON key file\
registryBase: required, string, for example, gcr.io, us.gcr.io, eu.gcr.io, asia.gcr.io

OR GitHubCredentials: object\
token: required, string

OR GitHubCrCredentials: object\
username: required, string\
password: required, string \
egistryBase: required, string, for example ghcr.io

OR GitHubEnterpriseCredentials: object\
token: required, string\
url: required, string

OR GitLabCredentials: object\
token: required, string\
url: string. for self-hosted GitLab only

OR GitLabCrCredentials: object\
username: required, string\
password: required, string\
registryBase: required, string, for example, your.gitlab.host

OR GoogleArtifactCrCredentials: object\
password: required, string, JSON key\
file registryBase: required, string, for example, us-east1-docker.pkg.dev, europe-west1-docker.pkg.dev

OR HarborCrCredentials: object\
username: required, string\
password: required, string\
registryBase: required, string, for example, your.harbor.host

OR NexusCrCredentials: object\
username: required, string\
password: required, string\
registryBase: required, string, for example, your.nexus.host

OR QuayCrCredentials: object\
username: required, string\
password: required, string\
registryBase: required, string, for example, quay.io, your.quay.host

### Snyk permissions needed

View Organization\
View Integrations\
Edit Integrations

## Update credentials for an existing non-brokered integration

### Command

```
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"gitlab\",
    \"credentials\": { \"token\": \"GITLAB_TOKEN\" }
}" \
''
```

### Response

```
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b"
}
```

### Possible values for Type (integration type, enum)

Same as [values for Set up Broker for an existing integration](examples-for-the-update-existing-integration-endpoint.md#possible-values-for-type-integration-type-enum).

### Credentials needed for the integration you are updating

Same as [credentials for Set up Broker for an existing integration](examples-for-the-update-existing-integration-endpoint.md#credentials-needed-for-the-integration-you-are-updating).

### Snyk permissions needed

View Organization\
View Integrations\
Edit Integrations

## Disable broker for an existing integration

### Command

```
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"github\",
    \"broker\": { \"enabled\": false },
    \"credentials\": { \"token\": \"GITHUB_TOKEN\" }
}" \
''
```

### Response

```
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b"
}
```
