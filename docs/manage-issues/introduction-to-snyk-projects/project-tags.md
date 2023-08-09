# Project tags

{% hint style="info" %}
**Feature availability**\
Project tags are available for Enterprise plans. See [pricing plans](https://snyk.io/plans/) for details.
{% endhint %}

## Project tags use and conditions

A Project tag allows you to add custom metadata to Snyk Projects.

### Using tags in the Web UI

In the Snyk Web UI, you can:

* [Create tags](project-tags.md#create-tags)
* [Apply and remove tags from Projects](project-tags.md#apply-and-remove-tags)
* [Filter Projects using tags](project-tags.md#filter-projects-listing-by-tags)

{% hint style="info" %}
Group and Organization admins can perform all these actions. Collaborators can perform the actions if they are in an Organization which is part of the Group.
{% endhint %}

### Adding and removing tags

You can add and remove tags using the API and the CLI:&#x20;

* [Snyk API Project tags endpoints](https://snyk.docs.apiary.io/#reference/projects/project-tags/add-a-tag-to-a-project)
* The Snyk CLI option `--project-tags`. See the [CLI commands and options summary](../../snyk-cli/cli-reference.md). for the commands that support this option.

### Tag conditions

The following conditions apply to Project tags:

* Keys are limited to 30 characters.
* Values are limited to 256 characters.
* Keys allow only alphanumerics and the following characters **`-`**, **`_`**
* Values allow these characters plus **`/`**, **`:`**, **`?`**, **`#`**, **`@`**, **`&`**, **`+`**, **`=`**, **`%`**, **`~`**
* You can create 1,000 unique key and value combinations per group, and apply ten unique tags per Project.
* Reusing a key and value combination does not add to the count.

## **Create tags**

To create a new tag for a Snyk Project:

1.  On the Project details page, under **TAGS** click **Add a key/value...**\\

    ***

    <figure><img src="../../.gitbook/assets/projects-tags_20sept2022.png" alt="Choose the option to create a Project tag"><figcaption><p>Choose the option to create a Project tag</p></figcaption></figure>
2. Add the new key and click **Enter**.
3. Add a new value and click **Enter**.

You have created a new tag. When a new tag is created, it is automatically applied to the Project it was created in. The tag can also be used for any other Project in the Group.

<figure><img src="../../.gitbook/assets/screenshot_2020-09-29_at_17.58.47.png" alt="Project details page showing tag applied"><figcaption><p>Project details page showing tag applied</p></figcaption></figure>

You can apply multiple Project tag values to the same Project tag key.

<figure><img src="../../.gitbook/assets/screenshot_2020-09-29_at_18.04.30.png" alt="Multiple Project tag values applied to the same key"><figcaption><p>Multiple Project tag values applied to the same key</p></figcaption></figure>

## **Apply and remove tags**

If a tag exists in your Group, you can apply it to any Snyk Project.

1. Click the **+** icon under TAGS.
2. You can either select a key from the list of recently used keys or type out the key for the tag you want to apply to the Project.
3. You can either select a value from the list of recently used values or type out the key value you want to apply to the Project.
4. After you select the value, the tag is applied to your Project.
5. To remove a tag from a Project, click the **x** for the tag.

<figure><img src="../../.gitbook/assets/screenshot_2020-09-29_at_18.14.44.png" alt="Select a key value for a tag"><figcaption><p>Select a key value for a tag</p></figcaption></figure>

## Filter Projects listing by tags

When **Group by none** (ungrouped) is applied to the Projects listing page, you can filter the list by tags using the option in the menu.

There is a filter by tag autocomplete. This is intentionally limited to a small number of results. If your tag is not displayed, enter more characters for the tag until it rises to the top of the results.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-24 at 08.23.14.png" alt="Filter by tag option"><figcaption><p>Filter by tag option</p></figcaption></figure>
