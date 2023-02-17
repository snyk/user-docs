# Configure integrations

## Introduction

You must give Snyk access to your environment, to allow Snyk to scan that environment. The type of integration you need depends on what systems you use, and what you want to scan. See [Integrate with Snyk](../../../integrations/) for more details.

This section describes some initial integrations configuration tips and suggestions.

### Initially disable global PR tests

As part of your rollout process, you may decide to globally disable Snyk test for pull request checks at the start and only enable them on the most important projects at first.

Under **Settings > Integrations**, **** select **Disable Pull Request check** at the **Git integration** screen for your Git platform/tool:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-02-01 at 11.35.03.png" alt=""><figcaption><p>Initially disable PR checks</p></figcaption></figure>

#### Enable tests for specific Projects

For specific Snyk Projects where you want to enable these tests, navigate to the Project settings and enable Snyk test for pull request checks **** for that Project:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-02-01 at 11.41.56.png" alt=""><figcaption><p>Enable tests for specific Snyk Projects</p></figcaption></figure>

See [Test your PRs for vulnerabilities before merging](../../../scan-application-code/run-pr-checks/pr-checks-for-snyk-open-source/test-your-prs-for-vulnerabilities-before-merging.md) for more details.

### Configure Jira integration

For Team Plan users, you can configure [Jira Integration](https://docs.snyk.io/integrations/notifications-ticketing-system-integrations/jira) to assist with logging tickets and addressing backlogged security issues.
