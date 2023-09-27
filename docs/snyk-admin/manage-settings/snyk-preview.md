# Snyk Preview

{% hint style="info" %}
[Snyk IDE plugins](../../integrations/ide-tools/) also have preview features. These preview features are separate from Snyk Preview and can be found in the documentation for the IDE-specific plugin.
{% endhint %}

Snyk Preview lets you "opt-in" to new features which may not be available to all customers by default.

Users with Admin permissions can use this option at the Organization and Group level.

### Enable or disable a feature

To enable a feature using Snyk Preview:

1. In either the Group or Organization level, click on **Settings** > **Snyk Preview**:
2. Click **Enable feature preview** to enable or disable the relevant feature.
3. Click **Save changes**.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-04 at 11.36.07.png" alt="Snyk Preview screen"><figcaption><p>Snyk Preview screen</p></figcaption></figure>

{% hint style="info" %}
After the feature is enabled at the Group level, all Organizations in the group have this feature, and it cannot be disabled individually for these Organizations.

Additionally, some features can only be made available at the Group level - these are indicated appropriately.
{% endhint %}

## Git repository cloning

{% hint style="warning" %}
Git repository cloning is in open beta for GitHub, GitHub Enterprise, GitLab, Bitbucket Server, Bitbucket Cloud App, Bitbucket Cloud (Legacy), and Azure Repos integrations.
{% endhint %}

If enabled, Git repository cloning allows Snyk to provide additional functionality, as well as improved reliability and accuracy when performing scans by ingesting clones of your source code repositories.

### How to enable Git repository cloning

This feature can be enabled through **Snyk Preview** in the Snyk app's **General** organization settings.

<figure><img src="https://lh4.googleusercontent.com/NeiM1iGKaUMiHC-qr8n3SjlNRCr8j33XO3M5PtAdMUJaIap6RNv1UwmpiVv1siDWRnE61v490VoLTP1uXL0gUVHQDLh7FK29vGQLSvCMhlmd2NZJnbWFt3xIOxzHO7Nw7SAQDGiMwLotub8y5HU2-vbyEiY9GzA4DXwRh3xXiib7z99lqHEDDShD9jQMfWjn" alt="Enable full Git repository cloning from the integration settings in the Snyk app"><figcaption><p>Enable full Git repository cloning from the integration settings in the Snyk app</p></figcaption></figure>

{% hint style="info" %}
[Learn more](../../more-info/how-snyk-handles-your-data.md#git-repository-cloning) about how Snyk performs Git repository cloning, [applicable contract terms](../../more-info/how-snyk-handles-your-data.md#git-cloning-applicable-contract-terms), and the safeguards in place to keep your data safe.
{% endhint %}
