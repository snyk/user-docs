# Backend requests with an internal certificate for Docker

By default, the Broker Client establishes HTTPS connections to the backend system: GitHub, BitBucket, Jira, or other. If your backend system is serving an internal certificate (signed by your own certificate authority (CA)), you can provide the CA certificate to the Broker Client.

For example, if your CA certificate is at `./private/ca.cert.pem`, provide it to the Docker container by mounting the folder and using the `NODE_EXTRA_CA_CERT` environment variable. See the following  example for Bitbucket:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BITBUCKET_USERNAME=username \
           -e BITBUCKET_PASSWORD=password \
           -e BITBUCKET=your.bitbucket-server.domain.com \
           -e BITBUCKET_API=your.bitbucket-server.domain.com/rest/api/1.0 \
           -e PORT=8000 \
           -e NODE_EXTRA_CA_CERTS=/private/ca.cert.pem \
           -v /local/path/to/private:/private \
       snyk/broker:bitbucket-server
```

Beginning with [Broker version 4.166.0 (2023-10-10)](https://github.com/snyk/broker/releases/tag/v4.166.0), the custom CA cert instruction is `NODE_EXTRA_CA_CERTS`  and this must be set as shown in order to use a custom CA. The `CA_CERT` environment variable is no longer in use for this purpose.

Note that this completely replaces the default CA Certificate List for any requests made to your backend system, so this must be the complete chain required by the certificate used by the backend system.

It must be `PEM`-formatted; `DER` is not supported. Supported certificate types are:

* `TRUSTED CERTIFICATE`
* `X509 CERTIFICATE`
* `CERTIFICATE`

An example follows.

```
-----BEGIN CERTIFICATE-----
<base64-encoded certificate>
-----END CERTIFICATE----
-----BEGIN CERTIFICATE-----
<base64-encoded certificate>
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
<base64-encoded certificate>
-----END CERTIFICATE-----
```
