# Set severity thresholds for CLI tests

To have better control over your tests, you can pass the severity-threshold flag to the **snyk test** command with one of the supported options \(low\|medium\|high\|critical\). With this flag, only vulnerabilities of **provided level or higher** will be reported.

```text
$ snyk test --severity-threshold=medium
```

{% hint style="info" %}
**NOTE**

Setting severity threshold to _Low_ currently has the same effect as running without specifying the threshold at all: all vulnerabilities are reported.
{% endhint %}

See also

* [Getting started with the CLI](getting-started-with-the-cli.md)
* [Our full CLI reference](https://support.snyk.io/hc/articles/360003812578#UUID-c88e66cf-431c-9ab1-d388-a8f82991c6e0)
* [Test your code from the CLI](https://support.snyk.io/hc/articles/360003812478#UUID-2e8464f9-1d8a-5f79-466d-2ca5c5cf9f22)
* [Test dev dependencies](https://support.snyk.io/hc/articles/360003812478#UUID-1070ae3e-5f30-cb4e-6350-a1c3f5c67511)
* [Test public repositories before use](https://support.snyk.io/hc/articles/360003851277#UUID-ba99a73f-110d-1f1d-9e7a-1bad66bf0996)
* [Test public npm packages before use](https://support.snyk.io/hc/articles/360003812498#UUID-0ab434a8-c1b5-873d-cbf6-7a61a07c4ad8)
* [Manage vulnerability results with the Snyk CLI wizard](https://support.snyk.io/hc/articles/360003851357#UUID-b401cc8a-a55a-2b74-d9e5-c92dd49ed58c)

