# Troubleshooting Snyk Reports

## Access to reporting

If reporting is not loading in the Snyk UI, follow these troubleshooting steps that may help resolve the issue.

{% hint style="info" %}
Allowlisting may be necessary due to some firewall settings. If allowlisting snyk.io was required for the initial implementation, additional allowlisting may be needed for new reporting access. Contact your account team for more information.
{% endhint %}

If you still have issues accessing the new reporting after using these troubleshooting tips, [submit a ticket](https://support.snyk.io/hc/en-us/requests/new) to Snyk support.

## Snyk Code

Resolved issues are issues detected in the previous scan snapshot that no longer appear in the current snapshot.\
\
Changes in the Snyk code **resolved** column between snapshots may occur for any of the following reasons:

* The issue was fixed between the two scans and is therefore marked as resolved.
* The engine rules were improved, which may change the scan results.\
  This may add new issues or resolve previously detected issues.
* The code changed in a way that caused the issue to be defined differently, thus changing the issue ID.\
  In this situation, the original issue ID is marked as **resolved,** and a new issue ID is created.

{% hint style="info" %}
Historical data for Snyk Code issue reports entered General Availability (GA) in November 2022.
{% endhint %}

## Data freshness

Data is available in reporting approximately one hour after a scan occurs. Data refreshes hourly and should, therefore, always be updated within two hours. Note the exception for ignores that follows.

## Ignores in reporting

After a Project rescan, issues that are ignored are reflected in reporting and the rest of the Snyk platform, including APIs. Likewise, Projects for which issues have been unignored are reflected in reporting after a rescan of the Project. This does not apply to an ignore in a `.snyk` file.

For information about Project test frequency, see [Project Settings](../../snyk-admin/snyk-projects/view-and-edit-project-settings.md).

## Known limitations

Tables in PDF exports are limited to 50 results. A link is available at the bottom of the table in the PDF export to view the report in the browser.

Session data is shared between browser tabs. Snyk recommends you have reporting tabs for one Group or one Organization open in the browser at any one time because having tabs open for different Groups or Organizations can result in unexpected behavior.

Customers may experience an error if hundreds of filter values are applied on a page. This is the result of the [stateful URLs for filtered views](getting-started-with-snyk-reports.md#stateful-urls-for-filtered-views) capability and a URL length limit. If you experience an error with many filter values applied, consider alternate filter mechanisms such as [Project collections](../../snyk-admin/snyk-projects/project-collections-groupings/project-collections.md).

## Filtered views

The way filtered views are managed in the URL may change over time.&#x20;

If the view changes, you can generate a new filtered URL in the Snyk Web UI and save it as a bookmark or share it with others.

## Deactivated Projects

Deactivated Projects and their respective results will not appear in the Reports area.
