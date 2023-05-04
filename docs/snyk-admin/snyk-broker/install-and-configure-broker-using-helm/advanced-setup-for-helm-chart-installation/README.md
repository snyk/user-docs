# Advanced setup for Helm Chart installation

When you set up Snyk Broker using Helm, you can set advanced parameters as explained here.

## Parameters for troubleshooting and providing your own certificate

For troubleshooting SSL inspection issues, you can set the `tlsRejectUnauthorized` parameter to `disable`.

```
--set tlsRejectUnauthorized=disable
```

To provide your own certificate (signed by your own CA), you can pass the file name to the `caCert` parameter. The file must reside within the Helm chart directory.

```
--set caCert=<CERT_NAME)>
```

If you would like yourBrroker to run as an HTTPS server, you can pass the files to the `httpsCert` and `httpsKey` parameters. The files must reside within the Helm chart directory.

```
--set httpsCert=<CERT_NAME> --set httpsKey=<CERT_KEY>
```

## Proxy settings for Broker Helm chart installation

To use the Helm chart behind a proxy, set the `httpProxy` and `httpsProxy` values.

```
--set httpsProxy=<PROXY_URL>
```

## Multi-tenant settings for Helm chart installation

### **Broker multi-tenant settings**

To use the Helm chart with different multi-tenant environments, set the `brokerServerUrl` to be one of the following URLs, depending which environment you are using:

Europe: `https://broker.eu.snyk.io`\
Australia: `https://broker.au.snyk.io`

```
--set brokerServerUrl=<BROKER_SERVER_URL>
```

### **Code Agent multi-tenant settings**

If you are using Code Agent, the `upstreamUrlCodeAgent` value must be one of the following URLs, depending which environment you are using:

Europe: `https://deeproxy.eu.snyk.io`\
Australia: `https://deeproxy.au.snyk.io`

```
--set upstreamUrlCodeAgent=<UPSTREAM_URL>
```

## Custom additional options for Broker Helm chart installation

If you need to inject additional option(s) via environment variables, use the override.yaml value file to add any additional environment variable(s) you may need.

Adding the `--values override.yaml` will load those values into your deployment. For example:

```
helm install snyk-broker-chart . \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --values override.yaml \
             -n snyk-broker --create-namespace
```

You can do the same inline without the override.yaml file if it is more convenient.

```
helm install snyk-broker-chart . \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set env[0].name=myEnvVarName \
             --set env[0].value=myEnvVarValue \
             --set env[1].name=myOtherEnvVarName \
             --set env[1].value=myOtherEnvVarValue \
             -n snyk-broker --create-namespace
```

## Credential pooling and Helm Chart installation

See [Credential pooling](../../install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker.md).
