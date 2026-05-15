# Parsing an input file

It can be difficult to understand the internal representation of your input files as you write your Rego code. As you will see when you learn [how to write a rule](writing-a-rule.md), the input value is a JSON-like object, but the input files could also be YAML, Terraform, or [Terraform Plan JSON Output](https://www.terraform.io/docs/internals/json-format.html). To help you understand how these are translated into JSON, Snyk provides a `parse` command.

You will need an IaC file to use as an input file. This input file can also be used when [testing the rules](testing-a-rule.md), which parses your files into JSON by default.

## Parsing Terraform files

Take, for example, the following Terraform file:

{% code title="example.tf" %}
```
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

```
snyk-iac-rules parse example.tf --format hcl2
```

This prints out the JSON, which you can use as guidance for writing your rules:

```
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

In Rego, accessing the `node_type` field would look like this:

```
input.resource.aws_redshift_cluster.example.node_type
```

## Parsing YAML files

Another example is the following YAML file, defining a Kubernetes resource:

{% code title="example.yaml" %}
```
apiVersion: v1
kind: Pod
metadata:
  name: example
spec:
  containers:
    - name: example
      image: example:latest
      securityContext:
        privileged: true 
```
{% endcode %}

To get the equivalent JSON format, run the parse command:

```
snyk-iac-rules parse example.yaml --format=yaml
```

This prints out the JSON, which you can use as guidance for writing your rules:

```
{
	"apiVersion": "v1",
	"kind": "Pod",
	"metadata": {
		"name": "example"
	},
	"spec": {
		"containers": [
			{
				"image": "example:latest",
				"name": "example",
				"securityContext": {
					"privileged": true
				}
			}
		]
	}
}
```

In Rego, accessing the `privileged` field would look like this:

```
input.spec.containers[0].securityContext.privileged
```

## Parsing Terraform Plan JSON output files

Another example is the following Terraform Plan JSON Output file, returned by the `terraform show -json ./plan/example.json.tfplan` command:

{% code title="example.json.tfplan" %}
```
{
  "format_version": "0.2",
  "terraform_version": "1.0.11",
  "planned_values": {
    "root_module": {
      "resources": [
        {
          "address": "aws_vpc.example",
          "mode": "managed",
          "type": "aws_vpc",
          "name": "example",
          "provider_name": "registry.terraform.io/hashicorp/aws",
          "schema_version": 1,
          "values": {
            "assign_generated_ipv6_cidr_block": false,
            "cidr_block": "10.0.0.0/16",
            "enable_dns_support": true,
            "instance_tenancy": "default",
            "tags": null
          },
          "sensitive_values": {
            "tags_all": {}
          }
        }
      ]
    }
  },
  "resource_changes": [
    {
      "address": "aws_vpc.example",
      "mode": "managed",
      "type": "aws_vpc",
      "name": "example",
      "provider_name": "registry.terraform.io/hashicorp/aws",
      "change": {
        "actions": [
          "create"
        ],
        "before": null,
        "after": {
          "assign_generated_ipv6_cidr_block": false,
          "cidr_block": "10.0.0.0/16",
          "enable_dns_support": true,
          "instance_tenancy": "default",
          "tags": null
        },
        "after_unknown": {
          "arn": true,
          "default_network_acl_id": true,
          "default_route_table_id": true,
          "default_security_group_id": true,
          "dhcp_options_id": true,
          "enable_classiclink": true,
          "enable_classiclink_dns_support": true,
          "enable_dns_hostnames": true,
          "id": true,
          "ipv6_association_id": true,
          "ipv6_cidr_block": true,
          "main_route_table_id": true,
          "owner_id": true,
          "tags_all": true
        },
        "before_sensitive": false,
        "after_sensitive": {
          "tags_all": {}
        }
      }
    }
  ],
  "configuration": {
    "provider_config": {
      "aws": {
        "name": "aws",
        "expressions": {
          "region": {
            "constant_value": "us-east-1"
          }
        }
      }
    },
    "root_module": {
      "resources": [
        {
          "address": "aws_vpc.example",
          "mode": "managed",
          "type": "aws_vpc",
          "name": "example",
          "provider_config_key": "aws",
          "expressions": {
            "cidr_block": {
              "constant_value": "10.0.0.0/16"
            }
          },
          "schema_version": 1
        }
      ]
    }
  }
}
```
{% endcode %}

To get the equivalent JSON format, run the parse command:

```
snyk-iac-rules parse example.json.tfplan --format=tf-plan
```

This prints out the JSON, which you can use as guidance for writing your rules:

```
{
	"data": {},
	"resource": {
		"aws_vpc": {
			"example": {
				"assign_generated_ipv6_cidr_block": false,
				"cidr_block": "10.0.0.0/16",
				"enable_dns_support": true,
				"instance_tenancy": "default",
				"tags": null
			}
		}
	}
}
```

In Rego, accessing the `tags` field would look like this:

```
input.resource.aws_vpc.example.tags
```
