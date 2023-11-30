# Set up an integration

{% hint style="info" %}
**Recap**\
You have [created a Snyk account](create-or-log-in-to-a-snyk-account.md). You now need to tell Snyk where to scan.
{% endhint %}

You must give Snyk access to your environment, to allow Snyk to scan that environment. The type of integration you need depends on what systems you use, and what you want to scan. See [Integrate with Snyk](../../integrate-with-snyk/) for information about aviaable integrators.

You can set up this integration:

* [Following a guided process](set-up-an-integration.md#guided-process-after-signup), immediately after creating a Snyk account.
* [Manually](set-up-an-integration.md#manual-process-any-time), at any point.

## Guided process (after signup)

Immediately after you [Create a Snyk account](create-or-log-in-to-a-snyk-account.md), you will see optional getting started walkthrough prompts. You can choose to provide some information to help Snyk guide your experience, and then follow the prompts to integrate your code repository for a seamless experience.

An example follows.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-16 at 9.36.53 AM.png" alt="Choose integration method"><figcaption><p>Choose integration method</p></figcaption></figure>

Fill in the details for the integration you select. If you select **GitHub**, fill in the details as shown.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-16 at 9.37.34 AM.png" alt="Set access permissions"><figcaption><p>Set access permissions</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-16 at 9.39.45 AM.png" alt="Configure automation settings &#x26; authenticate"><figcaption><p>Configure automation settings &#x26; authenticate</p></figcaption></figure>

To finish, [Import a Project](import-a-project.md) to scan:

<figure><img src="../../.gitbook/assets/image (248) (1).png" alt="Add your first project"><figcaption><p>Add your first project</p></figcaption></figure>

Alternatively, if you want to scan your code without authenticating to your source code repository, you can select the CLI integration. This allows you to run scans from your local machine and upload results to your Organization in Snyk.

Though GitHub, Bitbucket Cloud, and the CLI are shown with dedicated tiles, many other integrations are available through the **View all integrations** link.

## Manual process (any time)

You can add an integration to Snyk manually at any point. For more information, see [Integrate with Snyk](../../integrate-with-snyk/).

An example follows showing a Git repository integration

To scan code from a Git-based source code repository, you must integrate Snyk with a [Git repository integration](../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/). Snyk has pre-built integrations for GitHub, GitHub Enterprise, Bitbucket Cloud, and other repositories.

First, log in to the Snyk Web UI ([app.snyk.io](https://app.snyk.io)), and select **Integrations > Source control**.

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2022-07-26 at 13.26.22.png" alt="Select Source control integrations"><figcaption><p>Select Source control integrations</p></figcaption></figure>

</div>

{% hint style="info" %}
If an integration is already configured for your Organization, it is marked as **Configured**.
{% endhint %}

To connect Snyk with your GitHub repositories:

1. Choose whether to give Snyk access to both public and private repositories or only to public repositories.\
   The GitHub authorization screen opens.
2. In the GitHub authorization screen, click **Authorize Snyk** to provide Snyk with access to your repositories.
3. Enter your account credentials and save your details when prompted.

See [Git repository integrations (SCMs)](../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/) for more information.

## What's next?

You can now [import a Snyk Project](import-a-project.md), to tell Snyk what to scan for issues.
