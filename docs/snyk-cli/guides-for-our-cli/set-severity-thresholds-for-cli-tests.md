# Set severity thresholds for CLI tests

To have better control over your tests, you can pass the severity-threshold flag to the **snyk test** command with one of the supported options \(low\|medium\|high\|critical\). With this flag, only vulnerabilities of **provided level or higher** will be reported.

```text
$ snyk test --severity-threshold=medium
```

{% hint style="info" %}
Setting severity threshold to **Low** currently has the same effect as running without specifying the threshold at all: all vulnerabilities are reported.
{% endhint %}



