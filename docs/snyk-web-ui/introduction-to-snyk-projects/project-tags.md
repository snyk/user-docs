# Project tags

{% hint style="info" %}
**Feature availability**\
Project tags are available for Business and Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

A tag is a key and value combination which allows you to add additional custom metadata to projects, through the project issues page.

* Keys are limited to 30 characters
* Values are limited to 50 characters
* Both keys and values allow only alphanumerics and the following characters `-`, `_`
* You can create 1000 unique key and value combinations per group, and apply 10 unique tags per project
* Reusing a key and value combination does not add to the count

## Tag actions

Using the Snyk UI and Snyk API, you can perform the following actions:

* Create tags for projects
* Assign and remove tags from projects
* Filter projects in the project listing page

The Group and Organization admins can perform all actions, while a Collaborator can perform the actions if they are in an Org which is part of a Group.

## **Creating tags**

To create a new tag for a project:

1. On the project issues page, click **Add a key/value...**
2. Enter your new key, and then click Enter.
3. Enter your new value, and then click Enter.

You have created a new tag. When a new tag is created, it is assigned automatically to the project it was created in. The tag is also usable on any other project within your group.

![](../../.gitbook/assets/screenshot\_2020-09-29\_at\_17.58.47.png)

You can also apply multiple values to the same key:

![](../../.gitbook/assets/screenshot\_2020-09-29\_at\_18.04.30.png)

## **Assigning and removing tags from a project**

If a tag exists in your group, you can apply it to any project.

1. Click the "+" icon in the tags section.
2. You can either select a key from the list of recently used keys, or type out the key which is part of the tag you would like to assign to your project.
3. You can either select a value from the list of recently used values, or type out the value which is part of the tag you would like to assign to your project.
4. After you select the value, the tag is assigned to your project.
5. To remove a tag from a project, click the "X" icon on the tag.

![](../../.gitbook/assets/screenshot\_2020-09-29\_at\_18.14.44.png)

## Filter projects by tags

![](<../../.gitbook/assets/image (6) (2) (1).png>)

{% hint style="success" %}
**Filter by tag autocomplete**: this field is intentionally limited to a small amount of results--if your tag does not show up initially, enter more characters of the tag until it rises to the top of the results.
{% endhint %}
