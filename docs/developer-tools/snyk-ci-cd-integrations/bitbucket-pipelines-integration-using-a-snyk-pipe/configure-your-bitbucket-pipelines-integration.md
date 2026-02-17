# Configure your Bitbucket Pipelines integration

To enable Snyk to test and monitor your code as an integral part of your CI/CD workflow in Bitbucket, add the Snyk pipe into your `bitbucket-pipelines.yml` (YAML) file located in the root of your repository.&#x20;

This file defines all your build configurations, that is, the pipelines for your CI/CD workflow.

### API & Web scanning

Use Snyk API & Web to run dynamic scans of web apps and APIs from your Bitbucket pipelines. At a high level:

- Store `PROBELY_API_KEY` and `TARGET_ID` as secure repository variables.
- Add a step that installs the Snyk API & Web CLI and starts a scan.

Example (non-blocking scan):

```yaml
pipelines:
  default:
    - step:
        name: Dynamic scan (Snyk API & Web)
        image: python:3.11
        script:
          - pip install probely
          - SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key ${PROBELY_API_KEY})
          - echo "Started scan ${SCAN_ID}"
```

For blocking scans that wait for results and optionally fail on high issues, poll the scan status using `probely scans get`. For scanning internal apps or ephemeral environments, run with a scanning agent; see [Install a Scanning Agent](../../../implementation-and-setup/enterprise-setup/networking-and-connectivity/install-a-scanning-agent.md) and [Scan internal apps with an agent](../../../implementation-and-setup/enterprise-setup/networking-and-connectivity/scan-internal-apps-with-agent.md).
