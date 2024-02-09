# Plan Organization structure

## Introduction to structure

<div align="left">

<figure><img src="../../../.gitbook/assets/determine-account-structure.png" alt="Determine your account structure" width="563"><figcaption><p>Determine your account structure</p></figcaption></figure>

</div>

Snyk uses a hierarchical approach to managing assets, access, and rollup reporting.

* **A Snyk Group**: the top level
  * Typically this represents the account and is the company name.
  * The largest customers may have multiple groups representing each of their companies.
* **Snyk Organizations**: below the Group level, typically representing
  * Line of business
  * Git Organization or team structure
  * Types of application
  * Dev teams
* **Snyk Projects**
  * The import you have tested and monitored with Snyk.
  * A CLI scan that is being monitored, a container that is being monitored in a registry or Kubernetes, or an asset detected during a Git import, such as an open-source manifest, a code analysis scan, or an infrastructure as code file detected.

For more details, see [Manage Groups and Organizations](../../../snyk-admin/manage-groups-and-organizations/).

## Decide the structure

Deciding your structure is one of the earliest decisions you must make. When determining your account structure, structure Snyk Organizations to match your business needs.

Consider various factors when structuring Organizations, such as:

* Team-Based Structure: Link Developer Team A with Organization Team A.
* Product-Based Structure:
  * Set up separate Organizations for distinct parts of an application, like Payment Front-End and Payment Back-End.
  * Allow full-stack developers to join multiple organizations as needed.
* Git organization-based. : Some companies use a structure that mimics the organizations in Git, typically seen if a customer has 10+ organizations in their Git platform.

{% hint style="info" %}
If you intend to use the [api-import-tool](../../../snyk-api-info/other-tools/tool-snyk-api-import/), the Git organization-based approach will be your path forward.
{% endhint %}

By exploring these options and grasping your company's requirements, you can establish an account structure that enhances collaboration and efficiency.

## Additional information

See [Manage Organizations](../../../snyk-admin/manage-groups-and-organizations/manage-organizations.md) for more details.&#x20;
