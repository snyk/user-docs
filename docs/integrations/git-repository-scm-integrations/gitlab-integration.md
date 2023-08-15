# GitLab integration

{% hint style="info" %}
**Feature availability**\
The Snyk GitLab integration is available for [Snyk Enterprise plan](https://snyk.io/plans/) customers. [Snyk Broker](../../enterprise-setup/snyk-broker/) is required if you integrate from a private network.
{% endhint %}

## Prerequisites for GitLab integration

* GitLab versions 9.5 and above (API v4).
* A public or private GitLab Group or Project.

## Snyk GitLab integration features

The Snyk GitLab integration allows you to:

1. Check for vulnerabilities in your pull requests.&#x20;
2. Receive email alerts when new vulnerabilities that affect your repository arise and fixes for those vulnerabilities are shown.
3. Receive email alerts containing a new pull request if a new upgrade or patch is available for a vulnerability.
4. From the **Report** page or the **Project** page on the Snyk Web UI, trigger a Snyk pull request for the fixes listed.

{% hint style="info" %}
**GitLab webhooks** send out an event to Snyk when merge requests occur. This starts a series of other events, such as pulling Project files, running the test process, and posting the results to GitLab, all of which occur on the Snyk side.
{% endhint %}

## GitLab access tokens

To set up the GitLab integration with Snyk, create a GitLab access token and enter this into the Snyk application.&#x20;

Typically, the first user in a Snyk Organization, a Snyk admin account user who is also a GitLab Owner or Maintainer, sets up an integration with a **GitLab Personal Access Token** or **Group Access Token.** This token is then authenticated with GitLab, enabling access by Snyk to the repositories in that GitLab account.

* A **GitLab Personal Access Token** is used to perform actions on and manage personal GitLab projects individually. These differ from Group Access Tokens as they are attached to a user rather than a GitLab group.
* A **GitLab Group Access Token,** most commonly used for the Snyk GitLab integration, is used to perform actions for and manage, more than one GitLab project within a GitLab group.

To trigger the creation of fix pull requests manually, all users in a Snyk Organization can add and work with any related Snyk Projects, while the merge requests themselves will appear in GitLab as having been opened by the Snyk admin who set up the configuration.

## How to set up the Snyk GitLab integration

### Add a GitLab Personal Access Token in GitLab

1.  Generate a GitLab Personal Access Token in your GitLab instance.\
    Select your profile icon, then **Edit Profile > Access Tokens**.\
    Set the token name, for example, Snyk, and select the **api** scope.\
    Alternatively, use a [Group Access Token](https://docs.gitlab.com/ee/user/group/settings/group\_access\_tokens.html) to grant access to all GitLab projects in a GitLab group or subgroup without contributing to GitLab's licensed user count.\


    <figure><img src="../../.gitbook/assets/2023-08-01_10-31-25.png" alt=""><figcaption><p>Create a GitLab Personal Access Token with the api scope.</p></figcaption></figure>
2.  Navigate to the Snyk [**Integrations**](https://app.snyk.io/integrations) page, select the GitLab integration tile, and enter the URL of your GitLab instance and the token you generated.\


    <figure><img src="../../.gitbook/assets/2023-08-01_10-19-59.png" alt=""><figcaption><p>Add the URL of your GitLab instance and the generated Personal access token.</p></figcaption></figure>
3. Click **Save**.
4.  When the tile on the **Integrations** page indicates the integration is **Configured**, click the tile and select the GitLab projects you want to test or select **Add projects** from your **Snyk** **Dashboard**.\


    <div align="left">

    <figure><img src="../../.gitbook/assets/2023-07-31_14-27-14.png" alt="" width="241"><figcaption><p>The Snyk GitLab integration tile shows as Configured.</p></figcaption></figure>

    </div>

### Add a GitLab Group Access Token

Generating a GitLab Group Access Token has a different methodology, explained [below](gitlab-integration.md#create-a-group-access-token), which requires selecting the Maintainer role for access.

Selecting the **api** scope with a Maintainer role allows:

* Snyk to authenticate user accounts and create webhooks, enabling automation of  fix pull requests and Snyk tests on your pull requests.
* Continuous write access, enabling Snyk Organization users to trigger the creation of fix pull requests manually.
* Continuous read access, enabling Snyk to monitor your Projects and enabling you and other members of the Organization to manually re-trigger tests.

#### Create a GitLab Group Access Token

1. Locate your GitLab Group and select **Settings > Access Tokens**.
2. Enter a descriptive token name such as "Snyk token", select the **Maintainer** role, and check the **api** scope**.**

<figure><img src="../../.gitbook/assets/gitlab_group_token.png" alt="Generate GitLab group access token"><figcaption><p>Generate a GitLab group access token</p></figcaption></figure>

#### Add a GitLab Group Access Token to Snyk

1. Copy the token you generated from GitLab.
2. Navigate to the Snyk GitLab integration page by selecting the tile.
3.  Paste your GitLab Group Access Token into the Snyk application field, the same way you would add a GitLab Personal Access Token.\


    <figure><img src="../../.gitbook/assets/2023-08-01_10-19-59.png" alt=""><figcaption><p>Enter your GitLab Group Access Token into the Snyk application Personal access token field.</p></figcaption></figure>

## **Fix vulnerabilities with Snyk merge requests in GitLab**

When viewing a Snyk test report for a Project that you own or when looking at a Project that you are watching with Snyk, you see two options for fixing a vulnerability:

* **Open a fix Merge Request:** generate a Snyk merge request with the minimal changes needed to fix the vulnerabilities affecting the Project.
* **Fix this vulnerability:** generate a Snyk merge request that fixes only this vulnerability.

You can review the vulnerabilities that will be fixed, change your selection, and choose to ignore any vulnerabilities that cannot be fixed now before opening the merge request on the **Open a fix Merge Request** page.

<figure><img src="../../.gitbook/assets/uuid-8d2ef9cb-cd32-bf48-a827-32bb358a10ab-en.png" alt="Open a fix Merge Request"><figcaption><p>Open a fix Merge Request</p></figcaption></figure>

When you open a merge request via snyk.io, Snyk notifies you. An example follows:

<figure><img src="../../.gitbook/assets/uuid-5e9a4b58-4d87-06fb-0479-a308515d4b12-en.png" alt="Notification of merge request"><figcaption><p>Notification of merge request</p></figcaption></figure>

### Get a Snyk merge request for newly disclosed vulnerabilities that affect you

Whenever a vulnerability is disclosed that affects a Project you are watching, Snyk will email you about it and generate a Snyk merge request that addresses the vulnerabilities. You will receive a merge request similar to the preceding example.

### Get a Snyk merge request when new upgrades or patches are available

When no upgrade is available, you can ignore or patch the vulnerability; patching is available only for Node.js Projects. When a fix option has become available, for example, an upgrade for a vulnerability you previously ignored, Snyk notifies you about this via email and also generates a merge request with the new fix.

## How to disable the GitLab integration

{% hint style="warning" %}
Disabling a GitLab integration effectively removes all Snyk integrations and webhooks, along with the Snyk credentials, and deactivates the GitLab Projects in the Snyk Web UI.
{% endhint %}

The Project will be set to inactive, and you will no longer get alerts, pull requests, or Snyk tests on your pull requests. Again, the webhook that enables the integration for this repo will be removed.

You can restart watching at any time; however, re-initiating GitLab Projects for monitoring requires setting up the integration again.
