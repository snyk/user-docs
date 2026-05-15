# CircleCI migration

This page explains how to transition away from affected jobs.

## Updating snyk orb and using `iac test`

Customers using the `scan-iac` job will need to switch to using `snyk/scan` with the `iac test` command. For an example, see the [IaC scanning examples](https://github.com/snyk/snyk-orb/blob/v2.0.0/src/examples/quickstart-iac-scanning.yml) in the snyk-orb repository.

It is important to update the version of the snyk orb used to the latest version by updating the circleci config files. Currently, the latest Snyk orb is `snyk/snyk@2.1.0` .

## Use the Snyk orb to only install the Snyk CLI, and then run the Snyk CLI commands in your own steps <a href="#use-the-snyk-orb-to-only-install-the-snyk-cli-and-then-run-the-snyk-cli-commands-in-your-own-steps" id="use-the-snyk-orb-to-only-install-the-snyk-cli-and-then-run-the-snyk-cli-commands-in-your-own-steps"></a>

Instead of relying on predefined jobs, customers can use the Snyk orb to install the Snyk CLI and then run commands as their own steps. See this [install example](https://github.com/snyk/snyk-orb/blob/v2.0.0/src/examples/only-install.yml) in the snyk-orb repository.

In the case of replacing `scan-iac` job, an example config could look like the following:

```yaml
version: 2.1
orbs:
  node: circleci/node@5
  snyk: snyk/snyk@2.1.0
jobs:
  snyk_scan:
    docker:
      - image: cimg/node:lts
    steps:
      - checkout
      - run: npm ci
      - snyk/install
      - run:
          command: snyk iac test
          name: Run iac test 
```

## Direct CLI installation without using the Snyk orb <a href="#direct-cli-installation-without-using-the-snyk-orb" id="direct-cli-installation-without-using-the-snyk-orb"></a>

Alternatively, customers who prefer not to rely on the Snyk CircleCI orb or wish to have complete control over their pipelines can opt to install the Snyk CLI directly. An example follows:

```yaml
version: 2.1
jobs:
  snyk_scan:
    docker:
      - image: cimg/node:lts
    steps:
      - checkout
      - run: npm ci
      - run:
          name: Download Snyk CLI
          command: |
            curl -Lo snyk-linux https://downloads.snyk.io/cli/stable/snyk-linux
      - run:
          name: Download Snyk CLI SHA256 Checksum
          command: |
            curl -Lo snyk-linux.sha256 https://downloads.snyk.io/cli/stable/snyk-linux.sha256
      - run:
          name: Verify SHA256 Checksum
          command: |
            sha256sum -c snyk-linux.sha256
      - run:
          name: Install Snyk CLI
          command: |
            chmod +x snyk-linux
            ./snyk-linux --version
      - run:
          name: Run Snyk iac test
          command: |
            ./snyk-linux iac test
workflows:
  version: 2
  build_and_scan:
    jobs:
      - snyk_scan
```
