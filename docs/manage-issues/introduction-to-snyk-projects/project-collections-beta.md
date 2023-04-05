# Project collections (beta)

{% hint style="info" %}
**Feature availability**\
This feature is in Open Beta.\
Project collections are available for Enterprise plans. See [pricing plans](https://snyk.io/plans/) for details.
{% endhint %}

**Project collections** help you collect and organize your Projects so you can easily view and perform actions on them. A Project collection can consist of multiple Project types from different Targets.

**Project views** are flexible collections you can create based on the filters available on the Project listing page.

In the Snyk Web UI, you can:

* Create and configure a dynamic Project view that is based on filters
* Share a Project view with another member of your organization
* View reports on your Project views

{% hint style="info" %}
Any Oganization-level role can create and share Project views.

**Note:** You can only create and configure a Project collection when you apply the "Group by none" grouping in the Project listing page.

In addition, you will have more filter options available when the “Group by none” grouping is applied.
{% endhint %}

## Create and configure a Project view

1. On the right-hand side of the Project listing page beneath the **View Import Log** button, select **Group by none** from the drop-down list.
2. Click the drop-down next to the **FILTERS** drop-down, and click **Create new view**. If you already have filters selected, you can click **Save Changes** next to the drop-down and follow the next step.

<figure><img src="../../.gitbook/assets/image (393).png" alt="Create a new project view"><figcaption><p>Create a new Project view</p></figcaption></figure>

3. Enter a name for your Project view in the available field, then click **Create view.**

<figure><img src="../../.gitbook/assets/image (1) (7).png" alt="Enter a name for your Project view"><figcaption><p>Enter a name for your Project view</p></figcaption></figure>

Your new view is created and is automatically selected. You can click the icon next to your Project view to duplicate, rename, or delete it.

4. In the **FILTERS** drop-down, select the criteria that you would like to use to organize your view.

<figure><img src="../../.gitbook/assets/image (2) (4).png" alt="Select the filters that will be applied in your Project view"><figcaption><p>Select the filters that will be applied in your Project view</p></figcaption></figure>

5. After you have applied all of the filters that you want, click **Save changes** next to the drop-down. Your Project view will be updated.
6. If you want to configure a Project view, you can select different filters and click **Save changes** again. You can also discard any changes to a Project view before you save them to return it to its most recent saved state.

## Share a Project view with another member of your Organization

1. From the Project views drop-down, select the Project view that you would like to share with another member of your Organization.

<figure><img src="../../.gitbook/assets/image (3) (1).png" alt="Select the Project view that you want to share"><figcaption><p>Select the Project view that you want to share</p></figcaption></figure>

2. Once the Project view has loaded on the screen, copy the URL of the page. This URL can be shared with anyone in your Organization.

When you open the link for the view, the Project listing page is organized by the filters applied through the Project view with that URL.

{% hint style="info" %}
The name of the view will not appear in the Project views drop-down, and other users cannot modify your Project view.
{% endhint %}

## View reports on your Projects

1. From the Project views drop-down, select the Project view for which you would like to view the report.
2. Next to the **FILTERS** drop-down, click the **See report for these projects** button.

<figure><img src="../../.gitbook/assets/image (5) (1).png" alt="Generate a report for your Project view by clicking the See report for these projects button"><figcaption><p>Generate a report for your Project view by clicking the See report for these projects button</p></figcaption></figure>

The **Reports** page loads and shows a report of the Projects in your Project view.

<figure><img src="../../.gitbook/assets/image (6) (7).png" alt="A report is generated automatically with the filters selected in your Project view"><figcaption><p>A report is generated automatically with the filters selected in your Project view</p></figcaption></figure>

{% hint style="info" %}
The scan results you see in the Project listing page will be reflected in the reporting page roughly one (1) hour after the scan is complete.
{% endhint %}
