# Upgrading dependencies with automatic PRs

Once you have imported your preferred Git repositories, Snyk monitors those repos, regularly scanning them for vulnerability, license and dependency health issues. In addition to remediation advice, Snyk can then also automatically create pull requests \(PRs\) on your behalf in order to upgrade your dependencies based on the scan results.

Snyk currently supports this feature for npm, Yarn and Maven-Central projects through GitHub, GitHub Enterprise Server and BitBucket Cloud. For use with the Broker, your admin should first upgrade to v4.55.0 or later. See our docs for additional assistance when upgrading Broker.

![](../../.gitbook/assets/image%20%288%29%20%282%29%20%284%29%20%284%29%20%284%29%20%286%29%20%283%29%20%281%29%20%281%29.png)

## How it works

1. Integration is configured and users enable automatic upgrade PRs \(within the integration settings or [in the project settings](upgrading-dependencies-with-automatic-prs.md)\).
2. Snyk scans your projects as you import them and continues to monitor your projects, scanning on a regular basis thereafter.
3. Per scan, when new versions for your dependencies are identified, Snyk does the following:
   * Snyk creates automatic upgrade PRs \(frequency based on Snyk project settings\)
   * Snyk will not open a new upgrade PR for a dependency that is already changed \(upgraded or patched\) in another open Snyk PR.
   * Snyk opens separate PRs for each dependency.
   * Snyk will not create upgrade PRs for a repo that has 5 or more Snyk PRs open - if the limit of open PRs is reached, no new ones are created. This number can set to between 1-10 from the Settings. This limit only applies when creating upgrade PRs, but does count fix PRs. Fix PRs are not limited in this way.
   * By default, Snyk recommends only patch and minor upgrades, but major version upgrade can be enabled in the settings where the feature is enabled.
   * If the latest eligible version contains vulnerabilities not already found in your project, Snyk will not recommend an upgrade.
   * Snyk does not recommend upgrades to versions that are less than 21 days old. This is to avoid versions that introduce functional bugs and subsequently get unpublished, or versions that are released from a compromised account \(where the account owner has lost control to someone will malicious intent\).

## Supported languages and repos

Snyk currently supports this feature for npm, Yarn and Maven-Central projects through GitHub, GitHub Enterprise Server and BitBucket Cloud, including use of the Snyk Broker. For use with the Broker, your admin should first upgrade to v4.55.0 or later. See our docs for additional assistance when upgrading Broker.

## Upgrading dependencies with automatic PRs

Once you have imported your preferred Git repositories, Snyk monitors those repos, regularly scanning them for vulnerability, license and dependency health issues. In addition to remediation advice, Snyk can then also automatically create pull requests \(PRs\) on your behalf in order to upgrade your dependencies based on the scan results.

Snyk currently supports this feature for npm, Yarn and Maven-Central projects through GitHub, GitHub Enterprise Server and BitBucket Cloud. For use with the Broker, your admin should first upgrade to v4.55.0 or later. See our docs for additional assistance when upgrading Broker.

![](../../.gitbook/assets/image%20%288%29%20%282%29%20%284%29%20%284%29%20%284%29%20%286%29%20%283%29%20%281%29%20%285%29.png)

## Enable automatic dependency upgrade PRs for a specific project

Enable Snyk to regularly check your dependency health, recommend dependency upgrades and automatically submit PRs for upgrades on your behalf for a specific project.

Once configured, Snyk automatically creates PRs for all necessary dependencies as upgrades become available for the specific project.

{% hint style="info" %}
**Note**  
Settings on the project level override the settings on the organization level. Currently, we support all languages supported by the Git repositories that we integrate with: GitHub, GitLab, Bitbucket and Azure repos.
{% endhint %}

### **To configure automatic upgrade PRs for a specific project:**

**Prerequisites:**

* For use with Broker, your admin should first upgrade to v4.55.0.

**Steps:**

1. Navigate to the organization for which you would like to enable automatic upgrade PRs and then click Projects.
2. Navigate to the relevant project and click the Settings cog ![](../../.gitbook/assets/cog_icon.png) 
3. From the Settings area, click on the integration settings from the left panel menu.  **Note:** These settings only apply to integration for that one project. 
4. From settings that load, scroll to the **Automatic dependency upgrade pull requests** and click Disabled. 
5. From the options that appear:
   1. Snyk creates PRs up to a maximum of 10 open simultaneously - per repo. To limit this number further, select the maximum number of PRs from the dropdown list. For further information about this, read more about how it works.
   2. In the Dependencies to ignore field, enter the exact name of any dependencies that should not be handled as part of the automatic functionality. This field accepts only lower case letters.
6. Click **Update dependency upgrade settings**
7. Settings are saved

![](../../.gitbook/assets/image%20%287%29.png)

Every time Snyk scans this project now, it automatically submits upgrade PRs based on results. If a newer version is released for an existing Snyk upgrade PR or for an existing fix PR, the existing PR must be closed or merged before Snyk can raise a new PR.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

