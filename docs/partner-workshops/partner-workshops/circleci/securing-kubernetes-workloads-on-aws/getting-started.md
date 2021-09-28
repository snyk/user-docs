# Getting started

You will use the [AWS Command Line Interface \(AWS CLI\)](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html) as well as `kubectl` and `eksctl` for these exercises. Detailed documentation on installing the AWS CLI is available on [AWS Documentation](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html) as well as a [Getting started with `eksctl` guide](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html) that details installation steps for Linux and Windows. These examples will be based on macOS. Sample code, templates, and other resources are provided in a [Git repository](https://github.com/snyk-partners/snyk-circleci-eks) that accompanies this workshop. You are encouraged to [`clone`](https://github.com/snyk-partners/snyk-circleci-eks.git) or [fork](https://github.com/snyk-partners/snyk-circleci-eks/fork) this repository as we will reference that content throughout these exercises.

### Configure the local environment

From your terminal, let's verify that the AWS CLI is installed:

```bash
aws --version
```

{% hint style="danger" %}
**Stop!** Remember that you will need to run [`aws configure`](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) to configure your AWS CLI credentials.
{% endhint %}

#### Install Homebrew

If you don't already have it, [install Homebrew](https://docs.brew.sh/Installation.html) then install `eksctl`.

```bash
brew tap weaveworks/tap && \
brew update && \
brew install weaveworks/tap/eksctl
```

{% hint style="info" %}
Installing `eksctl` with Homebrew will also install the `kubectl` command-line utility.
{% endhint %}

Now, let's validate that the installation was successful.

```bash
eksctl version
```

You are now ready to create your Amazon EKS cluster and worker nodes!

#### 

