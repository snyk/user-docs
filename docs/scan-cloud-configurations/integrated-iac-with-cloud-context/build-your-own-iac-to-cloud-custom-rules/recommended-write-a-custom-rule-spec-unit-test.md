# (Recommended) Write a Custom Rule Spec (Unit Test)

Custom Rule Specs are unit tests for your Custom Rules. As you iterate on your Custom Rules, it is recommended that you maintain unit tests. These unit tests typically include sample terraform that represents valid and invalid test cases. The following is an example Custom Rule spec for a custom rule that enforces tag requirements: &#x20;

{% code title="/spec/rules/REQUIRED_S3_BUCKET_TAGS" %}
```
provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "valid" {
  bucket = "valid-bucket"

  tags = {
    owner           = "devteam1"
    classification  = "public"
  }
}

resource "aws_s3_bucket" "invalid" {
  bucket = "invalid-bucket"

  tags = {
    owner           = "devteam5"
    classification  = "available"
  }
}
```
{% endcode %}

## Creating a Custom Rule Spec template​​

Use the following CLI command:

```
snyk iac rules init
```

This command introduces a series of interactive prompts and allows you to set up your Project, rule, rule spec (for testing), and resource relations. By selecting `rule spec`,  you will get a series of interactive prompts to create the template for your rule spec. This results in a template based on Terraform, CloudFormation, Azure Resource Manager, or Kubernetes.

**Example output from interactive prompts:**

```
What do you want to initialize? rule spec
Choose a rule ID: RULE_1
Spec name: RULE_SPEC_1
Input type: tf
```

## Writing a Custom Rule spec

After the Custom Rule spec template is created, you must populate it with valid and invalid configurations, with the Input Type selected, such as Terraform.
