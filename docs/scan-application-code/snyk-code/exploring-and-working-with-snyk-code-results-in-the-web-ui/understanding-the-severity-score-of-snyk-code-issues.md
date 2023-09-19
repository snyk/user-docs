# Understanding the Severity Score of Snyk Code Issues

Snyk Code reports issues by severity levels: High, Medium, and Low. Snyk Code currently does not use the **Critical** severity level.

At the High level, Snyk determines severity based on the following factors:

* Qualities intrinsic to a vulnerability
* Evolution of vulnerability over a lifetime

In addition, Snyk incorporates a variety of quantitative and qualitative factors:

* Quantitative factors in Priority Score
  * Severity scores from other SAST products where information is publicly available
  * Severity scores from identifying similar vulnerabilities in the Snyk Vulnerability database
* Qualitative factors in Priority Score
  * The severity of source, direct versus indirect
  * Prevalence and impact of the sink
  * Security team experience and research
  * Customer feedback

{% hint style="info" %}
**Exceptions**

If a vulnerability is detected in code, filename, or folder with the word `test`, it is deemed a low-severity vulnerability. This applies to all languages.



The severity of CWEs may change depending on the environment.
{% endhint %}

## Example: CWE-22: Path Traversal

For CEW-22 Path Traversal, if the vulnerability occurs in a test, it is Low severity. If not, and it comes from a direct source, it is High severity. Otherwise, it is Low severity.

<figure><img src="../../../.gitbook/assets/image (81).png" alt="Decision flow chart for Priority Score CWE-22 Path Traversal"><figcaption><p>Decision flow chart for Priority Score CWE-22 Path Traversal</p></figcaption></figure>

## Example: CWE-601: Open Redirect

For CEW-2601 Open Redirect, if the vulnerability occurs in a test, it is Low severity. If not, and it comes from a direct source, it is Medium severity.

<figure><img src="../../../.gitbook/assets/image (5) (8).png" alt="Decision flow chart for Priority Score CWE-601 Oen Redirect"><figcaption><p>Decision flow chart for Priority Score CWE-601 Oen Redirect</p></figcaption></figure>

