# config.yml

### CircleCI Configuration Walk-through

{% hint style="info" %}
CircleCI orbs are reusable package of YAML configuration that condenses repeated pieces of config into a single line of code. Learn more about why to use an orb or explore use cases [here](https://circleci.com/orbs/).
{% endhint %}

#### The basics

We have defined our [`version`](https://circleci.com/docs/2.0/configuration-reference/#version) as [Config Reference 2.1](https://circleci.com/docs/reference-2-1/#section=configuration) as well as the [`orbs`](https://circleci.com/docs/2.0/configuration-reference/#orbs-requires-version-21) we intend to leverage for our job. We will be using the following orbs:

*  [`aws-eks`](https://circleci.com/orbs/registry/orb/circleci/aws-eks) for working with Amazon Elastic Container Service for Kubernetes \(Amazon EKS\).
*  [`aws-ecr`](https://circleci.com/orbs/registry/orb/circleci/aws-ecr) to build images and push them to the Amazon Elastic Container Registry.
*  [`kubernetes`](https://circleci.com/orbs/registry/orb/circleci/kubernetes) which is a collection of tools for working with Kubernetes on CircleCI.
*  [`snyk`](https://circleci.com/orbs/registry/orb/snyk/snyk) to find, fix and monitor known vulnerabilities in your app dependencies and docker image.

```yaml
version: 2.1

orbs:
  aws-eks: circleci/aws-eks@0.2.7
  aws-ecr: circleci/aws-ecr@6.8.2
  kubernetes: circleci/kubernetes@0.11.0
  snyk: snyk/snyk@0.0.10

defaults: &defaults
  docker:
    - image: circleci/node:9.11.2
  working_directory: ~/repo
```

In the interest of efficiency, we have also defined a few defaults for our project which we will later call from our jobs.

{% hint style="info" %}
Did you know that you can search the [CircleCI registry](https://circleci.com/orbs/registry/) for several _**certified**_ and _**partner**_ orbs to solve most use cases?
{% endhint %}

#### Workflows

Our example `config.yml` file also organizes and orchestrates our defined jobs using [`workflows`](https://circleci.com/docs/2.0/configuration-reference/#workflows) .

```yaml
workflows:
  build_and_deploy:
    jobs:
      - test_app
      - scan_app:
          requires:
            - test_app
      - build_and_scan_image:
          requires:
            - scan_app
      - build_and_push_image:
          requires:
            - build_and_scan_image
      - deploy_app:
          cluster-name: ${CIRCLE_PROJECT_REPONAME}
          aws-region: ${AWS_REGION_ENV_VAR_NAME}
          docker-image-name: "$AWS_ECR_ACCOUNT_URL_ENV_VAR_NAME/${CIRCLE_PROJECT_REPONAME}:${CIRCLE_SHA1}"
          version-info: "${CIRCLE_SHA1}"
          requires:
            - build_and_push_image
```





