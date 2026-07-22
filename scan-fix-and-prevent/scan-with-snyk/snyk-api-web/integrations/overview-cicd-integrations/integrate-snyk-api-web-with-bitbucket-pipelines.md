---
description: How to integrate Snyk API and Web with Bitbucket Pipelines
nav_context: classic
---

# Integrate with Bitbucket Pipelines

This guide provides step-by-step instructions for integrating Snyk API & Web into your Bitbucket Pipelines.

## Overview

This guide focuses on using the Snyk CLI to run scans. The following examples cover a complete end-to-end journey, from configuring your targets in the Snyk UI to running different scan scenarios in your pipeline.

## Prerequisites

Configure your scan targets and credentials in the Snyk application before you begin.

### Create a target in Snyk

In the Snyk app, navigate to the **Targets** menu and click **Add**. Fill out the form and click **Add** to create the new target.

{% hint style="info" %}
During this process, Snyk checks connectivity. If your target is internal or not yet deployed, you can bypass any warnings and add the target. For more details, visit How to add a Target.
{% endhint %}

Before you configure the integration in Bitbucket Pipelines, retrieve the unique target ID from Snyk.

1. In your Snyk dashboard, select **Targets**.
2. From the target list, select the target you want to integrate.
3. In your browser's address bar, copy the target ID. This is the string of characters immediately following /target/ in the URL.

{% hint style="info" %}
After you create a target, you must verify the target's domain. Otherwise, your scans are limited to lightning scans. To learn more, visit the importance of domain ownership verification.
{% endhint %}

### Create a Snyk API key

You need an API key with permissions to start a scan on your target. For instructions, visit How to generate an API key.

## Step 1: Add your API key and target ID to Bitbucket Pipelines

To run a scan, your pipeline must authenticate with Snyk and know which target to scan. Configure your Snyk API key and target ID as secure repository variables in your Bitbucket Pipelines project.

1. From your Bitbucket Pipelines project side menu, navigate to **Repository settings > Repository variables**.
2. Click **Add** and create an entry for your Snyk **API Key** (for example, **PROBELY\_API\_KEY**).
3. Click **Add** again to create a second entry for your **target ID** (for example, **TARGET\_ID**).

{% hint style="warning" %}
For enhanced security, always store sensitive values as Bitbucket CI/CD variables. Do not store variables directly in your bitbucket-pipelines.yml file, because they are saved in plain text and visible to anyone who can view the file.
{% endhint %}

## Step 2: Configure your pipeline

Create a `bitbucket-pipelines.yml` file at the root of your repository and add one of the following code examples based on your use case. You can also find all of them in the Snyk API & Web CI/CD examples repository on GitHub.

### Important note on these examples

The following YAML configurations are scanning steps to incorporate into your existing `bitbucket-pipelines.yml` file.

For example, your pipeline might already have steps to build your code, deploy to a QA environment, and run automated tests. You can add the Snyk scan as another step at any point that makes sense for your workflow, such as after you deploy to QA or staging.

### Run a scan on a target in non-blocking mode

This is the simplest configuration. It uses the Snyk CLI to run a scan on a remote target and does not wait for the scan to finish, so the pipeline completes quickly.

```yaml
# bitbucket-remote-app-non-blocking-mode.yaml
image: ubuntu:latest

pipelines:
  branches:
    main:
      - step:
          name: Install Snyk API & Web CLI and start a Scan
          caches:
            - pip
          script:
            - apt-get update && apt-get install -y python3-pip python3-venv jq
            - python3 -m venv venv
            - source ./venv/bin/activate
            - pip install probely # install Snyk API & Web CLI
            - probely targets get --api-key $PROBELY_API_KEY
              
            - |
              echo "Starting Scan..."
              for i in {1..20}; do
                echo "-----------------------------------"
                SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key $PROBELY_API_KEY)
                echo ${SCAN_ID}
                if [ -z "${SCAN_ID}" ]; then
                  echo "Scan didn't start... Retry start-scan"
                else
                  echo "Scan started with SCAN ID: ${SCAN_ID}"
                  break
                fi
                sleep 5
              done
              if [ -z "$SCAN_ID" ]; then
                echo "No Scan ID, aborting..."
                exit 1
              fi
          services:
            - docker
```

### Run a scan on a target in blocking mode

This configuration starts a scan and then polls the results, failing the pipeline if any high-severity vulnerabilities are found.

```yaml
# bitbucket-remote-app-blocking-mode.yaml
image: ubuntu:latest

pipelines:
  branches:
    main:
      - step:
          name: Install Snyk API & Web CLI, Scan and wait for scan to complete
          caches:
            - pip
          script:
            - apt-get update && apt-get install -y python3-pip python3-venv jq
            - python3 -m venv venv
            - source ./venv/bin/activate
            - pip install probely # install Snyk API & Web CLI
            - probely targets get --api-key $PROBELY_API_KEY
              
            - |
              echo "Starting Scan..."
              for i in {1..20}; do
                echo "-----------------------------------"
                SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key $PROBELY_API_KEY)
                echo ${SCAN_ID}
                if [ -z "${SCAN_ID}" ]; then
                  echo "Scan didn't start... Retry start-scan"
                else
                  echo "Scan started with SCAN ID: ${SCAN_ID}"
                  break
                fi
                sleep 5
              done
              if [ -z "$SCAN_ID" ]; then
                echo "No Scan ID, aborting..."
                exit 1
              fi
              
            - |
              echo "Waiting for scan to complete..."
              while true; do
                echo "-----------------------------------"
                SCAN_OUTPUT=$(probely scans get ${SCAN_ID} --api-key $PROBELY_API_KEY | tail -1)
                echo ${SCAN_OUTPUT}
                echo "-----------------------------------"
                SCAN_STATUS=$(probely scans get ${SCAN_ID} --api-key $PROBELY_API_KEY -o JSON | jq -r '.status')
                if [ "$SCAN_STATUS" == "started" ] || [ "$SCAN_STATUS" == "queued" ]; then
                  echo "Scan is running or queued!"
                else
                  echo "Scan is completed."
                  break
                fi
                sleep 30
              done
              
            # Optional logic - abort the pipeline if there are any HIGH risk vulnerabilities.
            - HIGH_VULNS=$(probely scans get $SCAN_ID --api-key $PROBELY_API_KEY -o JSON | jq -r '.highs')
            - echo "HIGH risk vulnerabilities ${HIGH_VULNS}"
            - |
              if [[ "$HIGH_VULNS" -gt 0 ]]; then
                echo "Scan has High risk vulnerabilities... aborting"
                exit 1
              else
                echo "Scan doesn't have High risk vulnerabilities"
              fi
          services:
            - docker
```

### Run a scan on an ephemeral (dynamic) target in blocking mode

This is a more advanced configuration for building, deploying, and scanning an application in a temporary environment. The pipeline creates the environment for a specific purpose and then automatically destroys it during the run.

Ephemeral environments require an agent token, target hostname, and target URL. Store them as variables for better security.

You also need to create a scanning agent in Snyk and configure your target to use it. This process requires the `scanning-agent/farcasterd-linux-amd64-0.4.3` file. For detailed instructions, visit How to install a Scanning Agent and How to scan internal applications.

{% hint style="info" %}
This code example uses Docker to create ephemeral environments. However, you can use any other solution to create your environment.
{% endhint %}

```yaml
# bitbucket-ephemeral-app-blocking-mode.yaml
pipelines:
  default:
    - step:
        name: Build and Test App
        image: python:3.9-slim-buster  # Use a Python image (Debian-based)
        services:
          - docker
        script:
          - apt-get update && apt-get install -y curl jq psmisc coreutils iputils-ping
          - pip3 install probely

          - export DOCKER_HOST=tcp://localhost:2375
          - docker network create custom-network
          - docker build -t test-app .
          - docker run --name test-app --hostname custom-web-app --network custom-network -p 8080:8080 -d test-app
          - CONTAINER_IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' test-app)
          - echo "${CONTAINER_IP} ${TARGET_HOSTNAME} ${TARGET_HOSTNAME}." | tee -a /etc/hosts
          - cat /etc/hosts # confirm host is on /etc/hosts

          # Wait for the app to start
          - |
            for i in {1..10}; do
              if curl -s ${TARGET_URL} > /dev/null; then
                echo "App is up!";
                break;
              fi
              echo "Waiting for the app to be ready...";
              sleep 2;
            done

          # Test application with curl
          - RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" ${TARGET_URL})
          - |
            if [[ "$RESPONSE" -ne 200 ]]; then
              echo "App test failed with HTTP status $RESPONSE";
              docker stop test-app
              docker rm test-app
              docker network rm custom-network
              exit 1;
            fi
          - curl -s -i ${TARGET_URL}
          - echo "App test passed with HTTP status $RESPONSE";

          # Run userspace agent
          - chmod +x scanning-agent/farcasterd-linux-amd64-0.4.3 
          - ./scanning-agent/farcasterd-linux-amd64-0.4.3 --token $AGENT_TOKEN &

          # Wait for the agent to start
          - sleep 40

          - probely targets get --api-key $PROBELY_API_KEY

          # Start Scan
          - |
            for i in {1..20}; do
              echo "-----------------------------------"
              SCAN_ID=$(probely targets start-scan $TARGET_ID -o IDS_ONLY --api-key $PROBELY_API_KEY)
              echo ${SCAN_ID}
              if [[ -z "$SCAN_ID" ]]; then
                echo "Scan didn't start... Retry start-scan"
              else
                echo "Scan started with SCAN ID ${SCAN_ID}";
                break;
              fi
              sleep 5
            done
            if [[ -z "$SCAN_ID" ]]; then
              echo "No Scan ID, aborting..."
              docker stop test-app
              docker rm test-app
              docker network rm custom-network
              killall farcasterd-linux-amd64-0.4.3
              exit 1
            fi

          # Wait for scan to end
          - |
            while true; do
              echo "-----------------------------------"
              SCAN_OUTPUT=$(probely scans get $SCAN_ID --api-key $PROBELY_API_KEY | tail -1)
              echo ${SCAN_OUTPUT}
              echo "-----------------------------------"
              SCAN_STATUS=$(probely scans get $SCAN_ID --api-key $PROBELY_API_KEY -o JSON | jq -r '.status')
              if [[ "$SCAN_STATUS" == "started" ]] || [[ "$SCAN_STATUS" == "queued" ]]; then
                echo "Scan is running or queued!";
              else
                echo "Scan is not running... finishing"
                break;
              fi
              sleep 30;
            done

          # Optional logic - abort the pipeline if there are any HIGH risk vulnerabilities.
          - HIGH_VULNS=$(probely scans get $SCAN_ID --api-key $PROBELY_API_KEY -o JSON | jq -r '.highs')
          - echo "HIGH vulnerabilities ${HIGH_VULNS}"
          - |
            if [[ "$HIGH_VULNS" -gt 0 ]]; then
              echo "Scan has High risk vulnerabilities... aborting"
              docker stop test-app
              docker rm test-app
              docker network rm custom-network
              killall farcasterd-linux-amd64-0.4.3
              exit 1
            else
              echo "Scan doesn't have High risk vulnerabilities"
            fi

          # Clean up
          - docker stop test-app
          - docker rm test-app
          - docker network rm custom-network
          - killall farcasterd-linux-amd64-0.4.3
```

### Advanced scenario: dynamic target creation

This example demonstrates how to use the Snyk CLI to dynamically check for, create, and then scan a target that does not already exist in the Snyk UI.

```yaml
# bitbucket-create-remote-app-blocking-mode.yaml
image: ubuntu:latest

pipelines:
  branches:
    main:
      - step:
          name: Install Snyk API & Web CLI, Scan and wait for scan to complete
          caches:
            - pip
          script:
            - apt-get update && apt-get install -y python3-pip python3-venv jq
            - python3 -m venv venv
            - source ./venv/bin/activate
            - pip install probely # install Snyk API & Web CLI
            - probely targets get --api-key $PROBELY_API_KEY

            - TARGET_REPO_LABEL_ID=$(probely target-labels get --f-search ${BITBUCKET_REPO_SLUG} -o IDS_ONLY --api-key ${PROBELY_API_KEY})
            - echo "TARGET LABEL => ${TARGET_REPO_LABEL_ID}"
            - |
              echo "Create target if does not exist"
              if [[ -z "${TARGET_REPO_LABEL_ID}" ]]; then
                echo -e "labels:\n  - name: ${BITBUCKET_REPO_SLUG}" > /tmp/label-${BITBUCKET_REPO_SLUG}.yaml
                echo "There is no label for this repo slug and no target... create target with label"
                TARGET_ID=$(probely targets add ${TARGET_URL} --target-name ${TARGET_URL} --target-type WEB -f /tmp/label-${BITBUCKET_REPO_SLUG}.yaml -o IDS_ONLY --api-key ${PROBELY_API_KEY})
              else
                TARGET_ID=$(probely targets get --f-search ${TARGET_URL} --f-label ${TARGET_REPO_LABEL_ID} -o IDS_ONLY --api-key ${PROBELY_API_KEY})
                if [[ -z "${TARGET_ID}" ]]; then
                  echo "There is no target ID, creating a new target"
                  echo -e "labels:\n  - name: ${BITBUCKET_REPO_SLUG}" > /tmp/label-${BITBUCKET_REPO_SLUG}.yaml
                  TARGET_ID=$(probely targets add ${TARGET_URL} --target-name ${TARGET_URL} --target-type WEB -f /tmp/label-${BITBUCKET_REPO_SLUG}.yaml -o IDS_ONLY --api-key ${PROBELY_API_KEY})
                fi
              fi

            - echo "Check TARGET_ID ${TARGET_ID}"
            - |
              if [[ -z "${TARGET_ID}" ]]; then
                echo "No TARGET_ID... Aborting..."
                exit 1
              fi
              
            - |
              echo "Starting Scan..."
              for i in {1..20}; do
                echo "-----------------------------------"
                SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key $PROBELY_API_KEY)
                echo ${SCAN_ID}
                if [ -z "${SCAN_ID}" ]; then
                  echo "Scan didn't start... Retry start-scan"
                else
                  echo "Scan started with SCAN ID: ${SCAN_ID}"
                  break
                fi
                sleep 5
              done
              if [ -z "$SCAN_ID" ]; then
                echo "No Scan ID, aborting..."
                exit 1
              fi
              
            - |
              echo "Waiting for scan to complete..."
              while true; do
                echo "-----------------------------------"
                SCAN_OUTPUT=$(probely scans get ${SCAN_ID} --api-key $PROBELY_API_KEY | tail -1)
                echo ${SCAN_OUTPUT}
                echo "-----------------------------------"
                SCAN_STATUS=$(probely scans get ${SCAN_ID} --api-key $PROBELY_API_KEY -o JSON | jq -r '.status')
                if [ "$SCAN_STATUS" == "started" ] || [ "$SCAN_STATUS" == "queued" ]; then
                  echo "Scan is running or queued!"
                else
                  echo "Scan is completed."
                  break
                fi
                sleep 30
              done
              
            # Optional logic - abort the pipeline if there are any HIGH risk vulnerabilities.
            - HIGH_VULNS=$(probely scans get $SCAN_ID --api-key $PROBELY_API_KEY -o JSON | jq -r '.highs')
            - echo "HIGH risk vulnerabilities ${HIGH_VULNS}"
            - |
              if [[ "$HIGH_VULNS" -gt 0 ]]; then
                echo "Scan has High risk vulnerabilities... aborting"
                exit 1
              else
                echo "Scan doesn't have High risk vulnerabilities"
              fi
          services:
            - docker
```

## Step 3: Run the pipeline and view the results

After you commit your `bitbucket-pipelines.yml` file, you can run the pipeline in Bitbucket Pipelines to test the integration.

After the scan completes, you can view the detailed findings in your Snyk dashboard.
