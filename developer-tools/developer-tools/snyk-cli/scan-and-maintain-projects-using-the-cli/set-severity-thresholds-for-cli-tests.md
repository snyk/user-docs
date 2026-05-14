# Severity thresholds for CLI tests

To improve control over your tests, you can use the `--severity-threshold` option for the `snyk test` command with one of the supported levels: `low|medium|high|critical`. With this option, only vulnerabilities of **the specified level or higher** are reported.

`$ snyk test --severity-threshold=medium`

{% hint style="info" %}
Setting `--severity-threshold` to `low` has the same effect as running the command without specifying the threshold; all vulnerabilities are reported.
{% endhint %}

Note: The `--severity-threshold` option is available with the `snyk test`, `snyk code`, `snyk container`, and `snyk iac test` commands. See the [CLI commands help](../commands/) pages for each command for details.
