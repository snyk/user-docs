# Fix vulnerabilities with Snyk Open Source

{% hint style="info" %}
See [fix-your-vulnerabilities.md](../../../../fixing-and-reporting-issues/starting-to-fix-vulnerabilities/fix-your-vulnerabilities.md "mention") for more details of general fix functions.
{% endhint %}

Snyk provides actionable fix advice for vulnerabilities in your open source libraries, using:

* [Automatic pull / merge requests (PRs / MRs)](./#automatic-pull-merge-requests).
* [Manual pull / merge requests](./#manual-pull-merge-requests-for-a-project-code).

{% content-ref url="../../../../fixing-and-reporting-issues/starting-to-fix-vulnerabilities/what-languages-do-we-support-fix-pull-requests-or-merge-requests.md" %}
[what-languages-do-we-support-fix-pull-requests-or-merge-requests.md](../../../../fixing-and-reporting-issues/starting-to-fix-vulnerabilities/what-languages-do-we-support-fix-pull-requests-or-merge-requests.md)
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

![](../../../../.gitbook/assets/image18.png)
