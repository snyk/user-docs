# Setting up the Code Agent to work with a Proxy Server

To use the Code Agent - Broker Client deployment in an infrastructure that uses a proxy,  add the following environment variables to the `docker run` command of the Code Agent:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.>
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
```

In addition, you will need to add these environment variables to the Broker Client component and a command to bypass the Code Agent, as described on [Setting up the Broker Client to work with a Proxy server](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2a-running-the-broker-client-without-the-code-snippet-display/setting-up-the-broker-client-to-work-with-a-proxy-server).

**Note**: For more information on using Docker containers with a proxy, see [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/).

&#x20;

To disable a certificate verification, for example, in the case of self-signed certificates, add the following parameter to the Code Agent `docker run` command:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```

&#x20;
