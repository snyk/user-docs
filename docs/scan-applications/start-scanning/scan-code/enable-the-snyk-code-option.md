# Enable the Snyk Code option

To start working with Snyk Code via the Web UI, CLI, IDE, or API, the **Snyk Code** option must be enabled in your Snyk Organization settings.

## Prerequisites to enabling the Snyk Code option in an Organization

To enable the **Snyk Code** option, you must have Admin permissions in your Snyk Organization.

Snyk Code is enabled by default for new users. However, verifying that the **Snyk Code** option is enabled in your Organization settings before importing repositories for analysis is highly recommended.

{% hint style="info" %}
Snyk Code is enabled by default in GitHub integration but disabled in other SCM integrations.
{% endhint %}

Follow these steps to enable Snyk Code:

1. If you are logged out of your Snyk Account, go to [Snyk.io](http://snyk.io) and log into your Snyk Account by clicking the **Log in** button.
2. In the Snyk Web UI, select the <img src="../../../.gitbook/assets/Org Settings button - Icon (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (6).png" alt="Settings icon" data-size="line"> **Settings > Snyk Code** option.
3. In the **Enable Snyk Code** section, change the setting to **Enabled**:

<figure><img src="../../../.gitbook/assets/image (358).png" alt="Enable Snyk Code setting"><figcaption><p>Enable Snyk Code setting</p></figcaption></figure>

4\. Click **Save changes**.\
\
Snyk Code is now enabled in your Organization settings.

{% hint style="warning" %}
After the Snyk Code option is enabled, Snyk Code scans and tests only new repositories that are imported to Snyk, and it will not scan repositories that already exist in Snyk. To apply the Snyk Code analysis to repositories that are already imported to Snyk, you must re-import these repositories.
{% endhint %}

The next step is integrating your Snyk Account with the Git repositories you want to analyze using Snyk Code.
