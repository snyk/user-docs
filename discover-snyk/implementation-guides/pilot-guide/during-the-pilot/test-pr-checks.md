---
description: How to enable and test Snyk PR Checks, which block pull requests that introduce new vulnerabilities
nav_context: new
---

# Test PR Checks

{% include "../../../.gitbook/includes/pilot-guide-navigation.md" %}

{% hint style="info" %}
Enabling PR Checks blocks Pull Requests that introduce new vulnerabilities. Snyk recommends including developers in the decision to enable PR Checks on actively developed repositories.
{% endhint %}

### Enable PR Checks

1. In the Snyk Web UI, use the scope selector at the top of the page to select your Organization.
2. Navigate to **Settings** > **Integrations**, then select your SCM integration, for example, **GitHub**.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXey76C-t0VJCjNUT9sOKfcbwxZR0mzyka0AMKwdaL1Sbp8HwS_rI0mRsU0maIyAe5zjeHfcMKkDZ9k_MguVPwddry4-a3MbBE_cdb1xJoR5Q5rx7SgCsbjJAzYEgxRcU-B5XeMFpg?key=i_CNrr-DvB8PGUAzq09BT3pc" alt="Snyk PR Status Checks option in the integration settings"><figcaption></figcaption></figure>

3. In the **Snyk PR Status Checks** section, enable PR Checks for both Open Source and Code, and define fail conditions for each.
4. Save the changes. If you have already imported repositories, apply the changes to all overridden Projects.
5. Enable inline comments for a more integrated developer experience. For details, visit [Pull Request experience](https://docs.snyk.io/scan-fix-and-prevent/prevent/pull-request-checks/pull-request-experience).

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXfNXo0IULol0ix0VcJ34oOd87JGOdtq4g49PyoUx_pVFpqj5E1GSz0j8Atiu0Ehyk6APwTHfx6xNPqa5ye9-2w9YEMSUwiAhpw0yFEVaecvalkF4eXQz01inGYGPSGEPJvUuIaWDA?key=i_CNrr-DvB8PGUAzq09BT3pc" alt="PR Checks enabled for Open Source and Code"><figcaption></figcaption></figure>

### Use PR Checks

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
