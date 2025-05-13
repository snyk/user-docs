# Assets policies

## Overview

With Policies, you can easily automate the process of adding business context and receiving notifications.

{% hint style="info" %}
After a policy is created, it is run in a maximum of 3 hours after creation, then once every 3 hours.

If your policy is set to run daily, then the policy is run 3 hours after the 24-hour period ends. You can always manually run a policy by using the Run button.
{% endhint %}

Access the Snyk Essentials policies by positioning yourself at the Group level, selecting **Policies**, then **Assets**.

The following video presents an overview of the types of policies you can create from the Policies view.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1737656954/snyk-learn/product-training-videos/Snyk_Essentials_and_Snyk_AppRisk_-5a_-_v1_-_Policy_Overview.mp4" %}
Overview of asset policies
{% endembed %}

{% hint style="info" %}
[Manage assets](../../../manage-assets/) and [assets policies](./) are interconnected. Before setting up any new policy, ensure you have reviewed and filtered your assets from the Inventory menu.
{% endhint %}

## Use Cases

You can create policies for organizing the assets, classifying them, and always being up to date with the latest information about an asset.\
Common use cases for policies include:

* [New asset notifications](use-cases-for-policies/notification-policy-use-case.md)
* [Asset classification](use-cases-for-policies/classification-policy-use-case.md)
* [Asset tagging](use-cases-for-policies/tagging-policy-use-case.md)
* [Coverage control](use-cases-for-policies/coverage-control-policy-use-case.md)

### New asset notifications

Notify members of the AppSec team when new assets meeting certain criteria are discovered. For example, you may send a Slack message to the infra team if new repository assets that leverage Terraform as a technology are detected by Snyk Essentials.

When setting up a notification action (email or Slack) for a policy, you can include a link to the relevant assets. Each notification will list all the assets impacted by the policy. You can view the assets individually, or you can see a summary of all the assets by clicking the **Click Here** option in the notification. The list of assets displayed in the email notification is automatically generated.

### Asset classification

Classify repository assets according to their business criticality from A (most critical) to D (least critical), based on asset properties such as name and tags. For example, you might indicate that any repositories that contain "customer-portal" in the name should be classified as A, given that the customer-portal app holds sensitive data.

### Asset tagging

Categorize and label repository assets with [asset tags](../../../manage-assets/assets-inventory-components.md#asset-tags) to filter the asset inventory.

* **GitHub custom properties** - lists the GitHub custom properties associated with your GitHub repository as a tag
* **User-defined tags** are customizable, as you can define their logic through [Assets Policies](./). For example, you can set tags to label a repository that comes from a specific source, such as GitHub. Tags associated with assets are identified in the UI with the **Asset policy tags** name.
* **System tags** are automatically assigned by Snyk based on asset names or detected keywords (for example, `codeowners`).

### Security coverage

Monitor if your assets are scanned by the selected security products. You can select one or multiple security products and also specify a timeframe for when the last scan should have taken place.
