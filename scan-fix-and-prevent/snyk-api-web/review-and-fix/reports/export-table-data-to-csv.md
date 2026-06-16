# Export table data to CSV

Whether you are performing a deep-dive audit or preparing a presentation for stakeholders, getting your data into your own workflow should be seamless.

The **Download CSV** feature allows you to export filtered, customized table data directly from your Snyk API & Web account. This eliminates the need for manual copy-pasting and allows you to manipulate your data in Excel, Google Sheets, or any BI tool of your choice.

## Prerequisites

Access is tied to table visibility. If you have permission to view a table, you have permission to export its data.

## How it works

The export tool is designed to provide you with as much flexibility as you need. When choosing to download a CSV file, you can customize the fields you want to export (by selecting them on the interface).

### Steps to export data

1. Navigate to the page that contains the table you want to export.
1. Click the **Download CSV** button above the table.
1. Select the fields you want to export.
1. Choose your delivery method. A prompt appears asking how you would like to receive your file:
   * **Wait for generation:** stay on the page while the system processes the request. Once finished, you can download the file directly to your device.
   * **Receive via email:** you can continue working on the app. Snyk API & Web sends you an email with the CSV file attached.

If you choose to wait for the generation but change your mind, you can still opt to have the completed report sent to your email once it is ready.

## Technical details (API)

For users leveraging the Snyk API & Web's API, the export functionality can be triggered programmatically. The API respects the same filtering and field-selection parameters as the Web interface, ensuring consistency across your automated workflows.

For more details on how to export data using the API, see [Export API Reference](https://developers.probely.com/api/reference/export).
