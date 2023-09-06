# Understanding the IaC CLI test results

{% hint style="info" %}
Beginning with Snyk CLI v. 1.939.0, the `snyk iac test` command offers additional results. This new offering is currently in Open Beta. To use it, install the latest version of the Snyk CLI, and enable the feature on the **Snyk Preview** page. For more information, see the instructions that follow.
{% endhint %}

Snyk Infrastructure as Code offers additional CLI `iac test` results in Snyk CLI v. 1.939.0 and later releases. For additional information on both the later and earlier versions, see the following documentation.

* [Snyk IaC CLI test results (v. 1.939.0 and later)](snyk-iac-cli-test-results-v.-1.939.0-and-later.md)
* [Snyk IaC CLI test results (v. 1.938.0 and earlier)](snyk-iac-cli-test-results-v.-1.938.0-and-earlier.md)

To use the new IaC CLI output results, install the latest Snyk CLI version or update to this version, for example, by entering:

```
npm install snyk -g 
```

For information about other methods for installing and updating the CLI, see [Install or update the Snyk CLI](../../../../snyk-cli/install-or-update-the-snyk-cli/).

In addition, to use the new IaC CLI output results, enable the option in the **Snyk Preview** page as follows:

1\. On the Snyk Web UI, select the Organization with the repositories you want to test using the CLI.

{% hint style="info" %}
By default, the CLI runs tests under your **Preferred Organization**, as defined in your **Account settings**. You can change your **Preferred Organization** or set another Organization for the CLI tests.
{% endhint %}

<figure><img src="../../../../.gitbook/assets/OS - Automatic Dependency Upgrade - Selecting Organization (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1)  (1).png" alt="Select the Organization for snyk iac test"><figcaption><p>Select the Organization for snyk iac test</p></figcaption></figure>

2\. When the required Organization is open, select **Org Settings**:

<figure><img src="../../../../.gitbook/assets/OS - Automatic Dependency Upgrade - Org Settings button (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1 (12).png" alt="Select Org Settings"><figcaption><p>Select Org Settings</p></figcaption></figure>

3\. On the **Org** **Settings** page, select **Snyk Preview** and navigate to the section **Additional CLI result information for Infrastructure as Code**:

<figure><img src="../../../../.gitbook/assets/IaC - CLI - New results - Enabling in Snyk Preview - Section.png" alt="dditional CLI result information for Infrastructure as Code"><figcaption><p>Additional CLI result information for Infrastructure as Code</p></figcaption></figure>

4\. In the **Additional CLI result information for Infrastructure as Code** section, move the slider to **Enabled** and click the **Save changes** button.
