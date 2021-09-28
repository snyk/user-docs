# Deploy CI/CD pipeline

## What is CI/CD?

Continuous integration \(CI\) and continuous delivery \(CD\) embodies a culture, set of operating principles, and collection of practices that enable application development teams to deliver features and functionality more frequently and reliably.

Continuous integration is a coding philosophy and set of practices that drive development teams to implement small changes and check in code to version control repositories as frequently as possible. Because most modern applications require developing code in different platforms and tools, the team needs a mechanism to integrate and validate its changes.

A goal of CI is to establish a consistent and automated way to build and test applications. With consistency in the integration process in place, teams are more likely to commit code changes more frequently, which leads to better collaboration and software quality through feedback loops.

Continuous delivery picks up where continuous integration ends. CD automates the delivery of applications to selected infrastructure environments. Most teams work with multiple environments other than the production, such as development and testing environments, and CD ensures there is an automated way to push code changes to them. CD automation in a modern application performs any necessary deployment tasks while adhering to principles such as "infrastructure as code" and immutability.

Continuous integration and delivery requires continuous testing because the objective is to deliver quality applications and code to users. Continuous testing is often implemented as a set of automated regression, performance, security, and other tests that are executed in the CI/CD pipeline.

A mature CI/CD practice has the option of implementing continuous deployment where application changes run through the CI/CD pipeline and passing builds are deployed directly to production environments.

## CI/CD pipeline being deployed in this workshop

![](../../../../.gitbook/assets/aws-pipeline.png)

1. Development and local testing: The developer will work on coding tasks with the Cloud9 IDE environment.  Once completed with the task the developer will commit her/his changes to the local git repository and test the changes.
2. Push to remote master branch: When the developer is satisfied with the software changes, the developer will push those changes to the remote master branch.  In this workshop this is the AWS CodeCommit Repo. 
3. AWS CodePipeline - Commit Event: CodePipeline will monitor AWS CodeCommit for any new commits.  When a new commit \(code change\) is detected a CodeBuild job will be triggered.
4. AWS CodePipeline - Build: Within CodeBuild a series of security tests will be instrumented to validate that code changes are not adding security risks to the application.  If security issues are detect, this phase of the CodePipeline process is ended and the pipeline process reports a failure status. If no security issues are detected the build process continues by building and packaging the application.
5. AWS CodePipeline - Postbuilld: Runs additional tests and if those test pass the container is pushed to Amazon ECR.  
6. AWS CodePipeline - Deploy: If the build process is successful, Codepipeline will update Elastic Container Service that there is a new image.  The new image will be deployed and monitored.  If all healthchecks pass the new deployment will become the primary service and the previous deployment will be gracefully shutdown. 
7. Monitor: ECS and the Application Load Balancer will continually monitor the health of the container and the application and will make the required adjustments to keep the minimum number of healthly container tasks running at all times.  

## Deploy CI/CD Pipeline

To deploy the pipeline, run the following commands in Cloud9's terminal

```bash
cd ~/environment/modernization-workshop/modules/30_workshop_app
aws cloudformation create-stack --stack-name WorkshopPipeline --template-body file://pipeline.yaml --capabilities CAPABILITY_NAMED_IAM
until [[ `aws cloudformation describe-stacks --stack-name "WorkshopPipeline" --query "Stacks[0].[StackStatus]" --output text` == "CREATE_COMPLETE" ]]; do  echo "The stack is NOT in a state of CREATE_COMPLETE at `date`";   sleep 30; done && echo "The Stack is built at `date` - Please proceed"
```

{% hint style="info" %}
This step takes approximately 1 minute and if successfully, you should see the message as below.
{% endhint %}

```text
The stack is NOT in a state of CREATE_COMPLETE at Sun Aug  4 05:46:27 UTC 2019
The stack is NOT in a state of CREATE_COMPLETE at Sun Aug  4 05:46:58 UTC 2019
The Stack is built at Sun Aug  4 05:47:29 UTC 2019 - Please proceed
```

At this point you should have a fully functioning CI/CD CodePipeline. If you head over to CodePipeline in the AWS console and click on the pipeline that begins with the name **WorkshopPipeline-Pipeline** you will see a similar screen to the one below.

![](../../../../.gitbook/assets/pipeline-view.png)

