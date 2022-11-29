# Troubleshooting Snyk Reports

## Access

If reporting is not loading in the Snyk UI, the following troubleshooting steps may help resolve the issue.

**Note:** Whitelisting may be necessary due to some firewall settings. If whitelisting snyk.io was required for the initial implementation, additional whitelisting may be needed for new reporting access. Contact your account team for more information.

If you still have issues accessing the new reporting after using these troubleshooting tips, [submit a ticket](https://support.snyk.io/hc/en-us/requests/new) to Snyk support.

## Snyk Code

Resolved issues are issues detected in the previous scan snapshot that no longer appear in the current snapshot.\
\
Changes in the Snyk code **resolved** column between snapshots may be caused by any of the following reasons:

* The issue was fixed between the two scans and is therefore marked as resolved.
* The engine rules were improved, which may change the scan results. \
  This may add new issues or resolve previously detected issues.
* The code changed in a way that caused the issue to be defined differently, thus changing the issue ID. \
  In this situation, the original issue ID is marked as **resolved** and a new issue ID is created.

{% hint style="info" %}
Historical data for Snyk Code issue reports officially entered General Availability (GA) in November 2022.
{% endhint %}

## Data freshness

Data is available in reporting approximately one hour after a scan occurs.

## Known limitations

Exporting to PDF currently includes 10 results in table views.  Improvements to export will be implemented over time.

## Filtered views

The way filtered views are managed in the URL may change over time. \
If the view changes, you can generate a new filtered URL in the Snyk Web UI and save it as a bookmark or share it with others.&#x20;

