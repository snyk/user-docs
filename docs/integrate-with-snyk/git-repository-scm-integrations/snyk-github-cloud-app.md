# Snyk GitHub Cloud App

{% hint style="warning" %}
The Snyk GitHub Cloud App is in [Open Beta](../../more-info/snyk-feature-release-process.md#open-beta). This feature must be enabled using a feature flag on the Group or Organization level. Contact [Snyk support](https://support.snyk.io/hc/en-us/requests/new) if you are interested in getting access.
{% endhint %}

{% hint style="info" %}
**Feature availability**

The GitHub Cloud App is available for all customers on GitHub Cloud, independent of plan type.

This feature is not yet supported for Snyk Broker.
{% endhint %}

## Prerequisites for GitHub Cloud App

* Snyk Organization Admin user role.
* GitHub Admin permissions for the selected repository. If the Snyk GitHub Cloud App is installed in a repository that requires organization permissions, you must be a GitHub organization owner.
* A public or private GitHub repository.

## Known limitations of the GitHub Cloud App

* Only one-to-one mapping between Snyk Organizations and GitHub organizations is supported.
* The GitHub Cloud App is only available in prod16.

## GitHub Cloud App benefits

The Snyk GitHub Cloud App improves upon many features as compared to the current GitHub integration, including role-based, granular access control, increased API rate limits, and the creation of an entry point for expanded and enhanced developer experiences.

* **RBAC (Role-Based Access Control) Compliance**: With the GitHub Cloud App, the access control mechanism is decoupled from individual user accounts. Instead, it is associated with the app entity itself. This separation allows for better management and enforcement of RBAC policies, as access control is handled at the application level rather than being tied to individual user accounts.
* **Granular access control**: The GitHub Cloud App allows for fine-grained control over access permissions at the repository level.&#x20;
* **Increased API rate limit**: The GitHub Cloud App provides higher rate limits, allowing Snyk to make a larger number of API requests. This increased limit will assist in handling large-scale use cases, such as mono-repos with a large number of Projects, GitHub Organizations with a large number of repositories, and more.
* **Enabler for an enhanced developer experience:**
  * Pull request checks: The GitHub Cloud App improves the PR Checks workflow by showing the test results directly on the Checks tab in GitHub.
  * Fix and upgrade pull requests: Pull requests initiated by Snyk would be performed on behalf of Snyk-bot rather than a service account individual user account.

## How to set up the Snyk GitHub Cloud App

1. Log in to your Snyk account and navigate to the Integrations section in the Snyk Organization where you would like to set up the GitHub Cloud App.
2.  Select the **GitHub Cloud App** tile.\


    <figure><img src="../../.gitbook/assets/2023-11-28_09-42-28 (1).png" alt=""><figcaption><p>Selection of the GitHub Cloud App tile</p></figcaption></figure>
3.  In the confirmation modal, select **Configure GitHub Cloud App.**\


    <figure><img src="../../.gitbook/assets/2023-11-28_09-44-21.png" alt="" width="375"><figcaption><p>Configuration notice for the GitHub Cloud App</p></figcaption></figure>
4.  You are taken to GitHub, where you can select the organization where you wish to install the app.\


    <figure><img src="../../.gitbook/assets/2023-11-28_09-45-45.png" alt="" width="375"><figcaption><p>Selection of the GitHub Organization to install the app into</p></figcaption></figure>
5. The next screen asks if you wish to install the app in all of the repositories belonging to the selected GitHub organization, or if you want to install the app in a select number of repositories belonging to a GitHub organization.&#x20;

<figure><img src="https://lh7-us.googleusercontent.com/izrSkGKUWpJYqBk4yOi4psfRqmNLJiH1LCun3RLwdIfdEUx8wmU5LomzYzvHCGf5Ak5WVAatbOYhDd489QCmSjJv58lYnizUnfH6HiMiI7xi5o0VfLHyDzCIMO5MdqNXxlOPgTR4pIWD6fhHrPEpC8o" alt="Install and Authorize settings for the GitHub organization you are installing the GitHub Cloud App into" width="375"><figcaption><p>Install and Authorize settings for the GitHub organization you are installing the GitHub Cloud App into</p></figcaption></figure>

{% hint style="warning" %}
The GitHub Cloud App will lose access to Snyk if it is uninstalled from the GitHub organization or if the repositories to which the app instance has access are edited.
{% endhint %}

## Feedback on the Snyk GitHub Cloud App

This feature is currently in Open Beta, so the functionality will likely evolve based on your feedback. If you would like to provide any feedback, contact your Account Manager or [Snyk support](https://support.snyk.io/hc/en-us/requests/new).



