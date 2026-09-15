---
description: How to run the Universal Broker client, including the deployment and credential prerequisites you need
nav_context: agnostic
---

# Running your Universal Broker client

{% hint style="info" %}
Ensure you have all of the [prerequisites](prerequisites-for-universal-broker.md) before running the Broker Client:

* The DEPLOYMENT\_ID, CLIENT\_ID, CLIENT\_SECRET for your Broker Deployment
* A credential reference associated with your deployment
* Valid integration credentials required by your connections such as MY\_GITHUB\_TOKEN If references are missing, the connection will not be established, and an error entry will be logged in the Broker client logs.
{% endhint %}

Run your Broker deployment on your container engine ([Docker Compose](running-your-universal-broker-client.md#docker-compose-example) or [Kubernetes cluster](running-your-universal-broker-client.md#helm)).

If you are not using `app.snyk.io` (for example, if you log into https://app.eu.snyk.io), you will need to target the Broker server for your region by using the following in your Docker run command `-e BROKER_SERVER_URL=https://broker.region.snyk.io \` .  For details, visit [Broker URLs](https://app.gitbook.com/s/ELvljsaLKPkSpffOkmsQ/regional-hosting-and-data-residency#broker-server-urls).

## Docker Compose example

1. Create a `.env` file with required and optional configuration variables:

```bash
  DEPLOYMENT_ID=<your-deployment-id>
  CLIENT_ID=<your-client-id>
  CLIENT_SECRET=<your-client-secret>
  PORT=8000
  # Add any credentials your integrations need, for example:
  MY_GITHUB_TOKEN=<secret>
  # Optional: override for EU or other environments
  BROKER_SERVER_URL=https://broker.eu.snyk.io
  BROKER_DISPATCHER_BASE_URL=https://api.eu.snyk.io
```

2. Copy this example file to `docker-compose.yaml`

```yaml
services:
  snyk-broker-universal-1:
    image: snyk/broker:universal
    environment:
      DEPLOYMENT_ID: ${DEPLOYMENT_ID}
      CLIENT_ID: ${CLIENT_ID}
      CLIENT_SECRET: ${CLIENT_SECRET}
      PORT: ${PORT:-8000}
      BROKER_SERVER_URL: ${BROKER_SERVER_URL:-https://broker.snyk.io}
      BROKER_DISPATCHER_BASE_URL: ${BROKER_DISPATCHER_BASE_URL:-https://api.snyk.io}
      # Declare any integration credentials (same as -e KEY=value in docker run)
      # Example: GITHUB_TOKEN, BROKER_CLIENT_VALIDATION_AUTH_HEADER, etc.
      GITHUB_TOKEN: ${MY_GITHUB_TOKEN}
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
      BROKER_SERVER_URL: ${BROKER_SERVER_URL:-https://broker.snyk.io}
      BROKER_DISPATCHER_BASE_URL: ${BROKER_DISPATCHER_BASE_URL:-https://api.snyk.io}
      GITHUB_TOKEN: ${MY_GITHUB_TOKEN}
    env_file:
      - .env
    ports:
      - "${EXTERNAL_PORT_2:-8001}:${PORT:-8000}"
    restart: unless-stopped
```

3. Run `docker compose up -d` to start the containers.

## Helm

A [Helm chart](https://github.com/snyk/snyk-universal-broker-helm) is available for use on Kubernetes clusters. Refer to the README for details.

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
