# Framework mappings

Every attacker goal Snyk scores is cross-referenced to four industry frameworks, so you can locate a model risk finding inside a control framework you already use instead of treating it as a separate vocabulary.

Use this page to map a finding to the control, control family, or risk category your governance process already tracks.

## The frameworks

* MITRE ATLAS, version 5.6.0. Adversarial tactics and techniques against AI systems. Identifiers take the form `AML.Txxxx`.
* OWASP Top 10 for Large Language Model Applications, 2025. Identifiers take the form `LLMxx`.
* OWASP Top 10 for Agentic Security Initiative (ASI), 2026. Identifiers take the form `ASIxx`.
* NIST AI Risk Management Framework, covering AI 100-2 E2025 for adversarial machine learning and AI 600-1 for generative AI risks.

## How the mappings relate

Two properties of the mapping explain most of the gaps you see in the tables.

Every indirect attacker goal shares indirect prompt injection as its delivery vector, so every indirect row also maps to `AML.T0051.001`, `LLM01`, and `NISTAML.015`. The indirect table lists only the additional identifiers specific to each goal.

The OWASP Top 10 for LLM Applications has no harmful-content category, because it is a security taxonomy rather than a content-safety one. Where a harmful-content goal lists `LLM01`, that identifier describes the jailbreak or injection used to bypass the model's guardrails, not the goal itself. Those rows are marked with an asterisk. The goal itself maps to NIST AI 600-1 and to the MITRE ATLAS external-harms and jailbreak techniques.

{% hint style="info" %}
Insecure code generation is intentionally unmapped to MITRE ATLAS. No adversary operates against the model, so no ATLAS technique applies.
{% endhint %}

## Direct attack surface

On the direct surface, the user is the adversary and controls the prompt.

<table><thead><tr><th width="151.6171875">Attacker goal</th><th width="147.46484375">MITRE ATLAS</th><th width="119.890625">OWASP LLM</th><th width="100.0859375">OWASP ASI</th><th>NIST AI 100-2 / 600-1</th></tr></thead><tbody><tr><td>Chemical &#x26; biological</td><td>AML.T0054, AML.T0048</td><td>LLM01*</td><td>None</td><td>NISTAML.04, AI 600-1 #1</td></tr><tr><td>Cybercrime &#x26; intrusion</td><td>AML.T0054, AML.T0048</td><td>LLM01*</td><td>None</td><td>NISTAML.04, AI 600-1 #9</td></tr><tr><td>Harassment &#x26; bullying</td><td>AML.T0054, AML.T0048.003</td><td>LLM01*</td><td>None</td><td>NISTAML.04, AI 600-1 #3</td></tr><tr><td>Harmful (general)</td><td>AML.T0054, AML.T0048</td><td>LLM01*</td><td>None</td><td>NISTAML.04, AI 600-1 #3</td></tr><tr><td>Illegal activities</td><td>AML.T0054, AML.T0048</td><td>LLM01*</td><td>None</td><td>NISTAML.04, AI 600-1 #3</td></tr><tr><td>Misinformation &#x26; disinformation</td><td>AML.T0048.002</td><td>LLM01*, LLM09</td><td>None</td><td>NISTAML.04, AI 600-1 #8</td></tr><tr><td>Copyright reproduction</td><td>AML.T0048.004</td><td>LLM02</td><td>None</td><td>AI 600-1 #10</td></tr><tr><td>Gender bias</td><td>AML.T0048.002</td><td>None</td><td>None</td><td>AI 600-1 #6</td></tr><tr><td>Race bias</td><td>AML.T0048.002</td><td>None</td><td>None</td><td>AI 600-1 #6</td></tr><tr><td>Religion bias</td><td>AML.T0048.002</td><td>None</td><td>None</td><td>AI 600-1 #6</td></tr><tr><td>Insecure code generation</td><td>None</td><td>LLM05</td><td>ASI05</td><td>NISTAML.027</td></tr><tr><td>System prompt extraction</td><td>AML.T0056, AML.T0069.002</td><td>LLM07</td><td>None</td><td>NISTAML.035</td></tr><tr><td>Tool extraction</td><td>AML.T0084, AML.T0069</td><td>LLM07</td><td>None</td><td>NISTAML.035</td></tr><tr><td>Model identification</td><td>AML.T0014, AML.T0069</td><td>LLM07</td><td>None</td><td>NISTAML.035</td></tr><tr><td>PII extraction</td><td>AML.T0057</td><td>LLM02</td><td>None</td><td>NISTAML.038</td></tr><tr><td>Role / purpose deviation</td><td>AML.T0054, AML.T0051</td><td>LLM01</td><td>ASI01</td><td>NISTAML.018</td></tr><tr><td>Unauthorized tool use</td><td>AML.T0053</td><td>LLM06</td><td>ASI02, ASI03</td><td>None</td></tr></tbody></table>

\* `LLM01` describes the elicitation vector rather than the goal. See the preceding section.

## Indirect attack surface

On the indirect surface, the attacker hides instructions in content the model reads while doing its job, and the user is the victim. Every goal in the following table also maps to `AML.T0051.001`, `LLM01`, and `NISTAML.015`.

<table><thead><tr><th width="151.65234375">Attacker goal</th><th width="146.9609375">MITRE ATLAS</th><th width="120.36328125">OWASP LLM</th><th width="99.98828125">OWASP ASI</th><th>NIST AI 100-2 / 600-1</th></tr></thead><tbody><tr><td>Credential / secret theft</td><td>AML.T0098, AML.T0086</td><td>LLM02</td><td>ASI03</td><td>NISTAML.039</td></tr><tr><td>System prompt extraction</td><td>AML.T0056</td><td>LLM07</td><td>None</td><td>NISTAML.035</td></tr><tr><td>Source code exfiltration</td><td>AML.T0086</td><td>LLM02</td><td>None</td><td>NISTAML.039</td></tr><tr><td>Private data leakage</td><td>AML.T0057, AML.T0086</td><td>LLM02</td><td>None</td><td>NISTAML.036</td></tr><tr><td>Command execution</td><td>AML.T0050, AML.T0053, AML.T0102</td><td>LLM05</td><td>ASI05</td><td>None</td></tr><tr><td>Data / file destruction</td><td>AML.T0101</td><td>LLM06</td><td>ASI02</td><td>None</td></tr><tr><td>Other unauthorized actions</td><td>AML.T0053</td><td>LLM06</td><td>ASI02</td><td>None</td></tr><tr><td>Code backdoors</td><td>AML.T0102</td><td>LLM05</td><td>ASI05, ASI09</td><td>NISTAML.027</td></tr><tr><td>Supply chain poisoning</td><td>AML.T0010, AML.T0011.001</td><td>LLM03</td><td>ASI04</td><td>None</td></tr><tr><td>Security control degradation</td><td>AML.T0031, AML.T0053</td><td>None</td><td>ASI02, ASI09</td><td>None</td></tr><tr><td>Decision override</td><td>AML.T0031, AML.T0067</td><td>None</td><td>ASI09</td><td>NISTAML.027</td></tr><tr><td>Phishing / credential harvesting</td><td>AML.T0052.000, AML.T0067</td><td>None</td><td>ASI09</td><td>None</td></tr><tr><td>Trust exploitation</td><td>AML.T0031</td><td>LLM09</td><td>None</td><td>NISTAML.027</td></tr></tbody></table>

## Related

To learn how Snyk scores each attacker goal and how the scores roll up, visit [Risk intelligence](./).

To learn how to raise an issue when a model crosses a score you choose, visit [Policies & issues](../../platform-surfaces/policies-and-issues.md).
