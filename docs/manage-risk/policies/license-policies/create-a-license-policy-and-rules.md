# Create a license policy and rules

{% hint style="info" %}
**Feature availability**

License policies are available only with Enterprise plans and apply only to Snyk Open Source scans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Group administrators can configure the following settings for each license:

* Level of severity: Values include `None`, `Low`, `Medium`, and `High`.
  * When `None` is selected, instructions cannot be inserted since licenses marked with `None` do not appear in any Snyk test results.
  * New licenses added by Snyk inherit the `Unknown` license type severity. In cases where this severity is not set to `None`, newly added licenses appear in Snyk test results.
* Legal instructions for developers: Enter free text to provide any necessary instructions for developers.
  * Snyk recommends describing the specific policy of your company, explaining the need for collaboration from your developers as well as providing them with step-by-step instructions if any are needed
  * Legal instructions will appear in the CLI results, in PR Checks, and on issue cards in the Project view

## View and edit license policies

If your company account has one or more Groups, follow these steps to view and edit the License policy settings:

1.  Use the **Switch Group** selector and choose a Group to open its overview.

    <figure><img src="../../../.gitbook/assets/license_choose-group_19oct2022.png" alt="Switch Group"><figcaption><p>Switch Group</p></figcaption></figure>
2.  Navigate to **Policies** > **Policy manager** > **License policies** > **Organization** and choose the policy you want to edit.

    <figure><img src="../../../.gitbook/assets/policy_license_18oct2022.png" alt="Choose the policy to update"><figcaption><p>Choose the policy to update</p></figcaption></figure>

The **License policy** screen displays the list of Organizations that the policy applies to, a policy description, and the licenses included in the policy.

You can edit the license severities and instructions.

<figure><img src="../../../.gitbook/assets/choose-org_customize_19oct2022.png" alt="License policy overview."><figcaption><p>License policy overview</p></figcaption></figure>

## Assign rules and severities for a license policy

1. In the **Policy manager**, navigate to **License policies** > **Organization**, and choose a policy link to open the **License policy** screen.
2. To set the severity for specific licenses, click the **Severity** selector on the **License policy** screen, and choose a **Severity** level to define which license issues you want to identify when Snyk tests run.\
   If you select a severity other than **None** and want additional instructions or fix recommendations to appear when that license issue is identified, select the instructions icon to the right of the **Severity** dropdown and enter the text for the license instruction.
3. Click **Add** or **Update** to confirm your changes.
4. Click **Submit** to save your policy.

<figure><img src="../../../.gitbook/assets/policy-severity-instructions-x_06oct2022.png" alt="Update license policy instructions"><figcaption><p>Update license policy instructions</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/policy-severity-instructions-2_06oct2022.png" alt="Submit policy instructions"><figcaption><p>Submit policy instructions</p></figcaption></figure>

The updated severities or instructions or both are automatically updated on Snyk servers. When the next scheduled test runs or when a user re-tests a Project, updated results are delivered according to these changes. See [License policy results](license-policy-results.md) for details.

## Set severity for  the Unknown license type

The `Unknown` license type indicates Snyk has not recognized a license for a given package version. You can set a severity level for **Unknown**, for example, to ensure these license issues are given a higher severity and so appear more prominently in results.

Scroll down the licenses list on the right of the policy you are editing, then select the Severity dropdown for the **Unknown** license type:

<div align="left"><figure><img src="../../../.gitbook/assets/Screenshot 2023-05-12 at 10.42.12.png" alt="Set severity for Unknown license type"><figcaption><p>Set severity for Unknown license type</p></figcaption></figure></div>

{% hint style="info" %}
Snyk applies the severity level set for **Unknown** to new licenses when Snyk adds them to your license policy.
{% endhint %}
