# Enable the Snyk Code option

To start working with Snyk Code via the Web UI, CLI, IDE, or API, you must enable the **Snyk Code** option in your Snyk Organization settings.

## Prerequisites

To enable **Snyk Code** in your Organization, you must have Admin permissions.

You can enable Snyk Code when you [Set up an integration for the first time](../../../getting-started/quickstart/set-up-an-integration.md). Otherwise, you can verify if the **Snyk Code** option is enabled in your Organization settings before importing repositories.

## Enable Snyk Code&#x20;

1. In the Snyk Web UI, navigate to **Settings** > **Snyk Code**.
2. In the **Enable Snyk Code** section, change the setting to **Enabled**:

<figure><img src="../../../.gitbook/assets/image (358).png" alt="Enable Snyk Code setting"><figcaption><p>Enable Snyk Code setting</p></figcaption></figure>

3\. Click **Save changes**.\
\
Snyk Code is now enabled in your Organization settings.

{% hint style="warning" %}
After the Snyk Code option is enabled, Snyk Code scans and tests only new repositories that are imported to Snyk, and it will not scan repositories that already exist in Snyk. To apply the Snyk Code analysis to repositories that are already imported to Snyk, you must re-import these repositories.
{% endhint %}

The next step is integrating your Snyk Account with the Git repositories you want to analyze using Snyk Code.
