# Examples of IaC custom rules

## Example of a simple boolean rule

{% hint style="info" %}
You can find a full example of this guide in [this OPA Playground](https://play.openpolicyagent.org/p/SCYndBjWxh) and the [snyk/custom-rules-example](https://github.com/snyk/custom-rules-example) repository.
{% endhint %}

Assume you have generated a new rule, `CUSTOM-RULE-1` using the SDK , that is, `snyk-iac-rules template --rule CUSTOM-RULE-1` and have a very simple fixture file containing a Terraform resource:

{% code title="rules/CUSTOM-RULE-1/fixtures/denied.tf" %}
```
resource "aws_redshift_cluster" "denied" {
  cluster_identifier = "tf-redshift-cluster"
  node_type          = "dc1.large"
  tags = {
  }
}
```
{% endcode %}

Now, modify the generated Rego to enforce resources tagged with an owner:

1. Create a variable `[name]` to enumerate across all of the `aws_redshift_cluster` resources. This variable can be named anything you like, for example, `i`, `j`, `name`, and so on.
2. Store this into the resource variable by assigning the value to it with a walrus operator `:=`; e.g. `resource := input.resource.aws_redshift_cluster[name]`
3. Check whether the owner tag exists for each resource; to do that, check if the path `resource.tags.owner` is defined. If it is undefined, it will evaluate as undefined. So, use the `NOT` keyword in front of it, which will evaluate to `TRUE`; for example,`not resource.tags.owner`

The modified Rego is:

{% code title="rules/CUSTOM-RULE-1/main.rego" %}
```
package rules

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    not resource.tags.owner

    msg := {
        "publicId": "CUSTOM-RULE-1",
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

{% hint style="info" %}
To understand how the Rego code evaluates the Terraform file provided earlier, have a look at how the SDK is able to [parse a fixture file](parsing-an-input-file.md) into JSON.
{% endhint %}

{% hint style="info" %}
Snyk recommends always validating that your rule is correct by [updating and running the unit tests](testing-a-rule.md).
{% endhint %}

The test for this rule verifies that the Rego rule can identify that the fixture at the beginning of this guide is invalid:

{% code title="rules/CUSTOM-RULE-1/main_test.rego" %}
```
package rules

import data.lib
import data.lib.testing

test_CUSTOM_RULE_1 {
	# array containing test cases where the rule is allowed
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]

	# array containing cases where the rule is denied
	denied_test_cases := [{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied].tags"],
		"fixture": "denied.tf",
	}]

	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("CUSTOM-RULE-1", "./rules/CUSTOM-RULE-1/fixtures", test_cases)
}
```
{% endcode %}

## Example with logical AND

Try and extend the preceding example  and update the rule to allow all cases that satisfy two conditions:

1. A resource has an “owner” tag\
   **AND**
2. A resource has a “description” tag

To test this new condition, generate a new rule `CUSTOM-RULE-2` using the `template` command and write the following fixture file:

{% code title="rules/CUSTOM-RULE-2/fixtures/denied.tf" %}
```
resource "aws_redshift_cluster" "denied" {
  cluster_identifier = "tf-redshift-cluster"
  node_type          = "dc1.large"
  tags = {
    owner = "team-123"
  }
}
```
{% endcode %}

Joining multiple expressions together expresses logical `AND`.

* You can do this with the `;` operator.
* Or, you can omit the `;` (`AND`) operator by splitting expressions across multiple lines.

{% hint style="info" %}
The logical AND is covered also in the [OPA documentation](https://www.openpolicyagent.org/docs/latest/#expressions-logical-and).
{% endhint %}

{% code title="rules/CUSTOM-RULE-2/main.rego" %}
```
package rules

aws_redshift_cluster_tags_present(resource) {
    resource.tags.owner
    resource.tags.description
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    not aws_redshift_cluster_tags_present(resource)
    msg := {
        "publicId": "CUSTOM-RULE-2",
        "title": "Missing a description and an owner from the tag",
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

{% hint style="info" %}
Snyk recommends always validating that your rule is correct by [updating and running the unit tests](testing-a-rule.md).
{% endhint %}

The test for this rule will look the same as the one for `CUSTOM-RULE-1`, but the name of the test and the first two arguments passed to the `testing.evaluate_test_cases` function will differ:

{% code title="rules/CUSTOM-RULE-2/main_test.rego" %}
```
package rules

import data.lib
import data.lib.testing

test_CUSTOM_RULE_2 {
	# array containing test cases where the rule is allowed
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]
	# array containing cases where the rule is denied
	denied_test_cases := [{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied].tags"],
		"fixture": "denied.tf",
	}]
	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("CUSTOM-RULE-2", "./rules/CUSTOM-RULE-2/fixtures", test_cases)
}
```
{% endcode %}

## Example with logical OR

You can also rewrite the rule above by combining the `NOT` operator with the OR functionality.

Update the example in a new rule `CUSTOM-RULE-3`, to deny all cases that fail either of the two conditions, to deny all `aws_redshift_cluster` resources that are missing either:

1. an “owner” tag , OR
2. A “description” tag

For this, use two new fixture files, one for each case:

{% code title="rules/CUSTOM-RULE-3/fixtures/denied1.tf" %}
```
resource "aws_redshift_cluster" "denied1" {
  cluster_identifier = "tf-redshift-cluster"
  node_type          = "dc1.large"
  tags = {
    owner = "team-123@corp-domain.com"
  }
}
```
{% endcode %}

{% code title="rules/CUSTOM-RULE-3/fixtures/denied2.tg" %}
```
resource "aws_redshift_cluster" "denied2" {
  cluster_identifier = "tf-redshift-cluster"
  node_type          = "dc1.large"
  tags = {
    description = "description",
  }
}
```
{% endcode %}

To express logical OR in Rego, define multiple rules or functions with the same name. This is also described in the OPA documentation for[ Logical OR](https://www.openpolicyagent.org/docs/latest/#logical-or).

First, add a function that will implement the `NOT` for each tag. Then, call this function with the resource:

{% code title="rules/CUSTOM-RULE-3/main.rego" %}
```
package rules

aws_redshift_cluster_tags_missing(resource) {
    not resource.tags.owner
}

aws_redshift_cluster_tags_missing(resource) {
    not resource.tags.description
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    aws_redshift_cluster_tags_missing(resource)
    msg := {
        "publicId": "CUSTOM-RULE-3",
        "title": "Missing a description or an owner from the tag",
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

This will successfully return all the rules that deny.

{% hint style="info" %}
Snyk recommends always validating that your rule is correct by [updating and running the unit tests](testing-a-rule.md).
{% endhint %}

The test for this rule will now contain multiple test cases, to show that the logical OR works as expected:

{% code title="rules/CUSTOM-RULE-3/main_test.rego" %}
```
package rules

import data.lib
import data.lib.testing

test_CUSTOM_RULE_3 {
	# array containing test cases where the rule is allowed
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]
	# array containing cases where the rule is denied
	denied_test_cases := [{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied1].tags"],
		"fixture": "denied1.tf",
	},{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied2].tags"],
		"fixture": "denied2.tf",
	}]
	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("CUSTOM-RULE-3", "./rules/CUSTOM-RULE-3/fixtures", test_cases)
}
```
{% endcode %}

## Example with strings

Extend this further and add a third condition. Deny all resources that are missing any of the following:

1. An “owner” tag , OR
2. A “description” tag, OR
3. The email of the owner does not belong to the “@corp-domain.com” domain

{% code title="rules/CUSTOM-RULE-4/main.rego" %}
```
package rules

aws_redshift_cluster_tags_missing(resource) {
    not resource.tags.owner
}

aws_redshift_cluster_tags_missing(resource) {
    not resource.tags.description
}

aws_redshift_cluster_tags_missing(resource) {
    not endswith(resource.tags.owner, "@corp-domain.com")
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    aws_redshift_cluster_tags_missing(resource)
    msg := {
        "publicId": "CUSTOM-RULE-4",
        "title": "Missing a description and an owner from tag, or owner tag does not comply with email requirements",
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

{% hint style="info" %}
Snyk recommends always validating that your rule is correct by [updating and running the unit tests](testing-a-rule.md).
{% endhint %}

The test for this rule will look very similar to the ones from previous example and will also require its own fixture file.

## Example with XOR

Now assume you want to add more complexity and check the following:

* If the tag type is a “user”, then the tag “email” exists as well.
* If not, assuming the other type is a “service”, we want it has a serviceDescription.
* These two will be mutually exclusive; if the first condition applies, the second one does not, and vice versa.

| Type    | Email | ServiceDescription |
| ------- | ----- | ------------------ |
| User    | YES   | NO                 |
| Service | NO    | YES                |

To do this, refactor your code to use a checkTags helper function. This can check whether there are any tags, but also check for the two conditions above with an OR.

{% code title="rules/CUSTOM-RULE-5/main.rego" %}
```
package rules

checkTags(resource){
    resource.tags.type == "user"
    not resource.tags.email
}

checkTags(resource){
    resource.tags.type == "service"
    not resource.tags.serviceDescription
}

checkTags(resource){
    count(resource.tags) == 0
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    checkTags(resource)   

    msg := {
        "publicId": "CUSTOM-RULE-5",
        "title": "Complex rule",
        "severity": "medium",
        "msg": sprintf("input.resource.aws_redshift_cluster[%v].tags", [name]),
        "issue": "",
        "impact": "",
        "remediation": "",
        "references": [],
    }
}
```
{% endcode %}

To convert this to an XOR you can use an `else` rule:

{% code title="rules/CUSTOM-RULE-5/main.rego" %}
```
package rules

checkUserTag(resource){
    not resource.tags.email
}

checkUserTag(resource){
    resource.tags.serviceDescription
}

checkServiceTag(resource){
    not resource.tags.serviceDescription
}

checkServiceTag(resource){
    resource.tags.email
}

checkTags(resource){
	count(resource.tags) == 0
}

checkTags(resource) {
    resource.tags.type == "user"
    checkUserTag(resource)
} else {
    resource.tags.type == "service"
    checkServiceTag(resource)
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
	checkTags(resource)
    msg := {
        "publicId": "CUSTOM-RULE-5",
        "title": "Missing the right tags from for a resource of type user or service",
        "severity": "medium",
        "msg": sprintf("input.resource.aws_redshift_cluster[%v].tags", [name]),
        "issue": "",
        "impact": "",
        "remediation": "",
        "references": [],
    }
}
```
{% endcode %}

If you want to try it out yourselves, use the same example in this [OPA Playground](https://play.openpolicyagent.org/p/1xcdj9kJRw).

{% hint style="info" %}
Snyk recommends always validating that your rule is correct by[ updating and running the unit tests](testing-a-rule.md).
{% endhint %}

The test for this rule will look very similar to the ones from the previous example and will also require its own fixture file.

## Examples with grouped resources

You can also iterate over many resources by adding them to an array of resources.

```
"resources": [
            "aws_iam_policy",
            "aws_iam_group_policy",
            "aws_iam_user_policy",
            "aws_iam_role_policy",
            "data.aws_iam_policy_document",
]
```

One way to leverage this is to implement denylist rules.

For example, you may want to ensure that if someone defines a Kubernetes ConfigMap, then they cannot use it to store sensitive information such as passwords, secret keys, and access tokens.

You can do that and expand what is defined as "sensitive information" over time by defining a group of sensitive tokens inside a denylist:

```
package rules

sensitive_denylist := [
	"pass",
	"secret",
	"key",
	"token",
]

check_sensitive(keys, denylist) {
	_ = keys[key]
	contains(key, denylist[_])
}

deny[msg] {
	input.kind == "ConfigMap"
	input.data = keys
	check_sensitive(keys, sensitive_denylist)
	msg := {
		"publicId": "CUSTOM-RULE-7",
		"title": "ConfigMap exposes sensitive data",
		"severity": "high",
		"msg": "input.data",
		"issue": "",
		"impact": "",
		"remediation": "",
		"references": [],
	}
}
```

Any key containing the substrings "pass", "secret", "key", and "token" can be considered sensitive and so should not be included in the ConfigMap.
