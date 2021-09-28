# Parsing an input file

It can be difficult to understand the internal representation of your input files as you write your Rego code. As we will see when we learn [how to write a rule](writing-a-rule.md),  the input value is a JSON-like object but the input files could also be YAML or Terraform. To help understand how these are translated into JSON we have provided a `parse` command.

You will need an IaC file to use as an input file. This input file can also be used when [testing the rules](testing-a-rule.md).

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

In Rego, accessing the `node_type` would look like: 

```text
input.resource.aws_redshift_cluster.example.node_type
```

