# Understanding the Severity Score of the Snyk Code Issues

## Methodology

Snyk Code currently reports issues by severity with three dimensions:

* High
* Medium
* Low

At a high level, we determine severity based on the following factors:

* Qualities intrinsic to a vulnerability
* Evolution of vulnerability over a lifetime

Additionally, we incorporate a variety of quantitative and qualitative factors:

### Quantitative:

* Severity scores from other SAST products where information is publicly available
* Severity scores from identifying similar vulnerabilities in Snyk VulnDB (our proprietary Snyk Vulnerability database)

### Qualitative:

* The severity of source (direct vs. indirect)
* Prevalence and Impact of sink
* Security Team Experience and Research
* Customer Feedback

{% hint style="info" %}
**Exceptions**

If a vulnerability is detected in code, filename, or folder with the word `test`, it will be deemed a low-severity vulnerability. This applies to all languages.



The severity of CWEs may change depending on the environment.
{% endhint %}

### Examples

### CWE-22: Path Traversal

<figure><img src="../../../.gitbook/assets/image (1).png" alt="Severity score for a path traversal."><figcaption><p>Severity score for a path traversal</p></figcaption></figure>

### CWE-601: Open Redirect

<figure><img src="../../../.gitbook/assets/image.png" alt="Severity score for an open redirect."><figcaption><p>Severity score for an open redirect</p></figcaption></figure>

