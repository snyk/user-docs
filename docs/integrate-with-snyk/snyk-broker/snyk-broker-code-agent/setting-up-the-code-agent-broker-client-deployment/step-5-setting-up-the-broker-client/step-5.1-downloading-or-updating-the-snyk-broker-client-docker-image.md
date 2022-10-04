# Step 5.1: Downloading or Updating the Snyk Broker Client – Docker image

**Note**: Before downloading the Broker Client Docker image, you must verify that your machine can run Docker containers.

The first step in setting up the Broker Client is pulling the Broker Client Docker image from Docker Hub. The Broker Client image should be downloaded to each machine that will run the Broker Client. After downloading it, this image is cached on the host machine.

The Docker image of the Broker Client is located at Docker Hub: [https://hub.docker.com/r/snyk/broker](https://hub.docker.com/r/snyk/broker)

**Note:** The Code Agent is only supported in the Broker Client version 4.108.0 and later versions. If you already have a running Broker Client, pull the latest update.

**To pull the Code Agent Docker image:**

* Run:

```
docker pull snyk/broker:<SCM_tag>
```

Where:

`SCM_tag` – a specific tag for each integrated SCM, as follows:

* **GitHub**:

```
docker pull snyk/broker:github-com
```

* **GitHub Enterprise**:

```
docker pull snyk/broker:github-enterprise
```

* **GitLab**:

```
docker pull snyk/broker:gitlab
```

* **Bitbucket Server/Data Center**:

```
docker pull snyk/broker:bitbucket-server
```

* **Azure Repo**:

```
docker pull snyk/broker:azure-repos
```

The download process of the Docker image of the Client Broker image starts.

For example:

<figure><img src="../../../../../.gitbook/assets/Client Broker - Pull image - example.png" alt=""><figcaption></figcaption></figure>

&#x20;Once the download is completed, you can verify that the Broker Client image was successfully downloaded to your machine.

**To verify the successful download of the Docker image of the Broker Client:**

* Run the Docker List command:

```
docker image ls
```

Your output should look similar to the following:

```
REPOSITORY           TAG                 IMAGE ID       CREATED       SIZE
snyk/broker          github-com          e999988aa7b7   7 days ago    252MB
snyk/broker          github-enterprise   0a8b4e6f518d   7 days ago   252MB
```

