# Container Registry Agent advanced configuration

## Container Registry Agent server **HTTPS configuration**

The Container Registry Agent (CRA) runs an HTTP server by default. CRA can be configured to run an HTTPS server for local connections. This requires providing an SSL certificate and a private key to the docker container at runtime.

For example, if your certificate files are found locally at `./private/container-registry-agent.crt` and `./private/container-registry-agent.key`, provide these files to the docker container by mounting the folder and using the `HTTPS_CERT` and `HTTPS_KEY` environment variables:

```
docker run --restart=always \
       -p 8081:8081 \
       -e SNYK_PORT=8081 \
       -e HTTPS_CERT=/private/container-registry-agent.crt \
       -e HTTPS_KEY=/private/container-registry-agent.key \
       snyk/container-registry-agent:latest
```

## **Container Registry and Broker client with an internal certificate**

By default, the Container Registry Agent establishes HTTPS connections to the Container Registry and Broker client. If your Container Registry or Broker client is serving an internal certificate (signed by your own CA), you can provide the CA certificate to the Container Registry Agent. For example, if your CA certificate is at `./private/ca.cert.pem`, provide it to the docker container by mounting the folder and using the `NODE_EXTRA_CA_CERTS` environment variable:

```
docker run --restart=always
-p 8081:8081
-e SNYK_PORT=8081
-e NODE_EXTRA_CA_CERTS=/private/ca.cert.pem
snyk/container-registry-agent:latest
```
