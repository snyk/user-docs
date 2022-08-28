# Setting up the Broker Client to work with a Proxy server

**Note**: These instructions are also applicable to the setup of the Broker Client with code snippets.

To use the Code Agent - Broker Client deployment in an infrastructure that uses a proxy, you need to verify that the requests to the Code Agent are NOT sent through the proxy, but bypassing it.

For the Broker Client component, add to the docker run command the following environment variables and a command to bypass the Code Agent:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.>
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
-e NO_PROXY=<code_agent_container_name>
```

In addition, you will need to add these environment variables to the Code Agent components, as described on page 32.

**Note**: For more information on using Docker containers with a proxy, see [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/).

