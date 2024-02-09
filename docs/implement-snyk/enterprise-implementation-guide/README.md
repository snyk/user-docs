# Enterprise implementation guide

Each business and environment is different. With that in mind, this guide aims to help an enterprise business to implement Snyk. The guide provides recommendations on implementing a large-scale rollout, focusing on the stages needed to help get towards an ideal rollout.

The guide starts by recognizing that most businesses:

* Have a backlog of issues in their existing software.
* Are continuously creating new software, and need to secure new code.&#x20;

{% hint style="info" %}
There are **typical timelines for implementation** depending on the size and scope of your business.

If your business is small and nimble, Snyk implementation can be achieved in days. You can start scanning with Snyk soon after purchasing, often using a Git integration and the [API Import Tool](../../snyk-api-info/other-tools/tool-snyk-api-import/). See the [Getting started](../../getting-started/) and [Start scanning](../../scan-with-snyk/start-scanning-using-the-cli-web-ui-or-api.md) sections for details of this type of process

However, for larger, more process-oriented enterprises, the implementation process may take weeks or months and it requires more detailed planning to succeed.&#x20;
{% endhint %}

If you want to enable the Application Security Posture Management for all your used products, you can use the [Snyk AppRisk Implementation Guide](../../manage-risk/snyk-apprisk/implementation-guide-for-snyk-apprisk/).

## Implementation strategy overview

This guide is composed of multiple phases, outlining a series of actions that align with three goals:

* [Achieve visibility](./#achieve-visibility)
* [Achieve prevention and drive developer adoption](./#achieve-prevention-and-drive-developer-adoption)
* [Fix the backlog and triage issues](./#fix-the-backlog-and-triage-issues)

### Achieve visibility

For large businesses, we recommend you first focus on visibility - getting a clear sense of the security issues, but without always fixing them.

{% hint style="info" %}
This does not stop you from fixing issues using Snyk. You can start fixing issues early, but the emphasis is to avoid blocking development early on, build trust, and slowly introduce gating in later phases, usually the prevention phase.
{% endhint %}

Visibility achieves a broad view of security across your application portfolio, avoids Snyk scans being seen as a blocker, and minimizes impact on development processes.&#x20;

This visibility helps build trust while rolling out Snyk.&#x20;

### Achieve prevention and drive developer adoption

Next is the prevention stage; stopping new security issues from being added to your applications. During this stage, you can put controls in place to allow developers to see issues in their pipelines using Pull Request (PR)/Merge Request (MR) checks, and checks in the pipeline that may block.&#x20;

As part of this, developers may use IDE plugins and other tools like [Snyk Advisor](https://snyk.io/advisor) to select secure packages, and [Snyk Learn](https://learn.snyk.io/) to educate on secure coding, security, and the product.

### Fix the backlog and triage issues

Finally, you can focus on fixing your backlog of security issues. This can take several forms:

* As part of the initial rollout, security or initial stakeholder may triage the initial results for the existing portfolio, create tickets for priority items to investigate or address, or have the teams do that for their applications as part of the weekly triage process.
* After getting visibility and achieving prevention, you can look at your backlog of issues.  For example, a weekly triage process with the key stakeholders can guide the teams on what to address.

## Use enhanced resources with Snyk

Snyk was built with developers in mind, providing:

* Tools to create secure applications using integrations for IDE, Git, and CI/CD.
* [Snyk Advisor](https://snyk.io/advisor) and other tools to make decisions.
* [Snyk Learn](https://learn.snyk.io) training materials on products, securing code, and best practices.&#x20;
* [Policies](../../scan-with-snyk/policies/) that allow security and compliance teams to provide direction.
