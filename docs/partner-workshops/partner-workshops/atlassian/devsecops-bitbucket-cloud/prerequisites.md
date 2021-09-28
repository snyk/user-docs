# Prerequisites

{% hint style="danger" %}
You are **responsible** for the cost of the AWS services used while running this workshop in your AWS account. Your account **must** have the ability to create new IAM roles and scope other IAM permissions.
{% endhint %}

{% hint style="info" %}
If you already have an AWS account, and have IAM Administrator access, go to the **Provision AWS services** section below. 
{% endhint %}

## Create an account

### Step 1

If you don't already have an AWS account with Administrator access: [create one now](http://docs.aws.amazon.com/connect/latest/adminguide/gettingstarted.html#sign-up-for-aws)

### Step 2

Once you have an AWS account, ensure you are following the remaining workshop steps as an **IAM user** with administrator access to the AWS account: [Create a new IAM user to use for the workshop](https://console.aws.amazon.com/iam/home?region=us-east-1#/users$new)

### Step 3

Enter the user details: 

![](../../../.gitbook/assets/iam-1-create-user.png)

### Step 4

Attach the AdministratorAccess IAM Policy: 

![](../../../.gitbook/assets/iam-2-attach-policy%20%281%29.png)

### Step 5

Click to create the new user: 

![](../../../.gitbook/assets/iam-3-create-user.png)

### Step 6

Take note of the login URL and save: 

![](../../../.gitbook/assets/iam-4-save-url.png)

## Provision AWS services

We have simplified and automated the process for provisioning the necessary AWS services needed for this workshop. This is possible through the [Snyk controller for Amazon EKS](https://github.com/aws-quickstart/quickstart-eks-snyk) AWS Quick Start. By clicking on the **Launch Stack** button below, you will be redirected to the AWS CloudFormation console where you will be prompted to complete the following steps:

* Create stack, click **Next**
* Specify stack details, click **Next**
* Configure stack options, click **Next**
* Scroll to bottom section under **Capabilities** and check both boxes and click **Create stack** 

When you are ready, [click here to deploy](https://us-west-2.console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/create/template?stackName=Amazon-EKS-with-Snyk&templateURL=https://aws-quickstart.s3.us-west-2.amazonaws.com/quickstart-amazon-eks/templates/amazon-eks-master.template.yaml)!

{% hint style="warning" %}
The installation takes several minutes. Please continue to the next section. 
{% endhint %}

