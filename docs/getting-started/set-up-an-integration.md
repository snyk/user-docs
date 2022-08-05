# Set up an integration

{% hint style="info" %}
**Recap**\
You have [created a Snyk account](create-a-snyk-account.md), you now need to tell Snyk where to scan.
{% endhint %}

You must give Snyk access to your environment, to allow Snyk to scan that environment. Access is provided by [Snyk integrations](../integrations/).

### Integration types

The type of integration you need depends on what systems you use, and what you want to scan.

#### Example: set up a source control integration

To scan open source libraries (using Snyk Open Source) or your application code (using Snyk Code) from a Git-based source code repository, you will need to connect using a [Git repository integration](../integrations/git-repository-scm-integrations/). Snyk has pre-built integrations for GitHub, GitHub Enterprise, Bitbucket Cloud and other repositories.&#x20;

First, log in to the Snyk Web UI ([app.snyk.io](https://app.snyk.io)), and select **Integrations > Source control**.

![](<../.gitbook/assets/Screenshot 2022-07-26 at 13.26.22.png>)

{% hint style="info" %}
If an integration is already configured for your Organization, it is marked as **Configured**.
{% endhint %}

Next, click the source control system (for example, [GitHub](../integrations/git-repository-scm-integrations/github-integration.md)) to integrate with Snyk.

Finally, to grant Snyk access permissions to the integrated source control system, enter your account credentials and save your details when prompted.

#### Example: set up a container integration

To scan your container images, choose a [container registry integration](../products/snyk-container/image-scanning-library/), to connect the registry with Snyk.

First, log in to the Snyk Web UI ([app.snyk.io](https://app.snyk.io)), and select **Integrations > Container registries**

![](<../.gitbook/assets/Screenshot 2022-07-26 at 13.16.05.png>)

Next, click the registry you want to integrate with Snyk, and enter details as prompted.

Finally, save the changes, to integrate this entry with Snyk.

### What's next?

You can now [import a Snyk Project](import-a-project.md), to tell Snyk what to scan for issues.
