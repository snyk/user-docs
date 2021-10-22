# Writing a rule

### How to generate a new rule

There are two options to get started:

1.  Use the `template` command to generate the required files for writing a rule:

    ```
    snyk-iac-rules template --rule <RULE-NAME>
    ```

    This generates the scaffolding for the rule. For more details, read the [documentation about the template command](../sdk-reference.md#template-options).
2. Create a Rego policy from scratch and match the expected file and folder structure on your own (**Note**: you will have to write your own Rego testing framework if you don't use the `template`: command)\
   `rules `\
   `└── my_rule `\
   `         ├── main.rego  `\
   `        └── main_test.rego`

### Structure of the rule

In Rego, you can write statements that allow or deny a request, such as:\
`allow { input.name == "alice" }` or `deny  { input.name == "alice" }`

{% hint style="info" %}
If the **`template`** command was used to generate the rules, then the default entry point is **`rules/deny`**. To override it and use a different name than `deny`, check the section [Bundling Rules](bundling-rules.md).
{% endhint %}

This is what a generated skeleton of a deny rule looks like when we run `snyk-iac-rules `template` --rule my_rule`:

{% code title="rules/my_rule/main.rego" %}
```
package rules

deny[msg] {
    input.spec.template.todo
    msg := {
        "publicId": "my_rule",
        "title": "Default title",
        "severity": "low",
        "issue": "",
        "impact": "",
        "remediation": "",
        "msg": "spec.template.todo",
        "references": [],
    }
}
```
{% endcode %}

{% hint style="warning" %}
You must follow this format of the **msg** property to ensure the output correctly displays on the Snyk IaC CLI.
{% endhint %}

The attributes are:

* **publicId:** a naming convention unique to yourselves, such as COMPANY-001. **This should not contain/start with “SNYK-”** to differentiate from the internal Snyk rules.
* **title:** a short title that should summarise the issue.
* **severity:** this can be one of **low/medium/high/critical.**
* **msg:** we recommend only changing the resource name and property e.g. `aws_s3_bucket[%s].tags` to match your example. The function `sprintf` is provided by Rego and enables us to provide a dynamic error message explaining exactly where the issue was found.

The following attributes are optional but can be used to enhance the scan results in the Snyk CLI:

* **issue:** a more detailed string explanation of what the exact issue is.
* **impact:** a more detailed string explanation of what the impact of not resolving this issue is.
* **remediation:** a more detailed string explanation of how to resolve the issue. We recommend providing a code snippet here.
* **references:** you can provide an array of strings with URLs to documentation

### Example of a rule

{% hint style="info" %}
For more examples, see[ Custom Rules Examples](examples.md).
{% endhint %}

For this example, we modified our templated rule to assign a `msg` when a resource does not have an `owner` tag:

{% code title="rules/my_rule/main.rego" %}
```
package rules

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    not resource.tags.owner

    msg := {
        "publicId": "my_rule",
        "title": "Missing an owner from tag",
        "severity": "medium",
        "issue": "",
        "impact": "",
        "remediation": "",
        "msg": sprintf("resource.aws_redshift_cluster[%s]", [name]),
        "references": [],
    }
}
```
{% endcode %}

### Limitations/Notes

* As we compile Rego policies into Wasm modules, you can only use built-in functions that support Wasm. There is a table at the bottom of the [Policy Reference Documentation](https://www.openpolicyagent.org/docs/latest/policy-reference/) that can help you identify those.
* A rule may be defined multiple times with the same name, either in a file, or in separate files under the same package, e.g:&#x20;

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

These rules are referred as `incremental` as each definition is additive. You can read more about Incremental Definitions [here](https://www.openpolicyagent.org/docs/latest/policy-language/#incremental-definitions). Note that these same named rules have to return a different value, or OPA will return an error. You can read more about Complete Definitions [here](https://www.openpolicyagent.org/docs/latest/policy-language/#complete-definitions).&#x20;



For more complex topics, check [how OPA resolves Conflict Resolution](https://www.openpolicyagent.org/docs/latest/faq/#conflict-resolution).&#x20;
