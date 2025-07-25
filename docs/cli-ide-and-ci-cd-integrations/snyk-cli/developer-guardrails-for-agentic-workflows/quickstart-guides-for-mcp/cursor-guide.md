# Cursor guide

You can add the Snyk MCP server to Cursor to secure code generated with agentic workflows through an LLM. This can be achieved in several ways. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Direct link (preferred)

* Click [this link](cursor://anysphere.cursor-deeplink/mcp/install?name=Snyk\&config=eyJjb21tYW5kIjoic255ayBtY3AgLXQgc3RkaW8ifQ==) to directly add the server to Cursor.
* Confirm installation in the Cursor settings by clicking on `Install`

<figure><img src="../../../../.gitbook/assets/image (44).png" alt=""><figcaption></figcaption></figure>

## Cursor MCP server directory

* Navigate to the list of MCP servers on the Cursor website and search for Snyk. Then install by clicking the `Add Snyk to Cursor` button.

<figure><img src="../../../../.gitbook/assets/image (43).png" alt=""><figcaption></figcaption></figure>

* Confirm the installation by clicking `Install` in the Cursor settings.

## Install manually in Cursor

* Open Cursor settings
* Navigate to **Tools & Integrations**
* Select **Add Custom MCP**&#x20;
* Add the following json snippet to the file in use. For this example, the assumption is that  Snyk is on the system path. If this is not applicable, then use the absolute path to the Snyk CLI.

```json5
{
  "mcpServers": {
    "Snyk": {
      "command": "snyk mcp -t stdio",
      "env": {}
    }
  }
}
```
