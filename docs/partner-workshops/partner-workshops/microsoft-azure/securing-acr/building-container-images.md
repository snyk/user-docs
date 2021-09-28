# Building container images

### Background

In **Securing AKS with Snyk**, we deployed an application to our Kubernetes cluster running on Microsoft Azure Kubernetes Service \(AKS\). We deployed this using a [manifest](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/) template file. The manifest provides a means to organize resource configurations that simplify our deployments. If we examine the `azure-vote.yaml` we see that we are defining the container images for our back-end and front-end applications as pulling `redis` and `microsoft/azure-vote-front:v1` respectively.

Later, when we scanned our Kubernetes workload, we found vulnerabilities in our Kubernetes security configuration as well as our container images. We were able to fix the security configuration issues detected. However, when we examined our container images, we noticed the following alert:

![](../../../.gitbook/assets/snyk_scan_06.png)

In this module, we will cover some best practices for building container images, storing these in a private registry like ACR, and monitoring those registries with Snyk.

### Building a Docker image

In Snyk’s [State of open source security report – 2019](https://snyk.io/blog/top-ten-most-popular-docker-images-each-contain-at-least-30-vulnerabilities/), we found that many of the popular Docker containers that are featured on the Docker Hub website are bundling images that contain many known vulnerabilities. We encourage you to read this report and as well as our recent blog, [10 Docker Image Security Best Practices](https://snyk.io/blog/10-docker-image-security-best-practices/) for a deeper dive on the subject. The following exercises are aligned with these best practices and we will cover them in stages starting with the concept of `minimal images`.

#### Dockerfile

Let's start with our back-end application. We see from our manifest that we are using the [Docker Official Image](https://docs.docker.com/docker-hub/official_repos/) for [redis](https://hub.docker.com/_/redis). We will take the first step in improving our security posture by building a fresh image from a `Dockerfile` and storing this in our private registry on ACR.

From the terminal, let's make sure we are at the root of our cloned repository by typing `pwd` and verifying the result displays `$HOME/snyk-azure-resources/`. Our repo contains a [`submodule`](https://git-scm.com/book/en/v2/Git-Tools-Submodules) that points to the source for the official redis image which happens to include a Dockerfile. To synchronize the submodule, type the following command:

```bash
git submodule update --recursive
```

Then, let's change directory by typing the following command:

```bash
cd app/redis/6.0/
```

From here, a simple `ls -a` command will show the contents and you will see a `Dockerfile` in that directory. You can open this in your favorite editor like [Microsoft Visual Studio Code](https://code.visualstudio.com/) and review the contents. The top few lines should look similar to this:

```text
FROM debian:buster-slim

# add our user and group first to make sure their IDs get assigned consistently, regardless of whatever dependencies get added
RUN groupadd -r -g 999 redis && useradd -r -g redis -u 999 redis
```

We will come back to this `Dockerfile` later, but for now, take note that the underlying base image is `debian:buster-slim`.

#### Docker build

Now, from the `app/redis/6.0/` directory, we will build and tag our own container image using the provided `Dockerfile`. To do so, we will run the [`docker build`](https://docs.docker.com/engine/reference/commandline/build/) command as follows:

```bash
docker build -t my-redis:v1 .
```

Upon successful completion we should see results similar to the following:

```text
Successfully built aa8130687a13
Successfully tagged my-redis:v1
```

#### Docker tag

We are almost ready to push our initial image to an Azure Container registry. However, before we proceed, we will need to run a few [`docker tag`](https://docs.docker.com/engine/reference/commandline/tag/) commands to tag our image with the fully qualified name of your ACR login server. We are going to make this interesting by querying the value of our ACR login server using the Azure CLI within our `docker tag` command. To do this, we will invoke `az acr show` and specify our registry name while formatting our output with the `-o tsv` parameter for plain text.

```bash
docker tag my-redis:v1 $(az acr show --name mySnykContainerRegistry --query loginServer --output tsv)/my-redis:v1
```

We tagged the first with a `version` and we will tag it a second time with `latest`.

```bash
docker tag my-redis:v1 $(az acr show --name mySnykContainerRegistry --query loginServer --output tsv)/my-redis:latest
```

Now, we will run a quick [`docker images`](https://docs.docker.com/engine/reference/commandline/images/) and make sure everything is correctly tagged.

```text
REPOSITORY                                                      TAG                 IMAGE ID            CREATED             SIZE
my-redis                                                        v1                  aa8130687a13        4 hours ago         104MB
mysnykcontainerregistry.azurecr.io/my-redis                     latest              aa8130687a13        4 hours ago         104MB
mysnykcontainerregistry.azurecr.io/my-redis                     v1                  aa8130687a13        4 hours ago         104MB
debian                                                          buster-slim         e5aad4204d00        7 days ago          69.2MB
```

#### Docker push

We made it. We are now ready to run [`docker push`](https://docs.docker.com/engine/reference/commandline/push/) on our tagged images and store these in ACR. Let's start with the first tagged image `v1`:

```bash
 docker push $(az acr show --name mySnykContainerRegistry --query loginServer --output tsv)/my-redis:v1
```

You should see output similar to what's below:

```text
The push refers to repository [mysnykcontainerregistry.azurecr.io/my-redis]
555634e64cd8: Pushed
12534748095c: Pushed
ea013d725bd6: Pushed
63689ea7c92e: Pushed
2c2aedafe0c2: Pushed
c2adabaecedb: Pushed
v1: digest: sha256:f8e1a610528d3fbd6f1b26fc2a2610c05fd07938b68e8aef7cb87ef1c9a6ec65 size: 1573
```

Next, our second tagged image, `latest`:

```bash
docker push $(az acr show --name mySnykContainerRegistry --query loginServer --output tsv)/my-redis:latest
```

You should see output similar to what's below:

```text
555634e64cd8: Layer already exists
12534748095c: Layer already exists
ea013d725bd6: Layer already exists
63689ea7c92e: Layer already exists
2c2aedafe0c2: Layer already exists
c2adabaecedb: Layer already exists
latest: digest: sha256:f8e1a610528d3fbd6f1b26fc2a2610c05fd07938b68e8aef7cb87ef1c9a6ec65 size: 1573
```

We can also verify these were indeed pushed to our registry by querying through the Azure CLI. 

```bash
az acr repository list --name mySnykContainerRegistry --output json
```

You should see output similar to what's below:

```text
[
  "my-redis"
]
```

To verify the list of tags in our `my-redis` repository matches what we pushed, we can run the following command:

```bash
az acr repository show-tags --name mySnykContainerRegistry --repository my-redis --output table
```

We should see output similar to what's below:

```text
Result
--------
latest
v1
```

Of course, we can also view the same results from the Azure portal:

![](../../../.gitbook/assets/acr_repository_01.png)

