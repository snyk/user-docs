---
description: How to identify business-critical applications, stakeholders, and success metrics before rolling out Snyk
nav_context: agnostic
---

# Discovery

## Identify business-critical applications

Identifying key applications early helps you find important contacts, define success metrics, and prioritize early. Even if your business has thousands of applications, can you identify a few key applications to benchmark progress and priority?

Large enterprises can import everything and collect this additional information in parallel to help prioritize work and measure success. This step does not need to block or delay the rollout. You can do it in parallel during the planning and implementation phases.

## Confirm points of contact and roles in small companies

To successfully implement Snyk, you must identify the skills needed and the stakeholders who will be involved. For example, you must identify people who can:

* Perform the initial setup and admin tasks.
* Trial the workflows and functionality as developers.
* Generate tokens with the necessary permissions for Git repositories and other integrations you are considering, such as Jira or container registries.
* Modify pipeline scripts.
* Provide input on open-source licensing, if licensing was a priority for purchasing. Otherwise, the default policy is often sufficient.
* Receive progress reports as stakeholders.

## Identify stakeholders in an enterprise with a RACI matrix

In smaller companies where a few individuals have the necessary access, involving the stakeholders is quick. For a small team or early startup, it can be as simple as finding who can create credentials for Snyk to access the system or update a CI/CD script. In a larger enterprise where only a single team uses Snyk, this tends to be more organizationally complex. A responsible, accountable, consulted, and informed (RACI) matrix helps you determine who is:

* Responsible: the person who does the work to complete the task or deliverable.
* Accountable: the approver who must sign off on work before it is considered complete.
* Consulted: those whose opinions are sought, two-way communication.
* Informed: kept up to date on progress, one-way communication.

For large enterprises, the following RACI matrix is useful to clearly define roles and responsibilities during rollout:

<table><thead><tr><th width="179">Task</th><th width="136">Champion</th><th width="146">Admin</th><th width="132">Security</th><th>DevOps</th></tr></thead><tbody><tr><td>Onboarding</td><td>Responsible</td><td>Responsible</td><td>Responsible</td><td>Responsible</td></tr><tr><td>Admin Training</td><td>Accountable</td><td>Responsible</td><td>Consulted</td><td>Responsible</td></tr><tr><td>Security Training</td><td>Responsible</td><td>Consulted</td><td>Accountable</td><td>Responsible</td></tr><tr><td>DevOps Training</td><td>Responsible</td><td>Consulted</td><td>Consulted</td><td>Accountable</td></tr><tr><td>Source Control/IDE/Pipeline Setup</td><td>Responsible</td><td>Responsible</td><td>Responsible</td><td>Accountable</td></tr><tr><td>License Policy Management</td><td>Responsible</td><td>Responsible</td><td>Responsible</td><td>Accountable</td></tr><tr><td>Security Triage</td><td>Responsible</td><td>Consulted</td><td>Accountable</td><td>Consulted</td></tr></tbody></table>

This ensures all stakeholders are engaged in suitable capacities without duplicated or neglected efforts. Following the matrix enables smooth collaboration, with individuals contributing specialized skills within their designated areas.

The clear delineation of duties promotes productivity, efficiency, and accountability. Overall, the RACI framework is an effective method for orchestrating a structured, well-coordinated onboarding process resulting in the successful implementation of Snyk.
