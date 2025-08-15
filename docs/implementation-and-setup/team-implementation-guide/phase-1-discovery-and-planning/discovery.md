# Discovery

## Identify business-critical applications

Identifying key applications early helps you identify important contacts to make, helps define success metrics, and helps early prioritization. Your business may have thousands of applications, but can you identify a few key applications to benchmark progress and priority?

For large enterprises, you can import everything, but that additional information can be collected in parallel to help prioritize work and measure success. This should not be a blocker or delay, it can be done in parallel in the planning and implementation phases.

## For implementations in smaller organizations/startups - confirm internal points of contact and roles

To successfully implement Snyk, you must identify the skills needed and the stakeholders who will be involved. For example, you must identify people who can:

* Perform the initial setup/admin tasks.
* The developers who may want to trial the workflows/functionality.
* Generate tokens with the necessary permissions for git repositories, and other integrations being considered like Jira, or container registries.
* Modify pipeline scripts.
* Who would want input on open-source licensing if this was a priority for purchasing, otherwise the default policy is often sufficient.
* Any stakeholders that need to be reported to.

## For implementations for a Team plan within an Enterprise - identify stakeholders with a RACI matrix

In smaller organizations where a few individuals may have the necessary access, involving the stakeholders can be quick and easy. For a small team, early startup, it can be as simple as finding who can create credentials for Snyk to access the system or update a CI/CD script. In a larger enterprise, where you may have a single team using Snyk, this tends to be more organizationally complex, and the following RACI matrix can assist in determining:

* **R**esponsible: the person ultimately accountable for carrying out the task or deliverable.
* **A**ccountable: the approver who must sign off on work before it is considered complete.
* **C**onsulted: those whose opinions are sought, two-way communication.
* **I**nformed: kept up-to-date on progress, one-way communication.

For large enterprises, the following RACI matrix is useful to clearly define roles and responsibilities during rollout:

<table><thead><tr><th width="179">Task</th><th width="136">Champion</th><th width="146">Admin</th><th width="132">Security</th><th>DevOps</th></tr></thead><tbody><tr><td>Onboarding</td><td>Responsible</td><td>Responsible</td><td>Responsible</td><td>Responsible</td></tr><tr><td>Admin Training</td><td>Accountable</td><td>Responsible</td><td>Consulted</td><td>Responsible</td></tr><tr><td>Security Training</td><td>Responsible</td><td>Consulted</td><td>Accountable</td><td>Responsible</td></tr><tr><td>DevOps Training</td><td>Responsible</td><td>Consulted</td><td>Consulted</td><td>Accountable</td></tr><tr><td>Source Control/IDE/PIPELINE  Setup</td><td>Responsible</td><td>Responsible</td><td>Responsible</td><td>Accountable</td></tr><tr><td>License Policy Management</td><td>Responsible</td><td>Responsible</td><td>Responsible</td><td>Accountable</td></tr><tr><td>Security Triage</td><td>Responsible</td><td>Consulted</td><td>Accountable</td><td>Consulted</td></tr></tbody></table>

This ensures all stakeholders are engaged in suitable capacities without duplicated or neglected efforts. Following the matrix enables smooth collaboration, with individuals contributing specialized skills within their designated areas.&#x20;

The clear delineation of duties promotes productivity, efficiency, and accountability. Overall, the RACI framework is an effective method for orchestrating a structured, well-coordinated onboarding process resulting in the successful implementation of Snyk.
