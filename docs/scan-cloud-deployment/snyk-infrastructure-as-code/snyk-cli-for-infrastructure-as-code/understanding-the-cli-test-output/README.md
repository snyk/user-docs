# Understanding the IaC CLI test results

{% hint style="info" %}
As of Snyk CLI v. 1.939.0, the `snyk iac test` command offers additional results. This new offering is currently in Open Beta. To use it, you need to install the latest version of the Snyk CLI, and to enable the feature in the **Snyk Preview** page. For more information, see the instructions that follow.
{% endhint %}

Snyk Infrastructure as Code offers additional CLI `iac test` results in Snyk CLI v. 1.939.0 and later releases. For additional information on both the later and earlier versions, see the following documentation.

{% content-ref url="snyk-iac-cli-test-results-v.-1.939.0-and-later.md" %}
[snyk-iac-cli-test-results-v.-1.939.0-and-later.md](snyk-iac-cli-test-results-v.-1.939.0-and-later.md)
{% endcontent-ref %}

{% content-ref url="snyk-iac-cli-test-results-v.-1.938.0-and-earlier.md" %}
[snyk-iac-cli-test-results-v.-1.938.0-and-earlier.md](snyk-iac-cli-test-results-v.-1.938.0-and-earlier.md)
{% endcontent-ref %}

To use the new IaC CLI output results, you need to install the latest Snyk CLI version or to update to this version by entering:

```
npm install snyk -g 
```

**Note**: For Install and Update commands for other CLI installation methods, see [Install or update the Snyk CLI](../../../../snyk-cli/install-the-snyk-cli/).

In addition, to use this IaC CLI offering, you need to enable this option in the **Snyk Preview** page, as follows:

1\. On the Snyk Web UI, open the organization whose repositories you want to test via the CLI:

**Note**: By default, the CLI will run tests under your **Preferred Organization**, as defined in your **Account settings**. You can change your **Preferred Organization**, or set another organizations for the CLI tests.

![](<../../../../.gitbook/assets/OS - Automatic Dependency Upgrade - Selecting Organization (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1)  (3).png>)

2\. Once the required organization is open, click the **Org Settings** button on the top menu:

![](<../../../../.gitbook/assets/OS - Automatic Dependency Upgrade - Org Settings button (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1  (5).png>)

3\. On the Org **Settings** page, select **Snyk Preview** on the left menu, and move to the section – **Additional CLI result information for Infrastructure as Code**:

![](<../../../../.gitbook/assets/IaC - CLI - New results - Enabling in Snyk Preview - Section.png>)

4\. On the **Additional CLI result information for Infrastructure as Code** section, move the slider to **Enabled**, and click the **Save changes** button.
