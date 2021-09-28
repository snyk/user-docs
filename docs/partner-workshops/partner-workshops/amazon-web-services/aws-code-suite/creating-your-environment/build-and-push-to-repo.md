# Build & push to repo

## Build & Push to AWS CodeCommit

We should now be ready to push our application to the AWS CodeCommit repo and the Amazon Elastic Container Repository

The following sets a new origin for the application repo to CodeCommit unicorn-store, configures a credential helper needed for CodeCommit, and pushes the source code to the repo. This step is necessary for an automated pipeline as CodeBuild will build the application directly from this repo.

```bash
cd ~/environment/modernization-workshop/
git remote set-url origin https://git-codecommit.us-west-2.amazonaws.com/v1/repos/modernization-workshop
git config --global credential.helper '!aws codecommit credential-helper $@'
git config --global credential.UseHttpPath true
git push origin master
```

{% hint style="info" %}
If successfully, you should see a similar message to the one below.
{% endhint %}

```text
Counting objects: 9525, done.
Compressing objects: 100% (5900/5900), done.
Writing objects: 100% (9525/9525), 33.75 MiB | 2.65 MiB/s, done.
Total 9525 (delta 3240), reused 9525 (delta 3240)
remote: processing 
To https://git-codecommit.us-west-2.amazonaws.com/v1/repos/modernization-workshop
 * [new branch]      master -> master
```

## Push to Amazon Elastic Container Repository \(ECR\)

Now it's time to compile and package your code. Copy and paste the below code into Cloud9's terminal window

```bash
cd ~/environment/modernization-workshop/app
docker build -t modernization-workshop .
docker tag modernization-workshop:latest $(aws ecr describe-repositories --repository-name modernization-workshop --query=repositories[0].repositoryUri --output=text):latest
eval $(aws ecr get-login --no-include-email)
docker push $(aws ecr describe-repositories --repository-name modernization-workshop --query=repositories[0].repositoryUri --output=text):latest
```

If you watch the screen you should see the docker image build process animating the terminal.

{% hint style="info" %}
If successfully, you should see the message as below. 
{% endhint %}

```text
The push refers to repository [1234567891011.dkr.ecr.us-west-2.amazonaws.com/modernization-workshop]
8d2f7b95f78d: Pushed 
82852e5eaa9d: Pushed 
9df07df94e41: Pushed 
aa90bcce39de: Pushed 
d9ff549177a9: Pushed 
latest: digest: sha256:4229b5fe142f6d321ef2ce16ff22070e410272ee140e7eec51540a823dcd315a size: 1369
```

