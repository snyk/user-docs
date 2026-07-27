---
description: How to use the sequence recorder for Snyk API and Web web targets
nav_context: classic
---

# Use sequence recorder

Record login and navigation sequences with the Snyk API & Web Sequence Recorder browser plugin to help the scanner test authenticated and complex areas of your web application.

## About the sequence recorder

The Snyk API & Web Sequence Recorder is a browser plugin that records sequences to be replayed during target scans.

A sequence is a set of actions and values that the scanner follows to access specific areas of your application. Use sequences for:

* Login sequences: record authentication flows, including complex multi-step logins or two-factor authentication
* Navigation sequences: record workflows to reach specific application states or sections that require user interactions

## Install the sequence recorder

The Snyk API & Web Sequence Recorder is available for Chrome, Firefox, and Edge browsers.

* Chrome: install the plugin from the [Chrome Web Store](https://chrome.google.com/webstore/detail/probely-recorder/hldkejmiceccmcfgicnfbgminlidgkph).
* Firefox: install the plugin from the [Firefox Add-ons store](https://addons.mozilla.org/en-US/firefox/addon/probely-recorder/).
* Edge: install the plugin from the [Edge Add-ons store](https://microsoftedge.microsoft.com/addons/detail/snyk-api-web-sequence-r/bgdnfkliglfichfflhlgfpifnfbdgdgm/).

## Record a sequence

After installing the Snyk API & Web Sequence Recorder plugin, follow these steps to record a sequence:

1. Open your browser in incognito mode to ensure a clean session without cached data or existing authentication.
2. Click the plugin icon in your browser toolbar. A pop-up window appears.
3. Enter your target URL and click **Start recording**.

Note the following before recording:

* If you are recording a login sequence, ensure you are **logged out** of your target before starting.
* If you are recording a navigation sequence that requires authentication, ensure you are **logged in** to your target before starting.

4. Perform the necessary steps for your sequence. Click links and buttons using your mouse cursor, simulating how a user would interact with your application.

Avoid adding unnecessary steps that are not part of the workflow you want to record. The plugin registers each action you make in the page.

5. Stop the recording after you complete all steps. Click the plugin icon again and select **Stop recording**.
6. Review and edit your sequence on the Review page. The plugin provides advanced features to modify your sequence if needed.
7. Copy or download your sequence for later use. You can copy the sequence to your clipboard or download it as a file.

## Use the recorded sequence

After recording a sequence, import it into Snyk API & Web:

For login sequences: navigate to your target's **Authentication** settings and configure the login sequence. Visit [Configure login sequence authentication](../configure-authentication/configure-login-sequence.md) for detailed instructions.

For navigation sequences: navigate to your target's **Scanner** settings and add the navigation sequence. Visit [Use navigation sequences](use-navigation-sequences.md) for detailed instructions.

## Source code

The Snyk API & Web Sequence Recorder is open source. Visit the [GitHub repository](https://github.com/Probely/sequence-recorder) to view the plugin source code.
