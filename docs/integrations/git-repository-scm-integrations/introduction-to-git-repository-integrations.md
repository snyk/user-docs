# Introduction to Git repositories

You can use Snyk functions to secure your application code, using [Snyk Code](../../scan-application-code/snyk-code/) for your own code, and [Snyk Open Source](../../scan-application-code/snyk-open-source/) for the open-source libraries you use:

<figure><img src="https://lh5.googleusercontent.com/DfNV0u45H2pscEybpGNWSBLFPbIUe-Tp-75iTNUnNFJQvkpow1pWr07HgWkzaE31f1XdH9wQfijKStwEyDIvF93J6rD0E9aWbrVeBEUQuh111VpnHssNuS0FGCQ-ugaSp3OYUz_fMwRjbZNQVbjvdYp0CYaQQyyEq4NoXCFda3HLtTc5WBVkKJ_emw" alt="Snyk functions to secure application code"><figcaption><p>Snyk functions to secure application code</p></figcaption></figure>

These functions are available at the following points in the development process:

## While developing

Developers can use Snyk to check for issues while writing code locally (on their machines), before any code changes are pushed to the central Git repository.

Developers can use:

* [Snyk CLI](../../snyk-cli/) for local checks on the developer’s machine.\
  **Training**: see [Introduction to the Snyk CLI](https://training.snyk.io/courses/intro-cli)​.
* [Snyk for IDEs](../ide-tools/) plugins for checks embedded in the developer IDEs.\
  **Training**: see [Introduction to using Snyk in an IDE](https://training.snyk.io/courses/introduction-to-using-snyk-in-an-ide).

## At code merge

When developers merge their code changes into their Git repository, Snyk can:

* **Run PR Checks**: scan for issues when a pull request (PR) is merged. By default, PR Checks acts to ensure that the attack surface of the application never increases, only failing when a PR adds a dependency with issues.
* **Run daily scans**: have Snyk by default run daily scans, if you imported Snyk Projects from your repo, to find any new problems in your current libraries, for example to find critical zero-day vulnerabilities quickly. This scanning occurs for all imported Projects, whether or not your teams are currently working on them. See [Walkthrough: code repository projects](../../getting-started/walkthrough-code-repository-projects/).
* **Create Jira tickets**: manage work on new issues discovered, to assign this work to developers in your team and to track progress on these issues. See the [Jira integration](../notifications-ticketing-system-integrations/jira.md) documentation.

## Automatically fix

Snyk can also suggest fixes by creating a PR for a number of reasons: to address a vulnerability, to address older dependencies, and to help address backlogged vulnerabilities over time. See [Fix your first vulnerability](../../getting-started/walkthrough-code-repository-projects/fix-your-first-vulnerability.md).

## During build

Snyk can again act as a “gate” when you are building the application, checking that the code is secure at this stage to:

* Check for issues automatically as part of the build process
* Prevent a build from being completed based on policies

You can choose to either report on issues (allowing the build to happen) or to stop the build if issues are encountered. You can also easily add criteria to this process (including exploitability, CVSS score, whether a fix exists), to focus on fixing the issues that matter to you.

Snyk also provides a number of [Snyk tools](../../snyk-api-info/other-tools/) to help with this process.
