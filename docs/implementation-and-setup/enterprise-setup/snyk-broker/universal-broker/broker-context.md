# Broker Contexts

A Broker Context is an optional configuration override for Broker Connections that you can apply to a Broker Integration. When a Context is applied, that Integration begins using the override's parameters instead of the Connection's default values.

Typical use case: A Connection is shared by many Orgs, with some Orgs using a different endpoint or credentials. You create a Context with the necessary parameters overrides, then apply that Context to the relevant Integrations for each Org.

- Connection: Defines how the Broker connects to an SCM (e.g. GitHub, Jira, etc). It has a type and a fixed set of parameters (URLs, tokens, etc.).
- Context: Belongs to one Connection and holds a overriding subset of that Connection’s parameters. Each key in the Context overrides the same key on the Connection when the Context is applied.
- Integration: An Organization’s use of a Connection (one Org + one Connection = one Integration). You can apply at most one Context per Integration.

## Configuring Broker Contexts

{% hint style="info" %}
Before starting ensure that you have installed the snyk-broker-config have configured your local set up correctly. Please see [Prerequisites for Universal Broker](prerequisites-for-universal-broker.md) for more info. You must also have at least one Broker Connection configured and integrated.
{% endhint %}

### Create a Context

Creates a new Broker Context. You choose the Connection, then which parameters to override, then provide the values (or credential references for sensitive fields) that you wish to override.

```bash
snyk-broker-config workflows contexts create
```

- You will be asked to pick a Connection.
- You then choose which parameters to override. Allowed parameters depend on the connection type (e.g. for jira: `jira_hostname`, `jira_username`, `jira_password` and `jira_pat` can all be overridden).
- For each chosen parameter you enter a value. For sensitive parameters the API expects a credential reference from the same deployment, you will be prompted to choose or create a new credential reference.

### Apply a Context to an Integration

Applies a Context's overrides to an Integration. You select the Context, then the Integration you want to override.

```bash
snyk-broker-config workflows contexts apply
```

- Only integrations for the same Connection as the Context are offered. 
- If you only have one Context then you won't be prompted to choose the Context.

### Withdraw a Context from an Integration

Stops an Integration from using the Context, the Integration falls back to the Connection’s default parameters. You select the Context, then one or more Integrations to withdraw from.

```bash
snyk-broker-config workflows contexts withdraw
```

### Delete a Context

```bash
snyk-broker-config workflows contexts delete
```
- Note that you cannot delete a Context that has been applied to an Integration. Withdraw the Context from any Integrations it has been applied to before deletion.

## Viewing Context Details

### List all Contexts for a Deployment

```bash
snyk-broker-config workflows contexts list
```
### Get a single Context

```bash
snyk-broker-config workflows contexts get
```

### See which Integrations use a specific Context

Shows which Integration and Org IDs currently have this Context applied.

```bash
snyk-broker-config workflows contexts usage
```
