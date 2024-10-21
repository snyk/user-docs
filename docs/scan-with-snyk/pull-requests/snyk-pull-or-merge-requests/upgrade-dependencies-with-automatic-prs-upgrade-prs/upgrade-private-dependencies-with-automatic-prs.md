# Upgrade private dependencies with automatic PRs

{% hint style="warning" %}
**Release status**&#x20;

Upgrade private dependencies with automatic PRs is currently in [Closed Beta](https://docs.snyk.io/getting-started/snyk-release-process).&#x20;

Contact your account manager to get access to this feature. After your account manager turns on the feature flag, you can activate this feature in the [Snyk Preview](https://docs.snyk.io/snyk-admin/snyk-preview).
{% endhint %}

## Private Packages API

To receive pull requests that upgrade your private dependencies, you need to inform Snyk about the private libraries you internally develop and the available versions for them. Snyk requires this information to compute the right upgrade for you. The data related to private packages can be sent using this [API](https://apidocs.snyk.io/experimental?version=2023-11-20\~experimental#post-/groups/-group\_id-/packages).

The API is available at the Group level, meaning that you only send once the details related to a package for your Group, and those details will propagate to all your repositories, Projects, and Organizations within that Group.

## Supported languages and SCMs

Snyk supports the **Automatic private dependency upgrade pull requests** feature for npm, Yarn, and Maven Central repositories with the following Source Control Managers (SCMs): GitHub, GitHub Enterprise, GitHub Cloud App, Bitbucket Server, Bitbucket Cloud, Bitbucket Connect, GitLab, and Azure Repos.

Navigate to the [Upgrade open source dependencies with automatic PRs](upgrade-open-source-dependencies-with-automatic-prs.md#supported-languages-and-scms) page for more details.

## How to enable the Automatic private dependency upgrade PRs option

You can configure Snyk to check your repositories, recommend dependency upgrades, and automatically submit PRs for upgrades for an entire Organization or a specific Project. After configuration, Snyk automatically creates PRs for all the necessary dependencies as upgrades become available for the scanned Projects.

By default, the Project Settings inherit the Organization Settings. However, when the Settings on the Organization and Project levels are different, the Project Settings override the Organization Settings.

{% hint style="info" %}
**Reminder**\
Automatic dependency upgrade PRs are available only for the following SCM integrations: GitHub, GitHub Enterprise, and Bitbucket Cloud.
{% endhint %}

### How to enable the Automatic private dependency upgrade PRs option for an entire Organization

Follow these steps to configure automatic upgrade PRs for an entire Organization:

1. On the Snyk Web UI, open the required Organization.
2. Navigate to **Settings > Organization Settings > Integrations,** find your configured SCM, and click **Edit settings** at the end of the row for that integration.
3. On the **Settings** page of the selected integration, navigate to the **Automatic dependency upgrade PRs** section.
4. In this section, perform the following actions:
   * Slider - change to **Enable**.
   * **Maximum number of open upgrade PRs allowed** – define how many open Snyk PRs a Project can have to also receive a dependency upgrade PR; the maximum is ten. When the limit of the open PRs is reached, no new upgrade PRs are created.
   * **Dependencies to ignore** – enter the exact name of the dependencies that should NOT be included in the **Automatic upgrade** operation. You can only enter lowercase letters.
   * **Include major version in upgrade recommendation** – select whether to include major version upgrades in the recommendations. By default, only patches and minor versions are included in the upgrade recommendations.
   * **Receive private packages upgrade** - select the option if you want to include your private dependencies to the list of dependencies included in the **Automatic upgrade** operation.
     * If you turn on this option, the PRs you select to be received will be split equally between 50% for open source upgrades and 50% for private upgrades.

<figure><img src="../../../../.gitbook/assets/image (441).png" alt="Enabling Automatic public and private dependency upgrade PFs"><figcaption><p>Enabling Automatic public and private dependency upgrade PFs</p></figcaption></figure>



5. To save and apply your changes, select one of the following from the **Save** dropdown:
   * **Save**  – your changes are saved and will be applied to all the Projects in the Organization that are configured to inherit these Settings from the Organization. Projects that have Custom Settings will not be influenced by this change.
   * **Save changes and apply to all overridden Projects** – your changes are saved and will be applied to all the Projects in the Organization. Projects that have Custom Settings will inherit these Organization Settings, and their Custom Settings will be overridden.

From now on, every time Snyk scans any Project in the Organization, it automatically submits Upgrade PRs if the scan discovers that an upgrade is available.

If a newer version is released for an existing Snyk Upgrade PR, or for an existing Fix PR, the existing PR must be closed or merged before Snyk can raise a new PR.

### How to enable the Automatic dependency upgrade PRs option for a specific Project

Navigate to the [How to enable the Automatic dependency upgrade PRs](upgrade-open-source-dependencies-with-automatic-prs.md#how-to-enable-the-automatic-dependency-upgrade-prs-option-for-a-specific-project) option for a specific Project page for more details.
