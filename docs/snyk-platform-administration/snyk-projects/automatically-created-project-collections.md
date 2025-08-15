# Automatically created Project collections

{% hint style="info" %}
**Release status**

Automated Collections are in Early Access and available only with Enterprise plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

See also [Limitations of Automated Collections](automatically-created-project-collections.md#limitations-of-automated-collections). You can enable Automated Collections in your Organization settings.

Scanning a repository through an SCM integration and rescanning it using the Snyk CLI creates duplicate Targets within the Snyk Web UI with duplicate Projects and issues. These may not be exact duplicates.&#x20;

With the option to create Project collections automatically, Projects from these duplicate Targets will be detected and grouped automatically into a new collection. This helps identify duplicates and allows filtering and reporting on the issues of your preferred code-scanning method.

{% hint style="info" %}
Automated Collections takes up to an hour to create or delete a collection.
{% endhint %}

## Enable automatic Project collections

{% hint style="info" %}
Bitbucket Server Projects cannot be grouped into automatically created collections because of their dynamic repository URL schema.
{% endhint %}

Automated Collections can be turned on under the **Organization settings** by users with the appropriate permissions: **Organization Admins**, **Group Admins**, or users with permissions to edit **Organization settings**.

<figure><img src="../../.gitbook/assets/enable auto-collections.png" alt="Managing Automated Collections under Organization Settings"><figcaption><p>Managing Automated Collections under Organization Settings</p></figcaption></figure>

After Automated Collections are enabled, all the Organization's Projects will be analyzed, and Projects with the same repository URL will be grouped automatically into a collection. All subsequent Project imports will also update the list of automatically created collections; there is no need to refresh the list manually.&#x20;

## Turn off Automated Collections

Turning off Automated Collections will remove all the automatically created collections without impacting any of the grouped Projects. Only the automated collections are deleted, not their contents.

## Automatically versus manually created Project collections

An automatically created collection has the same filtering and reporting options as a manually created collection:

* You can filter an automatically created collection to see only the result of your preferred scanning method, thus hiding duplicates.
* You can create reports from an automatically created collection and track issues from just one of the scanning methods and ignore the rest.

Automatically created collections have no management options available:

* You cannot rename an automated collection; its name will reflect the repository URL of the matched Projects.
* You cannot add new Projects to an automated collection. Its content updates itself when new Projects with the same repo URL are imported.
* You cannot delete an automatically created collection. You can turn off the Automated Collections feature, which will remove all automatically created collections.

## Limitations of Automated Collections

* Automated Collections does not detect [SAST](../../discover-snyk/getting-started/glossary.md#sast) scans pushed to the Snyk Web UI using the `snyk code test --report --project-name="name"` command.
* This feature supports only GitHub, GitHub Enterprise, GitLab, Bitbucket Cloud, and Azure integration scans.
* This feature does not support Snyk Container.
