# Broker Contexts

Use a Broker Context to override the configuration of a Broker Connection. When you apply a Context to a Broker Integration, the Integration uses the override parameters instead of the default values.

**Use case**

Use a Context when multiple Organizations share a Connection but require different endpoints or credentials. Create a Context with the necessary parameter overrides and apply it to the relevant Integrations for each Organization.

*   **Connection:** Defines how the Broker connects to a platform, for example, GitHub or Jira. A Connection includes a type and fixed parameters, such as URLs and tokens.
*   **Context:** Associates with a single Connection and contains a subset of parameters. These parameters override the corresponding Connection parameters.
*   **Integration:** Represents the use of a Connection by an Organization. An Integration consists of one Organization and one Connection. You can apply only one Context per Integration.

## Configuring Broker Contexts
**Prerequisites**

*   Install `snyk-broker-config` and configure your local environment. For more details, see [Prerequisites for Universal Broker](prerequisites-for-universal-broker.md).
*   Configure and integrate at least one Broker Connection.

### Create a Context

Create a new Broker Context. Select the Connection and parameters to override, enter the new values or credential references for sensitive fields.

```bash
snyk-broker-config workflows contexts create
```

When prompted, follow these steps:
1. Select a connection.
2. Select the parameters to override.
   Available parameters depend on the connection type. For example, for Jira, you can override `jira_hostname`, `jira_username`, `jira_password`, and `jira_pat`.
3. Enter a value for each parameter.
   For sensitive parameters, the API requires a credential reference from the same deployment. Select or create a new credential reference.

### Apply a Context to an integration

Apply context overrides to an integration. Select the context and the integration to override.

```bash
snyk-broker-config workflows contexts apply
```

- Snyk displays integrations only for the same Connection as the Context.
- If you have only one Context, Snyk does not prompt you to select a Context.

### Withdraw a Context from an integration

Stop an integration from using a context. The integration reverts to the default connection parameters. Select the context, then select one or more integrations to remove.

```bash
snyk-broker-config workflows contexts withdraw
```

### Delete a Context

```bash
snyk-broker-config workflows contexts delete
```
{% hint style="info" %}
You cannot delete a Context that is applied to an integration. First, withdraw the Context from all integrations it has been applied to before attempting to delete it.
{% endhint %}

## Viewing context details

### List all Contexts for a deployment

```bash
snyk-broker-config workflows contexts list
```
### Get a single Context

```bash
snyk-broker-config workflows contexts get
```

### See which integrations use a specific Context

Displays the integration and Organization IDs associated with this Context.

```bash
snyk-broker-config workflows contexts usage
```
