# Container Registry Agent advanced configuration using Docker

For instructions on installation of the Broker Container Registry Agent using Docker, see [Snyk Broker - Container Registry Agent](./). See [Install Broker for Container Registry using Helm](install-broker-for-container-registry-agent-using-helm.md) for instructions on the Helm method.

## Container Registry Agent server **HTTPS configuration**

The Container Registry Agent (CRA) runs an HTTP server by default. CRA can be configured to run an HTTPS server for local connections. This requires providing an SSL certificate and a private key to the Docker container at runtime.

For example, if your certificate files are found locally at `./private/container-registry-agent.crt` and `./private/container-registry-agent.key`, provide these files to the Docker container by mounting the folder and using the `HTTPS_CERT` and `HTTPS_KEY` environment variables:

```
docker run --restart=always \
       -p 8081:8081 \
       -e SNYK_PORT=8081 \
       -e HTTPS_CERT=/private/container-registry-agent.crt \
       -e HTTPS_KEY=/private/container-registry-agent.key \
       snyk/container-registry-agent:latest
```

## **Container Registry and Broker Client with an internal certificate**

By default, the Container Registry Agent establishes HTTPS connections to the Container Registry and Broker Client. If your Container Registry or Broker Client is serving an internal certificate (signed by your own CA), you can provide the CA certificate to the Container Registry Agent. For example, if your CA certificate is at `./private/ca.cert.pem`, provide it to the Docker container by mounting the folder and using the `NODE_EXTRA_CA_CERTS` environment variable:

<pre><code>docker run --restart=always \
       -p 8081:8081 \
<strong>       -e SNYK_PORT=8081 \
</strong>       -e NODE_EXTRA_CA_CERTS=/private/ca.cert.pem \
       snyk/container-registry-agent:latest
</code></pre>

## **Disable repository listing**

Some container registries do not support the catalog endpoint (`GET /v2/_catalog`) that lists repositories, or your organization may have permission restrictions on this endpoint. You can disable listing repos by setting the `SNYK_DISABLE_LIST_REPOS` environment variable.

When enabled, instead of calling the registry's catalog endpoint, an empty list is returned. This is useful for:

* Container registries that do not support the catalog endpoint (for example, GitHub Container Registry, GitLab Container Registry)
* Organizations with permission restrictions on the catalog endpoint
* Reducing unnecessary API calls to the container registry.

### Docker deployment

Add the environment variable to your Docker run command:

<pre class="language-markup"><code class="lang-markup">docker run --restart=always \
       -p 8081:8081 \
       -e SNYK_PORT=8081 \
<strong>       -e SNYK_DISABLE_LIST_REPOS=true \
</strong>       snyk/container-registry-agent:latest
</code></pre>

### Helm deployment

Add the following to your Helm values file:

```yaml
env:
  - name: SNYK_DISABLE_LIST_REPOS
    value: "true"
```

### Kubernetes deployment

Add the environment variable to your deployment specification:

```yaml
spec:
  containers:
    - name: container-registry-agent
      env:
        - name: SNYK_DISABLE_LIST_REPOS
          value: "true"
```
