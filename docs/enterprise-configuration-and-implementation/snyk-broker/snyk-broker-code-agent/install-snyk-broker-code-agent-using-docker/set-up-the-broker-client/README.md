# Set up the Broker client

Ensure that you do not currently have a running Broker client for the SCM you want to set up in the Snyk Organization where you want to do the setup. If there is already a running Broker Client,  remove it.

Then download or update the Snyk Broker client Docker image. The Code Agent is supported only in Broker Client version 4.108.0 and later versions. Pull the latest update.

Next, run the Docker container. You can run the Broker Client [with](run-the-broker-client-with-the-code-snippets-display.md) or [without](run-the-broker-client-without-the-code-snippet-display.md) a display of snippets of the Snyk Code results in the Web UI.&#x20;

You can **set up the Broker client to work with a proxy server**. To use the Code Agent in an infrastructure that uses a proxy, verify that the requests to the Code Agent are NOT sent through the proxy, but bypass it.

Add the following environment variable the `docker run` command for the Broker client and a command to bypass the proxy

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.>
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
-e NO_PROXY=<code_agent_container_name>
```

If your proxy requires username and password authentication, also the following environment variable:

```
-e PROXY_AUTH=userID:userPass
```

In addition, add these environment variables to the Code Agent components. See [Set up the Code Agent to work with a proxy server](../step-4-setting-up-the-code-agent.md#set-up-the-code-agent-to-work-with-a-proxy-server).

For more information on using Docker containers with a proxy, see [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/).
