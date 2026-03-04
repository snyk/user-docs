# HTTPS for Broker Client with Docker

{% hint style="info" %}
The HTTPS configuration for the Broker client is identical for GitHub, GitLab, and other SCM integrations. Only the Broker image and SCM-specific environment variables differ. \
\
HTTPS is optional and typically required only when your SCM provider or internal security policies require TLS for webhook or local connections.
{% endhint %}

The Broker Client runs an HTTP server by default. It can be configured to run an HTTPS server for local connections. This requires an SSL certificate and a private key to be provided to the Docker container at runtime.

For example, if your certificate files are found locally at `./private/broker.crt` and `./private/broker.key`, provide these files to the Docker container by mounting the folder and using the `HTTPS_CERT` and `HTTPS_KEY` environment variables:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e HTTPS_CERT=/private/broker.crt \
           -e HTTPS_KEY=/private/broker.key \
           -e BROKER_CLIENT_URL=https://my.broker.client:8000 \
           -v /local/path/to/private:/private \
       snyk/broker:github-com
```

{% hint style="info" %}
**This example uses the GitHub Broker image.**\
If you are using GitLab, Bitbucket, or another integration, replace the Broker image and the SCM token environment variable.\
The HTTPS configuration itself is identical across all Broker types.
{% endhint %}

`BROKER_CLIENT_URL` now has the HTTPS scheme.
