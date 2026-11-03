---
description: How to structure a Snyk Organization by team, product, or SCM to manage policies, reporting, and access
nav_context: agnostic
---

# Plan your Organization structure

{% hint style="success" %}
**Key decision:** Select the structure (team, product, or SCM-based) that best supports how you want to manage policies, report vulnerabilities, and define user access.
{% endhint %}

Snyk Organizations (Orgs) are the primary level for managing Projects and user access. Align your Org structure with your business, development, or security reporting needs.

| Structure                  | Description                                                                                               | Recommended for                                                                               | Avoid if                                                                            |
| -------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| **Team-based**             | Linked to a specific developer or security team.                                                          | Strict security ownership and reporting by team.                                              | Multiple teams need access to the exact same set of Projects (leads to complexity). |
| **Product-based**          | Set up for distinct applications or product lines (for example, Product A front end, Product A back end).   | Product-centric reporting and access control.                                                 | You have hundreds of small, unrelated products.                                     |
| **SCM organization-based** | Mimics your SCM hosting platform structure (for example, GitHub organizations, GitLab groups).            | Automated onboarding using the Snyk API, when you have more than 10 SCM-related integrations. | Your SCM organization names are not intuitive or do not align with business units. |
