# Notification policy

You can use the **Send Email** and **Send Slack Message** actions to notify you about changes that take place on your assets.

This use case demonstrates how to set up and receive a notification every time a new class A asset does not have Snyk security coverage.

To follow this example, you need to create four filters that find:

* **Filter 1**: Assets that are of type Repository.
* **Filter 2**: assets that are Class A.
* **Filter 3**: tags that include relevant programming languages (like Apex, ASP, C, C#, C++, CMake, Go, HTML, Java, JavaScript, Kotlin, PHP, Python, Ruby, Scala, Swift, TypeScript, VisualBasic, Handlebars, Makefile, Objective-C).
* **Filter 4**: Do not have Snyk Open Code or Snyk Open Source scan coverage.

After setting up the filter conditions, you need to choose the **Send Email** action.

{% hint style="info" %}
If you want to set a **Send Slack Message** action, then you can generate the Slack webhook by using the [Incoming WebHooks app](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks) or by [creating your own Slack app](https://api.slack.com/incoming-webhooks).
{% endhint %}

Customize the Send Email action to notify you with a link to the assets impacted by the set policy. You can do this by typing "/" inside the **Body** field of the **Send Email** action and selecting **Link to Assets**. After you save the policy, every notification received will list all the assets impacted by the policy.

<figure><img src="../../../../.gitbook/assets/image (183).png" alt="Snyk AppRisk - Set up the Links to Assets option from the Send Email action"><figcaption><p>Assets Policy - Set up the Links to Assets option from the Send Email action</p></figcaption></figure>

This is how your policy should look after all filters and actions are set.

<figure><img src="../../../../.gitbook/assets/image (182).png" alt="Snyk AppRisk - Setting up a Notification policy"><figcaption><p>Assets Policy - Setting up a notification policy</p></figcaption></figure>

You will receive an email notification after including the **Link to Assets** option in the Body field. You can access the assets from the notification individually, or view them in an aggregated form by clicking the **Click Here** link. The list of assets displayed in the email notification is automatically generated.

{% hint style="info" %}
After an email notification policy is created, it is run in a maximum of 3 hours after creation, then once every 3 hours.

If your policy is set to run daily, then the policy is run in a maximum of 3 hours after the 24-hour period ends. You can always manually run a policy by using the Run button.
{% endhint %}
