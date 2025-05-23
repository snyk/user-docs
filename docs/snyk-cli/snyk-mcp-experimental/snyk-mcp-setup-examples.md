# Snyk MCP setup examples

To add an MCP server to an Agentic IDE, consult the documentation for the AI system where you plan to integrate Snyk and review the specific MCP instructions. Examples of systems where you can integrate Snyk include [Windsurf's MCP](https://docs.windsurf.com/windsurf/mcp), [Qodo's MCP support](https://docs.qodo.ai/qodo-documentation/qodo-gen/qodo-gen-chat/agentic-mode/agentic-tools-mcps), and [VS Code MCP support](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).

## MCP setup examples using the `mcpconfig.json` file

This method can be used to set up Windsurf's MCP, as one example. These examples show how to add the Snyk MCP server in the `mcpconfig.json` file for each transport type. This method requires that the Snyk CLI is in your system path and can be invoked with the command `snyk`. If the CLI is not in your system path, you can provide the full path to the CLI.

### Add the Snyk MCP server using `stdio` transport in your `mcpconfig.json` file

```
{
  "mcpServers": {
    "Snyk Security Scanner": {
      "command": "/absolute/path/to/snyk",
      "args": [
        "mcp",
        "-t",
        "stdio",
        "--experimental"
      ],
      "env":{
      // optional CLI environment variables, e.g. SNYK_CFG_ORG, SNYK_TOKEN
      }
    }
  }
}
```

### Add the Snyk MCP server using `sse` transport in your `mcpconfig.json` file

If your MCP Client expects a URL, then you will need to start the MCP server in your terminal first by running `snyk mcp -t sse --experimental`&#x20;

This will output the base URL for your local SSE server. The `sse` endpoint lives on `http://baseUrl/sse.`

```
{
  "mcpServers": {
    "Snyk Security Scanner": {
      "url": "http://baseUrl/sse",
    }
  }
}
```

{% hint style="info" %}
SSE transport supports running the MCP server locally only. SSe does not support remote or hosted configurations.
{% endhint %}

## Qodo setup steps

1. Select the Agentic option for interacting with Qodo.

<div data-full-width="true"><figure><img src="../../.gitbook/assets/Screenshot 2025-04-24 at 09.56.26.png" alt="" width="563"><figcaption><p>Select Agentic</p></figcaption></figure></div>

2. Click **Connect more tools**.

<figure><img src="../../.gitbook/assets/Screenshot 2025-04-24 at 09.56.52.png" alt="" width="561"><figcaption><p>Connect more tools</p></figcaption></figure>

3. Click the **+** button to add a new MCP server.

<figure><img src="../../.gitbook/assets/Screenshot 2025-04-24 at 09.57.06.png" alt="" width="563"><figcaption><p>Plus button for Agentic Tools (MCP)</p></figcaption></figure>

4. Provide the required details for the Snyk MCP Server:

* The MCP Server name
* The path to the Snyk CLI
* The Snyk MCP command to [start the MCP server](snyk-mcp-setup-examples.md#starting-the-snyk-mcp-server).

<figure><img src="../../.gitbook/assets/Screenshot 2025-04-24 at 10.01.57.png" alt="" width="563"><figcaption><p>MCP server details</p></figcaption></figure>

5. Snyk Security should now be visible in your list of tools. You can expand the Snyk Security list to see the tools available with the Snyk MCP integration.

<figure><img src="../../.gitbook/assets/Screenshot 2025-04-24 at 10.02.14.png" alt="" width="563"><figcaption><p>Tools available with Snyk MCP integraton</p></figcaption></figure>

6. At this point, start interacting with the Snyk MCP and ask for your code to be scanned.

<figure><img src="../../.gitbook/assets/Screenshot 2025-04-24 at 10.02.59.png" alt="" width="563"><figcaption><p>Request to scan your code</p></figcaption></figure>
