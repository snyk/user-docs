# Plan Organization structure

## Introduction to structure

Snyk uses a hierarchical approach to managing assets, access, and rollup reporting.

* **A Snyk Group**: the top level
  * Typically, this represents the account and is the company name.
  * The largest customers may have multiple groups representing each of their companies.
* **Snyk Organizations**: below the Group level, typically representing
  * Line of business
  * Git organization or team structure
  * Types of application
  * Dev teams
* **Snyk Projects**
  * The import you have tested and monitored with Snyk.
  * A CLI scan that is being monitored, a container that is being monitored in a registry or Kubernetes, or an asset detected during a Git import, such as an open-source manifest, a code analysis scan, or an infrastructure as code file detected.

For more details, see [Manage Tenants, Groups and Organizations](../../../snyk-platform-administration/groups-and-organizations/).

## Decide the structure

Deciding your structure is one of the earliest decisions you must make. When determining your account structure, structure Snyk Organizations to match your business needs.

Consider various factors when structuring Organizations, such as:

* Team-based structure: Link developer team A with Organization team A.
* Product-based structure:
  * Set up separate Organizations for distinct parts of an application, for example, Payment Front-End and Payment Back-End.
  * Allow full-stack developers to join multiple Organizations as needed.
* Git organization-based. : Some companies use a structure that mimics the organizations in Git, typically seen if a customer has 10-plus organizations in their Git platform.

{% hint style="info" %}
If you intend to use the [api-import-tool](../../../scan-with-snyk/snyk-tools/tool-snyk-api-import/), the Git organization-based approach will be your path forward.
{% endhint %}

By exploring these options and grasping your company's requirements, you can establish an account structure that enhances collaboration and efficiency.

## Additional information

See [Manage Organizations](../../../snyk-platform-administration/groups-and-organizations/organizations/create-and-delete-organizations.md) for more details.&#x20;
