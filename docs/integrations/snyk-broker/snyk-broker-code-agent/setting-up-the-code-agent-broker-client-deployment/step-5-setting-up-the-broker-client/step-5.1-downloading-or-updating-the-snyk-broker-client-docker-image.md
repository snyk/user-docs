# Step 5.1: Downloading or Updating the Snyk Broker Client – Docker image

{% hint style="info" %}
Before downloading the Broker Client Docker image, verify that your machine can run Docker containers.
{% endhint %}

The first step in setting up the Broker Client is pulling the Broker Client Docker image from [Docker Hub](https://hub.docker.com/r/snyk/broker). Download the Broker Client image to each machine that will run the Broker Client. After you download it, this image is cached on the host machine.

{% hint style="info" %}
The Code Agent is supported only in the Broker Client version 4.108.0 and later versions. If you already have a running Broker Client, pull the latest update.
{% endhint %}

**To pull the Code Agent Docker image** run:

```
docker pull snyk/broker:<SCM_tag>
```

where `SCM_tag` is a specific tag for each integrated SCM, as follows:

**GitHub**:

```
docker pull snyk/broker:github-com
```

**GitHub Enterprise**:

```
docker pull snyk/broker:github-enterprise
```

**GitLab**:

```
docker pull snyk/broker:gitlab
```

**Bitbucket Server/Data Center**:

```
docker pull snyk/broker:bitbucket-server
```

**Azure Repo**:

```
docker pull snyk/broker:azure-repos
```

The download process for the Docker image of the Client Broker image starts, for example:

<figure><img src="../../../../../.gitbook/assets/Client Broker - Pull image - example.png" alt="Download for Docker image of Client Broker"><figcaption><p>Download for Docker image of Client Broker</p></figcaption></figure>

When the download is completed, you can verify that the Broker Client image was successfully downloaded to your machine.

**To verify the successful download of the Docker image of the Broker Client** run the Docker List command:

```
docker image ls
```

Your output is similar to the following:

```
REPOSITORY           TAG                 IMAGE ID       CREATED       SIZE
snyk/broker          github-com          e999988aa7b7   7 days ago    252MB
snyk/broker          github-enterprise   0a8b4e6f518d   7 days ago   252MB
```
