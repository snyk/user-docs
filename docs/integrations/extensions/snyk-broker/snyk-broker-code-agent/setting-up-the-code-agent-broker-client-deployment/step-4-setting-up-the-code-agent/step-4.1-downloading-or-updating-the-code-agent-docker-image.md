# Step 4.1: Downloading or Updating the Code Agent – Docker image

**Note**: Before downloading the Code Agent Docker image, you must verify that your machine can run Docker containers.

#### **Downloading the Code Agent Docker Image**

First, pull the Code Agent Docker image from Docker Hub. The Snyk Code Agent Docker image should be downloaded to each machine that will run the Code Agent. Docker images are usually cached on the host machine.

The Docker image of the Code Agent is located at Docker Hub: [https://hub.docker.com/r/snyk/code-agent/](https://hub.docker.com/r/snyk/code-agent/).

**To pull the Code Agent Docker image:**

In the terminal, enter:

```
docker pull snyk/code-agent
```

The download process of the Code Agent Docker image starts.

For example:

<figure><img src="../../../../../../.gitbook/assets/Code Agent - Pull docker image - New.png" alt=""><figcaption></figcaption></figure>

#### Updating the Code Agent Docker Image

Pull the Code Agent Docker image again. If using the `latest` tag, the image is automatically updated - otherwise, provide a new image tag:

```
docker pull snyk/code-agent:<image_tag>
```

Remove or stop the older Code Agent container.

Follow the steps in [step-4.2-running-the-code-agent-container.md](step-4.2-running-the-code-agent-container.md "mention") to start the new Code Agent container.
