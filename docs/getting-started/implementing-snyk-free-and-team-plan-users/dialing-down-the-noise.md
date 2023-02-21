# Reduce the noise

## Introduction: what is "noise"?

In cybersecurity, "noise" is the volume of information you receive as part of security analysis, including scans, alerts and notifications. The higher this noise level is, the more difficult it becomes to prioritize the important work.

When you adopt Snyk, you should reduce this noise, to avoid distractions being raised for your teams, so allowing your team to focus on the areas that matter most.

### Reduce noise with Snyk

Snyk products reduce noise levels by default. You can additionally configure and adjust Snyk settings as needed to keep focus, especially as a company grows and the number of applications increases.

This section describes Snyk functions that can help you reduce noise, and help ensure your success and engagement.

### Manage notifications

You can manage notifications both for your Organization and for your personal account, to avoid being overwhelmed by too many email notifications. See [Managing Notifications](../../user-and-group-management/notifications.md) for more details.

* **Organization Notification Settings**\
  In your settings for the organization and its members, you can determine if alerts are only for High and Critical, this is very useful to focus on the worst issues. This does not impact reporting, only alerts.
*   **Personal Notifications Settings**

    Navigate to your profile, on the top right of the screen, select notifications. Here you can choose which projects you get notified for. This is critical to ensure you are not getting messages for projects you are not responsible for, and your developers will appreciate knowing they can set that.

### Use filtering functions

Snyk provides criteria-based filtering functions that can help. For example:

* PR checks can fail if there’s a Critical/High issue, and if a fix is available (See [Run PR Checks](../../scan-application-code/run-pr-checks/))
* CI/CD & CLI - You can filter on **--severity-threshold**, but also if a fix is available (See [CLI Help](../../snyk-cli/commands/)).
* CI/CD & CLI - Using the [Snyk-Filter](https://github.com/snyk-tech-services/snyk-filter) plugin, you can report on issues if there’s an exploit available, CVSS score, Priority score etc.

### Avoid failing the build initially

It is common not to fail the build in the first month and to only use the **Snyk monitor** command to start. When you do start failing the build, using the criteria shown above is really helpful.&#x20;

See [Snyk test and snyk monitor in CI/CD integration](../../integrations/ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md) for details.

### Do not set automatic PRs or PR Checks

If you have a large number of projects, initially we recommend you turn both of these off, for all but the important projects. As your teams become more accustomed to this type of capability, you can then slowly turn it on across your repositories.

See [Upgrading dependencies with automatic PRs](../../products/snyk-open-source/dependency-management/upgrading-dependencies-with-automatic-prs.md) and [Run PR Checks](../../scan-application-code/run-pr-checks/) for details.
