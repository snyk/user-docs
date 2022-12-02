# AWS Quick start guide

## [Open the AWS Quick Start Guide](https://aws.amazon.com/quickstart/architecture/eks-snyk/)

You have the option of deploying the Snyk controller for Amazon EKS as an official [AWS Quick Start](https://aws.amazon.com/quickstart/architecture/eks-snyk/). This option eliminates the need for manual configuration. Deploying this Quick Start with default parameters into an existing Amazon EKS cluster builds the following environment.

![Snyk AWS Quick Start environment](<../../../.gitbook/assets/architecture (1).png>)

There are three deployment options to match most common use cases. These are as follows:

**1.** If you already have an Amazon EKS cluster running in your AWS account

[![cloudformation-launch-stack.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](../../../.gitbook/assets/cloudformation-launch-stack.png)](https://us-east-2.console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/template?stackName=Snyk-EKS\&templateURL=https://aws-quickstart.s3.us-east-1.amazonaws.com/quickstart-amazon-eks/submodules/quickstart-eks-snyk/templates/eks-snyk.template.yaml)

**2.** If you already have an Amazon Virtual Private Cloud (Amazon VPC) but need an Amazon EKS cluster _**with**_ the Snyk controller deployed to the cluster

[![cloudformation-launch-stack.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](../../../.gitbook/assets/cloudformation-launch-stack.png)](https://us-east-2.console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/template?stackName=Amazon-EKS-with-Snyk\&templateURL=https://aws-quickstart.s3.us-east-1.amazonaws.com/quickstart-amazon-eks/templates/amazon-eks-master-existing-vpc.template.yaml)

**3.** If you have neither an Amazon VPC or Amazon EKS cluster and need all services _**with**_ the Snyk controller deployed to the cluster

[![cloudformation-launch-stack.png - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](../../../.gitbook/assets/cloudformation-launch-stack.png)](https://us-east-2.console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/template?stackName=Amazon-EKS-with-Snyk\&templateURL=https://aws-quickstart.s3.us-east-1.amazonaws.com/quickstart-amazon-eks/templates/amazon-eks-master-existing-vpc.template.yaml)

**NOTE:** _Review the_ [_parameter reference_](https://github.com/aws-quickstart/quickstart-eks-snyk#parameter-reference) _prior to deployment._

## Snyk Security on AWS

When configuring Snyk to integrate with an Amazon Elastic Kubernetes Services (EKS) cluster, if you wish to scan images hosted on your Amazon Elastic Container Registry (ECR), you may also deploy the quick start [Snyk Security on AWS](https://aws.amazon.com/quickstart/architecture/snyk-security/) to enable this integration.

Refer to [Snyk Controller for Amazon EKS Quick Start Reference Deployment](https://aws-quickstart.github.io/quickstart-eks-snyk/).
