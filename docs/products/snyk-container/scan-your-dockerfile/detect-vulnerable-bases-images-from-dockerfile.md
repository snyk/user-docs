# Detect vulnerable base images from Dockerfile

Snyk detects vulnerable base images by scanning your Dockerfile when importing a Git repository. This allows you to examine security issues before building the image, so helps solve potential problems before they land in your registry or in production.

After you [integrate your Git repository to Snyk](https://support.snyk.io/hc/en-us/sections/360001138098-Git-repository-SCM-integrations), any Dockerfiles in it are automatically picked up and surfaced in the web UI as ‘projects’.

![](../../../.gitbook/assets/mceclip0-5-.png)

## Linking from a Dockerfile to its container images

You can also link from a Dockerfile to all container images built from it. This can be used to understand the security impact on your running applications, and understand which images can be better secured or need to be rebuilt when taking action and updating the Dockerfile base image.

![](../../../.gitbook/assets/mceclip3.png)

See [Base image detection](https://docs.snyk.io/snyk-container/getting-around-the-snyk-container-ui/base-image-detection) for more details about detecting vulnerable base images and fix recommendations
