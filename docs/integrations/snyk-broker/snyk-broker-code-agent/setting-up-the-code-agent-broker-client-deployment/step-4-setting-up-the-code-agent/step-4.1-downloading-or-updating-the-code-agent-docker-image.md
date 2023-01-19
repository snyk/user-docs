# Step 4.1: Downloading or Updating the Code Agent – Docker image

{% hint style="info" %}
Before downloading the Code Agent Docker image, you must verify that your machine can run Docker containers.
{% endhint %}

## **Downloading the Code Agent Docker Image**

First pull the Code Agent Docker image from Docker Hub. Download the Snyk Code Agent Docker image to each machine that will run the Code Agent. Docker images are usually cached on the host machine.

The Docker image of the Code Agent is located at [Docker Hub](https://hub.docker.com/r/snyk/code-agent/).

**To pull the Code Agent Docker image**, in the terminal, enter:

```
docker pull snyk/code-agent
```

The download process starts for the Code Agent Docker image, for example:

<figure><img src="../../../../../.gitbook/assets/Code Agent - Pull docker image - New.png" alt="Download process for Code Agent Docker image"><figcaption><p>Download process for Code Agent Docker image</p></figcaption></figure>

## Updating the Code Agent Docker Image

Pull the Code Agent Docker image again. If you are using the `latest` tag, the image is automatically updated. Otherwise provide a new image tag:

```
docker pull snyk/code-agent:<image_tag>
```

Remove or stop the older Code Agent container.

Follow the steps in [Running the Code Agent container](step-4.2-running-the-code-agent-container.md) to start the new Code Agent container.
