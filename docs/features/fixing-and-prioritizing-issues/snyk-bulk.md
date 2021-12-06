---
description: https://github.com/snyk-tech-services/snyk-bulk
---

# Snyk-Bulk

Snyk-Bulk is a tool that allows users to simulate a discover-oriented vulnerability scan outside of their CI/CD pipeline. Snyk-Bulk reproduces the snyk test/monitor --all-projects on a per-language basis.



### Use Cases:

#### Scanning for out-of-band repos

* With Snyk-Bulk, security teams who may not have direct access to developer build pipelines can still scan repos without interfering with the build process
* Snyk-Bulk will automatically build the dependencies when testing, allowing security teams to simulate the snyk-test/monitor command, providing actionable feedback to developer teams.

#### Testing for Monorepos

* Using snyk-bulk, monorepos can be scanned more quickly by scanning projects at a per-language basis, cutting down the time required to scan these large projects.&#x20;

### Example of using Snyk-Bulk to scan for a Python3 project:

Note: This example can also be found in the Snyk-Bulk project in github. For further details/instructions, please refer to the Snyk-Bulk github project linked above.

#### Goal:&#x20;

Use Snyk-Bulk to scan and test for a python project. We will build a python project using a Dockerfile located within the Snyk-Bulk project. We will then scan this project for vulnerabilites and verify within the Snyk UI.&#x20;

#### Pre-requisites:&#x20;

* Snyk Token (found [here](https://app.snyk.io/account))
* Docker

#### Steps:&#x20;

1.  Clone the github project for snyk-bulk

    `git clone` [`https://github.com/snyk-tech-services/snyk-bulk`](https://github.com/snyk-tech-services/snyk-bulk)``
2.  Access the snyk-bulk folder

    `cd snyk-bulk`
3.  Build the Docker container for Python using the Dockerfile in the Snyk-Bulk repository

    `docker build -t snyk-bulk:python3 -f Dockerfile-python .`
4.  `Set the SNYK_TOKEN variable`

    `export SNYK_TOKEN=<Snyk API token>`
5.  Run the following command to run snyk-test using snyk-bulk against the container that you built in step 3

    `docker run -it --env SNYK_TOKEN -v $(PWD):/test snyk-bulk:python3 --test --target /test`
6. Observe the results of the scan both in the CLI and in your Snyk Organization within the Snyk UI.&#x20;

