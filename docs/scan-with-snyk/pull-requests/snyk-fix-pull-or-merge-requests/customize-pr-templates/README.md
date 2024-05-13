# Customize PR templates

{% hint style="warning" %}
**Release status** \
Customize PR templates is currently in [Closed Beta](../../../../getting-started/snyk-release-process.md) and available only for Enterprise plans. The functionality is likely to evolve based on feedback, and there will be breaking changes.&#x20;

Contact your account manager to get access to this feature. After your account manager turns on the feature flag, you can activate this feature in the [Snyk Preview](../../../../snyk-admin/snyk-preview.md).
{% endhint %}

## Understand Customized PRs

When you fix or upgrade Snyk Open Source and Container Projects imported using the [SCM integrations](../../../../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/), Snyk raises pull requests (PRs) against your repository.&#x20;

Snyk has default templates for the title, description and commit message. These indicate what packages are being changed, which issues are being fixed, and many other details.

You may have your own standards and practices for submitting pull requests. For instance, if a pull request comes from Snyk, you may require the title to begin with `SNYK:`. This page identifies the areas of pull requests that you can customize and provides instructions on how to do so.

On the [supported languages, frameworks, and feature availability overview](../../../../scan-using-snyk/supported-languages-frameworks-and-feature-availability-overview/) page, you can find all the languages for which the Fix PR functionality is supported.

{% hint style="warning" %}
Snyk is looking initially for feedback on the variables and templating system. After the approach is validated, Snyk will look into building more robust authoring workflows using the API and UI interfaces.
{% endhint %}

## Enable the Customize PR feature

Follow these steps to enable the Configure Snyk Pull Requests feature:

1. Log in to your Snyk account.
2. Select **Settings**, then **Snyk Preview**.
3. Enable the **Configure Snyk Pull Requests** feature.\
   You can enable this feature at the Group level. See [Configure Automatic fix PRs](../create-automatic-fix-prs-for-backlog-issues-and-known-vulnerabilities.md) for more configuration details.

<figure><img src="../../../../.gitbook/assets/Config Snyk PRs.png" alt="Enable the Configure Snyk Pull Requests feature"><figcaption><p>Enable the Configure Snyk Pull Requests feature</p></figcaption></figure>

## **API support**

Customize your PR template automatically by making an API request; see [Create an API request with a Custom template](apply-a-custom-pr-template.md#create-an-api-request-with-a-custom-template). For more information, see the [Snyk API pull request templates endpoints](https://apidocs.snyk.io/?version=2023-10-13%7Ebeta#tag--Pull-Request-Templates).
