# Writing a rule

## Rules in Rego

Rules are written in Rego. When you are writing Rego, you do two things:

1. Write rules that make policy decisions. A rule is a conditional assignment.
2. Organize rules into policies. A policy is a set of rules with a hierarchical name

To learn more about the Policy Language, visit the [OPA Policy Language Documentation Page](https://www.openpolicyagent.org/docs/latest/policy-language/).

{% hint style="info" %}
You can also use the [OPA Playground](https://play.openpolicyagent.org) to try out Rego or run examples of this guide.
{% endhint %}

## How to generate a new rule

There are two options to get started:

1.  Use the `template` command to generate the required files for writing a rule:

    ```
    snyk-iac-rules template --rule <RULE-NAME> --format <hcl2|json|yaml|tf-plan>
    ```

    This generates the scaffolding for the rule, including fixture files based on the provided configuration format. For more details, read the [documentation about the template command](../sdk-reference.md#template-options).
2. Create a Rego policy from scratch and match the expected file and folder structure on your own:\
   `rules`\
   `└── my_rule`\
   `├── main.rego`\
   `└── main_test.rego`

{% hint style="info" %}
You will have to write your own Rego testing framework if you do not use the `template`command.
{% endhint %}

## Structure of the rule

In Rego, you can write statements that allow or deny a request, such as:\
`allow { input.name == "alice" }` or `deny { input.name == "alice" }`

{% hint style="info" %}
If the **`template`** command was used to generate the rules, then the default entry point is **`rules/deny`**. To override it and use a different name from `deny`, refer to the section [Bundling Rules](bundling-rules.md).
{% endhint %}

This is what a generated skeleton of a deny rule looks like when you run `snyk-iac-rules template --rule NEW-RULE --format hcl2`:

{% code title="rules/NEW-RULE/main.rego" %}
```
package rules

deny[msg] {
	resource := input.resource.test[name]
	resource.todo
	msg := {
		# Mandatory fields
		"publicId": "NEW-RULE",
		"title": "Default title",
		"severity": "low",
		"msg": sprintf("input.resource.test[%s].todo", [name]),
		# Optional fields
		"issue": "",
		"impact": "",
		"remediation": "",
		"references": [],
	}
}
```
{% endcode %}

{% hint style="warning" %}
You must follow this format of the **msg** property to ensure the output is displayed correctly from the Snyk IaC CLI.
{% endhint %}

The attributes are:

* **publicId:** a naming convention unique to your team, such as COMPANY-001. This should not contain/start with “SNYK-” to differentiate it from the internal Snyk rules.
* **title:** a short title that should summarise the issue.
* **severity:** this can be one of low/medium/high/critical.
* **msg:** Snyk recommends changing only the resource name and property, for example,`input.aws_s3_bucket[%s].tags` to match your example. The function `sprintf` is provided by Rego and enables Snyk to provide a dynamic error message explaining exactly where the issue was found.

The following attributes are optional but can be used to enhance the scan results in the Snyk CLI:

* **issue:** a more detailed string explanation of what the exact issue is.
* **impact:** a more detailed string explanation of what the impact of not resolving this issue is.
* **remediation:** a more detailed string explanation of how to resolve the issue. Snyk recommends providing a code snippet here.
* **references:** you can provide an array of strings with URLs to documentation

The generated test for the rule uses two generated Terraform files to verify if the correct `msg` field is returned by the rule for allowed and denied fixtures:

```
package rules

import data.lib
import data.lib.testing

test_NEW_RULE {
	# array containing test cases where the rule is allowed
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]

	# array containing cases where the rule is denied
	denied_test_cases := [{
		"want_msgs": ["input.resource.test[denied].todo"], # verifies that the correct msg is returned by the denied rule
		"fixture": "denied.tf",
	}]

	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("NEW-RULE", "./rules/NEW-RULE/fixtures", test_cases)
}
```

## Example of a rule

{% hint style="info" %}
For more examples, see[ Custom Rules Examples](examples-of-iac-custom-rules.md).
{% endhint %}

For this example, the templated rule was modified to assign a `msg` when a resource does not have an `owner` tag:

{% code title="rules/MY_RULE/main.rego" %}
```
package rules

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    not resource.tags.owner
	
    msg := {
        "publicId": "MY_RULE",
        "title": "Missing an owner from tag",
        "severity": "medium",
        "msg": sprintf("input.resource.aws_redshift_cluster[%s].tags", [name]),
        "issue": "",
        "impact": "",
        "remediation": "",
        "references": [],
    }
}
```
{% endcode %}

## Limitations and notes

* As Snyk compiles Rego policies into Wasm modules, you can only use built-in functions that support Wasm. There is a table at the bottom of the [Policy Reference Documentation](https://www.openpolicyagent.org/docs/latest/policy-reference/) that can help you identify those.
* A rule may be defined multiple times with the same name, either in a file or in separate files under the same package, for example,

```
packages rules

deny[msg] {
    resource.this
}
...

deny[msg] {
    resource.that
}
...
```

These rules are referred to as `incremental` because each definition is additive. You can read more about Incremental Definitions in the [Policy Reference Documentation](https://www.openpolicyagent.org/docs/latest/policy-language/#incremental-definitions). Note that these same named rules have to return a different value, or OPA will return an error. You can read more about Complete Definitions in the [Policy Reference Documentation](https://www.openpolicyagent.org/docs/latest/policy-language/#complete-definitions).

For more complex topics, check [how OPA resolves Conflict Resolution](https://www.openpolicyagent.org/docs/latest/faq/#conflict-resolution).
