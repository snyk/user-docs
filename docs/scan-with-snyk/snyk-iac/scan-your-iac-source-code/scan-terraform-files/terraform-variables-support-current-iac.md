# Terraform variables support

Support for Terraform (TF) variables is currently available only in the CLI. Snyk supports:

* [Input Variables](https://www.terraform.io/language/values/variables)
* [Local Values](https://www.terraform.io/language/values/locals)

At this time Snyk does not support [Output Values](https://www.terraform.io/language/values/outputs).

The CLI scans all directories and handles each directory that includes supported TF files as its own module. Each module that includes variables is dereferenced appropriately.

Supported TF file formats are `.tf`, `.tfvars`, `.auto.tfvars`. Snyk does not support variables that were set and defined using environment variables or the `--var` CLI option.

The scan handles [variable definition precedence](https://www.terraform.io/language/values/variables#variable-definition-precedence) in the same way that TF handles the precedence.

You can load an external variable definitions file by using the `--var-file` option, for example:

`snyk iac test myproject/staging/networking --var-file=myproject/vars.tf`

This loads the `vars.tf` definitions file from the `myproject` directory, dereferences any variables if they exist, and applies them to the context of the scanned path, `myproject/staging/networking` in this example.

For more information, see the `IAC test` [help](../../../../developer-tools/snyk-cli/commands/iac-test.md).

## Supported Terraform expressions

The following expressions are supported:

* [Arithmetic and Logical Operators](https://www.terraform.io/language/expressions/operators)
* [Strings and Templates](https://www.terraform.io/language/expressions/strings#strings-and-templates)
* [Conditional Expressions](https://www.terraform.io/language/expressions/conditionals)
* [For Expressions](https://www.terraform.io/language/expressions/for)
* [Splat Expressions](https://www.terraform.io/language/expressions/splat)

## Supported Terraform functions

The following functions are supported:

* Numeric functions - all functions
* String functions - all functions except `lower`, `regex`, `regexall`, `replace`, `substr`, `title`, `upper`
* Collection functions - `chunklist`, `concat`, `distinct`, `flatten`, `length`, `merge`, `reverse`, `sort`
* Encoding functions - `csvdecode`, `jsondecode`, `jsonencode`
* Date and Time functions - `formatdate`, `timeadd`

## Examples of Terraform variables

### **Variable handling in the correct precedence**

In the example that follows, we see that we configured a new resource, and we are using a variable named `remote_user_addr` to set its `cidr_blocks` value.

The variable is defined inside the `variables.tf` file with a default value, but the value is being overridden inside the `terraform.tfvars` file.

At the end, the value is set to `0.0.0.0/0`, and this causes the CLI to raise an issue.

```hcl
vpc.tf

resource "aws_security_group_rule" "ssh" {
  type              = "ingress"
  from_port         = 22
  to_port           = 22
  protocol          = "tcp"
  cidr_blocks       = [var.remote_user_addr]
  security_group_id = aws_security_group.allow.id
}
```

```hcl
variables.tf

variable "remote_user_addr" {
  type = string
  default = "11.0.0.0/24"
}
```

```hcl
terraform.tfvars

remote_user_addr = "0.0.0.0/0"
```

## **Conditional expression using variables**

In the following example, we are using local and input variables together with a conditional expression.

We are checking to see if `local.test` equals 0, and we are setting the `cidr_blocks` accordingly.

In our case `local.test` equals 0, and the value is set to the value of `var.remote_user_addr` which causes the CLI to raise an issue.

```hcl
vpc.tf

resource "aws_security_group_rule" "ssh" {
  type              = "ingress"
  from_port         = 22
  to_port           = 22
  protocol          = "tcp"
  cidr_blocks       = local.test == 0 ? [var.remote_user_addr] : ["11.0.0.0/24"]
  security_group_id = aws_security_group.allow.id
}

locals {
  test = 0
}
```

```hcl
variables.tf

variable "remote_user_addr" {
  type = string
  default = "0.0.0.0/0"
}
```
