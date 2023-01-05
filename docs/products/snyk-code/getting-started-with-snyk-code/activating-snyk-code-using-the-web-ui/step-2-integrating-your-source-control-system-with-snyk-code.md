# Step 2: Integrating your Source Control System with Snyk Code

{% hint style="info" %}
If your SCM is already integrated with your Snyk Account, and you do not want to add additional SCMs, you can skip this step and move to [Step 3: Importing repositories to Snyk for the Snyk Code](step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/).
{% endhint %}

After you enabled Snyk Code in your Snyk Org Settings, you need to integrate your Account with the Source Control Management system that contains the repositories you want to test.

**Notes**:

* Snyk Code temporarily clones your repositories for the code analysis. This requires appropriate permissions and HTTPS access to your SCM. For more information on how data is stored in Snyk, see [How Snyk handles your data](../../../../Snyk-processes/how-snyk-handles-your-data.md).
* For more details about integrations, see [Snyk Integrations](https://docs.snyk.io/integrations).

**To integrate your SCM with Snyk Code:**

1\. On the Snyk Web UI, select the **Integrations** > **Source control** menu option.

2\. From the available **Source control** options, select the Source Control system you want to integrate with Snyk Code:

**Notes**:

* The **Source control** page only displays SCMs that are supported by Snyk Code.
* If you already have an integrated SCM, it will be indicated as “**Configured**”. If you want to use this SCM, you can skip the rest of this section and move to [Step 3: Importing repositories to Snyk for the Snyk Code testing](step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/).

<figure><img src="../../../../.gitbook/assets/image (10) (1).png" alt=""><figcaption></figcaption></figure>

The **Account credentials** page of your selected Source Control system appears.

3\. To grant Snyk access permissions to the integrated Source Control, on the **Account credentials** page, enter your account credentials and save your details.

**Note**: For more information on the integration of Snyk with each of the available SCMs, see [Git repository (SCM) integrations](https://docs.snyk.io/integrations/git-repository-scm-integrations).

After your SCM is integrated with your Snyk Account, you can import from it the repositories you want to test for vulnerabilities using Snyk Code.
