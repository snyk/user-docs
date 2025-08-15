# Phase 6: Rolling out the prevention stage

After you gain visibility for your business-wide security issues, you can start to implement a prevention and gating system to stop new vulnerabilities from being added to your applications.

## Common prevention methods

The two functions that allow you to ‘prevent new issues are:

* Snyk tests on [Pull Request (PR) or Merge Request (MR) Checks](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/),  available for Open Source, with Snyk Code PR checks.
* Adding `snyk test` to your CI/CD pipelines; you may have already implemented `snyk monitor` to import your Projects as part of the pipeline. In addition, open source, code, infrastructure as code, and container vulnerabilities can all be gated.

Regardless of which function you use, Snyk suggests that you communicate these changes clearly to your developers before implementing any form of gating.

## Tips on prevention

### Block the differences

If you are not already blocking vulnerabilities, start by blocking the differences.

This eases developers into the process, as they take responsibility only for vulnerabilities directly related to their changes.&#x20;

* [Run PR Checks](../../../scan-with-snyk/pull-requests/pull-request-checks/), which include the option to block new issues.&#x20;
* [Snyk Delta](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-delta.md) can assist with differences in results at the CLI level.

### Communicate exception processes

It is important to ensure the teams know the exception processes and how to address exceptions if a PR is blocked or a build fails.&#x20;

For example:

* Let the team know who has the authority to override a PR check if a pass is mandatory.
* If a build fails, determine whether the issue can be ignored and the test rerun. Who can run it? Or can a script be run allowing that step to pass? Who can make that determination?

