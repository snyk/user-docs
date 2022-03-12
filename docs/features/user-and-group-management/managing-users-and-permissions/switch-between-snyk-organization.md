# Switch between Snyk organizations

**On snyk.io**

1. Choose the organization you want from the drop-down menu in the top navigation.
2. If you add projects on snyk.io through GitHub integration, the projects are added to the currently chosen organization.

**In the Snyk CLI**

1. If you have only your default organization, any projects you add or update by running `snyk monitor` are automatically associated with your default organization.
2.  If you have more than one organization, you can configure the organization with which newly added projects should be associated by running `snyk config set org=orgname`.

    **Note:**\
    `orgname` should match the name as displayed in the URL of your org in the snyk UI: \[\[[https://app.snyk.io/org/\[orgname\](https://app.snyk.io/org/\[orgname)\](https://app.snyk.io/org/\[orgname\]%28https://app.snyk.io/org/\[orgname%29)\\](https://app.snyk.io/org/\[orgname]\(https://app.snyk.io/org/\[orgname\)]\(https://app.snyk.io/org/\[orgname]\(https://app.snyk.io/org/\[orgname\)\)/)].
3. If you would like to override this global configuration for individual runs of `snyk monitor`, run `snyk monitor --org=orgname`.
