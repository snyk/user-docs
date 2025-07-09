# Project collections groupings

{% hint style="info" %}
**Feature availability**\
This feature is available only with Enterprise Plans. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

There are two ways to group your Projects: **Project views** and **Project collections**.&#x20;

**Project views** are flexible collections you can create based on the filters available on the Project listing page.

In the Snyk Web UI, you can:

* Create and configure a dynamic Project view that is based on filters
* Share links to views of the Project listing page that you use with another member of your Organization
* View reports on your Project views

{% hint style="info" %}
Any Organization-level role can create and share Project views.

You can create and configure a Project view only when you apply the **Group by none** grouping in the Project listing page.

In addition, more filter options are available when you apply **Group by none**.

Views are visible to and can be edited by anyone in your Organization.
{% endhint %}

For details, see the [Project views](project-views.md) page.

**Project collections** help you collect and organize your Projects so you can easily view and perform actions on them.

Project collections are static collections of Projects which can be built using a combination of filters and selecting individual Projects. A Project collection can consist of multiple Project types from different Targets.

Project collections live on the Organization level, which means they are available for all members of the Organization to view and work within.

{% hint style="info" %}
Project views can be created by any user in an Organization, while Project collections can be created by users with admin roles and custom roles with the appropriate permissions. The permissions that a custom role requires to be able to use collections are:

* Read
* Create
* Delete&#x20;
* Edit

For more information, see [Pre-defined roles](../../user-roles/pre-defined-roles.md).
{% endhint %}

In the Snyk Web UI and [API](../../../snyk-api/reference/collection.md), you can:&#x20;

* Create collections using filters and selecting Projects individually
* View all of the collections that have been created on a collections listing page
* Perform bulk actions such as deleting, deactivating, and activating Projects in a collection. Performing bulk actions is not exclusive to collections, and can be done on Projects in the Project listing page.

For details, see the [Project collections](project-collections.md) page.
