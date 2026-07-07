# Fail on authentication failure

In Snyk API & Web, you can force your scans to fail if authentication is unsuccessful.

To enable this feature:

1. Configure authentication for your Web or API target.
1. Select the **When login fails, fail the scan immediately and notify me** check box.
1. Click **Save**.

On your next scan, Snyk tries to authenticate. If authentication or reauthentication is not possible, the scan fails immediately, and Snyk notifies you by email so that you can take action.
