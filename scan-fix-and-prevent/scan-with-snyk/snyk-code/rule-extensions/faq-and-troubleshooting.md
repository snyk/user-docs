---
description: FAQ and troubleshooting guidance for Snyk Code Rule Extensions, including permissions and API rate limits
---

# FAQ & troubleshooting

This page provides answers to common questions about using Rule Extensions, along with troubleshooting guidance and permission requirements.

## General information

### What entitlements and permissions do I need to use this feature?

You must be an Enterprise customer, and you need the appropriate permissions on a [custom role](https://docs.snyk.io/snyk-platform-administration/user-roles/user-role-management). The exact set depends on whether you access Rule Extensions through the API or the in-product UI — see [Rule Extensions permissions](rule-extensions-permissions.md) for the full breakdown.

### What are the API rate limits?

Impact test requests — creating a test and retrieving its results — use a lower rate-limit bucket: **5 requests per second, 50 per minute, and 500 per hour**. All other Rule Extensions API endpoints use the standard Snyk REST API limits, which are considerably higher (currently 160 per second and 1,620 per minute). Exceeding a limit returns `429 Too Many Requests`.

### Which languages are supported by Rule Extensions?

Rule Extensions support most of the same languages as Snyk Code tests. The [Supported rules](supported-rules.md) page shows detailed support for combinations of rules, languages, and Rule Extensions. For more information on Snyk Code language support, refer to the [documentation](https://docs.snyk.io/supported-languages-package-managers-and-frameworks). For information about the roadmap, reach out to your account team.

## Migrating from the closed beta

If you used Rule Extensions during the closed beta, complete these steps for general availability:

* **Move to the GA API endpoints.** The closed-beta API endpoints will be retired 30 days after general availability. The GA endpoints are available now — see the [API reference](https://apidocs.snyk.io/?version=2026-03-25#get-/groups/-group_id-/sast/rule_extensions).
* **Update your custom-role permissions.** Grant your roles the `rule_extension.*` permissions they need, including **View Groups** and **View Organizations**. See [Rule Extensions permissions](rule-extensions-permissions.md).

## Rule Extensions

### How do I know if a Rule Extension is working, or the impact it is having?

You can preview the impact of a Rule Extension in the Snyk Web UI or with the [impact testing API](impact-testing.md). Use the smallest practical project for the test.

### What is the delay between publishing a Rule Extension and being able to see the results in a Code scan?

The maximum delay is nine (9) hours.

{% hint style="warning" %}
A delay of up to 9 hours can occur before a Rule Extension is recognized when Snyk scans an existing commit hash. To force the change to be picked up sooner, push a new commit to the repository.
{% endhint %}

### Can I apply the Rule Extension to a Group or an Organization?

A Rule Extension only affects scans where it is published and has an assignment to that scope. Use the [Assignments API](https://apidocs.snyk.io/?version=2026-03-25#post-/groups/-group_id-/rule_extensions/assignments) to assign a published extension to the whole Group or to specific Organizations under the Group.

### Is there a limit to the number of Rule Extensions I can create?

You can create a maximum of 1,000 published Rule Extensions for a Group.

### Can I update multiple extensions at the same time?

This functionality is not currently available.

### Can I delete multiple extensions?

You delete one Rule Extension at a time. Published extensions can only be deleted after all assignments are removed (see the [Assignments API](https://apidocs.snyk.io/?version=2026-03-25#post-/groups/-group_id-/rule_extensions/assignments)).

### What corner cases can stop a sanitizer from being recognized?

A Rule Extension only takes effect when Snyk Code can resolve your function to its FQN in the data flow. The most common reasons it doesn't:

* **Wildcard imports** (Java, Kotlin, Scala, C#) — `import pkg.*` prevents resolution. Use explicit imports.
* **Relative imports** (JavaScript/TypeScript) — `import { clean } from '../utils'` does not resolve. Use an absolute module path or the npm package name.
* **Implicit typing** — with type inference (for example, `val v = Validator()`), the engine cannot bind an instance method. Use an explicit type.
* **Missing `global::` prefix** (C#, VB.NET) — the FQN must start with `global::`.
* **Python decorators** — a function applied as a decorator (`@my_sanitizer`) cannot be used as a sanitizer.
* **Go factory or interface returns** — methods on structs returned by interface factories often fail to resolve. Instantiate the struct directly.
* **Chained or dynamic calls** (PHP, Ruby) — the FQN may resolve to the method name alone. Confirm with an impact test.
* **Non-public functions** — sanitizers must be public or exported.

See [Identify your sanitizer's FQN](identify-your-sanitizers-fqn.md) for per-language detail, and validate any uncertain case with [impact testing](impact-testing.md).

### How do I choose the right sanitizer type?

Match the type to how your function behaves — Flow Through (the return value is always clean), If True / If False (checked in a condition), or Any Usage (mutates in place or throws). Choosing **Any Usage** for a function that does not sanitize on every path can create **false negatives** by hiding a real vulnerability. When in doubt, run an impact test before publishing. See [Custom sanitizers](custom-sanitizers.md#types-of-sanitizers).

## Rule management

### What is the difference between `published` and `draft` for Rule Extensions?

* **Published** means the extension can be assigned and applied to Snyk Code scans (together with [assignments](https://apidocs.snyk.io/?version=2026-03-25#post-/groups/-group_id-/rule_extensions/assignments)).
* **Draft** keeps the configuration saved but it is not applied to scans until you publish it.

### Can published Rule Extensions be changed to draft?

Published Rule Extensions cannot be changed to draft. Create a new Rule Extension and delete the one that's no longer needed.

### How do I publish my draft Rule Extensions?

They can be created as published using the [Create a Rule Extension](https://apidocs.snyk.io/?version=2026-03-25#post-/groups/-group_id-/sast/rule_extensions) endpoint, or created as `draft` and later published with the [Update a Rule Extension by Rule Extension ID](https://apidocs.snyk.io/?version=2026-03-25#patch-/groups/-group_id-/sast/rule_extensions/-rule_extension_id-) endpoint by setting `attributes.status` to `published`.

### What data can I update for a rule?

The [Update a Rule Extension by Rule Extension ID](https://apidocs.snyk.io/?version=2026-03-25#patch-/groups/-group_id-/sast/rule_extensions/-rule_extension_id-) endpoint supports partial updates to `description`, `configuration` (sanitization `type` and `rule_keys`), and `status`. For **draft** extensions, you can also update `signature` (including `fully_qualified_name`). For **published** extensions, you **cannot** update `signature`.

### What do the different sanitizer types mean?

Sanitizer type refers to the way the input data is being sanitized. Snyk has identified four major types that are supported in Snyk scans, covering the different usage patterns. See [Custom sanitizers](custom-sanitizers.md).

### Can Snyk recover deleted Rule Extensions?

Deleted Rule Extensions cannot be recovered. If a Rule Extension is deleted, you must create a new Rule Extension.

### Can I change the fully qualified function name?

* For draft extensions, you can use the [Update a Rule Extension by Rule Extension ID](https://apidocs.snyk.io/?version=2026-03-25#patch-/groups/-group_id-/sast/rule_extensions/-rule_extension_id-) endpoint `signature` attribute to specify a new FQN.
* For published extensions, the signature cannot be updated. Create a new Rule Extension with the desired FQN and remove the old one when it is no longer needed.

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

### Will Rule Extensions be available across Groups (tenants) in the future?

This is a future consideration and is not currently supported.

## Troubleshooting

### Why is my code scan timing out?

Try running the test again, and if the issue persists, contact your account team or open a case with [Snyk support](https://support.snyk.io/s/).

### I accidentally deleted my rule; can I get it back?

Deleted Rule Extensions cannot be recovered. You must create a new one.

### I added a custom sanitizer but the false positives still show up; what should I do?

Snyk supports fully qualified function names or sanitizers that are uniquely identifiable across a repository or Project. If you are still having issues after ensuring your sanitizer meets these criteria, reach out to your account team or open a case with [Snyk support](https://support.snyk.io/s/).

For further assistance, contact your account team or refer to the [Snyk documentation](https://docs.snyk.io/).

## Glossary

| Term                  | Definition                                                                                                                     |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| **FQN**               | Fully Qualified Name — the complete identifier for a function including namespace, class, and method name.                     |
| **Rule Key**          | The unique identifier of a Snyk Code rule that the extension applies to — the value used in the `rule_keys` API field. See [Supported rules](supported-rules.md). |
| **Sanitizer**         | Name of the sanitizing function being added to Snyk Code rules.                                                                |
| **Sanitization Type** | Specifies the expected behavior of the sanitizing function. See [Custom sanitizers](custom-sanitizers.md) for details.         |
| **Scope**             | Where a published extension applies, as defined by [assignments](https://apidocs.snyk.io/?version=2026-03-25#post-/groups/-group_id-/rule_extensions/assignments) to the Group or Organizations. |
| **Status**            | Specifies whether the Rule Extension has been saved as a draft or published for test execution.                                |
