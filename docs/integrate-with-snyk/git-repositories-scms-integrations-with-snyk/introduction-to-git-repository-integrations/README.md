# Introduction to Git repositories integrations

You can use Snyk functions to secure your application code, using [Snyk Code](../../../scan-with-snyk/snyk-code/) for your own code and [Snyk Open Source](../../../scan-with-snyk/snyk-open-source/) for the open-source libraries you use.

<figure><img src="https://lh5.googleusercontent.com/DfNV0u45H2pscEybpGNWSBLFPbIUe-Tp-75iTNUnNFJQvkpow1pWr07HgWkzaE31f1XdH9wQfijKStwEyDIvF93J6rD0E9aWbrVeBEUQuh111VpnHssNuS0FGCQ-ugaSp3OYUz_fMwRjbZNQVbjvdYp0CYaQQyyEq4NoXCFda3HLtTc5WBVkKJ_emw" alt="Snyk functions to secure application code"><figcaption><p>Snyk functions to secure application code</p></figcaption></figure>

These functions are available at each point in the development process to test, fix, and monitor your code, as explained in the following sections.

{% hint style="info" %}
For additional introductory information about Git integrations, see the other articles in this section of the documentation and [Clone an integration across your Snyk Organizations](../../../enterprise-configuration/snyk-broker/clone-an-integration-across-your-snyk-organizations.md) in the Snyk Broker documentation.
{% endhint %}

## While developing

Developers can use Snyk to check for issues while writing code locally on their machines before any code changes are pushed to the central Git repository.

Developers can test, fix, and monitor using:

* [Snyk CLI](../../../snyk-cli/) for local checks on the developer’s machine.\
  **Training**: see [Introduction to the Snyk CLI](https://learn.snyk.io/lesson/snyk-cli/)​.
* [Snyk for IDEs](../../use-snyk-in-your-ide/) plugins for checks embedded in the developer IDEs.\
  **Training**: see [Introduction to using Snyk in an IDE](https://learn.snyk.io/lesson/snyk-in-an-ide/).

## At code merge

When developers merge their code changes into their Git repository, Snyk can:

* **Run PR Checks**: scan for issues when a pull request (PR) is merged. By default, PR Checks act to ensure that the attack surface of the application never increases, only failing when a PR adds a dependency with issues.
* **Run daily scans**: have Snyk, by default, run daily scans if you imported Snyk Projects from your repo, to find any new problems in your current libraries quickly, such as critical zero-day vulnerabilities. This scanning occurs for all imported Projects, whether or not your teams are currently working on them. See [Walkthrough: code repository Projects](../../../implement-snyk/walkthrough-code-repository-projects/).
* **Create Jira tickets**: manage work on new issues discovered, to assign this work to developers in your team, and track progress on these issues. See the [Jira integration](../../notification-and-ticketing-systems-integrations/jira-integration.md) documentation.

## Automatically fix

Snyk can also suggest fixes by creating a PR to address a vulnerability, address older dependencies, and help address backlogged vulnerabilities over time. See [Fix your first vulnerability](../../../implement-snyk/walkthrough-code-repository-projects/fix-your-first-vulnerability.md).

## During build

Snyk can again act as a “gate” when you are building the application, checking that the code is secure at this stage by checking for issues automatically as part of the build process. This prevents a build from being completed based on policies as needed.

You can choose to report on issues, allowing the build to happen, or to stop the build if issues are encountered. You can also easily add criteria to this process, including exploitability, CVSS score, and whether a fix exists, thus focusing on fixing the issues that matter to you.

Snyk provides a number of [Snyk tools](../../../snyk-api/snyk-tools/) to help with this process.

## Security gate and Deploy

After passing through these points and the security gate, applications and code can be deployed to production: traditional and PaaS, Serverless, and Registry. You can get alerts or create Jira tickets when a new vulnerability is discovered and use Snyk's monitor function and other capabilities to maintain security.
