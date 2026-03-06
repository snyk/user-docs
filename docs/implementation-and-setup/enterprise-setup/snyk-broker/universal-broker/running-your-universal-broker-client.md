# Running your Universal Broker client

Prerequisites:
Ensure you have the following environment variables set before running the Broker Client:
  - DEPLOYMENT_ID, CLIENT_ID, CLIENT_SECRET
  - Any integration credentials required by your connections (e.g. GITHUB_TOKEN)

If references are missing, the connection will not be established, and an error entry will be logged in the Broker client logs.


Run your Broker deployment on your container engine (see example Docker Compose file below) or Kubernetes cluster.

If you are not using broker.snyk.io, target the Broker server for your region by using the command `-e BROKER_SERVER_URL=https://broker.region.snyk.io \` . For details, see [Broker URLs](../../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).

## Docker Compose

### Usage 

1. Copy this file to docker-compose.yml
2. Create a .env file with required and optional variables:

  DEPLOYMENT_ID=<your-deployment-id>

  CLIENT_ID=<your-client-id>

  CLIENT_SECRET=<your-client-secret>

  PORT=8000

  Add any credentials your integrations need, for example:

  GITHUB_TOKEN=<secret>

  Optional: override for EU or other environments

  BROKER_SERVER_URL=https://broker.eu.snyk.io

  BROKER_DISPATCHER_BASE_URL=https://api.eu.snyk.io

3. Run: docker compose up -d

### Example Docker Compose file

```
services:
  snyk-broker-universal-1:
    image: snyk/broker:universal
    environment:
      DEPLOYMENT_ID: ${DEPLOYMENT_ID}
      CLIENT_ID: ${CLIENT_ID}
      CLIENT_SECRET: ${CLIENT_SECRET}
      PORT: ${PORT:-8000}
      BROKER_HA_MODE_ENABLED: "true"
      BROKER_SERVER_URL: ${BROKER_SERVER_URL:-https://broker.snyk.io}
      BROKER_DISPATCHER_BASE_URL: ${BROKER_DISPATCHER_BASE_URL:-https://api.snyk.io}
      GITHUB_TOKEN: ${MY_GH_TOKEN}
      # Pass through any integration credentials (same as -e KEY=value in docker run)
      # Example: GITHUB_TOKEN, BROKER_CLIENT_VALIDATION_AUTH_HEADER, etc.
    env_file:
      - .env
    ports:
      - "${EXTERNAL_PORT_1:-8000}:${PORT:-8000}"
    restart: unless-stopped

  snyk-broker-universal-2:
    image: snyk/broker:universal
    environment:
      DEPLOYMENT_ID: ${DEPLOYMENT_ID}
      CLIENT_ID: ${CLIENT_ID}
      CLIENT_SECRET: ${CLIENT_SECRET}
      PORT: ${PORT:-8000}
      BROKER_HA_MODE_ENABLED: "true"
      BROKER_SERVER_URL: ${BROKER_SERVER_URL:-https://broker.snyk.io}
      BROKER_DISPATCHER_BASE_URL: ${BROKER_DISPATCHER_BASE_URL:-https://api.snyk.io}
      GITHUB_TOKEN: ${MY_GH_TOKEN}
    env_file:
      - .env
    ports:
      - "${EXTERNAL_PORT_2:-8001}:${PORT:-8000}"
    restart: unless-stopped
```

## Helm

A [Helm chart](https://github.com/snyk/snyk-universal-broker-helm) is available for use on Kubernetes clusters. Refer to the readme for details.

Ensure that you first pull the Helm chart:

`helm pull oci://registry-1.docker.io/snyk/snyk-universal-broker`

Then run:

```
helm install my-snyk-broker oci://registry-1.docker.io/snyk/snyk-universal-broker \
  --set deploymentId='YOUR_DEPLOYMENT_ID' \
  --set clientId='YOUR_CLIENT_ID' \
  --set clientSecret='YOUR_CLIENT_SECRET' \
  --set credentialReferences.MY_GITHUB_TOKEN='YOUR_GITHUB_PAT' \
```

#### Secret Values

Integration types may require different SCM-specific authentication parameters. When setting your credential reference environment variable in your Broker deployment you may wish to consult the following list of secret values and their required formats:

| Integration Type   | Parameter Name                             | Format                                                           |
| ------------------ | ------------------------------------------ | ---------------------------------------------------------------- |
| Artifactory        | Artifactory URL                            | `<username>:<password>@<yourdomain.artifactory.com>/artifactory` |
| Azure Repos        | Azure Repos Token                          | Azure Repos PAT                                                  |
| Bitbucket Server   | Bitbucket Password                         | Alphanumeric password                                            |
|                    | Bitbucket PAT                              | Bitbucket Personal Access Token                                  |
| Container Registry | Azure CR Password                          | Alphanumeric password                                            |
|                    | Artifactory CR Password (ACR)              | Alphanumeric password                                            |
|                    | Docker Hub Password                        | Alphanumeric password                                            |
|                    | DigitalOcean CR Token                      | DigitalOcean PAT                                                 |
|                    | Amazon Elastic Container CR Role ARN (ECR) | AWS IAM role ARN                                                 |
|                    | Google Cloud Container (GCR) Password      | Alphanumeric password                                            |
|                    | GitHub CR Password                         | Alphanumeric password                                            |
|                    | GitLab CR Password                         | Alphanumeric password                                            |
|                    | Google Artifact CR Password                | Alphanumeric password                                            |
|                    | Harbor CR Password                         | Alphanumeric password                                            |
|                    | Nexus CR Password                          | Alphanumeric password                                            |
|                    | Quay CR Password                           | Alphanumeric password                                            |
| Github             | Github Token                               | Github Personal Access Token                                     |
| Github Enterprise  | Github Token                               | Github Personal Access Token                                     |
| Github Server App  | Github App Client ID                       | Github App Client ID                                             |
| Github Cloud App   | Github App Client ID                       | Github App Client ID                                             |
| Gitlab             | Gitlab Token                               | Gitlab Personal Access Token                                     |
| Jira               | Jira Password                              | Alphanumeric password                                            |
|                    | Jira PAT                                   | JIRA Personal Access Token                                       |
| Nexus              | Nexus Base URL                             | `https://<username>:<password>@<your.nexus.hostname>`            |
