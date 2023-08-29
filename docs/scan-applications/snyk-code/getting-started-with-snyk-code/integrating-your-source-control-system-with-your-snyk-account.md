# Integrating your Source Control System with your Snyk account

{% hint style="info" %}
If your SCM is already integrated with your Snyk Account, and you do not want to add additional SCMs, you can skip this step and move to [Importing repositories for Snyk Code testing](importing-repositories-for-snyk-code-testing.md).
{% endhint %}

After you enable Snyk Code in your Snyk Organization settings, to work in the Web UI or with the API but not the CLI, you must integrate your account with the Source Control Management system that contains the repositories you want to test.

Then you can import the required repositories to your Snyk account, and Snyk Code will automatically analyze them and show the analysis results.

{% hint style="info" %}
Snyk Code temporarily clones your repositories for code analysis. This requires appropriate permissions and HTTPS access to your SCM.

For more information on how data is stored in Snyk, see [How Snyk handles your data](../../../more-info/how-snyk-handles-your-data.md). For more details about integrations, see [Integrate with Snyk](../../../integrations/).
{% endhint %}

Follow these steps to integrate your SCM with your Snyk account

1\. In the Snyk Web UI, select the **Integrations** > **Source control** option.

2\. From the available **Source control** options, select the SCM system you want to integrate.

The **Source control** integrations display only SCMs that are supported by Snyk Code.

{% hint style="info" %}
If you already have an integrated SCM, it is indicated as **Configured**. If you want to use this SCM, you can continue with[ Importing repositories to Snyk for the Snyk Code testing](importing-repositories-for-snyk-code-testing.md).
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (106) (1) (1).png" alt=""><figcaption><p>Source control integrations showoing GitLab being selected</p></figcaption></figure>

3\. On the **Account credentials** page, enter your account credentials and save your details.

This grants access permissions to Snyk for the integrated SCM.

For more information on integrating Snyk with each of the available SCMs, see [Git repositories (SCMs)](../../../integrations/git-repository-scm-integrations/).

After your SCM is integrated with your Snyk account, you can import the repositories you want to test for vulnerabilities using Snyk Code.
