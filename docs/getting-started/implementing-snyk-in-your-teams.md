# Implementing Snyk in your teams

For security and DevOps professionals preparing to roll Snyk out to your development teams, here are some key considerations and tasks:

* [Make implementation decisions](implementing-snyk-in-your-teams.md#implementation-decisions).
* [Set up your first organization](implementing-snyk-in-your-teams.md#set-up-first-organization).
* [Learn to find, prioritize and fix issues](implementing-snyk-in-your-teams.md#learn-to-find-prioritize-and-fix-issues).

### Make implementation decisions

#### Where to implement Snyk

As you prepare for your teams to start adopting Snyk as part of a secure development workflow, you’ll want to decide where in your software development lifecycle you are implementing Snyk and which of the Snyk platform products you are using.&#x20;

More information:

* Docs: [integrations.md](../products/snyk-code/key-features/integrations.md "mention")
* Snyk Training: [Ways to integrate Snyk at your company](http://training.snyk.io/learn/course/ways-to-use-snyk/ways-to-integrate-snyk-at-your-company/snyk-across-the-sdlc?page=1)

#### How to structure your account

There are different ways to structure your account to align to the organization of your teams and desired project access. The highest level of the hierarchy within Snyk is a Group. Organizations, the next level of the hierarchy, manage access to projects. Projects are the hierarchy levels where Snyk scans for issues.&#x20;

More information:

* Docs: [user-and-group-management](../features/user-and-group-management/ "mention")
* Training: [Snyk account structure](http://training.snyk.io/learn/course/snyk-account-structure/account-structure-considerations/an-overview-of-account-structure?page=1)

#### How to access Snyk

There are a few different ways that users can authenticate into their Snyk accounts, including setting up single sign-on via your existing identity provider.&#x20;

More information:

* Docs: [authentication](../features/user-and-group-management/authentication/ "mention")
* Training: [SSO, authentication and user provisioning](http://training.snyk.io/learn/course/sso/authentication-to-snyk/an-overview-of-authentication-and-provisioning?page=1)

#### Plan for importing projects

Before importing projects, make sure your organizations in Snyk are configured appropriately. There are different ways to add projects in Snyk, including via an integration, the CLI or Snyk API.&#x20;

More information:

* Docs: [introduction-to-snyk-projects](introduction-to-snyk-projects/ "mention")
* Training: [Project import strategies](http://training.snyk.io/learn/course/project-import-strategies/project-import-considerations/about-projects?page=1)

### Set up your first organization

#### Plan for Snyk notifications

Snyk sends teams different types of alerts based on settings defined for the group and for the organization. It's highly recommended to define the default settings for the group and the first organization with most notifications disabled by default before you create additional organizations and import projects. Individual users can set up their own notification preferences to receive alerts for specific projects.&#x20;

More information:

* Docs: [notification-management.md](../features/user-and-group-management/notifications/notification-management.md "mention")
* Training: [Notifications](http://training.snyk.io/learn/course/notifications/snyk-notification-configuration/an-overview-of-notifications?page=1)

#### Set source code manager (SCM) configurations

Snyk includes a number of automations for Snyk Open Source when integrated with a source code manager. These automations are a great way to mature your developer security maturity. However, they can introduce frustration for developers if introduced too early in your journey. Make sure your settings align with your phase of adoption.&#x20;

More information:

* Docs: [git-repository-scm-integrations](../features/integrations/git-repository-scm-integrations/ "mention")
* Training: [Source Code Manager Configurations](http://training.snyk.io/learn/course/source-code-manager-configurations/integration-scenarios/overview?page=1)

{% hint style="info" %}
After the first organization is configured, you can copy the settings when creating new organizations.
{% endhint %}

#### Invite members and set ignore permissions

When it’s time to invite additional team members to Snyk, you will also determine how to handle ignoring issues in Snyk. It’s important to understand why you would ignore issues and who has permission to do so.&#x20;

More information:

* Docs: [managing-users-and-permissions](../features/user-and-group-management/managing-users-and-permissions/ "mention")
* Training: [Members and Permissions](http://training.snyk.io/learn/course/members-and-permissions/members-and-permissions/member-invitations?page=1)

### Learn to find, prioritize, and fix issues

#### Get familiar with the Snyk UI

The Snyk platform scans different types of projects (depending on which product you have purchased), displays and scores the scan results, and offers different types of fix advice or options (depending on several factors).&#x20;

More information:

* Docs: [getting-started-snyk-products.md](getting-started-snyk-products.md "mention")
* Training: [Introduction to the Snyk UI](http://training.snyk.io/learn/course/introduction-to-the-snyk-ui/snyk-log-in-details/log-in-tasks-and-considerations?page=1)

#### Get familiar with the Snyk CLI

The Snyk CLI uses the command line interface to work with Snyk in a local environment or via a CI/CD process.&#x20;

More information:

* Docs: [snyk-cli](../snyk-cli/ "mention")
* Training: [Introduction to the Snyk CLI](http://training.snyk.io/learn/course/intro-cli/snyk-cli/authentication-to-snyk-account)

#### Get familiar with Snyk in an IDE

IDE plugins for Snyk take advantage of the Snyk CLI to offer Snyk functionality in an individual development environment.&#x20;

More information:

* Docs: [ide-tools](../features/integrations/ide-tools/ "mention")
* Training: [Introduction to using Snyk in an IDE](http://training.snyk.io/learn/course/introduction-to-using-snyk-in-an-ide/snyk-plugin-for-ide/authenticate-plug-in-to-snyk?page=1)

