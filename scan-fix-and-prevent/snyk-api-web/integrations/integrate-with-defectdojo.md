# Integrate with DefectDojo

By connecting Snyk API & Web to your DefectDojo server, you can synchronize target scan results with a DefectDojo product of your choice.

The synchronization is uni-directional, meaning that a finding reported by Snyk API & Web is sent to DefectDojo, but if its state changes at DefectDojo, the matching finding at Snyk API & Web will not have its state updated.

If Snyk API & Web detects a change, it updates the DefectDojo finding. For instance, if the underlying vulnerability is fixed, Snyk API & Web detects it, sets the finding as fixed, and updates the DefectDojo finding to fixed as well.

The DefectDojo instance is set at the account level and enabled on demand for each target. Configurations such as which product and engagement to use are set in the target settings.

This integration supports DefectDojo versions 1.5.X and 1.6.X.

## Configure DefectDojo API key

Snyk API & Web needs the URL of your DefectDojo server and an API v2 Key to authenticate itself. The API key must belong to a staff user.

You can find your API v2 Key at `<your DefectDojo>/api/key-v2` or by clicking in the top right dropdown and click **API v2 Key**.

Copy the API key value.

Go to `https://plus.probely.app/integrations` and enter your DefectDojo URL and the copied API key. It looks like this:

<figure><img src="../../../../.gitbook/assets/integrate-with-defectdojo-configure-api.png" alt="DefectDojo integration form showing URL and API key fields"></figure>

Click **Save**.

Snyk API & Web tries to connect and authenticate to DefectDojo, and a success message appears. Done.

If the Snyk API & Web servers cannot connect or the API key is incorrect, an error is displayed. Review your configuration and ensure your server can receive connections from Snyk API & Web IPs.

## Choose your synchronization settings

You need to choose which targets to synchronize and how. To configure a target to use DefectDojo go to its settings at **Settings > Integrations** and then **DefectDojo**.

You see the following screen:

<figure><img src="../../../../.gitbook/assets/integrate-with-defectdojo-target-settings.png" alt="DefectDojo target synchronization settings form showing product, engagement, test, and options"></figure>

### Product

Choose which DefectDojo product to sync with.

### Engagement

Choose which engagement to sync with. The list only shows engagements for the selected product.

### Test

An optional name to identify Snyk API & Web scans. If empty, the target scans can be identified by the test type **Snyk API & Web Scan**. The test type is created automatically when the integration is configured.

### Set findings to active/verified

Sets the findings reported by Snyk API & Web to active/verified.

These are enabled by default to ensure findings get adequate visibility at DefectDojo. Non-active and non-verified findings might not be visible in the DefectDojo dashboards.

### Delete

Removes the configuration for this target. Findings already synchronized are not affected.
