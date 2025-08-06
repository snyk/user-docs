# Phase 1: Discovery and planning

## Discovery phase steps&#x20;

* [Connect with Snyk](connect-with-snyk.md): Connect with your Snyk account manager.
* [Conduct discovery](conduct-discovery.md): Identify stakeholders, integrations, and applications to monitor.
* [Plan Organization structure](plan-organization-structure.md): Decide how to organize and control access to content using the Snyk Organization structure.
* [Determine user roles](determine-member-roles.md): Decide whether your users need customized access to Snyk.
* [Decide SSO access](decide-sso-access.md): Identify initial Single Sign-On (SSO) settings.
* [Plan for success](plan-for-success.md): Decide how to judge the success of your rollout.
* [Choose rollout integrations](choose-rollout-integrations.md): Decide which integrations to implement initially.
* [Create rollout plan](create-rollout-plan.md): Create a high-level plan for rolling out Snyk in your business.

## General pre-rollout questions

Some initial questions follow that you can ask to assist in planning before starting the rollout. This is one way of deciding on your rollout process.

### Who is involved?

* Who will manage and oversee the project?
* Who will champion Snyk?
* Who will be the Group Administrator?

### What are your goals?

* Why did you choose Snyk?&#x20;
* Why are you implementing it now?

### How will your users use Snyk?&#x20;

How will you provision users and integrate Snyk with your platforms?

* Who will need access to Snyk?&#x20;
* What will they need access to?&#x20;
* Will access be restricted to certain Projects?
* Who can grant Snyk access to platforms like SSO and Git repositories?

### How will you structure your account?

* How will you group your Projects?
  * By developer teams?&#x20;
  * By product?&#x20;
  * By business unit?
* If by developer teams, are there some teams that would need access to the same Projects? If yes, think about a different structure to avoid confusion.
* How many [Snyk Organizations](../../../snyk-admin/groups-and-organizations/organizations/) do you need?&#x20;

### How will you measure success?&#x20;

* What KPIs will be tracked?
* How will you know that you are making progress?
* Are there key development projects that progress tracking should be aligned with, or at least included in the tracking, to measure progress against?

## Snyk Essentials considerations

Snyk Essentials is part of the Snyk Enterprise offering, and it provides discovery and visibility for your application assets and security tool coverage.&#x20;

When or before you use Snyk Essentials, you should consider the following items:

* Who would want coverage visibility or is accountable if an important application is not being monitored by security tools?
* Who would you notify, using automated policies, if a repository were missing coverage by a security tool?
* Are you using [Application context](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/), and are there fields that might be valuable in automating policies with Snyk Essentials?
  * Is it possible to categorize important applications in Git or CMDB (ServiceNow) using topics or fields, such as a PCI topic or tag?
  * Would this also reduce noise about test applications and internal applications by implementing an internal tag, topic, or naming convention?&#x20;
* Read the available examples of [common policies](../../../manage-risk/policies/assets-policies/#use-cases) that can be created using Snyk Essentials.&#x20;
