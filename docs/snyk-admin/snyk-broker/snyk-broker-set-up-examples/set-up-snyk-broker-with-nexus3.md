# Set up Snyk Broker with Nexus3

To use the Nexus 3 client with a Nexus 3 deployment, **run** `docker pull snyk/broker:nexus` tag. The following environment variables are needed to customize the Broker client:

* `BROKER_TOKEN` - the Snyk Boker token, obtained from your Artifactory integration settings view.
* `BASE_NEXUS_URL` - the URL of your Nexus 3 deployment, such as `https://[<user>:<pass>@]<your.nexus.hostname>`.
* `BROKER_CLIENT_VALIDATION_URL` - Nexus validation url, checked by Broker Client systemcheck endpoint. If Nexus user requires auth, use `$BASE_NEXUS_URL/service/rest/v1/status/check` (for example, `https://<user>:<pass>@<your.nexus.hostname>/service/rest/v1/status/check`) otherwise use `$BASE_NEXUS_URL/service/rest/v1/status` (for example, `https://<your.nexus.hostname>/service/rest/v1/status`).
* (Optional) `RES_BODY_URL_SUB` - This URL substitution is required for npm/Yarn integration and is the same as the URL of the Nexus without credentials appended with `/repository`, for example, `https://<your.nexus.hostname>/repository`

Use the following command to set up a fully configured Broker Client to use with Nexus3. You can run the Docker container by providing the relevant configuration:

```console
docker run --restart=always \
           -p 7341:7341 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BASE_NEXUS_URL=https://[<user>:<pass>@]<your.nexus.hostname> \
           -e BROKER_CLIENT_VALIDATION_URL=https://<your.nexus.hostname>/service/rest/v1/status[/check] \
           -e RES_BODY_URL_SUB=https://<your.nexus.hostname>/repository \
       snyk/broker:nexus
```

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Nexus3 integration.
