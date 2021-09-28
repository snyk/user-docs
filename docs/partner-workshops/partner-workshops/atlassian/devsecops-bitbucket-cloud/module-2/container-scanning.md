# Container Scanning

## Background

![](../../../../.gitbook/assets/snyk-container-01.png)

### Testing for vulnerabilities at different stages of the SDLC\*

There are a few strategies for implementing testing. The following chart provides a helpful summary of these. The examples in this workshop will focus on two of those presented below.

| Stage | Description | Cost | Feedback | Completeness |
| :--- | :--- | :--- | :--- | :--- |
| Local | Great for debugging and building up knowledge among developers but requires individual developer action, and no way to enforce | Medium | Fast | Low |
| CI/CD\*\* | Great as a gate, fast feedback for developers, requires per-pipeline implementation which will depend on how standardized pipeline management is in your organization. Possibly counter productive for low severity issues breaking the build, so need other feedback cycles as well. | Medium | Fast | Variable |
| Registry\*\*\* | Often a single owner, so easy to integrate and covers all first-party images, no matter how they were built. Potentially noisy because images may be unused. | Low | Medium | Medium |
| Production | An accurate picture of what you’re running, including third-party content, but potentially slow feedback cycles to development teams and the risk vulnerabilities can be exploited in running applications. | High | Slow | High |

## Container image scanning in your pipeline

Similar to the previous module on **Application Scanning**, in this module we are configuring our [bitbucket-pipelines.yml](https://bitbucket.org/snyk/patterns-library-atlassian-aws/src/192a4d2412a4330b9f634e9d45a546ec1add61fb/bitbucket-pipelines.yml#lines-32:56) file to build the Docker image for our application, scan the image, then push that image to our registry. Let's take a closer look at the container image scanning step:

```yaml
scan-push-image: &scan-push-image
  - step:
      name: "Scan and push container image"
      services:
        - docker
      script:
        - docker build -t $IMAGE ./app/goof/
        - docker tag $IMAGE $IMAGE:${BITBUCKET_COMMIT}
        - pipe: snyk/snyk-scan:0.4.3
          variables:
            SNYK_TOKEN: $SNYK_TOKEN
            LANGUAGE: "docker"
            IMAGE_NAME: $IMAGE
            PROJECT_FOLDER: "app/goof"
            TARGET_FILE: "Dockerfile"
            CODE_INSIGHTS_RESULTS: "true"
            SEVERITY_THRESHOLD: "high"
            DONT_BREAK_BUILD: "true"
            MONITOR: "false"
        - pipe: atlassian/aws-ecr-push-image:1.1.3
          variables:
            AWS_ACCESS_KEY_ID: '$AWS_ACCESS_KEY_ID'
            AWS_SECRET_ACCESS_KEY: '$AWS_SECRET_ACCESS_KEY'
            AWS_DEFAULT_REGION: '$AWS_DEFAULT_REGION'
            IMAGE_NAME: $IMAGE
            TAGS: '${BITBUCKET_COMMIT}'
```

Here, we are building our container image and tagging it, then leveraging the [Snyk Scan](https://bitbucket.org/product/features/pipelines/integrations?p=snyk/snyk-scan) pipe in our pipeline to perform a scan of the container image. We will keep the same values for `CODE_INSIGHTS_RESULTS`, `SEVERITY_THRESHOLD`, and `DONT_BREAK_BUILD`. We are also passing a few additional [supported variables](https://bitbucket.org/snyk/snyk-scan) relevant for our Snyk Pipe to understand we are requesting a container image scan instead of an application scan. These are namely setting `LANGUAGE` to `docker`, declaring the `IMAGE_NAME` and passing the appropriate **repository variable** as well as setting the `TARGET_FILE` to `Dockerfile`. Once the image has been scanned by Snyk and results available in Code Insights, we will call the Amazon ECR [push image pipe](https://bitbucket.org/atlassian/aws-ecr-push-image) to push to our repository.

**Footnotes:**

1. **\*Snyk Blog -** [Testing for vulnerabilities at different stages of the SDLC](https://snyk.io/blog/container-security-throughout-the-sdlc/)
2. **\*\***We will focus on this stage in our Bitbucket pipeline example.
3. **\*\*\***We will focus on this stage in our Snyk integration for Amazon ECR example.

