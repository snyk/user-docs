# Configure Group settings and policies

This page is designed to help you plan your Snyk account structure and policies at Group-level to ensure efficient asset management, precise access control, and accurate reporting.

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

* **Business impact**: Applications that handle sensitive data or generate high revenue.
* **Exposure**: Public-facing apps or those with high-risk network configurations.
* **Development activity**: Teams with active release cycles that can implement fixes quickly.

## Structure your account

{% hint style="success" %}
**Key decisions:**

* Select the structure (team, product, or SCM-based) that best supports how you want to manage policies, report vulnerabilities, and define user access.
* Define who can provision users and grant Snyk access to external platforms, such as Git repositories.
{% endhint %}

Snyk uses a hierarchical structure to manage all assets and security policies. This section will help you map your business to the Snyk architecture.

* **Understanding the hierarchy:** Learn the purpose of the Tenant, Group, Organization, and Project levels.
* **Group structure:** Decide how many top-level Group accounts your company requires.

To learn more, visit [Structure your account](structure-your-account.md).

## Authentication and access: set up SSO

Implement Single Sign-On (SSO) at the Group level before rolling Snyk out to your organization. While pilot teams often start with personal authentication, transitioning to SSO is required for broad adoption, consistent login access, and centralized control.

Choose a provisioning strategy that defines the user experience and access level for new users.

| Option                | Description                                                               | Recommended for                                             |
| --------------------- | ------------------------------------------------------------------------- | ----------------------------------------------------------- |
| **Open to all**       | Any user with a valid company email domain can join the Snyk Group.       | Rapid, frictionless onboarding.                             |
| **Require an invite** | Users must receive an email invitation to join the Group or Organization. | Strict control over license seats and access.               |
| **API provisioning**  | Use Snyk API endpoints to grant access and roles before a user logs in.   | Pre-defining custom roles and specific Organization access. |

To learn more, visit [Authentication and access](authentication-and-access.md).

## Authentication and access: role management

{% hint style="success" %}
**Key decisions:**

* Determine if your team leads can operate with the fixed permissions of an **Organization Admin** or if they require a restricted custom role.
* Identify if you need a "middle-ground" role, such as a Senior Developer who can, for example, override checks but cannot delete Projects.
* Decide who is responsible for provisioning new users as your Snyk footprint grows.
{% endhint %}

Snyk provides both pre-defined and custom roles to ensure users have the exact permissions they need.

* **Tenant-level roles:** Determine if you need centralized oversight roles (Tenant Admin, Viewer, or Member) for cross-group reporting and user management.
* **Group roles:** Review fixed pre-defined roles.
* **Custom roles:** Evaluate whether you need roles with highly granular, customized permissions.

To learn more, visit [Authentication and access.](authentication-and-access.md)

## Define policies

{% hint style="success" %}
**Key decisions:**

* Decide which conditions automatically increase or decrease the priority or severity of an issue to match your risk appetite.
* Decide which specific issues or types of issues are automatically ignored to reduce "noise" and save development time.
* Decide which specific license types to explicitly allow or disallow to avoid using packages with incompatible or problematic licenses.
* Use severity types to decide how to configure policies to match your specific legal and compliance requirements.
* Decide how to automate the governance, tracking, and remediation workflows for your assets to ensure continuous security visibility and compliance
{% endhint %}

Policies allow you to automate business context, compliance checks, and notification workflows.

* **License policies:** Decide which specific license types to explicitly allow or disallow to match your specific legal and compliance requirements.
* **License severity:** Understand and configure how Snyk assigns High or Medium severity to problematic commercial licenses.
* **Asset governance:** Decide how to automate the governance, tracking, and remediation workflows for your assets.
* **Policy creation:** Learn how to build asset policies using specific filters (like Name, Asset Type, or Tags) and actions (like assigning classification or triggering Slack notifications).

To learn more, visit [Define policies](define-policies.md).
