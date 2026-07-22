# Policies & issues

Policies govern how AI is used across your environment. When an asset violates a policy, Evo raises an issue that you triage and resolve in your standard security workflow.

## Policies

A policy has a name, a severity (Critical, High, Medium, or Low), one or more conditions matched against asset attributes, and optional remediation steps. A policy supports up to 12 conditions, one per attribute.

Create and edit policies from the Policies & issues page or with Evo chat. You can delete custom policies. Default policies are read-only.

### Default policies

Evo provides default policies that raise high or critical severity risks as issues with no setup. AI-SPM and Agent Supply Chain Security each include default policies.

### User-defined policies

You can also create custom policies. With custom policies, you target assets by their attributes — for example, disallow a specific model in your code, or disallow a specific MCP server in your code, on end users' machines, or both.

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

<table><thead><tr><th width="374">Field</th><th>Description</th></tr></thead><tbody><tr><td><strong>Policy name</strong></td><td>The name of the policy.</td></tr><tr><td><strong>Severity</strong></td><td>The severity level assigned to issues created by this policy: <strong>Critical</strong>, <strong>High</strong>, <strong>Medium</strong>, or <strong>Low</strong>.</td></tr><tr><td><strong>Conditions (match all)</strong><br></td><td><p>Add a condition for assets under the policy.</p><p>Use the dropdown menus to select an attribute, a condition, and a value.</p><p>Click <strong>+Add condition</strong> to add another condition to the same policy. You can create up to 12 conditions, one per attribute.</p></td></tr><tr><td><strong>Remediation steps</strong> <strong>(optional)</strong></td><td>Add remediation advice.</td></tr></tbody></table>

#### Delete a policy

{% hint style="warning" %}
You cannot delete default policies.
{% endhint %}

You can delete user-defined policies. To quickly identify critical enforcement, you can group policies by severity.

To delete a policy:

1. Select the policy you want to delete.
2. Click the ellipsis next to **Edit policy**.
3. Click **Delete**.

## Issues

An issue is a policy violation. View issues on the Policies & issues page under Issues, or on an asset in Inventory to see them in context.

Each issue shows its severity, the asset that triggered it, remediation advice, and the number of occurrences. The remaining details vary by issue type.
