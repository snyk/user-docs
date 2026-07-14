# FAQ & troubleshooting

This page provides answers to common questions about using Rule Extensions, along with troubleshooting guidance and permission requirements.

## General information

### What entitlements and permissions do I need to use this feature?

You must be an Enterprise customer with API entitlement and have the appropriate permissions configured through a custom role.

#### Required permissions

Rule Extensions are **managed at the Group level** and **tested at the Organization level**. Access is granted through a [custom role](https://docs.snyk.io/snyk-platform-administration/user-roles/custom-role-templates): create or edit a role and add only the permissions for the actions that the role needs to perform.

You do **not** need every permission listed below. Grant only what the role requires:

* A read-only reviewer needs **View Rule Extensions**, **View Groups**, and **View Organizations**.
* A user who manages extensions also needs the **Create**, **Edit**, and/or **Delete** permissions.
* A user who runs impact tests needs the Organization-level **Test Rule Extensions** permission.

When you configure the custom role in the Snyk Web UI, the Group-level permissions appear under the **Rule Extensions Management** permission group.

**Group-level permissions**

These permissions control who can manage rule extensions for the Group. **View Rule Extensions** is the base permission: the Create, Edit, and Delete permissions each also require it.

| Permission name        | Permission key                | Grants                                                                                                                            |
| ---------------------- | ----------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| View Rule Extensions   | `group.rule_extension.read`   | View the rule extensions in the Group. Required by all of the permissions below.                                                  |
| Create Rule Extensions | `group.rule_extension.create` | Create new rule extensions in the Group.                                                                                          |
| Edit Rule Extensions   | `group.rule_extension.edit`   | Update existing rule extensions in the Group.                                                                                     |
| Delete Rule Extensions | `group.rule_extension.delete` | Permanently delete rule extensions from the Group.                                                                                |
| View Groups            | `group.read`                  | View the Group that a rule extension belongs to. Required to manage Group-scoped rule extensions.                                 |
| View Organizations     | `group.org.list`              | See the Organizations a rule extension applies to. Required because rule extensions are scoped to Organizations within the Group. |

**Organization-level permission**

This permission is only needed to use the impact testing API, which runs a rule extension against a Project to preview its results.

| Permission name           | Permission key                          | Grants                                                                       |
| ------------------------- | --------------------------------------- | ----------------------------------------------------------------------------- |
| Test Rule Extensions | `org.rule_extension.project.test` | Run the impact test for a rule extension on a Project in the Organization. |

### What are the API rate limits?

The rate limits are 1 request per second, 10 requests per minute, and 100 requests per hour.

### Which languages are supported by rules extensions?

Rules extensions support most of the same languages as Snyk Code tests. The [Supported rules](supported-rules.md) page shows detailed support for combinations of rules, languages, and rule extensions. For more information on Snyk Code language support, refer to the [documentation](https://docs.snyk.io/supported-languages-package-managers-and-frameworks). For information about the roadmap, reach out to your account team.

## Rule extensions

### How do I know if a rule extension is working, or the impact it is having?

The best way to understand the impact of a rule extension is by using the [impact testing API](https://docs.snyk.io/developer-tools/snyk-api/reference/sastruleextensions#post-groups-group_id-sast-rule_extensions-tests). Use the smallest practical project for the test.

Snyk will be introducing **UI** functionality that allows you to test rule extension changes to a Project for the General Availability milestone.

### What is the delay between publishing a rule extension and being able to see the results in a Code scan?

The maximum delay is nine (9) hours.

{% hint style="warning" %}
A delay of up to 9 hours can occur before a rule extension is recognized when Snyk scans an existing commit hash. To force the change to be picked up sooner, push a new commit to the repository.
{% endhint %}

### Can I apply the rule extension to a Group or an Organization?

A rule extension only affects scans where it is published and has an assignment to that scope. Use the [Assignments API](https://docs.snyk.io/developer-tools/snyk-api/reference/assignments) to assign a published extension to the whole Group or to specific Organizations under the Group.

### Is there a limit to the number of rule extensions I can create?

You can create a maximum of 1,000 published rule extensions for a Group.

### Can I update multiple extensions at the same time?

This functionality is not currently available.

### Can I delete multiple extensions?

You delete one rule extension at a time. Published extensions can only be deleted after all assignments are removed (see the [Assignments API](https://docs.snyk.io/developer-tools/snyk-api/reference/assignments)).

## Rule management

### What is the difference between `published` and `draft` for rule extensions?

* **Published** means the extension can be assigned and applied to Snyk Code scans (together with [assignments](https://docs.snyk.io/developer-tools/snyk-api/reference/assignments)).
* **Draft** keeps the configuration saved but it is not applied to scans until you publish it.

### Can published rule extensions be changed to draft?

Published rule extensions cannot be changed to draft rule extensions. Create a new rule extension and delete the one that is no longer required.

### How do I publish my draft rule extensions?

They can be created as published using the [Create a rule extension](https://docs.snyk.io/developer-tools/snyk-api/reference/sastruleextensions#post-groups-group_id-sast-rule_extensions) endpoint, or created as `draft` and later published with the [Update a rule extension by rule extension ID](https://docs.snyk.io/developer-tools/snyk-api/reference/sastruleextensions#patch-groups-group_id-sast-rule_extensions-rule_extension_id) endpoint by setting `attributes.status` to `published`.

### What data can I update for a rule?

The [Update a rule extension by rule extension ID](https://docs.snyk.io/developer-tools/snyk-api/reference/sastruleextensions#patch-groups-group_id-sast-rule_extensions-rule_extension_id) endpoint supports partial updates to `description`, `configuration` (sanitization `type` and `rule_keys`), and `status`. For **draft** extensions, you can also update `signature` (including `fully_qualified_name`). For **published** extensions, you **cannot** update `signature`.

### What do the different sanitizer types mean?

Sanitizer type refers to the way the input data is being sanitized. Snyk has identified four major types that are supported in Snyk scans, covering the different usage patterns. See [Custom sanitizers](custom-sanitizers.md).

### Can Snyk recover deleted rule extensions?

Deleted rule extensions cannot be recovered. If a rule extension is deleted, you must create a new rule extension.

### Can I change the fully qualified function name?

* For draft extensions, you can use the [Update a rule extension by rule extension ID](https://docs.snyk.io/developer-tools/snyk-api/reference/sastruleextensions#patch-groups-group_id-sast-rule_extensions-rule_extension_id) endpoint `signature` attribute to specify a new FQN.
* For published extensions, the signature cannot be updated. Create a new rule extension with the desired FQN and remove the old one when it is no longer needed.

### Best practices for custom sanitizers

To ensure Snyk correctly identifies and applies your custom sanitizers, follow these guidelines for structuring your code.

#### Extract sanitizer logic to a named method

Snyk identifies sanitizers by their Fully Qualified Name (FQN). If your sanitization logic is written inline (for example, inside a larger function or as a lambda), Snyk cannot address it.

**Problematic code (inline logic)**

The sanitization logic is mixed with other code and has no unique addressable name.

{% code fullWidth="false" %}
```csharp
// Logic is inline and cannot be targeted by FQN
if (path.Contains("..")) {
    throw new Exception("Invalid path");
}
File.Delete(path);
```
{% endcode %}

**Recommended structure (extracted method)**

Extract the logic into a dedicated, named method.

{% code fullWidth="false" %}
```csharp
public class Sanitizer {
    // This method has an FQN: Namespace.Sanitizer.PathIsSafe
    public static bool PathIsSafe(string path) {
        return !path.Contains("..");
    }
}

// Usage
if (Sanitizer.PathIsSafe(path)) {
    File.Delete(path);
}
```
{% endcode %}

#### Avoid reconstructing arguments at the call site

When you pass a reconstructed value (like a concatenated string) directly to a sanitizer, the Snyk engine may treat the argument as a new, separate object from the one used in the sink (the dangerous function). This breaks the taint flow connection.

**Problematic code (reconstruction at call site)**

The string `addressLabelPath + pdf + ".pdf"` is constructed twice: once for the check and once for the delete. Snyk sees these as two different events.

{% code fullWidth="false" %}
```csharp
// The argument is reconstructed in the check
if (Sanitizer.PathIsSafe(addressLabelPath + pdf + ".pdf")) {
    // And reconstructed again in the sink
    System.IO.File.Delete(addressLabelPath + pdf + ".pdf");
}
```
{% endcode %}

**Recommended structure (assign to variable)**

Assign the constructed value to a variable first, then use that variable in both the sanitizer and the sink. This ensures Snyk tracks the same object.

{% code fullWidth="false" %}
```csharp
// Assign to a variable first
var fullPath = addressLabelPath + pdf + ".pdf";

// Use the variable in the check
if (Sanitizer.PathIsSafe(fullPath)) {
    // Use the same variable in the sink
    System.IO.File.Delete(fullPath);
}
```
{% endcode %}

## Upcoming plans and enhancements

### Will the API have Organization scope access in the future?

Snyk plans to revisit access and custom roles in the future.

### Will the rule extension scope be expanded to tenants in the future? That is, will we be able to create rule extensions across groups?

This is a future consideration and is not currently supported.

## Troubleshooting

### Why is my code scan timing out?

Try running the test again, and if the issue persists, contact your account team or open a case with [Snyk support](https://support.snyk.io/s/).

### I accidentally deleted my rule; can I get it back?

Deleted rule extensions cannot be recovered. You must create a new one.

### I added a custom sanitizer but the false positives still show up; what should I do?

Snyk supports fully qualified function names or sanitizers that are uniquely identifiable across a repository or Project. If you are still having issues after ensuring your sanitizer meets these criteria, reach out to your account team or open a case with [Snyk support](https://support.snyk.io/s/).

For further assistance, contact your account team or refer to the [Snyk documentation](https://docs.snyk.io/).

## Glossary

| Term                  | Definition                                                                                                                     |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| **FQN**               | Fully Qualified Name - the complete identifier for a function including namespace, class, and method name.                     |
| **Rule Name**         | Lists all rules that the extension applies to.                                                                                 |
| **Sanitizer**         | Name of the sanitizing function being added to Snyk Code rules.                                                                |
| **Sanitization Type** | Specifies the expected behavior of the sanitizing function. See [Custom sanitizers](custom-sanitizers.md) for details.         |
| **Scope**             | Where a published extension applies, as defined by [assignments](https://docs.snyk.io/developer-tools/snyk-api/reference/assignments) to the Group or Organizations. |
| **Status**            | Specifies whether the rule extension has been saved as a draft or published for test execution.                                |
