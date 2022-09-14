# Implementing Snyk: Business and Enterprise plan users

{% hint style="info" %}
See [Implementing Snyk: Free and Team plan users](implementing-snyk-free-and-team-plan-users/) for details of using Snyk with those [pricing plans](https://snyk.io/plans/).
{% endhint %}

There are a few key considerations and tasks to get started using Snyk, including:

* [Make implementation decisions](implementing-snyk-business-and-enterprise-plan-users.md#make-implementation-decisions)
* [Set up your first organization](implementing-snyk-business-and-enterprise-plan-users.md#set-up-your-first-organization)
* [Create additional organizations](implementing-snyk-business-and-enterprise-plan-users.md#create-additional-organizations)
* [Roll out Snyk to developers](implementing-snyk-business-and-enterprise-plan-users.md#roll-out-snyk-to-developers)
* [Find, prioritize, and fix issues](implementing-snyk-business-and-enterprise-plan-users.md#find-prioritize-and-fix-issues)

### Make implementation decisions

Get off to the right start with a few planning steps before inviting your teams and scanning your projects. There are a few decisions to make to get started:

* [where to implement Snyk in your software development life cycle (SDLC)](implementing-snyk-business-and-enterprise-plan-users.md#where-to-implement-snyk)
* [how to structure your account](implementing-snyk-business-and-enterprise-plan-users.md#how-to-structure-your-account)
* [how to access Snyk](implementing-snyk-business-and-enterprise-plan-users.md#how-to-access-snyk)
* [how you'll import projects](implementing-snyk-business-and-enterprise-plan-users.md#plan-for-importing-projects)

#### Where to implement Snyk

As you prepare for your teams to start adopting Snyk as part of a secure development workflow, you’ll want to decide where in your software development lifecycle you are implementing Snyk and which of the Snyk platform products you are using. You'll also want to roll Snyk out in phases based on where you are with developer security (awareness, visibility, preventing issues, fixing the backlog, optimizing).

More information:

* Docs: [Integrations](../products/snyk-code/introducing-snyk-code/key-features/integrations.md)
* Snyk Training: [Ways to integrate Snyk at your company](https://training.snyk.io/courses/ways-to-use-snyk)

#### How to structure your account

There are different ways to arrange the organizations in your account. Before inviting members to your account or scanning for issues, plan which organizations you'll need based on how you want to allow permissions and access to projects. The way you set policies are the next order of consideration. And how you want to report on projects are the third consideration.

More information:

* Docs: [User and group management](../features/user-and-group-management/)
* Training: [Snyk account structure](https://training.snyk.io/courses/snyk-account-structure)

#### How to access Snyk

There are a few different ways that users can authenticate into their Snyk accounts, such as with a GitHub or Google account. You may want to set up single sign-on (SSO) via your existing identity provider to streamline sign-ins and new user provisioning.

If using SSO, after you set it up, you'll need to remove any social login accounts from Snyk.

More information:

* Docs: [Setting up Single Sign-On (SSO) for authentication](../features/user-and-group-management/setting-up-sso-for-authentication/)
* Training: [SSO, authentication and user provisioning](https://training.snyk.io/courses/sso)

#### Plan for importing projects

There are different ways to add projects in Snyk, including via an integration, the CLI, or Snyk API. However, before importing projects, make sure your organizations in Snyk are configured appropriately.

More information:

* Docs: [Get started with your Snyk product and environment](https://docs.snyk.io/getting-started#step-2-get-started-with-your-snyk-product-and-environment)
* Training: [Project import strategies](https://training.snyk.io/courses/project-import-strategies)

### Set up your first organization

For a small team, you may only need one organization. For bigger teams, you'll identify more organizations, which are aligned to your needs, in the earlier decisions. When planning multiple organizations, you'll want to complete all of the configurations for the first organization, so that you can use those settings as a template for creating any other organizations via the Snyk Web UI or via an API.

#### Configure Git repository integrations

Snyk includes a number of automations for Snyk Open Source when integrated with a source code manager (SCM) on a Git repository. These automations are a great way to mature your developer security program. However, the automations can introduce frustration for developers if introduced too early in your journey. Make sure your settings align with your phase of adoption.

More information:

* Docs: [Git repository (SCM) integrations](../integrations/git-repository-scm-integrations/)
* Training: [Source Code Manager Configurations](https://training.snyk.io/courses/source-code-manager-configurations)

If you are using an on-premise source code manager, learn more about Snyk Broker:

* Docs: [Snyk Broker](https://docs.snyk.io/features/snyk-broker)

#### Define default license policy

Snyk can identify license compliance issues in your Open Source packages. The Default License Policy indicates the severity associated with the use of different licenses, along with an option to provide license instructions if Snyk finds these licenses in your projects.

More information:

* Docs: [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/products/snyk-open-source/licenses/getting-started-snyk-licensing-compliance)

#### Configure notifications

Snyk sends teams different types of alerts based on settings defined for the group and for the organization. It's highly recommended to define the default settings for the group and the first organization with most notifications disabled by default before you create additional organizations and import projects. Individual users can set up their own notification preferences to receive alerts for specific projects.

More information:

* Docs: [Managing notifications](../features/user-and-group-management/notifications.md)
* Training: [Notifications](https://training.snyk.io/courses/notifications)

#### Set ignore permissions

Before inviting additional team members to Snyk, determine who can ignore the vulnerabilities and license issues that Snyk identifies.

More information:

* Docs: [Set who can configure ignore settings](https://docs.snyk.io/features/fixing-and-prioritizing-issues/issue-management/ignore-issues#set-who-can-configure-ignore-settings)
* Training: [Members and Permissions](https://training.snyk.io/courses/members-and-permissions)

#### Define language settings

Based on the nuances of the tech stack you are using, you'll want to also set your language preferences.

More information:

* Docs: [Python Settings](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-python)
* Docs: [JavaScript Settings](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-javascript)
* Docs: [.NET Settings](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-.net)
* Docs: [PHP Settings](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-php)
* Docs: [Maven Settings](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support/snyk-for-java-gradle-maven)

#### Integrate private registries

If you are using private registries, you'll want to set up those integrations.

More information:

* Docs: [Private registry integrations](https://docs.snyk.io/integrations/private-registry-integrations)

#### Set up Jira integration

To allow Snyk users to create a Jira issue for vulnerabilities and license issues that Snyk identifies, you'll want to integrate Snyk with your Jira instance.

More information:

* Docs: [Jira](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/jira)

### Create additional organizations

Once the first organization has the desired configurations, you can copy it to create additional organizations. Other organizations can have different settings for the following as needed: source control manager integrations, license policy, notification settings, ignore permissions, language settings, and Jira integration.

More information:

* Docs: [Creating a new Snyk organization](https://docs.snyk.io/features/user-and-group-management/managing-groups-and-organizations/manage-snyk-organizations#creating-a-new-snyk-organization)
* Training: [Snyk account structure](https://training.snyk.io/courses/snyk-account-structure)

### Roll out Snyk to developers

#### Prepare teams for using Snyk

Our Developer Launch Package provides a number of resources to help you prepare for launching Snyk to a wider audience.

* Training: [Developer Launch Package](https://training.snyk.io/courses/launch-snyk-to-your-teams)

#### Invite members

Once your organization(s) are configured, you're ready to invite other users to Snyk.

More information:

* Docs: [Manage users in your organization](https://docs.snyk.io/features/user-and-group-management/managing-users-and-permissions/manage-users-in-your-organizations)
* Training: [Members and Permissions](https://training.snyk.io/courses/members-and-permissions)

### Find, prioritize, and fix issues

#### Use Snyk in an IDE

Empower developers to find and fix issues early in the development process by adding the Snyk plugin to their integrated development environment.

More information:

* Docs: [Snyk for IDEs](../ide-tools/)
* Training: [Introduction to using Snyk in an IDE](https://training.snyk.io/courses/introduction-to-using-snyk-in-an-ide)

#### Use Snyk in the CLI

The Snyk CLI provides a way to find security and license issues locally or in your CI/CD pipeline.

More information:

* Docs: [Snyk CLI](../snyk-cli/)
* Training: [Introduction to the Snyk CLI](https://training.snyk.io/courses/intro-cli)

#### Use Snyk in the Web UI

The Snyk Web UI scans different types of projects (depending on which product you have purchased), displays and scores the scan results, allows you to prioritize the results, and offers different types of fix advice or options (depending on several factors).

More information:

* Docs: [Snyk Web UI](../snyk-web-ui/)
* Training: [Introduction to Snyk](https://training.snyk.io/learning-paths/introduction-to-snyk), [Find and Fix with Snyk Open Source](https://training.snyk.io/learning-paths/find-and-fix-with-snyk-open-source), [Introduction to the Snyk UI](https://training.snyk.io/courses/introduction-to-the-snyk-ui)
