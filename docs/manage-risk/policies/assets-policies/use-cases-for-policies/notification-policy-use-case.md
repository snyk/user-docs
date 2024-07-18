# Notification policy - Use case

You can use the **Send Email** and **Send Slack Message** actions to notify you about changes that take place on your assets.

This use case demonstrates how to set up and receive a notification every time a new class A asset does not have Snyk security coverage.

To follow this example, you need to create four filters that find:

* **Filter 1**: Assets that are of type Repository.

<figure><img src="../../../../.gitbook/assets/image (206).png" alt="Filter configuration for asset type" width="352"><figcaption><p>Filter configuration for asset type</p></figcaption></figure>

* **Filter 2**: assets that are Class A.

<figure><img src="../../../../.gitbook/assets/image (223).png" alt="Filter configuration for asset class" width="354"><figcaption><p>Filter configuration for asset class</p></figcaption></figure>

* **Filter 3**: tags that include relevant programming languages (like Apex, ASP, C, C#, C++, CMake, Go, HTML, Java, JavaScript, Kotlin, PHP, Python, Ruby, Scala, Swift, TypeScript, VisualBasic, Handlebars, Makefile, Objective-C).

<figure><img src="../../../../.gitbook/assets/image (236).png" alt="Filter configuration for tags" width="349"><figcaption><p>Filter configuration for tags</p></figcaption></figure>

* **Filter 4**: Do not have Snyk Open Code or Snyk Open Source scan coverage.

<figure><img src="../../../../.gitbook/assets/Coverage filter.png" alt="Filter configuration for coverage control" width="353"><figcaption><p>Filter configuration for coverage control</p></figcaption></figure>

After setting up the filter conditions, you need to choose the **Send Slack Message** action. You can generate the Slack webhook by using the  [Incoming WebHooks app](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks) or by [creating your own Slack app](https://api.slack.com/incoming-webhooks).

This is how your policy should look after all filters and actions are set.

<figure><img src="../../../../.gitbook/assets/image (289).png" alt="AppRisk - Setting up a Notification policy"><figcaption><p>Snyk AppRisk - Setting up a Notification policy</p></figcaption></figure>
