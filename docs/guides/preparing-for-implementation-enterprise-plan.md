# Preparing for implementation: Enterprise plan

{% hint style="info" %}
Please read the [Getting started with Snyk: Enterprise plan](getting-started-with-snyk-enterprise-plan.md) guide before you read this guide.
{% endhint %}

In this guide, we’ll look at a few planning steps and walk through some initial configuration options that you should complete before inviting more teams and scanning additional applications.&#x20;

{% hint style="info" %}
If you have been using Snyk on the Free/Team Plan and are looking for guidance on upgrading to the Enterprise plan, read the [Upgrading to Enterprise Plan](upgrading-to-enterprise-plan.md) guide.
{% endhint %}

### Topics to consider

There are several topics that you should consider before you launch Snyk to other teams in your organization. Here are some areas to consider:

* Where will you implement Snyk in your Software Development Life Cycle (SDLC)?
* How will you structure your Snyk account?
* How will you and your team access Snyk?
* How will you define license and security policies?
* How will you scan your applications?
* How will you configure your first organization?

The next section of this guide will discuss each of these topics. After you make these decisions and configure your first Organization, you can configure additional Organizations and invite members.

### Implementation checklist

As you work through the rest of this guide, use the following checklist to ensure you have prepared for your implementation.

* [ ] Set up SSO (if applicable).
* [ ] Determine Organization structure and naming convention.
* [ ] Define notification defaults (for the Group, and for the first Organization).
* [ ] Define success metrics based on program maturity.
* [ ] Identify influential developers as security champions/early adopters.
* [ ] Define communications and training plan.
* [ ] Set up first Organization.
* [ ] Develop project import strategy and begin importing projects.
* [ ] Create additional organizations and begin importing projects.

### Where to implement Snyk in your SDLC

As you prepare for your teams to start adopting Snyk as part of a secure development workflow, you should decide where to use Snyk in your SDLC, and what you want to scan (such as application code, open-source library dependencies, container registries, etc.)

You will also want to roll Snyk out in phases depending on how far you and your teams have progressed with developer security (awareness, visibility, preventing issues, fixing the backlog, optimizing). See the [Ways to integrate Snyk at your company](https://training.snyk.io/courses/ways-to-use-snyk) training course for more details.

### How to structure your Snyk account

For Enterprise plans, the hierarchy for your Snyk account includes a Group at the highest level, Organizations within that Group, and Projects within those Organizations.&#x20;

<figure><img src="https://lh6.googleusercontent.com/yt7XP5scEUM8_rRMTZOYUvuUPOP23QPoFw551WLspjaU1tUtCvo8kOhPHJ-Z2tYIIoijMmcYraKqnLtDlQIzgeX5ZaeclHJWrGA9pcaprWK6hG898FV6BJ2Pzx4PUsRkewXv_F5M58ruiyVJu9pTnw" alt="Snyk account hierarchy"><figcaption><p>Snyk account hierarchy</p></figcaption></figure>

{% hint style="info" %}
For very large implementations of Snyk, you may also choose to implement multiple groups, such as to align top levels with separate business units. Please contact your Account Executive if you are interested in discussing this option.
{% endhint %}

Consider this hierarchy, as well as deciding which members of your teams need access to specific Snyk Projects, when you decide how to align your Snyk Organizations with your internal structure. You’ll want to decide how to structure your account before inviting members or scanning a large number of applications.

To make these decisions, ask yourself:

* Which team members can access specific repositories?
* How do you want to apply policies to Snyk Projects for prioritizing and automating tests?
* How you want to report on scans?

For more help addressing these questions, see the free, self-paced Training course [Snyk account structure](https://training.snyk.io/courses/snyk-account-structure).

### How your teams log in to Snyk

There are a few different ways that users can authenticate into their Snyk accounts, such as with a GitHub or Google account.&#x20;

<figure><img src="https://lh5.googleusercontent.com/EYz6tsF_VfAbsLnz4dZyPa_BH40wJw2tX4sfKhyCD7EakbIahU3KtajVtkoxcpzLrM3m--x-EXFUz--qLlIDj5dFBvSmGM6qSmteM3K-7r6vWPBndJT-blCz_EQq6KU-zaMk_KHKZzBXjpDw5260wg" alt="Logging in to Snyk"><figcaption><p>Logging in to Snyk</p></figcaption></figure>

#### Single Sign-On (SSO)

You may want to set up SSO using your existing identity provider to streamline sign-ins and new user provisioning.

If using SSO, after you set it up, you'll need to remove any social login accounts from Snyk so that all user access to your Group is controlled via SSO.

See [Setting up Single Sign-on (SSO) for authentication](https://docs.snyk.io/user-and-group-management/setting-up-sso-for-authentication) for more details.

### Define policies

The Policy manager allows you to define license and security policies at the Group level.&#x20;

License policies define how your Organizations identify license compliance issues in your Open Source packages and change the severity of issues associated with use of those licenses.&#x20;

See [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/products/snyk-open-source/licenses/getting-started-snyk-licensing-compliance) for more details.

Security policies define the way your Organizations handle specific issues or exploit maturity levels. A policy can programmatically increase an issue’s severity rating, decrease the severity rating, or ignore the issue.&#x20;

See [Getting started with security policies](https://docs.snyk.io/manage-issues/security-policies/getting-started-with-security-policies) for more details.

### How you’ll scan your applications

Snyk Projects are the components Snyk tests, along with the related configuration and metadata. Each target you want to scan (repos, container images, Dockerfiles, configuration files, source code) may include more than one Project.

See [Introduction to Snyk Projects](https://docs.snyk.io/manage-issues/introduction-to-snyk-projects) for more details.

There are different ways to add Projects in Snyk, including via an integration, the Snyk CLI, or Snyk API.&#x20;

Before you start importing Projects, make sure your Organizations in Snyk are configured appropriately. You’ll want to determine how you want to configure your first Organization and complete all of the settings first. Then you can use that Organization as a template for configuring other Organizations. Then you can begin scanning applications.&#x20;

### Configure your first Organization

Configure your first Organization before you start scanning applications and inviting members, to ensure the best experience.

To save time, we recommend setting up your first Organization fully so that this Organization can serve as a template when creating other Organizations via the Snyk Web UI or via an API.

Follow these steps to ensure you have completed the configuration.

#### Set ignore permissions

Snyk allows you to ignore issues in a few different ways. An ignored issue is not deleted. It is only removed from the filtered list of open issues.

Before inviting additional team members to Snyk, determine who can ignore the vulnerabilities and license issues that Snyk identifies.&#x20;

Go to **Settings > General** to specify the ignore permissions.

{% hint style="info" %}
Snyk recommends that only Snyk Admins are allowed to ignore issues, and that all ignores require a reason to be recorded.
{% endhint %}

#### Configure Git repository integrations

Snyk includes a number of automated processes for Snyk Open Source when integrated with a source code manager (SCM) on a Git repository. These automated processes are a great way to mature your developer security program. However, the automated processes can introduce frustration for developers if introduced too early in your journey. Make sure your settings align with your phase of adoption.

You may decide to disable Snyk test for pull requests and other automated processes until your teams are ready.&#x20;

Use the following steps to disable these automated processes for the Organization until you are ready to implement them. Note that all of these settings may not be displayed, depending on the specific Integration you are configuring.

1. Open the Organization in the Snyk UI.
2. Select **Integrations** from the left navigation pane.
3. Select the cog icon adjacent to integration you want to configure (for example, GitHub).
4. In the **Snyk test for pull requests** section, disable the following:
   1. **Open Source Security & Licenses**
   2. **Code Analysis**
   3. **Automatic fix pull requests**, both **New vulnerabilities** and **Known vulnerabilities (backlog)**
   4. &#x20;**Automatically Update Dockerfile base images**
   5. **Automatic dependency upgrade pull requests**
   6. **Pull request assignees for private repos**

See the[ Source Code Manager Configurations](https://training.snyk.io/courses/source-code-manager-configurations?query=Source%20code) training course for details of how to match these integration settings with your security maturity phase, to begin enabling the settings when your teams are ready.

Also see [Snyk Integrations](https://docs.snyk.io/integrations) for more details.

If you are using an on-premise source control manager, you must also configure and deploy the [Snyk Broker](../user-and-group-management/snyk-broker/).&#x20;

#### Set up Jira integration

Integrate your organization with Jira to assist with logging tickets and addressing backlogged security issues.&#x20;

Also see [Jira integration](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/jira).

#### Set up Slack integration

Integrate your organization with a Slack channel to notify your teams of issues.

Also see [Slack integration](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/slack-integration).

#### Define language settings

Based on the nuances of the tech stack you are using, you'll want to also set your language preferences. For more detail on the nuances of using Snyk with specific circumstances, visit the Guide specific to your language.

#### Integrate private registries

If you are using private registries, you'll want to set up those integrations. See [Private registry integrations](https://docs.snyk.io/integrations/private-registry-integrations).

#### Enable Snyk Code

When you create a new Organization, Snyk Code (SAST) scanning is disabled by default. We recommend enabling this product before you import your first projects to Snyk.

1. Select the **Settings > Snyk Code** option.
2. Click the toggle to enable Snyk Code, then click **Save changes**.

If you imported projects prior to enabling, you will need to re-import them.

See [Getting started with Snyk Code](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code) for more details.

#### Enable Infrastructure as code

Snyk Infrastructure as Code enables your developers to scan their configuration files for misconfigurations. Confirm or enable that Infrastructure as Code scans are enabled using the following steps:

1. Open the Organization in the Snyk UI.
2. Select Settings from the left navigation pane.
3. Select Infrastructure as code.
4. Select Enabled under Detect configuration files, and save the changes.

If you imported projects prior to enabling, you will need to re-import them.

Also see [Snyk Infrastructure as Code ](https://docs.snyk.io/scan-cloud-deployment/snyk-infrastructure-as-code)for more details.

#### Configure notifications

Snyk sends teams different types of alerts based on settings defined for the Group and for the Organization. It's highly recommended to define the default settings for the Group and the first Organization with most notifications disabled by default, before you create additional Organizations and start scanning applications.&#x20;

If you want alerts to be sent by default for projects imported into the Organization, you can have Snyk send notifications for either vulnerabilities and/or license issues. You can also limit the notifications to only High and Critical severity issues.

{% embed url="https://thoughtindustries-1.wistia.com/medias/o41yfkfu34" %}
Organization notification video
{% endembed %}

Encourage individuals to set up their own notification preferences, if they want to customize how they receive alerts for specific projects.

{% embed url="https://thoughtindustries-1.wistia.com/medias/x1750wshlg" %}
Personal notification video
{% endembed %}

#### Invite members

When your Organization is configured, you are ready to invite others to use Snyk. There are two ways in the Snyk UI to invite members. You can invite by entering their email addresses, or add existing Group members to this Organization. Both are demonstrated in the following video.

{% embed url="https://thoughtindustries-1.wistia.com/medias/qqmkcaequj" %}
Invite members video
{% endembed %}

### Create additional Organizations

After you configure your first Organization, you can use it as a template for creating additional Organizations.&#x20;

To create a new Organization, open the Organization switcher in the navigation panel and select **Create new Organization**.

Provide the name for the new Organization. See [Manage Snyk organizations](../snyk-admin/managing-groups-and-organizations/manage-snyk-organizations.md) for more details.

#### Using an existing Organization as a template

You can use an existing Organization as a template. Make sure you have completed the configurations for the Organization before copying it.&#x20;

Select the Organization from the list. The following settings will be copied from the selected Organization:

* Source control integrations (GitHub, GitLab, BitBucket)
* Container registry integrations (ACR, Docker Hub, ECR, GCR)
* Container orchestrator integrations (Kubernetes)
* PaaS and Serverless integrations (Heroku, AWS Lambda)
* Notification integrations (Slack, Jira)
* Policies
* Ignore settings
* Language settings
* Infrastructure as Code settings
* Snyk Code settings

The new Organization will not use the same settings from the copied Organization for the following:

* Service accounts
* Members
* Projects
* Notification preferences

Any of the Organization settings that you configured for the first Organization can then be customized for the new Organization.

### Ready to roll out Snyk to your teams?

The [Launch Snyk to your teams](https://training.snyk.io/courses/launch-snyk-to-your-teams) course in Snyk Training provides strategies, communication templates, and checklists for rolling out Snyk to a wider audience.
