# Create a Custom Rule

With the Snyk CLI, you can create a template for a Custom Rule, which includes a file titled `main.rego`. Using Rego, you can populate the rule that can then be consistently applied across the Snyk Platform. When the rule is finished, it can look like this example rule for enforcing a password policy:&#x20;

{% code title="rules/NEW_PASSWORD_POLICY/main_test.rego" %}
```
package rules.NEW_PASSWORD_POLICY

import data.snyk

input_type := "cloud_scan"

metadata := {
	"id": "NEW_PASSWORD_POLICY",
	"severity": "high",
	"title": "Increase the number of characters in our password policy",
	"description": "All of our password policies now require a minimum of 17 characters instead of the CIS recommendation of 14 characters",
	"product": ["cloud"],
}

password_pol := snyk.resources("aws_iam_account_password_policy")[_]

deny[info] {
	count(password_pol) < 1 
	info := {
		"message": "This account does not contain a password policy",
		"resource": password_pol
		}
}

deny[info] {
	password_pol.minimum_password_length < 17
	info := {"resource": password_pol}
}

resources[info] {
	info := {"resource": password_pol}
}
```
{% endcode %}

## Creating a Custom Rule template

Use the following CLI command:

```
snyk iac rules init
```

This command introduces a series of interactive prompts and allows you to set up your Project, rule, rule spec (for testing), and resource relations. ​​By selecting `rule`, you will get a series of interactive prompts to create the template for your rule as a file named `main.rego`.

**Example output from interactive prompts:**

```
What do you want to initialize? rule
Rule ID: NEW_PASSWORD_POLICY
Title: Increase the number of characters in our password policy
Severity: high
Description: All of our password policies now require a minimum of 17 characters instead of the CIS recommendation of 14 characters
Is this rule intended for Snyk IaC, Snyk Cloud, or both? cloud
Input type: tf
Does this rule need more than one resource type? No
Resource type: aws_iam_account_password_policy
```

## Writing a Custom Rule

After the template is created, the rule’s logic must be defined, and the rule must be written in [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/). For details about additional metadata that can be added to your rule, see the [metadata page](https://github.com/snyk/policy-engine/blob/main/docs/policy\_spec.md#metadata) in the Snyk policy engine repository.
