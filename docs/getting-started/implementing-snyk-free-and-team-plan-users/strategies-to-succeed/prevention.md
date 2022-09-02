# Prevention

When you operationalize security as part of development, you can break down the problem into:

* **Backlog**: addressing issues you already have.
* **Prevention**: new development that might introduce new issues.

Prevention is about both ensuring new code is written right the first time, and not introducing new vulnerabilities to your pipeline/production.

### Achieving prevention

This can be achieved a number of key ways:&#x20;

* [Snyk Learn](https://learn.snyk.io/) teaches developers how to stay secure with interactive lessons exploring vulnerabilities across a variety of languages and ecosystems.&#x20;
* [Snyk Advisor](https://snyk.io/advisor/) helps developers by informing them and helping choose better open source packages and containers to base their development work on.&#x20;
* [Snyk for IDEs](../../../ide-tools/) and [Snyk CLI](../../../snyk-cli/) allow developers to test their code locally as they work on it, before they submit it.&#x20;
* [Pull Request](https://docs.snyk.io/integrations/git-repository-scm-integrations/snyk-checks-on-pull-requests) checks: turn these on your most important projects to start, this gives visibility without developers having to do anything extra. As adoption grows, add PR Checks on all your projects.
* [Snyk CI/CD integration](../../../integrations/ci-cd-integrations/): after using [snyk monitor](https://docs.snyk.io/snyk-cli/commands/monitor) for visibility, move to using [snyk test](https://docs.snyk.io/snyk-cli/commands/test) to prevent vulnerabilities from being introduced. Note: we recommend initially configuring your CI pipelines so that [snyk test](https://docs.snyk.io/snyk-cli/commands/test) does not block your builds. You can adjust this later when you are comfortable with the configuration.
  * Focus on Critical, High, and other criteria, like is it fixable, as needed. This is achieved via parameters you can pass to the CLI & CI/CD plugins.

{% hint style="info" %}
You must define an exception path if a Pull Request (PR) is blocked, you have a designated “Snyk admin” who can “mark as successful” in the Snyk PR check results if the business/security decides to allow that PR to move forward.
{% endhint %}

### What's next?

You can now start [addressing the backlog](addressing-the-backlog.md).
