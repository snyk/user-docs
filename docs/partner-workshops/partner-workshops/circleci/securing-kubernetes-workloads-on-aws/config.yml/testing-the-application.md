# Testing the application

Our project includes a sample application from another project. For this we have leveraged [`git submodule`](https://git-scm.com/book/en/v2/Git-Tools-Submodules) to keep a Git repository as a subdirectory of another Git repository. In this example, we are working with [Snyk's vulnerable demo app](https://github.com/snyk/goof) which may be found in our repo under the `./submodules` directory. 

The next section of our `config.yml` defines a few `jobs` starting with `test_app` which will run some commands in our build environment. Here, we are importing the parameters defined earlier under `defaults` such as the Docker image for our build: `circleci/node:9.11.2`. We will also invoke `git submodule` and `npm install` the application. Lastly, because we will want this artifact downstream, we will call `persist_to_workspace` to reference it later on. 

```yaml
jobs:
  test_app:
    <<: *defaults
    steps:
      - checkout
      - run:
          name: "pull submodules"
          command: |
            git submodule init
            git submodule update --recursive
      - run:
          name: "run test"
          command: |
            cd submodules/goof
            npm install
      - persist_to_workspace:
          root: .
          paths:
            - .
```

{% hint style="info" %}
You can read more about basic concepts to help you understand how CircleCI manages your CICD pipelines [here](https://circleci.com/docs/2.0/concepts/). Also checkout the blog post [Persisting Data in Workflows: When to Use Caching, Artifacts, and Workspaces](https://circleci.com/blog/persisting-data-in-workflows-when-to-use-caching-artifacts-and-workspaces/) to learn more about how to move data into and out of jobs.
{% endhint %}

