# Anthropic's Claude Enterprise

{% hint style="info" %}
**Early Access** — This integration is available in early access. Contact your Snyk account team to enable it.
{% endhint %}

Integrate Snyk Evo with Anthropic's Claude Enterprise to ingest a snapshot of your Claude Enterprise organization. This integration surfaces AI assets, such as Claude models in use and organization-approved MCP servers, directly in your Evo inventory.

## Prerequisites

* Snyk Evo with the Anthropic Claude Enterprise integration enabled. Contact your Snyk account team to request access.
* A Claude Enterprise account.
* Primary Owner role in Claude Enterprise. Only the Primary Owner can create a Compliance Access Key.

## Create a Compliance Access Key in Claude Enterprise

1. Log in to [claude.ai](https://claude.ai) as the Primary Owner of your organization.
2. Navigate to **Settings > Data and Privacy > Compliance access keys**.
3. Click **Create key**.
4. Enter a name for the key, for example, _Snyk Evo_.
5. Select the following three scopes:
   * `read:compliance_activities`
   * `read:compliance_user_data`
   * `read:compliance_org_data`
6. Click **Create**.
7. Copy the key and store it in a secure location.

{% hint style="warning" %}
You cannot change scopes after you create a key. If you select the wrong scopes, delete the key and create a new one. Do not select `delete:compliance_user_data`. Evo does not require this scope.
{% endhint %}

{% hint style="warning" %}
Evo does not store your Compliance Access Key between sessions. Paste the key each time you run an ingestion. If you lose the key, return to Claude Enterprise and create a new one.
{% endhint %}

## Ingest a snapshot in Evo

1. In Evo, navigate to **AI platform integrations**.
2. On the **Ingest Anthropic compliance snapshot** page, paste your Compliance Access Key into the **Compliance key** field.
3. Click **Ingest snapshot**.

Evo calls Claude's Compliance API and ingests the models and MCP servers of your Claude Enterprise organization as native Evo assets. When the ingestion completes, a **Snapshot ingested** confirmation appears, and the page displays **Snapshot last ingested: Just now**.

{% hint style="info" %}
Ingestion is synchronous. For large workspaces, this process takes several seconds.
{% endhint %}

## Refresh your data

Evo does not automatically sync with Claude Enterprise. To update your inventory with the latest data, navigate to the **AI platform integrations** page, click **Ingest new snapshot**, and enter your Compliance Access Key.

## View your inventory

After a successful ingestion, the **AI platform integrations** page displays two tabs: Models and MCP servers.

### Models

The Models tab lists every Claude model detected in use across your organization.

| Column        | Description                                                          |
| ------------- | -------------------------------------------------------------------- |
| Name          | Claude model identifier                                              |
| Added         | When the model first appeared in your organization's data            |
| Projects      | Number of Claude Projects associated with this model                 |
| Organizations | Number of Claude Enterprise organizations where this model is active |

Select a model to open its detail page. The detail page shows usage statistics (Organizations, Projects, Used in chats, Users) and three tabs:

* Usage: Lists the organizations using this model, with the date added and the creator who first associated it.
* Risk profile: Evo's AI Risk DB scores the model across vulnerability categories on a 0–1000 index (lower is better):

| Category                 | Description                                                                                                                                                    |
| ------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Bias & discrimination    | Measures fairness and discriminatory behavior across demographics.                                                                                             |
| Insecure code generation | Measures how often the model generates code containing vulnerabilities such as SQL injection, XSS, OS injection, path traversal, or improper input validation. |
| Sensitive data exposure  | Tests whether the model memorizes and reproduces personally identifiable information (PII) from its training data.                                             |
| Attack reconnaissance    | Measures susceptibility to assisting with information gathering for attacks.                                                                                   |
| Safety guardrail bypass  | Tests resistance to attempts to circumvent safety controls.                                                                                                    |

* Capabilities: Displays model metadata including vendor, model family, model variant, tasks, modalities, license type, cost tier, and a link to the Anthropic model documentation.

### MCP servers

The MCP servers tab lists every MCP server added to your Claude Enterprise organization.

| Column        | Description                                                           |
| ------------- | --------------------------------------------------------------------- |
| Name          | Name of the MCP server (for example, _PagerDuty_)                     |
| Added         | Date the server was added to the Claude Enterprise organization       |
| Projects      | Number of Claude Projects associated with this server                 |
| Organizations | Number of Claude Enterprise organizations where this server is active |

Select an MCP server to open its detail page. The detail page shows usage statistics (Added, Organizations, Private projects, Public projects, Users) and two tabs:

* Usages: Lists the organizations using this MCP server, with the date added and the creator's email.
* Tools & permissions: Displays tool-level permission restrictions set by your Claude Enterprise administrators. The Tools table shows:

| Column          | Description                                   |
| --------------- | --------------------------------------------- |
| Name            | Tool name (for example, _delete\_task_)       |
| Date            | When the restriction was set                  |
| Set by          | The administrator who applied the restriction |
| Max permissions | Permission level for this tool                |

If no tool-level policies are configured in Claude Enterprise, this table shows **No data**. To set tool permissions, update your MCP server configuration in Claude Enterprise and click **Ingest new snapshot** again.

## Troubleshooting

| Issue                                     | Likely cause                                                    | Resolution                                                                                                                               |
| ----------------------------------------- | --------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| **Ingest snapshot** button is disabled    | **Compliance key** field is empty                               | Paste your Compliance Access Key into the **Compliance key** field.                                                                      |
| Ingestion succeeds but Models tab shows 0 | No chat activity in your Claude Enterprise organization         | Generate activity in Claude Enterprise, then click **Ingest new snapshot**.                                                              |
| MCP servers tab shows 0                   | No MCP servers are added to your Claude Enterprise organization | Add MCP servers in Claude Enterprise administrator settings, then click **Ingest new snapshot**.                                         |
| Ingestion fails                           | Key is invalid or revoked                                       | Check the key status in Claude Enterprise under **Settings > Data and Privacy > Compliance access keys** and create a new key if needed. |
| Ingestion fails with a scope error        | Key lacks required scopes                                       | Delete the key in Claude Enterprise and create a new one with all three required scopes.                                                 |
