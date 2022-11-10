# Setting a license policy

Administrators can configure the following settings for each license:

* The level of severity— values include **None**, **Low**, **Medium**, and **High**.
  * When **None** is selected, instructions cannot be inserted since licenses marked with **None** do not appear in any Snyk test results.
  * New licenses added by Snyk inherit the **Unknown** license type severity. In cases where this severity is not set to **None**, newly added licenses  appear in Snyk test results.
* **Legal instructions for developers**—enter free text to provide any necessary instructions for developers.
  * We recommend describing your company’s specific policy, explaining the need for collaboration from your developers as well as providing them step-by-step instructions if any are needed
  * Legal instructions will appear in the CLI results and on issue cards within the project view

{% hint style="info" %}
**Feature availability**\
For customers with a Business plan: if no group is available, org administrators can make license policy changes; if a group is available, group administrators can make license policy changes.&#x20;

For customers with an Enterprise plan, group administrators can create or modify a policy. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

## Access license policies&#x20;

If your __ company's account has one or more groups, to access the License policy settings: &#x20;

1.  Use the **Switch group** selector and choose a Group to open its overview. \
    &#x20;

    <figure><img src="../../../.gitbook/assets/license_choose-group_19oct2022.png" alt=""><figcaption></figcaption></figure>
2.  Go to **Policies > Policy manager > License polices > Organization** and choose the policy you want to update.\


    <figure><img src="../../../.gitbook/assets/policy_license_18oct2022.png" alt=""><figcaption></figcaption></figure>

The **License policy** screen displays the list of Organizations that the policy applies to, a policy description, and the licenses included in the policy.&#x20;

You can edit the license severities and instructions.&#x20;

![](../../../.gitbook/assets/choose-org\_customize\_19oct2022.png)

## Assign rules and severities for a license policy

1. In **Policy manager > License polices > Organization**, choose a policy link to open the **License policy** screen.&#x20;
2. To set the severity for specific licenses, click the **Severity** selector in the **License policy** screen.&#x20;
3. To enter an explanation and recommendations for fixes, click the **Instructions** icon (to the right of the **Severity** selector) and enter your text.
4.  Click **Add** to save your changes to the instructions

    Once added, the **Add instructions** link changes to **Edit instructions.**
5. Click **Submit** to save your policy.

<figure><img src="../../../.gitbook/assets/policy-severity-instructions-x_06oct2022.png" alt=""><figcaption><p>Open the <strong>Policy</strong> modal, set severity, and add or update instructions</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/policy-severity-instructions-2_06oct2022.png" alt=""><figcaption><p>Add or update instructions and <strong>Submit</strong> to save your policy changes</p></figcaption></figure>

The updated severities or instructions (or both) are automatically updated on Snyk servers. When the next scheduled test runs, or when a user re-tests a project, updated results are delivered according to these changes.
