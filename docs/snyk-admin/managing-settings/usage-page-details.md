# Usage page details

Click on **Settings** > **Usage** to view Snyk usage details, including:

* [Test usage](usage-page-details.md#test-usage): the number of tests used.
* [Contributing developers](usage-page-details.md#contributing-developers): the number of developers contributing to Projects.
* [Projects](usage-page-details.md#projects): Project test usage settings.

### Test usage

The **Test Usage** section shows how many tests you are using over the current billing period:

<figure><img src="../../.gitbook/assets/test-usage.png" alt="Test usage data"><figcaption><p>Test usage data</p></figcaption></figure>

{% hint style="info" %}
Test limits vary for Snyk products and plans. See the [plans page](https://snyk.io/plans/) for details.
{% endhint %}

{% hint style="info" %}
See [What counts as a test?](https://support.snyk.io/hc/en-us/articles/360000925418-What-counts-as-a-test-) for details of how Snyk counts tests.
{% endhint %}

### Contributing developers

{% hint style="info" %}
Currently, the integrations for which we have developer counts are GitHub, GitHub Enterprise, GitLab and the Snyk CLI.
{% endhint %}

Snyk defines contributing developers as developers having made a commit to a private repo monitored by Snyk in the last 90 days.

The **Contributing developers for Git and CLI integrations** section shows contributing developer counts, both at the org level and the group level.

The counts indicate the number of contributing developers to the default branch of the private repos connected with the integration.

We do not count contributions to public (open source) repos currently as our pricing model is based on the number of contributing developers to private repositories.

For example:

<figure><img src="../../.gitbook/assets/image__10_.png" alt="Contributing developers results"><figcaption><p>Contributing developers results</p></figcaption></figure>

* **Total unique contributors across all integrations:** the number of contributors across all the integrations in your Snyk account. Contributing developers are only counted once, even if they have contributed to multiple integrations or multiple repositories.
* **Breakdown by integration**: the number of contributors, Organizations, and repos in that integration.

#### Contributor emails

Each contributor is counted by the **author** email field, which is set within the local git configuration in the developer’s machine.

### Projects

The **Projects** section shows test usage settings for your Projects:

#### Bulk actions

For **Bulk actions**, select relevant Projects, then select to **Delete**, **Activate** or **Deactivate** the selected projects:

<figure><img src="../../.gitbook/assets/usage-projects-bulk-actions.png" alt="Bulk actions on Projects"><figcaption><p>Bulk actions on Projects</p></figcaption></figure>

#### Set test frequency

You can set the frequency of testing for each Project.&#x20;

{% hint style="info" %}
For Code analysis projects, you can only set weekly tests. To test your code daily, contact [Support](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

For each entry, you can select the frequency of testing for that Project (never, daily, or weekly)

<figure><img src="../../.gitbook/assets/usage-projects-single.png" alt="Select test frequency"><figcaption><p>Select test frequency</p></figcaption></figure>

Click **Deactivate** to never test, and also remove webhooks and stop showing the Project’s results in reporting.
