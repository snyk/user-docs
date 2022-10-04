# Switch between Snyk Organizations

**On the Snyk Web UI**

1. Choose the Organization you want from the drop-down menu in the top navigation.
2. If you add Projects through GitHub integration, these Projects are added to the currently chosen Organization.

**In the Snyk CLI**

1. If you have only your default Organization, any Projects you add or update by running `snyk monitor` are automatically associated with your default Organization.
2. If you have more than one Organization, you can configure the organization with which newly added Projects should be associated by running `snyk config set org=ORG_ID`.
3. If you would like to override this global configuration for individual runs of `snyk monitor`, run `snyk monitor --org=ORG_ID`.

The default is the `<ORG_ID>` that is the current preferred Organization in your [Account settings](https://app.snyk.io/account).

For more information see [How to select the organization to use in the CLI](https://support.snyk.io/hc/en-us/articles/360000920738-How-to-select-the-organization-to-use-in-the-CLI).
