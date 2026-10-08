---
description: How Snyk Code security rules work and how they are updated
nav_context: agnostic
---

# Snyk Code security rules

{% hint style="info" %}
Snyk Code rules are updated continuously. The list expands continually, and the rules may change to provide the best protection and security solutions for your code.
{% endhint %}

This page lists all security rules used by Snyk Code when scanning your source code for vulnerabilities.

Each rule includes the following information.

* **Rule Name**: The Snyk name of the rule.
* **Languages**: The programming languages to which this specific rule applies. Note that there might be two rules with the same name that apply to different languages.
* **CWEs**: The [CWE numbers](https://cwe.mitre.org/) the rule covers.
* **Security Categories**: The [OWASP Top 10](https://owasp.org/Top10/2025/) (2025 edition) category the rule maps to, when applicable. This column also notes whether the rule appears in the [CWE Top 25](https://cwe.mitre.org/top25/), and any applicable [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) (2023) or [OWASP Mobile Top 10](https://owasp.org/www-project-mobile-top-10/) (2024) categories.

{% hint style="info" %}
\* XML listed in the language column applies only to NuGet XML files.&#x20;
{% endhint %}

## Snyk Code secrets detection

Snyk Code secrets detection is part of Snyk Code SAST. It is a separate capability from [Snyk Secrets](../../snyk-secrets/README.md), which is a dedicated secrets scanner with its own engine and findings.

The Snyk Code rules report hardcoded secrets and credentials in these cases:

* A hardcoded user name, password, or key assigned to a variable, field, or argument whose name indicates a credential, for example `username`, `password`, or `key`.
* A secret with a known format, such as an AWS or GitHub key, matched by a regular expression.
* A hardcoded string passed as a credential argument to a method, for example the user name and password arguments of a database connection call.

The rules look at the name and the shape of the code, and they check the value, for example its length. They do not verify that a value is a working secret. A value read from an environment variable or a configuration store is not reported, because the code holds no literal.

Snyk Code reads source code files in supported languages and does not read comments. Findings appear as **Hardcoded Secret** or **Use of Hardcoded Credentials** issues, with CWE-259, CWE-321, CWE-547, or CWE-798, depending on the rule. For the rules in each language, see the language pages in this section.

To scan all plain text files for secrets, use [Snyk Secrets](../../snyk-secrets/README.md). To suppress a Snyk Code finding, see [Consistent Ignores for Snyk Code](../../../manage-risk/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code/README.md).
