# Application Scanning in your pipeline



The bitbucket-pipelines.yml defines your Pipelines builds configuration. If you're new to Pipelines you can learn more about how to get started [here](https://support.atlassian.com/bitbucket-cloud/docs/get-started-with-bitbucket-pipelines/).

In this workshop, we have provided you with a sample bitbucket-pipelines.yml file containing distinct steps mapped to our workflow:

* Build the code
* Build the docker image
* Scan the container image
* Deploy the container image to a registry
* Deploy the container image to a Kubernetes cluster

Let's start by examining the first steps to build the code and container:

```yaml
image: atlassian/default-image:2

build-app: &build-app
  - step:
      name: Build and Test
      caches:
        - maven
      script:
        - mvn -B verify --file pom.xml
      after-script:
          # Collect checkstyle results, if any, and convert to Bitbucket Code Insights.
        - pipe: atlassian/checkstyle-report:0.2.0

atlassian-security-scan: &atlassian-security-scan
  - step:
      name: Security Scan
      script:
        # Run a security scan for sensitive data.
        # See more security tools at https://bitbucket.org/product/features/pipelines/integrations?&category=security
        - pipe: atlassian/git-secrets-scan:0.4.3
```

In this example, we perform a few activities to illustrate a representative pipeline.  The first stage is **build-app,** and it's primary purpose is to run a maven build with checkstyle to generate a jar file.  The next stage is entitled **atlassian-security-scan** and we utilize the Atlassian pipe to scan for git secrets.  These two steps are include for convenience and illustrative of a best practice to ensure things are working.  We will not cover their details in the workshop, and we encourage you to review if you are interested.

The next block includes Snyk scan where we use two pipes to scan and push an image.

```
snyk-scan-push-image: &snyk-scan-push-image
  - step:
      name: "Push container image"
      services:
        - docker
      script:
        - docker build -t $IMAGE .
#        - docker tag $IMAGE $IMAGE:${BITBUCKET_COMMIT}
        - docker tag $IMAGE $IMAGE:latest
        - pipe: snyk/snyk-scan:0.5.1
          variables:
            SNYK_TOKEN: $SNYK_TOKEN
            LANGUAGE: "docker"
            IMAGE_NAME: $IMAGE
            TARGET_FILE: "Dockerfile"
            CODE_INSIGHTS_RESULTS: "true"
            SEVERITY_THRESHOLD: "high"
            DONT_BREAK_BUILD: "true"
            MONITOR: "true"
        - pipe: atlassian/aws-ecr-push-image:1.1.3
          variables:
            AWS_ACCESS_KEY_ID: '$AWS_ACCESS_KEY_ID'
            AWS_SECRET_ACCESS_KEY: '$AWS_SECRET_ACCESS_KEY'
            AWS_DEFAULT_REGION: '$AWS_DEFAULT_REGION'
            IMAGE_NAME: $IMAGE
#            TAGS: '${BITBUCKET_COMMIT}'
            TAGS: latest
```

The first part of the script builds the docker image with commands already familiar to docker users.

The next block is the [Snyk Scan](https://bitbucket.org/product/features/pipelines/integrations?p=snyk/snyk-scan) pipe for Atlassian Bitbuckets, with configuration parameters.  The details for configuring the pipe are available online, but you can see the human-readable names help us easily understand what is going on.  Let's take a closer look at a few of these:

1. `SNYK_TOKEN` is being passed into our pipe as a repository variable previously defined in the \[**Bitbucket Configuration**]\() module.
2.  `PROJECT_FOLDER` is the folder in which the project resides and normally defaults to `.`. However, in our example, we have set this to `app/goof` and

    are passing this as an [artifact](https://support.atlassian.com/bitbucket-cloud/docs/use-artifacts-in-steps/) to other steps in our pipeline.
3. `CODE_INSIGHTS_RESULTS` defaults to `false`. However, we will want to create [Code Insight report with Snyk](https://snyk.io/blog/enhanced-security-for-bitbucket-cloud-development/) test results so we have set this to `true`.
4. `SEVERITY_THRESHOLD` reports on issues equal or higher of the provided level. The default is `low` but in our case, we are interested only in `high` so we have defined this variable accordingly.
5. `DONT_BREAK_BUILD` the default is `false` and this is to be expected. Under normal circumstances we would want to break our build if issues our found. However, for the purpose of this workshop we set the value to `true`.

{% hint style="info" %}
You can run **Snyk security scans** on your _pull requests_ and view results in **Code Insights** with the help of a brand new [Snyk Security Connect App on the Atlassian Marketplace](https://marketplace.atlassian.com/apps/1222359/snyk-for-bitbucket-cloud?hosting=cloud\&tab=overview). It's easy to get started and you can install the app with just a few clicks.
{% endhint %}

The last block is for an EKS deployment.  We include the block in the workshop to complete the picture of a deployment to a destination environment.  The pipeline will automatically deploy the container to the named pipeline based on the declarations below.

```
deploy-image: &deploy-image
  - step:
      name: "Push container image"
      script:
      - pipe: atlassian/aws-eks-kubectl-run:2.2.0
        variables:
          CLUSTER_NAME: $AWS_EKS_CLUSTER
          KUBECTL_COMMAND: 'apply'
          RESOURCE_PATH: 'k8s'
```

