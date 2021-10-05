# Deploy ECS Fargate service

## Deploy Fargate Service

In the following set of commands we are going to use CloudFormation to deploy services that will allow our Unicorn Store application to service traffic from the Internet. The CloudFormation template sets up an ECS Cluster, a Service, Task Definition, Task, and Application Load Balancer.

```bash
cd ~/environment/modernization-workshop/modules/30_workshop_app
aws cloudformation create-stack --stack-name WorkshopECS --template-body file://ecs-fargate.yaml --capabilities CAPABILITY_NAMED_IAM
until [[ `aws cloudformation describe-stacks --stack-name "WorkshopECS" --query "Stacks[0].[StackStatus]" --output text` == "CREATE_COMPLETE" ]]; do  echo "The stack is NOT in a state of CREATE_COMPLETE at `date`";   sleep 30; done && echo "The Stack is built at `date` - Please proceed"
```

{% hint style="info" %}
This step takes approximately 3 minutes and if successfully, you should see the message as below.
{% endhint %}

```text
The stack is NOT in a state of CREATE_COMPLETE at Sun Aug  4 05:34:25 UTC 2019
The stack is NOT in a state of CREATE_COMPLETE at Sun Aug  4 05:34:55 UTC 2019
The stack is NOT in a state of CREATE_COMPLETE at Sun Aug  4 05:35:26 UTC 2019
The stack is NOT in a state of CREATE_COMPLETE at Sun Aug  4 05:35:57 UTC 2019
The stack is NOT in a state of CREATE_COMPLETE at Sun Aug  4 05:36:27 UTC 2019
The stack is NOT in a state of CREATE_COMPLETE at Sun Aug  4 05:36:58 UTC 2019
The Stack is built at Sun Aug  4 05:37:28 UTC 2019 - Please proceed
```

To test, run the following query and copy the URL you obtain from the output into the address bar of a web browser. You should see something similar to the image below.

```bash
aws elbv2 describe-load-balancers --names="Modernization-Workshop-LB" --query="LoadBalancers[0].DNSName" --output=text
```

