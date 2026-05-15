# Structure your account

## Plan your Organization structure

{% hint style="success" %}
**Key decision:** Select the structure (team, product, or SCM-based) that best supports how you want to manage policies, report vulnerabilities, and define user access.
{% endhint %}

Snyk Organizations (Orgs) are the primary level for managing Projects and user access. Align your Org structure with your business, development, or security reporting needs.

| Structure                  | Description                                                                                               | Recommended for                                                                               | Avoid if                                                                            |
| -------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| **Team-based**             | Linked to a specific developer or security team.                                                          | Strict security ownership and reporting by team.                                              | Multiple teams need access to the exact same set of projects (leads to complexity). |
| **Product-based**          | Set up for distinct applications or product lines (for example, product a front-end, product a back-end). | Product-centric reporting and access control.                                                 | You have hundreds of small, unrelated products.                                     |
| **SCM Organization-based** | Mimics your SCM hosting platform structure (for example Github orgs, Gitlab groups).                      | Automated onboarding using the Snyk API, when you have more than 10 SCM-related integrations. | Your SCM organization names are not intuitive or don't align with business units.   |
