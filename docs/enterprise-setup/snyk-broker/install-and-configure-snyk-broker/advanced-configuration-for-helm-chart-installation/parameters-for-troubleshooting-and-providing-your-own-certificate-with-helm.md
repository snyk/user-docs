# Parameters for troubleshooting and providing your own certificate with Helm

For troubleshooting SSL inspection issues, you can set the `tlsRejectUnauthorized` parameter to `disable`.

```
--set tlsRejectUnauthorized=disable
```

To provide your own certificate (signed by your own CA), you can pass the file name to the `caCert` parameter. The file must reside within the Helm chart directory.

```
--set caCert=<CERT_NAME)>
```

If you would like your Brroker to run as an HTTPS server, you can pass the files to the `httpsCert` and `httpsKey` parameters. The files must reside within the Helm chart directory.

```
--set httpsCert=<CERT_NAME> --set httpsKey=<CERT_KEY>
```

