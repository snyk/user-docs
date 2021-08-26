# Usage page details

Click on settings ![cog\_icon.png](../../.gitbook/assets/cog_icon.png)

 &gt; **Usage** to view:

* [Test usage](usage-page-details.md): the numbers of tests used.
* [Contributing developers](usage-page-details.md) number of developers contributing to projects.
* [Projects](usage-page-details.md): project test usage settings.

The **Test Usage** section shows how many tests you are using over the current billing period:

![](../../.gitbook/assets/test-usage.png)



{% hint style="info" %}
Test limits vary for Snyk products and plans. See the [plans page](https://snyk.io/plans/) for details.
{% endhint %}

{% hint style="info" %}
See [What counts as a test?](https://support.snyk.io/hc/en-us/articles/360000925418-What-counts-as-a-test-) for details of how Snyk counts tests.
{% endhint %}

## View contributing developers

{% hint style="info" %}
Currently, the integrations for which we have developer counts are GitHub, GitHub Enterprise, GitLab and the Snyk CLI.
{% endhint %}

Snyk defines contributing developers as developers having made a commit to a private repo monitored by Snyk in the last 90 days.

The **Contributing developers for Git and CLI integrations** section shows contributing developer counts, both at the org level and the group level.

The counts indicate the number of contributing developers to the default branch of the private repos connected with the integration.

We do not count contributions to public \(open source\) repos currently as our pricing model is based on the number of contributing developers to private repositories.

For example:

![image\_\_10\_.png](https://support.snyk.io/hc/article_attachments/4403676877585/image__10_.png)


The **total unique contributors across all integrations** count shows the number of contributors across all the integrations in your Snyk account. Contributing developers are only counted once, even if they have contributed to multiple integrations or multiple repositories.

The **Breakdown by integration** section shows the number of contributors, orgs, and repos in that integration.

### Contributor emails

Each contributor is counted by the **author** email field, which is set within the local git configuration in the developer’s machine.

**Privacy**

Snyk does not store any actual git user emails in its database but instead uses the _hashes_ of the emails for counting contributing developers. Snyk uses the hash of the local git author email, that is, the email that is set locally in the user's git config, whether for one of the Source Control Management integrations \(such as GitHub\), or from the Snyk CLI.

## View projects

The **Projects** section shows test usage settings for your projects:

### Bulk actions

For **Bulk actions**, select relevant projects, then select to **Delete**, **Activate** or **Deactivate** the selected projects:

![Usage-projects-bulk-actions.png](https://support.snyk.io/hc/article_attachments/4403674675985/Usage-projects-bulk-actions.png)


### Set test frequency

You can set the frequency of testing for each project.

For each entry, you can select the frequency of testing for that project \(never, daily, or weekly\)

![usage-projects-single.png](https://support.snyk.io/hc/article_attachments/4403676922769/usage-projects-single.png)


Click **Deactivate** to never test, and also remove webhooks and stop showing the project’s results in reporting.

