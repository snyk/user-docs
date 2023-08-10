# Parameters for troubleshooting and providing your own certificate with Helm

For troubleshooting SSL inspection issues, you can set the `tlsRejectUnauthorized` parameter to `disable`.

```
--set tlsRejectUnauthorized=disable
```

To provide your own certificate (signed by your own CA), you can pass the file name to the `caCert` parameter. The file must reside within the Helm chart directory.

```
--set caCert=<CERT_NAME)>
```

If you want your Broker to run as an HTTPS server, you can pass the files to the `httpsCert` and `httpsKey` parameters. The files must reside within the Helm chart directory.

```
--set httpsCert=<CERT_NAME> --set httpsKey=<CERT_KEY>
```

For more information about using your own certificate, see [Backend requests with an internal certificate for Docker](../advanced-configuration-for-snyk-broker-docker-installation/backend-requests-with-an-internal-certificate-for-docker.md) and [HTTPS for Broker Client with Docker](../advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker.md).
