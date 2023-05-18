# Step 1: Enabling the Snyk Code option

To start working with Snyk Code either via the Web UI, CLI, IDE, or API, the **Snyk Code** option must be enabled in your Snyk Org Settings.

**Notes**:

* To enable the **Snyk Code** option, you must have Admin permissions in your Snyk Org.
* Snyk Code is enabled by default for new users. However, it is highly recommended that you verify that the **Snyk Code** option is enabled in your Org Settings, before importing repositories for analysis.
* Snyk Code is enabled by default in GitHub integration, but it is disabled in other SCM integrations.

**To enable Snyk Code:**

1\. \[If you are logged out of your Snyk Account] Go to [Snyk.io](http://snyk.io), and log into your Snyk Account by clicking the **Log in** button.

2\. On the Snyk Web UI, select the <img src="../../../../.gitbook/assets/Org Settings button - Icon (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (6).png" alt="" data-size="line"> **Settings > Snyk Code** menu option.

3\. In the **Enable Snyk Code** section, change the setting to **Enabled**:

<figure><img src="../../../../.gitbook/assets/image (358).png" alt=""><figcaption></figcaption></figure>

4\. Click **Save changes**.\
\
Snyk Code is now enabled in your Org Settings.

{% hint style="warning" %}
After the Snyk Code option is enabled, Snyk Code will only scan and test new repositories that are imported to Snyk, and it will not scan repositories that already exist in Snyk. So to apply the Snyk Code analysis to repositories that are already imported to Snyk, you need to re-import these repositories.
{% endhint %}

The next step is integrating your Snyk Account with the Git repositories you want to analyze using Snyk Code.
