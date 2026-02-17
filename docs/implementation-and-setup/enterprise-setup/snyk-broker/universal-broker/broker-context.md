# Broker Contexts

Use a Broker Context to override the configuration of a Broker Connection. When you apply a Context to a Broker Integration, the Integration uses the override parameters instead of the default values.

**Use case**

Use a Context when multiple Organizations share a Connection but require different endpoints or credentials. Create a Context with the necessary parameter overrides and apply it to the relevant Integrations for each Organization.

*   **Connection:** Defines how the Broker connects to a platform, for example, GitHub or Jira. A Connection includes a type and fixed parameters, such as URLs and tokens.
*   **Context:** Associates with a single Connection and contains a subset of parameters. These parameters override the corresponding Connection parameters.
*   **Integration:** Represents the use of a Connection by an Organization. An Integration consists of one Organization and one Connection. You can apply only one Context per Integration.

## Configuring Broker Contexts

{% hint style="info" %}
Before starting ensure that you have installed the snyk-broker-config and have configured your local environment correctly. Please see [Prerequisites for Universal Broker](prerequisites-for-universal-broker.md) for more info. You must also have at least one Broker Connection configured and integrated.
{% endhint %}

### Create a Context

Creates a new Broker Context. You choose the Connection, then which parameters to override, then provide the values (or credential references for sensitive fields) that you wish to override.

```bash
snyk-broker-config workflows contexts create
```

- You will be asked to pick a Connection.
- You then choose which parameters to override. Allowed parameters depend on the connection type (e.g. for jira: `jira_hostname`, `jira_username`, `jira_password` and `jira_pat` can all be overridden).
- For each chosen parameter you enter a value. For sensitive parameters the API expects a credential reference from the same deployment, you will be prompted to choose or create a new credential reference.

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
- Note that you cannot delete a Context that has been applied to an Integration. Withdraw the Context from any Integrations it has been applied to before deletion.

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
