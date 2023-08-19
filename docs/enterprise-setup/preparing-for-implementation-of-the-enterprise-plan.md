# Preparing for implementation of the Enterprise plan

{% hint style="info" %}
Before using this page, see [Getting started with the Snyk: Enterprise plan](getting-started-with-the-snyk-enterprise-plan.md) for information about getting a team involved in using Snyk.
{% endhint %}

This page looks at a few planning steps and provides a walkthrough of some initial configuration options to complete before inviting more teams and scanning additional applications.

{% hint style="info" %}
If you have been using Snyk on the Free or Team Plan and are looking for guidance on upgrading to the Enterprise plan, see [Upgrading to Enterprise Plan](upgrading-to-the-enterprise-plan.md)
{% endhint %}

## Considerations before launching Snyk to additional teams

There are several questions to consider before you launch Snyk to multiple teams in your Organization:

* Where will you implement Snyk in your Software Development Life Cycle (SDLC)?
* How will you structure your Snyk account?
* How will you and your team log in to Snyk?
* How will you define license and security policies?
* How will you scan your applications?
* How will you configure your first Organization?

After you answer these questions and configure your first Organization, you can configure additional Organizations and invite members.

## Checklist for preparing to implement the Enterprise plan

Use this checklist to ensure you have prepared for your implementation.

* [ ] Determine where to implement Snyk in your SDLC.
* [ ] Determine how you will structure your Snyk account: your Snyk Organization structure and naming convention and Projects in Organizations.
* [ ] Determine how you and your team will log in to Snyk.
* [ ] Set up SSO if applicable.
* [ ] Define license and security policies.
* [ ] Develop your Project import strategy: how you will scan your applications.
* [ ] Set up your first Organization.\
  Note: There are many Settings to establish, including notification defaults for the Group and for the first Organization.
* [ ] Invite members.
* [ ] Begin importing Projects
* [ ] Create additional Organizations.
* [ ] Roll out Snyk to your teams.
  * [ ] Identify influential developers as security champions and early adopters.
  * [ ] Define your communications and training plan.
  * [ ] Define success metrics based on program maturity.

## Where to implement Snyk in your SDLC

As you prepare for your teams to start adopting Snyk as part of a secure development workflow, you should decide where to use Snyk in your SDLC, and what you want to scan, for example, application code, open-source library dependencies, container registries, and so on.

You will also want to roll Snyk out in phases depending on how far you and your teams have progressed with developer security: awareness, visibility, preventing issues, fixing the backlog, optimizing) See the [Ways to integrate Snyk at your company](https://training.snyk.io/courses/ways-to-use-snyk) training course for more details.

## How to structure your Snyk account

For Enterprise plans, the hierarchy for your Snyk account includes a Group at the highest level, Organizations within that Group, and Projects within those Organizations.

<figure><img src="https://lh6.googleusercontent.com/yt7XP5scEUM8_rRMTZOYUvuUPOP23QPoFw551WLspjaU1tUtCvo8kOhPHJ-Z2tYIIoijMmcYraKqnLtDlQIzgeX5ZaeclHJWrGA9pcaprWK6hG898FV6BJ2Pzx4PUsRkewXv_F5M58ruiyVJu9pTnw" alt="Snyk account hierarchy"><figcaption><p>Snyk account hierarchy</p></figcaption></figure>

{% hint style="info" %}
For very large implementations of Snyk, you may also choose to implement multiple Groups, for example, to align top levels with separate business units. Contact your Account Executive if you are interested in discussing this option.
{% endhint %}

Consider this hierarchy, as well as which members of your teams need access to specific Snyk Projects, when you decide how to align your Snyk Organizations with your internal structure. You will want to decide how to structure your account before inviting members or scanning a large number of applications.

To make these decisions, consider the following:

* Which team members have access to specific repositories?
* How do you want to apply policies to Snyk Projects for prioritizing and automating tests?
* How do you want to report on scans?

For more help addressing these questions, see the free, self-paced training course [Snyk account structure](https://training.snyk.io/courses/snyk-account-structure).

## How your teams log in to Snyk

There are a few different ways that users can authenticate to log in to their Snyk accounts, such as with a GitHub or Google account.

<figure><img src="https://lh5.googleusercontent.com/EYz6tsF_VfAbsLnz4dZyPa_BH40wJw2tX4sfKhyCD7EakbIahU3KtajVtkoxcpzLrM3m--x-EXFUz--qLlIDj5dFBvSmGM6qSmteM3K-7r6vWPBndJT-blCz_EQq6KU-zaMk_KHKZzBXjpDw5260wg" alt="Logging in to Snyk"><figcaption><p>Logging in to Snyk</p></figcaption></figure>

## Enable Single Sign-On (SSO) if applicable

You may want to set up SSO using your existing identity provider to streamline sign-ins and new user provisioning.

If you use SSO, after you set it up, you must remove any social login accounts from Snyk so that all user access to your Group is controlled through SSO.

See [Set up Single Sign-on (SSO) for authentication](using-single-sign-on-sso-for-authentication/) for more details.

## Define license and security policies

The Policy manager allows you to define license and security policies at the Group level.

License policies define how your Organizations identify license compliance issues in your Open Source packages and change the severity of issues associated with use of those licenses.

See [Getting Started with Snyk License Compliance Management](../scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance.md) for more details.

Security policies define the way your Organizations handle specific issues or exploit maturity levels. A policy can programmatically increase an issue’s severity rating, decrease the severity rating, or ignore the issue.

See [Security policies](../manage-issues/policies/security-policies/) for more details.

## How you will scan your applications

Snyk Projects are the components that Snyk tests, along with the related configuration and metadata. Each target you want to scan: repos, container images, Dockerfiles, configuration files, source code, may include more than one Project.

See the [introduction to Snyk Projects](../manage-issues/snyk-projects/) for more details. Before you add Projects to Snyk, you wlll define a Project import strategy.

There are different ways to add Projects in Snyk, including using an integration, the Snyk CLI, or the Snyk API.

Make sure your Organizations in Snyk are configured appropriately. Determine how you want to configure your first Organization and complete all of the settings first. Then you can use that Organization as a template for configuring other Organizations so you can begin scanning applications.

## Set up your first Organization

Configure your first Organization before you start scanning applications and inviting members, to ensure the best experience.

To save time, Snyk recommends setting up your first Organization fully so that this Organization can serve as a template when you create other Organizations using the Snyk Web UI or  an API.

Follow these steps to ensure you have completed the configuration.

### Set ignore permissions

Snyk allows you to ignore issues in a few different ways. An ignored issue is not deleted. It is only removed from the filtered list of open issues.

Before inviting additional team members to Snyk, determine who can ignore the vulnerabilities and license issues that Snyk identifies.

Go to **Settings > General** to specify the ignore permissions.

{% hint style="info" %}
Snyk recommends that only Snyk Admins be allowed to ignore issues and that all ignores must have a reason specified.
{% endhint %}

### Configure Git repository integrations

Snyk includes a number of automated processes for Snyk Open Source when integrated with a source code manager (SCM) on a Git repository. These automated processes are a great way to develop the maturity of your developer security program. However, these automated processes can introduce frustration for developers if introduced too early in your journey. Make sure your settings align with your phase of adoption.

You may decide to disable `snyk test` for pull requests and other automated processes until your teams are ready.

Use the following steps to disable these automated processes for the Organization until you are ready to implement them. Note that all of these settings may not be displayed, depending on the specific Integration you are configuring.

1. Open the Organization in the Snyk UI.
2. Select **Integrations** from the left navigation pane.
3. Select the cog icon adjacent to integration you want to configure, for example, GitHub.
4. In the **Snyk test for pull requests** section, disable the following:
   1. **Open Source Security & Licenses**
   2. **Code Analysis**
   3. **Automatic fix pull requests**, both **New vulnerabilities** and **Known vulnerabilities (backlog)**
   4. **Automatically Update Dockerfile base images**
   5. **Automatic dependency upgrade pull requests**
   6. **Pull request assignees for private repos**

See the[ Source Code Manager Configurations](https://training.snyk.io/courses/source-code-manager-configurations?query=Source%20code) training course for details on how to match these integration settings with your security maturity phase, to begin enabling the settings when your teams are ready.

See also [Snyk Integrations](../integrations/) for more details.

If you are using an on-premise source control manager, you must also configure and deploy the [Snyk Broker](snyk-broker/).

### Set up Jira integration

Integrate your Organization with Jira to assist with logging tickets and addressing backlogged security issues. For details, see [Jira integration](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/jira).

### Set up Slack integration

Integrate your Organization with a Slack channel to notify your teams of issues. For details, see [Slack integration](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/slack-integration).

### Define language settings

Based on the nuances of the tech stack you are using, you will want to set your language preferences. For more detail on the nuances of using Snyk in specific circumstances, see the [Guide ](../guides/)specific to your language.

### Integrate private registries

If you are using private registries, you will want to set up those integrations. See [Private registry integrations](../integrations/package-repository-integrations/).

### Enable Snyk Code

When you create a new Organization, Snyk Code (SAST) scanning is disabled by default. Snyk recommends enabling this product before you import your first Projects to Snyk.

1. Select the **Settings > Snyk Code** option.
2. Click the toggle to **enable Snyk Code**; then click **Save changes**.

If you imported Projects prior to enabling, you must re-import them.

See [Getting started with Snyk Code](../scan-application-code/snyk-code/getting-started-with-snyk-code/) for more details.

### Enable Infrastructure as Code

Snyk Infrastructure as Code enables your developers to scan their configuration files for misconfigurations. Enable Infrastructure as Code scans as follows:

1. Open the Organization in the Snyk UI.
2. Select **Settings** from the left navigation pane.
3. Select **Infrastructure as Code**.
4. Select **Enabled** under **Detect configuration files**, and save the changes.

If you imported Projects prior to enabling, you must re-import them.

For more information, see [Snyk Infrastructure as Code](../scan-configurations/snyk-infrastructure-as-code/).

### Configure notifications

Snyk sends teams different types of alerts based on settings defined for the Group and for the Organization. Snyk highly recommends **defining the default settings for the Group and the first Organization with most notifications disabled by default**, before you create additional Organizations and start scanning applications.

If you want alerts to be sent by default for Projects imported into the Organization, you can have Snyk send notifications for either vulnerabilities or license issues, or for both. You can also limit the notifications to only High and Critical severity issues. The following video demonstrates how to set notification defaults.

{% embed url="https://thoughtindustries-1.wistia.com/medias/o41yfkfu34" %}
Organization notification video
{% endembed %}

Encourage individuals to set up their own notification preferences, if they want to customize how they receive alerts for specific Projects.

{% embed url="https://thoughtindustries-1.wistia.com/medias/x1750wshlg" %}
Personal notification video
{% endembed %}

## Invite members

When your Organization is configured, you are ready to invite others to use Snyk. There are two ways in the Snyk UI to invite members. You can invite by entering their email addresses, or add existing Group members to this Organization. Both are demonstrated in the following video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/qqmkcaequj" %}
Invite members video
{% endembed %}

## Create additional Organizations

After you configure your first Organization, you can **use it as a template** for creating additional Organizations.

To create a new Organization:

1. Open the Organization switcher in the navigation panel and select **Create new Organization**.
2. Provide the name of the new Organization. See [Manage Snyk organizations](../snyk-admin/manage-groups-and-organizations/manage-organizations.md) for more details.

After you have created an Organization that you plan to use as a template, complete the setup to configure all of the settings in the way you want them for another Organization. To use **an existing Organization as a template**, select the Organization from the list.

The following settings will be copied from the selected Organization to the new Organization.

* Source control integrations (GitHub, GitLab, BitBucket)
* Container registry integrations (ACR, Docker Hub, ECR, GCR)
* Container orchestrator integrations (Kubernetes)
* PaaS and Serverless integrations (Heroku, AWS Lambda)
* Notification integrations (Slack, Jira)
* Policies
* Ignore settings
* Language settings
* Snyk Code settings
* Infrastructure as Code settings

The new Organization will not use the settings from the copied Organization for the following:

* Service accounts
* Members
* Projects
* Notification preferences

You can then customize any of the Organization settings that you configured for the first Organization for the new Organization.

## Roll out Snyk to your teams

When you roll out Snyk to your teams, you will:

* Identify influential developers as security champions and early adopters.
* Define your communications and training plan.
* Define success metrics based on program maturity.

The [Launch Snyk to your teams](https://training.snyk.io/courses/launch-snyk-to-your-teams) course in Snyk Training provides strategies, communication templates, and checklists for rolling out Snyk to a wider audience.
