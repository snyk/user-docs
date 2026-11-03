---
description: Guide to implementing Snyk for a team on the Team plan, covering discovery, configuration, visibility, and a fix strategy
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Team implementation guide

Use Snyk to accelerate the performance of your team. This guide helps you implement Snyk for your team. The Team plan applies to teams of up to 10 members.

This guide assumes that most businesses:

* have a backlog of issues in their existing software
* are continuously creating new software and need to secure new code

## Typical timelines

After your Snyk Organization is set up, you can immediately start gaining visibility into your code by integrating your code repositories (with PR Checks disabled), CI/CD pipelines, or container registries.

To minimize disruption, Snyk recommends a gradual rollout of your "shift left" security strategy, focusing on backlog remediation and prevention. Key recommendations include providing developers with access to the IDE and piloting on a single Project before expanding best practices to the wider team.

## Implementation strategy overview

This guide consists of multiple phases. The phases outline actions that configure your account and actions outside the system, aligned with the following goals:

* [Achieve visibility](./#achieve-visibility)
* [Achieve prevention and drive developer adoption](./#achieve-prevention-and-drive-developer-adoption)
* [Fix the backlog and triage issues](./#fix-the-backlog-and-triage-issues)

### Achieve visibility

If you focus on visibility first, you can get a clear sense of the security issues, but without always fixing them.

{% hint style="info" %}
This does not stop you from fixing issues using Snyk. You can start fixing issues early, but the emphasis is to avoid blocking development early on, build trust, and slowly introduce gating in later phases, usually the prevention phase. This is true of the smallest and largest teams — communication is key.
{% endhint %}

Visibility achieves a broad view of security across your application portfolio, avoids Snyk scans being seen as a blocker, and minimizes impact on development processes.

This visibility helps build trust while rolling out Snyk. With the Team plan, this means onboarding your Projects through an SCM integration and disabling PR Checks and automatic PRs in the integration settings. Choose an important Project and enable PR Checks after you communicate with the relevant team members. A later phase of this guide covers this in detail.

### Achieve prevention and drive developer adoption

Next is the prevention stage, where you stop new security issues from being added to your applications. During this stage, you can put controls in place so that developers see issues through pull request (PR) and merge request (MR) checks, and through pipeline checks that can block builds.

As part of this, developers can use IDE plugins and other tools such as [Snyk Advisor](https://snyk.io/advisor) to select secure packages, and [Snyk Learn](https://learn.snyk.io/) to learn about secure coding, security, and the product. Developers commonly download and use IDE plugins. Provide guides that indicate which settings to use and what to fix first, often Critical and High severity issues where fixes are available.

### Fix the backlog and triage issues

Finally, you can focus on fixing your backlog of security issues. This can take several forms:

* As part of the initial rollout, the security team or initial stakeholders can triage the initial results for the existing portfolio and create tickets for priority items to investigate or address. Alternatively, they can have the teams do this for their applications as part of the weekly triage process.
* After gaining visibility and achieving prevention, you can review your backlog of issues. For example, a weekly triage process with the key stakeholders can guide the teams on what to address.

## Use enhanced resources with Snyk

Snyk was built with developers in mind, providing:

* tools to create secure applications using IDE, SCM, and CI/CD integrations
* [Snyk Advisor](https://snyk.io/advisor) and other tools to make decisions
* [Snyk Learn](https://learn.snyk.io) training materials on products, securing code, and best practices
