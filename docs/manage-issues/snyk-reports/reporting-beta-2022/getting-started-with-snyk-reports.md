# Getting started with Snyk Reports

Reports are available during Open Beta on the **Reports** tab in the Snyk navigation. To get access to the **Reporting Beta,** a Snyk admin can activate the reporting in [Snyk Preview](https://docs.snyk.io/features/user-and-group-management/managing-settings/snyk-preview).

You can navigate to different reports from the dropdown in the upper left. The report currently selected is displayed.

## Snyk reporting filters

All reports are filterable by a number of attributes, enabling users to create the appropriate horizontal or vertical slice of the business for their use case. You can see the available filters on the dropdown furthest right at the top.

If you do not select or enter any values for a particular key, the filter is not applied.

### Stateful URLs for filtered views

Every time a filter value is applied, the app.snyk.io URL is updated to persist the state of the page. You can bookmark and copy the URL, and share it with anyone who has the appropriate Snyk Organization or Group access. For easy sharing, use the copy URL button in the upper right corner of the list of reports.

### Snyk reporting filter logic

Within a given filter, all values selected are separated by an OR operator. For example, if you select the values `Critical` and `High` for the `Issue Severity` filter, Snyk displays issues that have a severity of either `critical` or `high`.\\

<figure><img src="https://lh5.googleusercontent.com/OztJ-7tKSZpq2KPKXa6Sjr2vH2GHABcMKOdDNlw6MVShpcpnk7GZHZWz6mE3EeQBGB8SjHVUrVkeMC4IqA1YUhFSa41aNS8z_ZRsHeBxgRnMrp-YPwadg0u9hxuAtStePCrBWp9QV0GswTVozsLR7w7432etIbB6ij6t1t_Dkj5-udb9ofzTfiQw4UqtYuw7FjMEUg" alt="Issue severity values selected"><figcaption><p>Issue severity values selected</p></figcaption></figure>

Filters are separated by an AND operator. For example, if you select the `Critical` value for the `Issue Severity` filter and the `Resolved` value for the `Issue Status` filter, Snyk displays issues that are both `critical` severity and `resolved`.

## Exporting Snyk Reports

You can export reports to PDF and download tabular data within a report to CSV.

### Export Snyk Reports to PDF

Use the **Export to PDF** button at the top right to download a PDF of not only the content of a report, but also its context, including who ran the report, when it was exported, the scope being viewed (Snyk Organization or Group), and any filters applied.

Exporting a PDF allows sharing a report with a user who will not or should not authenticate into the Snyk app, such as an executive or external auditor. **Export to PDF** offers point-in-time attestation with the necessary context.

The following shows an example of a PDF export.

<figure><img src="https://lh3.googleusercontent.com/e_QTKFb7mnnuRr6ljPTIOYb3zUPRosEZ2SgSRjGWiIViVD0uTIzyIhb3Cu6rIxgIn85zkvrdQ4AE1uKj_H3hL77kyhpYcIW8j-JxPYcNscluE9_O2BbbTRpIQuHBo8pFNQ_cLpuiTJAtKHfOP37lsRbWsbhmi_uyhNJZf5gZo66yWuspaApOFyVq3w" alt="Example of PDF export of Snyk report"><figcaption><p>Example of PDF export of Snyk Report</p></figcaption></figure>

For tabular data on reports, only ten results are shown in the PDF export. Snyk plans to extend the row count and enable links to the report for full details.

### Download Snyk Report to CSV file

You can use the **Download CSV** button at the right in the report to download data presented in tables to CSV. This information can be used, for example, for prioritization purposes or a one-time analysis of data in Excel or a similar tool.

All columns displayed in the UI are included in the CSV output. In addition, any columns that contain hyperlinks in the UI are split into two columns: one containing the text and the other containing the linked URL.

There is no row limit, but there is a 5GB file size limit.

## Column sorting

Sort columns within tables by clicking the arrows next to the column header. Click once to sort in ascending order, twice to sort in descending order, and a third time to remove sort from that column. Multi-column sorting is supported.

When columns are sorted, the app.snyk.io URL is updated to persist the state of the page, allowing for bookmarking, copying, and sharing.

## Modifying Snyk Report columns

In some reports, tables may include an option to modify columns. When this option is available, you can use it to select the columns to display in the UI. The export features (PDF and CSV) respect the selected columns.

When columns are modified, the app.snyk.io URL is updated to persist the state of the page, allowing for bookmarking, copying, and sharing.
