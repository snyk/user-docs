# Generate and use audit log reports

Generate audit log reports and understand their content.

Audit logs record user activity in Snyk API & Web, which lets you reproduce the timeline of events that users originated. This is useful for identifying and analyzing behaviors or issues and taking the necessary actions.

This article is divided into the following sections:

* Generate the audit log report.
* Understand the information in the audit log report.
* Mapping of object model values and business names.

## Generate the audit log report

To analyze the audit log, first generate the report:

1. Open the **Settings** dropdown on the bottom-left corner of the navigation bar and select **Audit Log**.
2. Fill in the criteria of the report. You can download the list of actions performed by a single user or by all users in a specific period. When you are done, click **Export**.

## Understand the information in the audit log report

Events listed in the report describe which action was done, on which object, when, and by whom. Update actions also include information about the fields that changed and the old and new values.

The report is in CSV format and you can open it in any application that handles spreadsheets.

Snyk sorts the events by date, starting with the oldest event, and displays all event details in columns:

* **user\_email** - Email of the user that originated the event.
* **date** - Date and time of the event, in ISO 8601 UTC format.
* **action** - Action of the event, which can be:
  * create - Snyk created a new object.
  * update - Snyk updated an existing object.
  * delete - Snyk deleted an existing object.
* **object\_model** - Type of object. Visit the next section for more information.
* **object\_id** - Identifier of the object.
* **field** - In update actions, this shows the name of the property or parameter for which the value changed.
* **old\_value** - In update actions, this shows the value before the change.
* **new\_value** - In update actions, this shows the new value of the property or parameter.

## Mapping of object model values and business names

The following table contains the mapping of the **object\_model** values to the business names:

| Object\_model value       | Business Name       |
| ------------------------- | ------------------- |
| vulndb \| scope           | Target              |
| domains \| domain         | Domain              |
| vulndb \| assessment      | Scan                |
| vulndb \| asset           | Extra Host          |
| vulndb \| asset sequences | Navigation Sequence |
| vulndb \| finding         | Finding             |
