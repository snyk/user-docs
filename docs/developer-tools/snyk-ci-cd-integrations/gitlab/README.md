# GitLab CI/CD: Snyk API & Web scanning

Use Snyk API & Web to run dynamic scans of web apps and APIs in your GitLab pipelines.

## Prerequisites
- Create a target in Snyk API & Web and note the target ID.
- In GitLab, add CI/CD variables for `PROBELY_API_KEY` and `TARGET_ID`.
- For internal apps or ephemeral environments, run with a scanning agent; see:
  - [Install a Scanning Agent](../../../implementation-and-setup/enterprise-setup/networking-and-connectivity/install-a-scanning-agent.md)
  - [Scan internal apps with an agent](../../../implementation-and-setup/enterprise-setup/networking-and-connectivity/scan-internal-apps-with-agent.md)

## Example: Non-blocking scan job
This job starts a scan and exits without waiting for results.

```yaml
stages:
  - dast

dast_scan:
  stage: dast
  image: python:3.11
  script:
    - pip install probely
    - |
      SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key ${PROBELY_API_KEY})
      echo "Started scan ${SCAN_ID}"
```

## Example: Blocking scan job (wait for results)
This variation waits for the scan to finish and fails the job if high-severity findings are present.

```yaml
stages:
  - dast

dast_scan_blocking:
  stage: dast
  image: python:3.11
  script:
    - pip install probely
    - |
      SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key ${PROBELY_API_KEY})
      echo "Started scan ${SCAN_ID}"
      # Wait until scan finishes
      while true; do
        STATUS=$(probely scans get ${SCAN_ID} --api-key ${PROBELY_API_KEY} -o JSON | jq -r '.status')
        echo "Status: ${STATUS}"
        if [ "${STATUS}" = "started" ] || [ "${STATUS}" = "queued" ]; then
          sleep 30
        else
          break
        fi
      done
      # Optional: fail on high issues
      HIGHS=$(probely scans get ${SCAN_ID} --api-key ${PROBELY_API_KEY} -o JSON | jq -r '.highs')
      echo "High issues: ${HIGHS}"
      if [ "${HIGHS}" -gt 0 ]; then
        echo "Failing due to high issues"
        exit 1
      fi
```

## Tips
- Use masked/protected CI/CD variables for secrets.
- Coordinate scan timing with your deploy stage (for example, scan a staging URL after deployment).
- For ephemeral environments, ensure DNS/hosts or agent routing makes the app reachable to Snyk API & Web.
