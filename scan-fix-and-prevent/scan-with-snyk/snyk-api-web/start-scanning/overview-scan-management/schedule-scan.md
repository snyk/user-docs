---
description: How to schedule a Snyk API and Web scan
nav_context: classic
---

# Schedule scan

You can set up a schedule so that scans start automatically on specific dates and times.

Creating a scheduled scan involves two steps:

1. Select the action to create scheduled scans.
1. Configure the scheduled scan by filling out a form.

## Select the action to schedule a scan

To start creating a scheduled scan for a single target, navigate to the **Targets** section and do one of the following:

* From the scan actions available on the page, click the caret icon within the **Scan** button, and select **Schedule scan**.
* Alternatively, navigate to the target details page (any tab), click the caret icon within the **Scan** button, and select **Schedule scan**.

To start creating a scheduled scan for multiple targets, navigate to the **Targets** section and do the following:

1. Select the check boxes for the targets in the list. If all selected targets have scans in the same state, the button for bulk scan actions becomes available.
1. Click the caret icon within the **Scan** button, and select **Schedule scan**.

## Configure the scheduled scan

Fill out the form to configure the scheduled scan:

* **Start date** - The date on which the scan starts.
* **Hour and Minute** - The time at which the scans start.
* **Recurrence** - The recurrence of the scan:
  * **None** - The scan runs only once.
  * **Daily** - The scan runs every day.
  * **Weekly** - The scan runs every week on the day of the week of the **Start date**.
  * **Monthly/Quarterly** - The scan runs every month or quarter on the day (number) defined in the **Start date**. In this case, Snyk displays a **Repeat scans every** check box to configure a different day:
    * In the first dropdown, select the week of the month (**First**, **Second**, **Third**, **Fourth**, or **Last**).
    * In the second dropdown, select the day of the week (**Monday**, **Tuesday**, and so on).
* **Override the target settings** - Define the target settings to use:
  * If not selected, the scheduled scan uses the settings defined for each target.
  * If selected, choose the settings the scheduled scan uses for all targets. You can configure only the settings that are common to all targets.

This option is available only if all targets are of the same type: either Web or API.

After completing the form, click **Add** to create the scheduled scan.

## Manage scheduled scans

There are two areas in Snyk API & Web to list and manage scheduled scans:

* **Global to the account** - List and manage scheduled scans for all the targets of the account. Navigate to the **Scans** section and select the **Scheduled** tab.
* **A specific target** - List and manage scheduled scans for a specific target. Navigate to the **Targets** section, click the target to show its details, and click the **Scheduled Scans** tab.
