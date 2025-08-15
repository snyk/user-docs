# Failing of builds in Snyk CLI

The `snyk test` command has the following options for failing your builds:

`--severity-threshold=low|medium|high|critical` Report only vulnerabilities of the specified level or higher.

`--fail-on=all` Fail when there is at least one vulnerability that can be either upgraded or patched.

`--fail-on=upgradable` Fail when there is at least one vulnerability that can be upgraded.

`--fail-on=patchable` Fail when there is at least one vulnerability that can be patched or vulnerabilities that can be either patched or upgraded.

The Snyk CLI on its own does not have the capability natively to fail tests on more complex use cases. Here are some ways to achieve more complex pass/fail criteria.

## Combining security policies with --severity-threshold

[Security policies](../../../manage-risk/policies/) provide the capability to change the severity of a vulnerability, when a Project is tested against an Organization using that policy and the severity matches specific criteria. You could, for example, change the severity of a vulnerability from high to low, and if you run `snyk test` with the option `--severity-threshold=medium|high`, this previously high severity vulnerability no longer fails the build.

{% hint style="info" %}
Security policies do not have all attributes available for criteria matching. Refer to the security policy configuration to see what is available.
{% endhint %}

An example follows of `snyk test` with the option `--severity-threshold=high` running against a default Organization with no policy applied to it.

![Test against a default organization with no policy applied to it](https://gblobscdn.gitbook.com/assets%2F-MVXKdrh-jY3KDGPs8lQ%2F-MZT_W3O1oFyMAzF9g3s%2F-MZTrc0D6NjT6VlS1jmU%2Fimage.png?alt=media\&token=27e0ee8c-147f-4942-ada4-08de07f67c40)

The following example shows the`snyk test` command with the option `--severity-threshold=high` running against an Organization with a policy that downgrades this particular vulnerability severity to `low`. There are no vulnerabilities found.

![Test against an organization with a policy applied](../../../.gitbook/assets/test-organization-with-policy-applied.png)

## Companion tools

The CLI has companion tools for Open Source scanning  `snyk-delta` and `snyk-filter`.

`snyk-delta` finds the delta of vulnerabilities between the current test and a previously monitored snapshot. This tool is available from npmjs.org, and may be pulled into your CI/CD pipeline by running `npm install -g snyk-delta`.

`snyk-filter` provides for user-defined pass/fail criteria based on any available data in the `snyk test` JSON output. This tool is available from npmjs.org and may be pulled into your CI/CD pipeline by running `npm install -g snyk-filter`.

### Fail current build only if new vulnerabilities are being introduced

#### inline mode for `snyk delta`

This is a simple example: `snyk test --json --print-deps | snyk-delta`

This example possibly points to a specific snapshot by specifying Organization and Project coordinates.

`snyk test --json --print-deps | snyk-delta --baselineOrg xxx --baselineProject xxx`

#### Standalone `snyk delta`

`snyk-delta --baselineOrg xxx --baselineProject xxx --currentOrg xxx --currentProject xxx`

Refer to the [snyk-delta project on GitHub](https://github.com/snyk-tech-services/snyk-delta) for more information.

### Fail build for CVSS score higher than a specified score with `snyk filter`

`snyk test --json | snyk-filter -f /path/to/example-cvss-9-or-above.yml`

### Custom criteria and filtering with `snyk filter`

`snyk-filter` can use any combination of criteria available in the `snyk test` JSON output.

You may also have different criteria for display from what will fail the build. This allows you to do things like display all vulnerabilities in the test output while failing only on some specific criteria.

Refer to the [snyk-filter project on GitHub](https://github.com/snyk-tech-services/snyk-filter) for examples and more information.
