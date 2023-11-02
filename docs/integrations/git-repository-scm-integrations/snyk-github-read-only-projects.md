# Snyk GitHub Read-only Projects

Snyk offers GitHub Read-only Projects, providing the ability to monitor a public GitHub repository that is not owned by your Organization.

## How GitHub Read-only Projects work

Adding a read-only Project lets you track the vulnerabilities in a Project you are considering using as a dependency, a Project you are already using as a stand-alone independent tool within your business, or any other public repository where you do not need to actively prevent or fix issues using Snyk.

The repository is tested daily using your Organization's GitHub credentials. These automated tests are not counted as part of the test limits of your Snyk plan.

Unlike Projects imported through the Snyk GitHub integration, Projects that are imported or monitored with the read-only status cannot do the following:

* Use automatic retesting when a pull request is merged.
* Commit tests on any PR raised to detect and optionally block new vulnerabilities from being introduced.
* Use [automated fix PRs](../../scan-using-snyk/snyk-open-source/open-source-basics/fix-pull-requests-for-new-vulnerabilities.md) to recommend minimal changes to fix vulnerabilities.
* Use [automated dependency upgrade PRs](../../scan-using-snyk/snyk-open-source/open-source-basics/upgrading-dependencies-with-automatic-prs.md) to keep dependencies up to date, avoid new vulnerabilities, and simplify fixing those that are found.
* Use manual Fix PRs generated through Snyk to address specific issues chosen by the user.

## How to monitor a public GitHub repository

You can import a read-only Project using the **Add project** **> Monitor public GitHub repos** menu in the **Dashboard** and **Projects** tabs, or by going to [Monitor public GitHub repositories](https://app.snyk.io/add/github-readonly).

<figure><img src="../../.gitbook/assets/screen_shot_2020-06-09_at_14.27.40.png" alt="Add project, Monitor public GitHub repos"><figcaption><p>Add project, Monitor public GitHub repos</p></figcaption></figure>

1. Enter a public repository to monitor, following the format _owner/repository_.
2. When you have entered a valid repository name, click **+ Add repo**.\
   The repository is quickly tested for a supported manifest file.
3. Enter the public repositories you want to monitor and select **Import N repository/ies**.

<figure><img src="../../.gitbook/assets/github_readonly_steps 2 &#x26; 3_18july2022.png" alt="Add repo and Import repository or repositories"><figcaption><p>Add repo and Import repository or repositories</p></figcaption></figure>
