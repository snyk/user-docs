# Configure Environment

## Install the AWS Command Line Interface \(AWS CLI\)

You will use the [AWS Command Line Interface \(AWS CLI\)](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html) as well as `kubectl` and `eksctl` for these exercises. Detailed documentation on installing the AWS CLI is available on [AWS Documentation](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html) as well as a [Getting started with eksctl guide](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html) that details installation steps for Linux and Windows. These examples will be based on macOS. Sample code, templates, and other resources are provided in a [Bitbucket repository](https://bitbucket.org/snyk/patterns-library-atlassian-aws) that accompanies this workshop.

### Verify the AWS CLI is installed

Let's make sure the AWS CLI installed correctly.

```bash
aws --version
```

{% hint style="info" %}
Remember that you will need to run `aws configure` to setup your AWS credentials.
{% endhint %}

## Installing eksctl, the official CLI for Amazon EKS

We are going to install `eksctl` with [Homebrew](https://docs.brew.sh/Installation.html).

```bash
brew tap weaveworks/tap && \
brew update && \
brew install weaveworks/tap/eksctl
```

{% hint style="info" %}
Installing `eksctl` with Homebrew will also install the `kubectl` command-line utility. 
{% endhint %}

## Installing the Snyk CLI

The Snyk CLI authenticates your machine with your Snyk account. This tool will help you find and fix known vulnerabilities in your dependencies, both manually and also as part of your continuous integration build server.

```bash
brew tap snyk/tap && \
brew update && \
brew install snyk
```

To associate your Snyk account with the CLI, you must first authenticate your machine. No repository permissions are needed at this stage. Simply run the following command:

```bash
snyk auth
```

A web browser tab will open, redirecting you to authenticate the CLI for use with your account. Click `Authenticate`. When the authentication has completed, you may return to your terminal and continue working.

## Creating an Amazon ECR repository

Before you can push your Docker images to Amazon ECR, you must [create a repository](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-create.html) to store them in. You can create Amazon ECR repositories with the AWS Management Console, or with the AWS CLI and AWS SDKs. For this workshop, we will create the repository with the AWS Management Console:

1. Open the Amazon ECR console at [https://console.aws.amazon.com/ecr/repositories](https://console.aws.amazon.com/ecr/repositories).
2. From the navigation bar, choose the Region to create your repository in.
3. In the navigation pane, choose Repositories.
4. On the Repositories page, choose Create repository.
5. For Repository name, enter a unique name for your repository.
6. For Tag immutability, choose the tag mutability setting for the repository. Repositories configured with immutable tags will prevent image tags from being overwritten. For more information, see Image tag mutability.
7. For Scan on push, choose the image scanning setting for the repository. Repositories configured to scan on push will start an image scan whenever an image is pushed, otherwise image scans need to be started manually. For more information, see Image scanning.
8. Choose Create repository.

## Additional lab resources

The exercises contained in this workshop will include a combination of commands or code snippets that will be shared within the specific module pages as well as templates and source code available in a public Bitbucket repository. Once your Bitbucket Cloud account is setup, you will copy these resources into your account. To do so, please follow these steps:

### Step 1 - Fork the repository

Click [here](https://bitbucket.org/snyk/patterns-library-atlassian-aws/fork) to fork the `upstream` repository into your Bitbucket account. For detailed instructions on how to [fork a respository](https://support.atlassian.com/bitbucket-cloud/docs/fork-a-repository/), please review Atlassian's documentation.

### Step 2 - Clone your fork locally

Once **Step 1** is complete, you will need to `clone` your forked repository. Please review Atlassian's documentation on how to [clone a repository](https://confluence.atlassian.com/bitbucket/clone-a-repository-223217891.html) for detailed instructions.

