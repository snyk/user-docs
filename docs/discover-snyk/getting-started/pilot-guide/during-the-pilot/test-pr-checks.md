# Test PR Checks

{% include "../../../../.gitbook/includes/pilot-guide-toc.md" %}

{% hint style="info" %}
Enabling PR Checks blocks Pull Requests that introduce new vulnerabilities. Snyk recommends including developers in the decision to enable PR Checks on actively developed repositories.
{% endhint %}

## Enable PR Checks

Follow these steps to enable the PR Checks feature:

* Open the Snyk Web UI
* Navigate to the Organization-level
* Open **Integrations**
* Select the Settings icon of your integration

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXey76C-t0VJCjNUT9sOKfcbwxZR0mzyka0AMKwdaL1Sbp8HwS_rI0mRsU0maIyAe5zjeHfcMKkDZ9k_MguVPwddry4-a3MbBE_cdb1xJoR5Q5rx7SgCsbjJAzYEgxRcU-B5XeMFpg?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

* Navigate to the **Snyk PR Status Checks** option. Enable it for both Open Source and Code, and define fail conditions for each of them.&#x20;
* Save the changes and apply them to all overridden Projects if you have already imported your repositories.&#x20;
* Enable inline comments for a more integrated developer experience. See the [Pull Request experience ](../../../../scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience.md)page for more details.&#x20;

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXfNXo0IULol0ix0VcJ34oOd87JGOdtq4g49PyoUx_pVFpqj5E1GSz0j8Atiu0Ehyk6APwTHfx6xNPqa5ye9-2w9YEMSUwiAhpw0yFEVaecvalkF4eXQz01inGYGPSGEPJvUuIaWDA?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

## Use PR Checks

After PR Checks are enabled, you will begin to see new PRs decorated with three additional Snyk checks:

* code/snyk: Snyk Code vulnerabilities
* license/snyk: Open Source license issues
* security/snyk: Open Source vulnerabilities\


<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXfLnXujzF5V7LkTvXBaKbFsnRtrc2HGhM37B1Ij58a8O_UcLNamCOEf6aMckN27doPl2Vj3RUpdcWlAphmKKGt7QpNKTzcP62QaD3jITRBe8kta0whoIyp5Tk660eScfa6XwzVM?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

Try introducing a vulnerability in the PR so that you can walk through a scenario where a check fails. From the PR, clicking on the details of any failed check presents the list of issues that have been introduced in the PR.

Click into the full details of the issue to better understand the vulnerability and get remediation advice. In the situation where you need to force the check to pass, either to get a hotfix out or if you accept the risk, you can click **Mark as successful in SCM**. This button is only available to Org Admins. For Org Collaborators, this option is grayed out. See the [User role management](../../../../snyk-platform-administration/user-roles/user-role-management.md) page for more details.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXeoGMM6QgygOIslguUEnVLg4cDB7CaerOKKY44EjPjgD6GL3znT8QS2-9-T8G5d6MEkLA4fPaGTth5saE_UuB4wiA08Wnep-RPFr39_u1G6sm2Mr0XWQYyyCVMiBrgfZJ1k8DVqjQ?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

With ‌inline comments enabled, you will see comments added for each Snyk Code vulnerability identified.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXesXBwkre6Tst7YogJVBdoXPS3DwSe_D4i2zvLirtRjd7wwDGVcT07oNaaJ1PykRHUQir4xi0nRGtcEYAE96KtZiFrfA4HK7l4-rwVCUYp8SqS-xVkQ22b8eOh1QaSGQnpHeziz?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

Try fixing the vulnerability in a follow-up commit, push the commit, and verify that the Snyk checks re-run and are passing.

{% hint style="info" %}


Additional Resources

* [Snyk PR Checks](../../../../scan-with-snyk/pull-requests/pull-request-checks/)
* [Product Training: Snyk PR Checks in your PR/MR](https://learn.snyk.io/lesson/checking-your-code-with-pr-checks/?ecosystem=general)
{% endhint %}

