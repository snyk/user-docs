# Create a Snyk App using the Snyk CLI

{% hint style="warning" %}
All `apps` subcommands are accessible only behind the `--experimental` flag and the behavior can change at any time, without prior notice. Use all the commands with caution
{% endhint %}

You can use the Snyk CLI to create [Snyk Apps](./) by running `snyk apps create`.

You can pass the Snyk App-related data either by using the options to pass it to the Snyk CLI or by using the `--interactive` mode, for example `snyk apps --experimental --interactive`.

All Snyk Apps related subcommands are grouped under the top-level `apps` command, for example, `snyk apps create`.

## Sub-commands of snyk apps

To learn about all the available subcommands under `snyk apps` command, use the `--help` option, `snyk apps --help.`

## Options for snyk apps

`--interactive`

Use the command in interactive mode.

`--org=<ORG_ID>`

Specify the `<ORG_ID>` under which to create the Snyk App. Required for the `create` command.

`--name=<SNYK_APP_NAME>`

The name of the Snyk App that to be displayed to the user during the authentication flow. Required for the `create` command.

`--redirect-uris=<REDIRECT_URIS>`

A comma-separated list of redirect URIs. This forms a list of allowed redirect URIs to call \_\_ back after authentication. Required for the `create` command.

`--scopes=<SCOPES>`

A comma-separated list of scopes required by your Snyk App. This forms a list of scopes that your app is allowed to request during authorization. Required for the `create` command.

`--context=<CONTEXT>`

The context your Snyk App will install using. Can be either 'tenant' or 'user', will default to 'tenant' if not specified. A Snyk App using the 'tenant' context will act as a bot user so it is not tied to any particular user and hence will persist even if the installing user leaves an organisation. In contrast a Snyk App under the 'user' context will perform actions as the installing user, this should only be used if your Snyk App is performing operations that are specific to individual users. If in doubt use 'tenant'.

## Examples

Create a Snyk App

`snyk apps create --experimental --org=48ebb069-472f-40f4-b5bf-d2d103bc02d4 --name='My Awesome App' --redirect-uris=https://example1.com,https://example2.com --scopes=apps:beta`

or

`snyk apps create --experimental --interactive`
