# Enterprise implementation guide

Rolling out a developer security platform across a large organization requires a consistent, repeatable approach to account configuration. In this guide, you will learn how to streamline your enterprise rollout in Snyk by creating a dedicated template Organization and altering settings for each individual Organization created from that original template.

To have a successful Snyk rollout, you need to:

1. [Understand the Snyk hierarchy.](./#snyk-hierarchy)
2. [Configure Group settings and policies](configure-group-settings-and-policies/):
   1. [Structure your account](configure-group-settings-and-policies/structure-your-account.md).
   2. [Authentication and access](configure-group-settings-and-policies/authentication-and-access.md).
   3. [Define policies](configure-group-settings-and-policies/define-policies.md).
3. [Create a template Organization](create-a-template-organization/):
   1. [Structure your account](create-a-template-organization/structure-your-account.md).
   2. [Authentication and access](create-a-template-organization/authentication-and-access.md).
   3. [Connect your development tools](create-a-template-organization/connect-your-development-tools.md).
4. [Create your Snyk structure](create-your-snyk-structure.md).
5. [Gain issues visibility with Projects](phase-3-gain-visibility/).
6. [Initial team rollout](initial-team-rollout.md).
7. [Automate prevention measures](automate-prevention-measures.md).

By the end of this guide, you will have successfully configured a baseline template Org, used it to easily copy your global settings and integrations to provision new Organizations, and learned how to customize individual settings in those new environments to support specific team workflows.

{% hint style="info" %}
To understand how AI is used at Snyk and how this may affect your implementation decisions, visit the [AI Data and Governance](../../snyk-data-and-governance/how-snyk-incorporates-generative-ai-into-the-platform.md) page.
{% endhint %}

## Confirm points of contact

{% hint style="success" %}
**Key decision**: Determine which teams manage your identity provider (IdP) and source control management (SCM) systems. Snyk requires permissions that typically live outside your security team. Identifying these stakeholders early on prevents your rollout from stalling mid-configuration.
{% endhint %}

To implement Snyk successfully, you must identify the necessary skills and the internal stakeholders required for the rollout.

For example, identify people who can perform the following tasks:

* Create the required single sign-on (SSO) connections.
* Generate tokens with the necessary permissions for SCM repositories.
* Configure permissions for other integrations, such as container registries or CI/CD pipelines.

## Map business critical applications

{% hint style="success" %}
**Key decision**: Select three to five high-priority applications to serve as initial implementation benchmarks. Starting at a smaller scale enables you to validate the configuration and surface any integration issues before a large-scale rollout.
{% endhint %}

To identify priority applications, categorize them based on:

* **Business impact**: Applications that handle sensitive data or generate high revenue.&#x20;
* **Exposure**: Public-facing apps or those with high-risk network configurations.&#x20;
* **Development activity**: Teams with active release cycles that can implement fixes quickly.

## Snyk hierarchy

Plan your Snyk account structure to ensure efficient asset management, precise access control, and accurate reporting. Snyk uses a hierarchical structure to manage all assets and security policies.

Snyk organizes work into four levels:

| Level            | Typical representation | Purpose                                           |
| ---------------- | ---------------------- | ------------------------------------------------- |
| **Tenant**       | Your entire company    | The overarching instance managed by Snyk.         |
| **Group**        | Business entity        | The top level of your administrative control.     |
| **Organization** | Team or product        | The key level for managing Projects and policies. |
| **Project**      | Imported asset         | The individual item Snyk tests and monitors.      |

## Configure Group settings and policies

As a Group Admin or above you need to decide on which settings to enable and what to configure at the Group level. This is so your business entity is represented at the top-level of control and you can cleanly create an Organization template from these configurations.

To learn more, visit [Configure Group settings and policies](configure-group-settings-and-policies/).

## Create a template Organization

Create a template Organization to clone settings and integrations. This ensures consistency across your Snyk structure and reduces manual configuration.

Snyk does not have a dedicated template feature. Instead, you can create a standard Organization named "Template" and configure its settings.

To learn more, visit [Create a template Organization](create-a-template-organization/).

## Create your Snyk structure based on the template Organization

{% hint style="success" %}
**Key decision**: Determine how you want to structure your Organizations. Do you want to reflect your business structure by having an Organization per team, per business unit, or by repository? This is important because your chosen Org structure defines how your teams work in Snyk, controls who sees what, and determines how policies and reporting are scoped.
{% endhint %}

You can use an existing Organization as a model to create a new Organization, and apply the settings and integrations.&#x20;

Use the API templating functionality to ensure consistent settings when creating Organizations at scale. There are two scenarios where you can apply this:

* **Mirroring existing structures**: Use the `snyk-api-import` tool to replicate an existing source, for example, GitHub Organizations.&#x20;
* **Direct API creation**: Use the `Create a new organization` endpoint and provide a `sourceOrgId` to apply the template.

To learn more, visit [Create your Snyk structure](create-your-snyk-structure.md).

## Gain visibility by importing repositories

Gaining visibility into your organization's security starts with importing your repositories so Snyk can actively monitor your code, dependencies, containers, and infrastructure.

Depending on your tech stack and workflows, you can choose from three primary import methods:

* SCM integrations for automated scanning
* Snyk CLI for granular control within your CI/CD pipelines
* Snyk API for large-scale programmatic automation

Once you import your Projects, you can apply specific tags and attributes to easily categorize, filter, and generate targeted reports across your entire portfolio.

To learn more, visit [Gain visibility by importing repositories](phase-3-gain-visibility/).

## Initial team rollout

Successfully rolling out Snyk to your development team requires clear communication, targeted education, and seamless workflow integration.

Start by using customizable templates to announce the launch, tailoring your messaging and feature rollout to match your team's comfort level with security automation.

Next, empower your developers with persona-specific training through Snyk Learn to build both foundational security knowledge and Snyk product expertise.&#x20;

Finally, reinforce a shift-left security strategy by deploying Snyk IDE plugins, equipping developers to find and fix vulnerabilities directly within their preferred coding environment before the code ever reaches the CI/CD pipeline.

To learn more, visit [Initial team rollout](initial-team-rollout.md).

## Automate prevention measures

Drive user adoption and secure coding practices by thoughtfully introducing Snyk to your development teams.

By combining clear communication, tailored training through Snyk Learn, and the deployment of IDE plugins, you can empower your developers to actively find and fix vulnerabilities directly within their daily workflows.

To learn more, visit [Automate prevention measures](automate-prevention-measures.md).
