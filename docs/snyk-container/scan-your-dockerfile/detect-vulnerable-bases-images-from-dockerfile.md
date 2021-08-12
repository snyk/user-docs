# Detect vulnerable bases images from Dockerfile

##  Detect vulnerable bases images from Dockerfile

Snyk detects vulnerable base images by scanning your Dockerfile when importing a Git repository. This allows you to examine security issues before building the image, so helps solve potential problems before they land in your registry or in production. 

After you [integrate your Git repository to Snyk](https://support.snyk.io/hc/en-us/sections/360001138098-Git-repository-SCM-integrations), any Dockerfiles in it are automatically picked up and surfaced in the web UI as ‘projects’.

![mceclip0.png](https://support.snyk.io/hc/article_attachments/360012913838/mceclip0.png)

### Linking from a Dockerfile to its container images

You can also link from a Dockerfile to all container images built from it. This can be used to understand the security impact on your running applications, and understand which images can be better secured or need to be rebuilt when taking action and updating the Dockerfile base image.

![mceclip3.png](https://support.snyk.io/hc/article_attachments/360012839537/mceclip3.png)

See [Base image detection](https://support.snyk.io/hc/en-us/articles/4405540674577-Base-image-detection) for more details about detecting vulnerable base images and remediation recommendations

