# Switch between Snyk organizations

**On snyk.io**

1. Choose the organization you want from the drop-down menu in the top navigation.
2. If you add projects on snyk.io through GitHub integration, the projects are added to the currently chosen organization.

**In the Snyk CLI**

1. If you have only your default organization, any projects you add or update by running `snyk monitor` are automatically associated with your default organization.
2. If you have more than one organization, you can configure the organization with which newly added projects should be associated by running `snyk config set org=ORG_ID`.
3. If you would like to override this global configuration for individual runs of `snyk monitor`, run `snyk monitor --org=ORG_ID`.

The default is the `<ORG_ID>` that is the current preferred organization in your [Account settings](https://app.snyk.io/account).

For more information see the article [How to select the organization to use in the CLI](https://support.snyk.io/hc/en-us/articles/360000920738-How-to-select-the-organization-to-use-in-the-CLI).
