# Snyk Preview

{% hint style="info" %}
[Snyk IDE plugins](../../integrate-with-snyk/ide-tools/) also have preview features. These preview features are separate from Snyk Preview and can be found in the documentation for the IDE-specific plugin.
{% endhint %}

Snyk Preview lets you enable new features that may not be available to all customers by default.

Users with Admin permissions can use Snyk Preview at the Organization and Group level.

## Enable or disable a feature

To enable a feature using Snyk Preview:

1. At either the Group or Organization level, select **Settings** > **Snyk Preview.**
2. Select **Enable feature preview** to enable or disable a feature.
3. Click **Save changes**.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-04 at 11.36.07.png" alt="Snyk Preview screen"><figcaption><p>Snyk Preview screen</p></figcaption></figure>

{% hint style="info" %}
After the feature is enabled at the Group level, all Organizations in the Group have this feature, and it cannot be disabled individually for these Organizations.

Additionally, some features can only be made available at the Group level - these are indicated appropriately.
{% endhint %}

## Enable Git repository cloning

{% hint style="info" %}
Git repository cloning is in open beta for GitHub, GitHub Enterprise, GitLab, Bitbucket Server, Bitbucket Cloud App, Bitbucket Cloud (Legacy), and Azure Repos integrations.
{% endhint %}

When enabled, Git repository cloning allows Snyk to provide additional functionality, and improved reliability and accuracy in performing scans by ingesting clones of your source code repositories.

Git cloning can be enabled through **Snyk Preview** in the **General** Organization settings.

<figure><img src="https://lh4.googleusercontent.com/NeiM1iGKaUMiHC-qr8n3SjlNRCr8j33XO3M5PtAdMUJaIap6RNv1UwmpiVv1siDWRnE61v490VoLTP1uXL0gUVHQDLh7FK29vGQLSvCMhlmd2NZJnbWFt3xIOxzHO7Nw7SAQDGiMwLotub8y5HU2-vbyEiY9GzA4DXwRh3xXiib7z99lqHEDDShD9jQMfWjn" alt=".Enable full Git repository cloning from the Snyk settings"><figcaption><p>Enable full Git repository cloning from the Snyk settings</p></figcaption></figure>

{% hint style="info" %}
For more information bout how Snyk performs Git repository cloning, applicable contract terms, and safeguards in place to keep your data safe, see the [Git repository cloning section](../../more-info/how-snyk-handles-your-data.md#snyk-integrations-git-repository-cloning) on the page "How Snyk handles your data."
{% endhint %}
