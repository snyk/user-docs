---
description: How to export table data to CSV in Snyk API and Web
nav_context: classic
---

# Export table data to CSV

Whether you are performing a deep-dive audit or preparing a presentation for stakeholders, getting your data into your own workflow should be seamless.

The **Download CSV** feature lets you export filtered, customized table data directly from your Snyk API & Web account. This eliminates the need for manual copy-pasting and lets you manipulate your data in Excel, Google Sheets, or any BI tool of your choice.

## Prerequisites

Access is tied to table visibility. If you have permission to view a table, you have permission to export its data.

## How it works

The export tool provides as much flexibility as you need. When you download a CSV file, you can customize the fields you want to export by selecting them in the interface.

### Steps to export data

1. Navigate to the page that contains the table you want to export.
1. Click **Download CSV** above the table.
1. Select the fields you want to export.
1. Choose your delivery method. A prompt appears asking how you want to receive your file:
   * Wait for generation: stay on the page while the system processes the request. After it finishes, you can download the file directly to your device.
   * Receive by email: continue working in the app. Snyk sends you an email with the CSV file attached.

If you choose to wait for the generation but change your mind, you can still have the completed report sent to your email after it is ready.

## Technical details (API)

If you use the Snyk API, you can trigger the export functionality programmatically. The API respects the same filtering and field-selection parameters as the web interface, ensuring consistency across your automated workflows.

For more details on how to export data using the API, visit the [Export API reference](https://developers.probely.com/api/reference/export).
