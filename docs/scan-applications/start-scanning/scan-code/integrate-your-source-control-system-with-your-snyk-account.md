# Integrate your Source Control System with your Snyk account

{% hint style="info" %}
If your SCM is already integrated with your Snyk Account, and you do not want to add additional SCMs, you can skip this step and move to [Importing repositories for Snyk Code testing](import-repositories-to-scan-with-snyk-code.md).
{% endhint %}

After you enable Snyk Code in your Snyk Organization settings to work in the Web UI or with the API but not the CLI, you must integrate your account with the Source Control Management system that contains the repositories you want to test.

Then you can import the required repositories to your Snyk account, and Snyk Code automatically analyzes them and displays the analysis results.

{% hint style="info" %}
Snyk Code temporarily clones your repositories for code analysis. This requires appropriate permissions and HTTPS access to your SCM.

For more information on how data is stored in Snyk, see [How Snyk handles your data](../../../more-info/how-snyk-handles-your-data.md). For more details about integrations, see [Integrate with Snyk](../../../integrations/).
{% endhint %}

To integrate your SCM with your Snyk account:

1\. In the Snyk Web UI, navigate to **Settings** > **Integrations** > **Source control**.

{% hint style="info" %}
If you already have an integrated SCM, it is marked as **Configured**. If you want to use the configured SCM, continue with[ Importing repositories to Snyk for the Snyk Code testing](import-repositories-to-scan-with-snyk-code.md).
{% endhint %}

2\. From the available options, select the SCM system you want to integrate by clicking **Edit settings**.

{% hint style="info" %}

{% endhint %}

The **Source control** integrations display only SCMs that are supported by Snyk Code.

<figure><img src="../../../.gitbook/assets/code_source_control_options.png" alt=""><figcaption><p>Source control options for Snyk Code</p></figcaption></figure>

3\. On the integration page, enter your account credentials and save your details.

This grants Snyk access permissions for the integrated SCM.

For more information on integrating Snyk with each of the available SCMs, see [Git repositories (SCMs)](../../../integrations/git-repository-scm-integrations/).

After you have integrated the SCM with your Snyk account, you can import the repositories you want to scan using Snyk Code.
