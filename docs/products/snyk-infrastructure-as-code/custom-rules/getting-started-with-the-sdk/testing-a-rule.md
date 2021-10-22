# Testing a rule

Open the `main_test.rego` file created in the previous tutorial and configure the `fixture` field with the name of the file inside your `rules/my_rule/fixtures/` folder.

Create files to store your resources under `rules/my_rule/fixtures`. These files can have any name, so take for example `denied.tf` and `allowed.tf`:

{% code title="rules/my_rule/fixtures/denied.tf" %}
```
resource "aws_redshift_cluster" "denied" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
}
```
{% endcode %}

{% code title="rules/my_rule/fixtures/allowed.tf" %}
```
resource "aws_redshift_cluster" "allowed" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  tags {
    owner = "snyk"
  }
}
```
{% endcode %}

In the `want_msgs` field of the test case, you should add the msg fields of the resources that you expect that your deny rule will evaluate/return, e.g. `["resource.aws_redshift_cluster[denied]"]`.

{% code title="rules/my_rule/main_test.rego" %}
```
package rules

import data.lib
import data.lib.testing

test_my_rule {
		# array containing test cases where the rule is allowed
		allowed_test_cases := [{
				"want_msgs": [],
				"fixture": "allowed.tf",
		}]

		# array containing cases where the rule is denied
		denied_test_cases := [{
				"want_msgs": ["resource.aws_redshift_cluster[denied]"], # verifies that the correct msg is returned by the denied rule
				"fixture": "denied.tf",
		}]

		test_cases := array.concat(allowed_test_cases, denied_test_cases)
    testing.evaluate_test_cases("my_rule", "./rules/my_rule/fixtures"", test_cases)
}
```
{% endcode %}

To run all tests, run the following command:

```
 snyk-iac-rules test
```

If your tests pass successfully, you will see an output similar to the following:

```
PASS: 3/3
```

However, if any of them fail, you will see an output similar to the following:

```
data.rules.test_my_rule: FAIL (1.12234ms)
FAIL: 2/3
```

If you have more than one rule in your `rule/` folder you can target a specific test by running the following command:

```
snyk-iac-rules test --run test_CUSTOM_1
```

This will output:

```
Executing Rego test cases...
data.rules.my_rule: FAIL (1.040468ms)
--------------------------------------------------------------------------------
FAIL: 1/1
```

If you need more details about it, add the` --explain notes` option:

```
 snyk-iac-rules test --run test_my_rule --explain notes
```

This will output more details to debug the failed test.

{% hint style="info" %}
If you have more than your generated rules in the current folder consider using the  `--ignore` flag to exclude the folders and files irrelevant to testing (make sure to not exclude `lib/` and `rules` if you used the `template` command). This can speed up the tests and also avoids running into problems where Rego is trying to evaluate non-Rego files.
{% endhint %}
