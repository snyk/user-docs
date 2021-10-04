---
title: Cleanup
chapter: true
weight: 100
---

# Cleanup

## AWS Cleanup

![](../../../.gitbook/assets/hardhat.png)

{% hint style="danger" %}
In order to prevent charges to your account we recommend cleaning up the infrastructure that was created. If you plan to keep things running so you can examine the workshop a bit more please remember to do the cleanup when you are done. It is very easy to leave things running in an AWS account, forget about it, and then accrue charges.
{% endhint %}

In the AWS console, go to [CloudFormation](https://us-west-2.console.aws.amazon.com/cloudformation/home?region=us-west-2), click the following:

![](../../../.gitbook/assets/aws-account-cleanup.png)

1. `Snyk-Security-ECR` stack and then `Delete`.  
2. `Amazon-EKS-with-Snyk` stack and then `Delete`.

You do not need to choose the stacks for the workshop with the **Nested** tag, they will automatically be deleted with the `Amazon-EKS-with-Snyk` stack.

{% hint style="warning" %}
This process will take 15-30 minutes, to be sure to verify that none of the `Amazon-EKS-with-Snyk` or `Snyk-Security-ECR` stacks are listed in CloudFormation and you are done.
{% endhint %}

