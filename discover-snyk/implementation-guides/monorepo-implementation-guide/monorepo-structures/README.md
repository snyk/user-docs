---
description: How to tailor your Snyk scanning strategy to the structure of your monorepo
nav_context: agnostic
---

# Monorepo structures

Monorepo structures vary based on organizational requirements, team scale, and deployment needs. Understanding which model your monorepo follows helps you tailor your CLI scanning strategy.

{% hint style="success" %}
**Key decision**: Identify which structure your monorepo most closely matches. If it does not fit neatly into one category, combine best practices from several.
{% endhint %}

The following structures are the four most common.

| Structure             | Description                                                                                                                                                                                                      |
| --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Functional/Layered** | Separates deployable applications from shared logic to ensure a stable, predictable foundation for multiple services.                                                                                           |
| **Domain-centric**    | Aligns code with business units by grouping all product components into a single directory, reducing cross-team friction. Often adopted as a company-wide single source of truth that houses multiple domains. |
| **App-centric**       | Organizes projects by runtime environment (for example, browser versus backend) to streamline CI/CD for specific deployment targets.                                                                             |
| **Flat/Namespace**    | A decentralized, peer-to-peer layout that relies on clear naming conventions for high visibility across the ecosystem.                                                                                           |

Select the structure that matches your monorepo:

1. [Functional/Layered monorepo](functional-layered-monorepo.md)
2. [Domain-centric monorepo](domain-centric-monorepo.md)
3. [App-centric monorepo](app-centric-monorepo.md)
4. [Flat/Namespace monorepo](flat-namespace-monorepo.md)
