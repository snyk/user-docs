---
description: How to create a template Snyk Organization to standardize account structure, access control, and reporting
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Create a template Organization

Use this section to plan your Snyk account structure at the Organization level to ensure efficient asset management, precise access control, and accurate reporting. This Organization becomes the template you copy to create additional Organizations, so that all your Organizations share standard settings.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1775661970/4._Creating_your_Template_Organization_vf1x0b.mp4" %}
Create a template Organization video guide
{% endembed %}

The following sections outline the key topics and decisions covered in this part of the guide.

## Plan your Organization structure

{% hint style="success" %}
**Key decisions:**

* Select the structure (team, product, or SCM-based) that best supports how you want to manage policies, report vulnerabilities, and define user access.
* Define who can provision users and grant Snyk access to external platforms, such as Git repositories.
{% endhint %}

Snyk uses a hierarchical structure to manage all assets and security policies. This section helps you map your business to the Snyk architecture.

* Organization structure: Determine whether a team-based, product-based, or SCM organization-based Organization structure best supports your policy and access needs.

To learn more, visit [Plan your Organization structure](structure-your-account.md).

## Authentication and access

{% hint style="success" %}
**Key decisions:**

* Determine if your team leads can operate with the fixed permissions of an **Organization Admin** or if they require a restricted custom role.
* Identify if you need a "middle-ground" role, such as a Senior Developer who can, for example, override checks but cannot delete Projects.
* Decide who is responsible for provisioning new users as your Snyk footprint grows.
{% endhint %}

Snyk provides both pre-defined and custom roles to ensure users have the exact permissions they need.

* Organization roles: Review fixed pre-defined roles.
* Role alignment: Learn how to assign roles based on your chosen Organization structure.

To learn more, visit [Authentication and access](authentication-and-access.md).

## Connect your development tools

{% hint style="success" %}
**Key decisions:**

* Architecture and connectivity: Select your primary SCM platform and identify where your production container registries (for example, Docker Hub, Amazon ECR) reside. Determine if you need to deploy Snyk Broker (using Docker or Kubernetes) to securely bypass firewalls, segment network traffic, or if you require custom proxy/CA configurations.
* Authentication and access management: Commit to using dedicated, least-privilege service accounts rather than personal tokens (OAuth/PAT) to guarantee stable connections. Decide whether to use broad, centralized Group-level credentials for global asset discovery or require unique tokens at the Organization level. Verify you have the correct administrative roles to manage this inventory.
* Developer workflow and remediation: Define your PR check behaviors and automated fix strategies to match specific team workflows. Establish a monitoring frequency that provides adequate security visibility without overwhelming your developers' capacity to remediate.
* Snyk Code enablement: Decide whether to roll out Snyk Code globally or phase it in for high-priority teams. You must enable Snyk Code before importing your first Projects. If you manage Organizations at a large scale, plan to use the Snyk API rather than the web UI for this step.
* Project import strategy: Choose the import method that provides the best dependency resolution for your specific programming languages to ensure accurate scanning.
{% endhint %}

By configuring your core tools, source control integrations, and default security behaviors now, you establish a standardized baseline that you can clone, either manually or using the API, across your entire business.

To learn more, visit [Connect your development tools](connect-your-development-tools.md).
