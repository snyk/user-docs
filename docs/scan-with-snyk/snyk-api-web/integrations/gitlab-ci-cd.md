# Integrate Snyk API & Web with GitLab CI/CD

This guide provides step-by-step instructions for integrating Snyk API & Web into your GitLab CI/CD pipelines.

<span class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&:nth-child(n+2)]:hidden lg:[&:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9"><span class="text-lg leading-6">C</span></span>

<span class="text-body-secondary-color"></span>

This guide focuses on using the Snyk API & Web CLI to run scans. The examples below cover the complete end-to-end journey, from configuring your targets to running scans in different scenarios.

# **Prerequisites**

Before you begin, you must configure your scan targets and credentials in the Snyk API & Web application.

## **Create a target in Snyk API & Web**

In the Snyk API & Web app, go to the **Targets** menu and click the **Add** button. Fill out the form and click **Add** to create the new target.

**Note:** During this process, we check for connectivity. If your target is internal or not yet deployed, you can bypass any warnings and add the target regardless. For more details, see [How to add a Target](https://help.probely.com/en/articles/5733114-how-to-add-a-target).

Before configuring the integration in GitLab, make sure to retrieve the unique **target ID** from Snyk API & Web.

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In your Snyk API & Web dashboard, select **Targets**.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    From the target list, select the target you want to integrate.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In your browser's address bar, copy the **target ID**. This is the string of characters immediately following /target/ in the URL.

    </div>

**Note:** After creating a target, it is mandatory to verify your target’s domain. Otherwise, your scans are only limited to lightning scans. To learn more, see our documentation about the [importance of domain ownership verification](https://help.probely.com/en/articles/3285602-why-do-we-require-you-to-verify-the-ownership-of-a-domain).

## **Create a Snyk API & Web API Key**

You need an API key with permissions to start a scan on your target. For instructions, see [How to generate an API key](https://help.probely.com/en/articles/8592281-how-to-generate-an-api-key).

# **Step 1: Add your API Key and target ID to GitLab**

To run a scan, your pipeline needs to authenticate with Snyk API & Web and know which target to scan. You must configure your Snyk API & Web **API Key** and the specific **target ID** as <a href="https://docs.gitlab.com/ci/variables/#define-a-cicd-variable-in-the-ui" rel="nofollow noopener noreferrer" target="_blank">secure CI/CD variables</a> in your GitLab project.

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    From your GitLab project side menu, navigate to **Settings \> CI/CD**.\
    ​

    </div>

    <div class="intercom-interblocks-image intercom-interblocks-align-left">

    <a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1803464792/bbe9cbeb82cd7dea108a069af87d/Integrate_CICD_GitLab_variables_1_markup.png?expires=1769985000&amp;signature=7b111b75e8e62bfc8aa6246c52e18ed0c82c203ad1b2fbd58674807cfdb56b51&amp;req=dSgnFc14mYZWW%2FMW1HO4zUhqETjJhQzJCsqHneRYrZKNGa6sGROcElcRaKxV%0AJyWQ%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/ccdb3e7e27da4490c3df350ff4bc5f8f96f9ef88.png" width="3024" height="1718" /></a>

    </div>

    <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Find the **Variables** section and expand it.\
    ​

    </div>

    <div class="intercom-interblocks-image intercom-interblocks-align-left">

    <a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1803465528/f07c8129da47bc12d4e122fcd61e/Integrate_CICD_GitLab_variables_2_markup.png?expires=1769985000&amp;signature=2e772cb77a09dcfb9fd14fb47ca1886f248d021f79a7981a74392a01101a735b&amp;req=dSgnFc14mIRdUfMW1HO4zSv4H%2F%2FSdXZzEvxgsXj1%2F55bVdX%2BqzaZ1IdPANRa%0AWBJx%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/143166f3bb333d3af77de268537847f7b9e3e7d8.png" width="3020" height="1712" /></a>

    </div>

    <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click **Add variable** and create an entry for your Snyk API & Web **API Key**, for example, **PROBELY_API_KEY**.\
    ​

    </div>

    <div class="intercom-interblocks-image intercom-interblocks-align-left">

    <a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1803466432/4e168a4b95ce8ecd1bb74489f22f/Integrate_CICD_GitLab_variables_3_markup.png?expires=1769985000&amp;signature=d0e3f1ed7e8010fced994b578fde07a07ef8f9b66b4ddae4c380f7cab2cd4fbd&amp;req=dSgnFc14m4VcW%2FMW1HO4zSl%2FQVA%2Ff3ezTBhj1tPhMqTGFUxzRJ8DY5MBiBtJ%0AYLc9%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/6a7c10448adaffd28b76daccb0ab58682e4951bb.png" width="3020" height="1716" /></a>

    </div>

    <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    </div>

    <div class="intercom-interblocks-image intercom-interblocks-align-left">

    <a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1803467774/f9c03b99191432946800c40043e5/Integrate_CICD_GitLab_variables_4.jpg?expires=1769985000&amp;signature=9024f190e9221c998a287157be9bdbc360a616e42c602b0ca40e257965a16724&amp;req=dSgnFc14moZYXfMW1HO4zU%2Bg9t5Pt7yhFBeO6c0MkkRglePcaOhq1zjPsNGT%0A2mnp%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/131c762104299ad5d0c825270e3873caa808f12e.jpg" width="1600" height="513" /></a>

    </div>

    <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    </div>

4.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click **Add variable** again to create a second entry for your **target ID**, for example, **TARGET_ID**.

    </div>

**Important:** For enhanced security, always store sensitive values as GitLab <a href="https://docs.gitlab.com/ci/variables/#define-a-cicd-variable-in-the-ui" rel="nofollow noopener noreferrer" target="_blank">CI/CD variables</a>. Storing variables directly in your `.gitlab-ci.yml` file is not recommended, as they are saved in plain text and visible to anyone who can view the file.\
​

# **Step 2: Configure your pipeline**

Create a `.gitlab-ci.yml` file at the root of your repository and add one of the following code examples based on your use case. You can also find all of them in our <a href="https://github.com/Probely/cicd-pipeline-scan-examples/tree/main/cicd-examples/gitlab" rel="nofollow noopener noreferrer" target="_blank">Snyk API &amp; Web CI/CD examples repository on GitHub</a>.

## **Important note on these examples**

The YAML configurations below are scanning steps designed to be incorporated into your existing `.gitlab-ci.yml` file.

For example, your pipeline might already have steps to build your code, deploy to a QA environment, and run automated tests. You can add the Snyk API & Web scan as another step at any point that makes sense for your workflow, such as after you deploy to QA or staging.

## **Run a scan on a target in non-blocking mode**

This is the simplest way to get started. It runs a scan on a remotely accessible environment without blocking the pipeline.\
​

    # gitlab-remote-app-non-blocking-mode.yaml
    stages:
      - scan

    scan:
      stage: scan
      image: python:latest
      script:
        - pip install probely
        - probely targets get --api-key ${PROBELY_API_KEY}
        - |
          for i in {1..20}; do
            echo "-----------------------------------"
            SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key ${PROBELY_API_KEY})
            echo ${SCAN_ID}
            if [[ -z "$SCAN_ID" ]]; then
              echo "Scan didn't start... Retry start-scan"
            else
              echo "Scan started with SCAN ID: ${SCAN_ID}";
              break;
            fi
            sleep 5
          done
          if [[ -z "$SCAN_ID" ]]; then
            echo "No Scan ID, aborting..."
            exit 1
          fi

## **Run a scan on a target in blocking mode**

This configuration builds on the first example by adding steps to poll the scan status and block the pipeline if high-severity vulnerabilities are found.

    # gitlab-remote-app-blocking-mode.yaml
    stages:
      - scan

    scan:
      stage: scan
      image: python:3.11-bullseye
      script:
        - apt-get update && apt-get install -y jq
        # Install Snyk API & Web CLI
        - pip install probely
        - probely targets get --api-key ${PROBELY_API_KEY}
        - |
          for i in {1..20}; do
            echo "-----------------------------------"
            SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key ${PROBELY_API_KEY})
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
            exit 1
          fi

        - |
          while true; do
            echo "-----------------------------------"
            SCAN_OUTPUT=$(probely scans get ${SCAN_ID} --api-key ${PROBELY_API_KEY} | tail -1)
            echo ${SCAN_OUTPUT}
            echo "-----------------------------------"
            SCAN_STATUS=$(probely scans get ${SCAN_ID} --api-key ${PROBELY_API_KEY} -o JSON | jq -r '.status')
            if [[ "$SCAN_STATUS" == "started" ]] || [[ "$SCAN_STATUS" == "queued" ]]; then
              echo "Scan is running or queued!";
            else
              echo "Scan is not running... finishing"
              break;
            fi
            sleep 30;
          done

        - HIGH_VULNS=$(probely scans get ${SCAN_ID} --api-key ${PROBELY_API_KEY} -o JSON | jq -r '.highs')
        - echo "HIGH risk vulnerabilities ${HIGH_VULNS}"
        - |
          if [[ "$HIGH_VULNS" -gt 0 ]]; then
            echo "Scan has High risk vulnerabilities... aborting"
            exit 1
          else
            echo "Scan doesn't have High risk vulnerabilities"
          fi

\
**Run a scan on an ephemeral (dynamic) target in blocking mode**
----------------------------------------------------------------

This is a more advanced configuration for building, deploying, and scanning an application in a temporary environment that is created for a specific purpose and then automatically destroyed during the pipeline run.

Using ephemeral environments requires agent token, target hostname and target URL. We recommend storing them as variables for better security.

You also need to create a scanning agent in Snyk API & Web and configure your target to use it. This process requires the `scanning-agent/farcasterd-linux-amd64-0.4.3` file. For detailed instructions, see <a href="https://www.google.com/search?q=link-to-article" rel="nofollow noopener noreferrer" target="_blank">How to install a Scanning Agent</a> and <a href="https://www.google.com/search?q=link-to-article" rel="nofollow noopener noreferrer" target="_blank">How to scan internal applications</a>.

**Note:** In this code example, Docker is used to create ephemeral environments. However, you can use any other solution to create your environment.

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1803479389/e0748bbc480674bdcd9533032fab/Integrate_CICD_GitLab_variables_3.png?expires=1769985000&amp;signature=4e893e00c42412deb88a3bd291c504febf9de70b5fa4976134dfbb006316b6e6&amp;req=dSgnFc15lIJXUPMW1HO4zbgzn4lT3OzBNGpcA0d4FE0geMiUa6oFnS0eLZwF%0A7RKx3dzXHjdGZYgQ6hQ%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/d34ac6aa1559aec496a740b429351304494be058.png" width="3020" height="1716" /></a>

    # gitlab-ephemeral-app-blocking-mode.yaml
    stages:
      - build-and-test

    build-and-test:
      stage: build-and-test
      image: docker:latest
      services:
        - docker:dind
      variables:
        DOCKER_HOST: tcp://docker:2375
      script:
        - apk add --no-cache curl jq python3 py3-pip
        - python3 -m venv venv
        - source ./venv/bin/activate
        # Install Snyk API & Web CLI
        - pip install probely
        - probely targets get --api-key ${PROBELY_API_KEY}

        - docker network create custom-network

        - docker build -t test-app . 
        - docker run --name test-app --hostname custom-web-app --network custom-network -p 0.0.0.0:8080:8080 -d test-app

        - cat /etc/hosts # current /etc/hosts
        
        - CONTAINER_IP=$(grep -i 'docker' /etc/hosts | head -1 | awk '{print $1}')
        - echo "Container IP from /etc/hosts is $CONTAINER_IP"
        - echo "${CONTAINER_IP} ${TARGET_HOSTNAME} ${TARGET_HOSTNAME}." | tee -a /etc/hosts  # Add to /etc/hosts
        - cat /etc/hosts  # Confirm host was added

        - |
          for i in {1..10}; do  # Wait for the app to start
            if curl -s ${TARGET_URL} > /dev/null; then
              echo "App is up!";
              break;
            fi
            echo "Waiting for the app to be ready...";
            sleep 2;
          done

        # Test the application
        - RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" ${TARGET_URL})
        - |
          if [[ "$RESPONSE" -ne 200 ]]; then
            echo "App test failed with HTTP status ${RESPONSE}";
            exit 1;
          fi
        - curl -s -i ${TARGET_URL}
        - echo "App test passed with HTTP status ${RESPONSE}";

        # Run userspace agent
        - chmod +x scanning-agent/farcasterd-linux-amd64-0.4.3 
        - ./scanning-agent/farcasterd-linux-amd64-0.4.3 --token ${AGENT_TOKEN} &

        # Wait for the agent to start
        - sleep 40

        # Start Scan
        - |
          for i in {1..20}; do  # Start Snyk API & Web scan
            echo "-----------------------------------"
            SCAN_ID=$(probely targets start-scan ${TARGET_ID} -o IDS_ONLY --api-key ${PROBELY_API_KEY})
            echo ${SCAN_ID}
            if [[ -z "${SCAN_ID}" ]]; then
              echo "Scan didn't start... Retry start-scan"
            else
              echo "Scan started with SCAN ID ${SCAN_ID}";
              break;
            fi
            sleep 5
          done
        - |
          if [[ -z "${SCAN_ID}" ]]; then
            echo "No Scan ID, aborting..."
            exit 1
          fi

        # Wait for scan to end
        - |
          while true; do
            echo "-----------------------------------"
            SCAN_OUTPUT=$(probely scans get ${SCAN_ID} --api-key ${PROBELY_API_KEY} | tail -1)
            echo ${SCAN_OUTPUT}
            echo "-----------------------------------"
            SCAN_STATUS=$(probely scans get ${SCAN_ID} --api-key ${PROBELY_API_KEY} -o JSON | jq -r '.status')
            if [[ "$SCAN_STATUS" == "started" ]] || [[ "$SCAN_STATUS" == "queued" ]]; then
              echo "Scan is running or queued!";
            else
              echo "Scan is not running... finishing"
              break;
            fi
            sleep 30;
          done

        # Check for high vulnerabilities
        - HIGH_VULNS=$(probely scans get ${SCAN_ID} --api-key ${PROBELY_API_KEY} -o JSON | jq -r '.highs')
        - echo "HIGH vulnerabilities  ${HIGH_VULNS}"
        - |
          if [[ "$HIGH_VULNS" -gt 0 ]]; then
            echo "Scan has High vulnerabilities... aborting"
            exit 1
          else
            "Scan doesn't have high vulnerabilities"
          fi

        # Clean up
        - docker stop test-app
        - docker rm test-app
        -dockernetworkrmcustom-network

# **Step 3: Run the pipeline and view the results**

After committing your `.gitlab-ci.yml` file, you can run the pipeline in GitLab to test the integration.

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1803480949/3cb968e6931131a656d705f46455/Integrate_CICD_GitLab_run_pipeline_in_GitLabUI.png?expires=1769985000&amp;signature=70dc51517e65973e832c3258f947c281afedeca92fb01d2e39bdf26c7473312a&amp;req=dSgnFc12nYhbUPMW1HO4zW1OXCj9W7alMQThDLA75YREJzfXBoaAijIf4yaw%0AfEGcG1H5jfSkrUMpx4k%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/153ed72df215434b0854d558b0e81240dd867245.png" width="2261" height="1056" /></a>

Once the scan is complete, you can view the detailed findings in your Snyk API & Web dashboard.

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1803481591/805937d7a3448be09aeed8f63742/Integrate_CICD_GitLab_SnykAPIWeb_dashboard.png?expires=1769985000&amp;signature=ad13832a92f788979611f8fdd2eae3494f795f5a38efd70b79b71910b8266fd5&amp;req=dSgnFc12nIRWWPMW1HO4zTPCi49G04oeJkbCta96mQPA4BIH9CxswUDfRnlL%0AzkkSGsI1AW1sbJIiptE%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/3d84cfe05a543c839aaafeaf087e6aa059186d42.png" width="2270" height="1252" /></a>

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F12692938-integrate-snyk-api-web-with-gitlab-ci-cd&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

