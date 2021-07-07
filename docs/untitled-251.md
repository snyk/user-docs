# Detect vulnerable bases images from Dockerfile

##  Detect vulnerable bases images from Dockerfile

Snyk detects vulnerable base images by scanning your Dockerfile when importing a Git repository. This allows you to examine security issues before building the image, so helps solve potential problems before they land in your registry or in production. 

After you [integrate your Git repository to Snyk](https://support.snyk.io/hc/en-us/sections/360001138098-Git-repository-SCM-integrations), any Dockerfiles in it are automatically picked up and surfaced in the web UI as ‘projects’.

In a Dockerfile project, you can find the relevant metadata of the Dockerfile and the base image used. If the base image is an [Official Docker image](https://docs.docker.com/docker-hub/official_images/), the results include recommendations for upgrades to resolve some of the discovered vulnerabilities.

This allows you to see vulnerability counts in minor and major upgrades, as well as in alternative base images which might have fewer vulnerabilities. With that, you can decide whether to upgrade your base image, and which one will be the best fit.

See the section on [base image recommendations](https://support.snyk.io/hc/en-us/articles/360003915938-Analysis-and-remediation-for-your-images-from-the-Snyk-app) for more information. 

You can find the base image vulnerabilities in your Dockerfile project, sorted by their priority score, together with useful filters to help you extract issues.

You can link from a Dockerfile to all container images built from it. This can be used to understand the security impact on your running applications, and understand which images can be better secured or need to be rebuilt when taking action and updating the Dockerfile base image.

