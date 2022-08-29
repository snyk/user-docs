# Step 4.1: Downloading or Updating the Code Agent – Docker image

**Note**: Before downloading the Code Agent Docker image, you must verify that your machine can run Docker containers.

The first step in setting up the Code Agent is pulling the Code Agent Docker image from Docker Hub. The Snyk Code Docker image should be downloaded to each machine that will run the Code Agent. After downloading it, this image is cached on the host machine. If you already have a Docker image of the Code Agent on your machine, downloading the Docker image again will update your existing version.

The Docker image of the Code Agent is located at Docker Hub: [https://hub.docker.com/r/snyk/code-agent/](https://hub.docker.com/r/snyk/code-agent/)

**To pull the Code Agent Docker image:**

* In the terminal, enter:

```
docker pull snyk/code-agent
```

The download process of the Code Agent Docker image starts.

For example:

<figure><img src="../../../../../.gitbook/assets/Code Agent - Pull docker image - New.png" alt=""><figcaption></figcaption></figure>

****
