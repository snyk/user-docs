# Configure your CircleCI integration

After a user adds a project to CircleCI, and adds the [Snyk Orb](https://circleci.com/developer/orbs/orb/snyk/snyk) to the configuration file, then when a build runs, the Snyk Orb is used as well.

## Scan

1. Scans app dependencies or container images for vulnerabilities or licensing issues, and lists them.
2. If Snyk finds vulnerabilities, it does one of the following (based on configuration):
   * Fails the build
   * Lets the build complete

## Monitor

Optionally, if the build completes successfully and **MONITOR** is set to True in the Snyk step, then Snyk saves a snapshot of the project dependencies from the Snyk app, where you can view the dependency tree displaying all of the issues, and you can receive alerts for new issues found in the existing app version.

## Protect (optional)

(For Node.js projects only) Optionally, set **PROTECT** to **True** and if a `.snyk` policy file exists and you have the [`@snyk/protect`](https://www.npmjs.com/package/@snyk/protect) package in your dependencies, Snyk applies patches specified in the `.snyk` file.

## Prerequisites

1. Create a Snyk account and retrieve the Snyk API token from your **Account settings**.
2. Import the relevant repo into CircleCI.
3. Go to `Settings -> Security -> Orb security settings` and make sure you allow to opt-in to third party Orbs.
4. Make sure your configuration (`config.yml`) file follows version 2.1.
5. Add the required variables to CircleCI (e.g. Snyk API token as `API_TOKEN`)
