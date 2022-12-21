# Switch between Snyk Organizations

**On the Snyk Web UI**

1. Choose the Organization you want using **Switch org** <img src="../../../.gitbook/assets/compare_arrows.png" alt="" data-size="line"> in the left navigation menu.&#x20;
2. If you add Projects through GitHub integration, these Projects are added to the currently chosen Organization.

**In the Snyk CLI**

1. If you have only your default Organization, any Projects you add or update by running `snyk monitor` are automatically associated with your default Organization.
2. If you have more than one Organization, you can configure the organization with which newly added Projects should be associated by running `snyk config set org=ORG_ID`.
3. If you would like to override this global configuration for individual runs of `snyk monitor`, run `snyk monitor --org=ORG_ID`.

The default `<ORG_ID>` is the currently preferred Organization in your [Account settings](https://app.snyk.io/account).

For more information see [How to select the organization to use in the CLI.](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/how-to-select-the-organization-to-use-in-the-cli)
