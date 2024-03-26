# Upgrade dependencies with automatic PRs

After you import Git repositories, Snyk continuously monitors these repositories by scanning them for vulnerability, license, and dependency health issues. In addition to providing fix advice, Snyk creates pull requests (PRs) according to your configuration settings.

<figure><img src="../../../.gitbook/assets/Upgrade Broker.png" alt="Snyk Bot conversation card in GitHub reporting PR raised"><figcaption><p>Snyk Bot conversation card in GitHub reporting PR raised</p></figcaption></figure>

## Supported languages and SCMs

Snyk supports the **Automatic dependency upgrade pull requests** feature for npm, Yarn, and Maven Central repositories with the following Source Control Managers (SCMs): GitHub, GitHub Enterprise, and Bitbucket Cloud.

You can also use this feature with Snyk Broker. To use this feature, you must upgrade Snyk Broker to v. 1.4.55.0 or later. For more information, see [Upgrade the Snyk Broker client](../../../enterprise-configuration-and-implementation/snyk-broker/upgrade-the-snyk-broker-client.md).

## Automatic dependency (upgrade) PRs

Automatic dependency or upgrade PRs work as follows.

1. The **Automatic dependency upgrade pull requests** option must be enabled in [the Integration Settings at the Organization level](upgrade-dependencies-with-automatic-prs.md#enabling-the-automatic-dependency-upgrade-prs-option-for-an-entire-organization) or in the Project Settings.
2. When you import repositories, Snyk scans the repositories and provides scan results. Snyk then continues to monitor your Open Source Projects, scanning them on a regular basis. The re-scan frequency is based on the schedule set in the Project Settings.
3. For each scan, when new versions for your dependencies are identified, Snyk creates automatic upgrade PRs.
   * Snyk does not open a new upgrade PR for a dependency that is already upgraded or patched in another open Snyk PR.
   * Snyk opens separate PRs for each dependency.
   * By default, Snyk does not create upgrade PRs for a Project that has five or more open Snyk PRs. After the limit of open PRs is reached, no new PRs are created. This limit can be set on the Integration or Project Settings to be between 1-10. This limit applies only to Upgrade PRs, but it does count Fix PRs. However, Fix PRs are not limited in this way.
   * By default, Snyk recommends only patches and minor upgrades. However, recommendations for major version upgrades can be requested on the **Settings** page where the feature is enabled.
   * If the latest eligible version contains vulnerabilities that are not found yet in your Project, Snyk will not recommend an upgrade.
   * Snyk does not recommend upgrades to versions that are less than 21 days old. This is to avoid versions that introduce functional bugs and are subsequently unpublished or versions that are released from a compromised account, an account for which the account owner has lost control to someone with malicious intent.

## How to enable the Automatic dependency upgrade PRs option

You can configure Snyk to check your dependency health regularly, recommend dependency upgrades, and automatically submit PRs for upgrades for an entire Organization or a specific Project. After configuration, Snyk automatically creates PRs for all the necessary dependencies as upgrades become available for the scanned Projects.

By default, the Project Settings inherit the Organization Settings. However, when the Settings on the Organization and Project levels are different, the Project Settings override the Organization Settings.

{% hint style="info" %}
**Reminder**\
Automatic dependency upgrade PRs are available only for the following SCM integrations: GitHub, GitHub Enterprise, and Bitbucket Cloud.
{% endhint %}

### How to enable the Automatic dependency upgrade PRs option for an entire Organization

Follow these steps to configure automatic upgrade PRs for an entire Organization:

1. On the Snyk Web UI, open the required Organization.
2.  Navigate to **Settings > Organization Settings > Integrations,** find your configured SCM, and click **Edit settings** at the end of the row for that integration.

    <figure><img src="../../../.gitbook/assets/scm-integration_edit settings_20oct2022.png" alt="Editing integration settings."><figcaption><p>Editing integration settings</p></figcaption></figure>
3. On the **Settings** page of the selected integration, navigate to the **Automatic dependency upgrade PRs** section.
4. In this section, perform the following actions:
   * Slider - change to **Enable**.
   * **Maximum number of open upgrade PRs allowed** – define how many open Snyk PRs a Project can have to also receive a dependency upgrade PR; the maximum is ten. When the limit of the open PRs is reached, no new upgrade PRs are created.
   * **Include major version in upgrade recommendation** – select whether to include major version upgrades in the recommendations. By default, only patches and minor versions are included in the upgrade recommendations.
   * **Dependencies to ignore** – enter the exact name of the dependencies that should NOT be included in the **Automatic upgrade** operation. You can only enter lowercase letters.

<figure><img src="../../../.gitbook/assets/Github-integrations-auto-dep-PRs.png" alt="Enabling Automatic dependency upgrade PFs"><figcaption><p>Enabling Automatic dependency upgrade PFs</p></figcaption></figure>

5. To save and apply your changes, select one of the following from the **Save** dropdown:
   * **Save**  – your changes are saved and will be applied to all the Projects in the Organization that are configured to inherit these Settings from the Organization. Projects that have Custom Settings will not be influenced by this change.
   * **Save changes and apply to all overridden Projects** – your changes are saved and will be applied to all the Projects in the Organization. Projects that have Custom Settings will inherit these Organization Settings, and their Custom Settings will be overridden.

From now on, every time Snyk scans any Project in the Organization, it automatically submits Upgrade PRs if the scan discovers that an upgrade is available.

If a newer version is released for an existing Snyk Upgrade PR, or for an existing Fix PR, the existing PR must be closed or merged before Snyk can raise a new PR.

### How to enable the Automatic dependency upgrade PRs option for a specific Project

The Settings on the Project level override the Settings on the Organization level. However, the Settings on the Organization level can override the Custom Settings of a Project if they are configured after the Project-level customization and are saved with the option **Apply changes to all overridden projects**.

Follow these steps to configure automatic upgrade PRs for a specific Project:

1. From the Snyk Web UI, open the Organization that includes the Project you want to configure.
2. In the list of Projects, locate and expand the **Project** for which you want to enable automatic upgrade PRs.
3.  Click the **Project settings** at the end of the Project row.\


    <figure><img src="../../../.gitbook/assets/Screenshot 2023-04-27 at 12.55.40.png" alt="Project settings on the Project jow"><figcaption><p>Project settings on the Project jow</p></figcaption></figure>
4. On the **Project** **Settings** page, select the integration you are using.\
   <img src="../../../.gitbook/assets/image (1) (1) (1) (1).png" alt="Select the integratiion" data-size="original">Project
5. On the **Integration** page, scroll to the **Automatic dependency upgrade pull requests** section and choose one of the following:
   * **Inherit from Integration settings** – apply the Integration Settings of the Organization to the selected Project.\
     If the **Automatic dependency upgrade PRs option is disabled for the Organization**, this option will also be disabled for the Project.
   * **Customize for only this Project** – apply specific settings of the **Automatic dependency upgrade PRs** option on the Project. If you select this option:
     * Change the slider to **Enabled**.
     * In **Include major version in upgrade recommendation,** select one of the available options to define whether major version upgrades will be included in the recommendations.\
       By default, only patches and minor versions are included in the upgrade recommendations.
     * In **Limit Snyk to this many dependency upgrade PRs open simultaneously,** define how many open Snyk PRs a Project can have to receivee also a dependency upgrade PR. You can set a number between 1 and 10.\
       When the limit of the open PRs is reached, no new upgrade PRs are created.\
       By default, to _also_ receive a dependency upgrade PR, a Project can have _up to four_ open PRs.
     * In **Dependencies to ignore**, enter the exact name of the dependencies to _exclude_ from the **Automatic upgrade** operation.\
       You can only enter lowercase letters.
     * Click **Update dependency upgrade settings** to save your changes.

<figure><img src="../../../.gitbook/assets/auto-dependency-prs_1dec2022.png" alt="Automatic dependency upgrade pull requests settings at the Project level."><figcaption><p>Automatic dependency upgrade pull requests settings at the Project level</p></figcaption></figure>

After you have completed these steps, Snyk scans the Project and automatically submits Upgrade PRs if the scan discovers that an upgrade is available. If a newer version is released for an existing Snyk Upgrade PR or an existing Fix PR, the existing PR must be closed or merged before Snyk can raise a new PR.
