# Push image to registry

Next, we call the `circleci/aws-ecr` orb once more, but this time we are invoking the [`build-and-push-image`](https://circleci.com/orbs/registry/orb/circleci/aws-ecr#commands-build-and-push-image) command to authenticate to our AWS account and push our image.

```yaml
  build_and_push_image:
    <<: *defaults
    steps:
      - attach_workspace:
          at: ~/repo
      - setup_remote_docker
      - aws-ecr/build-and-push-image:
          account-url: AWS_ECR_ACCOUNT_URL_ENV_VAR_NAME
          aws-access-key-id: ACCESS_KEY_ID_ENV_VAR_NAME
          aws-secret-access-key: SECRET_ACCESS_KEY_ENV_VAR_NAME
          region: AWS_REGION_ENV_VAR_NAME
          repo: ${CIRCLE_PROJECT_REPONAME}
          create-repo: true
          checkout: true
          dockerfile: Dockerfile
          path: ./submodules/goof/
          tag: ${CIRCLE_SHA1}
```

{% hint style="info" %}
By setting `create-repo` to `true` the orb will create the repo if it does not exist, but subsequent runs of the job will simply update the image to the existing repo.
{% endhint %}

