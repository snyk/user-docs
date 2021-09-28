# Cleanup

In order to prevent charges to your account we recommend cleaning up the infrastructure that was created. If you plan to keep things running so you can examine the workshop a bit more please remember to do the cleanup when you are done. It is very easy to leave things running in an AWS account, forget about it, and then accrue charges.

{% hint style="info" %}
You will need to manually delete some resources before you delete the CloudFormation stacks so please do the following steps in order. With the CloudFormation Stacks, delete one at a time and validate the stack is removed before deleting the next stack.
{% endhint %}

```bash
# Delete S3 Bucket
aws s3 rm s3://$(aws s3api list-buckets --query 'Buckets[?starts_with(Name, `workshoppipeline-artifactbucket`) == `true` ].Name' --output text) --recursive

# Delete Log Group
aws logs delete-log-group --log-group-name ModernizationWorkshop

# Delete ECR Repository
aws ecr delete-repository --repository-name modernization-workshop --force

# Delete CloudFormation Pipeline and ECS Stacks
aws cloudformation delete-stack --stack-name WorkshopPipeline
aws cloudformation delete-stack --stack-name WorkshopECS
```

Now remove the WorkshopServices stack

```bash
aws cloudformation delete-stack --stack-name WorkshopServices
```

Finally, close the cloud9 window and manually verify deletion of previous stacks and delete the final stack. In the AWS console, go to CloudFormation. Ensure **WorkshopPipeline, WorkshopECS, and WorkshopServices** have all been removed. Once verified, click `ModernizationWorkshop` stack and then `Delete`

Verify that none of the Workshop\* stacks are listed in CloudFormation and you are done.

