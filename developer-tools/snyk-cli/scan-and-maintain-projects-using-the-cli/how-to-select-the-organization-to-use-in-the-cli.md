# How to select the Organization to use in the CLI

When you run commands with the CLI such as `snyk monitor` and `snyk test`, Snyk uses your `Preferred Organization`, which can be configured in your account. If you want commands to run against a different Organization, you can specify that Organization either globally or individually.

**Note**: The Snyk CLI help specifies using `--org=<ORG_ID>`; this article specifies using the `orgslugname`. The `ORG_ID` works in both the CLI and the API. The organization slug name works in the CLI, but not in the API.

## Specify an Organization globally

Run `snyk config set org=orgslugname`. Note: `orgslugname` must match the slug name as displayed in the URL of your org in the Snyk UI: `https://app.snyk.io/org/[orgslugname]`.

## Specify an Organization individually

To override this global configuration for individual runs, specify the Organization as follows:

`snyk monitor --org=orgslugname` or `snyk test --org=orgslugname`.

Example: `$ snyk test --org=my-team`
