# Snyk Secrets

The Snyk Secrets scanning tool provides accurate scanning across all repositories and includes governance features to prevent secret leaks.

Snyk Secrets provides the following features:

* Scans with low false-positive and high-recall rates using high entropy checks, machine learning (ML) semantic and contextual analysis, and regular expression (regex) pattern matching.
* Scans all plain text files (excluding binaries and lock files).
* Scans for secrets using the IDE (Visual Studio Code, Visual Studio, Eclipse and JetBrains), CLI (for local pre-commit and CI/CD scanning), and SCM (using recurring tests and scanning on import).
* Creates ignore requests to generate Secrets Consistent Ignores and governs the requests using the Ignore Approval Workflow.
* Creates [reports](https://docs.snyk.io/scan-fix-and-prevent/prevent/analytics/reports-tab) and analytics for findings and includes them when you export issue reports using the API.

{% hint style="info" %}
**Fake or AI generated secrets**

The Snyk Secrets scanner has low false-positive rates and will not detect any fake or AI generated secrets. Snyk recommends that you scan repos with genuine credentials.
{% endhint %}

## Prerequisites

Before you use Snyk Secrets, ensure you meet the following prerequisites:

* You must enable the Secrets setting on the organization level.
* Enable [Code Consistent Ignores](https://docs.snyk.io/scan-fix-and-prevent/fix/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code) to use Secrets Consistent Ignores and the Ignore Approval Workflow.
* Update your [IDE](https://docs.snyk.io/developer-tools/integrations/snyk-ide-plugins-and-extensions/release-and-support-policy-for-snyk-ide-plugins) and [CLI](https://docs.snyk.io/developer-tools/snyk-cli/snyk-cli/releases-and-channels-for-the-snyk-cli) to the latest versions to scan for secrets.
* Workspaces must be enabled to leverage Secret scanning. You can [read more](https://docs.snyk.io/developer-tools/integrations/scm-integrations/workspaces) on Workspaces and how to enable them.&#x20;

{% hint style="info" %}
To use Snyk Secrets, consider these system parameters and behaviors:

* Ignore requests are managed outside of pull request checks. For detailed scanning behavior, see [Secrets Pull Request checks](https://docs.snyk.io/scan-fix-and-prevent/prevent/pull-request-checks/secrets-pull-request-checks).
* Recurring tests do not generate email notifications.
* Snyk Secrets does not support binary and lock files.
* Repositories support up to 800 findings per scan. If an error alerts you that a repository exceeds this threshold, you can exclude the directories and files that hold known, non-production secrets by committing a `.snyk` file to your repository. Snyk applies these exclusions to both CLI and SCM scans. For details, see [Use the `.snyk` file with Snyk Secrets](../../manage-risk/policies/the-.snyk-file.md#use-the-.snyk-file-with-snyk-secrets). In the CLI, you can also use the `--exclude` option or scan sub-directories individually.
{% endhint %}
