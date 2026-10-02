---
description: How to enable and test Snyk PR Checks, which block pull requests that introduce new vulnerabilities
nav_context: classic
---

{% include "../../../.gitbook/includes/new-navigation-banner.md" %}

# Test PR Checks

{% include "../../../.gitbook/includes/pilot-guide-navigation.md" %}

{% hint style="info" %}
Enabling PR Checks blocks Pull Requests that introduce new vulnerabilities. Snyk recommends including developers in the decision to enable PR Checks on actively developed repositories.
{% endhint %}

## Enable PR Checks

Follow these steps to enable the PR Checks feature:

* Open the Snyk Web UI
* Navigate to the Organization-level
* Open **Integrations**
* Select the Settings icon of your integration

* Navigate to the **Snyk PR Status Checks** option. Enable it for both Open Source and Code, and define fail conditions for each of them.
* Save the changes and apply them to all overridden Projects if you have already imported your repositories.
* Enable inline comments for a more integrated developer experience. See the [Pull Request experience ](https://docs.snyk.io/scan-fix-and-prevent/prevent/pull-request-checks/pull-request-experience)page for more details.

## Use PR Checks

After PR Checks are enabled, you will begin to see new PRs decorated with three additional Snyk checks:

* code/snyk: Snyk Code vulnerabilities
* license/snyk: Open Source license issues
* security/snyk: Open Source vulnerabilities

Try introducing a vulnerability in the PR so that you can walk through a scenario where a check fails. From the PR, clicking on the details of any failed check presents the list of issues that have been introduced in the PR.

Click into the full details of the issue to better understand the vulnerability and get remediation advice. In the situation where you need to force the check to pass, either to get a hotfix out or if you accept the risk, you can click **Mark as successful in SCM**. This button is only available to Org Admins. For Org Collaborators, this option is grayed out. See the [User role management](https://docs.snyk.io/platform-administration/user-management/user-role-management) page for more details.

With ‌inline comments enabled, you will see comments added for each Snyk Code vulnerability identified.

Try fixing the vulnerability in a follow-up commit, push the commit, and verify that the Snyk checks re-run and are passing.

{% hint style="info" %}
Additional Resources

* [Snyk PR Checks](https://docs.snyk.io/scan-fix-and-prevent/prevent/pull-request-checks)
* [Product Training: Snyk PR Checks in your PR/MR](https://learn.snyk.io/lesson/checking-your-code-with-pr-checks/?ecosystem=general)
{% endhint %}
