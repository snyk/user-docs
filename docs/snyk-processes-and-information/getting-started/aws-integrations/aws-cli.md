# AWS CLI

## Install the AWS Command Line Interface (AWS CLI)

You will use the [AWS Command Line Interface (AWS CLI)](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html) as well as `kubectl` and `eksctl` for these exercises. Detailed documentation on installing the AWS CLI is available on [AWS Documentation](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html) as well as a [Getting started with eksctl guide](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html) that details installation steps for Linux and Windows. These examples will be based on macOS. Sample code, templates, and other resources are provided in a [Bitbucket repository](https://bitbucket.org/snyk/patterns-library-atlassian-aws) that accompanies this workshop.

### Verify the AWS CLI is installed

Let's make sure the AWS CLI installed correctly.

```bash
aws --version
```

{% hint style="info" %}
Remember that you will need to run `aws configure` to setup your AWS credentials.
{% endhint %}

##
