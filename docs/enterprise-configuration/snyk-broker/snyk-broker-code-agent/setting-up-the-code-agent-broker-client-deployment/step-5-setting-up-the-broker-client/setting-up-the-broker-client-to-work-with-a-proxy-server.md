# Setting up the Broker Client to work with a proxy server

{% hint style="info" %}
These instructions also apply to setup of the Broker Client with code snippets.
{% endhint %}

To use the Code Agent - Broker Client deployment in an infrastructure that uses a proxy, verify that the requests to the Code Agent are NOT sent through the proxy, but bypass it.

For the Broker Client component, to the docker run command add the following environment variable and a command to bypass the Code Agent:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.>
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
-e NO_PROXY=<code_agent_container_name>
```

If your proxy requires username and password authentication, add also the following environment variable:

```
-e PROXY_AUTH=userID:userPass
```

In addition, add these environment variables to the Code Agent components. See [Setting up the Code Agent to work with a Proxy Server](../step-4-setting-up-the-code-agent/setting-up-the-code-agent-to-work-with-a-proxy-server.md).

For more information on using Docker containers with a proxy, see [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/).
