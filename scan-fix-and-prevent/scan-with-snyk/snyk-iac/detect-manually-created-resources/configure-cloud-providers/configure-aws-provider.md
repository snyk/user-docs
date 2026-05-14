# Configure AWS provider

## Authentication for AWS provider

To use `iac describe`, set up credentials to make authenticated requests to AWS. As you do for the AWS CLI, use [credentials and configuration](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) settings declared as user environment variables or in local AWS configuration files.

The `iac describe` command supports a [named profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html). By default, the CLI uses the settings found in the profile named `default`. You can override an individual setting by declaring the supported environment variables such as `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_PROFILE` and so on.

If you are using an [IAM role](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html) as an authorization tool, which is considered a good practice, you can still use `iac describe` by defining a profile for the role in your `~/.aws/config` file.

```
[profile snykrole]
role_arn = arn:aws:iam::123456789012:role/<NAMEOFTHEROLE>
source_profile = user # profile to assume the role
region = eu-west-3
```

You can now use `iac describe` by overriding the profile setting.

```
$ AWS_PROFILE=snykrole snyk iac describe
```

## Custom credentials to read a state on an S3 backend[​](https://docs.driftctl.com/0.22.0/providers/aws/authentication#custom-credentials-to-read-a-state-on-an-s3-backend) <a href="#custom-credentials-to-read-a-state-on-an-s3-backend" id="custom-credentials-to-read-a-state-on-an-s3-backend"></a>

If you want to use a different set of AWS credentials to read your state on S3, you can override each specific AWS environment variable with the `DCTL_S3_` prefix. The purpose is to have the choice to read a state in a different region from your infrastructure. Remember to use your usual AWS credentials to read the resources of your actual infrastructure.

```
# Export a dedicated AWS named profile (or any other AWS environment variables) to read your state in your S3 backend
$ export DCTL_S3_PROFILE="s3reader"
# Export the usual AWS named profile
$ export AWS_PROFILE="snykrole"
$ snyk iac describe --from="tfstate+s3://mybucket/terraform.tfstate"

# You can also use a specific region to authenticate to the S3 bucket
$ DCTL_S3_REGION=us-east-1 snyk iac describe --from="tfstate+s3://mybucket/terraform.tfstate"
```

## Terraform custom role​ <a href="#terraform-custom-role" id="terraform-custom-role"></a>

The following code represents the custom role you can assume to run `iac describe` written in HCL.

```
data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    effect = "Allow"
    actions   = ["sts:AssumeRole"]
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }
  }
}

data "aws_iam_policy_document" "policy" {
  statement {
    effect = "Allow"
    actions   = [
      "apigateway:GET",
      "cloudformation:DescribeStacks",
      "cloudformation:GetTemplate",
      "cloudfront:GetDistribution",
      "cloudfront:ListDistributions",
      "cloudfront:ListTagsForResource",
      "ec2:DescribeAddresses",
      "ec2:DescribeImages",
      "ec2:DescribeInstanceAttribute",
      "ec2:DescribeInstances",
      "ec2:DescribeInstanceCreditSpecifications",
      "ec2:DescribeInternetGateways",
      "ec2:DescribeKeyPairs",
      "ec2:DescribeNetworkAcls",
      "ec2:DescribeRouteTables",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeSnapshots",
      "ec2:DescribeTags",
      "ec2:DescribeVolumes",
      "ec2:DescribeVpcs",
      "ec2:DescribeVpcAttribute",
      "ec2:DescribeVpcClassicLink",
      "ec2:DescribeVpcClassicLinkDnsSupport",
      "ec2:DescribeSubnets",
      "ec2:DescribeNatGateways",
      "ec2:DescribeLaunchTemplates",
      "ecr:DescribeRepositories",
      "ecr:ListTagsForResource",
      "iam:GetPolicy",
      "iam:GetPolicyVersion",
      "iam:GetRole",
      "iam:GetRolePolicy",
      "iam:GetUser",
      "iam:GetUserPolicy",
      "iam:ListAccessKeys",
      "iam:ListAttachedRolePolicies",
      "iam:ListAttachedUserPolicies",
      "iam:ListPolicies",
      "iam:ListRolePolicies",
      "iam:ListRoles",
      "iam:ListUserPolicies",
      "iam:ListUsers",
      "kms:DescribeKey",
      "kms:GetKeyPolicy",
      "kms:GetKeyRotationStatus",
      "kms:ListAliases",
      "kms:ListKeys",
      "kms:ListResourceTags",
      "lambda:GetEventSourceMapping",
      "lambda:GetFunction",
      "lambda:GetFunctionCodeSigningConfig",
      "lambda:ListEventSourceMappings",
      "lambda:ListFunctions",
      "lambda:ListVersionsByFunction",
      "rds:DescribeDBInstances",
      "rds:DescribeDBSubnetGroups",
      "rds:ListTagsForResource",
      "route53:GetHostedZone",
      "route53:ListHostedZones",
      "route53:ListResourceRecordSets",
      "route53:ListTagsForResource",
      "route53:ListHealthChecks",
      "route53:GetHealthCheck",
      "s3:GetAccelerateConfiguration",
      "s3:GetAnalyticsConfiguration",
      "s3:GetBucketAcl",
      "s3:GetBucketCORS",
      "s3:GetBucketLocation",
      "s3:GetBucketLogging",
      "s3:GetBucketNotification",
      "s3:GetBucketObjectLockConfiguration",
      "s3:GetBucketPolicy",
      "s3:GetBucketRequestPayment",
      "s3:GetBucketTagging",
      "s3:GetBucketVersioning",
      "s3:GetBucketWebsite",
      "s3:GetEncryptionConfiguration",
      "s3:GetInventoryConfiguration",
      "s3:GetLifecycleConfiguration",
      "s3:GetMetricsConfiguration",
      "s3:GetReplicationConfiguration",
      "s3:ListAllMyBuckets",
      "s3:ListBucket",
      "sqs:GetQueueAttributes",
      "sqs:ListQueueTags",
      "sqs:ListQueues",
      "sns:ListTopics",
      "sns:GetTopicAttributes",
      "sns:ListTagsForResource",
      "sns:ListSubscriptions",
      "sns:ListSubscriptionsByTopic",
      "sns:GetSubscriptionAttributes",
      "dynamodb:ListTables",
      "dynamodb:DescribeTable",
      "dynamodb:DescribeGlobalTable",
      "dynamodb:ListTagsOfResource",
      "dynamodb:DescribeTimeToLive",
      "dynamodb:DescribeTableReplicaAutoScaling",
      "dynamodb:DescribeContinuousBackups",
      "rds:DescribeDBClusters",
      "rds:DescribeGlobalClusters",
      "application-autoscaling:DescribeScheduledActions",
      "autoscaling:DescribeLaunchConfigurations"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_role" "snyk_assume_role" {
  name = "snyk_assume_role"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

resource "aws_iam_role_policy" "snyk_policy" {
  name = "snyk_policy"
  role = aws_iam_role.snyk_assume_role.id
  policy = data.aws_iam_policy_document.policy.json
}
```

## ​CloudFormation template <a href="#cloudformation-template" id="cloudformation-template"></a>

Deploy this CloudFormation template to create the limited permission role that you can use according to the authentication guide in the preceding sections of this page.

[![Launch Stack](https://cdn.rawgit.com/buildkite/cloudformation-launch-stack-button-svg/master/launch-stack.svg)](https://console.aws.amazon.com/cloudformation/home?#/stacks/quickcreate?stackName=driftctl-stack\&templateURL=https://driftctl-cfn-templates.s3.eu-west-3.amazonaws.com/driftctl-role.yml)

When the stack is deployed, attach the following policy to your IAM user. This allows the user to assume only the role specified. For more information about granting a user access to assume a role, see the [AWS Identity and Access Management User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_permissions-to-switch.html).

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Resource": "arn:aws:iam::<IDOFYOURACCOUNT>:role/snyk_assume_role"
    }
  ]
}
```

There is no automatic way to update the CloudFormation template from the Snyk side because you launched this template from your AWS account. Therefore you must update the template yourself to use the most recent Snyk role.

### Update the CloudFormation template using the AWS console

* In the [AWS CloudFormation console](https://console.aws.amazon.com/cloudformation), from the list of stacks, select the **snyk stack**.
* In the stack details pane, choose **Update**.
* Select **Replace current template** and specify the Snyk **Amazon S3 URL** `https://driftctl-cfn-templates.s3.eu-west-3.amazonaws.com/driftctl-role.yml`; click **Next.**
* On the **Specify stack details** and the **Configure stack options** pages, click **Next**.
* In the **Change set preview** section, check that AWS CloudFormation will make the changes.
* Because the Snyk template contains one IAM resource, select **I acknowledge that this template may create IAM resources**.
* To finish, click **Update stack**.

### Update the CloudFormation template using the AWS CLI

Use the following command:

```
$ aws cloudformation update-stack --stack-name SNYK_STACK_NAME --template-url https://driftctl-cfn-templates.s3.eu-west-3.amazonaws.com/snyk-role.yml --capabilities CAPABILITY_NAMED_IAM
```

## Least privilege policy​ <a href="#least-privileged-policy" id="least-privileged-policy"></a>

The `iac describe` command needs access to your cloud provider account so that it can list resources on your behalf.

As the AWS documentation recommends, the policy that follows grants only the required permissions.

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Resource": "*",
            "Action": [
                "apigateway:GET",
                "cloudformation:DescribeStacks",
                "cloudformation:GetTemplate",
                "cloudfront:GetDistribution",
                "cloudfront:ListDistributions",
                "cloudfront:ListTagsForResource",
                "ec2:DescribeAddresses",
                "ec2:DescribeImages",
                "ec2:DescribeInstanceAttribute",
                "ec2:DescribeInstances",
                "ec2:DescribeInstanceCreditSpecifications",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeKeyPairs",
                "ec2:DescribeNetworkAcls",
                "ec2:DescribeRouteTables",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSnapshots",
                "ec2:DescribeTags",
                "ec2:DescribeVolumes",
                "ec2:DescribeVpcs",
                "ec2:DescribeVpcAttribute",
                "ec2:DescribeVpcClassicLink",
                "ec2:DescribeVpcClassicLinkDnsSupport",
                "ec2:DescribeSubnets",
                "ec2:DescribeNatGateways",
                "ec2:DescribeLaunchTemplates",
                "ecr:DescribeRepositories",
                "ecr:ListTagsForResource",
                "iam:GetPolicy",
                "iam:GetPolicyVersion",
                "iam:GetRole",
                "iam:GetRolePolicy",
                "iam:GetUser",
                "iam:GetUserPolicy",
                "iam:ListAccessKeys",
                "iam:ListAttachedRolePolicies",
                "iam:ListAttachedUserPolicies",
                "iam:ListPolicies",
                "iam:ListRolePolicies",
                "iam:ListRoles",
                "iam:ListUserPolicies",
                "iam:ListUsers",
                "kms:DescribeKey",
                "kms:GetKeyPolicy",
                "kms:GetKeyRotationStatus",
                "kms:ListAliases",
                "kms:ListKeys",
                "kms:ListResourceTags",
                "lambda:GetEventSourceMapping",
                "lambda:GetFunction",
                "lambda:GetFunctionCodeSigningConfig",
                "lambda:ListEventSourceMappings",
                "lambda:ListFunctions",
                "lambda:ListVersionsByFunction",
                "rds:DescribeDBInstances",
                "rds:DescribeDBSubnetGroups",
                "rds:ListTagsForResource",
                "route53:GetHostedZone",
                "route53:ListHostedZones",
                "route53:ListResourceRecordSets",
                "route53:ListTagsForResource",
                "route53:ListHealthChecks",
                "route53:GetHealthCheck",
                "s3:GetAccelerateConfiguration",
                "s3:GetAnalyticsConfiguration",
                "s3:GetBucketAcl",
                "s3:GetBucketCORS",
                "s3:GetBucketLocation",
                "s3:GetBucketLogging",
                "s3:GetBucketNotification",
                "s3:GetBucketObjectLockConfiguration",
                "s3:GetBucketPolicy",
                "s3:GetBucketRequestPayment",
                "s3:GetBucketTagging",
                "s3:GetBucketVersioning",
                "s3:GetBucketWebsite",
                "s3:GetEncryptionConfiguration",
                "s3:GetInventoryConfiguration",
                "s3:GetLifecycleConfiguration",
                "s3:GetMetricsConfiguration",
                "s3:GetReplicationConfiguration",
                "s3:ListAllMyBuckets",
                "s3:ListBucket",
                "sqs:GetQueueAttributes",
                "sqs:ListQueueTags",
                "sqs:ListQueues",
                "sns:ListTopics",
                "sns:GetTopicAttributes",
                "sns:ListTagsForResource",
                "sns:ListSubscriptions",
                "sns:ListSubscriptionsByTopic",
                "sns:GetSubscriptionAttributes",
                "dynamodb:ListTables",
                "dynamodb:DescribeTable",
                "dynamodb:DescribeGlobalTable",
                "dynamodb:ListTagsOfResource",
                "dynamodb:DescribeTimeToLive",
                "dynamodb:DescribeTableReplicaAutoScaling",
                "dynamodb:DescribeContinuousBackups",
                "rds:DescribeDBClusters",
                "rds:DescribeGlobalClusters",
                "application-autoscaling:DescribeScalableTargets",
                "application-autoscaling:DescribeScalingPolicies",
                "application-autoscaling:DescribeScheduledActions",
                "autoscaling:DescribeLaunchConfigurations"
            ]
        }
    ]
}
```
