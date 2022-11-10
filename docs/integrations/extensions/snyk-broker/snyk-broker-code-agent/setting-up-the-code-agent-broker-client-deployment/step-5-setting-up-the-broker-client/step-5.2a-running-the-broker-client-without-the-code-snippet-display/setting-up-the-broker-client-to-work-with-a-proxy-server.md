# Setting up the Broker Client to work with a Proxy server

**Note**: These instructions are also applicable to the setup of the Broker Client with code snippets.

To use the Code Agent - Broker Client deployment in an infrastructure that uses a proxy, you need to verify that the requests to the Code Agent are NOT sent through the proxy, but bypassing it.

For the Broker Client component, add to the docker run command the following environment variables and a command to bypass the Code Agent:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.>
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
-e NO_PROXY=<code_agent_container_name>
```

If your proxy requires username and password authentication, please add the following additional environment variable:

```
-e PROXY_AUTH=userID:userPass
```

In addition, you will need to add these environment variables to the Code Agent components. See [Setting up the Code Agent to work with a Proxy Server](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-4-setting-up-the-code-agent/setting-up-the-code-agent-to-work-with-a-proxy-server).

**Note**: For more information on using Docker containers with a proxy, see [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/).

