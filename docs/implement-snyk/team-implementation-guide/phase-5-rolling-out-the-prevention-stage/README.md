# Phase 5: Rolling out the prevention stage

After you gain visibility on your security issues, you can now start to implement a prevention/gating system, to stop new vulnerabilities being added to your applications.

## Common prevention methods

Below are the two common areas that allow you to "prevent new issues":&#x20;

* Snyk tests on pull request (PR)/merge checks (MR),  currently available for Open Source.&#x20;
* Adding "Snyk test" to your CI/CD pipelines (you may have already implemented "Snyk monitor" to import your Projects as part of the pipeline).&#x20;
  * Additionally, open source, code, infrastructure as code, and container vulnerabilities can all be gated.

In either case, Snyk suggests that you communicate these changes clearly to your developers before implementing any form of gating.

## Tips

### Block the differences

If you are not already blocking vulnerabilities, start by blocking the difference.

This eases developers into the process, as they only take responsibility for vulnerabilities directly related to their changes.&#x20;

To block new issues, you can use PR checks. For details, see [Run PR Checks](../../../scan-with-snyk/pull-requests/pull-request-checks/).

### Communicate exception processes

It is important to ensure the teams know the exception processes and how to address if a PR is blocked or a build fails.&#x20;

For example:

* Let the team know who has the authority to override a PR check if a pass is mandatory.
* If a build fails, can the issue be ignored and the test rerun? Who can run it? Or can a script be run allowing that step to pass? Who can make that determination?

