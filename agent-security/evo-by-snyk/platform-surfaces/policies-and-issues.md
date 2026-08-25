---
nav_context: classic
---

# Policies & issues

Policies govern how AI is used across your environment. When an asset violates a policy, Evo raises an issue that you triage and resolve in your standard security workflow.

## Policies

A policy has a name, a severity (Critical, High, Medium, or Low), one or more conditions matched against asset attributes, and optional remediation steps. A policy supports up to 12 conditions, one per attribute.

Severity is fixed per policy. A policy does not change severity based on the score it matches, so to raise different severities at different scores, create a policy for each.

Create and edit policies from the Policies & issues page or with Evo chat. You can delete custom policies. Default policies are read-only.

{% hint style="info" %}
Users with a Tenant role and full Evo access can create, edit, and delete policies. Users restricted to specific Organizations can read policies and issues, but they cannot change them. Visit [Access and authentication](../access-and-authentication.md) for more information.
{% endhint %}

### Default policies

Evo provides default policies that raise high or critical severity risks as issues with no setup. AI-SPM and Agent Supply Chain Security each include default policies.

Snyk default policies for model risk target attacker goals, the most specific level of the model risk taxonomy. They raise:

* a **High** issue from a score of 500
* a **Critical** issue from a score of 750

{% hint style="info" %}
A score in a severity range does not always raise an issue. If no default policy exists for that attacker goal at that severity, the score appears on the model's **Risk profile** tab with no issue attached. You can create your own policy against any attacker goal or sub-category at any threshold.
{% endhint %}

To learn how model risk scores are calculated, visit [Risk intelligence](../ai-spm/risk-intelligence/).

### User-defined policies

You can also create custom policies. With custom policies, you target assets by their attributes — for example, disallow a specific model in your code, or disallow a specific MCP server in your code, on end users' machines, or both.

For model risk, you can target an attacker's goal or a sub-category. Impact categories cannot be targeted.

Target a sub-category when its attacker goals are interchangeable for you: the same fix, the same owner, and the specific goal does not change what you do about it. Bias is the common example, where you care that a model shows bias rather than whether it is gender, race, or religious bias.

Avoid pairing an attacker's goal with the sub-category that contains it. Both fire, and you get two issues for one problem.

#### Create a policy

You can create a policy from the **Policies** page or using Evo chat.

<details>

<summary>Create a policy from the <strong>Policies</strong> page</summary>

To create a policy from the **Policies** page:

1. Click **Create policy**.
2. Complete the form fields.
   1. Use the dropdown menus to select an attribute, a condition, and a value.
   2. Click **+Add condition** to add another condition to the same policy. You can create up to 12 conditions, one per attribute.
3. Click **Create**.

</details>

<details>

<summary>Create a policy using Evo chat</summary>

To create policies using Evo chat, you can ask it to create a policy with specific criteria. For example, you can ask it to "create a policy that raises a critical issue when you detect _\[model name]_".

The **Policy agent** then navigates to the policy creation form. Ensure the policy details are correct and click **Create** to create the policy.

Based on the policy, Evo evaluates scan results and creates issues when matches occur.

A newly created policy produces issues immediately after creation.

</details>

#### Edit a policy

From the **Policies** page, select the policy you want to edit and click **Edit policy**. The following fields are available:

<table><thead><tr><th width="259.921875">Field</th><th>Description</th></tr></thead><tbody><tr><td><strong>Policy name</strong></td><td>The name of the policy.</td></tr><tr><td><strong>Severity</strong></td><td>The severity level assigned to issues created by this policy: <strong>Critical</strong>, <strong>High</strong>, <strong>Medium</strong>, or <strong>Low</strong>.</td></tr><tr><td><strong>Conditions (match all)</strong><br></td><td><p>Add a condition for assets under the policy.</p><p>Use the dropdown menus to select an attribute, a condition, and a value.</p><p>Click <strong>+Add condition</strong> to add another condition to the same policy. You can create up to 12 conditions, one per attribute.</p></td></tr><tr><td><strong>Remediation steps</strong> <strong>(optional)</strong></td><td>Add remediation advice.</td></tr></tbody></table>

#### Delete a policy

{% hint style="warning" %}
You cannot delete default policies.
{% endhint %}

You can delete user-defined policies. To quickly identify critical enforcement policies, you can group them by severity.

To delete a policy:

1. Select the policy you want to delete.
2. Click the ellipsis next to **Edit policy**.
3. Click **Delete**.

#### Policies created against the Risk index

The Model Risk Score replaces the Risk index. Policies you created against a Risk index category stop raising issues and need to be recreated.

{% hint style="warning" %}
The five Risk index categories do not map onto the new taxonomy, and the two scores are not equivalent. You cannot carry a Risk index threshold across. Set the threshold again against the attacker's goal or sub-category you want to govern.
{% endhint %}

For what the new taxonomy contains and how to read a score, visit [Risk intelligence](../ai-spm/risk-intelligence/).

## Issues

An issue is a policy violation. View issues on the Policies & issues page under Issues, or on an asset in Inventory to see them in context.

Each issue shows its severity, the triggering asset, remediation advice, and the number of occurrences. The remaining details vary by issue type.
