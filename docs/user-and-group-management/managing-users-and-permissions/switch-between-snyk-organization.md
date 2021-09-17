# Switch between Snyk organization

**On snyk.io**

1. Choose the organization you want from the drop-down menu in the top navigation.
2. If you add projects on snyk.io via GitHub integration, they will be added to the currently chosen organization.

**In the Snyk CLI**

1. If you have only your default organization, any projects you add or update by running `snyk wizard` or `snyk monitor` will be automatically associated with your default organization.
2. If you have more than one organization, you can configure which organization newly added projects should be associated with by running `snyk config set org=orgname`.

   **Note**  
   `orgname` should match the name as displayed in the URL of your org in the snyk UI: \[[https://app.snyk.io/org/\[orgname\]\(https://app.snyk.io/org/\[orgname\)\](https://app.snyk.io/org/[orgname]%28https://app.snyk.io/org/[orgname%29\)\].

3. If you would like to override this global configuration for individual runs of `snyk wizard` or `snyk monitor`, run `snyk monitor --org=orgname` or `snyk wizard --org=orgname`.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

