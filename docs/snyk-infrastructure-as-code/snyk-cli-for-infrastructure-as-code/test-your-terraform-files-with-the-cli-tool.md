# Test your Terraform files with the CLI tool

With Snyk Infrastructure as Code, you can scan both your static configuration files and Terraform Plan output using the CLI. 

|  | **Terraform configuration files** | **Terraform Plan file** |
| :--- | :--- | :--- |
| **Identify configuration issues** | Yes | Yes |
| **Process Variables** | No | Yes |
| **Scan Terraform Modules** | No | Yes - public & private |

### Terraform Configuration Files

You can scan the configuration files, e.g. \`main.tf\` using the CLI. 

Any declared variables or external modules that are referenced will not be considered. 

### To scan configuration files:

You can specify either a file name or a whole directory

```text
snyk iac test main.tf
snyk iac test .
```

### Terraform Plan

Terraform Plan is the step run between writing your configuration files and deploying those changes.

`$ terraform plan` identifies the changes that need to be made to your target environment in order to match your desired state. 

As part of this planning stage, all variables and Terraform Modules that are used in your targeted terraform deployment are taken into consideration. 

If you have written a custom terraform module and are referencing it in your deployment, then it will be included in the terraform plan output and scanned accordingly. 

This means the Terraform plan output provides a complete artefact to be scanned from a security perspective. 

You can now scan this artefact using the Snyk IaC CLI as of version **1.594.0**

This file is not sent to Snyk to be processed, it is scanned locally within the CLI and does not leave your machine. 

### To scan Terraform Plan output:

Provide the path to your Terraform Plan output which must be stored as a valid JSON file. 

```text
snyk iac test tf-plan.json
```

By default, we will scan the changes that will be made to your infrastructure, not the full infrastructure. 

You can change this behaviour by providing the `--scan=` flag

* `--scan=resource-changes` is the default behaviour. This will scan only the changes that are going to be made to your infrastructure if you ran `$ terraform apply`
* `--scan=planned-values` will scan the full planned state, providing results of the existing infrastructure plus changes that will be made. 

If you do not already have your terraform plan output saved as JSON file, you may need to follow these steps:

```text
terraform plan -out=tfplan.binary
terraform show -json tfplan.binary > tf-plan.json
```

You can name the tf-plan.json file according to your needs. 

These files are considered sensitive and is not recommended to commit them to source control. 

### Troubleshooting

#### There are differences between scanning the static files & plan output

This could be due to the following

* **Variables** - Terraform Plan output considers the values stored in variables
* **Terraform Modules** - Terraform Plan output will include any configuration issues found from Terraform Modules that you may be using
* **Delta** - By default, scanning the Terraform Plan output will only scan for configuration issues on the changes that will be made, not the whole deployment. Whereas the static scan looks at all of the files. Try re-running the scan with `--scan=planned-values` appended

```text
Note: the flag --experimental is not required anymore when testing your Terraform projects.
```

If you have found a discrepancy that you cannot explain with the above, please raise a support ticket. 

