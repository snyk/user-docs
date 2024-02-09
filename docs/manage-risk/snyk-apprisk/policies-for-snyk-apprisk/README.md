# Policies for Snyk AppRisk

## Overview

With Snyk AppRisk policies, you can easily automate the process of adding business context and receiving notifications. After a policy is created, it is run in a maximum of 30 minutes after creation, then once every 30 minutes. You can always manually run a policy by using the Run button.

## **Use Cases**

You can create policies for organizing the assets, classifying them, and always being up to date with the latest information about an asset.\
Common use cases for policies include:

* [New asset notifications](use-cases-for-policies/notification-policy-use-case.md)
* [Asset classification](use-cases-for-policies/classification-policy-use-case.md)
* [Asset tagging](use-cases-for-policies/tagging-policy-use-case.md)
* [Coverage control](use-cases-for-policies/coverage-control-policy-use-case.md)

### New asset notifications

&#x20;Notify members of the AppSec team when new assets meeting certain criteria are discovered. For example, you may send a Slack message to the infra team if new repository assets that leverage Terraform as a technology are detected by Snyk AppRisk.

### Asset classification

Classify repository assets according to their business criticality from A (most critical) to D (least critical), based on asset properties such as name and tags. For example, you might indicate that any repositories that contain “customer-portal” in the name should be classified as A, given that the customer-portal app holds sensitive data.

### Asset tagging

&#x20;Categorize and label repository assets with flexible tags, which can be used to filter the asset inventory.

### Security coverage

Monitor if your assets are scanned by the selected security products. You can select one or multiple security products and also specify a timeframe for when the last scan should have taken place.
