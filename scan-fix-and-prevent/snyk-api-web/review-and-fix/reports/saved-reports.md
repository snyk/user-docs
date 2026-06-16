# Saved reports

With Snyk API & Web, you can download reports of specific scans (scan reports), or reports based on search criteria you define, that can comprise multiple targets (saved reports).

This feature (available on the Enterprise plan only) allows you to, for instance, get a report of all High findings across all targets from a specific team.

## Generate a saved report

To generate these reports, go to the global list of **Findings**, apply the search terms and filters wanted, and click the **Reports** button from the top right corner of the page.

<figure><img src="../../../../.gitbook/assets/saved-reports-findings-list.png" alt="Global Findings list with Reports button in top right corner"></figure>

Once you do, a dropdown menu appears with 3 options:

<figure><img src="../../../../.gitbook/assets/saved-reports-dropdown-menu.png" alt="Reports dropdown menu showing Generate now, Save report, and Manage reports options"></figure>

### Generate now

This generates a PDF report of the findings listed, taking into account the search and filters applied on the interface.

Similarly to what happens with scan reports, you can choose to send the report by email, or wait for it to be generated and then download it.

<figure><img src="../../../../.gitbook/assets/saved-reports-generate-now-dialog.png" alt="Generate report dialog showing options to receive by email or download when ready"></figure>

<figure><img src="../../../../.gitbook/assets/saved-reports-generate-now-confirmation.png" alt="Confirmation message after report generation request"></figure>

### Save report

This saves a new report (that takes into account the search and filters applied on the interface) and lists it under Manage reports.

Reports can be saved with or without recurrence. If recurrence is configured, you automatically receive the respective report in your email address at the "next report" date.

<figure><img src="../../../../.gitbook/assets/saved-reports-save-report-dialog.png" alt="Save report dialog showing name, recurrence, and next report date fields"></figure>

### Manage reports

This option redirects to a new page where you can manage the reports already saved using the previous option.

<figure><img src="../../../../.gitbook/assets/saved-reports-manage-reports-page.png" alt="Manage reports page showing saved reports with Download and Delete buttons"></figure>

* You can click the **Findings list** link in the Findings column to access the global Findings list filtered by the same criteria that were used when saving the report. This gives you an overview of the findings that match the criteria applied and you can even tweak this pre-filtered view to generate new reports with similar filters, if needed.
* You can click the **Download** button to download the current report taking into account the filters and search defined. This means every time the report is downloaded, the results may be different (that is, the findings from the report are always up to date).
* You can click the **Delete** button to delete the report. If you do, any upcoming reports are canceled and you no longer receive them by email.
