# Set up an integration

{% hint style="info" %}
**Recap**\
You have [created a Snyk account](create-a-snyk-account/). You now need to tell Snyk where to scan.
{% endhint %}

You must give Snyk access to your environment, to allow Snyk to scan that environment. The type of integration you need depends on what systems you use, and what you want to scan. See [Integrate with Snyk](../../integrations/) for a full list of integrations available.

You can set up this integration:

* [Following a guided process](set-up-an-integration.md#guided-process-after-signup), immediately after creating a Snyk account.
* [Manually](set-up-an-integration.md#manual-process-any-time), at any point.

### Guided process (after signup)

Immediately after you [Create a Snyk account](create-a-snyk-account/), you will see some optional getting started walkthrough prompts. You can choose to provide some information to help Snyk guide your experience, and then follow the prompts to integrate your code repository for a seamless experience.

For example:

<figure><img src="../../.gitbook/assets/image (496).png" alt=""><figcaption><p>Choose integration method</p></figcaption></figure>

If you select **GitHub**, then fill in the details as prompted.

<figure><img src="../../.gitbook/assets/image (175).png" alt=""><figcaption><p>Set access permissions</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (529).png" alt=""><figcaption><p>Configure automation settings &#x26; authenticate</p></figcaption></figure>

Finally, you can [Import a Project](import-a-project.md) to scan:

<figure><img src="../../.gitbook/assets/image (248).png" alt=""><figcaption><p>Add your first project</p></figcaption></figure>

### Manual process (any time)

You can add an integration to Snyk manually at any point (See [Integrate with Snyk](../../integrations/) for more details).

#### Example: Git repository integration

To scan code from a Git-based source code repository, you must integrate Snyk to a [Git repository integration](../../integrations/git-repository-scm-integrations/). Snyk has pre-built integrations for GitHub, GitHub Enterprise, Bitbucket Cloud, and other repositories.

First, log in to the Snyk Web UI ([app.snyk.io](https://app.snyk.io)), and select **Integrations > Source control**.

<figure><img src="../../.gitbook/assets/Screenshot 2022-07-26 at 13.26.22.png" alt="Select Source control integrations"><figcaption><p>Select Source control integrations</p></figcaption></figure>

{% hint style="info" %}
If an integration is already configured for your Organization, it is marked as **Configured**.
{% endhint %}

Next, click the source control system, for example, GitHub, to integrate with Snyk.

Finally, to grant Snyk access permissions to the integrated source control system, enter your account credentials and save your details when prompted.

See [Git repository integrations (SCMs)](../../integrations/git-repository-scm-integrations/) for more information.

### What's next?

You can now [import a Snyk Project](import-a-project.md), to tell Snyk what to scan for issues.
