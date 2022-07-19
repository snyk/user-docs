# GitHub Read-only Projects

Snyk offers the ability to monitor a public GitHub repository without granting any permissions through your own Snyk account.

## How it works:

This allows you to track the vulnerabilities in a project you are considering using as a dependency, are using as a stand-alone independent tool within your business, or any other public repository where you do not need to actively prevent or fix issues using Snyk's tool.

The repository will be automatically tested daily using Snyk's own GitHub credentials, and tests do not count against any test limits you may have.

As public projects imported in this way are not considered to be "yours", they do not appear in [Snyk's Reports section](https://app.snyk.io/reports).

Unlike projects imported through the Snyk GitHub integration, projects imported/monitored in this way are not eligible for:

* Automatic retesting when a pull request is merged
* Commit tests on any PR raised, to detect (and optionally block) new vulnerabilities from being introduced
* Automated fix PRs to recommend minimal changes to fix vulnerabilities - find out more
* Automated dependency upgrade PRs, to keep dependencies up to date and help avoid new vulnerabilities and simplify fixing those that are found.
* Manual Fix PRs generated through Snyk to address specific issues chosen by the user

Projects can be imported in this way both [during onboarding](github-read-only-projects.md), or [post-onboarding to Snyk](github-read-only-projects.md) (during normal ongoing use).

## During onboarding to Snyk:

During onboarding, selecting GitHub as the source from which to import projects allows you to set up the full GitHub integration (and take advantage of the vulnerability prevention and fix functionality available), or opt to proceed without granting Snyk permissions via the link at the bottom.

![](../../../.gitbook/assets/github\_readonly\_setup-1\_18july2022.png)

1. Enter a repository to monitor, following the format _owner/repository_.
2. When a valid repository name is entered, click **+ Add repo**. \
   The repository is quickly tested for a supported manifest file.
3. Enter the public repositories you want to monitor, and click "Import X repository/ies"

![](<../../../.gitbook/assets/github\_readonly\_steps 2 & 3\_18july2022.png>)

## Post-onboarding

Read-only projects can be imported regardless of whether a full GitHub integration is set up, because they don't rely on your GitHub permissions.

You can import a read-only project via the **Add project** **> Monitor public GitHub repos** menu in the **Dashboard** and **Projects** tabs, or by going to [https://app.snyk.io/add/github-readonly](https://app.snyk.io/add/github-readonly).

![](../../../.gitbook/assets/screen\_shot\_2020-06-09\_at\_14.27.40.png)

1. As in the previous section, enter a public repository to monitor, following the format _owner/repository_.
2. When a valid repository name is entered, click **+ Add repo**. \
   The repository is quickly tested for a supported manifest file.
3. Enter the public repositories you want to monitor and click "Import X repository/ies".

![](<../../../.gitbook/assets/github\_readonly\_steps 2 & 3\_18july2022.png>)
