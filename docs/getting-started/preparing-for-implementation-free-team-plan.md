# Preparing for implementation: Free / Team plan

{% hint style="info" %}
Please read the [Getting started with Snyk: Free / Team plan](getting-started-with-snyk-free-team-plan.md) guide before you read this guide.
{% endhint %}

## Introduction

In this guide, we’ll look at a few planning steps before you implement Snyk beyond your local project.

You should make several decisions before you implement Snyk.

Ask yourself:

* Where will you implement Snyk in your Software Development Life Cycle (SDLC)?
* How will you structure your Snyk account?
* How will you and your team access Snyk?
* How will you scan your applications?
* How will you configure your organization?

After you make these decisions and configure your Snyk Organization, you can [invite members](preparing-for-implementation-free-team-plan.md#invite-members).

### Where to implement Snyk in your SDLC

As you prepare for your team to start adopting Snyk as part of a secure development workflow, you should decide where to use Snyk in your SDLC, and what you want to scan (such as application code, open-source library dependencies, container registries, etc.)

You'll also want to roll Snyk out in phases depending on how far you and your team have progressed with developer security (awareness, visibility, preventing issues, fixing the backlog, optimizing).

See the [Ways to integrate Snyk at your company](https://training.snyk.io/courses/ways-to-use-snyk) course in Snyk Training for more details.

### How to structure your Snyk account

For Free and Team plan accounts, the hierarchy for your Snyk account includes a single Organization and Projects within that Organization. Your team members can access all Projects in your Organization.

{% hint style="info" %}
If you do not wish to keep the default Organization name, which is based on your signup name, you can rename the Organization. See [Managing groups & organizations](../snyk-admin/manage-groups-and-organizations/) for details.
{% endhint %}

### How you and your team log in to Snyk

There are a few different ways that users can authenticate into their Snyk accounts, such as with a GitHub or Google account.

<figure><img src="https://lh4.googleusercontent.com/snMHeJzIlnECB82n6BUAr0ssYW9iIfNdDNzvgSqqNOVOjr84x6C0CSijMuXefU5HXEzT1AVaKU-KjccG2s0qvoIdzbvcYUSvEfUIZf9o5X_fjswdW56YYujMEX6A0Jdl_OzwYsWgRyyIdVDq8qVV3lM" alt="Authenticate your account"><figcaption><p>Authenticate your account</p></figcaption></figure>

{% hint style="warning" %}
The option to use single sign-on (SSO) via your existing identity provider to sign in is available only on Enterprise plans.
{% endhint %}

### How you’ll scan your applications

Snyk Projects are the components Snyk tests, along with the related configuration and metadata. Each target you want to scan (repos, container images, Dockerfiles, configuration files, source code) may include more than one Project. See [Introduction to Snyk Projects](../scan-application-code/snyk-code/snyk-code-local-engine.md) for more details.

There are different ways to scan applications in Snyk, including from a [Git repository integration](walkthrough-code-repository-projects/), using the [Snyk CLI](../snyk-cli/), or the [Snyk API](../snyk-api/).

Before you start scanning more applications, make sure your Organization in Snyk is configured appropriately.

### Configure your Organization

The Snyk Organization is the entity in your Snyk account that holds Snyk Projects, and controls how Snyk scans those Projects. Before you start scanning additional applications and inviting members, you should configure your Organization settings to ensure the best experience.

#### Configure Git repository integrations

Snyk includes a number of automations for Snyk Open Source when integrated with a source code manager (SCM) on a Git repository. These automations are a great way to mature your developer security program. However, the automations can introduce frustration for developers if introduced too early in your journey. Make sure your settings align with your phase of adoption.

You may decide to disable Snyk test for pull requests and other automations until your team is ready.

To disable or customize automations, go to **Settings > Integrations**, and select the appropriate code repository to open the configuration for that integration.

Choose the settings under **Default Snyk test for pull requests**:

* Open Source Security & Licenses
* Automatic fix pull requests
* Manually fix pull requests
* Automatically update Dockerfile base images
* Automatic dependency upgrade pull requests
* Pull request assignees for private repos
* Auto-detect Dockerfiles

See the[ Source Code Manager Configurations](https://training.snyk.io/courses/source-code-manager-configurations?query=Source%20code) training course for details of how to match these integration settings with your security maturity phase.

Also see [Snyk Integrations](https://docs.snyk.io/integrations) for more details.

#### Define default license policy

As part of Team and Enterprise plans, Snyk can identify license compliance issues in your Open Source packages. The Default License Policy indicates the severity associated with the use of different licenses, along with an option to provide license instructions if Snyk finds these licenses in your projects. You can customize the default license policy for your needs.

Also see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/products/snyk-open-source/licenses/getting-started-snyk-licensing-compliance) for more details.

#### Enable Snyk Code

Snyk Code enables your developers to scan their own application source code for issues.

Free plan users and users that have purchased Snyk Code as part of a paid plan are prompted to allow source code scanning when first adding a code repository. You can enable it later (or confirm it’s available).

Go to **Settings > Snyk Code**. Select Enabled and Save changes.

{% hint style="warning" %}
If you imported Projects prior to enabling Snyk Code, you will need to re-import them.
{% endhint %}

See [Getting started with Snyk Code](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code) for more details.

#### Configure notifications

Snyk sends teams different types of alerts based on settings defined for the Organization. It's highly recommended to define the default settings for the Organization with most notifications disabled by default before you scan additional applications.

If you want alerts to be sent by default for projects imported into the Organization, you can have Snyk send notifications for either vulnerabilities and/or license issues. You can also limit the notifications to only High and Critical severity issues.

{% embed url="https://thoughtindustries-1.wistia.com/medias/o41yfkfu34" %}
Organization notification video
{% endembed %}

Encourage individuals to set up their own notification preferences, if they want to customize how they receive alerts for specific projects.

{% embed url="https://thoughtindustries-1.wistia.com/medias/x1750wshlg" %}
Personal notification video
{% endembed %}

#### Set ignore permissions

Snyk allows you to ignore issues in a few different ways. An ignored issue is not deleted. It is only removed from the filtered list of open issues.

Before inviting additional team members to Snyk, determine who can ignore the vulnerabilities and license issues that Snyk identifies.

Go to **Settings > General** to specify the ignore permissions.

{% hint style="info" %}
Snyk recommends that only Snyk Admins are allowed to ignore issues, and that all ignores require a reason for the ignore.
{% endhint %}

#### Define language settings

Based on the nuances of the tech stack you are using, you'll want to also set your language preferences. For more detail on the nuances of using Snyk with specific circumstances, visit the Guide specific to your language.

#### Set up Jira integration

Integrate your Organization with Jira to assist with logging tickets and addressing backlogged security issues.

Also see [Jira](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/jira).

### Invite members

When your organization is configured, you are ready to invite others to use Snyk. For Free and Team plans, you can invite members using their email addresses, demonstrated in the following video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/qqmkcaequj" %}
invite members video
{% endembed %}
