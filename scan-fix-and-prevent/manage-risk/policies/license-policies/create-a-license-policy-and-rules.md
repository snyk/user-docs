# Create a license policy and rules

{% hint style="info" %}
**Feature availability**

License policies are available only with Enterprise plans and apply only to Snyk Open Source scans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

Customers can configure the following settings for each license:

* Level of severity: Values include `None`, `Low`, `Medium`, and `High`.
  * New licenses added by Snyk default to a severity of `None` and **do not inherit** the `Unknown` license severity.&#x20;
  * When `None` is selected, instructions cannot be inserted since licenses marked with `None` **do not appear** in any Snyk test results.

## View and edit license policies

To update license severities under your License Policies:

Navigate to **Policies** > **Policy manager** > **License policies** > **Organization** and choose the policy you want to edit.

The **License policy** screen displays a list of licenses and their currently configured severities.

You can edit the license severities and instructions.

## Assign rules and severities for a license policy

1. In the **Policy manager**, navigate to **License policies** > **Organization**, and choose a policy link to open the **License policy** screen.
2. To set the severity for specific licenses, click the **Severity** selector on the **License policy** screen, and choose a **Severity** level to define which license issues you want to identify when Snyk tests run.\
   If you select a severity other than **None** and want additional instructions or fix recommendations to appear when that license issue is identified, select the instructions icon to the right of the **Severity** dropdown and enter the text for the license instruction.
3. Click **Add** or **Update** to confirm your changes.
4. Click **Submit** to save your policy.

The updated severities or instructions or both are automatically updated on Snyk servers. When the next scheduled test runs or when a user re-tests a Project, updated results are delivered according to these changes. See [License policy results](license-policy-results.md) for details.

## Set severity for the Unknown license type

The `Unknown` license type indicates Snyk has not recognized a license for a given package version. You can set a severity level for `Unknown`, for example, to ensure these license issues are given a higher severity and so appear more prominently in results.

Scroll down the licenses list on the right of the policy you are editing, then select the Severity dropdown for the **Unknown** license type.

{% hint style="info" %}
When `None` is selected, instructions cannot be inserted since licenses marked with `None` do not appear in any Snyk test results.
{% endhint %}
