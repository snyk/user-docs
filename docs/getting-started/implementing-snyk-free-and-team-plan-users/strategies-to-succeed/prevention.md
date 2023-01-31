# Prevention

## Introduction to prevention

Prevention aims to ensure new code is written right the first time, to prevent new code changes from introducing new security issues.

To operationalize security (adopt good security as standard) in your development workflow, you need to adopt:

* **Prevention**: avoid security issues with new development changes.
* **Backlog**: address issues you already have in your codebase (see [Addressing the Backlog](addressing-the-backlog.md)).

### Achieving prevention

Prevention can be achieved using:

* [Snyk Learn](https://learn.snyk.io/) teaches developers how to stay secure, with interactive lessons exploring vulnerabilities across a variety of languages and ecosystems.
* [Snyk Advisor](https://snyk.io/advisor/) helps developers by informing them and helping them choose better open-source packages and containers to base their development work on.
* [Snyk for IDEs](../../../ide-tools/) and [Snyk CLI](../../../snyk-cli/) allow developers to test code locally, before they submit changes.
* [PR Checks](../../../scan-application-code/run-pr-checks/): turn these on your most important projects to start, this gives visibility without developers having to do anything extra. As adoption grows, add PR Checks on all your projects.
* [Snyk CI/CD integration](../../../integrations/ci-cd-integrations/): initially use [snyk monitor](../../../snyk-cli/commands/monitor.md) for visibility, then use [snyk test](../../../snyk-cli/commands/test.md) to prevent vulnerabilities from being introduced. \
  **Note**: we recommend initially configuring your CI pipelines so that [snyk test](../../../snyk-cli/commands/test.md) does not block your builds. You can adjust this later when you are comfortable with the configuration.
  * Focus on Critical, High, and other criteria, like is it fixable, as needed. This is achieved via parameters you can pass to the CLI & CI/CD plugins.

{% hint style="info" %}
If a Pull Request (PR) is blocked, you must have a designated Snyk admin who can override this block, if the business decides to allow that PR to progress.
{% endhint %}

### What's next?

You can now start [addressing the backlog](addressing-the-backlog.md).
