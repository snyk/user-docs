# Container image scanning

Next, we will call the [`circleci/aws-ecr`](https://circleci.com/orbs/registry/orb/circleci/aws-ecr) orb to build our image from the `Dockerfile` in our project repository and we will call [`snyk/scan`](https://circleci.com/orbs/registry/orb/snyk/snyk#commands-scan) command to find vulnerabilities in our base image. Again, we have made a few choices here such as setting `fail-on-issues` to `false` and setting our `severity-threshold` to `high`.  

```yaml
  build_and_scan_image:
    <<: *defaults
    steps:
      - attach_workspace:
          at: ~/repo
      - setup_remote_docker
      - aws-ecr/build-image:
          account-url: AWS_ECR_ACCOUNT_URL_ENV_VAR_NAME
          dockerfile: Dockerfile
          path: ./submodules/goof/
          repo: ${CIRCLE_PROJECT_REPONAME}
          tag: ${CIRCLE_SHA1}
      - snyk/scan:
          docker-image-name: '$AWS_ECR_ACCOUNT_URL_ENV_VAR_NAME/${CIRCLE_PROJECT_REPONAME}:${CIRCLE_SHA1}'
          fail-on-issues: false
          monitor-on-build: true
          project: '${CIRCLE_PROJECT_REPONAME}/${CIRCLE_BRANCH}-container'
          severity-threshold: high
          target-file: ./submodules/goof/Dockerfile
          token-variable: SNYK_TOKEN
```

