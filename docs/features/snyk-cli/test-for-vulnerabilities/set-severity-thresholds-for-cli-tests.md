# Set severity thresholds for CLI tests

To improve control over your tests, you can use the severity-threshold option for the `snyk test` **** command with one of the supported options (low|medium|high|critical). With this option, only vulnerabilities of **the specified level or higher** are reported.

`$ snyk test --severity-threshold=medium`

{% hint style="info" %}
Setting severity threshold to `low` currently has the same effect as running the command without specifying the threshold; all vulnerabilities are reported.
{% endhint %}
