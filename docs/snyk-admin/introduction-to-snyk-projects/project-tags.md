# Project tags

{% hint style="warning" %}
**Release status**&#x20;

Project tags are available only for Enterprise plans.

See [Pricing plans](https://snyk.io/plans).
{% endhint %}

## Project tags use and conditions

### Using tags in the Web UI

{% hint style="info" %}
Group and Organization admins can perform the actions explained on this page. Collaborators can perform the actions if they are in an Organization that is part of the Group.
{% endhint %}

The Project tags feature allows you to add custom metadata to Snyk Projects. You can create and delete tags, apply tags to and remove tags from Projects, and filter Projects using tags.

You can perform these actions in the Snyk Web UI, as shown on this page.

You can [Add a tag to a Project](https://snyk.docs.apiary.io/#reference/projects/project-tags/add-a-tag-to-a-project) and  [Remove a tag from a Project ](https://snyk.docs.apiary.io/#reference/projects/remove-project-tag/remove-a-tag-from-a-project)in the Snyk API v1. You can also [List all tags in a Group](https://snyk.docs.apiary.io/#reference/groups/list-members-in-a-group) using the Snyk API v1.

You can set values for a tag applied to a Project and clear the Project tags using the Snyk CLI option `--project-tags`. See the [CLI commands and options summary](../../snyk-cli/cli-commands-and-options-summary.md) for the commands that support this option.

{% hint style="info" %}
If a Project tag is no longer assigned to a Project, the tag will no longer exist.&#x20;
{% endhint %}

### Tag conditions

The following conditions apply to Project tags:

* Keys are limited to 30 characters.
* Values are limited to 256 characters.
* Keys allow only alphanumerics and the following characters **`-`**, **`_`**
* Values allow these characters plus **`/`**, **`:`**, **`?`**, **`#`**, **`@`**, **`&`**, **`+`**, **`=`**, **`%`**, **`~`**
* Reusing a key and value combination does not add to the count.

## **How to create tags**

To create a new tag for a Snyk Project:

1.  On the Project details page, under **TAGS,** click **Add a key/value...**

    ***

    <figure><img src="../../.gitbook/assets/projects-tags_20sept2022.png" alt="Choose the option to create a Project tag"><figcaption><p>Choose the option to create a Project tag</p></figcaption></figure>
2. Add the new key and click **Enter**.
3. Add a new value and click **Enter**.

You have created a new tag. When a new tag is created, it is automatically applied to the Project it was created in. The tag can also be used for any other Project in the Group.

<figure><img src="../../.gitbook/assets/screenshot_2020-09-29_at_17.58.47.png" alt="Project details page showing tag applied"><figcaption><p>Project details page showing tag applied</p></figcaption></figure>

You can apply multiple Project tag values to the same Project tag key.

<figure><img src="../../.gitbook/assets/screenshot_2020-09-29_at_18.04.30.png" alt="Multiple Project tag values applied to the same key"><figcaption><p>Multiple Project tag values applied to the same key</p></figcaption></figure>

After you create a tag, it can be applied to other Projects in the Snyk Group.

## How to delete tags

To delete a Project tag from the Group, use the  [Delete a tag from a Group](https://snyk.docs.apiary.io/#reference/groups/delete-tag-from-group/delete-tag-from-group) endpoint in the Snyk API v1. You can also [List all tags in a Group](https://snyk.docs.apiary.io/#reference/groups/list-members-in-a-group) using the Snyk API v1.

The [Delete a tag from a Group](https://snyk.docs.apiary.io/#reference/groups/delete-tag-from-group/delete-tag-from-group) endpoint has the option in the Body to specify `"force": false` or `"force": true`_._ If you specify `"force": true`, the tag will be removed from any Projects to which it is applied, and it will then be deleted. If you specify `"force": false` and the tag is still applied to any Projects, error 403 “the tag has entities” occurs. Otherwise, tag deletion should succeed.&#x20;

{% hint style="info" %}
A tag will only exist if it has been created and applied to a Project or Projects. If a tag is

removed from all Projects, it will no longer exist.
{% endhint %}

## **How to apply and remove tags**

If a tag exists in your Group, you can apply it to any Snyk Project.

1. On the Project details page, click the **+** icon under **TAGS**.
2. You can either select a key from the list of recently used keys or type out the key for the tag you want to apply to the Project.
3. You can either select a value from the list of recently used values or type out the key value you want to apply to the Project.
4. After you select the value, the tag is applied to your Project.
5. To remove a tag from a Project, click the **x** for the tag.\
   If you remove a tag and it has not been assigned to any other Project, the tag will no longer exist.

<figure><img src="../../.gitbook/assets/screenshot_2020-09-29_at_18.14.44.png" alt="Select a key value for a tag"><figcaption><p>Select a key value for a tag</p></figcaption></figure>

## How to filter the Projects listing by tags

When **Group by none** (ungrouped) is applied to the Projects listing page, you can filter the list by tags using the option in the menu.

There is a filter by tag autocomplete. This is intentionally limited to a small number of results. If your tag is not displayed, enter more characters for the tag until it rises to the top of the results.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-24 at 08.23.14.png" alt="Filter by tag option"><figcaption><p>Filter by tag option</p></figcaption></figure>
