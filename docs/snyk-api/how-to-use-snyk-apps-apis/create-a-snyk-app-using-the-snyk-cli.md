# Create a Snyk App using the Snyk CLI

{% hint style="warning" %}
Release status\
All `snyk apps` subcommands are available only behind the `--experimental` flag and the behavior can change at any time, without prior notice. Use all the commands with caution.
{% endhint %}

You can use the Snyk CLI to create Snyk Apps by running `snyk apps create`. There are two ways to use the command.

The first is the normal mode, for example:

`snyk apps create --experimental --org=48ebb069-472f-40f4-b5bf-d2d103bc02d4 --name='My Awesome App' --redirect-uris=https://example1.com,https://example2.com --scopes=apps:beta`

The second is the interactive mode, which prompts you to enter all the values in a similar way as with the normal mode. The following is an example of the interactive mode:

`snyk apps create --experimental --interactive`

```
snyk apps create --experimental --interactive

? Name of the Snyk App (visible to users when they install the Snyk App)? My Awesome Snyk App
? Your Snyk App's redirect URIs (comma-separated list.  Ex: https://example1.com,https://example2.com)?:  https://example1.com
? Your Snyk App's permission scopes (comma-separated list.  Ex: org.read)?:  apps:beta
? Please provide the org id under which you want to create your Snyk App:  48ebb069-472f-40f4-b5bf-d2d103bc02d4
```

## Options for `snyk apps create`

`--interactive`

Use the `snyk apps create` command in interactive mode.

`--org=<ORG_ID>`

Specify the `<ORG_ID>` under which to create the Snyk App. Required for the `create` command.

`--name=<SNYK_APP_NAME>`

The name to be displayed to the end-user when the user authorizes the App. Required for the `create` command.

`--redirect-uris=<REDIRECT_URIS>`

A comma-separated list of redirect URIs. This forms a list of allowed redirect URIs to call back after authentication. Required for the `create` command.

`--scopes=<SCOPES>`

A comma-separated list of scopes required by your Snyk App. This forms a list of scopes that your app is allowed to request during authorization. Required for the `create` command.

`--context=<CONTEXT>`

The `context` your Snyk App will use when installed.

Can be either `tenant` or `user`. The default is `tenant` if `context` is not specified.

A Snyk App that has the `tenant` context will act as a bot user so it is not tied to any individual user and thus will persist even if the installing user leaves the Snyk Organization. In contrast, a Snyk App that has the `user` context will perform actions as the installing user. Specify the `user` context only if your Snyk App is performing operations that are specific to individual users. If there is any doubt, use `tenant`.

## Sub-commands of `snyk apps`

All `snyk apps` subcommands are used with the command, for example, `snyk apps create`.

To learn about all the available subcommands under `snyk apps` command, use the `--help` option, `snyk apps --help.`
