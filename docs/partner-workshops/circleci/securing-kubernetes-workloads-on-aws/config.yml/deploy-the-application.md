# Deploy the application

Now, we are ready to push our application. Here we are calling the `aws-eks/python3` [executor](https://circleci.com/docs/2.0/configuration-reference/#executors-requires-version-21) which will define the environment in which the steps of our job will be run. 

```yaml
  deploy_app:
    executor: aws-eks/python3
    parameters:
      cluster-name:
        type: string
      docker-image-name:
        type: string
      version-info:
        type: string
      aws-region:
        type: string
        default: ""
    steps:
      - checkout
      - run:
          name: Create deployment manifest
          command: |
            BUILD_DATE=$(date '+%Y%m%d%H%M%S')
            cat deployment/goof-deployment-template.yaml |\
               sed "s|DOCKER_IMAGE_NAME|<< parameters.docker-image-name >>|g"\
                 > deployment/goof-deployment.yaml
      - aws-eks/update-kubeconfig-with-authenticator:
          cluster-name: << parameters.cluster-name >>
          install-kubectl: true
          aws-region: << parameters.aws-region >>
      - kubernetes/create-or-update-resource:
          resource-file-path: "deployment/goof-deployment.yaml"
          get-rollout-status: true
          resource-name: deployment/goof
      - kubernetes/create-or-update-resource:
          resource-file-path: "deployment/goof-service.yaml"
```

In this job, we are passing a few [parameters](https://circleci.com/docs/2.0/configuration-reference/#parameters-requires-version-21) such as the Kubernetes `cluster-name` and the `docker-image-name` which references the image we recently created. We are using these values to authenticate to our Kubernetes cluster and also to replace the value of our image in our Kubernetes manifest with the [Amazon ECR repository URL](https://docs.aws.amazon.com/AmazonECR/latest/userguide/docker-pull-ecr-image.html) and tag. Then,  will invoke the `kubernetes/create-or-update-resources` command to apply our Kubernetes manifests for the deployment and service.

