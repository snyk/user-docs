# Step 1: Enabling the Snyk Code option

To start working with Snyk Code either via the Web UI, CLI, IDE, or API, the **Snyk Code** option must be enabled in your Snyk Org Settings.

**Notes**:

* To enable the **Snyk Code** option, you must have Admin permissions in your Snyk Org.
* Snyk Code is enabled by default for new users. However, it is highly recommended that you verify that the **Snyk Code** option is enabled in your Org Settings, before importing repositories for analysis.
* Snyk Code is enabled by default in GitHub integration, but it is disabled in other SCM integrations.

**To enable Snyk Code:**

1\. \[If you are logged out of your Snyk Account] Go to [Snyk.io](http://snyk.io), and log into your Snyk Account by clicking the **Log in** button:

![](<../../../../.gitbook/assets/Snyk Code - Log in button.png>)

2\. On the Snyk Web UI, click the **Org Settings** button<img src="../../../../.gitbook/assets/Org Settings button - Icon (1) (1) (1) (2).png" alt="" data-size="line">on the top menu:

![](<../../../../.gitbook/assets/Snyk Code - Org Settings button (1).png>)

3\. On the **Settings** page, select **Snyk Code** on the left menu. Then, on the **Enable Snyk Code** section, change the setting to **Enabled**:

![](<../../../../.gitbook/assets/Snyk Code - Settings - Enable Snyk Code option .png>)

4\. Click the **Save changes** button.\
\
Snyk Code is now enabled in your Org Settings.

<mark style="color:red;">**Important!**</mark> <mark style="color:red;">After the</mark> <mark style="color:red;">**Snyk Code**</mark> <mark style="color:red;">option is enabled, Snyk Code will only scan and test new repositories that are imported to Snyk, and it will NOT scan repositories that already exist in Snyk. Therefore, if you want to apply the Snyk Code analysis to repositories that are already imported to Snyk, you need to</mark> [<mark style="color:red;">re-import them</mark>](step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/re-importing-existing-repositories-for-the-snyk-code-test.md)<mark style="color:red;">.</mark>

The next step is integrating your Snyk Account with the SCM that contains the repositories you want to analyze using Snyk Code.
