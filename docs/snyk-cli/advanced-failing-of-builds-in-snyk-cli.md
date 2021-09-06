# Advanced failing of builds in Snyk CLI

The Snyk CLI provides the following options when failing your builds:

```text
--severity-threshold=low|medium|high
```

Only report vulnerabilities of provided level or higher.

```text
--fail-on=all|upgradable|patchable
```

Only fail when there are vulnerabilities that can be fixed.

```text
--fail-on=all
```

fails when there is at least one vulnerability that can be either upgraded or patched.

```text
--fail-on=upgradable
```

fails when there is at least one vulnerability that can be upgraded.

```text
--fail-on=patchable
```

fails when there is at least one vulnerability that can be patched. If vulnerabilities do not have a fix and this option is being used, tests will pass.

The Snyk CLI on its own does not currently have the capability natively to fail tests on more complex use cases. Here are some ways to achieve more complex pass/fail criteria.

## Combining security policies with --severity-threshold

[Security policies](https://snyk.gitbook.io/user-docs/fixing-and-prioritizing-issues/policies) provide the capability to change a vulnerability's severity if it matches specific criteria when a project is tested against an organization using that policy. You could, for example, change the severity of a vulnerability from high to low, and if performing a snyk test with the CLI with

```text
 --severity-threshold=medium|high
```

this previously high severity vulnerability will no longer fail the build.

{% hint style="info" %}
Security policies do not have all attributes available for criteria matching. Please refer to the security policy configuration for what's available as it will be added to over time.
{% endhint %}

Here is an example of a snyk test, using --severity-threshold=high running against a default organization with no policy applied to it.

![](https://gblobscdn.gitbook.com/assets%2F-MVXKdrh-jY3KDGPs8lQ%2F-MZT_W3O1oFyMAzF9g3s%2F-MZTrc0D6NjT6VlS1jmU%2Fimage.png?alt=media&token=27e0ee8c-147f-4942-ada4-08de07f67c40)

Here is an example of a `snyk test`, using `--severity-threshold=high`, running against an organization with a policy that downgrades this particular vulnerability severity to _low_

![](https://gblobscdn.gitbook.com/assets%2F-MVXKdrh-jY3KDGPs8lQ%2F-MZT_W3O1oFyMAzF9g3s%2F-MZTuPF3Uat7DSSnTKFD%2Fimage.png?alt=media&token=914fd76f-bd9f-4170-8d96-b32026ae19ee)  
Since we lowered the severity of the original vulnerability with the policy, it no longer breached the severity threshold of high, resulting in a passing test.

## Companion tools

The rest of this article discusses use cases using snyk-delta or snyk-filter, which are open source companion tools for the Snyk CLI.

snyk-delta finds the delta of vulnerabilities between the current test and a previously monitored snapshot.

It is available from npmjs.org, and may be pulled into your CI/CD pipeline using

```text
npm install -g snyk-delta
```

snyk-filter provides for user-defined pass/fail criteria based on any available data in the snyk test JSON output.

It is available from npmjs.org and may be pulled into your CI/CD pipeline using npm install

```text
npm install -g snyk-filter
```

## Fail current build only if new vulnerabilities are being introduced

## Inline mode

```text
snyk test --json --print-deps | snyk-delta
```

Possibly point to a specific snapshot by specifying org + project coordinates

```text
snyk test --json --print-deps | snyk-delta --baselineOrg xxx --baselineProject xxx
```

## Standalone

```text
snyk-delta --baselineOrg xxx --baselineProject xxx --currentOrg xxx --currentProject xxx
```

Please refer to the snyk-delta project on GitHub for more information.

## Fail build for CVSS score higher than ...

```text
snyk test --json | snyk-filter -f /path/to/example-cvss-9-or-above.yml
```

## Custom criteria and filtering

snyk-filter can utilize any combination of criteria available in the snyk test JSON output.

You may also have different criteria for display from what will fail the build. This allows you to do things like display all vulns in the test output, while failing only on some specific criteria.

{% hint style="info" %}
Examples are provided in the snyk-filter project on Github [here](https://github.com/snyk-tech-services/snyk-filter). Please refer to the [snyk-filter project on GitHub](https://github.com/snyk-tech-services/snyk-filter) for more information.
{% endhint %}

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

