# Integrate with DefectDojo

By connecting Snyk API & Web to your DefectDojo server, you can synchronize target scan results with a DefectDojo product of your choice.

The synchronization is unidirectional. Snyk sends a finding it reports to DefectDojo, but if the state of that finding changes in DefectDojo, Snyk does not update the state of the matching finding.

When Snyk detects a change, it updates the DefectDojo finding. For example, when the underlying vulnerability is fixed, Snyk detects this, sets the finding as fixed, and updates the DefectDojo finding to fixed.

You set the DefectDojo instance at the account level and enable it on demand for each target. You set configurations such as which product and engagement to use in the target settings.

This integration supports DefectDojo versions 1.5.X and 1.6.X.

## Configure DefectDojo API key

Snyk needs the URL of your DefectDojo server and an API v2 Key to authenticate. The API key must belong to a staff user.

You can find your API v2 Key at `<your DefectDojo>/api/key-v2` or by clicking the top-right dropdown and selecting **API v2 Key**.

Copy the API key value.

Navigate to `https://plus.probely.app/integrations` and enter your DefectDojo URL and the copied API key. It looks like this:

Click **Save**.

Snyk connects and authenticates to DefectDojo, and a success message appears.

If the Snyk servers cannot connect or the API key is incorrect, an error appears. Review your configuration and ensure your server can receive connections from Snyk IPs.

## Choose your synchronization settings

Choose which targets to synchronize and how. To configure a target to use DefectDojo, navigate to its settings at **Settings > Integrations** and then **DefectDojo**.

You see the following screen:

### Product

Select which DefectDojo product to synchronize with.

### Engagement

Select which engagement to synchronize with. The list shows only engagements for the selected product.

### Test

An optional name to identify Snyk scans. If empty, the test type **Snyk API & Web Scan** identifies the target scans. Snyk creates the test type automatically when you configure the integration.

### Set findings to active/verified

Sets the findings that Snyk reports to active/verified.

These are enabled by default to ensure findings get adequate visibility in DefectDojo. Non-active and non-verified findings can be hidden in the DefectDojo dashboards.

### Delete

Removes the configuration for this target. Findings already synchronized are not affected.
