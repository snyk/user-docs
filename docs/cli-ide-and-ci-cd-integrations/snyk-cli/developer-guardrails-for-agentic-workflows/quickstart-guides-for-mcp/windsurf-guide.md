# Windsurf guide

MCP setup using the `mcp_config.json` file can be used to set up Windsurf's MCP, as one example.&#x20;

These examples show how to add the Snyk MCP server in the `mcp_config.json` file for each transport type.

This method requires that the Snyk CLI be in your system path and can be invoked with the command `snyk`. If the CLI is not in your system path, you can provide the full path to the CLI.

## Add the Snyk MCP server using `stdio` transport in your `mcp_config.json` file

```
{
  "mcpServers": {
    "Snyk Security Scanner": {
      "command": "/absolute/path/to/snyk",
      "args": [
        "mcp",
        "-t",
        "stdio"
      ],
      "env":{
      // optional CLI environment variables, e.g. SNYK_CFG_ORG, SNYK_TOKEN
      }
    }
  }
}
```

## Add the Snyk MCP server using `sse` transport in your `mcp_config.json` file

{% hint style="info" %}
SSE transport supports running the MCP server locally only. SSE does not support remote or hosted configurations.
{% endhint %}

If your MCP Client expects a URL, then you will need to start the MCP server in your terminal first by running `snyk mcp -t sse`&#x20;

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
