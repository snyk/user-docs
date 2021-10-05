# Creating your environment

{% hint style="danger" %}
You are responsible for the cost of the AWS services used while running this workshop in your AWS account.
{% endhint %}

In order for you to succeed in this workshop, you will need to run through a few steps in order to properly setup and configure your environment. These steps will include provisioning some services, installing some tools, and downloading some dependencies as well. We will begin with [AWS Cloud9](https://aws.amazon.com/cloud9/). Technically, you should be able to complete many of the steps in these modules if you have a properly configured terminal. However, in order to avoid the _"works on my machine"_ response you've surely experienced at some point in your career, I strongly encourage you to proceed with launching Cloud9.

[AWS Cloud9](https://aws.amazon.com/cloud9/) is a cloud-based integrated development environment \(IDE\) that lets you write, run, and debug your code with just a browser. It includes a code editor, debugger, and terminal. Cloud9 comes prepackaged with essential tools for popular programming languages, including JavaScript, Python, PHP, and more, so you don’t need to install files or configure your development machine to start new projects.

## Deploy & Launch AWS Cloud9

[Click here to deploy using CloudFormation template](https://console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/new?stackName=ModernizationWorkshop&templateURL=https://modernization-workshop-west-2.s3-us-west-2.amazonaws.com/devops/cfn/modernization-workshop.yaml)

* Create stack click, **Next**
* Specify stack details, click **Next**
* Configure stack options, click **Next**
* Review UnicornDevSecOpsWorkshop, scroll to bottom section under **Capabilities** and check both boxes and click **Create stack** 

> The deployment process takes approximately 2-3 minutes to complete. In the meantime, you can review the [deployment guide](https://aws-quickstart.s3.amazonaws.com/quickstart-cloud9-ide/doc/aws-cloud9-cloud-based-ide.pdf) while you wait.

Once the installation is complete, go to Cloud9 within the console and click on **Open IDE** on the name that begins with WorkshopIDE.

## Clone the source repository for this workshop

Now we want to clone the repository that contains all the content and files you need to complete this workshop.

```bash
cd ~/environment && \
git clone https://github.com/jamesbland123/workshop-sample modernization-workshop
cd modernization-workshop
git submodule init
git submodule update
```

## Increase AWS Cloud9 disk/storage

```bash
cd ~/environment/modernization-workshop/modules/20_getting_started
./resize.sh 50
```

## Update and install some tools

This step updates and installs various tools needed to complete the workshop. Feel free to look at the script if you are curious about what gets updated and installed.

```bash
./getting_started.sh
```

Next, lets source .bashrc to add .net PATH to our current working environment

```bash
. ~/.bashrc
```

