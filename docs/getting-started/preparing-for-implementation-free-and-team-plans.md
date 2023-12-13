# Preparing for implementation: Free and Team plans

{% hint style="info" %}
See [Getting started with Snyk: Free and Team plans](getting-started-with-snyk-free-team-plan.md) before you read this guide.
{% endhint %}

These pages explain a few planning steps to follow before you implement Snyk beyond your local Project.

You have several decisions to make before you implement Snyk.

* Where will you implement Snyk in your Software Development Life Cycle (SDLC)?
* How will you structure your Snyk account?
* How will you and your team log in to Snyk?
* How will you scan your applications?
* How will you configure your Organization?

After you make these decisions and configure your Snyk Organization, you can [invite members](preparing-for-implementation-free-and-team-plans.md#invite-members).

## Where to implement Snyk in your SDLC

As you prepare for your team to start adopting Snyk as part of a secure development workflow, decide where to use Snyk in your SDLC, and what you want to scan, for example, application code, open-source library dependencies, container registries,  and so on.

You will want to roll Snyk out in phases depending on how far you and your team have progressed with developer security: awareness, visibility, preventing issues, fixing the backlog, and optimizing.

See the [Ways to integrate Snyk at your company](https://learn.dev.snyk.io/lesson/integrate-snyk-at-your-company/) course for more details.

## How to structure your Snyk account

For Free and Team plan accounts, the hierarchy for your Snyk account includes a single Organization and Projects within that Organization. Your team members have access to all Projects in your Organization.

{% hint style="info" %}
If you do not wish to keep the default Organization name, which is based on your signup name, you can rename the Organization. See [Manage groups and organizations](../snyk-admin/manage-groups-and-organizations/) for details.
{% endhint %}

## How you and your team log in to Snyk

There are a few different ways that users can authenticate into their Snyk accounts, such as with a GitHub or Google account.

<figure><img src="https://lh4.googleusercontent.com/snMHeJzIlnECB82n6BUAr0ssYW9iIfNdDNzvgSqqNOVOjr84x6C0CSijMuXefU5HXEzT1AVaKU-KjccG2s0qvoIdzbvcYUSvEfUIZf9o5X_fjswdW56YYujMEX6A0Jdl_OzwYsWgRyyIdVDq8qVV3lM" alt="Authenticate your account"><figcaption><p>Authenticate your account</p></figcaption></figure>

{% hint style="warning" %}
The option to use single sign-on (SSO) via your existing identity provider to sign in is available only on Enterprise plans.
{% endhint %}

## How you will scan your applications

Snyk Projects are the components that Snyk tests, along with the related configuration and metadata. Each target you want to scan (repository, container images, Dockerfiles, configuration files, source code) may include more than one Project. See [Snyk Projects](../snyk-admin/snyk-projects/) for details.

There are different ways to scan applications in Snyk, including from a [Git repository integration](walkthrough-code-repository-projects/), using the [Snyk CLI](../snyk-cli/), or the [Snyk API](../snyk-api/).

Before you start scanning more applications, ensure your Snyk Organization s configured appropriately.

## Configure your Organization

The Snyk Organization is the entity in your Snyk account that holds Snyk Projects and controls how Snyk scans those Projects. Before you start scanning additional applications and inviting members, you should configure your Organization settings to ensure the best experience.

### Configure Git repository integrations

Snyk includes a number of automations for Snyk Open Source when integrated with a source code manager (SCM) on a Git repository. These automations are a great way to enhance your developer security program. However, automation can introduce frustration for developers if introduced too early in your journey. Ensure your settings align with your phase of adoption.

You may decide to disable Snyk test for pull requests and other automations until your team is ready.

To disable or customize automation, navigate to **Settings > Integrations**, and select the appropriate code repository to open the configuration for that integration.

Choose the settings under **Default Snyk test for pull requests**:

* Open Source Security & Licenses
* Automatic fix pull requests
* Manually fix pull requests
* Automatically update Dockerfile base images
* Automatic dependency upgrade pull requests
* Pull request assignees for private repos
* Auto-detect Dockerfiles

See the [Source Code Manager Configurations](https://learn.dev.snyk.io/lesson/configure-snyk-scm/) course for details on how to match these integration settings with your security maturity phase.

See also [Integrate with Snyk](../integrate-with-snyk/) for more details.

### Define your default license policy

As part of Team and Enterprise plans, Snyk can identify license compliance issues in your Open Source packages. The default license policy specifies the severity associated with the use of different licenses, along with an option to provide license instructions if Snyk finds these licenses in your Projects. You can customize the default license policy for your needs.

See also [Snyk License Compliance Management](../scan-using-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md) for more details.

### Enable Snyk Code

Snyk Code enables your developers to scan their application source code for issues.

Free plan users and users who have purchased Snyk Code as part of a paid plan are prompted to allow source code scanning when first adding a code repository. You can enable it later or confirm it is available.

Navigate to **Settings > Snyk Code**. Select Enabled and Save changes.

{% hint style="warning" %}
If you have imported Projects prior to enabling Snyk Code, you must re-import them.
{% endhint %}

See [Activate Snyk Code using the Web UI](../scan-using-snyk/start-scanning-using-the-cli-web-ui-or-api/scan-code/activate-snyk-code-using-the-web-ui.md) for more details.

### Configure notifications

Snyk sends teams different types of alerts based on settings defined for the Organization. It is highly recommended that you define the default settings for the Organization with most notifications disabled by default before you scan additional applications.

If you want alerts to be sent by default for Projects imported into the Organization, you can have Snyk send notifications for vulnerabilities, license issues, or both. You can also limit the notifications to only high and critical severity issues.

{% embed url="https://thoughtindustries-1.wistia.com/medias/o41yfkfu34" %}
Organization notification video
{% endembed %}

Encourage individuals to set up their own notification preferences if they want to customize how they receive alerts for specific Projects.

{% embed url="https://thoughtindustries-1.wistia.com/medias/x1750wshlg" %}
Personal notification video
{% endembed %}

### Set ignore permissions

Snyk allows you to ignore issues in a few different ways. An ignored issue is not deleted. It is only removed from the filtered list of open issues.

Before inviting additional team members to Snyk, determine who can ignore the vulnerabilities and license issues that Snyk identifies.

Navigate to **Settings > General** to specify the ignore permissions.

{% hint style="info" %}
Snyk recommends that only Snyk Admins be allowed to ignore issues and that all ignores require a reason for the ignore.
{% endhint %}

### Define language settings

Based on the nuances of the tech stack you are using, you will want to set your language preferences. For more details on the nuances of using Snyk with specific circumstances, visit the pages specific to your language.

### Set up Jira integration

Integrate your Organization with Jira to assist with logging tickets and addressing backlogged security issues.

See also [Jira integration](../integrate-with-snyk/notification-and-ticketing-systems-integrations/jira-integration.md).

## Invite members

When your Organization is configured, you are ready to invite others to use Snyk. For Free and Team plans, you can invite members using their email addresses, asdemonstrated in the following video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/qqmkcaequj" %}
invite members video
{% endembed %}
