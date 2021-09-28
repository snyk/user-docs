# Module 2

## Red Hat Quay private container registry

You can use Quay.io to automate your container builds, with integration to GitHub, Bitbucket, and more. Robot accounts allow you to lock down automated access and audit each deployment. In these examples, we will use a Quay.io cloud account to store our images which we will build, scan and push to the registry using the CLI.

### Scanning container images for vulnerabilities

First, we will [`docker login`](https://docs.docker.com/engine/reference/commandline/login/) into our Quay registry. If you are new to Quay, it is recommended that you complete the [tutorial](https://quay.io/tutorial/) where you will learn how to generate an encrypted password to use with the Docker CLI.

{% hint style="info" %}
The Docker CLI stores passwords entered on the command line in **plaintext**. It is therefore highly recommended to generate an an encrypted version of your password to use for `docker login`.
{% endhint %}

Run the command below, substituting `<QUAY_USER>` with your unique Quay.io username and `<GUID>` with your encrypted password.

```bash
docker login -u="<QUAY_USER>" -p="<GUID>" quay.io
```

Next, from your terminal, navigate to the directory of the `goof` repo you previously cloned. At the root of this directory you will see a [`Dockerfile`](https://github.com/snyk/goof/blob/master/Dockerfile). We will build and tag an image from this Dockerfile using the [`docker build`](https://docs.docker.com/engine/reference/commandline/build/) command as show below. Again, replace `<QUAY_USER>` with your username.

```bash
docker build -t <QUAY_USER>/goof .
```

At this point, you would normally push your image to a registry. However, we will perform one additional and critical step before we do so. We will scan our container image using Snyk to identify vulnerabilities and provide base image recommendations. To do this, we will run the following command in our terminal:

```bash
snyk monitor --docker quay.io/<QUAY_USER>/goof \
--file=submodules/goof/Dockerfile --org=<SNYK_ORG> \
--project-name=quay-cli-container-image
```

Let's take a moment and break down what we are doing in the above command. The `monitor` command will record the state of dependencies and any vulnerabilities to your account on [snyk.io](https://snyk.io) so you can later review those results.

![](../../.gitbook/assets/quay-cli-container-scan.gif)

Once complete, we are ready to push the image to our Quay registry and will do so by invoking the [`docker push`](https://docs.docker.com/engine/reference/commandline/push/) command as show below:

```bash
docker push quay.io/<QUAY_USER>/goof
```

Upon successful completion you should see results similar to the following:

![](../../.gitbook/assets/docker-push.gif)

We can optionally verify our image was indeed successfully pushed by logging into Quay.io and confirming the appropriate image has been tagged.

![](../../.gitbook/assets/rh-quay.png)

When we ran `snyk monitor` against our Dockerfile, we passed two parameters: `--org` and `--project-name`. This action allowed the results from our CLI to be saved to our Snyk account. A report of the findings is now available by logging into your Snyk account and viewing the imported projects.

![](../../.gitbook/assets/container-image-scan-results.png)

To receive base image remediation advice, including major, minor and alternative upgrades as well as advice when you need to rebuild your image, integrate with your preferred Git repository and import the repo that contains the relevant Dockerfile. You also need to [add the Dockerfile to the image](https://support.snyk.io/hc/en-us/articles/360003916218-Adding-your-Dockerfile-and-test-your-base-image). 

Once complete, you will see **recommendations for base image upgrade**.

![](../../.gitbook/assets/base-image-recommend.png)

