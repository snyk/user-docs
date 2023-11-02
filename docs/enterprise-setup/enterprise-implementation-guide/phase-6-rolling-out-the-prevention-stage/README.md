# Phase 6: Rolling out the prevention stage

After you gain visibility on your business-wide security issues, you can now start to implement a prevention/gating system, to stop new vulnerabilities being added to your applications.

## Common prevention methods

The two common areas that allow you to ‘prevent new issues’ are&#x20;

* Snyk tests on pull request (PR)/merge checks (MR),  currently available for Open Source.&#x20;
  * Snyk Code PR checks are currently in Beta.
* Adding ‘Snyk test’ to your CI/CD pipelines (you may have already implemented ‘Snyk monitor’ to import your projects as part of the pipeline).&#x20;
  * Additionally, open source, code, infrastructure as code and container vulnerabilities can all be gated.

In either case, we suggest that you communicate these changes clearly to your developers before implementing any form of gating.

## Tips

### Block the differences

If you are not already blocking vulnerabilities, start by blocking the difference.

This eases developers into the process, as they only take responsibility for vulnerabilities directly related to their changes.&#x20;

* [Run PR Checks](../../../scan-with-snyk/run-pr-checks/): these have the option to block new issues.&#x20;
* [Snyk Delta](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-delta.md) can assist with differences in results at the CLI level.

### Communicate exception processes

It is important to ensure the teams know the exception processes, and how to address if a PR is blocked or a build fails.&#x20;

For example:

* Let the team know who has the authority to override a PR check if a pass is mandatory
* If a build fails, can the issue be ignored and the test rerun? Who can run it? Or can a script be run allowing that step to pass? Who can make that determination?

