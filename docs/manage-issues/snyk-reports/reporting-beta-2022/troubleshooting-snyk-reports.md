# Troubleshooting Snyk Reports

## Access to reporting

If reporting is not loading in the Snyk UI, follow these troubleshooting steps that may help resolve the issue.

**Note:** Whitelisting may be necessary due to some firewall settings. If whitelisting snyk.io was required for the initial implementation, additional whitelisting may be needed for new reporting access. Contact your account team for more information.

If you still have issues accessing the new reporting after using these troubleshooting tips, [submit a ticket](https://support.snyk.io/hc/en-us/requests/new) to Snyk support.

## Snyk Code

Resolved issues are issues detected in the previous scan snapshot that no longer appear in the current snapshot.\
\
Changes in the Snyk code **resolved** column between snapshots may be caused for any of the following reasons:

* The issue was fixed between the two scans and is therefore marked as resolved.
* The engine rules were improved, which may change the scan results.\
  This may add new issues or resolve previously detected issues.
* The code changed in a way that caused the issue to be defined differently, thus changing the issue ID.\
  In this situation, the original issue ID is marked as **resolved** and a new issue ID is created.

{% hint style="info" %}
Historical data for Snyk Code issue reports entered General Availability (GA) in November 2022.
{% endhint %}

## Data freshness

Data is available in reporting approximately one hour after a scan occurs.

## Ignores in reporting

Issues that are ignored are reflected in reporting, and in the rest of the Snyk platform including APIs, after a project rescan.

## Known limitations

Tables in PDF exports are limited to 50 results. A link is available at the bottom of the table in the PDF export to view the report in the browser.

Session data is shared between browser tabs. Snyk recommends you have reporting tabs for one Group or one Org open in the browser at any one time, because having tabs open for different Groups or Orgs can result in unexpected behavior.

## Filtered views

The way filtered views are managed in the URL may change over time.\
If the view changes, you can generate a new filtered URL in the Snyk Web UI and save it as a bookmark or share it with others.

## Deactivated projects

Deactivated projects and their respective results will not appear in the Reports area.
