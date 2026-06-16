# Schedule scan

You can set up a schedule so that scans start automatically on specific dates and times.

Creating a scheduled scan involves two steps:

1. Select the action to create scheduled scans.
1. Configure the scheduled scan by filling out a form.

## Select the action to schedule a scan

To start the creation of a scheduled scan for a single target, go to the Targets section and do one of the following:

* From the scan actions available on the page, click the caret icon within the **Scan** button, and choose **Schedule scan**.
* Alternatively, go to the target details page (any tab), click the caret icon within the **Scan** button, and choose **Schedule scan**.

To start the creation of a scheduled scan for multiple targets, go to the Targets section and do the following:

1. Tick the checkboxes to select the targets from the list. If all selected targets have scans in the same state, the button for bulk scan actions becomes available.
1. Click the caret icon within the **Scan** button, and choose **Schedule scan**.

## Configure the scheduled scan

Fill out the form to configure the scheduled scan:

* **Start date** - The date on which the scan starts.
* **Hour and Minute** - The time at which the scans start.
* **Recurrence** - The recurrence of the scan:
  * **None** - The scan runs only once.
  * **Daily** - The scan runs every day.
  * **Weekly** - The scan runs every week on the day of the week of the **Start date**.
  * **Monthly/Quarterly** - The scan runs every month or quarter on the day (number) defined in the **Start date**. In this case, a **Repeat scans every** checkbox is displayed to configure a different day:
    * In the first dropdown, choose the week of the month (First, Second, Third, Fourth, or Last).
    * In the second dropdown, choose the day of the week (Monday, Tuesday, and so on).
* **Override the target settings** - Define the target settings to use:
  * If not checked, the scheduled scan will use the settings defined for each target.
  * If checked, choose the settings the scheduled scan will use for all targets. Only the settings that are common to all targets can be configured.

This option is only available if all targets are of the same type: either Web or API.

After completing the form, click **Add** to create the scheduled scan.

## Manage scheduled scans

There are two areas in Snyk API & Web to list and manage scheduled scans:

* **Global to the account** - List and manage scheduled scans for all the targets of the account. Go to the **Scans** section and select the **Scheduled** tab.
* **A specific target** - List and manage scheduled scans for a specific target. Go to the **Targets** section, click the target to show its details, and click the **Scheduled Scans** tab.
