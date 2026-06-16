# Fail on authentication failure

Snyk API & Web allows you to force your scans to fail if authentication is unsuccessful.

To enable this feature:

1. Configure authentication for your Web or API target.
1. Select the checkbox that says **When login fails, fail the scan immediately and notify me**.
1. Click **Save**.

On your next scan, Snyk API & Web will try to authenticate. If this is not possible, or if reauthentication is not possible, the scan fails immediately, and you will be notified by email so that you can take action.
