# Switch the scan profile

Switch between different scan profiles for your targets.

Snyk API & Web provides a variety of [scan profiles](built-in-scan-profiles.md) that you can choose from, depending on how thoroughly you want to scan your target.

You may decide to run a **Normal** scan on your target to start with and then decide to run a **Full** scan later after fixing some vulnerabilities. Or it might be the case that, at a specific moment in time, you want to run a **Safe** scan, in order to reduce the possible impact on your target.

All of this is possible. The only thing you need to do is switch the scan profile before you start a new scan.

## Switch scan profile in the Web UI

To switch the scan profile in the Snyk API & Web interface:

1. Access your target settings and click the **Profile** tab. A list of all available scan profiles is displayed according to your target's type and current verification state.
1. Choose the scan profile you want and click **Save**.

The next scan will use the selected profile.

## Switch scan profile using the API

If you are using the Snyk API & Web API, you can also send a different scan profile in the request the next time you start a target scan. Learn more about this through the [API documentation](https://developers.probely.com/).

Regardless of how you choose your scan profile and start a target scan, make sure to use the profile that makes the most sense to you and your target.
