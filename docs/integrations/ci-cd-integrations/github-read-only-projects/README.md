# GitHub Read-only Projects

Snyk offers the ability to monitor a public GitHub repository that is not owned by your organization.

## How it works:

Adding a read-only project lets you track the vulnerabilities in a project you're considering using as a dependency, are already using as a stand-alone independent tool within your business, or any other public repository where you do not need to actively prevent or fix issues using Snyk's tool.

The repository is tested daily using your organization's GitHub credentials. These automated tests are not counted as part of the test limits related to your Snyk plan.

Unlike projects imported through the Snyk GitHub integration, projects that are imported or monitored with the read-only status cannot:

* Use automatic retesting when a pull request is merged
* Commit tests on any PR raised, to detect (and optionally block) new vulnerabilities from being introduced
* Use [automated fix PRs ](https://docs.snyk.io/products/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities)to recommend minimal changes to fix vulnerabilities
* Use [automated dependency upgrade PRs](https://docs.snyk.io/products/snyk-open-source/dependency-management/upgrading-dependencies-with-automatic-prs), to keep dependencies up to date and help avoid new vulnerabilities and simplify fixing those that are found.
* Use manual Fix PRs generated through Snyk to address specific issues chosen by the user

## Monitoring a public repository

You can import a read-only project via the **Add project** **> Monitor public GitHub repos** menu in the **Dashboard** and **Projects** tabs, or by going to [https://app.snyk.io/add/github-readonly](https://app.snyk.io/add/github-readonly).

![](../../../.gitbook/assets/screen\_shot\_2020-06-09\_at\_14.27.40.png)

1. Enter a public repository to monitor, following the format _owner/repository_.
2. When a valid repository name is entered, click **+ Add repo**. \
   The repository is quickly tested for a supported manifest file.
3. Enter the public repositories you want to monitor and click "Import X repository/ies".

![](<../../../.gitbook/assets/github\_readonly\_steps 2 & 3\_18july2022.png>)
