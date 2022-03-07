# Terraform variables support

## Summary

Support for Terraform variables is currently only available in the CLI and in partial way.

Snyk currently supports:

1. [Input Variables](https://www.terraform.io/language/values/variables)
2. [Local Values](https://www.terraform.io/language/values/locals)

At this time Snyk does not support:

1. [Output Values](https://www.terraform.io/language/values/outputs)
2. Any variables that are being used in external modules

Supported TF file formats are `.tf`, `.tfvars`,`.auto.tfvars`.

## Supported expressions

The following expressions are currently supported:

* [Arithmetic and Logical Operators](https://www.terraform.io/language/expressions/operators)
* [Strings and Templates](https://www.terraform.io/language/expressions/strings#strings-and-templates)
* [Conditional Expressions](https://www.terraform.io/language/expressions/conditionals)
* [For Expressions](https://www.terraform.io/language/expressions/for)
* [Splat Expressions](https://www.terraform.io/language/expressions/splat)

## Supported functions

The following functions are currently supported:

* Numeric Functions - all functions
* String Functions - all functions except `lower`, `regex`, `regexall`, `replace`, `substr`, `title`,`upper`
* Collection Functions - `chunklist`, `concat`, `distinct`, `flatten`, `length`, `merge`, `reverse`, `sort`
* Encoding Functions - `csvdecode`, `jsondecode`, `jsonencode`
* Date and Time Functions - `formatdate`, `timeadd`
