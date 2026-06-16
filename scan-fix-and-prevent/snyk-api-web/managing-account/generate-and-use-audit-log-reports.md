# Generate and use audit log reports

Generate audit log reports and understand their content.

Audit logs contain the record of user activity in Snyk API & Web, allowing the reproduction of the timeline of events originated by users, which can be useful for identifying and analyzing behaviors or issues and taking the necessary actions.

This article is divided into the following sections:

* Generate the audit log report.
* Understand the information in the audit log report.
* Mapping of object model values and business names.

## Generate the audit log report

To analyze the audit log, first, we need to generate the report:

1. Open the **Settings** dropdown menu on the bottom-left corner of the navigation bar and select **Audit Log**.

<figure><img src="../../../.gitbook/assets/managing-account-generate-and-use-audit-log-reports-settings-audit-log.png" alt="Settings menu with Audit Log option"></figure>

1. Fill in the criteria of the report. You can download the list of actions performed by a single user or by all users in a specific period. When you are done, click on **Export**.

<figure><img src="../../../.gitbook/assets/managing-account-generate-and-use-audit-log-reports-export.png" alt="Audit log export criteria and Export button"></figure>

## Understand the information in the audit log report

Essentially, events listed in the report describe which action was done, on which object, when, and by whom. In update actions, there is also information about the fields that changed, the old and new values.

The report is in CSV format and can be opened in any application that handles spreadsheets.

<figure><img src="../../../.gitbook/assets/managing-account-generate-and-use-audit-log-reports-csv.png" alt="Example audit log CSV file"></figure>

The events are sorted by date (starting with the oldest event), and all event details are displayed in columns:

* **user\_email** - Email of the user that originated the event.
* **date** - Date and time of the event, in ISO 8601 UTC format.
* **action** - Action of the event, that can be:
  * create - A new object was created.
  * update - An existing object was updated.
  * delete - An existing object was deleted.
* **object\_model** - Type of object. Visit the next section for more information.
* **object\_id** - Identifier of the object.
* **field** - In update actions, this shows the name of the property or parameter for which the value changed.
* **old\_value** - In update actions, this shows the value before the change.
* **new\_value** - In update actions, this shows the new value of the property or parameter.

## Mapping of object model values and business names

The following table contains the mapping of the **object\_model** values to the business names:

| Object\_model value | Business Name |
| --- | --- |
| vulndb \| scope | Target |
| domains \| domain | Domain |
| vulndb \| assessment | Scan |
| vulndb \| asset | Extra Host |
| vulndb \| asset sequences | Navigation Sequence |
| vulndb \| finding | Finding |
