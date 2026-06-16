# Integrate with GitHub Actions

This guide provides step-by-step instructions for integrating Snyk API & Web into your GitHub Actions pipelines.

## Overview

To foster automation between systems, you can trigger target scans directly from your GitHub repository on events like a push to your main branch.

Two integration methods are offered:

* **Snyk API & Web Action:** A simple way to quickly add a scan to your workflow.
* **Snyk API & Web CLI:** A more flexible and powerful method that gives you full control over complex scenarios like blocking builds and scanning ephemeral applications.

## Prerequisites

Before you begin, you must configure your scan targets and credentials in the Snyk API & Web application.

### Create a target in Snyk API & Web

In the Snyk API & Web app, go to the Targets menu and click **Add**. Fill out the form and click **Add** to create the new target.

{% hint style="info" %}
During this process, connectivity is checked. If your target is internal or not yet deployed, you can bypass any warnings and add the target regardless. For more details, visit How to add a Target.
{% endhint %}

Before configuring the integration in GitHub Actions, make sure to retrieve the unique target ID from Snyk API & Web.

1. In your Snyk API & Web dashboard, select **Targets**.
2. From the target list, select the target you want to integrate.
3. In your browser's address bar, copy the target ID. This is the string of characters immediately following /target/ in the URL.

{% hint style="info" %}
After creating a target, it is mandatory to verify your target's domain. Otherwise, your scans are only limited to lightning scans. To learn more, see the importance of domain ownership verification.
{% endhint %}

### Create a Snyk API & Web API key

You need an API key with permissions to start a scan on your target. For instructions, visit How to generate an API key.

## Step 1: Add your API key and target ID to GitHub

To allow GitHub Actions to communicate with Snyk API & Web, you must store your credentials as secure repository secrets.

1. In your GitHub repository, navigate to the **Settings** tab.
2. In the side menu, expand **Secrets and variables** and click **Actions**.
3. Click **New repository secret**.
4. Create a secret named **PROBELY\_API\_KEY** and paste the API Key you generated as the value.

## Step 2: Configure your pipeline

### Option 1: Use the Snyk API & Web GitHub Action

This integration method uses the official GitHub Action.

```yaml
# Sample workflow for scanning a target with Snyk API & Web
on:
  # Runs on pushes targeting the default branch
  push:
    branches: ["main"]
  # Allows you to run this workflow manually from the Actions tab
  workflow_dispatch:
jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - name: Scan with Snyk API & Web
        id: snyk-scan
        uses: Probely/probely-github-action@main
        with:
          api-key: ${{ secrets.PROBELY_API_KEY }}
          target-id: "<TARGET_ID>"
          region: "eu"
```

Depending on the region where your Snyk API & Web instance is located, you need to specify a region value. The options are: `eu`, `us`, `au`.

Remember to replace `<TARGET_ID>` with the target ID you copied in the Prerequisites step.

### Option 2: Use the Snyk API & Web CLI

For more control over the pipeline, such as blocking builds or scanning ephemeral applications, you can use the Snyk API & Web CLI directly.

Create a `probely.yml` file at the root of your repository, under the `.github/workflows/` directory, and add one of the following code examples based on your use case. You can also find all of them in the Snyk API & Web CI/CD examples repository on GitHub.

#### Important note on these examples

The YAML configurations below are scanning steps designed to be incorporated into your existing `probely.yml` file.

For example, your pipeline might already have steps to build your code, deploy to a QA environment, and run automated tests. You can add the Snyk API & Web scan as another step at any point that makes sense for your workflow, such as after you deploy to QA or staging.

#### Run a scan on a target in non-blocking mode

This workflow installs the CLI and starts a scan, but does not wait for the results, allowing your pipeline to complete quickly.

```yaml
# github-remote-app-non-blocking-mode.yaml
name: CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      # Step 1: Install Snyk API & Web CLI
      - name: Install Snyk API & Web CLI
        run: |
          # Install Snyk API & Web CLI
          pip install probely
          # Test probely GET TARGETS
          probely targets get --api-key ${{ secrets.PROBELY_API_KEY }}
        
      # Step 2: Start Scan
      - name: Start Scan
        run: |
          for i in {1..20}; do
            echo "-----------------------------------"
            SCAN_ID=$(probely targets start-scan ${{ vars.TARGET_ID }} -o IDS_ONLY --api-key ${{ secrets.PROBELY_API_KEY }})
            echo ${SCAN_ID}
            if [ -f ${SCAN_ID} ]; then
              echo "Scan didn't start... Retry start-scan"
            else
              echo "Scan started with SCAN ID: ${SCAN_ID}";
              echo "SCAN_ID=${SCAN_ID}" >> $GITHUB_ENV
              break;
            fi
            sleep 5
          done
          if [ -f $SCAN_ID ]; then
            echo "No Scan ID, aborting..."
            exit 1
          fi
```

#### Run a scan on a target in blocking mode

This workflow starts a scan and then polls the results, failing the build if any high-severity vulnerabilities are found.

```yaml
# github-remote-app-blocking-mode.yaml
name: CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      # Step 1: Install Snyk API & Web CLI
      - name: Install Snyk API & Web CLI
        run: |
          # Install Snyk API & Web CLI
          pip install probely
          # Test probely GET TARGETS
          probely targets get --api-key ${{ secrets.PROBELY_API_KEY }}
        
      # Step 2: Start Scan
      - name: Start Scan
        run: |
          for i in {1..20}; do
            echo "-----------------------------------"
            SCAN_ID=$(probely targets start-scan ${{ vars.TARGET_ID }} -o IDS_ONLY --api-key ${{ secrets.PROBELY_API_KEY }})
            echo ${SCAN_ID}
            if [ -f ${SCAN_ID} ]; then
              echo "Scan didn't start... Retry start-scan"
            else
              echo "Scan started with SCAN ID: ${SCAN_ID}";
              echo "SCAN_ID=${SCAN_ID}" >> $GITHUB_ENV
              break;
            fi
            sleep 5
          done
          if [ -f $SCAN_ID ]; then
            echo "No Scan ID, aborting..."
            exit 1
          fi

      # Step 3: Wait for scan to end
      - name: Wait for scan to end
        run: |
          # Wait until scan ends
          while true; do
            echo "-----------------------------------"
            SCAN_OUTPUT=$(probely scans get ${SCAN_ID} --api-key ${{ secrets.PROBELY_API_KEY }} | tail -1)
            echo ${SCAN_OUTPUT}
            echo "-----------------------------------"
            SCAN_STATUS=$(probely scans get ${SCAN_ID} --api-key ${{ secrets.PROBELY_API_KEY }} -o JSON | jq -r '.status')
            if [ $SCAN_STATUS == "started" ] || [ $SCAN_STATUS == "queued" ]; then
              echo "Scan is running or queued!";
            else
              echo "Scan is not running... finishing"
              break;
            fi
            sleep 30;
          done

      # Step 4: Optional logic - abort the pipeline if there are any HIGH risk vulnerabilities.
      - name: Check for High risk vulnerabilities
        run: |
          # Wait until scan ends
          HIGH_VULNS=$(probely scans get $SCAN_ID --api-key ${{ secrets.PROBELY_API_KEY }} -o JSON | jq -r '.highs')
          echo "HIGH risk vulnerabilities ${HIGH_VULNS}"
          if [[ "$HIGH_VULNS" -gt 0 ]]; then
            echo "Scan has High risk vulnerabilities... aborting"
            exit 1
          else
            echo "Scan doesn't have High risk vulnerabilities"
          fi
```

#### Run a scan on an ephemeral (dynamic) target in blocking mode

This is a more advanced configuration for building, deploying, and scanning an application in a temporary environment that is created for a specific purpose and then automatically destroyed during the pipeline run.

Using ephemeral environments requires agent token, target hostname, and target URL. Store them as variables for better security.

You also need to create a scanning agent in Snyk API & Web and configure your target to use it. This process requires the `scanning-agent/farcasterd-linux-amd64-0.4.3` file. For detailed instructions, visit How to install a Scanning Agent and How to scan internal applications.

{% hint style="info" %}
In this code example, Docker is used to create ephemeral environments. However, you can use any other solution to create your environment.
{% endhint %}

```yaml
# github-ephemeral-app-blocking-mode.yaml
name: CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      # Step 1
      - name: Checkout code
        uses: actions/checkout@v4

      # Step 2: Set up Docker
      - name: Set up Docker
        uses: docker/setup-buildx-action@v2

      # Step 3: Create a custom Docker network
      - name: Create custom Docker network
        run: docker network create custom-network

      # Step 4: Build and run the Docker app container
      - name: Build and run app container
        run: |
          # Build the Docker image
          docker build -t test-app .

          # Run the Docker container with a custom hostname
          docker run --name test-app \
            --hostname custom-web-app \
            --network custom-network \
            -p 8080:8080 \
            -d test-app

      # Additional steps omitted for brevity
      # See full example in the Snyk API & Web CI/CD examples repository
```

## Step 3: Run the pipeline and view the results

After committing your workflow file, the scan is triggered automatically on a push to the main branch. You can also run the workflow manually.

1. In your GitHub repository, go to the **Actions** tab.
2. Under **All workflows**, find and select your scan workflow.
3. Click the **Run workflow** dropdown, then click **Run workflow** to trigger a new scan.

Once the scan is complete, you can view the detailed findings in your Snyk API & Web dashboard.
