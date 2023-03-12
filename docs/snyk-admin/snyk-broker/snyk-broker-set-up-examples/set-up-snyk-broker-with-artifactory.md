# Set up Snyk Broker with Artifactory

To use the Broker client with an Artifactory deployment, **run** `docker pull snyk/broker:artifactory` tag. The following environment variables are needed to customize the Broker client:

* `BROKER_TOKEN` - the Snyk Broker token, obtained from your Artifactory integration settings view.
* `ARTIFACTORY_URL` - the URL of your Artifactory deployment, such as `<yourdomain>.artifactory.com/artifactory`.

Use the following command to set up a fully configured Broker Client to use with Artifactory. You can run the Docker container by providing the relevant configuration:

```console
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e ARTIFACTORY_URL=<yourdomain>.artifactory.com/artifactory \
       snyk/broker:artifactory
```

As an alternative to using the Docker run command, you can use a derived Docker image to set up the Broker Client integration. See [Derived Docker images](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md) for the environment variables to override for the Artifactory integration.
