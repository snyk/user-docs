# Create a Snyk App using the Snyk CLI

{% hint style="warning" %}
All `apps` sub-commands are only accessible behind the `--experimental` flag and the behavior can change at any time, without prior notice. Use all the commands with caution
{% endhint %}

You can use the Snyk CLI to create [Snyk Apps](../features/integrations/snyk-apps/) by running `snyk apps create`.

You can pass the Snyk App related data either by using the options to pass it to the Snyk CLI or by using the `--interactive` mode, for example `snyk apps --experimental --interactive`.

All Snyk Apps related sub-commands are grouped under the top-level `apps` command, for example, `snyk apps create`.

## Sub-commands of snyk apps

To learn about all the available sub-commands under `snyk apps` command, use the `--help` option, for example, `snyk apps --help.`

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

## Examples

Create a Snyk App

`snyk apps create --experimental --org=48ebb069-472f-40f4-b5bf-d2d103bc02d4 --name='My Awesome App' --redirect-uris=https://example1.com,https://example2.com --scopes=apps:beta`

or

`snyk apps create --experimental --interactive`
