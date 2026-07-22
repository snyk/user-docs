---
description: How to switch the Snyk API and Web scan profile
---

# Switch the scan profile

Switch between different scan profiles for your targets.

Snyk API & Web provides a variety of [scan profiles](built-in-scan-profiles.md) that you can choose from, depending on how thoroughly you want to scan your target.

For example, you can run a **Normal** scan on your target to start with, then run a **Full** scan later after fixing some vulnerabilities. You can also run a **Safe** scan at a specific time to reduce the possible impact on your target. To do any of this, switch the scan profile before you start a new scan.

## Switch scan profile in the Web UI

To switch the scan profile in the Snyk API & Web interface:

1. Access your target settings and click the **Profile** tab. Snyk displays a list of all available scan profiles according to your target's type and current verification state.
1. Choose the scan profile you want and click **Save**.

The next scan uses the selected profile.

## Switch scan profile using the API

If you are using the Snyk API & Web API, you can also send a different scan profile in the request the next time you start a target scan. Learn more in the [API documentation](https://developers.probely.com/).
