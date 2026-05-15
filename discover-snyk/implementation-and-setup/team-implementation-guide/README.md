# Team implementation guide

Accelerate your team performance by using Snyk. This guide aims to help you implement Snyk for your team. The team plan applies to teams of up to 10 members.&#x20;

We start with the awareness that most businesses:

* Have a backlog of issues in their existing software
* Are continuously creating new software and need to secure new code.&#x20;

## **Typical timelines**

Once your Snyk Organization is set up, you can immediately start gaining visibility into your code by integrating your code repositories (with PR checks disabled), CI/CD pipelines, or container registries.

To minimize disruption, Snyk recommends a gradual rollout of your "shift left" security strategy, focusing on backlog remediation and prevention. Key recommendations include providing developers with access to the IDE and piloting on a single Project before expanding best practices to the wider team.

## Implementation strategy overview

This guide is composed of multiple phases, outlining a series of actions configuring your account, as well as actions outside the system, that align with the following goals:

* [Achieve visibility](./#achieve-visibility)
* [Achieve prevention and drive developer adoption](./#achieve-prevention-and-drive-developer-adoption)
* [Fix the backlog and triage issues](./#fix-the-backlog-and-triage-issues)

### Achieve visibility

If you focus on visibility first, you can get a clear sense of the security issues, but without always fixing them.

{% hint style="info" %}
This does not stop you from fixing issues using Snyk. You can start fixing issues early, but the emphasis is to avoid blocking development early on, build trust, and slowly introduce gating in later phases, usually the prevention phase. This is true of the smallest or largest teams - communication is key.
{% endhint %}

Visibility achieves a broad view of security across your application portfolio, avoids Snyk scans being seen as a blocker, and minimizes impact on development processes.&#x20;

This visibility helps build trust while rolling out Snyk. With the Team plan, this equates to onboarding your projects through Git repository and disabling PR Checks/Auto PRs in the integration settings. Choose an important project and enable PR checks after communicating with the relevant team members. This guide details this later on.

### Achieve prevention and drive developer adoption

Next is the prevention stage. You should stop new security issues from being added to your applications. During this stage, you can put controls in place to allow developers to see issues in their pipelines using Pull Request (PR)/Merge Request (MR) checks, and checks in the pipeline that may block.&#x20;

As part of this, developers may use IDE plugins and other tools like [Snyk Advisor](https://snyk.io/advisor) to select secure packages and [Snyk Learn](https://learn.snyk.io/) to educate on secure coding, security, and the product. It's quite common to see developers download and use IDE plugins. Provide guides indicating the settings they should use and guidelines on what they should fix to start often Criticals and Highs, where fixes are available.

### Fix the backlog and triage issues

Finally, you can focus on fixing your backlog of security issues. This can take several forms:

* As part of the initial rollout, security or initial stakeholder may triage the initial results for the existing portfolio, create tickets for priority items to investigate or address, or have the teams do that for their applications as part of the weekly triage process.
* After getting visibility and achieving prevention, you can look at your backlog of issues.  For example, a weekly triage process with the key stakeholders can guide the teams on what to address.

## Use enhanced resources with Snyk

Snyk was built with developers in mind, providing:

* Tools to create secure applications using integrations for IDE, Git, and CI/CD.
* [Snyk Advisor](https://snyk.io/advisor) and other tools to make decisions.
* [Snyk Learn](https://learn.snyk.io) training materials on products, securing code, and best practices.&#x20;
