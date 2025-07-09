# Organization-level integrations

At the Group level, you can set up and customize your Snyk Essentials integrations from the **Integrations** page where the following SCMs are available:

The Integrations page at the Group level shows all active integrations, including any data automatically synced from your existing Snyk Organizations, and provides access to the Integrations page.

You can use the wildcards while setting up your integration using the Integrations page:

* Open the **Integrations** page.&#x20;
* Select the **SCM** tag and search for GitHub or Azure DevOps.&#x20;
* Click the **Add** button.
* In the **Organizations** field, add the Organization details by using the `*` symbol. For example, using  `*snyk` integrates all SCM Organizations that have Snyk in their name.
* All the Organizations that match with the wildcard, `*` symbol will be added.&#x20;

### Adding an integration

Add integrations and populate Snyk Essentials with data from SCM tools.

{% hint style="info" %}
If you modify the permissions and scopes after the initial configuration, it is essential to either initiate an import or implement a change within the repository. This action allows Snyk to acknowledge and incorporate ‌updates effectively.
{% endhint %}

Snyk provides seamless integrations with various source control management systems such as GitHub, GitLab, Bitbucket, and Azure Repos at the Organizational level. These integrations enable you to automatically scan for vulnerabilities and receive actionable insights to enhance the security of your repositories. By integrating at this level, you can ensure comprehensive protection for all your Projects within the Organization.

Snyk can integrate with the following SCMs at the Organization-level to help you track, monitor, and fix the issues and vulnerabilities in your code:

* [GitHub](github.md)
* [GitHub Enterprise](github-enterprise.md)
* [GitHub Cloud App](github-cloud-app.md)
* [GitHub Server App](github-server-app.md)
* [GitHub Read-only Projects](github-read-only-projects.md)
* [GitLab](gitlab.md)
* [Bitbucket Cloud](bitbucket-cloud.md)
* [Bitbucket Cloud App](bitbucket-cloud-app.md)
* [Bitbucket Data Center/Server](bitbucket-data-center-server.md)
* [Azure Repositories (TFS)](azure-repositories-tfs.md)
