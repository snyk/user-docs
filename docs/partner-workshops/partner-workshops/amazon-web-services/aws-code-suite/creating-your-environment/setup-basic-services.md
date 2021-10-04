# Setup basic services

## Introduction

Up until now, we have been going through various steps to setup our environment. Installing tools and other necessary steps to make sure we progress through the modules without any issues. Now, we are ready to begin deploying the infrastructure that will support our Unicorn Store application.

## Basic Services CloudFormation Stack

We are going to setup some basic services such as an AWS CodeCommit and Amazon ECR services.

{% hint style="info" %}
This step takes approximately 2 minutes.
{% endhint %}

Copy and paste the following into Cloud9's terminal to launch a CloudFormation stack

```bash
cd ~/environment/modernization-workshop/modules/30_workshop_app
aws cloudformation create-stack --stack-name WorkshopServices --template-body file://services.yaml --capabilities CAPABILITY_NAMED_IAM
until [[ `aws cloudformation describe-stacks --stack-name "WorkshopServices" --query "Stacks[0].[StackStatus]" --output text` == "CREATE_COMPLETE" ]]; do  echo "The stack is NOT in a state of CREATE_COMPLETE at `date`";   sleep 30; done && echo "The Stack is built at `date` - Please proceed"
```

The output should look like the window below

```text
The stack is NOT in a state of CREATE_COMPLETE at Sun Sep  8 05:53:33 UTC 2019
The Stack is built at Sun Sep  8 05:54:04 UTC 2019 - Please proceed 
```

