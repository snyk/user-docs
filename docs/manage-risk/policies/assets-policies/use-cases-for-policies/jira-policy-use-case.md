# Jira policy - Use case

{% hint style="warning" %}
Jira action is in Closed Beta and available with Snyk AppRisk Essentials and Snyk AppRisk Pro. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

You can use the Jira action to create tickets for specific assets.

## Create A Jira Policy

This use case demonstrates how to set up and create a ticket every time a new class A asset does not have Snyk security coverage.

To follow this example, you need to create four filters that find:

### Filter 1

Filter by assets that are of type Repository.

<figure><img src="https://lh7-us.googleusercontent.com/X8FfYhea96ROTPfbiIQq1YyJHoSllVHZ0eIx_2ZkNJ7uPytMGBArWCJETI7pp5F__Hk7MoomAEcarxKidtsuLnM4J38Ch-qDEc7pC5NNHA3o8noBS3O47i83uc8vPOq3yHW_Nzgb7aw_vIFZaTApXao" alt="Filter assets of type Repository"><figcaption><p>Filter assets of type Repository</p></figcaption></figure>

### Filter 2

Filter by assets that are Class A.

<figure><img src="https://lh7-us.googleusercontent.com/thd74IsJiqDkqXZFGaIk59eNn0iMrpdKRdOPNcNg6-01l10B_YKb6LXOMLX4Tj4_fpmJWQ6Clqq9OURl2zArKFpf62F8vEY_D6KOoobTO6zOfkUgYiCKlgOKRP0bEQj0vz9Aoz95yAc_Ccg4lbC-ftQ" alt="Filter assets that are classified as Class A"><figcaption><p>Filter assets that are classified as Class A</p></figcaption></figure>

### Filter 3

Filter by tags that include relevant programming languages. You can use the following language tags: Apex, ASP, C, C#, C++, CMake, Go, HTML, Java, JavaScript, Kotlin, PHP, Python, Ruby, Scala, Swift, TypeScript, VisualBasic, Handlebars, Makefile, Objective-C.

<figure><img src="https://lh7-us.googleusercontent.com/kLR5gD6IYoqQK5eYPAOTBf3NWFlzzTP5aVrPziJ9pYaZYlZXvFCPy60cTnJ5UursV8LGr6luBP9Ab5sVypKmN1zgW81ZOtNaPbQjwYTCpSNedKhUX_WxJULF6iBl5O-uxnwlmm5BgK7aQgAcAVphBy0" alt="Filter configuration for language tags"><figcaption><p>Filter configuration for language tags</p></figcaption></figure>

### Filter 4

You have to exclude Snyk Code or Snyk Open Source from the scan coverage.

<figure><img src="https://lh7-us.googleusercontent.com/JMmpQw6q8VGy9a9EcNLv4meNZr3BYZ3iPJ8dVU92WCWlpgR9-DSFMRLB_FsSF-72E79Fo5c7nWijcta-UcqBivcDycs3wt0OEsmIfwbgBktBMvn_TFUfiA2sCHaydzb6juE_zmTybzl7t_5a62KFZW8" alt="Filter for coverage control "><figcaption><p>Filter for coverage control </p></figcaption></figure>



After setting up the filter conditions, you need to select the **Jira action**.&#x20;

* Information for the **Project Key** and **Issue Type** fields will be pulled in as a list using the Snyk AppRisk Jira Integration. _**(link to Jira integration in AppRisk**_)
* Add a description in the **Summary** field.

<figure><img src="https://lh7-us.googleusercontent.com/bDI-bFWUeq6BnW_wBIrfeuXH6a-DsfmU7HaJUmo6Z-C2jO3MkzTEvdazqzPmk8MCb63MN1ksPTPpn3XTmEJC_kemnRzlqmGGpkpi5OcyjAc-UwgYfkpy4n8Guu9BQtM-4e8jTLoB0n6btq2qqR6IdrE" alt="Create Jira Ticket action"><figcaption><p>Create Jira Ticket action</p></figcaption></figure>

* Click the **Next** button.
* Select the reporter of the ticket from the **Reported** field.
* Assign the ticket by using the options from the **Assigned** filed.
* Select a date from the **Due Date** field.
* The rest of the fields are optional. After you add all the information, click the **Apply** button.

\


<figure><img src="../../../../.gitbook/assets/image (386).png" alt="Create a Jira Ticket action"><figcaption><p>Create a Jira Ticket action</p></figcaption></figure>

This is how your policy should look after all filters and actions are set.

<figure><img src="../../../../.gitbook/assets/Screenshot 2024-04-25 at 13.38.40.png" alt=""><figcaption></figcaption></figure>

When a ticket is created, a label of the Policy ID will be added to it.&#x20;

## Modify a Jira Policy

You can modify a Jira Policy that you previously created.&#x20;

* Navigate to the **Policies** view and select the **Assets** tab.
* Click your Jira Policy, and then click **Edit**.&#x20;
* After you finish all the changes, click the **Save** button.

<div data-full-width="false">

<figure><img src="../../../../.gitbook/assets/image (388).png" alt="" width="324"><figcaption></figcaption></figure>

</div>

### Modify the Jira Action

* For Jira, Issues already opened will be modified according to the changes made.&#x20;
  * For example, if they edit the description, the existing tickets’ description will also be modified.
* For New matches, the Jira policy should create new Jira Issues&#x20;
* If the user does not want to modify previous Jira Issues, they should create a new policy

### Modify the Filters

* Jira Issues that were already opened will stay the same&#x20;
* For New matches, the Jira policy should create new Jira Issues&#x20;
* Tickets that their Asset no longer match will remain the same opened without changes

{% hint style="info" %}
If the label of the Policy ID is removed from the Jira Ticket,  the ticket will not be modified.
{% endhint %}

## Delete a Jira Policy

You can delete a Jira Policy that you previously created.&#x20;

* Navigate to the **Policies** view and select the **Assets** tab.
* Hover over your Jira Policy, and then click **Delete**.&#x20;

{% hint style="info" %}
If a policy is deleted, any previously linked tickets will remain open and no longer be linked to the policy.
{% endhint %}

<figure><img src="../../../../.gitbook/assets/image (389).png" alt="" width="364"><figcaption></figcaption></figure>

{% hint style="warning" %}
To delete the remaining open tickets, filter the tickets in Jira by the policy label and use the bulk actions.&#x20;
{% endhint %}

## Troubleshoot

### Editing required fields error

If you add required fields from an Issue Type in Jira will cause an error in the Policy. You can prevent this from happening by following the next steps:

* Disable the policy by using the toggle option from the Policies view, Assets tab, and setting it as Disabled.
* Add a field by editing the Jira policy.
* Modify the Policy by following the steps from the [Modify a Jira Policy](jira-policy-use-case.md#modify-a-jira-policy) section.
* Re-enable the policy by using the toggle option from the Policies view, Assets tab, and setting it as Enabled.

### Removing fields error

If you remove fields from an Issue Type in Jira, it will cause an error in the Policy. You can prevent this from happening by following one of the next steps:

* Delete the existing Policy by following the steps from the [Delete a Jira Policy](jira-policy-use-case.md#delete-a-jira-policy) section.
* Add back the removed field from the Issue Type in Jira.
* Create a new policy by following the steps from the [Create a Jira Policy](jira-policy-use-case.md#create-a-jira-policy) section.
