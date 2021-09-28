# Parsing an input file

It can be difficult to understand the internal representation of your input files as you write your Rego code. As we will see when we learn [how to write a rule](writing-a-rule.md), the `input` object is treated like a JSON object, but the input files may be YAML or Terraform. Rego can scan these files, but to help with writing the Rego from scratch, we have provided a `parse` command.

You will need an IaC file to use as an input file. This input file can also be used when [testing the rules](testing-a-rule.md), but for now it is helpful to write your rules.

For example, the following Terraform file:

{% code title="example.tf" %}
```text
resource "aws_redshift_cluster" "example" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
}
```
{% endcode %}

To get the equivalent JSON format, run the parse command:

```text
snyk-iac-rules parse example.tf
```

This prints out the JSON, which you can use as guidance for writing your rules:

```text
{
	"resource": {
		"aws_redshift_cluster": {
			"example": {
				"cluster_identifier": "tf-redshift-cluster",
				"cluster_type": "single-node",
				"database_name": "mydb",
				"master_password": "Mustbe8characters",
				"master_username": "foo",
				"node_type": "dc1.large"
			}
		}
	}
}
```

