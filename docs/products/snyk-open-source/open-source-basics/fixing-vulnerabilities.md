# Fixing vulnerabilities

Snyk provides actionable fix advice for vulnerabilities: see [remediate-your-vulnerabilities.md](../../../features/fixing-and-prioritizing-issues/issue-management/remediate-your-vulnerabilities.md "mention") for more details.

Snyk supports workflows to fix vulnerabilities using:

* [Automatic pull / merge requests (PRs / MRs)](fixing-vulnerabilities.md#automatic-pull-merge-requests).
* [Manual pull / merge requests](fixing-vulnerabilities.md#manual-pull-merge-requests-for-a-project-code).

{% content-ref url="../../../features/fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/what-languages-do-we-support-fix-pull-requests-or-merge-requests.md" %}
[what-languages-do-we-support-fix-pull-requests-or-merge-requests.md](../../../features/fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/what-languages-do-we-support-fix-pull-requests-or-merge-requests.md)
{% endcontent-ref %}

### **Automatic pull / merge requests**

For projects imported via an SCM (Source Code Manager), Snyk offers the following types of automated pull / merge requests:

* [Fix pull requests for new vulnerabilities](https://docs.snyk.io/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities)
* [Fix pull requests to clear the backlog of vulnerabilities in priority order](fix-pull-requests-for-known-vulnerabilities-backlog.md)
* [Dependency upgrade pull requests](https://docs.snyk.io/snyk-open-source/dependency-management/upgrading-dependencies-with-automatic-prs)

### Manual pull / merge requests for a project code

To generate a PR / MR directly from your project, using the Snyk UI:

1. Navigate to your project from the project list
2. Select the file
3. Select **Open a Fix PR/MR** or **Fix this vulnerability**
4. A preview screen appears, showing you what fixes will be applied
5. Click **Open a Fix PR** on this screen to generate the pull request

![](../../../.gitbook/assets/image18.png)
