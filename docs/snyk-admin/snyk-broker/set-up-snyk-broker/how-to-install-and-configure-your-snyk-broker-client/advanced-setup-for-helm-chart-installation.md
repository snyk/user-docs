# Advanced setup for Helm Chart installation

## Parameters for troubleshooting and providing your own certificate

There is also the ability to set more advanced parameters. For troubleshooting SSL inspection issues, you can set the `tlsRejectUnauthorized` parameter to `disable`.

```
--set tlsRejectUnauthorized=disable
```

To provide your own certificate (signed by your own CA) - you can pass the file name (it needs to reside within the helm chart directory) to the `caCert` parameter.

```
--set caCert=<CERT_NAME)>
```

If you would like your broker to run as an HTTPS server, you can pass the files (they need to reside within the helm chart directory) to the `httpsCert` and `httpsKey` parameters

```
--set httpsCert=<CERT_NAME> --set httpsKey=<CERT_KEY>
```

## Proxy settings for Broker Helm Chart installation

To use this chart behind a proxy, set the `httpProxy` and `httpsProxy` values.

```
--set httpsProxy=<PROXY_URL>
```

## Multi-tenant settings for Helm Chart installation

### **Broker multi-tenant settings**

To use this chart with different multi-tenant environments, set the `brokerServerUrl` to be one of the following URLs depending which environment you are using:

Europe: `https://broker.eu.snyk.io`\
Australia: `https://broker.au.snyk.io`

```
--set brokerServerUrl=<BROKER_SERVER_URL>
```

### **Code Agent multi-tenant settings**

If using Code Agent, this requires `upstreamUrlCodeAgent` value to be one of the following URLs depending which environment you are using:

Europe: `https://deeproxy.eu.snyk.io`\
Australia: `https://deeproxy.au.snyk.io`

```
--set upstreamUrlCodeAgent=<UPSTREAM_URL>
```

## Custom additional options for Broker Helm Chart installation

If you need to inject additional option(s) via environment variables, use the override.yaml value file to add any additional environment variable you may need.

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

The same can be achieved inline, without the override.yaml file if more convenient.

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
