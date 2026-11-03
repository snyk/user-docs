---
description: 'Phase 5 of the Snyk team implementation: roll out prevention and gating to stop new vulnerabilities entering your applications'
nav_context: agnostic
---

# Phase 5: Rolling out the prevention stage

After you gain visibility into your security issues, you can start to implement a prevention and gating system to stop new vulnerabilities from being added to your applications.

## Common prevention methods

The following are the two common ways to prevent new issues:

* Snyk tests on pull request (PR) and merge request (MR) checks, available for Snyk Open Source and Snyk Code.
* Adding `snyk test` to your CI/CD pipelines (you may have already implemented `snyk monitor` to import your Projects as part of the pipeline).
  * You can gate open source, code, infrastructure as code, and container vulnerabilities.

In either case, Snyk recommends that you communicate these changes clearly to your developers before implementing any form of gating.

## Tips

### Block the differences

If you are not already blocking vulnerabilities, start by blocking the difference.

This eases developers into the process, as they only take responsibility for vulnerabilities directly related to their changes.

To block new issues, you can use PR Checks. For details, visit [Run PR Checks](https://docs.snyk.io/scan-fix-and-prevent/prevent/pull-request-checks).

### Communicate exception processes

Ensure that the teams know the exception processes and what to do if a PR is blocked or a build fails.

For example:

* Let the team know who has the authority to override a PR check if a pass is mandatory.
* If a build fails, can the issue be ignored and the test rerun? Who can run it? Or can a script be run allowing that step to pass? Who can make that determination?
