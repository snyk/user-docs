# Enterprise implementation guide

Each business and environment is different. With that in mind, this guide aims to help an enterprise business to implement Snyk. The guide provides recommendations on implementing a large-scale rollout, focusing on the stages needed to help get to an ideal rollout.

The guide starts by recognizing that most businesses:

* Have a backlog of issues in their existing software.
* Are continuously creating new software and need to secure new code.

{% hint style="info" %}
There are **typical timelines for implementation** depending on the size and scope of your business.

If your business is small and nimble, Snyk implementation can be achieved in days. You can start scanning with Snyk soon after purchasing, often using a Git integration and the [API Import Tool](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/). See the [Getting started](../../discover-snyk/getting-started/) and [Start scanning](../../scan-with-snyk/start-scanning.md) sections for details of this type of process.

However, for larger, more process-oriented enterprises, the implementation process may take weeks or months and it requires more detailed planning to succeed.
{% endhint %}

The Snyk Essentials plan is included in the Snyk Enterprise plan, meaning that you will gain access to the following features:

* [Coverage control](../../manage-risk/policies/assets-policies/use-cases-for-policies/coverage-and-coverage-gap-policies.md) for SCM.
* [Policy](../../manage-risk/policies/assets-policies/) creation to automatically trigger specific actions.
* [Backstage file](../../developer-tools/scm-integrations/application-context-for-scm-integrations/) customization for SCM integrations.
* [Customized analytics](../../manage-risk/analytics/) and reports about the application.

{% hint style="info" %}
If you want to upgrade your plan to Snyk AppRisk, please contact your salesperson.\
On the [Snyk AppRisk ](broken-reference)page, you can find more details about the available features.
{% endhint %}

## Implementation strategy overview

This guide is composed of multiple phases, outlining a series of actions that align with three goals:

* [Achieve visibility](./#achieve-visibility)
* [Achieve prevention and drive developer adoption](./#achieve-prevention-and-drive-developer-adoption)
* [Fix the backlog and triage issues](./#fix-the-backlog-and-triage-issues)

### Achieve visibility

For large businesses, Snyk recommends that you first focus on visibility - getting a clear sense of the security issues, but without always fixing them.

{% hint style="info" %}
This does not stop you from fixing issues using Snyk. You can start fixing issues early, but the emphasis is to avoid blocking development early on, build trust, and slowly introduce gating in later phases, usually the prevention phase.
{% endhint %}

Visibility achieves a broad view of security across your application portfolio, prevents Snyk scans from being seen as a blocker, and minimizes impact on development processes.

This visibility helps build trust while rolling out Snyk.

### Achieve prevention and drive developer adoption

Next is the prevention stage; stopping new security issues from being added to your applications. During this stage, you can put controls in place to allow developers to see issues in their pipelines using Pull Request (PR)/Merge Request (MR) checks, and checks in the pipeline that may block.

As part of this, developers may use IDE plugins and other tools like [Snyk Advisor](https://snyk.io/advisor) to select secure packages and [Snyk Learn](https://learn.snyk.io/) to educate on secure coding, security, and the product.

### Fix the backlog and triage issues

Finally, you can focus on fixing your backlog of security issues. This can take several forms:

* As part of the initial rollout, security or the initial stakeholder may triage the initial results for the existing portfolio, create tickets for priority items to investigate or address, or have the teams do that for their applications as part of the weekly triage process.
* After getting visibility and achieving prevention, you can look at your backlog of issues. For example, a weekly triage process with the key stakeholders can guide the teams on what to address.

## Use enhanced resources with Snyk

Snyk was built with developers in mind, providing:

* Tools to create secure applications using integrations for IDE, Git, and CI/CD.
* [Snyk Advisor](https://snyk.io/advisor) and other tools to make decisions.
* [Snyk Learn](https://learn.snyk.io) training materials on products, securing code, and best practices.
* [Policies](../../manage-risk/policies/) that allow security and compliance teams to provide direction.
