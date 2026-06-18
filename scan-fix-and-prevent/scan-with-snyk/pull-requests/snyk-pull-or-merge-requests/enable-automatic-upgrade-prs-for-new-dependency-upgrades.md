# Enable automatic upgrade PRs for new dependency upgrades

{% hint style="info" %}
**Feature availability**

This feature is supported for the following SCM integrations: GitHub, GitHub Enterprise, GitHub Cloud App, Bitbucket Server, Bitbucket Cloud, Bitbucket Connect, GitLab, and Azure Repos.
{% endhint %}

Keeping dependencies up-to-date is crucial for security, performance, and compatibility. Snyk simplifies the process by scanning Projects for outdated dependencies and proposing updates through automated pull requests. This lets teams:

* Review and merge changes in a controlled way.
* Keep Projects aligned with the latest package versions.
* Automate testing before merging dependency changes.

If you use Snyk to update open-source and private dependencies, you reduce manual work and lower security risk. You also keep your PR backlog focused, because Snyk closes obsolete Fix PRs automatically by default when their targeted issues are already resolved.

## Defining automatic upgrade PRs

Automatic dependency or upgrade PRs work as follows.

1. The **Automatic dependency upgrade pull requests** option must be enabled in [the Integration Settings at the Organization level](enable-automatic-upgrade-prs-for-new-dependency-upgrades.md#how-to-enable-the-automatic-dependency-upgrade-prs-option-for-an-entire-organization) or in the Project Settings.
2. When you import repositories, Snyk scans the repositories and provides scan results. Snyk then continues to monitor your Open Source Projects, scanning them on a regular basis. The re-scan frequency is based on the schedule set in the Project Settings.
3. For each scan, when new versions for your dependencies are identified, Snyk creates automatic upgrade PRs.
   * Snyk opens separate PRs for each dependency
   * By default, Snyk does not create upgrade PRs for a Project that has five or more open Snyk PRs. After the limit is reached, Snyk does not create new upgrade PRs. You can set this limit between 1 and 10 in the integration or Project settings. This limit applies only to upgrade PR creation, but the count includes Fix PRs and Backlog PRs. Automatic closure of obsolete Fix PRs helps free capacity for new upgrade PRs.
   * By default, Snyk only recommends minor upgrades through automatic Upgrade PRs. However, recommendations for major version upgrades can be turned on in **Settings.**
   * If the latest eligible version contains vulnerabilities that are not found yet in your Project, Snyk does not recommend an upgrade.
   * Snyk does not recommend upgrades to versions that are less than 21 days old. Versions under this threshold have a higher risk of unknown functional bugs, unpublishing, or release through a compromised account.

## Supported languages and SCMs

Automatic dependency upgrades are supported for [certain package managers](../../snyk-open-source/manage-vulnerabilities/troubleshoot-fixing-vulnerabilities-with-snyk-open-source.md) and repositories with the following Source Control Managers (SCMs): GitHub, GitHub Enterprise, GitHub Cloud App, Bitbucket Server, Bitbucket Cloud, Bitbucket Connect, GitLab, and Azure Repos.

You can also use this feature with Snyk Broker. To use this feature, you must upgrade Snyk Broker to v. 1.4.55.0 or later. For more information, see [Upgrade the Snyk Broker client](https://app.gitbook.com/s/IgtgtomLQ2TUgSKOMSAm/snyk-broker/update-the-snyk-broker-client).

## Enabling automatic upgrade PRs

After importing a Git repository, Snyk continuously monitors for vulnerabilities, license, and dependency health issues. In addition to fix advice through the web UI, Snyk can automatically create pull requests (PRs) to solve for a variety of vulnerability types according to your configuration settings.

You can configure Snyk to regularly check your dependency health, recommend dependency upgrades, and automatically submit PRs for upgrades for either an entire Organization or a specific Project. After configuration, Snyk automatically creates PRs for all applicable dependencies as upgrades become available for scanned Projects.

{% hint style="info" %}
Automatic dependency upgrade PRs are available only for the following SCM integrations: GitHub, GitHub Enterprise, and Bitbucket Cloud.
{% endhint %}

Snyk also evaluates open Fix PRs during each recurring test. If a Fix PR no longer targets any active issues, Snyk closes it automatically and adds a comment in your source control manager.

