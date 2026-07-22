# Evo chat

## What is Evo chat?

Evo chat is a conversational interface for working with the Evo platform. Use it to explore your inventory, create and manage policies, understand risk, and generate reports using natural language. You can refine results with follow-up questions, and ask about what you see in the Inventory, Policies, Issues, and Reports views.

Evo chat asks for your approval before it creates or updates a policy. It does not delete policies.

## Use cases

{% hint style="warning" %}
Evo chat does not support:

* Deleting a policy (components are read-only)
* Reporting using email
{% endhint %}

The following examples show prompts you can use with Evo chat to perform tasks. Evo chat provides guidance if you attempt an unsupported operation.

### Discovery agent

Use the Discovery agent to find items in your inventory. Example prompts include:

* "Show all AI models in my repositories."
* "What AI components are in the \[repo-name] repository?"
* "List all datasets in my Organization."
* "List datasets scanned in the last 24 hours."
* "What components failed their last scan?"

### Policy agent

Using the Policy agent, Evo chat updates policies and displays them in the UI for your approval. The agent does not automatically create or update policies; it requires human approval.

Example prompts include:

* "Create a policy to block dataset \[dataset-name]."
* "Add a policy to disallow \[name] licensed datasets."
* "Block all models from \[name]."
* "Disallow datasets with license \[license-name]."

### Policy agent informed by Risk Intelligence agent

The Policy agent is plugged in to the Risk Intelligence agent to create more elaborate policies.

* "Create a high-severity policy for \[ECI-score] greater or equal than \[score]."
* "Create a policy that detects all issues with model\_bias\_score higher than 500."
* "Create a policy that detects all datasets that have data leakage score higher than 300."
* "Create a policy that detects all models that have ECI score higher than 300 and license \[license-name]."

### Policy compliance and issues

Example prompts:

* "Show all policy violations."
* "What are my critical severity issues?"
* "List high-severity AI component issues."
* "How many policy violations do I have?"
* "Show issues by severity."

## Follow-up interactions

Evo chat supports iterative narrowing. Start with a broad search and filter the results using follow-up prompts, as shown in the following examples.

Evo chat is also UI-aware. You can ask it about information in the **Inventory**, **Policies**, **Issues**, and **Reports** views.

### Iterative filtering

Chat example:

1. User: "Show me all AI components."
2. Evo chat: Returns 500 components.
3. User: "Filter those to just models."
4. Evo chat: Returns 50 models.
5. User: "Show only \[name] models."
6. Evo chat: Returns 3 \[name] models.

### Progressive analysis

Chat example:

1. User: "How many policy violations do I have?"
2. Evo chat: "You have 252 violations."
3. User: "Group them by severity."
4. Evo chat: Displays 31 critical, 109 high, and 112 medium issues.
5. User: "Show me the critical ones."
6. Evo chat: Displays 31 critical issues.

## Guidelines for interacting with Evo chat

Use these guidelines to optimize chat interactions:

* Keep conversations short: Long conversations increase the context window size.
* Be specific: Provide detailed prompts to receive accurate information.
* Start new topics: To change the topic, click **+** to start a new chat.
* Refresh the window: If the chat returns unexpected output, refresh the window.
* Expect non-deterministic results: Results and wording differ between sessions.
* Use natural language: For example, use prompts like "Which datasets have the \[name] license?" or "List models scanned in the last 24 hours."
* Specify sorting: For example, use prompts like "Show the top ten most common AI components" or "List issues by highest severity first."
* Define aggregations: For example, use prompts like "Group issues by repository" or "Count models by vendor."
* Filter issues by type: Use these exact terms:
  * agent, application, dataset, http-frontend, http-server, identity-provider, messaging, model, nosql-database, object-storage, package, prompt, repo, secrets-manager, service, sql-database, tool
  * mcp-client, mcp-server, mcp-resource
